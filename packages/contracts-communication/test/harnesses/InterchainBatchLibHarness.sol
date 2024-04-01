// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, InterchainBatchLib} from "../../contracts/libs/InterchainBatch.sol";

contract InterchainBatchLibHarness {
    function constructLocalBatch(uint64 dbNonce, bytes32 batchRoot) external view returns (InterchainBatch memory) {
        return InterchainBatchLib.constructLocalBatch(dbNonce, batchRoot);
    }

    function encodeBatch(InterchainBatch memory batch) external pure returns (bytes memory) {
        return InterchainBatchLib.encodeBatch(batch);
    }

    function decodeBatch(bytes calldata batch) external pure returns (InterchainBatch memory) {
        return InterchainBatchLib.decodeBatch(batch);
    }

    function decodeBatchFromMemory(bytes memory batch) external pure returns (InterchainBatch memory) {
        return InterchainBatchLib.decodeBatchFromMemory(batch);
    }

    function batchKey(InterchainBatch memory batch) external pure returns (bytes32) {
        return InterchainBatchLib.batchKey(batch);
    }
}
