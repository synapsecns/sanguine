// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ISynapseGasOracle} from "./ISynapseGasOracle.sol";

interface ISynapseGasOracleV1 is ISynapseGasOracle {
    /// @notice Struct defining the gas data for the remote chain.
    /// @param calldataPrice    The price of 1 byte of calldata in the remote chain's wei.
    /// @param gasPrice         The gas price of the remote chain, in remote chain's wei.
    /// @param nativePrice      The price of the remote chain's native token in Ethereum Mainnet's wei.
    struct RemoteGasData {
        uint256 calldataPrice;
        uint256 gasPrice;
        uint256 nativePrice;
    }

    error SynapseGasOracleV1__NotRemoteChainId(uint256 chainId);
    error SynapseGasOracleV1__NativePriceNotSet(uint256 chainId);
    error SynapseGasOracleV1__NativePriceZero();

    /// @notice Allows the contract owner to set the native token price of the local chain.
    /// @dev Could only be called by the contract owner. Will revert if the native token price is 0.
    /// @param nativePrice      The price of the local chain's native token in Ethereum Mainnet's wei.
    function setLocalNativePrice(uint256 nativePrice) external;

    /// @notice Allows the contract owner to set the gas data for a remote chain.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the native token price is 0, or if the chain id is not a remote chain id.
    /// @param chainId          The chain id of the remote chain.
    /// @param data             The gas data for the remote chain.
    function setRemoteGasData(uint256 chainId, RemoteGasData memory data) external;

    /// @notice Allows the contract owner to set the price of remote chain's calldata.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the chain id is not a remote chain id, or if native token price for the chain is 0.
    /// @param chainId          The chain id of the remote chain.
    /// @param calldataPrice    The price of 1 byte of calldata in the remote chain's wei.
    function setRemoteCallDataPrice(uint256 chainId, uint256 calldataPrice) external;

    /// @notice Allows the contract owner to set the gas price of the remote chain.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the chain id is not a remote chain id, or if native token price for the chain is 0.
    /// @param chainId          The chain id of the remote chain.
    /// @param gasPrice         The gas price of the remote chain, in remote chain's wei.
    function setRemoteGasPrice(uint256 chainId, uint256 gasPrice) external;

    /// @notice Allows the contract owner to set the price of the remote chain's native token.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the chain id is not a remote chain id, or if the price is 0.
    /// @param chainId          The chain id of the remote chain.
    /// @param nativePrice      The price of the remote chain's native token in Ethereum Mainnet's wei.
    function setRemoteNativePrice(uint256 chainId, uint256 nativePrice) external;

    /// @notice Gets the price of the local chain's native token in Ethereum Mainnet's wei.
    function getLocalNativePrice() external view returns (uint256);

    /// @notice Gets the gas data for a remote chain.
    /// @dev Will revert if the chain id is not a remote chain id.
    /// @param chainId          The chain id of the remote chain.
    function getRemoteGasData(uint256 chainId) external view returns (RemoteGasData memory);
}
