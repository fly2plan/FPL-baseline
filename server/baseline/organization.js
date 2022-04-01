const express = require('express')
const router = express.Router()

const { v4: uuidv4 } = require('uuid');

const { orgEventType } = require('./messaging/eventType.js');
const FPLMessageProducer = require('./messaging/producer.js');

const { orgRegistry } = require('./organizationRegistry');
const { getKeys } = require('./privacy/service.js');


router.get('/:id', (req, res) => {
  if (orgRegistry.has(req.params.id)) {
    return res.json({ id: req.params.id, name: orgRegistry.get(req.params.id).name, pk: orgRegistry.get(req.params.id).pk });
  };

  res.sendStatus(404);
})

router.post('', async (req, res) => {
  const id = uuidv4();
  if (req.body.name) {
    let {sk, pk} = await getKeys()
    let org = { id, name: req.body.name, pk: pk };

    const producer = new FPLMessageProducer('orgReg', orgEventType);
    await producer.queue(org, orgEventType);

    return res.json({ id: id, sk: sk });
  }

  res.sendStatus(400);
})

module.exports = {
  organizationRouter: router,
}