const Web3Utils = require('web3-utils');

module.exports = {
  deploy: async function (callback) {
    var self = this;
    self.FPLVerifier.setProvider(self.web3.currentProvider);
    self.ACKVerifier.setProvider(self.web3.currentProvider);
    self.Shield.setProvider(self.web3.currentProvider);

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

      const FPLVerifierInstance = await self.FPLVerifier.new({ from: self.account, gas: 3000000 });
      const ACKVerifierInstance = await self.ACKVerifier.new({ from: self.account, gas: 3000000 });
      const shieldInstance = await self.Shield.new(FPLVerifierInstance.address, ACKVerifierInstance.address, 0, { from: self.account, gas: 3000000 });

      callback(shieldInstance.address);
    })
  },
  verify: async function (type, a, b, c, input, shieldAddress, callback) {
    var self = this;

    self.Shield.setProvider(self.web3.currentProvider);
    const deployed = await self.Shield.at(shieldAddress);
    console.log(deployed)

    // TODO: for now just hashing input, should something else be put in the commitment?
    const commitment = Web3Utils.sha3(Web3Utils.toHex(input));
    if (type === 'fpl') {
      const res = await deployed.verifyAndPushFPL.call(a, b, c, input, commitment, { from: self.account });
      callback(res);
    } else if (type === 'ack') {
      const res = await deployed.verifyAndPushACK.call(a, b, c, input, commitment, { from: self.account });
      callback(res);
    };
  },
}
