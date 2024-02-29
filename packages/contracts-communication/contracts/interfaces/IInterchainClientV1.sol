// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainClientV1 {
    error InterchainClientV1__FeeAmountTooLow(uint256 actual, uint256 required);
    error InterchainClientV1__IncorrectDstChainId(uint256 chainId);
    error InterchainClientV1__IncorrectMsgValue(uint256 actual, uint256 required);
    error InterchainClientV1__NoLinkedClient(uint256 chainId);
    error InterchainClientV1__NotEnoughResponses(uint256 actual, uint256 required);
    error InterchainClientV1__NotEVMClient(bytes32 client);
    error InterchainClientV1__NotRemoteChainId(uint256 chainId);
    error InterchainClientV1__TxAlreadyExecuted(bytes32 transactionId);
    error InterchainClientV1__TxNotExecuted(bytes32 transactionId);

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
     * @param options Execution options for the message sent, encoded as bytes, currently gas limit + native gas drop.
     * @param message The message being sent.
     * @return transactionId The ID of the transaction that was sent.
     * @return dbNonce The database nonce of the written entry for the transaction.
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
        payable
        returns (bytes32 transactionId, uint256 dbNonce);

    function interchainSendEVM(
        uint256 dstChainId,
        address receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        payable
        returns (bytes32 transactionId, uint256 dbNonce);

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

    /// @notice Writes the proof of execution for a transaction into the InterchainDB.
    /// @dev Will revert if the transaction has not been executed.
    /// @param transactionId    The ID of the transaction to write the proof for.
    /// @return dbNonce         The database nonce of the written entry for the proof.
    function writeExecutionProof(bytes32 transactionId) external returns (uint256 dbNonce);

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

    /// @notice Returns the fee for sending an Interchain message.
    /// @param dstChainId           The chain ID of the destination chain.
    /// @param srcExecutionService  The address of the execution service to use for the message.
    /// @param srcModules           The source modules involved in the message sending.
    /// @param options              Execution options for the message sent, currently gas limit + native gas drop.
    /// @param message              The message being sent.
    function getInterchainFee(
        uint256 dstChainId,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        view
        returns (uint256);

    /// @notice Returns the address of the executor for a transaction that has been sent to the local chain.
    function getExecutor(bytes calldata transaction) external view returns (address);

    /// @notice Returns the address of the executor for a transaction that has been sent to the local chain.
    function getExecutorById(bytes32 transactionId) external view returns (address);

    /// @notice Returns the address of the linked client (as bytes32) for a specific chain ID.
    /// @dev Will return 0x0 if no client is linked for the chain ID.
    function getLinkedClient(uint256 chainId) external view returns (bytes32);

    /// @notice Returns the EVM address of the linked client for a specific chain ID.
    /// @dev Will return 0x0 if no client is linked for the chain ID.
    /// Will revert if the client is not an EVM client.
    function getLinkedClientEVM(uint256 chainId) external view returns (address);
}
