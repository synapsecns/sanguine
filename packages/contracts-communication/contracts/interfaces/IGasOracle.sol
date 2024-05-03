// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IGasOracle {
    function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) external view returns (uint256);

    function estimateTxCostInLocalUnits(
        uint64 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        returns (uint256);
    function estimateTxCostInRemoteUnits(
        uint64 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        returns (uint256);
}
