// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainDB {
    /// @notice Write data to the Interchain DataBase, and verify it on the destination chain
    /// using the provided Interchain Modules.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// @param destChainId  The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase
    function writeData(uint256 destChainId, address[] memory srcModules, bytes32 dataHash) external payable;

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @param destChainId  The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function getInterchainFee(uint256 destChainId, address[] memory srcModules) external view returns (uint256);

    /// @notice Verify the data written on specific source chain by a specific writer
    /// using the provided Interchain Modules.
    /// @dev The returned array of timestamps has the same length as the `dstModules` array,
    /// and its values are the block timestamps at which the data hash was confirmed by the corresponding module.
    /// Note: zero value indicates that the module has not confirmed the data hash.
    /// @param srcChainId   The chain id of the source chain
    /// @param srcWriter    The address of the writer on the source chain
    /// @param dataHash     The hash of the data written on the source chain
    /// @param dstModules   The destination chain addresses of the Interchain Modules to use for verification
    /// @return moduleConfirmedAt   The block timestamp at which the data hash was confirmed by each module,
    ///                             or zero if the module has not confirmed the data hash.
    function readData(
        uint256 srcChainId,
        bytes32 srcWriter,
        bytes32 dataHash,
        address[] memory dstModules
    )
        external
        view
        returns (uint256[] memory moduleConfirmedAt);
}
