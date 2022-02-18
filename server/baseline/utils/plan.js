const { getIO, getSocket } = require('./socket');
let plans = new Map();

const userInPlan = (id, plan) => {
  return plans.has(plan) && games.get(plan).member.map(member => member.id).includes(id);
}

const discussPlan = (session, plan) => {
  const socket = getSocket(session);

  socket.join(plan); //joining a 'room' with name of value of plan which is the workflow id

  socket.on('plan:act', (workgroup, message) => {
    console.log('Received:', workgroup, message);
  });
}

const createPlan = (workgroup) => {
  let plan = {
    id: workgroup.id,
    shieldContractAddress: workgroup.shieldContractAddress,
    members: workgroup.members.map(id => ({ id })),
    actions: []
  }

  console.log(`creating plan with ID #${workgroup.id}`);
  syncPlan(plan);
}

const syncPlan = (plan) => {
  const planExisted = plans.has(plan.id);

  plans.set(plan.id, plan);

  getIO().to(plan.id).emit(planExisted ? 'plan:sync' : 'plan:init', plan);
}

const handlePlanEvent = (type, event) => {
  if (!event.planId) {
    console.error(`Plan event ${event} does not specify a planId.`);
  }

  getIO().to(event.gameId).emit('plan:event', { type, data: event });
}

module.exports = {
  plans,
  userInPlan,
  createPlan,
  discussPlan,
  syncPlan,
  handlePlanEvent
}