
// SPDX-License-Identifier: AML
// 
// Copyright 2017 Christian Reitwiessner
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

// 2019 OKIMS

pragma solidity ^0.8.0;
import "./Pairing.sol";

contract FPLVerifier {

    using Pairing for *;

    uint256 constant SNARK_SCALAR_FIELD = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    uint256 constant PRIME_Q = 21888242871839275222246405745257275088696311157297823662689037894645226208583;

    struct VerifyingKey {
        Pairing.G1Point alfa1;
        Pairing.G2Point beta2;
        Pairing.G2Point gamma2;
        Pairing.G2Point delta2;
        Pairing.G1Point[7] IC;
    }

    struct Proof {
        Pairing.G1Point A;
        Pairing.G2Point B;
        Pairing.G1Point C;
    }

    function verifyingKey() internal pure returns (VerifyingKey memory vk) {
        vk.alfa1 = Pairing.G1Point(uint256(19321890128367905184434241644252805434982917455888436181805793028319383774437), uint256(4386775038361489281590108606690518227978437370878316028174752722436827759846));
        vk.beta2 = Pairing.G2Point([uint256(8401791724555996120565831114396790194520668036439398896819840040297472214510), uint256(13249245637011303111716086687658136990368589272788322911132208331678690090471)], [uint256(16742713093044928053603875235800701671871195035716684543973382116463006103382), uint256(15769966133244087532402098219333178484165172254244126374181424345679108831970)]);
        vk.gamma2 = Pairing.G2Point([uint256(11532654318723019872555547845181276135648708258850494041207898030884092940422), uint256(16650784525006791260477630620879223882575881798583617007468414658316505284298)], [uint256(9591297511777709180964881615524315888767055987752884396638648276336308403806), uint256(19612665326469428162394338333275772529779462609229281169111191931228845829435)]);
        vk.delta2 = Pairing.G2Point([uint256(6090293905091741500680178602146073103061984102731882681290441422695810722507), uint256(17158478396609353009648048095717770489699928726518731864286028880985010422349)], [uint256(21485404515319583993086912272549588386529855217747414411454503296081263987911), uint256(5096899510847759889370613146034874261833290218355620939673295670241601006021)]);   
        vk.IC[0] = Pairing.G1Point(uint256(16662534240970033836209704398539085728857626071359723072695066827428173171933), uint256(17214954667206756160629211903950556265780752211791271325898932407360388799176));   
        vk.IC[1] = Pairing.G1Point(uint256(7729321909017916710666300265121077873439567768759185755789962308246261067918), uint256(9502225879459181729459066119665777009012293724547648558652028229809871972648));   
        vk.IC[2] = Pairing.G1Point(uint256(12871826694067987631294848642423847867888989827308945387485253976455325823526), uint256(14340026469355098831952084484491303354591469875782046747477194535912949018771));   
        vk.IC[3] = Pairing.G1Point(uint256(20664558432514978866061228724034118980908084963129363225398721803164671864805), uint256(17482298650183674563231796391060094994248412211732239699208942729097950618115));   
        vk.IC[4] = Pairing.G1Point(uint256(16425079678664072063004681530801109502075856012814947379408924624590395795141), uint256(18847659541992561115862789442298296049980418045259607746380665068561141865198));   
        vk.IC[5] = Pairing.G1Point(uint256(3348787764364562820162381672997198062673726478024063795095895948137983642916), uint256(17606568328608858444160928045361669247526642343203438461226739850957774567846));   
        vk.IC[6] = Pairing.G1Point(uint256(5958823451454246609945000633843062942142813763081318711745232454815612279953), uint256(727591298227723475671478867460540643596293205422677321658775831831722368768));
    }
    
    /*
     * @returns Whether the proof is valid given the hardcoded verifying key
     *          above and the public inputs
     */
    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[6] memory input
    ) public view returns (bool r) {

        Proof memory proof;
        proof.A = Pairing.G1Point(a[0], a[1]);
        proof.B = Pairing.G2Point([b[0][0], b[0][1]], [b[1][0], b[1][1]]);
        proof.C = Pairing.G1Point(c[0], c[1]);

        VerifyingKey memory vk = verifyingKey();

        // Compute the linear combination vk_x
        Pairing.G1Point memory vk_x = Pairing.G1Point(0, 0);

        // Make sure that proof.A, B, and C are each less than the prime q
        require(proof.A.X < PRIME_Q, "verifier-aX-gte-prime-q");
        require(proof.A.Y < PRIME_Q, "verifier-aY-gte-prime-q");

        require(proof.B.X[0] < PRIME_Q, "verifier-bX0-gte-prime-q");
        require(proof.B.Y[0] < PRIME_Q, "verifier-bY0-gte-prime-q");

        require(proof.B.X[1] < PRIME_Q, "verifier-bX1-gte-prime-q");
        require(proof.B.Y[1] < PRIME_Q, "verifier-bY1-gte-prime-q");

        require(proof.C.X < PRIME_Q, "verifier-cX-gte-prime-q");
        require(proof.C.Y < PRIME_Q, "verifier-cY-gte-prime-q");

        // Make sure that every input is less than the snark scalar field
        for (uint256 i = 0; i < input.length; i++) {
            require(input[i] < SNARK_SCALAR_FIELD,"verifier-gte-snark-scalar-field");
            vk_x = Pairing.plus(vk_x, Pairing.scalar_mul(vk.IC[i + 1], input[i]));
        }

        vk_x = Pairing.plus(vk_x, vk.IC[0]);

        return Pairing.pairing(
            Pairing.negate(proof.A),
            proof.B,
            vk.alfa1,
            vk.beta2,
            vk_x,
            vk.gamma2,
            proof.C,
            vk.delta2
        );
    }
}
