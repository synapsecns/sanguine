// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// TODO: remove this
abstract contract ExecutionServiceEvents {
    event ExecutionRequested(bytes32 indexed transactionId, address client);
    event ExecutorEOAUpdated(address indexed executorEOA);
    event GasOracleUpdated(address indexed gasOracle);
    event InterchainClientUpdated(address indexed interchainClient);
}
