// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ThresholdECDSAModuleEvents {
    event VerifierAdded(address verifier);
    event VerifierRemoved(address verifier);
    event ThresholdChanged(uint256 threshold);

    event FeeCollectorChanged(address feeCollector);
    event GasOracleChanged(address gasOracle);
}
