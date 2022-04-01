const axios = require('axios');
const Web3 = require('web3')
const truffle_connect = require('../../connection/truffle_connect')
const endpoint = process.env.PROOF_SERVICE

function bootstrap() {
    console.log("bootstrap() started")
    console.log(endpoint)

    const pollingDelay = 1000;
    let files = null;
    let state = null;

    return new Promise(function (resolve, reject) {

        async function pollForVerifier() {
            console.log("Polling...")
            try {   
                const response = await axios.get(`http://${endpoint}/contracts`)
                state = response.data.state
                files = response.data.files;
                console.log('bootstrap', files)
                continueCheckingForCompletion();
            } catch (error) {
                if (error.response) {
                    reject(error)
                } else if (error.request) {
                    continueCheckingForCompletion()
                } else {
                    reject(error)
                }
            };
        }

        function continueCheckingForCompletion() {
            if (state !== "finished") {
                setTimeout(pollForVerifier, pollingDelay);
            }
            else {
                console.log("bootstrap() finished")
                resolve({ files });
            }
        }

        continueCheckingForCompletion();
    });
}

async function getKeys() {
    const response = await axios.get(`http://${endpoint}/keys`)
    const {sk, pk} = response.data
    return {sk, pk}
}

async function getFPLProof(FPL, privateKey) {
    let response = null
    try {
        response = await axios.post(`http://${endpoint}/prove/fpl`, {
            fpl: FPL,
            sk: privateKey
        })
    } catch(e) {
        return response
    }
    return response.data
}

async function getACKProof(FPLProofInputs, privateKey) {
    let response = null
    try {
        response = await axios.post(`http://${endpoint}/prove/ack`, {
            proof: FPLProofInputs,
            sk: privateKey
        })
    } catch(e) {
        return response
    }
    return response.data
}

async function getFormattedProof(type, proof, inputs) {

    // const BNString = (string) => {
    //     //return '0x' + Web3Utils.toBN(string).toString(16, 32)
    // }

    // const stringsToBNString = (data) => {
    //     if (Array.isArray(data)) {
    //         for (const [index, element] of data.entries()) {
    //             if (typeof element === 'string') {
    //                 data[index] = BNString(element)
    //             } else {
    //                 stringsToBNString(element)
    //             }
    //         }
    //     } else if (typeof data === 'object') {
    //         for (const [key, value] of Object.entries(data)) {
    //             if (typeof value === 'string') {
    //                 data[key] = BNString(value)
    //             } else {
    //                 stringsToBNString(value)
    //             }
    //         }
    //     }
    // }
    
    let response = null
    try {
        response = await axios.post(`http://${endpoint}/format/${type}`, {
            proof,
            inputs
        })
        const {hash, sig, pk} = response.data.inputs
        response.data.inputs = [...pk, ...sig, hash]
    } catch(e) {
        return [response, e]
    } 
    return [response.data, undefined]
}

module.exports = {
    bootstrap,
    getKeys,
    getFPLProof,
    getACKProof,
    getFormattedProof
}