const { getFPLProof, getACKProof } = require('./service')

class SyncProof {
    constructor(preimage) {
        this.preimage = preimage;
        this.hash = null;
        this.pk = null;
        this.signature = null;
        this.proof = null;
        if (new.target === SyncProof) {
            throw new TypeError("Cannot construct " + new.target.name + " instances directly");
        };
    };
}

class FPLSyncProof extends SyncProof {

    constructor(fpl, hash, pk, signature, proof) {
        super(fpl);
        this.fpl = JSON.parse(JSON.stringify(fpl));
        this.hash = hash;
        this.pk = pk;
        this.signature = signature;
        this.proof = proof;
    };

    async createProof(privateKey) {
        const {proof, inputs} = await getFPLProof(this.fpl, privateKey);
        const { hash, sig, pk } = inputs
        this.proof = proof;
        this.pk = pk
        this.signature = sig
        this.hash = hash
        return proof
    }

    getProofAndInputs() {
        return {
            raw: this.fpl,
            proof: this.proof,
            inputs: {
                hash: this.hash,
                sig: this.signature,
                pk: this.pk, 
            }
        };
    };

}

class ACKSyncProof extends SyncProof {

    constructor(fplProofInputs, hash, pk, signature, proof) {
        super(fplProofInputs)
        this.fplProofInputs = fplProofInputs;
        this.hash = hash;
        this.pk = pk;
        this.signature = signature;
        this.proof = proof;
    };

    async createProof(privateKey) {
        const {proof, inputs} = await getACKProof(this.fplProofInputs, privateKey);
        const { hash, sig, pk } = inputs
        this.proof = proof;
        this.pk = pk
        this.signature = sig
        this.hash = hash
        return proof
    }

    getProofAndInputs() {
        return {
            raw: this.hash,
            proof: this.proof,
            inputs: {
                ack_hash: this.hash,
                nm_sig: this.signature,
                nm_pk: this.pk,
                fpl_hash: this.fplProofInputs.inputs.hash,
                ao_sig: this.fplProofInputs.inputs.sig,
                ao_pk: this.fplProofInputs.inputs.pk
            }
        };
    }
}

module.exports = {
    FPLSyncProof,
    ACKSyncProof
}