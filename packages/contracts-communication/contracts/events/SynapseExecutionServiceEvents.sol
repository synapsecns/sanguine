// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract SynapseExecutionServiceEvents {
    event ExecutorEOASet(address executorEOA);
    event GasOracleSet(address gasOracle);
    event GlobalMarkupSet(uint256 globalMarkup);

    event ExecutionRequested(bytes32 indexed transactionId, address client);
}
