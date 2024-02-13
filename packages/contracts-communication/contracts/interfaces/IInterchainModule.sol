// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IInterchainDB} from "./IInterchainDB.sol";

/// @notice Every Module may opt a different method to confirm the verified entries on destination chain,
/// therefore this is not a part of a common interface.
interface IInterchainModule {
    /// @notice Verify an entry in the Interchain DataBase.
    /// Note: a fee is paid to the module for verification, and could be retrieved by using `getModuleFee`.
    /// Note: this will eventually trigger `InterchainDB.confirmEntry(entry)` function on destination chain,
    /// with no guarantee of ordering.
    /// @dev Could be only called by the Interchain DataBase contract.
    /// @param destChainId  The chain id of the destination chain
    /// @param entry        The entry to verify
    function verifyEntry(uint256 destChainId, IInterchainDB.InterchainEntry memory entry) external payable;

    /// @notice Get the Module fee for verifying an entry on the specified destination chain
    /// @param destChainId  The chain id of the destination chain
    function getModuleFee(uint256 destChainId) external view returns (uint256);
}
