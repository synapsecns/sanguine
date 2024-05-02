// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainTransaction, InterchainTxDescriptor} from "../libs/InterchainTransaction.sol";

interface IInterchainClientV1 {
    // TODO: recheck the explicitness of the TxReadiness enum
    enum TxReadiness {
        Ready,
        AlreadyExecuted,
        BatchAwaitingResponses,
        BatchConflict,
        ReceiverNotICApp,
        ReceiverZeroRequiredResponses,
        TxWrongDstChainId,
        UndeterminedRevert
    }

    // TODO: standardize error names across interfaces
    error InterchainClientV1__BatchConflict(address module);
    error InterchainClientV1__FeeAmountTooLow(uint256 actual, uint256 required);
    error InterchainClientV1__IncorrectDstChainId(uint64 chainId);
    error InterchainClientV1__IncorrectMsgValue(uint256 actual, uint256 required);
    error InterchainClientV1__InvalidTransactionVersion(uint16 version);
    error InterchainClientV1__NoLinkedClient(uint64 chainId);
    error InterchainClientV1__NotEnoughGasSupplied();
    error InterchainClientV1__NotEnoughResponses(uint256 actual, uint256 required);
    error InterchainClientV1__NotEVMClient(bytes32 client);
    error InterchainClientV1__NotRemoteChainId(uint64 chainId);
    error InterchainClientV1__ReceiverNotICApp(address receiver);
    error InterchainClientV1__TxAlreadyExecuted(bytes32 transactionId);
    error InterchainClientV1__TxNotExecuted(bytes32 transactionId);
    error InterchainClientV1__ZeroAddress();
    error InterchainClientV1__ZeroExecutionService();
    error InterchainClientV1__ZeroReceiver();
    error InterchainClientV1__ZeroRequiredResponses();

    /// @notice Allows the contract owner to set the address of the Guard module.
    /// Note: batches marked as invalid by the Guard could not be used for message execution,
    /// if the app opts in to use the Guard.
    /// @param guard_       The address of the Guard module.
    function setDefaultGuard(address guard_) external;

    /**
     * @notice Sets the linked client for a specific chain ID.
     * @dev Stores the address of the linked client in a mapping with the chain ID as the key.
     * @param chainId The chain ID for which the client is being set.
     * @param client The address of the client being linked.
     */
    function setLinkedClient(uint64 chainId, bytes32 client) external;

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
     * @return desc The descriptor of the sent transaction:
     * - transactionId: the ID of the transaction that was sent.
     * - dbNonce: the database nonce of the batch containing the written entry for transaction.
     * - entryIndex: the index of the written entry for transaction within the batch.
     */
    function interchainSend(
        uint64 dstChainId,
        bytes32 receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        payable
        returns (InterchainTxDescriptor memory desc);

    function interchainSendEVM(
        uint64 dstChainId,
        address receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        payable
        returns (InterchainTxDescriptor memory desc);

    /**
     * @notice Executes a transaction that has been sent via the Interchain.
     * @dev The transaction must have been previously sent and recorded.
     * Transaction data includes the requested gas limit, but the executors could specify a different gas limit.
     * If the specified gas limit is lower than requested, the requested gas limit will be used.
     * Otherwise, the specified gas limit will be used.
     * This allows to execute the transactions with requested gas limit set too low.
     * @param gasLimit          The gas limit to use for the execution.
     * @param transaction       The transaction data.
     * @param proof             The Merkle proof for transaction execution, fetched from the source chain.
     */
    function interchainExecute(
        uint256 gasLimit,
        bytes calldata transaction,
        bytes32[] calldata proof
    )
        external
        payable;

    /// @notice Writes the proof of execution for a transaction into the InterchainDB.
    /// @dev Will revert if the transaction has not been executed.
    /// @param transactionId    The ID of the transaction to write the proof for.
    /// @return dbNonce         The database nonce of the batch containing the written proof for transaction.
    /// @return entryIndex      The index of the written proof for transaction within the batch.
    function writeExecutionProof(bytes32 transactionId) external returns (uint64 dbNonce, uint64 entryIndex);

    /**
     * @notice Checks if a transaction is executable.
     * @dev Determines if a transaction meets the criteria to be executed based on:
     * - If approved modules have written to the InterchainDB
     * - If the threshold of approved modules have been met
     * - If the optimistic window has passed for all modules
     * @param transaction       The InterchainTransaction struct to be checked.
     * @param proof             The Merkle proof for transaction execution, fetched from the source chain.
     * @return bool Returns true if the transaction is executable, false otherwise.
     */
    function isExecutable(bytes calldata transaction, bytes32[] calldata proof) external view returns (bool);

    /// @notice Returns the readiness status of a transaction to be executed.
    /// @dev Some of the possible statuses have additional arguments that are returned:
    /// - Ready: the transaction is ready to be executed.
    /// - AlreadyExecuted: the transaction has already been executed.
    ///   - `firstArg` is the transaction ID.
    /// - BatchAwaitingResponses: not enough responses have been received for the transaction.
    ///   - `firstArg` is the number of responses received.
    ///   - `secondArg` is the number of responses required.
    /// - BatchConflict: one of the modules have submitted a conflicting batch.
    ///   - `firstArg` is the address of the module.
    ///   - This is either one of the modules that the app trusts, or the Guard module used by the app.
    /// - ReceiverNotICApp: the receiver is not an Interchain app.
    ///  - `firstArg` is the receiver address.
    /// - ReceiverZeroRequiredResponses: the app config requires zero responses for the transaction.
    /// - TxWrongDstChainId: the destination chain ID does not match the local chain ID.
    ///   - `firstArg` is the destination chain ID.
    /// - UndeterminedRevert: the transaction will revert for another reason.
    ///
    /// Note: the arguments are abi-encoded bytes32 values (as their types could be different).
    function getTxReadinessV1(
        InterchainTransaction memory icTx,
        bytes32[] calldata proof
    )
        external
        view
        returns (TxReadiness status, bytes32 firstArg, bytes32 secondArg);

    /// @notice Returns the fee for sending an Interchain message.
    /// @param dstChainId           The chain ID of the destination chain.
    /// @param srcExecutionService  The address of the execution service to use for the message.
    /// @param srcModules           The source modules involved in the message sending.
    /// @param options              Execution options for the message sent, currently gas limit + native gas drop.
    /// @param messageLen           The length of the message being sent.
    function getInterchainFee(
        uint64 dstChainId,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        uint256 messageLen
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
    function getLinkedClient(uint64 chainId) external view returns (bytes32);

    /// @notice Returns the EVM address of the linked client for a specific chain ID.
    /// @dev Will return 0x0 if no client is linked for the chain ID.
    /// Will revert if the client is not an EVM client.
    function getLinkedClientEVM(uint64 chainId) external view returns (address);
}
