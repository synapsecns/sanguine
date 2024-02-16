// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IThresholdECDSAModuleEvents {
    event VerifierAdded(address verifier);
    event VerifierRemoved(address verifier);
    event ThresholdChanged(uint256 threshold);
    event GasOracleChanged(address gasOracle);
}
