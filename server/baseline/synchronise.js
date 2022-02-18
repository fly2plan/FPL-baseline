const express = require('express');
const router = express.Router();

const { hash } = require('./utils/hash');
const { generateProof, verifyProof } = require('./privacy/proof');
const truffle_connect = require('../connection/truffle_connect')

const { targetEventType, proofEventType, gameEventType } = require('./messaging/eventType')
const KafkaProducer = require('./messaging/producer');

const { plans, userInPlan, updatePlan } = require('./plan');

router.get('/:id', (req, res) => {
  if (plans.has(req.params.id)) {
    let plan = plans.get(req.params.id);

    return res.json({ plan });
  }

  res.sendStatus(404);
})

router.put('/hash/:id', async (req, res) => {
  const id = req.body.id;
  const planID = req.params.id;

  if (plans.has(planID)) {
    if (userInPlan(id, planID)) {
      let plan = plans.get(planID);

      const member = plan.members.find(member => member.id === id);

      if (member.hash && member.hash > 0) {
        return res.status(409).send('Member hash already set.');
      }

      const formHash = await hash(req.body);
      member.hash = formHash;
      updateGame(game);

      const gameProducer = new KafkaProducer('game', gameEventType);
      await gameProducer.queue({ id: game.id, shieldContractAddress: game.shieldContractAddress, players: game.players });
      return res.sendStatus(200);
    }

    return res.status(403).send('Action not permitted.');
  }

  res.status(404).send(`Game #${game} does not exist`);
})

router.post('/target', async (req, res) => {
  const targetProducer = new KafkaProducer('battleship', targetEventType);
  await targetProducer.queue(req.body);
  res.sendStatus(200);
})

router.post('/proof', async (req, res) => {
  // const inputSignals = {
  //   shipX: req.body.shipX,
  //   shipY: req.body.shipY,
  //   shipO: req.body.shipO,
  //   shipHash: req.body.shipHash,
  //   targetX: req.body.targetX,
  //   targetY: req.body.targetY
  // }
  const [h1, h2] = req.body.hash
  const proof = await generateProof(h1, h2)
  const result = await verifyProof(proof)

  if (result) {
    const proofMsg = {
      proof,
      input,
      gameId: req.body.gameId,
      playerId: req.body.playerId
    }
    const proofProducer = new KafkaProducer('proof', proofEventType);
    await proofProducer.queue(proofMsg);
    res.sendStatus(200);

    return;
  }

  res.sendStatus(409);
})

router.post('/verify', async (req, res) => {
  const shieldAddress = games.get(req.body.gameId).shieldContractAddress
  await truffle_connect.verify(req.body.proof, req.body.input, shieldAddress, () => {
    res.send('verified');
  });
})

module.exports = {
  battleshipRouter: router,
}