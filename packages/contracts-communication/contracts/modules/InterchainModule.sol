// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IInterchainDB} from "../interfaces/IInterchainDB.sol";
import {IInterchainModule} from "../interfaces/IInterchainModule.sol";

/// @notice Common logic for all Interchain Modules.
abstract contract InterchainModule is IInterchainModule {
    /// @notice The address of the Interchain DataBase contract: used for verifying the entries.
    address public immutable INTERCHAIN_DB;

    constructor(address interchainDB) {
        INTERCHAIN_DB = interchainDB;
    }

    /// @notice Request the verification of an entry from the Interchain DataBase by the module.
    /// Note: a fee is paid to the module for verification, and could be retrieved by using `getModuleFee`.
    /// Note: this will eventually trigger `InterchainDB.verifyRemoteEntry(entry)` function on destination chain,
    /// with no guarantee of ordering.
    /// @dev Could be only called by the Interchain DataBase contract.
    /// @param dstChainId       The chain id of the destination chain
    /// @param versionedEntry   The versioned entry to verify
    function requestEntryVerification(uint64 dstChainId, bytes memory versionedEntry) external payable {
        if (msg.sender != INTERCHAIN_DB) {
            revert InterchainModule__CallerNotInterchainDB(msg.sender);
        }
        if (dstChainId == block.chainid) {
            revert InterchainModule__ChainIdNotRemote(dstChainId);
        }
        uint256 requiredFee = _getModuleFee(dstChainId);
        if (msg.value < requiredFee) {
            revert InterchainModule__FeeAmountBelowMin({feeAmount: msg.value, minRequired: requiredFee});
        }
        // Note: we don't emit an event here, the derived contract could emit an event if needed.
        _relayDBEntry(dstChainId, versionedEntry);
    }

    /// @notice Get the Module fee for verifying an entry on the specified destination chain.
    /// @param dstChainId   The chain id of the destination chain
    function getModuleFee(uint64 dstChainId) external view returns (uint256) {
        return _getModuleFee(dstChainId);
    }

    /// @dev Should be called once the Module has verified the entry and needs to signal this
    /// to the InterchainDB.
    function _verifyRemoteEntry(bytes memory versionedEntry) internal {
        IInterchainDB(INTERCHAIN_DB).verifyRemoteEntry(versionedEntry);
    }

    // solhint-disable no-empty-blocks
    /// @dev Internal logic to relay a DB entry to the destination chain.
    /// Following checks have been done at this point:
    /// - Entry is a valid versioned entry coming from the Interchain DataBase.
    /// - Enough fees have been paid for the verification.
    ///
    /// Derived contracts should implement the logic so that eventually the destination counterpart
    /// of this module calls `_verifyRemoteEntry(versionedEntry)`.
    function _relayDBEntry(uint64 dstChainId, bytes memory versionedEntry) internal virtual;

    /// @dev Internal logic to get the module fee for verifying an entry on the specified destination chain.
    function _getModuleFee(uint64 dstChainId) internal view virtual returns (uint256);
}
