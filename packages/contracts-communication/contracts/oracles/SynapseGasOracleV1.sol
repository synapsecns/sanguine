// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseGasOracleV1Events} from "../events/SynapseGasOracleV1Events.sol";
import {ISynapseGasOracle, IGasOracle} from "../interfaces/ISynapseGasOracle.sol";
import {ISynapseGasOracleV1} from "../interfaces/ISynapseGasOracleV1.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract SynapseGasOracleV1 is Ownable, SynapseGasOracleV1Events, ISynapseGasOracleV1 {
    uint256 internal _localNativePrice;

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
    function setRemoteGasData(uint256 chainId, RemoteGasData memory data) external {}

    /// @inheritdoc ISynapseGasOracleV1
    function setRemoteCallDataPrice(uint256 chainId, uint256 calldataPrice) external {}

    /// @inheritdoc ISynapseGasOracleV1
    function setRemoteGasPrice(uint256 chainId, uint256 gasPrice) external {}

    /// @inheritdoc ISynapseGasOracleV1
    function setRemoteNativePrice(uint256 chainId, uint256 nativePrice) external {}

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
    function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) external view returns (uint256) {}

    /// @inheritdoc IGasOracle
    function estimateTxCostInLocalUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        returns (uint256)
    {}

    /// @inheritdoc IGasOracle
    function estimateTxCostInRemoteUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        returns (uint256)
    {}

    /// @inheritdoc ISynapseGasOracleV1
    function getLocalNativePrice() external view returns (uint256) {
        return _localNativePrice;
    }

    /// @inheritdoc ISynapseGasOracleV1
    function getRemoteGasData(uint256 chainId) external view returns (RemoteGasData memory) {}
}
