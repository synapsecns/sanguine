// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1Events} from "./events/InterchainClientV1Events.sol";

import {IExecutionService} from "./interfaces/IExecutionService.sol";
import {IInterchainApp} from "./interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "./interfaces/IInterchainClientV1.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";

import {AppConfigV1, AppConfigLib} from "./libs/AppConfig.sol";
import {BatchingV1Lib} from "./libs/BatchingV1.sol";
import {InterchainBatch} from "./libs/InterchainBatch.sol";
import {
    InterchainTransaction, InterchainTxDescriptor, InterchainTransactionLib
} from "./libs/InterchainTransaction.sol";
import {OptionsLib, OptionsV1} from "./libs/Options.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";
import {VersionedPayloadLib} from "./libs/VersionedPayload.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title InterchainClientV1
 * @dev Implements the operations of the Interchain Execution Layer.
 */
contract InterchainClientV1 is Ownable, InterchainClientV1Events, IInterchainClientV1 {
    using AppConfigLib for bytes;
    using OptionsLib for bytes;
    using VersionedPayloadLib for bytes;

    /// @notice Version of the InterchainClient contract. Sent and received transactions must have the same version.
    uint16 public constant CLIENT_VERSION = 1;

    /// @notice Address of the InterchainDB contract, set at the time of deployment.
    address public immutable INTERCHAIN_DB;

    /// @dev Address of the InterchainClient contract on the remote chain
    mapping(uint64 chainId => bytes32 remoteClient) internal _linkedClient;
    /// @dev Executor address that completed the transaction. Address(0) if not executed yet.
    mapping(bytes32 transactionId => address executor) internal _txExecutor;

    constructor(address interchainDB, address owner_) Ownable(owner_) {
        INTERCHAIN_DB = interchainDB;
    }

    // @inheritdoc IInterchainClientV1
    function setLinkedClient(uint64 chainId, bytes32 client) external onlyOwner {
        _linkedClient[chainId] = client;
        emit LinkedClientSet(chainId, client);
    }

    // @inheritdoc IInterchainClientV1
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
        returns (InterchainTxDescriptor memory desc)
    {
        return _interchainSend(dstChainId, receiver, srcExecutionService, srcModules, options, message);
    }

    // @inheritdoc IInterchainClientV1
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
        returns (InterchainTxDescriptor memory desc)
    {
        bytes32 receiverBytes32 = TypeCasts.addressToBytes32(receiver);
        return _interchainSend(dstChainId, receiverBytes32, srcExecutionService, srcModules, options, message);
    }

    // TODO: Handle the case where receiver does not implement the IInterchainApp interface (or does not exist at all)
    // @inheritdoc IInterchainClientV1
    function interchainExecute(
        uint256 gasLimit,
        bytes calldata transaction,
        bytes32[] calldata proof
    )
        external
        payable
    {
        InterchainTransaction memory icTx = _assertCorrectVersion(transaction);
        bytes32 transactionId = keccak256(transaction);
        _assertExecutable(icTx, transactionId, proof);
        _txExecutor[transactionId] = msg.sender;

        OptionsV1 memory decodedOptions = icTx.options.decodeOptionsV1();
        if (msg.value != decodedOptions.gasAirdrop) {
            revert InterchainClientV1__IncorrectMsgValue(msg.value, decodedOptions.gasAirdrop);
        }
        // We should always use at least as much as the requested gas limit.
        // The executor can specify a higher gas limit if they wanted.
        if (decodedOptions.gasLimit > gasLimit) gasLimit = decodedOptions.gasLimit;
        // Check the the Executor has provided big enough gas limit for the whole transaction.
        if (gasleft() <= gasLimit) {
            revert InterchainClientV1__NotEnoughGasSupplied();
        }
        // Pass the full msg.value to the app: we have already checked that it matches the requested gas airdrop.
        IInterchainApp(TypeCasts.bytes32ToAddress(icTx.dstReceiver)).appReceive{gas: gasLimit, value: msg.value}({
            srcChainId: icTx.srcChainId,
            sender: icTx.srcSender,
            dbNonce: icTx.dbNonce,
            entryIndex: icTx.entryIndex,
            message: icTx.message
        });
        emit InterchainTransactionReceived(
            transactionId, icTx.dbNonce, icTx.entryIndex, icTx.srcChainId, icTx.srcSender, icTx.dstReceiver
        );
    }

    /// @inheritdoc IInterchainClientV1
    function writeExecutionProof(bytes32 transactionId) external returns (uint64 dbNonce, uint64 entryIndex) {
        address executor = _txExecutor[transactionId];
        if (executor == address(0)) {
            revert InterchainClientV1__TxNotExecuted(transactionId);
        }
        bytes memory proof = abi.encode(transactionId, executor);
        (dbNonce, entryIndex) = IInterchainDB(INTERCHAIN_DB).writeEntry(keccak256(proof));
        emit ExecutionProofWritten(transactionId, dbNonce, entryIndex, executor);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    // @inheritdoc IInterchainClientV1
    function isExecutable(bytes calldata encodedTx, bytes32[] calldata proof) external view returns (bool) {
        InterchainTransaction memory icTx = _assertCorrectVersion(encodedTx);
        // Check that options could be decoded
        icTx.options.decodeOptionsV1();
        bytes32 transactionId = keccak256(encodedTx);
        _assertExecutable(icTx, transactionId, proof);
        return true;
    }

    // @inheritdoc IInterchainClientV1
    function getExecutor(bytes calldata encodedTx) external view returns (address) {
        return _txExecutor[keccak256(encodedTx)];
    }

    // @inheritdoc IInterchainClientV1
    function getExecutorById(bytes32 transactionId) external view returns (address) {
        return _txExecutor[transactionId];
    }

    // @inheritdoc IInterchainClientV1
    function getInterchainFee(
        uint64 dstChainId,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        uint256 messageLen
    )
        external
        view
        returns (uint256 fee)
    {
        _assertLinkedClient(dstChainId);
        if (srcExecutionService == address(0)) {
            revert InterchainClientV1__ZeroExecutionService();
        }
        // Check that options could be decoded on destination chain
        options.decodeOptionsV1();
        // Verification fee from InterchainDB
        fee = IInterchainDB(INTERCHAIN_DB).getInterchainFee(dstChainId, srcModules);
        // Add execution fee from ExecutionService
        uint256 payloadSize = InterchainTransactionLib.payloadSize(options.length, messageLen);
        fee += IExecutionService(srcExecutionService).getExecutionFee(dstChainId, payloadSize, options);
    }

    /// @inheritdoc IInterchainClientV1
    function getLinkedClient(uint64 chainId) external view returns (bytes32) {
        if (chainId == block.chainid) {
            revert InterchainClientV1__NotRemoteChainId(chainId);
        }
        return _linkedClient[chainId];
    }

    /// @inheritdoc IInterchainClientV1
    function getLinkedClientEVM(uint64 chainId) external view returns (address linkedClientEVM) {
        if (chainId == block.chainid) {
            revert InterchainClientV1__NotRemoteChainId(chainId);
        }
        bytes32 linkedClient = _linkedClient[chainId];
        linkedClientEVM = TypeCasts.bytes32ToAddress(linkedClient);
        // Check that the linked client address fits into the EVM address space
        if (TypeCasts.addressToBytes32(linkedClientEVM) != linkedClient) {
            revert InterchainClientV1__NotEVMClient(linkedClient);
        }
    }

    /// @notice Gets the V1 app config and trusted modules for the receiving app.
    function getAppReceivingConfigV1(address receiver)
        external
        view
        returns (AppConfigV1 memory config, address[] memory modules)
    {
        bytes memory encodedConfig;
        (encodedConfig, modules) = IInterchainApp(receiver).getReceivingConfig();
        config = encodedConfig.decodeAppConfigV1();
    }

    /// @notice Decodes the encoded options data into a OptionsV1 struct.
    function decodeOptions(bytes memory encodedOptions) external view returns (OptionsV1 memory) {
        return encodedOptions.decodeOptionsV1();
    }

    /// @notice Encodes the transaction data into a bytes format.
    function encodeTransaction(InterchainTransaction memory icTx) public pure returns (bytes memory) {
        return VersionedPayloadLib.encodeVersionedPayload({
            version: CLIENT_VERSION,
            payload: InterchainTransactionLib.encodeTransaction(icTx)
        });
    }

    // ═════════════════════════════════════════════════ INTERNAL ══════════════════════════════════════════════════════

    /// @dev Internal logic for sending a message to another chain.
    function _interchainSend(
        uint64 dstChainId,
        bytes32 receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        internal
        returns (InterchainTxDescriptor memory desc)
    {
        _assertLinkedClient(dstChainId);
        if (receiver == 0) {
            revert InterchainClientV1__ZeroReceiver();
        }
        if (srcExecutionService == address(0)) {
            revert InterchainClientV1__ZeroExecutionService();
        }
        // Check that options could be decoded on destination chain
        options.decodeOptionsV1();
        uint256 verificationFee = IInterchainDB(INTERCHAIN_DB).getInterchainFee(dstChainId, srcModules);
        if (msg.value < verificationFee) {
            revert InterchainClientV1__FeeAmountTooLow(msg.value, verificationFee);
        }
        (desc.dbNonce, desc.entryIndex) = IInterchainDB(INTERCHAIN_DB).getNextEntryIndex();
        InterchainTransaction memory icTx = InterchainTransactionLib.constructLocalTransaction({
            srcSender: msg.sender,
            dstReceiver: receiver,
            dstChainId: dstChainId,
            dbNonce: desc.dbNonce,
            entryIndex: desc.entryIndex,
            options: options,
            message: message
        });
        desc.transactionId = keccak256(encodeTransaction(icTx));
        // Sanity check: nonce returned from DB should match the nonce used to construct the transaction
        {
            (uint64 dbNonce, uint64 entryIndex) = IInterchainDB(INTERCHAIN_DB).writeEntryWithVerification{
                value: verificationFee
            }(icTx.dstChainId, desc.transactionId, srcModules);
            assert(dbNonce == desc.dbNonce && entryIndex == desc.entryIndex);
        }
        uint256 executionFee;
        unchecked {
            executionFee = msg.value - verificationFee;
        }
        IExecutionService(srcExecutionService).requestTxExecution{value: executionFee}(
            icTx.dstChainId,
            InterchainTransactionLib.payloadSize(options.length, message.length),
            desc.transactionId,
            options
        );
        emit InterchainTransactionSent(
            desc.transactionId,
            icTx.dbNonce,
            icTx.entryIndex,
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

    /// @dev Asserts that the transaction is executable.
    function _assertExecutable(
        InterchainTransaction memory icTx,
        bytes32 transactionId,
        bytes32[] calldata proof
    )
        internal
        view
    {
        bytes32 linkedClient = _assertLinkedClient(icTx.srcChainId);
        if (icTx.dstChainId != block.chainid) {
            revert InterchainClientV1__IncorrectDstChainId(icTx.dstChainId);
        }
        if (_txExecutor[transactionId] != address(0)) {
            revert InterchainClientV1__TxAlreadyExecuted(transactionId);
        }
        // Construct expected batch based on interchain transaction data
        InterchainBatch memory batch = InterchainBatch({
            srcChainId: icTx.srcChainId,
            dbNonce: icTx.dbNonce,
            batchRoot: BatchingV1Lib.getBatchRoot({
                srcWriter: linkedClient,
                dataHash: transactionId,
                entryIndex: icTx.entryIndex,
                proof: proof
            })
        });
        (bytes memory encodedAppConfig, address[] memory approvedDstModules) =
            IInterchainApp(TypeCasts.bytes32ToAddress(icTx.dstReceiver)).getReceivingConfig();
        AppConfigV1 memory appConfig = encodedAppConfig.decodeAppConfigV1();
        if (appConfig.requiredResponses == 0) {
            revert InterchainClientV1__ZeroRequiredResponses();
        }
        uint256 responses = _getFinalizedResponsesCount(approvedDstModules, batch, appConfig.optimisticPeriod);
        if (responses < appConfig.requiredResponses) {
            revert InterchainClientV1__NotEnoughResponses(responses, appConfig.requiredResponses);
        }
    }

    /// @dev Asserts that the chain is linked and returns the linked client address.
    function _assertLinkedClient(uint64 chainId) internal view returns (bytes32 linkedClient) {
        if (chainId == block.chainid) {
            revert InterchainClientV1__NotRemoteChainId(chainId);
        }
        linkedClient = _linkedClient[chainId];
        if (linkedClient == 0) {
            revert InterchainClientV1__NoLinkedClient(chainId);
        }
    }

    /**
     * @dev Calculates the number of responses that are considered finalized within the optimistic time period.
     * @param approvedModules       Approved modules that could have confirmed the entry.
     * @param batch                 The Interchain Batch to confirm.
     * @param optimisticPeriod      The time period in seconds within which a response is considered valid.
     * @return finalizedResponses   The count of responses that are finalized within the optimistic time period.
     */
    function _getFinalizedResponsesCount(
        address[] memory approvedModules,
        InterchainBatch memory batch,
        uint256 optimisticPeriod
    )
        internal
        view
        returns (uint256 finalizedResponses)
    {
        for (uint256 i = 0; i < approvedModules.length; ++i) {
            uint256 confirmedAt = IInterchainDB(INTERCHAIN_DB).checkBatchVerification(approvedModules[i], batch);
            // checkVerification() returns 0 if entry hasn't been confirmed by the module, so we check for that as well
            if (confirmedAt != 0 && confirmedAt + optimisticPeriod < block.timestamp) {
                ++finalizedResponses;
            }
        }
    }

    /// @dev Asserts that the transaction version is correct. Returns the decoded transaction for chaining purposes.
    function _assertCorrectVersion(bytes calldata versionedTx)
        internal
        pure
        returns (InterchainTransaction memory icTx)
    {
        uint16 version = versionedTx.getVersion();
        if (version != CLIENT_VERSION) {
            revert InterchainClientV1__InvalidTransactionVersion(version);
        }
        icTx = InterchainTransactionLib.decodeTransaction(versionedTx.getPayload());
    }
}
