// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IInterchainModule} from "./IInterchainModule.sol";

interface IThresholdECDSAModule is IInterchainModule {
    // ═══════════════════════════════════════════════ PERMISSIONED ════════════════════════════════════════════════════

    /// @notice Adds a new verifier to the module.
    /// @dev Could be only called by the owner. Will revert if the verifier is already added.
    /// @param verifier     The address of the verifier to add
    function addVerifier(address verifier) external;

    /// @notice Removes a verifier from the module.
    /// @dev Could be only called by the owner. Will revert if the verifier is not added.
    /// @param verifier     The address of the verifier to remove
    function removeVerifier(address verifier) external;

    /// @notice Sets the threshold of the module.
    /// @dev Could be only called by the owner. Will revert if the threshold is zero.
    /// @param threshold    The new threshold value
    function setThreshold(uint256 threshold) external;

    /// @notice Sets the address of the gas oracle to be used for estimating the verification fees.
    /// @dev Could be only called by the owner. Will revert if the gas oracle is not a contract.
    /// @param gasOracle    The address of the gas oracle contract
    function setGasOracle(address gasOracle) external;

    // ══════════════════════════════════════════════ PERMISSIONLESS ═══════════════════════════════════════════════════

    /// @notice Verifies an entry using a set of verifier signatures.
    /// If the threshold is met, the entry will be marked as verified in the Interchain DataBase.
    /// @param encodedEntry The encoded entry to verify
    /// @param signatures   An array of signatures used to verify the entry
    function verifyEntry(bytes calldata encodedEntry, bytes[] calldata signatures) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Gets the threshold of the module.
    /// This is the minimum number of signatures required for verification.
    function getThreshold() external view returns (uint256);
}
