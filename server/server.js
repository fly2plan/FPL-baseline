const express = require('express');
const app = express();
const port = process.env.PORT || 3000;
const http = require('http')
const server = http.createServer(app)
const Web3 = require('web3');
const contract = require('@truffle/contract');
const truffle_connect_specifier = './connection/truffle_connect.js';
const bodyParser = require('body-parser');
const cors = require('cors');
const fs = require('fs')

const { socketConnection } = require('./baseline/utils/socket');
socketConnection(server);

const { bootstrap } = require('./baseline/privacy/service')

bootstrap().then(({ files }) => {
  if (!fs.existsSync('./build/contracts/')) {
    fs.mkdirSync('./build/contracts/', { recursive: true });
  }
  for (const file of files) {
    fs.writeFileSync(`./build/contracts/${file.name}`, file.content, { flag: 'a+' }, (err) => {
      if (err) throw err;
    })
    console.log(`${file.name} successfully written`)
  }
}).then(() => {
  import(truffle_connect_specifier).then(truffle_connect => {
    server.listen(port, async () => {

      const fpl_verifier_artifact = require('./build/contracts/FPLVerifier.json')
      const ack_verifier_artifact = require('./build/contracts/ACKVerifier.json')
      const shield_artifact = require('./build/contracts/Shield.json')

      truffle_connect.default.web3 = new Web3(new Web3.providers.HttpProvider("http://ganache-cli:8545"));
      truffle_connect.default.FPLVerifier = contract(fpl_verifier_artifact)
      truffle_connect.default.ACKVerifier = contract(ack_verifier_artifact)
      truffle_connect.default.Shield = contract(shield_artifact)

      console.log("Express Listening at http://localhost:" + port);
    });
  })
})

const { organizationRouter } = require('./baseline/organization');
const { workgroupRouter } = require('./baseline/workgroup');
const { workflowRouter } = require('./baseline/workflow');

const KafkaConsumer = require('./baseline/messaging/consumer.js');

KafkaConsumer.consume().then(() => {
  console.log('consume start successful');
}).catch(err => {
  console.log('consume start err ', err);
})

app.use(cors());

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

app.use('/', express.static('public_static'));

app.use('/api/organization', organizationRouter)
app.use('/api/workgroup', workgroupRouter)
app.use('/api/workflow', workflowRouter)