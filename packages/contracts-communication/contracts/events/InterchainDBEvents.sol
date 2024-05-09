// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainDBEvents {
    /// @notice Emitted when a local entry (a specific digest by the specific writer) is written to the database.
    /// @param dbNonce      The nonce of written entry.
    /// @param srcWriter    The address of the writer.
    /// @param digest       The written data digest.
    /// @param entryValue   The value of the written entry: keccak256(abi.encode(srcWriter, digest)).
    event InterchainEntryWritten(uint64 indexed dbNonce, bytes32 indexed srcWriter, bytes32 digest, bytes32 entryValue);

    /// @notice Emitted when a remote entry is verified by the Interchain Module.
    /// @param module       The address of the Interchain Module that verified the entry.
    /// @param srcChainId   The ID of the source chain.
    /// @param dbNonce      The nonce of the verified entry.
    /// @param entryValue   The value of the verified entry: keccak256(abi.encode(srcWriter, digest)).
    event InterchainEntryVerified(
        address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 entryValue
    );

    /// @notice Emitted when a local entry is requested to be verified on a remote chain
    /// using the set of Interchain Modules.
    /// @param dstChainId   The ID of the destination chain.
    /// @param dbNonce      The nonce of the entry to be verified.
    /// @param srcModules   The addresses of the Interchain Modules that will verify the entry.
    event InterchainEntryVerificationRequested(uint64 indexed dstChainId, uint64 indexed dbNonce, address[] srcModules);
}
