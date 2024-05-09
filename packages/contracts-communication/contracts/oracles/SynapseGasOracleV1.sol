// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseGasOracleV1Events} from "../events/SynapseGasOracleV1Events.sol";
import {ISynapseGasOracleV1} from "../interfaces/ISynapseGasOracleV1.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract SynapseGasOracleV1 is Ownable, SynapseGasOracleV1Events, ISynapseGasOracleV1 {
    /// @dev Price of the local chain's native token, expressed in Ethereum Mainnet's wei.
    uint256 internal _localNativePrice;
    /// @dev Gas data for tracked remote chains:
    /// - calldataPrice     The price of 1 byte of calldata in the remote chain's wei.
    /// - gasPrice          The gas price of the remote chain, in remote chain's wei.
    /// - nativePrice       The price of the remote chain's native token in Ethereum Mainnet's wei.
    mapping(uint64 chainId => RemoteGasData data) internal _remoteGasData;

    /// @dev Checks that the chain ID is not the local chain ID.
    modifier onlyRemoteChainId(uint64 chainId) {
        if (block.chainid == chainId) {
            revert SynapseGasOracleV1__ChainIdNotRemote(chainId);
        }
        _;
    }

    /// @dev Checks that the native token price is set for a remote chain ID.
    modifier onlyNativePriceSet(uint64 chainId) {
        if (_remoteGasData[chainId].nativePrice == 0) {
            revert SynapseGasOracleV1__NativePriceNotSet(chainId);
        }
        _;
    }

    /// @dev Checks that the native token price is non-zero.
    modifier onlyNonZeroNativePrice(uint256 nativePrice) {
        if (nativePrice == 0) {
            revert SynapseGasOracleV1__NativePriceZero();
        }
        _;
    }

    constructor(address owner_) Ownable(owner_) {}

    // ════════════════════════════════════════════════ ONLY OWNER ═════════════════════════════════════════════════════

    /// @notice Allows the contract owner to set the native token price of the local chain.
    /// @dev Could only be called by the contract owner. Will revert if the native token price is 0.
    /// @param nativePrice      The price of the local chain's native token in Ethereum Mainnet's wei.
    function setLocalNativePrice(uint256 nativePrice) external onlyOwner onlyNonZeroNativePrice(nativePrice) {
        if (_localNativePrice != nativePrice) {
            _localNativePrice = nativePrice;
            emit NativePriceSet(SafeCast.toUint64(block.chainid), nativePrice);
        }
    }

    /// @notice Allows the contract owner to set the gas data for a remote chain.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the native token price is 0, or if the chain id is not a remote chain id.
    /// @param chainId          The chain id of the remote chain.
    /// @param data             The gas data for the remote chain.
    function setRemoteGasData(
        uint64 chainId,
        RemoteGasData memory data
    )
        external
        onlyOwner
        onlyRemoteChainId(chainId)
        onlyNonZeroNativePrice(data.nativePrice)
    {
        _setRemoteCallDataPrice(chainId, data.calldataPrice);
        _setRemoteGasPrice(chainId, data.gasPrice);
        _setRemoteNativePrice(chainId, data.nativePrice);
    }

    /// @notice Allows the contract owner to set the price of remote chain's calldata.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the chain id is not a remote chain id, or if native token price for the chain is 0.
    /// @param chainId          The chain id of the remote chain.
    /// @param calldataPrice    The price of 1 byte of calldata in the remote chain's wei.
    function setRemoteCallDataPrice(
        uint64 chainId,
        uint256 calldataPrice
    )
        external
        onlyOwner
        onlyRemoteChainId(chainId)
        onlyNativePriceSet(chainId)
    {
        _setRemoteCallDataPrice(chainId, calldataPrice);
    }

    /// @notice Allows the contract owner to set the gas price of the remote chain.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the chain id is not a remote chain id, or if native token price for the chain is 0.
    /// @param chainId          The chain id of the remote chain.
    /// @param gasPrice         The gas price of the remote chain, in remote chain's wei.
    function setRemoteGasPrice(
        uint64 chainId,
        uint256 gasPrice
    )
        external
        onlyOwner
        onlyRemoteChainId(chainId)
        onlyNativePriceSet(chainId)
    {
        _setRemoteGasPrice(chainId, gasPrice);
    }

    /// @notice Allows the contract owner to set the price of the remote chain's native token.
    /// @dev Could only be called by the contract owner.
    /// Will revert if the chain id is not a remote chain id, or if the price is 0.
    /// @param chainId          The chain id of the remote chain.
    /// @param nativePrice      The price of the remote chain's native token in Ethereum Mainnet's wei.
    function setRemoteNativePrice(
        uint64 chainId,
        uint256 nativePrice
    )
        external
        onlyOwner
        onlyRemoteChainId(chainId)
        onlyNonZeroNativePrice(nativePrice)
    {
        _setRemoteNativePrice(chainId, nativePrice);
    }

    // ════════════════════════════════════════════════ ONLY MODULE ════════════════════════════════════════════════════

    // solhint-disable no-empty-blocks
    /// @notice Allows Synapse Module to pass the gas data from a remote chain to the Gas Oracle.
    /// @dev Could only be called by Synapse Module.
    /// @param srcChainId        The chain id of the remote chain.
    /// @param data              The gas data from the remote chain.
    function receiveRemoteGasData(uint64 srcChainId, bytes calldata data) external {
        // The V1 version has this function as a no-op, hence we skip the permission check.
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Gets the gas data for the local chain.
    function getLocalGasData() external view returns (bytes memory) {
        // The V1 version has this function as a no-op.
    }
    // solhint-enable no-empty-blocks

    /// @notice Convert a value from the native token of a remote chain to the local native token.
    /// @dev Will revert if no price is available for the remote chain.
    /// @param remoteChainId        The chain id of the remote chain.
    /// @param value                The value to convert.
    function convertRemoteValueToLocalUnits(
        uint64 remoteChainId,
        uint256 value
    )
        external
        view
        onlyRemoteChainId(remoteChainId)
        onlyNativePriceSet(remoteChainId)
        returns (uint256)
    {
        // This will revert if the local native price is not set.
        return _convertRemoteValueToLocalUnits(remoteChainId, value);
    }

    /// @notice Estimate the cost of execution a transaction on a remote chain,
    /// and convert it to the local native token.
    /// @dev Will revert if no price is available for the remote chain.
    /// @param remoteChainId        The chain id of the remote chain.
    /// @param gasLimit             The gas limit of the transaction.
    /// @param calldataSize         The size of the transaction calldata.
    function estimateTxCostInLocalUnits(
        uint64 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        onlyRemoteChainId(remoteChainId)
        onlyNativePriceSet(remoteChainId)
        returns (uint256)
    {
        uint256 remoteTxCost = _estimateTxCostInRemoteUnits(remoteChainId, gasLimit, calldataSize);
        // This will revert if the local native price is not set.
        return _convertRemoteValueToLocalUnits(remoteChainId, remoteTxCost);
    }

    /// @notice Estimate the cost of execution a transaction on a remote chain,
    /// and return it as is in the remote chain's native token.
    /// @dev Will revert if no price is available for the remote chain.
    /// @param remoteChainId        The chain id of the remote chain.
    /// @param gasLimit             The gas limit of the transaction.
    /// @param calldataSize         The size of the transaction calldata.
    function estimateTxCostInRemoteUnits(
        uint64 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        onlyRemoteChainId(remoteChainId)
        onlyNativePriceSet(remoteChainId)
        returns (uint256)
    {
        // This will NOT revert if the local native price is not set, and we are fine with that.
        return _estimateTxCostInRemoteUnits(remoteChainId, gasLimit, calldataSize);
    }

    /// @notice Gets the price of the local chain's native token in Ethereum Mainnet's wei.
    function getLocalNativePrice() external view returns (uint256) {
        return _localNativePrice;
    }

    /// @notice Gets the gas data for a remote chain.
    /// @dev Will revert if the chain id is not a remote chain id.
    /// @param chainId          The chain id of the remote chain.
    function getRemoteGasData(uint64 chainId) external view onlyRemoteChainId(chainId) returns (RemoteGasData memory) {
        return _remoteGasData[chainId];
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Updates the calldata price for the given remote chain, no-op if the price is already set.
    function _setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) internal {
        if (_remoteGasData[chainId].calldataPrice != calldataPrice) {
            _remoteGasData[chainId].calldataPrice = calldataPrice;
            emit CalldataPriceSet(chainId, calldataPrice);
        }
    }

    /// @dev Updates the gas price for the given remote chain, no-op if the price is already set.
    function _setRemoteGasPrice(uint64 chainId, uint256 gasPrice) internal {
        if (_remoteGasData[chainId].gasPrice != gasPrice) {
            _remoteGasData[chainId].gasPrice = gasPrice;
            emit GasPriceSet(chainId, gasPrice);
        }
    }

    /// @dev Updates the native token price for the given remote chain, no-op if the price is already set.
    function _setRemoteNativePrice(uint64 chainId, uint256 nativePrice) internal {
        if (_remoteGasData[chainId].nativePrice != nativePrice) {
            _remoteGasData[chainId].nativePrice = nativePrice;
            emit NativePriceSet(chainId, nativePrice);
        }
    }

    /// @dev Converts value denominated in remote chain's units to local chain's units.
    /// Note: the check for non-zero remote native token price is done outside this function.
    function _convertRemoteValueToLocalUnits(
        uint64 remoteChainId,
        uint256 remoteValue
    )
        internal
        view
        returns (uint256)
    {
        if (_localNativePrice == 0) {
            revert SynapseGasOracleV1__NativePriceNotSet(SafeCast.toUint64(block.chainid));
        }
        return (remoteValue * _remoteGasData[remoteChainId].nativePrice) / _localNativePrice;
    }

    /// @dev Estimates the transaction cost in remote chain's units.
    /// Note: the check for non-zero remote native token price is done outside this function.
    function _estimateTxCostInRemoteUnits(
        uint64 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        internal
        view
        returns (uint256)
    {
        return gasLimit * _remoteGasData[remoteChainId].gasPrice
            + calldataSize * _remoteGasData[remoteChainId].calldataPrice;
    }
}
