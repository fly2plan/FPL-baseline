// This file is MIT Licensed.
//
// Copyright 2017 Christian Reitwiessner
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
pragma solidity ^0.8.0;

import "./Pairing.sol";

contract ACKVerifier {
    using Pairing for *;
    struct VerifyingKey {
        Pairing.G1Point alpha;
        Pairing.G2Point beta;
        Pairing.G2Point gamma;
        Pairing.G2Point delta;
        Pairing.G1Point[] gamma_abc;
    }
    struct Proof {
        Pairing.G1Point a;
        Pairing.G2Point b;
        Pairing.G1Point c;
    }
    function verifyingKey() pure internal returns (VerifyingKey memory vk) {
        vk.alpha = Pairing.G1Point(uint256(0x10bf0fec471841bce2d8b9e2a6dc3b2f2431e672dd5751997059ff1d8f16a627), uint256(0x21bfcf7d98faceaac6cbd80a4a8fb06019373701ac7f9f08c3f466f9b9b367c1));
        vk.beta = Pairing.G2Point([uint256(0x12706d317ff5f71ff4f58aab46ea6535d8386b3a3c65788c3d6300499f41afb9), uint256(0x1c9145c0ab7f77a68ee62fbf077a7f543b7f98c6139131f5138f1e01fed2c3e9)], [uint256(0x0a8730dc43b5c339819951b42be97199a0ae25c03be0ddf770d351c59da543c1), uint256(0x05902dd2e1a7875c772d12b8d4f5f5e553e614dd7c6eb9819191ccae0b959d71)]);
        vk.gamma = Pairing.G2Point([uint256(0x196c4c588bdd33d9de68092462af5578767100b97a3f9734cb63619495d3d742), uint256(0x1c55dfef540672f2ec0cc2586f4c8a1cc9010c73d7f6ae05e7a49a65140108e8)], [uint256(0x181a46abcda8b25ee3c42629278574e55ab681285be7214727c9e0165068fb67), uint256(0x1df9aef27443415c50519dbf081ffb9ffe14d07f1b8323cab7074c14456084eb)]);
        vk.delta = Pairing.G2Point([uint256(0x29bbe3b661eb40f1ea590f1cb0b0ac2341105dda1fe415fe7a4007c11c10b8ed), uint256(0x161278ef2bc17f7a75f1a5983baa8de7f02afa601338d8bcda900092e2869f6a)], [uint256(0x169a2f332aed48a71611c15fa2c7f3fc4d87e19f7c9644a16bd36feac747454e), uint256(0x0152f2dc356ceee9b635d22df38497249f8d6762cad9a5c155e1ea6dc9e42f3a)]);
        vk.gamma_abc = new Pairing.G1Point[](5);
        vk.gamma_abc[0] = Pairing.G1Point(uint256(0x1f3952d8f1423f8952d2957957986b1efdfec7c4c71ddfbaf56087b005c69d68), uint256(0x2a3227d0ce0d57531132b43a80bafd1632f60ca1e80b185f3ec3d3d9dfd74449));
        vk.gamma_abc[1] = Pairing.G1Point(uint256(0x017ca30353761d5a4f3d000564fb06a1fdc04d1b7fe6efd536cff979c2e35892), uint256(0x060fa1c9927c740cb366708dabcfe935dc05793667507e9dcba5e6908ed12b55));
        vk.gamma_abc[2] = Pairing.G1Point(uint256(0x265a428fb8a58cbf0788ebd8a9a2ccb37a49a2a4e57978820aa72ebd5620f63e), uint256(0x2f5a6f15ab808daeb69ecffc6571c299ba6a46d779a97944d7920af88a65a641));
        vk.gamma_abc[3] = Pairing.G1Point(uint256(0x18572c6765d85cdb1fcddfce5e55de9ef2b49411ea1d6788ba5c243f03510b9b), uint256(0x1ea46d52fbbb4909c0befde2a4d7f13f14ebecd6651f20d6c12eabc469a3dfba));
        vk.gamma_abc[4] = Pairing.G1Point(uint256(0x004dfe5f91715e74eaec2bbac25f024def95f31fc0aff12a78d817e0f6ad4913), uint256(0x06704e8702b874579789568f22e91b4af9983b42072bab247089c1397d9af717));
    }
    function verify(uint[] memory input, Proof memory proof) internal view returns (uint) {
        uint256 snark_scalar_field = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
        VerifyingKey memory vk = verifyingKey();
        require(input.length + 1 == vk.gamma_abc.length);
        // Compute the linear combination vk_x
        Pairing.G1Point memory vk_x = Pairing.G1Point(0, 0);
        for (uint i = 0; i < input.length; i++) {
            require(input[i] < snark_scalar_field);
            vk_x = Pairing.addition(vk_x, Pairing.scalar_mul(vk.gamma_abc[i + 1], input[i]));
        }
        vk_x = Pairing.addition(vk_x, vk.gamma_abc[0]);
        if(!Pairing.pairingProd4(
             proof.a, proof.b,
             Pairing.negate(vk_x), vk.gamma,
             Pairing.negate(proof.c), vk.delta,
             Pairing.negate(vk.alpha), vk.beta)) return 1;
        return 0;
    }
    function verifyTx(
            Proof memory proof, uint[4] memory input
        ) public view returns (bool r) {
        uint[] memory inputValues = new uint[](4);
        
        for(uint i = 0; i < input.length; i++){
            inputValues[i] = input[i];
        }
        if (verify(inputValues, proof) == 0) {
            return true;
        } else {
            return false;
        }
    }
}
