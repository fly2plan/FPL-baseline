const { FPLSyncProof, ACKSyncProof } = require("./privacy/proof");
const { getIO } = require("./utils/socket");
const { updateWorkgroup, workgroupRegistry } = require("./workgroupRegistry");
const { getWorkstep, workstepExists } = require("./workstepRegistry");

let workflowRegistry = new Map();

const SYNC = 'SYNC';
const VALIDATE = 'VALIDATE';

const workflowExists = (id) => {
  return workflowRegistry.has(id);
};

const getWorkflow = (id) => {
  return workflowRegistry.get(id);
}

const updateWorkflow = (workflow) => {
  if (!workflowExists(workflow.id)) {
    workflowRegistry.set(workflow.id, {[workflow.type]: [workflow.workstepId]});
    let workgroup = workgroupRegistry.get(workflow.workgroupId);
    console.log(workgroup, workflow);
    workgroup.workflows.push(workflow.id);
    updateWorkgroup(workgroup);
    return;
  };
  if (!workstepExists(workflow.workstepId)) {
    workflowRegistry.set(workflow.id, workflowRegistry.get(workflow.id)[workflow.type].push(workflow.workstepId));
  };
};

const addSyncWorkstep = (workflowId, id, {...content}) => {
  const {fpl, ack, pk} = content;
  let proof = {};
  if (fpl !== undefined) {
    proof = new FPLSyncProof(fpl.raw, pk, fpl.signature, fpl.proof); //instantiate fplSyncProof which calls private methods to fill its fields;
  } else if (ack !== undefined) {
    const fplSyncProof = getWorkstep(getWorkflow(workflowId)['SYNC'][0]);
    proof = new ACKSyncProof(fplSyncProof, pk);
  }
  addWorkstep({id, step: proof});
};

const updateSyncWorkflow = (workgroupId, id, workstepId) => {
  updateWorkflow({workgroupId, id, type: 'SYNC', workstepId });
};

const broadcastSyncWorkstep = (workgroupId, workflowId) => {
  const workflow = workflowRegistry.get(workflowId);
  getIO().to(workgroupId).emit('workflow:sync', getWorkstep(workflow[SYNC][workflow[SYNC].length - 1]).getProofArtifacts());
};

const handleSyncWorkflow = (syncWorkflow) => {
  let { workgroupId, workflowId, workstepId, ...workflow } = syncWorkflow;
  if (!workstepExists(workstepId)) {
    addSyncWorkstep(workflowId, workstepId, workflow.SYNC);
  };
  updateSyncWorkflow(workgroupId, workflowId, workstepId);
  broadcastSyncWorkstep(workgroupId, workflowId);
};

module.exports = {
  workflowRegistry,
  workflowExists,
  getWorkflow,
  handleSyncWorkflow,
  updateSyncWorkflow,
};