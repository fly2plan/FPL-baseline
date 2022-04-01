// SPDX-License-Identifier: CC0
pragma solidity ^0.8.0;

/**
@title IVerifier
@dev Example Verifier Implementation
*/
interface IFPLVerifier {

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[6] memory input
    ) external returns (bool r);

}

interface IACKVerifier {

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[12] memory input
    ) external returns (bool r);
}