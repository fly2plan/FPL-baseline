// SPDX-License-Identifier: CC0
pragma solidity ^0.8.0;
import "./Pairing.sol";

library S {
    struct Proof {
        Pairing.G1Point a;
        Pairing.G2Point b;
        Pairing.G1Point c;
    }
}

/**
@title IVerifier
@dev Example Verifier Implementation
*/
interface IFPLVerifier {

    function verifyTx(
        S.Proof memory proof,
        uint[2] memory input
    ) external returns (bool r);

}

interface IACKVerifier {

    function verifyTx(
        S.Proof memory proof,
        uint[4] memory input
    ) external returns (bool r);
}