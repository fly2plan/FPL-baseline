const { initialize } = require('zokrates-js/node')
const fs = require('fs')
const path = require('path')
const CryptoJS = require('crypto-js')

const FPLProgramPath = path.join(__dirname, '/circuits/artifacts/fpl/fpl')
const FPLabi = require('./circuits/artifacts/fpl/abi.json')
const FPLProvingKeyPath = path.join(__dirname, '/circuits/artifacts/fpl/proving.key')
const FPLVerificationKeyPath = path.join(__dirname, '/circuits/artifacts/fpl/verification.key')

const ACKProgramPath = path.join(__dirname, '/circuits/artifacts/ack/ack')
const ACKabi = require('./circuits/artifacts/ack/abi.json')
const ACKProvingKeyPath = path.join(__dirname, '/circuits/artifacts/ack/proving.key')
const ACKVerificationKeyPath = path.join(__dirname, '/circuits/artifacts/ack/verification.key')

const TestProgramPath = path.join(__dirname, '/circuits/artifacts/test/out')
const Testabi = require('./circuits/artifacts/test/abi.json')
const TestProvingKeyPath = path.join(__dirname, '/circuits/artifacts/test/proving.key')
const TestVerificationKeyPath = path.join(__dirname, '/circuits/artifacts/test/verification.key')

const { bitGroups, castToBinary, joinPreimageDigits, digitsToDecimal, decimalToHexString, prependWithZeros, hashToHexPair, hashToIntPair } = require('../utils/hash')

// const FPL = null
// const ACK = null

// const fileSystemResolver = (from, to) => {
//     const location = path.resolve(path.dirname(path.resolve(from)), to);
//     console.log(from, to, location);
//     const source = fs.readFileSync(location).toString();
//     return { source, location };
// };

// const Initialise = async () => {
//     [FPL, ACK] = await initialize().then(zokratesProvider => {
//         const fpl_source = fs.readFileSync(path.join(__dirname, '/circuits/fpl.zok')).toString();
//         const fpl_options = {
//             location: "./circuits/fpl.zok", // location of the root module
//             resolveCallback: fileSystemResolver
//         };
//         const ack_source = fs.readFileSync(path.join(__dirname, '/circuits/ack.zok')).toString();
//         const ack_options = {
//             location: "./circuits/ack.zok",
//             resolveCallback: fileSystemResolver
//         }
//         console.log('creating fpl artifacts...')
//         const fpl_artifacts = zokratesProvider.compile(fpl_source, fpl_options);
//         console.log('creating ack artifacts...')
//         const ack_artifacts = zokratesProvider.compile(ack_source, ack_options);
//         console.log('creating fpl keypair...')
//         const fpl_keypair = zokratesProvider.setup(fpl_artifacts.program);
//         console.log('creating ack keypair...')
//         const ack_keypair = zokratesProvider.setup(ack_artifacts.program);
//         return [{ artifacts: fpl_artifacts, keypair: fpl_keypair }, { artifacts: ack_artifacts, keypair: ack_keypair }];
//     });
//     return true;
// }

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
        //Hash sections of the document, ensuring that 512 bits is not exceeded
        // const h1pre = joinPreimageDigits([
        //     this.preimage.ai,
        //     this.preimage.fr,
        //     this.preimage.tof,
        //     this.preimage.noa,
        //     this.preimage.toa,
        //     this.preimage.wtc,
        //     this.preimage.eandc.rc_n_aae.s,
        //     this.preimage.eandc.surveillance.s,
        //     this.preimage.eandc.surveillance.ssr.mode_a_c.a,
        //     this.preimage.eandc.surveillance.ssr.mode_s.a,
        //     this.preimage.eandc.surveillance.ads.c.a,
        //     this.preimage.eandc.surveillance.ads.b.a,
        //     this.preimage.t,
        //     this.preimage.da,
        //     this.preimage.cs,
        //     this.preimage.de,
        //     this.preimage.teet,
        //     this.preimage.ada
        // ]);
        // let h1preSplit = [h1pre.slice(0, 37), h1pre.slice(37, 74), h1pre.slice(74, 111), h1pre.slice(111, 148)];
        // let h1prePadded = h1preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h1preComplete = CryptoJS.enc.Hex.parse(h1prePadded);
        // const h1 = CryptoJS.SHA256(h1preComplete).toString();
        // const [h1_a, h1_b] = hashToHexPair(h1);

        // const h2pre = joinPreimageDigits([this.preimage.eandc.rc_n_aae.a]);
        // let h2preSplit = [h2pre.slice(0, 24), h2pre.slice(24, 49), h2pre.slice(49, 73), h2pre.slice(73, 98)];
        // let h2prePadded = h2preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h2preComplete = CryptoJS.enc.Hex.parse(h2prePadded);
        // const h2 = CryptoJS.SHA256(h2preComplete).toString();
        // const [h2_a, h2_b] = hashToHexPair(h2);

        // const h3pre = joinPreimageDigits([this.preimage.r.slice(0, 152)]);
        // let h3preSplit = [h3pre.slice(0, 38), h3pre.slice(38, 76), h3pre.slice(76, 114), h3pre.slice(114, 156)];
        // let h3prePadded = h3preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h3preComplete = CryptoJS.enc.Hex.parse(h3prePadded);
        // const h3 = CryptoJS.SHA256(h3preComplete).toString();
        // const [h3_a, h3_b] = hashToHexPair(h3);

        // const h4pre = joinPreimageDigits([this.preimage.r.slice(152, 304)]);
        // let h4preSplit = [h4pre.slice(0, 38), h4pre.slice(38, 76), h4pre.slice(76, 114), h4pre.slice(114, 156)];
        // let h4prePadded = h4preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h4preComplete = CryptoJS.enc.Hex.parse(h4prePadded);
        // const h4 = CryptoJS.SHA256(h4preComplete).toString();
        // const [h4_a, h4_b] = hashToHexPair(h4);

        // const h5pre = joinPreimageDigits([this.preimage.oi.slice(0, 152)]);
        // let h5preSplit = [h5pre.slice(0, 38), h5pre.slice(38, 76), h5pre.slice(76, 114), h5pre.slice(114, 156)];
        // let h5prePadded = h5preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h5preComplete = CryptoJS.enc.Hex.parse(h5prePadded);
        // const h5 = CryptoJS.SHA256(h5preComplete).toString();
        // const [h5_a, h5_b] = hashToHexPair(h5);

        // const h6pre = joinPreimageDigits([this.preimage.d, this.preimage.acm]);
        // let h6preSplit = [h6pre.slice(0, 38), h6pre.slice(0, 76), h6pre.slice(76, 114), h6pre.slice(114, 156)];
        // let h6prePadded = h6preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h6preComplete = CryptoJS.enc.Hex.parse(h6prePadded);
        // const h6 = CryptoJS.SHA256(h6preComplete).toString();
        // const [h6_a, h6_b] = hashToHexPair(h6);

        // const h7pre = joinPreimageDigits([
        //     this.preimage.adas[0],
        //     this.preimage.adas[1],
        //     this.preimage.e,
        //     this.preimage.pob,
        //     this.preimage.er,
        //     this.preimage.se,
        //     this.preimage.pn,
        // ]);
        // let h7preSplit = [h7pre.slice(0, 24), h7pre.slice(24, 48), h7pre.slice(48, 72), h7pre.slice(72, 96)];
        // let h7prePadded = h7preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h7preComplete = CryptoJS.enc.Hex.parse(h7prePadded);
        // const h7 = CryptoJS.SHA256(h7preComplete).toString();
        // const [h7_a, h7_b] = hashToHexPair(h7);

        // const h8pre = joinPreimageDigits([
        //     this.preimage.j,
        //     this.preimage.re
        // ]);
        // let h8preSplit = [h8pre.slice(0, 24), h8pre.slice(24, 48), h8pre.slice(48, 72), h8pre.slice(72, 96)];
        // let h8prePadded = h8preSplit.map((digits,) => prependWithZeros(decimalToHexString(digitsToDecimal(digits)))).reduce((previousString, currentString) => previousString.concat(currentString));
        // let h8preComplete = CryptoJS.enc.Hex.parse(h8prePadded);
        // const h8 = CryptoJS.SHA256(h8preComplete).toString();
        // const [h8_a, h8_b] = hashToHexPair(h8);

        // //Combine hashes into one hash
        // const h12 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse([h1_a, h1_b, h2_a, h2_b].reduce((previousArray, currentArray) => previousArray.concat(currentArray)))).toString();
        // const h34 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse([h3_a, h3_b, h4_a, h4_b].reduce((previousArray, currentArray) => previousArray.concat(currentArray)))).toString();
        // const h56 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse([h5_a, h5_b, h6_a, h6_b].reduce((previousArray, currentArray) => previousArray.concat(currentArray)))).toString();
        // const h78 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse([h7_a, h7_b, h8_a, h8_b].reduce((previousArray, currentArray) => previousArray.concat(currentArray)))).toString();

        // const [h12_a, h12_b] = hashToHexPair(h12);
        // const [h34_a, h34_b] = hashToHexPair(h34);
        // const [h56_a, h56_b] = hashToHexPair(h56);
        // const [h78_a, h78_b] = hashToHexPair(h78);

        // const h1234 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse([h12_a, h12_b, h34_a, h34_b].reduce((previousArray, currentArray) => previousArray.concat(currentArray)))).toString();
        // const h5678 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse([h56_a, h56_b, h78_a, h78_b].reduce((previousArray, currentArray) => previousArray.concat(currentArray)))).toString();

        // const [h1234_a, h1234_b] = hashToHexPair(h1234);
        // const [h5678_a, h5678_b] = hashToHexPair(h5678);

        // const h12345678 = CryptoJS.SHA256(CryptoJS.enc.Hex.parse([h1234_a, h1234_b, h5678_a, h5678_b].reduce((previousArray, currentArray) => previousArray.concat(currentArray)))).toString();
        // return h12345678;

        //Hash sections of the document, ensuring that 512 bits is not exceeded
        // fillFPLBlanks([
        //     [['ai'], 7],
        //     [['fr'], 1],
        //     [['tof'], 1],
        //     [['noa'], 2],
        //     [['toa'], 4],
        //     [['wtc'], 1],
        //     [['eandc', 'rc_n_aae', 's'], 4],
        //     [['eandc', 'surveillance', 's'], 1],
        //     [['eandc', 'surveillance', 'ssr', 'mode_a_c', 'a'], 2],
        //     [['eandc', 'surveillance', 'ssr', 'mode_s', 'a'], 7],
        //     [['eandc', 'surveillance', 'ads', 'c', 'a'], 4],
        //     [['eandc', 'surveillance', 'ads', 'b', 'a'], 12],
        //     [['t'], 4],
        //     [['da'], 4],
        //     [['cs'], 5],
        //     [['cl'], 5],
        //     [['r'], 152],
        //     [['de'], 4],
        //     [['teet'], 4],
        //     [['ada'], 4],
        //     [['adas', '0'], 4],
        //     [['adas', '1'], 4],
        //     [['oi'], 152],
        //     [['e'], 6],
        //     [['pob'], 5],
        //     [['er'], 4],
        //     [['se'], 4],
        //     [['j'], 10],
        //     [['d'], 38],
        //     [['acm'], 38],
        //     [['re'], 38],
        //     [['pn'], 19]
        // ]);

        castToBinary([
            [this.preimage.ai, 64],
            [this.preimage.fr, 8],
            [this.preimage.tof, 8],
            [this.preimage.noa, 16],
            [this.preimage.toa, 32],
            [this.preimage.wtc, 8],
            [this.preimage.eandc.rc_n_aae.s, 16],
            [this.preimage.eandc.surveillance.s, 8],
            [this.preimage.eandc.surveillance.ssr.mode_a_c.a, 16],
            [this.preimage.eandc.surveillance.ssr.mode_s.a, 64],
            [this.preimage.eandc.surveillance.ads.b.a[0], 64],
            [this.preimage.eandc.surveillance.ads.b.a[1], 64],
            [this.preimage.eandc.surveillance.ads.c.a, 32],
            [this.preimage.eandc.rc_n_aae.a[0], 64],
            [this.preimage.t, 32],
            [this.preimage.eandc.rc_n_aae.a[1], 64],
            [this.preimage.eandc.rc_n_aae.a[2], 64],
            [this.preimage.eandc.rc_n_aae.a[3], 64],
            [this.preimage.eandc.rc_n_aae.a[4], 64],
            [this.preimage.da, 32],
            [this.preimage.de, 32],
            [this.preimage.teet, 32],
            [this.preimage.ada, 32],
            [this.preimage.cs, 64],
            [this.preimage.cl, 64],
            [this.preimage.r[0], 64],
            [this.preimage.r[1], 64],
            [this.preimage.r[2], 64],
            [this.preimage.r[3], 64],
            [this.preimage.r[4], 64],
            [this.preimage.r[5], 64],
            [this.preimage.r[6], 64],
            [this.preimage.r[7], 64],
            [this.preimage.r[8], 64],
            [this.preimage.r[9], 64],
            [this.preimage.r[10], 64],
            [this.preimage.r[11], 64],
            [this.preimage.r[12], 64],
            [this.preimage.r[13], 64],
            [this.preimage.r[14], 64],
            [this.preimage.r[15], 64],
            [this.preimage.adas[0], 32],
            [this.preimage.adas[1], 32],
            [this.preimage.oi[0], 64],
            [this.preimage.oi[1], 64],
            [this.preimage.oi[2], 64],
            [this.preimage.oi[3], 64],
            [this.preimage.oi[4], 64],
            [this.preimage.oi[5], 64],
            [this.preimage.oi[6], 64],
            [this.preimage.oi[7], 64],
            [this.preimage.e, 64],
            [this.preimage.pob, 64],
            [this.preimage.er, 32],
            [this.preimage.se, 32],
            [this.preimage.j[0], 64],
            [this.preimage.j[1], 64],
            [this.preimage.d[0], 64],
            [this.preimage.d[1], 64],
            [this.preimage.d[2], 64],
            [this.preimage.d[3], 64],
            [this.preimage.acm[0], 64],
            [this.preimage.acm[1], 64],
            [this.preimage.acm[2], 64],
            [this.preimage.acm[3], 64],
            [this.preimage.re[0], 64],
            [this.preimage.re[1], 64],
            [this.preimage.re[2], 64],
            [this.preimage.re[3], 64],
            [this.preimage.pn[0], 64],
            [this.preimage.pn[1], 64],
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
        ]);
        //console.log(groups);

        let h = []
        for (let i = 1; i < 9; i++) {
            let hashPreimage = ''
            hashPreimage = hashPreimage.concat(groups[`h${i}_a`]);
            hashPreimage = hashPreimage.concat(groups[`h${i}_b`]);
            hashPreimage = hashPreimage.concat(groups[`h${i}_c`]);
            hashPreimage = hashPreimage.concat(groups[`h${i}_d`]);
            //console.log(i, hashPreimage);
            h.push(CryptoJS.enc.Hex.parse(BigInt('0b' + hashPreimage).toString(16)));
        }

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

        const segmentFields = (arrayOfFields) => {
            for (let field of arrayOfFields) {
                let segments = []
                while (field.length >= 19) {
                    let segment = field.slice(0, 19);
                    segments.push(segment)
                    field = field.slice(19, field.length);
                };
                if (segments.length == 1) {
                    field = segments[0]
                } else if (segments.length > 2) {
                    field = segments
                };
            };
        };

        const segmentFields = (object) => {
            const splitString = (string) => {
                let segments = [];
                while (string.length >= 19) {
                    let segment = string.slice(0, 19);
                    segments.push(segment);
                    string = string.slice(19, element.length);
                };
                if (segments.length == 1) {
                    string = segments[0];
                } else if (segments.length > 2) {
                    string = segments;
                };
                return segments;
            }

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
                        convertToASCII(value);
                    } else if (typeof value == 'string') {
                        object[key] = value.split('').map(c => c.charCodeAt(0)).join('');
                    };
                };
            };
        }

        segmentFields([
            this.preimage.ai,
            this.preimage.fr,
            this.preimage.tof,
            this.preimage.noa,
            this.preimage.toa,
            this.preimage.wtc,
            this.preimage.eandc.rc_n_aae.s,
            this.preimage.eandc.surveillance.s,
            this.preimage.eandc.surveillance.ssr.mode_a_c.a,
            this.preimage.eandc.surveillance.ssr.mode_s.a,
            this.preimage.eandc.surveillance.ads.b.a,
            this.preimage.eandc.surveillance.ads.c.a,
            this.preimage.eandc.rc_n_aae.a,
            this.preimage.t,
            this.preimage.da,
            this.preimage.de,
            this.preimage.teet,
            this.preimage.ada,
            this.preimage.cs,
            this.preimage.cl,
            this.preimage.r,
            this.preimage.adas[0],
            this.preimage.adas[1],
            this.preimage.oi,
            this.preimage.e,
            this.preimage.pob,
            this.preimage.er,
            this.preimage.se,
            this.preimage.j,
            this.preimage.d,
            this.preimage.acm,
            this.preimage.re,
            this.preimage.pn,
        ]);

        console.log(this.preimage);
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
                for (const [, value] of Object.entries(object)) {
                    if (typeof value === 'object') {
                        binAsHex(value);
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

    // To be done client side
    // async #createSignature(privateKey) {
    //     const [ signature_R, signature_S ] = await pycryptoJs.sign(privateKey, this.hash);
    //     return {
    //         R: signature_R,
    //         S: signature_S
    //     };
    // }

    async createProof(signature) {
        // this.signature = signature;
        // const program = fs.readFileSync(FPLProgramPath);
        // const provingKey = fs.readFileSync(FPLProvingKeyPath);
        // this.#preparePreimage()
        // console.log(this.preimage);
        // this.proof = await initialize().then(zokratesProvider => {
        //     const { witness, } = zokratesProvider.computeWitness({ program, abi: FPLabi }, [
        //         this.preimage,
        //         hashToIntPair(this.hash),
        //         this.signature,
        //         this.pk
        //     ]);
        //     console.log('test3')
        //     return zokratesProvider.generateProof(program, witness, provingKey);
        // });
        // return this.proof;
        this.signature = signature;
        const program = fs.readFileSync(TestProgramPath);
        const provingKey = fs.readFileSync(TestProvingKeyPath);
        this.proof = await initialize().then(zokratesProvider => {
            const { witness, } = zokratesProvider.computeWitness({ program, abi: Testabi }, [
                {
                    "a": "0x2a",
                    "b": {
                        "a": "42"
                    }
                },
                [
                    true,
                    false
                ],
                "42"
            ]);
            return zokratesProvider.generateProof(program, witness, provingKey);
        });
        return this.preimage;
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
            Array(122).fill('0') + ['65', '67', '75'].join(),
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
        // const [result] = await pycryptoJs.verify(this.pk, this.hash, signature.R, signature.S);
        // if (result) {
        //     this.signature = signature;
        //     const program = new Uint8Array(fs.readFileSync(ACKProgramPath));
        //     const provingKey = new Uint8Array(fs.readFileSync(ACKProvingKeyPath));
        //     this.proof = await initialize().then(zokratesProvider => {
        //         const { witness, } = zokratesProvider.computeWitness({ program, ACKabi }, [
        //             this.FPLSyncProof.preimage,
        //             hashToIntPair(this.FPLSyncProof.hash),
        //             this.FPLSyncProof.signature,
        //             this.FPLSyncProof.pk,
        //             hashToIntPair(this.hash),
        //             this.signature,
        //             this.pk
        //         ]);
        //         return zokratesProvider.generateProof(program, witness, provingKey);
        //     });
        // };
        // return null;
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
        return this.proof
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

const generateProof = async (h1, h2) => {
    const program = new Uint8Array(fs.readFileSync(programPath))
    const provingKey = new Uint8Array(fs.readFileSync(provingKeyPath))
    return await initialize().then(zokratesProvider => {
        const { witness, _output } = zokratesProvider.computeWitness({ program, abi }, [
            {
                "ai": Array(14).fill('7'),
                "fr": Array(2).fill('7'),
                "tof": Array(2).fill('7'),
                "noa": Array(4).fill('7'),
                "toa": Array(8).fill('7'),
                "wtc": Array(2).fill('7'),
                "e_and_c": {
                    "rc_n_aae": {
                        "s": Array(4).fill('7'),
                        "a": Array(98).fill('7')
                    },
                    "surveillance": {
                        "s": Array(2).fill('7'),
                        "ssr": {
                            "mode_a_c": {
                                "a": Array(4).fill('7')
                            },
                            "mode_s": {
                                "a": Array(14).fill('7')
                            }
                        },
                        "ads": {
                            "c": {
                                "a": Array(8).fill('7')
                            },
                            "b": {
                                "a": Array(24).fill('7')
                            }
                        }
                    }
                },
                "t": Array(8).fill('7'),
                "da": Array(8).fill('7'),
                "cs": Array(10).fill('7'),
                "cl": Array(10).fill('7'),
                "r": Array(304).fill('7'),
                "de": Array(8).fill('7'),
                "teet": Array(8).fill('7'),
                "ada": Array(8).fill('7'),
                "adas": [Array(8).fill('7'), Array(8).fill('7')],
                "oi": Array(152).fill('7'),
                "e": Array(12).fill('7'),
                "pob": Array(10).fill('7'),
                "er": Array(8).fill('7'),
                "se": Array(8).fill('7'),
                "j": Array(20).fill('7'),
                "d": Array(20).fill('7'),
                "acm": Array(76).fill('7'),
                "re": Array(76).fill('7'),
                "pn": Array(38).fill('7'),
            },
            [h1, h2]
        ]);
        return zokratesProvider.generateProof(program, witness, provingKey)
    })
}

const verifyFPLProof = async (proof) => {
    const FPLVerificationKey = new Uint8Array(fs.readFileSync(FPLVerificationKeyPath))
    return await initialize().then(zokratesProvider => {
        return zokratesProvider.verify(FPLVerificationKey, proof)
    })
}

const verifyACKProof = async (proof) => {
    const ACKVerificationKey = new Uint8Array(fs.readFileSync(ACKVerificationKeyPath))
    return await initialize().then(zokratesProvider => {
        return zokratesProvider.verify(ACKVerificationKey, proof)
    })
}

module.exports = {
    FPLSyncProof,
    ACKSyncProof,
    verifyFPLProof,
    verifyACKProof
}