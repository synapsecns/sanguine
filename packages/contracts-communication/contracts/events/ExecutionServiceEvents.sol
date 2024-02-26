// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ExecutionServiceEvents {
    event ExecutionRequested(
        bytes32 indexed transactionId
    );
    event ExecutorEOAUpdated(address indexed executorEOA);
    event GasOracleUpdated(address indexed gasOracle);
    event InterchainClientUpdated(address indexed interchainClient);
}
