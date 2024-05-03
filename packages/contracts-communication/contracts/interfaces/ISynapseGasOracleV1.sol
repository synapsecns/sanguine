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

    error SynapseGasOracleV1__ChainIdNotRemote(uint64 chainId);
    error SynapseGasOracleV1__NativePriceNotSet(uint64 chainId);
    error SynapseGasOracleV1__NativePriceZero();

    function setLocalNativePrice(uint256 nativePrice) external;
    function setRemoteGasData(uint64 chainId, RemoteGasData memory data) external;
    function setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) external;
    function setRemoteGasPrice(uint64 chainId, uint256 gasPrice) external;
    function setRemoteNativePrice(uint64 chainId, uint256 nativePrice) external;

    function getLocalNativePrice() external view returns (uint256);
    function getRemoteGasData(uint64 chainId) external view returns (RemoteGasData memory);
}
