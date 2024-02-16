// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IInterchainDB} from "./IInterchainDB.sol";
import {InterchainEntry} from "../libs/InterchainEntry.sol";

/// @notice Every Module may opt a different method to confirm the verified entries on destination chain,
/// therefore this is not a part of a common interface.
interface IInterchainModule {
    /// @notice Request the verification of an entry in the Interchain DataBase by the module.
    /// Note: a fee is paid to the module for verification, and could be retrieved by using `getModuleFee`.
    /// Note: this will eventually trigger `InterchainDB.verifyEntry(entry)` function on destination chain,
    /// with no guarantee of ordering.
    /// @dev Could be only called by the Interchain DataBase contract.
    /// @param destChainId  The chain id of the destination chain
    /// @param entry        The entry to verify
    function requestVerification(uint256 destChainId, InterchainEntry memory entry) external payable;

    /// @notice Get the Module fee for verifying an entry on the specified destination chain
    /// @param destChainId  The chain id of the destination chain
    function getModuleFee(uint256 destChainId) external view returns (uint256);
}
