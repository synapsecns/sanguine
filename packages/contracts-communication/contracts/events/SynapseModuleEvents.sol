// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract SynapseModuleEvents {
    /// @notice Emitted when a verifier is added. The verifier signatures are required to verify a batch.
    /// @param verifier         The address of the verifier.
    event VerifierAdded(address verifier);

    /// Emitted when a verifier is removed.
    /// @param verifier         The address of the verifier.
    event VerifierRemoved(address verifier);

    /// @notice Emitted when a threshold is set.
    /// The threshold is the minimum number of verifiers required to verify a batch.
    /// @param threshold        The threshold value.
    event ThresholdSet(uint256 threshold);

    /// @notice Emitted when a fee collector is set. The fee collector collects fees from the module.
    /// @param feeCollector     The address of the fee collector.
    event FeeCollectorSet(address feeCollector);

    /// @notice Emitted when a gas oracle is set. The gas oracle will be used to estimate the gas cost of
    /// verifying a batch on the remote chain.
    /// @param gasOracle        The address of the gas oracle.
    event GasOracleSet(address gasOracle);

    /// @notice Emitted when the gas limit estimate is set for a chain.
    /// @param chainId          The chain ID of the chain.
    /// @param gasLimit         The gas limit estimate for verifying a batch on the chain.
    event VerifyGasLimitSet(uint64 chainId, uint256 gasLimit);

    /// @notice Emitted when the claim fee fraction is set. This fraction of the fees will be paid
    /// to the caller of the `claimFees` function.
    /// This encourages rational actors to call the function as soon as claim fee is higher than the gas cost.
    /// @param claimFeeFraction The fraction of the fees to be paid to the claimer (100% = 1e18)
    event ClaimFeeFractionSet(uint256 claimFeeFraction);

    /// @notice Emitted when fees are claimed to the fee collector address.
    /// @param feeCollector     The address of the fee collector.
    /// @param collectedFees    The amount of fees collected.
    /// @param claimer          The address of the claimer (who called `claimFees`)
    /// @param claimerFee       The amount of fees claimed by the claimer.
    event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee);

    /// @notice Emitted when the gas data from the gas oracle is sent to the remote chain.
    /// @param dstChainId       The chain ID of the destination chain.
    /// @param data             The encoded gas data.
    event GasDataSent(uint64 dstChainId, bytes data);

    /// @notice Emitted when the gas data from the remote chain is received.
    /// @param srcChainId       The chain ID of the source chain.
    /// @param data             The encoded gas data.
    event GasDataReceived(uint64 srcChainId, bytes data);
}
