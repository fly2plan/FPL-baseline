const express = require('express');
const { FPLSyncProof, ACKSyncProof } = require('./privacy/proof');
const { SYNC, updateSyncWorkflow, getWorkflow } = require('./workflowRegistry');
const { workflowSyncEventType } = require('./messaging/eventType');
const router = express.Router();
const { v4: uuidv4 } = require('uuid');
const { addWorkstep, getWorkstep } = require('./workstepRegistry');
const { getShieldAddress } = require('./workgroupRegistry');
const FPLMessageProducer = require('./messaging/producer');

router.use((req, res, next) => {
    if (res.locals.id !== undefined) {
        if (req.path !== '/hash/fpl') {
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

        // const [h1, h2] = fplSyncProof.hash;
        // const proof = await generateProof(h1, h2);

        // const producer = new KafkaProducer('workflowSync', workflowSyncEventType);
        // await producer.queue({ workgroupId, id, [SYNC]: { fpl: { raw, hash } } }, workflowSyncEventType);

        // Store fplSyncProof as workflow step (workstep) then access again.

        addWorkstep({ id: workstepId, step: fplSyncProof });
        updateSyncWorkflow(workgroupId, workflowId, workstepId);

        res.status(200).send({ workflowId, workstepId, artifacts: fplSyncProof.getProofArtifacts() });
    } else if (type == 'ack') {
        const workflowId = req.body.workflowId;
        const workflow = getWorkflow(workflowId);
        const fplSyncProof = getWorkstep(workflow['SYNC'][0]);
        const ackSyncProof = new ACKSyncProof(fplSyncProof, pk);
        const workstepId = uuidv4();

        // const [h1, h2] = fplSyncProof.hash;
        // const proof = await generateProof(h1, h2);

        // const producer = new KafkaProducer('workflowSync', workflowSyncEventType);
        // await producer.queue({ workgroupId, id, [SYNC]: { fpl: { raw, hash } } }, workflowSyncEventType);

        // Store fplSyncProof as workflow step (workstep) then access again.
        addWorkstep({ id: workstepId, step: ackSyncProof });
        updateSyncWorkflow(workgroupId, workflowId, workstepId);

        res.status(200).send({ workflowId, workstepId, artifacts: ackSyncProof.getProofArtifacts() });
    } else {
        return res.status(404).send('Unknown object type');
    };
});

router.use('/prove/:type', async (req, res) => {
    const workgroupId = req.body.workgroupId;
    const workflowId = req.body.workflowId;
    const workstepId = req.body.workstepId;
    const type = req.params.type;
    let syncProof = null;
    let proof = null;
    if (type == 'fpl') {
        const sig_ao = req.body.sig_ao;
        syncProof = getWorkstep(workstepId);
        proof = await syncProof.createProof(sig_ao);
    } else if (type == 'ack') {
        const sig_nm = req.body.sig_nm;
        syncProof = getWorkstep(workstepId);
        proof = await syncProof.createProof(sig_nm);
    } else {
        return res.status(404).send('Unknown object type');
    };
    if (proof !== null) {
        // const producer = new KafkaProducer('workflowSync', workflowEventType);
        // await producer.queue({ workgroupId, id, type: SYNC, [SYNC]: { fpl: { raw, proof, hash, signature } } }, workflowEventType);

        // const producer = new KafkaProducer('workflowSync', workflowSyncEventType);
        // await producer.queue({ workgroupId, workflowId, workstepId }, workflowSyncEventType);

        // const producer = new FPLMessageProducer('workflowSync', workflowSyncEventType);
        // await producer.queue({ workgroupId, workflowId, workstepId, [SYNC]: { [type]: syncProof.getProofArtifacts() } }, workflowSyncEventType);
    } else {
        return res.status(405).send('Proof generation failed');
    };
    //res.sendStatus(200);
    res.status(200).send(proof);
});

router.use('/verify/:proof', async (req, res) => {
    const workgroupId = req.body.workgroupId;
    const type = req.params.proof;
    const proof = req.body.proof;
    if (type == 'fpl') {
        fpl = req.body.raw;
        // const fplSyncProof = new fplSyncProof(fpl, null);
        // if (fplSyncProof.getProofArtifacts().hash !== req.body.hash) {
        //     return res.status(405).send('Given fpl hash and generated fpl hash disagree');
        // };
        const shieldAddress = getShieldAddress(workgroupId);
        await truffle_connect.verify(proof.proof, proof.inputs, 'fpl', shieldAddress, () => {
            res.send('verified');
        });
    } else if (type == 'ack') {
        const shieldAddress = getShieldAddress(workgroupId);
        await truffle_connect.verify(proof.proof, proof.inputs, 'ack', shieldAddress, () => {
            res.send('verified');
        });
    } else {
        return res.status(404).send('Unknown proof type');
    };
})

module.exports = {
    workflowSyncRouter: router,
}