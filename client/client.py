from flask import Flask, request

from os import urandom
from math import ceil, log2

from zokrates_pycrypto.eddsa import PrivateKey, PublicKey
from zokrates_pycrypto.field import FQ
from zokrates_pycrypto.babyjubjub import JUBJUB_L



def get_rand_n():
    mod = JUBJUB_L
    nbytes = ceil(ceil(log2(mod)) / 8)
    rand_n = int.from_bytes(urandom(nbytes), "little")
    return rand_n

def sk_from_seed(key_seed):
    key = FQ(key_seed)
    sk = PrivateKey(key)
    return sk

app = Flask(__name__)

@app.route('/keys', methods=['GET'])
def get_keys():
    key_seed = get_rand_n()
    sk = sk_from_seed(key_seed)
    pk = PublicKey.from_private(sk)
    return {'sk': f'{key_seed}', 'pk': [f'{pk.p.x.n}', f'{pk.p.y.n}']}


@app.route('/sign/<message>', methods=['POST'])
def sign_message(message):
    request_data = request.get_json()
    message = message.replace('"', '')
    message = message + ''.join(['0' for _ in range(64)])
    sk = sk_from_seed(int(request_data['sk']))
    sig = sk.sign(bytes.fromhex(message))
    sig_R, sig_S = sig
    return {'R': [f'{sig_R.x}', f'{sig_R.y}'], 'S': f'{sig_S}'}
