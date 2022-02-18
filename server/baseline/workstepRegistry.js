const { FPLSyncProof, ACKSyncProof } = require("./privacy/proof");

let workstepRegistry = new Map();

const workstepExists = (id) => {
    return workstepRegistry.has(id);
};

const getWorkstep = (id) => {
    return workstepRegistry.get(id);
}

const addWorkstep = (workstep) => {
    if(workstepExists(workstep.id)) {
        return;
    };
    workstepRegistry.set(workstep.id, workstep.step)
}

module.exports = {
    addWorkstep,
    getWorkstep,
    workstepExists
}