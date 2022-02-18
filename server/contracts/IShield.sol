// SPDX-License-Identifier: CC0
pragma solidity ^0.8.0;
import "./IVerifier.sol";

interface IShield {

    function getFPLVerifier() external view returns (address);
    function getACKVerifier() external view returns (address);

    function verifyAndPushFPL(
        S.Proof memory proof,
        uint[2] memory input,
        bytes32 _newCommitment
    ) external returns (bool);

    function verifyAndPushACK(
        S.Proof memory proof,
        uint[4] memory input,
        bytes32 _newCommitment
    ) external returns (bool);
}