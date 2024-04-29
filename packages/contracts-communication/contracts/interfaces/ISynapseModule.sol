// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IInterchainModule} from "./IInterchainModule.sol";

interface ISynapseModule is IInterchainModule {
    error SynapseModule__GasOracleNotContract(address gasOracle);
    error SynapseModule__GasOracleNotSet();
    error SynapseModule__ZeroAddress();

    // ═══════════════════════════════════════════════ PERMISSIONED ════════════════════════════════════════════════════

    /// @notice Adds a new verifier to the module.
    /// @dev Could be only called by the owner. Will revert if the verifier is already added.
    /// @param verifier     The address of the verifier to add
    function addVerifier(address verifier) external;

    /// @notice Adds a list of new verifiers to the module.
    /// @dev Could be only called by the owner. Will revert if any of the verifiers is already added.
    /// @param verifiers    The list of addresses of the verifiers to add
    function addVerifiers(address[] calldata verifiers) external;

    /// @notice Removes a verifier from the module.
    /// @dev Could be only called by the owner. Will revert if the verifier is not added.
    /// @param verifier     The address of the verifier to remove
    function removeVerifier(address verifier) external;

    /// @notice Removes a list of verifiers from the module.
    /// @dev Could be only called by the owner. Will revert if any of the verifiers is not added.
    /// @param verifiers    The list of addresses of the verifiers to remove
    function removeVerifiers(address[] calldata verifiers) external;

    /// @notice Sets the threshold of the module.
    /// @dev Could be only called by the owner. Will revert if the threshold is zero.
    /// @param threshold    The new threshold value
    function setThreshold(uint256 threshold) external;

    /// @notice Sets the address of the fee collector, which will have the verification fees forwarded to it.
    /// @dev Could be only called by the owner.
    /// @param feeRecipient_   The address of the fee collector
    function setFeeRecipient(address feeRecipient_) external;

    /// @notice Sets the fraction of the accumulated fees to be paid to caller of `claimFees`.
    /// This encourages rational actors to call the function as soon as claim fee is higher than the gas cost.
    /// @dev Could be only called by the owner. Could not exceed 1%.
    /// @param claimerFraction  The fraction of the fees to be paid to the claimer (100% = 1e18)
    function setClaimerFraction(uint256 claimerFraction) external;

    /// @notice Sets the address of the gas oracle to be used for estimating the verification fees.
    /// @dev Could be only called by the owner. Will revert if the gas oracle is not a contract.
    /// @param gasOracle_   The address of the gas oracle contract
    function setGasOracle(address gasOracle_) external;

    /// @notice Sets the estimated gas limit for verifying a batch on the given chain.
    /// @dev Could be only called by the owner.
    /// @param chainId      The chain ID for which to set the gas limit
    /// @param gasLimit     The new gas limit
    function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) external;

    // ══════════════════════════════════════════════ PERMISSIONLESS ═══════════════════════════════════════════════════

    /// @notice Verifies a batch from the remote chain using a set of verifier signatures.
    /// If the threshold is met, the batch will be marked as verified in the Interchain DataBase.
    /// @dev List of recovered signers from the signatures must be sorted in the ascending order.
    /// @param encodedBatch The encoded batch to verify
    /// @param signatures   Signatures used to verify the batch, concatenated
    function verifyRemoteBatch(bytes calldata encodedBatch, bytes calldata signatures) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Returns the address of the gas oracle used for estimating the verification fees.
    function gasOracle() external view returns (address);

    /// @notice Returns the list of verifiers for the module.
    function getVerifiers() external view returns (address[] memory);

    /// @notice Gets the threshold of the module.
    /// This is the minimum number of signatures required for verification.
    function getThreshold() external view returns (uint256);

    /// @notice Checks if the given account is a verifier for the module.
    function isVerifier(address account) external view returns (bool);

    /// @notice Returns the estimated gas limit for verifying a batch on the given chain.
    /// Note: this defaults to DEFAULT_VERIFY_GAS_LIMIT if not set.
    function getVerifyGasLimit(uint64 chainId) external view returns (uint256);
}
