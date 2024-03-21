// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainModuleEvents {
    event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash);

    event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash);
}
