// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, InterchainBatchLib} from "../../contracts/libs/InterchainBatch.sol";

contract InterchainBatchLibHarness {
    function constructLocalBatch(uint256 dbNonce, bytes32 batchRoot) external view returns (InterchainBatch memory) {
        return InterchainBatchLib.constructLocalBatch(dbNonce, batchRoot);
    }

    function decodeVersionedBatchFromMemory(bytes memory versionedBatch)
        external
        view
        returns (uint16, InterchainBatch memory)
    {
        return InterchainBatchLib.decodeVersionedBatchFromMemory(versionedBatch);
    }

    function decodeVersionedBatch(bytes calldata versionedBatch)
        external
        pure
        returns (uint16, InterchainBatch memory)
    {
        return InterchainBatchLib.decodeVersionedBatch(versionedBatch);
    }

    function encodeVersionedBatch(
        uint16 dbVersion,
        InterchainBatch memory batch
    )
        external
        pure
        returns (bytes memory)
    {
        return InterchainBatchLib.encodeVersionedBatch(dbVersion, batch);
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
