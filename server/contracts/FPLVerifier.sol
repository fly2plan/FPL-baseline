// This file is MIT Licensed.
//
// Copyright 2017 Christian Reitwiessner
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
pragma solidity ^0.8.0;

import "./Pairing.sol";

contract FPLVerifier {
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
        vk.alpha = Pairing.G1Point(uint256(0x0f2415ca9ebef8e71ee5e047abc7cb7ec6f8252b50fca4be0e990149a23f232f), uint256(0x0fa2dd0174172a31a8ecdb2ccce86ca6a105a27e68fefd34015c17e7d97ba68b));
        vk.beta = Pairing.G2Point([uint256(0x1e800cc7541d7257ccdc8b0ea293ca2b79ca3d4b5f58c96b64d371e505309aae), uint256(0x300e26bab47631448f7fd6ac35eecc27019d4a5164a62556961201e41f79cc5c)], [uint256(0x2d08502e5de1043a3e085126ad27222b62f865ac1929fa67300f9a618927b1f5), uint256(0x051e1ae6d1d7588b683ae4edfe10653ee0077a39424dde8a5890fc33f65ac2ed)]);
        vk.gamma = Pairing.G2Point([uint256(0x2618c188949a6fa6250d982b3bf9d220c978b125a57a238a2329679469c3ab08), uint256(0x01b4abe7eb9173ed34f42c9f8c84e259e869a073776eee9d50c27305f9cc4b57)], [uint256(0x1ed378dc1376cbf62c33c23623e0f997f8fa39e8c92d6884fa8ee3fbfc0b1d4a), uint256(0x252692af68db8195b919c22e379c2dd9e474f996faeefcefa37909ea1dad8087)]);
        vk.delta = Pairing.G2Point([uint256(0x21bf0a1a0ed4040c07e97e1fc7cd25eb69dea894f48f66fb4a2a6701448a61ee), uint256(0x1c043046263e3b4a9004f7c80922ea9c32582f56b3d8ae53cb2355d851a8b8e0)], [uint256(0x05fcb41f482a57862e6f5e00da5827aeef083b0744a96269e26f90c4a28fcc9f), uint256(0x178664d3742b65cab6e9dd6bc491d7d22ae552ffcf7cfe99573db13a6b96f614)]);
        vk.gamma_abc = new Pairing.G1Point[](3);
        vk.gamma_abc[0] = Pairing.G1Point(uint256(0x060bd7288653b54c57f2add5a00b93ed752c93b471548cb88eae9a0568424eb2), uint256(0x0ed9960e1ad780b7da68d88821f16d5f514052784c4556a6faaf3e6997e7e434));
        vk.gamma_abc[1] = Pairing.G1Point(uint256(0x28e700d587715a426a9fd8f2dbff00edc90e19154e2ba83e40a37c3dab425278), uint256(0x0ecdada252c6428dc9c0c2d4b7e4aa3978d9cc6fbf51ac56aad3061d9e421cb6));
        vk.gamma_abc[2] = Pairing.G1Point(uint256(0x246780c36bde457bdff77c0ee409aa6fba6521c2ba008842bf0c44f5da70cab8), uint256(0x09e3e7e6be95b5e57f4a64df4cf2ffe0890df684d2cdb1e2bdd3243ab1f790f1));
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
            Proof memory proof, uint[2] memory input
        ) public view returns (bool r) {
        uint[] memory inputValues = new uint[](2);
        
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
