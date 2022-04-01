// SPDX-License-Identifier: CC0
pragma solidity ^0.8.0;

import "./lib/MerkleTreeSHA256.sol";
import "./IShield.sol";
import "./IVerifier.sol";

contract Shield is IShield, MerkleTreeSHA256 {
    // CONTRACT INSTANCES:
    IFPLVerifier private FPLverifier; // the FPL verification smart contract
    IACKVerifier private ACKverifier; // the ACK verification smart contract

    // FUNCTIONS:
    constructor(address _FPLverifier, address _ACKverifier, uint _treeHeight) public MerkleTreeSHA256(_treeHeight) {
        FPLverifier = IFPLVerifier(_FPLverifier);
        ACKverifier = IACKVerifier(_ACKverifier);
    }

    // returns the FPL verifier contract address that this shield contract uses for FPL proof verification
    function getFPLVerifier() external override view returns (address) {
        return address(FPLverifier);
    }

    // returns the ACK verifier contract address that this shield contract uses for ACK proof verification
    function getACKVerifier() external override view returns (address) {
        return address(ACKverifier);
    }

    function verifyAndPushFPL(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[6] memory input,
        bytes32 _newCommitment
    ) external override returns (bool) {

        // verify the proof
        bool result = FPLverifier.verifyProof(a, b, c, input);
        require(result, "The proof failed verification in the verifier contract");

        // update contract states
        insertLeaf(_newCommitment); // recalculate the root of the merkleTree as it's now different
        return true;
    }

    function verifyAndPushACK(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[12] memory input,
        bytes32 _newCommitment
    ) external override returns (bool) {

        // verify the proof
        bool result = ACKverifier.verifyProof(a, b, c, input);
        require(result, "The proof failed verification in the verifier contract");

        // update contract states
        insertLeaf(_newCommitment); // recalculate the root of the merkleTree as it's now different
        return true;
    }
}