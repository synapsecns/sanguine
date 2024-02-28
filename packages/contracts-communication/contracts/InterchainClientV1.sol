// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1Events} from "./events/InterchainClientV1Events.sol";

import {IExecutionFees} from "./interfaces/IExecutionFees.sol";
import {IExecutionService} from "./interfaces/IExecutionService.sol";
import {IInterchainApp} from "./interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "./interfaces/IInterchainClientV1.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";

import {AppConfigV1, AppConfigLib} from "./libs/AppConfig.sol";
import {InterchainEntry} from "./libs/InterchainEntry.sol";
import {InterchainTransaction, InterchainTransactionLib} from "./libs/InterchainTransaction.sol";
import {OptionsLib, OptionsV1} from "./libs/Options.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title InterchainClientV1
 * @dev Implements the operations of the Interchain Execution Layer.
 */
contract InterchainClientV1 is Ownable, InterchainClientV1Events, IInterchainClientV1 {
    using AppConfigLib for bytes;
    using OptionsLib for bytes;

    /// @notice Address of the InterchainDB contract, set at the time of deployment.
    address public immutable INTERCHAIN_DB;

    /// @notice Address of the contract that handles execution fees. Can be updated by the owner.
    address public executionFees;

    /// @notice Address of the InterchainClient contract on the remote chain
    mapping(uint256 chainId => bytes32 remoteClient) public linkedClients;
    /// @dev Executor address that completed the transaction. Address(0) if not executed yet.
    mapping(bytes32 transactionId => address executor) internal _txExecutor;

    constructor(address interchainDB) Ownable(msg.sender) {
        INTERCHAIN_DB = interchainDB;
    }

    // @inheritdoc IInterchainClientV1
    function setExecutionFees(address executionFees_) external onlyOwner {
        executionFees = executionFees_;
    }

    // @inheritdoc IInterchainClientV1
    function setLinkedClient(uint256 chainId, bytes32 client) external onlyOwner {
        linkedClients[chainId] = client;
    }

    // @inheritdoc IInterchainClientV1
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
        returns (bytes32 transactionId, uint256 dbNonce)
    {
        return _interchainSend(dstChainId, receiver, srcExecutionService, srcModules, options, message);
    }

    // @inheritdoc IInterchainClientV1
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
        returns (bytes32 transactionId, uint256 dbNonce)
    {
        bytes32 receiverBytes32 = TypeCasts.addressToBytes32(receiver);
        return _interchainSend(dstChainId, receiverBytes32, srcExecutionService, srcModules, options, message);
    }

    // TODO: Handle the case where receiver does not implement the IInterchainApp interface (or does not exist at all)
    // TODO: Save the executor address outside of the contract to pass the data back to the source chain
    // @inheritdoc IInterchainClientV1
    function interchainExecute(uint256 gasLimit, bytes calldata transaction) external payable {
        InterchainTransaction memory icTx = InterchainTransactionLib.decodeTransaction(transaction);
        bytes32 transactionId = _assertExecutable(icTx);
        _txExecutor[transactionId] = msg.sender;

        OptionsV1 memory decodedOptions = icTx.options.decodeOptionsV1();
        if (msg.value != decodedOptions.gasAirdrop) {
            revert InterchainClientV1__IncorrectMsgValue(msg.value, decodedOptions.gasAirdrop);
        }
        // We should always use at least as much as the requested gas limit.
        // The executor can specify a higher gas limit if they wanted.
        if (decodedOptions.gasLimit > gasLimit) gasLimit = decodedOptions.gasLimit;
        // Pass the full msg.value to the app: we have already checked that it matches the requested gas airdrop.
        IInterchainApp(TypeCasts.bytes32ToAddress(icTx.dstReceiver)).appReceive{gas: gasLimit, value: msg.value}({
            srcChainId: icTx.srcChainId,
            sender: icTx.srcSender,
            dbNonce: icTx.dbNonce,
            message: icTx.message
        });
        emit InterchainTransactionReceived(
            transactionId, icTx.dbNonce, icTx.srcChainId, icTx.srcSender, icTx.dstReceiver
        );
    }

    /// @inheritdoc IInterchainClientV1
    function writeExecutionProof(bytes32 transactionId) external returns (uint256 dbNonce) {
        address executor = _txExecutor[transactionId];
        if (executor == address(0)) {
            revert InterchainClientV1__TxNotExecuted(transactionId);
        }
        bytes memory proof = abi.encode(transactionId, executor);
        dbNonce = IInterchainDB(INTERCHAIN_DB).writeEntry(keccak256(proof));
        emit ExecutionProofWritten(transactionId, dbNonce, executor);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    // @inheritdoc IInterchainClientV1
    function isExecutable(bytes calldata encodedTx) external view returns (bool) {
        InterchainTransaction memory icTx = InterchainTransactionLib.decodeTransaction(encodedTx);
        _assertExecutable(icTx);
        return true;
    }

    // @inheritdoc IInterchainClientV1
    function getExecutor(bytes calldata encodedTx) external view returns (address) {
        InterchainTransaction memory icTx = InterchainTransactionLib.decodeTransaction(encodedTx);
        return _txExecutor[icTx.transactionId()];
    }

    // @inheritdoc IInterchainClientV1
    function getExecutorById(bytes32 transactionId) external view returns (address) {
        return _txExecutor[transactionId];
    }

    // @inheritdoc IInterchainClientV1
    // TODO: tests
    function getInterchainFee(
        uint256 dstChainId,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        view
        returns (uint256 fee)
    {
        // Verification fee from InterchainDB
        fee = IInterchainDB(INTERCHAIN_DB).getInterchainFee(dstChainId, srcModules);
        // Add execution fee, if ExecutionService is provided
        if (srcExecutionService != address(0)) {
            // Construct a mock InterchainTransaction to calculate the execution fee.
            // We don't care about values for static fields, as we are only interested in the payload size.
            InterchainTransaction memory icTx = InterchainTransactionLib.constructLocalTransaction({
                srcSender: address(0),
                dstReceiver: 0,
                dstChainId: dstChainId,
                dbNonce: 0,
                options: options,
                message: message
            });
            fee += IExecutionService(srcExecutionService).getExecutionFee(dstChainId, abi.encode(icTx).length, options);
        }
    }

    /// @notice Encodes the transaction data into a bytes format.
    function encodeTransaction(InterchainTransaction memory icTx) external pure returns (bytes memory) {
        return icTx.encodeTransaction();
    }

    /// @notice Decodes the encoded options data into a OptionsV1 struct.
    function decodeOptions(bytes memory encodedOptions) external pure returns (OptionsV1 memory) {
        return encodedOptions.decodeOptionsV1();
    }

    // ═════════════════════════════════════════════════ INTERNAL ══════════════════════════════════════════════════════

    /// @dev Internal logic for sending a message to another chain.
    function _interchainSend(
        uint256 dstChainId,
        bytes32 receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        internal
        returns (bytes32 transactionId, uint256 dbNonce)
    {
        if (dstChainId == block.chainid) {
            revert InterchainClientV1__IncorrectDstChainId(dstChainId);
        }
        // TODO: should check options for being correctly formatted
        uint256 verificationFee = IInterchainDB(INTERCHAIN_DB).getInterchainFee(dstChainId, srcModules);
        if (msg.value < verificationFee) {
            revert InterchainClientV1__FeeAmountTooLow(msg.value, verificationFee);
        }
        uint256 executionFee;
        unchecked {
            executionFee = msg.value - verificationFee;
        }
        dbNonce = IInterchainDB(INTERCHAIN_DB).getDBNonce();
        InterchainTransaction memory icTx = InterchainTransactionLib.constructLocalTransaction({
            srcSender: msg.sender,
            dstReceiver: receiver,
            dstChainId: dstChainId,
            dbNonce: dbNonce,
            options: options,
            message: message
        });
        transactionId = icTx.transactionId();
        // Sanity check: nonce returned from DB should match the nonce used to construct the transaction
        assert(
            dbNonce
                == IInterchainDB(INTERCHAIN_DB).writeEntryWithVerification{value: verificationFee}(
                    icTx.dstChainId, transactionId, srcModules
                )
        );
        IExecutionFees(executionFees).addExecutionFee{value: executionFee}(icTx.dstChainId, transactionId);
        // TODO: consider disallowing the use of empty srcExecutionService
        if (srcExecutionService != address(0)) {
            IExecutionService(srcExecutionService).requestExecution({
                dstChainId: dstChainId,
                // TODO: there should be a way to calculate the payload size without encoding the transaction
                txPayloadSize: icTx.encodeTransaction().length,
                transactionId: transactionId,
                executionFee: executionFee,
                options: options
            });
            address srcExecutorEOA = IExecutionService(srcExecutionService).executorEOA();
            IExecutionFees(executionFees).recordExecutor(icTx.dstChainId, transactionId, srcExecutorEOA);
        }
        emit InterchainTransactionSent(
            transactionId,
            icTx.dbNonce,
            icTx.dstChainId,
            icTx.srcSender,
            icTx.dstReceiver,
            verificationFee,
            executionFee,
            icTx.options,
            icTx.message
        );
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Asserts that the transaction is executable. Returns the transactionId for chaining purposes.
    function _assertExecutable(InterchainTransaction memory icTx) internal view returns (bytes32 transactionId) {
        if (icTx.srcChainId == block.chainid) {
            revert InterchainClientV1__IncorrectSrcChainId(icTx.srcChainId);
        }
        if (icTx.dstChainId != block.chainid) {
            revert InterchainClientV1__IncorrectDstChainId(icTx.dstChainId);
        }
        transactionId = icTx.transactionId();
        if (_txExecutor[transactionId] != address(0)) {
            revert InterchainClientV1__TxAlreadyExecuted(transactionId);
        }
        // Construct expected entry based on icTransaction data
        InterchainEntry memory icEntry = InterchainEntry({
            srcChainId: icTx.srcChainId,
            dbNonce: icTx.dbNonce,
            srcWriter: linkedClients[icTx.srcChainId],
            dataHash: transactionId
        });
        (bytes memory encodedAppConfig, address[] memory approvedDstModules) =
            IInterchainApp(TypeCasts.bytes32ToAddress(icTx.dstReceiver)).getReceivingConfig();
        AppConfigV1 memory appConfig = encodedAppConfig.decodeAppConfigV1();
        uint256 responses = _getFinalizedResponsesCount(approvedDstModules, icEntry, appConfig.optimisticPeriod);
        revert(string(abi.encode(responses)));
//        if (responses < appConfig.requiredResponses) {
//            revert InterchainClientV1__NotEnoughResponses(responses, appConfig.requiredResponses);
//        }
    }

    /**
     * @dev Calculates the number of responses that are considered finalized within the optimistic time period.
     * @param approvedModules       Approved modules that could have confirmed the entry.
     * @param icEntry               The InterchainEntry to confirm.
     * @param optimisticPeriod      The time period in seconds within which a response is considered valid.
     * @return finalizedResponses   The count of responses that are finalized within the optimistic time period.
     */
    function _getFinalizedResponsesCount(
        address[] memory approvedModules,
        InterchainEntry memory icEntry,
        uint256 optimisticPeriod
    )
        internal
        view
        returns (uint256 finalizedResponses)
    {
        for (uint256 i = 0; i < approvedModules.length; ++i) {
            uint256 confirmedAt = IInterchainDB(INTERCHAIN_DB).readEntry(approvedModules[i], icEntry);
            // readEntry() returns 0 if entry hasn't been confirmed by the module, so we check for that as well
            if (confirmedAt != 0 && confirmedAt + optimisticPeriod <= block.timestamp) {
                ++finalizedResponses;
            }
        }
    }
}
