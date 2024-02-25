// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ExecutionServiceEvents {
    event ExecutionRequested(
        uint256 indexed dstChainId,
        uint256 indexed txPayloadSize,
        bytes32 indexed transactionId,
        uint256 executionFee,
        bytes options
    );
    event ExecutorEOAUpdated(address indexed executorEOA);
    event GasOracleUpdated(address indexed gasOracle);
    event InterchainClientUpdated(address indexed interchainClient);
}
