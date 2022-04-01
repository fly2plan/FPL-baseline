const express = require('express');
const { FPLSyncProof, ACKSyncProof } = require('./privacy/proof');
const { SYNC, updateSyncWorkflow, getWorkflow } = require('./workflowRegistry');
const { workflowSyncFPLEventType } = require('./messaging/eventType');
const router = express.Router();
const { v4: uuidv4 } = require('uuid');
const { addWorkstep, getWorkstep } = require('./workstepRegistry');
const { getShieldAddress } = require('./workgroupRegistry');
const FPLMessageProducer = require('./messaging/producer');
const truffle_connect = require('../connection/truffle_connect');
const { getFormattedProof } = require('./privacy/service');

router.use((req, res, next) => {
    if (res.locals.id !== undefined) {
        if (req.path !== '/prove/fpl') {
            return res.sendStatus(403);
        };
    };
    next();
});

router.use('/hash/:type', (req, res) => {
    const pk = req.body.pk;
    const workgroupId = req.body.workgroupId;
    const type = req.params.type;
    const workstepId = uuidv4();
    if (type == 'fpl') {
        const workflowId = res.locals.id;
        const fpl = req.body.fpl;
        const fplSyncProof = new FPLSyncProof(fpl, pk); //instantiate fplSyncProof which calls private methods to fill its fields;

        // Store fplSyncProof as workflow step (workstep).
        addWorkstep({ id: workstepId, step: fplSyncProof });
        updateSyncWorkflow(workgroupId, workflowId, workstepId);

        res.status(200).send({ workflowId, workstepId, artifacts: fplSyncProof.getProofArtifacts() });
    } else if (type == 'ack') {
        const workflowId = req.body.workflowId;
        const workflow = getWorkflow(workflowId);
        const fplSyncProof = getWorkstep(workflow['SYNC'][0]);
        const ackSyncProof = new ACKSyncProof(fplSyncProof, pk);

        addWorkstep({ id: workstepId, step: ackSyncProof });
        updateSyncWorkflow(workgroupId, workflowId, workstepId);

        res.status(200).send({ workflowId, workstepId, artifacts: ackSyncProof.getProofArtifacts() });
    } else {
        return res.status(404).send('Unknown object type');
    };
});

router.use('/prove/:type', async (req, res) => {
    const workgroupId = req.body.workgroupId;
    const workflowId = res.locals.id ? res.locals.id : req.body.workflowId;
    const workstepId = uuidv4();
    const type = req.params.type;
    let syncProof = null;
    let proof = null;
    if (type == 'fpl') {
        const fpl = req.body.fpl
        const sk = req.body.sk;
        syncProof = new FPLSyncProof(fpl)
        proof = await syncProof.createProof(sk);
        addWorkstep({ id: workstepId, step: syncProof });
    } else if (type == 'ack') {
        const sk = req.body.sk
        const workflow = getWorkflow(workflowId)
        const workstep = getWorkstep(workflow[SYNC][workflow[SYNC].length - 1]).getProofAndInputs()
        const { raw, inputs } = workstep
        const {pk, sig, hash} = inputs
        syncProof = new ACKSyncProof({fpl: raw, pk, sig, hash})
        proof = await syncProof.createProof(sk)
    } else {
        return res.status(404).send('Unknown object type');
    };
    if (proof !== null) {
        const eventType = type == 'fpl' ? workflowSyncFPLEventType: workflowSyncACKEventType
        const event = type == 'fpl' ? 'workflowSyncFPL' : 'workflowSyncACK'
        const producer = new FPLMessageProducer(event, eventType);
        await producer.queue({ workgroupId, workflowId, workstepId, [SYNC]: { [type]: syncProof.getProofAndInputs() } }, eventType);
    } else {
        return res.status(405).send('Proof generation failed');
    };
    res.sendStatus(200);
});

router.use('/verify/:proof', async (req, res) => {
    const workgroupId = req.body.workgroupId
    const workflowId = req.body.workflowId
    const workflow = getWorkflow(workflowId)
    if (workflow === undefined) {
        return res.status(404).send('Workflow not found');
    }
    const workstep = getWorkstep(workflow[SYNC][workflow[SYNC].length - 1]).getProofAndInputs()
    if (workstep === undefined) {
        return res.status(404).send('Workflow is empty');
    }
    const { proof, inputs } = workstep
    const shieldAddress = getShieldAddress(workgroupId);
    const type = req.params.proof;
    if (type === 'fpl' || type === 'ack') {
        let formattedProof = await getFormattedProof(type, proof, inputs)
        if (formattedProof[1] !== undefined) {
            return res.status(500).send(`Error formatting proof: ${formattedProof[1]}`)
        }
        const {a, b, c} = formattedProof[0].proof
        const input = formattedProof[0].inputs
        await truffle_connect.verify(type, a, b, c, input, shieldAddress, (result) => {
            return res.send(result);
        });
    } else {
        return res.status(404).send('Unknown proof type');
    };
})

module.exports = {
    workflowSyncRouter: router,
}
