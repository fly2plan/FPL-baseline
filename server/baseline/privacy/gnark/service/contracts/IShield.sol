// SPDX-License-Identifier: CC0
pragma solidity ^0.8.0;
import "./IVerifier.sol";

interface IShield {

    function getFPLVerifier() external view returns (address);
    function getACKVerifier() external view returns (address);

    function verifyAndPushFPL(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[6] memory input,
        bytes32 _newCommitment
    ) external returns (bool);

    function verifyAndPushACK(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[12] memory input,
        bytes32 _newCommitment
    ) external returns (bool);
}