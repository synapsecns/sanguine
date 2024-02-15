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

    error InterchainDB__ConflictingEntries(InterchainEntry existingEntry, bytes32 dataHash);
    error InterchainDB__EntryDoesNotExist(address writer, uint256 writerNonce);
    error InterchainDB__IncorrectFeeAmount(uint256 actualFee, uint256 expectedFee);
    error InterchainDB__NoModulesSpecified();
    error InterchainDB__SameChainId();

    /// @notice Write data to the Interchain DataBase as a new entry.
    /// Note: there are no guarantees that this entry will be available for reading on any of the remote chains.
    /// Use `verifyEntry` to ensure that the entry is available for reading on the destination chain.
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @return writerNonce The writer-specific nonce of the written entry
    function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce);

    /// @notice Request the given Interchain Modules to verify the already written entry on the destination chain.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// Note: this method is permissionless, and anyone can request verification for any entry.
    /// @dev Will revert if the entry with the given nonce does not exist.
    /// @param destChainId   The chain id of the destination chain
    /// @param writer        The address of the writer on the source chain
    /// @param writerNonce   The nonce of the writer on the source chain
    /// @param srcModules    The source chain addresses of the Interchain Modules to use for verification
    function requestVerification(
        uint256 destChainId,
        address writer,
        uint256 writerNonce,
        address[] memory srcModules
    )
        external
        payable;

    /// @notice Write data to the Interchain DataBase,
    /// and request the given Interchain Modules to verify it on the destination chain.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// Note: additional verification for the same entry could be later done using `requestVerification`.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param destChainId  The chain id of the destination chain
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    /// @return writerNonce The writer-specific nonce of the written entry
    function writeEntryWithVerification(
        uint256 destChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint256 writerNonce);

    /// @notice Allows the Interchain Module to verify the entry coming from a remote source chain.
    /// @param entry        The Interchain Entry to confirm
    function verifyEntry(InterchainEntry memory entry) external;

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @dev Will revert if the empty array of modules is provided.
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
    /// and verify it on the destination chain using the provided Interchain Module.
    /// Note: returned zero value indicates that the module has not verified the entry.
    /// @param entry        The Interchain Entry to read
    /// @param dstModule    The destination chain addresses of the Interchain Modules to use for verification
    /// @return moduleVerifiedAt   The block timestamp at which the entry was verified by the module,
    ///                             or zero if the module has not verified the entry.
    function readEntry(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        returns (uint256 moduleVerifiedAt);
}
