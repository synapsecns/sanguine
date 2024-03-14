// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseGasOracleV1Events} from "../events/SynapseGasOracleV1Events.sol";
import {ISynapseGasOracle, IGasOracle} from "../interfaces/ISynapseGasOracle.sol";
import {ISynapseGasOracleV1} from "../interfaces/ISynapseGasOracleV1.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract SynapseGasOracleV1 is Ownable, SynapseGasOracleV1Events, ISynapseGasOracleV1 {
    uint256 internal _localNativePrice;
    mapping(uint256 chainId => RemoteGasData data) internal _remoteGasData;

    /// @dev Checks that the chain ID is not the local chain ID.
    modifier onlyRemoteChainId(uint256 chainId) {
        if (block.chainid == chainId) {
            revert SynapseGasOracleV1__NotRemoteChainId(chainId);
        }
        _;
    }

    /// @dev Checks that the native token price is set for a remote chain ID.
    modifier onlyNativePriceSet(uint256 chainId) {
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

    /// @inheritdoc ISynapseGasOracleV1
    function setLocalNativePrice(uint256 nativePrice) external onlyOwner onlyNonZeroNativePrice(nativePrice) {
        if (_localNativePrice != nativePrice) {
            _localNativePrice = nativePrice;
            emit NativePriceSet(block.chainid, nativePrice);
        }
    }

    /// @inheritdoc ISynapseGasOracleV1
    function setRemoteGasData(
        uint256 chainId,
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

    /// @inheritdoc ISynapseGasOracleV1
    function setRemoteCallDataPrice(
        uint256 chainId,
        uint256 calldataPrice
    )
        external
        onlyOwner
        onlyRemoteChainId(chainId)
        onlyNativePriceSet(chainId)
    {
        _setRemoteCallDataPrice(chainId, calldataPrice);
    }

    /// @inheritdoc ISynapseGasOracleV1
    function setRemoteGasPrice(
        uint256 chainId,
        uint256 gasPrice
    )
        external
        onlyOwner
        onlyRemoteChainId(chainId)
        onlyNativePriceSet(chainId)
    {
        _setRemoteGasPrice(chainId, gasPrice);
    }

    /// @inheritdoc ISynapseGasOracleV1
    function setRemoteNativePrice(
        uint256 chainId,
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
    /// @inheritdoc ISynapseGasOracle
    function receiveRemoteGasData(uint256 srcChainId, bytes calldata data) external {
        // The V1 version has this function as a no-op, hence we skip the permission check.
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ISynapseGasOracle
    function getLocalGasData() external view returns (bytes memory) {
        // The V1 version has this function as a no-op.
    }
    // solhint-enable no-empty-blocks

    /// @inheritdoc IGasOracle
    function convertRemoteValueToLocalUnits(
        uint256 remoteChainId,
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

    /// @inheritdoc IGasOracle
    function estimateTxCostInLocalUnits(
        uint256 remoteChainId,
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

    /// @inheritdoc IGasOracle
    function estimateTxCostInRemoteUnits(
        uint256 remoteChainId,
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

    /// @inheritdoc ISynapseGasOracleV1
    function getLocalNativePrice() external view returns (uint256) {
        return _localNativePrice;
    }

    /// @inheritdoc ISynapseGasOracleV1
    function getRemoteGasData(uint256 chainId) external view returns (RemoteGasData memory) {
        return _remoteGasData[chainId];
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Updates the calldata price for the given remote chain, no-op if the price is already set.
    function _setRemoteCallDataPrice(uint256 chainId, uint256 calldataPrice) internal {
        if (_remoteGasData[chainId].calldataPrice != calldataPrice) {
            _remoteGasData[chainId].calldataPrice = calldataPrice;
            emit CalldataPriceSet(chainId, calldataPrice);
        }
    }

    /// @dev Updates the gas price for the given remote chain, no-op if the price is already set.
    function _setRemoteGasPrice(uint256 chainId, uint256 gasPrice) internal {
        if (_remoteGasData[chainId].gasPrice != gasPrice) {
            _remoteGasData[chainId].gasPrice = gasPrice;
            emit GasPriceSet(chainId, gasPrice);
        }
    }

    /// @dev Updates the native token price for the given remote chain, no-op if the price is already set.
    function _setRemoteNativePrice(uint256 chainId, uint256 nativePrice) internal {
        if (_remoteGasData[chainId].nativePrice != nativePrice) {
            _remoteGasData[chainId].nativePrice = nativePrice;
            emit NativePriceSet(chainId, nativePrice);
        }
    }

    /// @dev Converts value denominated in remote chain's units to local chain's units.
    /// Note: the check for non-zero remote native token price is done outside this function.
    function _convertRemoteValueToLocalUnits(
        uint256 remoteChainId,
        uint256 remoteValue
    )
        internal
        view
        returns (uint256)
    {
        if (_localNativePrice == 0) {
            revert SynapseGasOracleV1__NativePriceNotSet(block.chainid);
        }
        return (remoteValue * _remoteGasData[remoteChainId].nativePrice) / _localNativePrice;
    }

    /// @dev Estimates the transaction cost in remote chain's units.
    /// Note: the check for non-zero remote native token price is done outside this function.
    function _estimateTxCostInRemoteUnits(
        uint256 remoteChainId,
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
