// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainDB {
    /// @notice Struct representing an entry in the Interchain DataBase
    /// @param srcChainId   The chain id of the source chain
    /// @param srcWriter    The address of the writer on the source chain
    /// @param writerNonce  The nonce of the writer on the source chain
    /// @param dataHash     The hash of the data written on the source chain
    struct InterchainEntry {
        uint256 srcChainId;
        bytes32 srcWriter;
        uint256 writerNonce;
        bytes32 dataHash;
    }

    /// @notice Write data to the Interchain DataBase, and verify it on the destination chain
    /// using the provided Interchain Modules.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @param destChainId  The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function writeEntry(bytes32 dataHash, uint256 destChainId, address[] memory srcModules) external payable;

    /// @notice Allows the Interchain Module to mark the entry as confirmed.
    /// @param entry        The Interchain Entry to confirm
    function confirmEntry(InterchainEntry memory entry) external;

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @param destChainId  The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function getInterchainFee(uint256 destChainId, address[] memory srcModules) external view returns (uint256);

    /// @notice Get the Interchain Entry by the writer and the writer nonce.
    /// @dev Will revert if the entry with the given nonce does not exist.
    /// @param writer       The address of the writer on this chain
    /// @param writerNonce  The nonce of the writer's entry on this chain
    function getEntry(address writer, uint256 writerNonce) external view returns (InterchainEntry memory);

    /// @notice Get the nonce of the writer on this chain.
    /// @param writer       The address of the writer on this chain
    function getWriterNonce(address writer) external view returns (uint256);

    /// @notice Read the data written on specific source chain by a specific writer,
    /// and verify it on the destination chain using the provided Interchain Modules.
    /// @dev The returned array of timestamps has the same length as the `dstModules` array,
    /// and its values are the block timestamps at which the entry was confirmed by the corresponding module.
    /// Note: zero value indicates that the module has not confirmed the entry.
    /// @param entry        The Interchain Entry to read
    /// @param dstModules   The destination chain addresses of the Interchain Modules to use for verification
    /// @return moduleConfirmedAt   The block timestamp at which the entry was confirmed by each module,
    ///                             or zero if the module has not confirmed the entry.
    function readEntry(
        InterchainEntry memory entry,
        address[] memory dstModules
    )
        external
        view
        returns (uint256[] memory moduleConfirmedAt);
}
