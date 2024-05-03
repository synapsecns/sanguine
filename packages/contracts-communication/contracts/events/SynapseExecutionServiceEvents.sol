// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// TODO: regenerate `sin-executor` before redeployment
abstract contract SynapseExecutionServiceEvents {
    /// @notice Emitted when the executor EOA is set.
    /// This address will be getting the execution fees that the service earns.
    /// @param executorEOA   The address of the executor EOA.
    event ExecutorEOASet(address executorEOA);

    /// @notice Emitted when the gas oracle is set.
    /// This gas oracle will be used to estimate the gas cost of the transactions.
    /// @param gasOracle     The address of the gas oracle.
    event GasOracleSet(address gasOracle);

    /// @notice Emitted when the global markup is set. This markup will be added to the gas cost of the transactions.
    /// Zero markup means that the Execution Service charges the exact gas cost estimated by the GasOracle.
    /// The markup is denominated in Wei, 1e18 being 100%.
    /// @param globalMarkup  The global markup value.
    event GlobalMarkupSet(uint256 globalMarkup);

    /// @notice Emitted when the execution of a transaction is requested.
    /// @param transactionId The unique identifier of the transaction.
    /// @param client        The address of the Interchain Client that requested the execution.
    /// @param executionFee  The fee paid for the execution.
    event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee);
}
