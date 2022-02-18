const { getSocket, getIO } = require("./utils/socket");

let workgroupRegistry = new Map();

const workgroupExists = (id) => {
  return workgroupRegistry.has(id)
};

const joinWorkgroup = (workgroupId, session) => {
  const socket = getSocket(session);

  socket.join(workgroupId); //joining a 'room' with name of value of game which is the workgroup id

  socket.on('workgroup:action', (workgroup, message) => {
    console.log('Received:', workgroup, message);
  });
}

const updateWorkgroup = (workgroup) => {
  if (workgroupRegistry.has(workgroup.id)) {
    console.log(`updating workgroup with id ${workgroup.id}`);
  } else {
    console.log('adding new workgroup ', workgroup);
  }

  workgroupRegistry.set(workgroup.id, workgroup);

  if (workgroup.members.length > 1) {
    broadcastWorkgroup(workgroup.id);
  };

}

const broadcastWorkgroup = (workgroupId) => {
  const workgroup = workgroupRegistry.get(workgroupId);

  if (!workgroup) {
    console.log(`broadcastWorkgroup - workgroup ${workgroupId} not found`);
    return;
  };

  getIO().to(workgroupId).emit('workgroup:update', workgroup);
}

const getShieldAddress = (workgroupId) => {
  const workgroup = workgroupRegistry.get(workgroupId);

  if (!workgroup) {
    console.log(`getShieldAddress - workgroup ${workgroupId} not found`);
    return;
  }


  return workgroup.shieldContractAddress;
}

module.exports = {
  workgroupRegistry,
  workgroupExists,
  updateWorkgroup,
  joinWorkgroup,
  broadcastWorkgroup,
  getShieldAddress
}