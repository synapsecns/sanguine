// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IGasOracle {
    /// @notice Convert a value from the native token of a remote chain to the local native token.
    /// @dev Will revert if no price is available for the remote chain.
    /// @param remoteChainId        The chain id of the remote chain.
    /// @param value                The value to convert.
    function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) external view returns (uint256);

    /// @notice Estimate the cost of execution a transaction on a remote chain,
    /// and convert it to the local native token.
    /// @dev Will revert if no price is available for the remote chain.
    /// @param remoteChainId        The chain id of the remote chain.
    /// @param gasLimit             The gas limit of the transaction.
    /// @param calldataSize         The size of the transaction calldata.
    function estimateTxCostInLocalUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        returns (uint256);

    /// @notice Estimate the cost of execution a transaction on a remote chain,
    /// and return it as is in the remote chain's native token.
    /// @dev Will revert if no price is available for the remote chain.
    /// @param remoteChainId        The chain id of the remote chain.
    /// @param gasLimit             The gas limit of the transaction.
    /// @param calldataSize         The size of the transaction calldata.
    function estimateTxCostInRemoteUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        returns (uint256);
}
