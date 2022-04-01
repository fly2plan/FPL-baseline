const express = require('express')
const router = express.Router()

const { v4: uuidv4 } = require('uuid');

const { organizationExists, getOrg } = require('./organizationRegistry');

const truffle_connect = require('../connection/truffle_connect')

const { workgroupRegistry, joinWorkgroup } = require('./workgroupRegistry')

const { workgroupEventType } = require('./messaging/eventType');
const FPLMessageProducer = require('./messaging/producer');


/*
 * We chose to use a truncated ID with only ID_LENGTH characters for the workgroup ID
 * to have simpler UX when joining workgroup.
 * It is implemented by taking first ID_LENGTH characters of uuid.
 */
const ID_LENGTH = 36;

router.post('', async (req, res) => {
  if (!req.body.session || !organizationExists(req.body.id)) {
    return res.sendStatus(403);
  }

  await truffle_connect.deploy(async function (shieldContractAddress) {
    let id = uuidv4().slice(0, ID_LENGTH);
    while (workgroupRegistry.has(id)) {
      id = uuidv4().slice(0, ID_LENGTH);
    }

    let workgroup = {
      id: id,
      members: [{ id: req.body.id, pk: getOrg(req.body.id).pk }],
      shieldContractAddress,
      workflows: []
    };

    joinWorkgroup(workgroup.id, req.body.session)

    const producer = new FPLMessageProducer('workgroupReg', workgroupEventType);
    await producer.queue({ ...workgroup }, workgroupEventType)

    res.json({ id: id });
  });
})

router.post('/join/:id', async (req, res) => {
  if (!req.body.session || !organizationExists(req.body.id)) {
    return res.sendStatus(403);
  }

  const org = getOrg(req.body.id)
  const id = req.params.id

  if (workgroupRegistry.has(id)) {
    let workgroup = workgroupRegistry.get(id);

    if (workgroup.members.length == 1) {
      if (workgroup.members[0] !== req.body.id) {
        workgroup = {
          id: workgroup.id,
          shieldContractAddress: workgroup.shieldContractAddress,
          members: [
            workgroup.members[0],
            { id: org.id, pk: org.pk }
          ],
          workflows: []
        }

        joinWorkgroup(workgroup.id, req.body.session)

        const producer = new FPLMessageProducer('workgroupReg', workgroupEventType);
        await producer.queue({ ...workgroup }, workgroupEventType);

        return res.json({ id: id });
      }

      return res.status(409).json({ error: 'User in workgroup' });
    }

    return res.status(409).json({ error: 'Workgroup full' });
  }

  res.sendStatus(404);
})


module.exports = {
  workgroupRouter: router,
}