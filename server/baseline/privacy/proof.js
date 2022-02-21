const { initialize } = require('zokrates-js/node')
const fs = require('fs')
const path = require('path')
const CryptoJS = require('crypto-js')

const FPLProgramPath = path.join(__dirname, '/circuits/artifacts/fpl/out')
const FPLabi = require('./circuits/artifacts/fpl/abi.json')
const FPLProvingKeyPath = path.join(__dirname, '/circuits/artifacts/fpl/proving.key')
const FPLVerificationKeyPath = path.join(__dirname, '/circuits/artifacts/fpl/verification.key')

const ACKProgramPath = path.join(__dirname, '/circuits/artifacts/ack/out')
const ACKabi = require('./circuits/artifacts/ack/abi.json')
const ACKProvingKeyPath = path.join(__dirname, '/circuits/artifacts/ack/proving.key')
const ACKVerificationKeyPath = path.join(__dirname, '/circuits/artifacts/ack/verification.key')

const { bitGroups, hashToIntPair } = require('../utils/hash')

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

    #createHash() {
        const castToBinary = (arrays) => {
            const fillBits = (bits, maxBits) => {
                if (bits.length < maxBits) {
                    bits = bits.concat(Array(maxBits - bits.length).fill(0).join(''));
                };
                return bits;
            };

            for (let [keys, maxBits] of arrays) {
                let field = this.preimage;
                for (const [i, key] of keys.entries()) {
                    let nextField = field[key];
                    if (i === keys.length - 1) {
                        if (Array.isArray(nextField)) {
                            let newField = [];
                            for (let element of nextField) {
                                element = BigInt(element).toString(2);
                                element = fillBits(element, maxBits);
                                newField.push(element);
                            };
                            nextField = newField;
                        } else if (typeof nextField === 'string') {
                            nextField = BigInt(nextField).toString(2);
                            nextField = fillBits(nextField, maxBits);
                        };
                        field[key] = nextField;
                    } else {
                        field = field[key];
                    }
                }
            };
        };

        castToBinary([
            [['ai'], 64],
            [['fr'], 8],
            [['tof'], 8],
            [['noa'], 16],
            [['toa'], 32],
            [['wtc'], 8],
            [['eandc', 'rc_n_aae', 's'], 16],
            [['eandc', 'surveillance', 's'], 8],
            [['eandc', 'surveillance', 'ssr', 'mode_a_c', 'a'], 16],
            [['eandc', 'surveillance', 'ssr', 'mode_s', 'a'], 64],
            [['eandc', 'surveillance', 'ads', 'b', 'a'], 64],
            [['eandc', 'surveillance', 'ads', 'c', 'a'], 32],
            [['eandc', 'rc_n_aae', 'a'], 64],
            [['t'], 32],
            [['da'], 32],
            [['de'], 32],
            [['teet'], 32],
            [['ada'], 32],
            [['cs'], 64],
            [['cl'], 64],
            [['r'], 64],
            [['adas'], 32],
            [['oi'], 64],
            [['e'], 64],
            [['pob'], 64],
            [['er'], 32],
            [['se'], 32],
            [['j'], 64],
            [['d'], 64],
            [['acm'], 64],
            [['re'], 64],
            [['pn'], 64],
        ]);

        const groups = bitGroups([
            [[
                this.preimage.ai,
                this.preimage.fr,
                this.preimage.tof,
                this.preimage.noa,
                this.preimage.toa,
            ], 'h1_a'],
            [[
                this.preimage.wtc,
                this.preimage.eandc.rc_n_aae.s,
                this.preimage.eandc.surveillance.s,
                this.preimage.eandc.surveillance.ssr.mode_a_c.a,
                this.preimage.eandc.surveillance.ssr.mode_s.a,
            ], 'h1_b'],
            [[
                this.preimage.eandc.surveillance.ads.b.a[0],
                this.preimage.eandc.surveillance.ads.b.a[1],
            ], 'h1_c'],
            [[
                this.preimage.eandc.surveillance.ads.c.a,
                this.preimage.eandc.rc_n_aae.a[0],
                this.preimage.t,
            ], 'h1_d'],
            [[
                this.preimage.eandc.rc_n_aae.a[1],
                this.preimage.eandc.rc_n_aae.a[2],
            ], 'h2_a'],
            [[
                this.preimage.eandc.rc_n_aae.a[3],
                this.preimage.eandc.rc_n_aae.a[4],
            ], 'h2_b'],
            [[
                this.preimage.da,
                this.preimage.de,
                this.preimage.teet,
                this.preimage.ada,
            ], 'h2_c'],
            [[
                this.preimage.cs,
                this.preimage.cl,
            ], 'h2_d'],
            [[
                this.preimage.r[0],
                this.preimage.r[1],
            ], 'h3_a'],
            [[
                this.preimage.r[2],
                this.preimage.r[3],
            ], 'h3_b'],
            [[
                this.preimage.r[4],
                this.preimage.r[5],
            ], 'h3_c'],
            [[
                this.preimage.r[6],
                this.preimage.r[7],
            ], 'h3_d'],
            [[
                this.preimage.r[8],
                this.preimage.r[9],
            ], 'h4_a'],
            [[
                this.preimage.r[10],
                this.preimage.r[11],
            ], 'h4_b'],
            [[
                this.preimage.r[12],
                this.preimage.r[13],
            ], 'h4_c'],
            [[
                this.preimage.r[14],
                this.preimage.r[15],
            ], 'h4_d'],
            [[
                this.preimage.adas[0],
                this.preimage.adas[1],
                this.preimage.oi[0],
            ], 'h5_a'],
            [[
                this.preimage.oi[1],
                this.preimage.oi[2],
            ], 'h5_b'],
            [[
                this.preimage.oi[3],
                this.preimage.oi[4],
            ], 'h5_c'],
            [[
                this.preimage.oi[5],
                this.preimage.oi[6],
            ], 'h5_d'],
            [[
                this.preimage.oi[7],
                this.preimage.e,
            ], 'h6_a'],
            [[
                this.preimage.pob,
                this.preimage.er,
                this.preimage.se,
            ], 'h6_b'],
            [[
                this.preimage.j[0],
                this.preimage.j[1],
            ], 'h6_c'],
            [[
                this.preimage.d[0],
                this.preimage.d[1],
            ], 'h6_d'],
            [[
                this.preimage.d[2],
                this.preimage.d[3],
            ], 'h7_a'],
            [[
                this.preimage.acm[0],
                this.preimage.acm[1],
            ], 'h7_b'],
            [[
                this.preimage.acm[2],
                this.preimage.acm[3],
            ], 'h7_c'],
            [[
                this.preimage.re[0],
                this.preimage.re[1],
            ], 'h7_d'],
            [[
                this.preimage.re[2],
                this.preimage.re[3],
            ], 'h8_a'],
            [[
                this.preimage.pn[0],
                this.preimage.pn[1],
            ], 'h8_b'],
            [[
                this.preimage.eandc.rc_n_aae.a[5],
            ], 'h8_c'],
        ]);

        let h = []
        for (let i = 1; i < 8; i++) {
            let hashPreimage = ''
            hashPreimage = hashPreimage.concat(groups[`h${i}_a`], groups[`h${i}_b`], groups[`h${i}_c`], groups[`h${i}_d`]);
            //console.log(i, hashPreimage);
            h.push(CryptoJS.enc.Hex.parse(BigInt('0b' + hashPreimage).toString(16)).toString());
        }
        let lastHashPreimage = groups['h8_a'].concat(groups['h8_b'], groups['h8_c'], Array(128).fill('0').join(''))
        h.push(CryptoJS.enc.Hex.parse(BigInt('0b' + lastHashPreimage).toString(16)).toString())

        //Combine hashes into one hash
        const h12 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse(h[0].concat(h[1]))).toString();
        const h34 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse(h[2].concat(h[3]))).toString();
        const h56 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse(h[4].concat(h[5]))).toString();
        const h78 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse(h[6].concat(h[7]))).toString();

        const h1234 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse(h12.concat(h34))).toString();
        const h5678 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse(h56.concat(h78))).toString();

        const h12345678 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse(h1234.concat(h5678))).toString();
        return h12345678;
    }

    #parsePreimage() {
        //Right pad FPL fields with spaces
        const fillFPLBlanks = (fplFields) => {
            for (const [fplField, maxLength] of fplFields) {
                let field = this.preimage;
                for (const item of fplField) {
                    field = field[`${item}`];
                };
                if (field.length < maxLength) {
                    field = field + Array(maxLength - field.length).fill(' ').join('');
                };
            };
        };

        fillFPLBlanks([
            [['ai'], 7],
            [['fr'], 1],
            [['tof'], 1],
            [['noa'], 2],
            [['toa'], 4],
            [['wtc'], 1],
            [['eandc', 'rc_n_aae', 's'], 4],
            [['eandc', 'surveillance', 's'], 1],
            [['eandc', 'surveillance', 'ssr', 'mode_a_c', 'a'], 2],
            [['eandc', 'surveillance', 'ssr', 'mode_s', 'a'], 7],
            [['eandc', 'surveillance', 'ads', 'c', 'a'], 4],
            [['eandc', 'surveillance', 'ads', 'b', 'a'], 12],
            [['t'], 4],
            [['da'], 4],
            [['cs'], 5],
            [['cl'], 5],
            [['r'], 152],
            [['de'], 4],
            [['teet'], 4],
            [['ada'], 4],
            [['adas', '0'], 4],
            [['adas', '1'], 4],
            [['oi'], 152],
            [['e'], 6],
            [['pob'], 5],
            [['er'], 4],
            [['se'], 4],
            [['j'], 10],
            [['d'], 38],
            [['acm'], 38],
            [['re'], 38],
            [['pn'], 19]
        ]);

        const convertToASCII = (object) => {
            if (Array.isArray(object)) {
                for (const [index, element] of object.entries()) {
                    if (typeof element === 'object') {
                        convertToASCII(element);
                    } else if (typeof element === 'string') {
                        object[index] = element.split('').map(c => c.charCodeAt(0)).join('');
                    };
                };
            } else if (typeof object === 'object') {
                for (const [key, value] of Object.entries(object)) {
                    if (typeof value === 'object') {
                        convertToASCII(value);
                    } else if (typeof value == 'string') {
                        object[key] = value.split('').map(c => c.charCodeAt(0)).join('');
                    };
                };
            };
        };

        convertToASCII(this.preimage);

        const segmentFields = (object) => {
            const splitString = (string) => {
                let segments = [];
                while (string.length >= 19) {
                    let segment = string.slice(0, 19);
                    segments.push(segment);
                    string = string.slice(19, string.length);
                };
                if (string.length != 0) {
                    segments.push(string);
                };
                if (segments.length == 1) {
                    string = segments[0];
                } else if (segments.length > 1) {
                    string = segments;
                };
                return string;
            };

            if (Array.isArray(object)) {
                for (const [index, element] of object.entries()) {
                    if (typeof element === 'object') {
                        segmentFields(element);
                    } else if (typeof element === 'string') {
                        let segments = splitString(element);
                        object[index] = segments;
                    };
                };
            } else if (typeof object === 'object') {
                for (const [key, value] of Object.entries(object)) {
                    if (typeof value === 'object') {
                        segmentFields(value);
                    } else if (typeof value == 'string') {
                        let segments = splitString(value);
                        object[key] = segments;
                    };
                };
            };
        }

        segmentFields(this.preimage);
    };

    #preparePreimage() {
        const binAsHex = (object) => {
            if (Array.isArray(object)) {
                for (const [index, element] of object.entries()) {
                    if (typeof element === 'object') {
                        binAsHex(element);
                    } else if (typeof element === 'string') {
                        object[index] = '0x' + BigInt('0b' + element).toString(16);
                    };
                };
            } else if (typeof object === 'object') {
                for (const [key, value] of Object.entries(object)) {
                    if (typeof value === 'object') {
                        binAsHex(value);
                    } else if (typeof value === 'string') {
                        object[key] = '0x' + BigInt('0b' + value).toString(16);
                };
                }
            }
        };

        binAsHex(this.preimage);
    }

    constructor(fpl, pk, signature, proof) {
        super(fpl);
        this.fpl = JSON.parse(JSON.stringify(fpl));
        this.#parsePreimage();
        this.hash = this.#createHash();
        this.#preparePreimage()
        this.pk = pk;
        this.signature = signature;
        this.proof = proof;
    };

    async createProof(signature) {
        this.signature = signature;
        const program = fs.readFileSync(FPLProgramPath);
        const provingKey = fs.readFileSync(FPLProvingKeyPath);
        console.log(this.preimage.eandc.surveillance.ads.b.a, hashToIntPair(this.hash), this.signature, this.pk);
        this.proof = await initialize().then(zokratesProvider => {
            const { witness, } = zokratesProvider.computeWitness({ program, abi: FPLabi }, [
                this.preimage,
                hashToIntPair(this.hash),
                this.signature,
                this.pk
            ]);
            return zokratesProvider.generateProof(program, witness, provingKey);
        });
        return this.proof;
    }

    getProofArtifacts() {
        return {
            raw: this.fpl,
            hash: this.hash,
            signature: this.signature,
            proof: this.proof,
        };
    };

}

class ACKSyncProof extends SyncProof {

    #createHash() {
        const hpre = CryptoJS.enc.Hex.parse([
            Array(122).fill('0') + ['65', '67', '75'].join(''),
            BigInt(this.preimage[0]).toString(16),
            BigInt(this.preimage[1]).toString(16),
            Array(128).fill('0').join('')
        ].reduce((previousArray, currentArray) => previousArray.concat(currentArray)));
        const h = CryptoJS.SHA256(hpre);
        return h;
    };

    constructor(fplSyncProof, pk) {
        super(fplSyncProof.hash);
        this.fplSyncProof = fplSyncProof;
        this.pk = pk;
        this.hash = this.#createHash();
        this.signature = null;
        this.proof = null;
    };

    async createProof(signature) {
        this.signature = signature;
        const program = new s.readFileSync(ACKProgramPath);
        const provingKey = new fs.readFileSync(ACKProvingKeyPath);
        this.proof = await initialize().then(zokratesProvider => {
            const { witness, } = zokratesProvider.computeWitness({ program, abi: ACKabi }, [
                this.FPLSyncProof.preimage,
                hashToIntPair(this.FPLSyncProof.hash),
                this.FPLSyncProof.signature,
                this.FPLSyncProof.pk,
                hashToIntPair(this.hash),
                this.signature,
                this.pk
            ]);
            return zokratesProvider.generateProof(program, witness, provingKey);
        });
        return this.proof;
    };

    getProofArtifacts() {
        return {
            raw: this.fplSyncProof.hash,
            hash: this.hash,
            signature: this.signature,
            proof: this.proof,
        };
    };
}

const verifyFPLProof = async (proof) => {
    const FPLVerificationKey = fs.readFileSync(FPLVerificationKeyPath);
    return await initialize().then(zokratesProvider => {
        return zokratesProvider.verify(FPLVerificationKey, proof);
    });
};

const verifyACKProof = async (proof) => {
    const ACKVerificationKey = fs.readFileSync(ACKVerificationKeyPath);
    return await initialize().then(zokratesProvider => {
        return zokratesProvider.verify(ACKVerificationKey, proof);
    });
};

module.exports = {
    FPLSyncProof,
    ACKSyncProof,
    verifyFPLProof,
    verifyACKProof
}