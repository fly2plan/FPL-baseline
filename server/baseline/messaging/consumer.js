const KafkaConfig = require('./config.js');
const { orgEventType, workgroupEventType, workflowSyncEventType } = require('./eventType.js');
let kafkaConfig = new KafkaConfig();

const { insertOrg } = require('../organizationRegistry');
const { updateWorkgroup } = require('../workgroupRegistry');
const { handleSyncWorkflow } = require('../workflowRegistry');

async function consume() {
  const stream = kafkaConfig.consumer();

  stream.on('data', function(data) {
    console.log('message received', data);
    switch(data.topic) {
      case 'orgReg':
        insertOrg(orgEventType.fromBuffer(data.value))
        break;
      case 'workgroupReg': 
        updateWorkgroup(workgroupEventType.fromBuffer(data.value))
        break;
      case 'workflowSync':
        const syncWorkflow = workflowSyncEventType.fromBuffer(data.value);
        handleSyncWorkflow(syncWorkflow);
        break;
      default:
        console.warn('unsupported topic received ', data.topic)
      };
  });
};

module.exports = {
  consume
};