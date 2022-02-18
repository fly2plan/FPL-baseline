const contract = require('truffle-contract');
const shield_artifact = require('../build/contracts/Shield.json')
const fpl_verifier_artifact = require('../build/contracts/FPLVerifier.json')
const ack_verifier_artifact = require('../build/contracts/ACKVerifier.json')

const Web3Utils = require('web3-utils');

var Shield = contract(shield_artifact)
var FPLVerifier = contract(fpl_verifier_artifact)
var ACKVerifier = contract(ack_verifier_artifact)

module.exports = {
  deploy: async function(callback) {
    var self = this;
    FPLVerifier.setProvider(self.web3.currentProvider);
    ACKVerifier.setProvider(self.web3.currentProvider);
    Shield.setProvider(self.web3.currentProvider);

    self.web3.eth.getAccounts(async (err, accs) => {
      if (err != null) {
        console.log("There was an error fetching your accounts.", err);
        callback([]);
        return;
      }
      if (accs.length == 0) {
        console.log("Couldn't get any accounts! Make sure your Ethereum client is configured correctly.");
        return;
      }
      self.accounts = accs;
      self.account = self.accounts[2];

      const FPLVerifierInstance = await FPLVerifier.new({from: self.account, gas: 3000000});
      const ACKVerifierInstance = await ACKVerifier.new({from: self.account, gas: 3000000});
      const shieldInstance = await Shield.new(FPLVerifierInstance.address, ACKVerifierInstance.address, 0, {from: self.account, gas: 3000000});
      callback(shieldInstance.address);
    })
  },
  verify: async function(proof, input, shieldAddress, type, callback) {
    var self = this;

    Shield.setProvider(self.web3.currentProvider);
    const deployed = await Shield.at(shieldAddress);

    // TODO: for now just hashing input, should something else be put in the commitment?
    const commitment = Web3Utils.sha3(Web3Utils.toHex(input));
    if (type === 'fpl') {
      const res = await deployed.verifyAndPushFPL.call(proof, input, commitment, { from: self.account });
      callback(res);
    } else if (type === 'ack') {
      const res = await deployed.verifyAndPushACK.call(proof, input, commitment, { from: self.account });
      callback(res);
    };
  },
}
