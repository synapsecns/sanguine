// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract SynapseGasOracleV1Events {
    /// @notice Emitted when the calldata price is set.
    /// @param chainId       The chain ID of the chain.
    /// @param calldataPrice The price of 1 byte of calldata in the remote chain's wei.
    event CalldataPriceSet(uint64 indexed chainId, uint256 calldataPrice);

    /// @notice Emitted when the gas price is set.
    /// @param chainId   The chain ID of the chain.
    /// @param gasPrice  The gas price of the remote chain, in remote chain's wei.
    event GasPriceSet(uint64 indexed chainId, uint256 gasPrice);

    /// @notice Emitted when the native price is set.
    /// @param chainId      The chain ID of the chain.
    /// @param nativePrice  The price of the remote chain's native token in Ethereum Mainnet's wei.
    event NativePriceSet(uint64 indexed chainId, uint256 nativePrice);
}
