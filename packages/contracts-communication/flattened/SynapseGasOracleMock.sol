// SPDX-License-Identifier: MIT
pragma solidity =0.8.20 ^0.8.0;

// contracts/interfaces/IGasOracle.sol

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

// contracts/interfaces/ISynapseGasOracle.sol

interface ISynapseGasOracle is IGasOracle {
    /// @notice Allows Synapse Module to pass the gas data from a remote chain to the Gas Oracle.
    /// @dev Could only be called by Synapse Module.
    /// @param srcChainId        The chain id of the remote chain.
    /// @param data              The gas data from the remote chain.
    function receiveRemoteGasData(uint256 srcChainId, bytes calldata data) external;

    /// @notice Gets the gas data for the local chain.
    function getLocalGasData() external view returns (bytes memory);
}

// test/mocks/SynapseGasOracleMock.sol

contract SynapseGasOracleMock is ISynapseGasOracle {
    function receiveRemoteGasData(uint256 srcChainId, bytes calldata data) external {}

    function getLocalGasData() external view returns (bytes memory) {}

    function convertRemoteValueToLocalUnits(
        uint256 remoteChainId,
        uint256 value
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInLocalUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInRemoteUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}
}
