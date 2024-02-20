// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainDBEvents {
    // TODO: figure out indexing
    event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash);
    event InterchainEntryVerified(
        address module, uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash
    );

    event InterchainVerificationRequested(uint256 destChainId, uint256 dbNonce, address[] srcModules);
}
