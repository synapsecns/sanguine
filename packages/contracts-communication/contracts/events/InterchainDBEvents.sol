// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainDBEvents {
    // TODO: figure out indexing
    event InterchainEntryWritten(uint64 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash);

    event InterchainBatchVerified(address module, uint64 srcChainId, uint256 dbNonce, bytes32 batchRoot);

    event InterchainBatchVerificationRequested(
        uint64 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules
    );
}
