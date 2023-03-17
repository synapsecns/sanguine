// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TREE_DEPTH } from "../../contracts/libs/Constants.sol";

import { AttestationProofGenerator } from "./proof/AttestationProofGenerator.t.sol";
import { HistoricalProofGenerator } from "./proof/HistoricalProofGenerator.t.sol";

abstract contract SynapseProofs {
    HistoricalProofGenerator internal originGen;
    AttestationProofGenerator internal summitGen;

    constructor() {
        clear();
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseProofs() external {}

    /// @notice Clears proof generators
    function clear() public {
        originGen = new HistoricalProofGenerator();
        summitGen = new AttestationProofGenerator();
    }

    function insertMessage(bytes memory message) public {
        originGen.insert(keccak256(message));
    }

    function getLatestProof(uint256 index) public view returns (bytes32[TREE_DEPTH] memory proof) {
        return originGen.getLatestProof(index);
    }

    function getProof(uint256 index, uint256 count)
        public
        view
        returns (bytes32[TREE_DEPTH] memory)
    {
        return originGen.getProof(index, count);
    }

    function getRoot(uint256 count) public view returns (bytes32) {
        return originGen.getRoot(count);
    }

    function acceptSnapshot(bytes[] memory snapshotStates) public {
        summitGen.acceptSnapshot(snapshotStates);
    }

    function genSnapshotProof(uint256 index) public view returns (bytes32[] memory) {
        return summitGen.generateProof(index);
    }

    function getSnapshotRoot() public view returns (bytes32) {
        return summitGen.root();
    }

    function getSnapshotHeight() public view returns (uint8) {
        // Extra element in the proof list is "right sub-leaf of Origin State"
        return uint8(summitGen.height() + 1);
    }
}
