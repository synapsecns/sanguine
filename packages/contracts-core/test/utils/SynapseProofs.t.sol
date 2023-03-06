// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../contracts/libs/Constants.sol";
import { HistoricalProofGenerator } from "./proof/HistoricalProofGenerator.t.sol";

abstract contract SynapseProofs {
    HistoricalProofGenerator internal proofGen;

    constructor() {
        clear();
    }

    /// @notice Clears HistoricalProofGenerator
    function clear() public {
        proofGen = new HistoricalProofGenerator();
    }

    function insertMessage(bytes memory message) public {
        proofGen.insert(keccak256(message));
    }

    function getLatestProof(uint256 index)
        public
        view
        returns (bytes32[ORIGIN_TREE_DEPTH] memory proof)
    {
        return proofGen.getLatestProof(index);
    }

    function getProof(uint256 index, uint256 count)
        public
        view
        returns (bytes32[ORIGIN_TREE_DEPTH] memory)
    {
        return proofGen.getProof(index, count);
    }

    function getRoot(uint256 count) public view returns (bytes32) {
        return proofGen.getRoot(count);
    }
}
