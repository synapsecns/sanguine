// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract SynapseModuleEvents {
    event VerifierAdded(address verifier);
    event VerifierRemoved(address verifier);
    event ThresholdChanged(uint256 threshold);

    event FeeCollectorChanged(address feeCollector);
    event GasOracleChanged(address gasOracle);
    event VerifyGasLimitChanged(uint256 chainId, uint256 gasLimit);

    event ClaimFeeFractionChanged(uint256 claimFeeFraction);
    event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee);
}
