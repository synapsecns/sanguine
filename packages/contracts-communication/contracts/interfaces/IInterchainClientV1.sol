// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainClientV1 {
    error InterchainClientV1__AlreadyExecuted(bytes32 transactionId);
    error InterchainClientV1__IncorrectMsgValue(uint256 actual, uint256 required);
    error InterchainClientV1__NotEnoughResponses(uint256 actual, uint256 required);

    /**
     * @notice Sets the address of the ExecutionFees contract.
     * @dev Only callable by the contract owner or an authorized account.
     * @param executionFees_ The address of the ExecutionFees contract.
     */
    function setExecutionFees(address executionFees_) external;

    /**
     * @notice Sets the linked client for a specific chain ID.
     * @dev Stores the address of the linked client in a mapping with the chain ID as the key.
     * @param chainId The chain ID for which the client is being set.
     * @param client The address of the client being linked.
     */
    function setLinkedClient(uint256 chainId, bytes32 client) external;

    /**
     * @notice Sets the address of the InterchainDB contract.
     * @dev Only callable by the contract owner or an authorized account.
     * @param _interchainDB The address of the InterchainDB contract.
     */
    function setInterchainDB(address _interchainDB) external;

    /**
     * @notice Sends a message to another chain via the Interchain Communication Protocol.
     * @dev Charges a fee for the message, which is payable upon calling this function:
     * - Verification fees: paid to every module that verifies the message.
     * - Execution fee: paid to the executor that executes the message.
     * Note: while a specific execution service is specified to request the execution of the message,
     * any executor is able to execute the message on destination chain, earning the execution fee.
     * @param dstChainId The chain ID of the destination chain.
     * @param receiver The address of the receiver on the destination chain.
     * @param srcExecutionService The address of the execution service to use for the message.
     * @param srcModules The source modules involved in the message sending.
     * @param options Execution options for the message sent, encoded as bytes, currently primarily gas limit + native gas drop.
     * @param message The message being sent.
     */
    function interchainSend(
        uint256 dstChainId,
        bytes32 receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        payable;

    /**
     * @notice Executes a transaction that has been sent via the Interchain.
     * @dev The transaction must have been previously sent and recorded.
     * Transaction data includes the requested gas limit, but the executors could specify a different gas limit.
     * If the specified gas limit is lower than requested, the requested gas limit will be used.
     * Otherwise, the specified gas limit will be used.
     * This allows to execute the transactions with requested gas limit set too low.
     * @param gasLimit          The gas limit to use for the execution.
     * @param transaction       The transaction data.
     */
    function interchainExecute(uint256 gasLimit, bytes calldata transaction) external payable;

    /**
     * @notice Checks if a transaction is executable.
     * @dev Determines if a transaction meets the criteria to be executed based on:
     * - If approved modules have written to the InterchainDB
     * - If the threshold of approved modules have been met
     * - If the optimistic window has passed for all modules
     * @param transaction The InterchainTransaction struct to be checked.
     * @return bool Returns true if the transaction is executable, false otherwise.
     */
    function isExecutable(bytes calldata transaction) external view returns (bool);
}
