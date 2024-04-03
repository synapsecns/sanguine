// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract SynapseGasOracleV1Events {
    event CalldataPriceSet(uint256 indexed chainId, uint256 calldataPrice);
    event GasPriceSet(uint256 indexed chainId, uint256 gasPrice);
    event NativePriceSet(uint256 indexed chainId, uint256 nativePrice);
}
