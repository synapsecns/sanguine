// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the GasOracle contract
abstract contract GasOracleEvents {
    /**
     * @notice Emitted when gas data is updated for the domain
     * @param domain        Domain of chain the gas data is for
     * @param paddedGasData Padded encoded gas data
     */
    event GasDataUpdated(uint32 domain, uint256 paddedGasData);
}
