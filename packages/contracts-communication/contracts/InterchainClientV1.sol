// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1Events} from "./events/InterchainClientV1Events.sol";

import {IExecutionService} from "./interfaces/IExecutionService.sol";
import {IInterchainApp} from "./interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "./interfaces/IInterchainClientV1.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";

import {AppConfigV1, AppConfigLib, APP_CONFIG_GUARD_DISABLED, APP_CONFIG_GUARD_DEFAULT} from "./libs/AppConfig.sol";
import {InterchainEntry, InterchainEntryLib, ENTRY_UNVERIFIED, ENTRY_CONFLICT} from "./libs/InterchainEntry.sol";
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
    using TypeCasts for address;
    using TypeCasts for bytes32;
    using VersionedPayloadLib for bytes;

    /// @notice Version of the InterchainClient contract. Sent and received transactions must have the same version.
    uint16 public constant CLIENT_VERSION = 1;

    /// @notice Address of the InterchainDB contract, set at the time of deployment.
    address public immutable INTERCHAIN_DB;

    /// @notice Address of the Guard module used to verify the validity of entries.
    /// Note: entries marked as invalid by the Guard could not be used for message execution,
    /// if the app opts in to use the Guard.
    address public defaultGuard;

    /// @notice Address of the default module to use to verify the validity of entries.
    /// Note: this module will be used for the apps that define an empty module list in their config.
    address public defaultModule;

    /// @dev Address of the InterchainClient contract on the remote chain
    mapping(uint64 chainId => bytes32 remoteClient) internal _linkedClient;
    /// @dev Executor address that completed the transaction. Address(0) if not executed yet.
    mapping(bytes32 transactionId => address executor) internal _txExecutor;

    constructor(address interchainDB, address owner_) Ownable(owner_) {
        INTERCHAIN_DB = interchainDB;
    }

    /// @notice Allows the contract owner to set the address of the Guard module.
    /// Note: entries marked as invalid by the Guard could not be used for message execution,
    /// if the app opts in to use the Guard.
    /// @param guard            The address of the Guard module.
    function setDefaultGuard(address guard) external onlyOwner {
        if (guard == address(0)) {
            revert InterchainClientV1__GuardZeroAddress();
        }
        defaultGuard = guard;
        emit DefaultGuardSet(guard);
    }

    /// @notice Allows the contract owner to set the address of the default module.
    /// Note: this module will be used for the apps that define an empty module list in their config.
    /// @param module           The address of the default module.
    function setDefaultModule(address module) external onlyOwner {
        if (module == address(0)) {
            revert InterchainClientV1__ModuleZeroAddress();
        }
        defaultModule = module;
        emit DefaultModuleSet(module);
    }

    /// @notice Sets the linked client for a specific chain ID.
    /// Note: only Interchain Entries written by the linked client could be used for message execution.
    /// @param chainId          The chain ID for which the client is being set.
    /// @param client           The address of the client being linked.
    function setLinkedClient(uint64 chainId, bytes32 client) external onlyOwner {
        _linkedClient[chainId] = client;
        emit LinkedClientSet(chainId, client);
    }

    /// @notice Sends a message to another chain via the Interchain Communication Protocol.
    /// @dev Charges a fee for the message, which is payable upon calling this function:
    /// - Verification fees: paid to every module that verifies the message.
    /// - Execution fee: paid to the executor that executes the message.
    /// Note: while a specific execution service is specified to request the execution of the message,
    /// any executor is able to execute the message on destination chain.
    /// @param dstChainId           The chain ID of the destination chain.
    /// @param receiver             The address of the receiver on the destination chain.
    /// @param srcExecutionService  The address of the execution service to use for the message.
    /// @param srcModules           The source modules involved in the message sending.
    /// @param options              Execution options for the message sent, encoded as bytes,
    ///                             currently gas limit + native gas drop.
    /// @param message              The message to be sent.
    /// @return desc                The descriptor of the sent transaction:
    /// - transactionId: the ID of the transaction that was sent.
    /// - dbNonce: the database nonce of the entry containing the transaction.
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

    /// @notice A thin wrapper around `interchainSend` that allows to specify the receiver address as an EVM address.
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
        bytes32 receiverBytes32 = receiver.addressToBytes32();
        return _interchainSend(dstChainId, receiverBytes32, srcExecutionService, srcModules, options, message);
    }

    /// @notice Executes a transaction that has been sent via the Interchain Communication Protocol.
    /// Note: The transaction must be proven to be included in one of the InterchainDB entries.
    /// Note: Transaction data includes the requested gas limit, but the executors could specify a different gas limit.
    /// If the specified gas limit is lower than requested, the requested gas limit will be used.
    /// Otherwise, the specified gas limit will be used.
    /// This allows to execute the transactions with requested gas limit set too low.
    /// @param gasLimit          The gas limit to use for the execution.
    /// @param transaction       The transaction data.
    function interchainExecute(uint256 gasLimit, bytes calldata transaction) external payable {
        InterchainTransaction memory icTx = _assertCorrectTransaction(transaction);
        bytes32 transactionId = keccak256(transaction);
        _assertExecutable(icTx, transactionId);
        _txExecutor[transactionId] = msg.sender;

        OptionsV1 memory decodedOptions = icTx.options.decodeOptionsV1();
        if (msg.value != decodedOptions.gasAirdrop) {
            revert InterchainClientV1__MsgValueMismatch(msg.value, decodedOptions.gasAirdrop);
        }
        // We should always use at least as much as the requested gas limit.
        // The executor can specify a higher gas limit if they wanted.
        if (decodedOptions.gasLimit > gasLimit) gasLimit = decodedOptions.gasLimit;
        // Check the the Executor has provided big enough gas limit for the whole transaction.
        uint256 gasLeft = gasleft();
        if (gasLeft <= gasLimit) {
            revert InterchainClientV1__GasLeftBelowMin(gasLeft, gasLimit);
        }
        // Pass the full msg.value to the app: we have already checked that it matches the requested gas airdrop.
        IInterchainApp(icTx.dstReceiver.bytes32ToAddress()).appReceive{gas: gasLimit, value: msg.value}({
            srcChainId: icTx.srcChainId,
            sender: icTx.srcSender,
            dbNonce: icTx.dbNonce,
            message: icTx.message
        });
        emit InterchainTransactionReceived({
            transactionId: transactionId,
            dbNonce: icTx.dbNonce,
            srcChainId: icTx.srcChainId,
            srcSender: icTx.srcSender,
            dstReceiver: icTx.dstReceiver
        });
    }

    /// @notice Writes the proof of execution for a transaction into the InterchainDB.
    /// @dev Will revert if the transaction has not been executed.
    /// @param transactionId    The ID of the transaction to write the proof for.
    /// @return dbNonce         The database nonce of the entry containing the written proof for transaction.
    function writeExecutionProof(bytes32 transactionId) external returns (uint64 dbNonce) {
        address executor = _txExecutor[transactionId];
        if (executor == address(0)) {
            revert InterchainClientV1__TxNotExecuted(transactionId);
        }
        bytes memory proof = abi.encode(transactionId, executor);
        dbNonce = IInterchainDB(INTERCHAIN_DB).writeEntry(keccak256(proof));
        emit ExecutionProofWritten({transactionId: transactionId, dbNonce: dbNonce, executor: executor});
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Determines if a transaction meets the criteria to be executed based on:
    /// - If approved modules have verified the entry in the InterchainDB
    /// - If the threshold of approved modules have been met
    /// - If the optimistic window has passed for all modules
    /// - If the Guard module (if opted in) has not submitted an entry that conflicts with the approved modules
    /// @dev Will revert with a specific error message if the transaction is not executable.
    /// @param encodedTx        The encoded transaction to check for executable status.
    function isExecutable(bytes calldata encodedTx) external view returns (bool) {
        InterchainTransaction memory icTx = _assertCorrectTransaction(encodedTx);
        // Check that options could be decoded
        icTx.options.decodeOptionsV1();
        bytes32 transactionId = keccak256(encodedTx);
        _assertExecutable(icTx, transactionId);
        return true;
    }

    /// @notice Returns the readiness status of a transaction to be executed.
    /// @dev Some of the possible statuses have additional arguments that are returned:
    /// - Ready: the transaction is ready to be executed.
    /// - AlreadyExecuted: the transaction has already been executed.
    ///   - `firstArg` is the transaction ID.
    /// - EntryAwaitingResponses: not enough responses have been received for the transaction.
    ///   - `firstArg` is the number of responses received.
    ///   - `secondArg` is the number of responses required.
    /// - EntryConflict: one of the modules have submitted a conflicting entry.
    ///   - `firstArg` is the address of the module.
    ///   - This is either one of the modules that the app trusts, or the Guard module used by the app.
    /// - ReceiverNotICApp: the receiver is not an Interchain app.
    ///  - `firstArg` is the receiver address.
    /// - TxWrongDstChainId: the destination chain ID does not match the local chain ID.
    ///   - `firstArg` is the destination chain ID.
    /// - UndeterminedRevert: the transaction will revert for another reason.
    ///
    /// Note: the arguments are abi-encoded bytes32 values (as their types could be different).
    // solhint-disable-next-line code-complexity
    function getTxReadinessV1(InterchainTransaction memory icTx)
        external
        view
        returns (TxReadiness status, bytes32 firstArg, bytes32 secondArg)
    {
        bytes memory encodedTx = encodeTransaction(icTx);
        try this.isExecutable(encodedTx) returns (bool) {
            return (TxReadiness.Ready, 0, 0);
        } catch (bytes memory errorData) {
            bytes4 selector;
            (selector, firstArg, secondArg) = _decodeRevertData(errorData);
            if (selector == InterchainClientV1__TxAlreadyExecuted.selector) {
                status = TxReadiness.AlreadyExecuted;
            } else if (selector == InterchainClientV1__ResponsesAmountBelowMin.selector) {
                status = TxReadiness.EntryAwaitingResponses;
            } else if (selector == InterchainClientV1__EntryConflict.selector) {
                status = TxReadiness.EntryConflict;
            } else if (selector == InterchainClientV1__ReceiverNotICApp.selector) {
                status = TxReadiness.ReceiverNotICApp;
            } else if (selector == InterchainClientV1__DstChainIdNotLocal.selector) {
                status = TxReadiness.TxWrongDstChainId;
            } else {
                status = TxReadiness.UndeterminedRevert;
                firstArg = 0;
                secondArg = 0;
            }
        }
    }

    /// @notice Returns the address of the executor for a transaction that has been sent to the local chain.
    function getExecutor(bytes calldata encodedTx) external view returns (address) {
        return _txExecutor[keccak256(encodedTx)];
    }

    /// @notice Returns the address of the executor for a transaction that has been sent to the local chain.
    function getExecutorById(bytes32 transactionId) external view returns (address) {
        return _txExecutor[transactionId];
    }

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
        returns (uint256 fee)
    {
        _assertLinkedClient(dstChainId);
        if (srcExecutionService == address(0)) {
            revert InterchainClientV1__ExecutionServiceZeroAddress();
        }
        // Check that options could be decoded on destination chain
        options.decodeOptionsV1();
        // Verification fee from InterchainDB
        fee = IInterchainDB(INTERCHAIN_DB).getInterchainFee(dstChainId, srcModules);
        // Add execution fee from ExecutionService
        uint256 payloadSize = InterchainTransactionLib.payloadSize(options.length, messageLen);
        fee += IExecutionService(srcExecutionService).getExecutionFee(dstChainId, payloadSize, options);
    }

    /// @notice Returns the address of the linked client (as bytes32) for a specific chain ID.
    /// @dev Will return 0x0 if no client is linked for the chain ID.
    function getLinkedClient(uint64 chainId) external view returns (bytes32) {
        if (chainId == block.chainid) {
            revert InterchainClientV1__ChainIdNotRemote(chainId);
        }
        return _linkedClient[chainId];
    }

    /// @notice Returns the EVM address of the linked client for a specific chain ID.
    /// @dev Will return 0x0 if no client is linked for the chain ID.
    /// Will revert if the client is not an EVM client.
    function getLinkedClientEVM(uint64 chainId) external view returns (address linkedClientEVM) {
        if (chainId == block.chainid) {
            revert InterchainClientV1__ChainIdNotRemote(chainId);
        }
        bytes32 linkedClient = _linkedClient[chainId];
        linkedClientEVM = linkedClient.bytes32ToAddress();
        // Check that the linked client address fits into the EVM address space
        if (linkedClientEVM.addressToBytes32() != linkedClient) {
            revert InterchainClientV1__LinkedClientNotEVM(linkedClient);
        }
    }

    /// @notice Decodes the encoded options data into a OptionsV1 struct.
    function decodeOptions(bytes memory encodedOptions) external view returns (OptionsV1 memory) {
        return encodedOptions.decodeOptionsV1();
    }

    /// @notice Gets the V1 app config and trusted modules for the receiving app.
    function getAppReceivingConfigV1(address receiver)
        public
        view
        returns (AppConfigV1 memory config, address[] memory modules)
    {
        // First, check that receiver is a contract
        if (receiver.code.length == 0) {
            revert InterchainClientV1__ReceiverNotICApp(receiver);
        }
        // Then, use a low-level static call to get the config and modules
        (bool success, bytes memory returnData) =
            receiver.staticcall(abi.encodeCall(IInterchainApp.getReceivingConfig, ()));
        if (!success || returnData.length == 0) {
            revert InterchainClientV1__ReceiverNotICApp(receiver);
        }
        bytes memory encodedConfig;
        (encodedConfig, modules) = abi.decode(returnData, (bytes, address[]));
        config = encodedConfig.decodeAppConfigV1();
        // Fallback to the default module if the app has no modules
        if (modules.length == 0) {
            modules = new address[](1);
            modules[0] = defaultModule;
        }
        // Fallback to "all responses" if the app requires zero responses
        if (config.requiredResponses == 0) {
            config.requiredResponses = modules.length;
        }
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
            revert InterchainClientV1__ReceiverZeroAddress();
        }
        if (srcExecutionService == address(0)) {
            revert InterchainClientV1__ExecutionServiceZeroAddress();
        }
        // Check that options could be decoded on destination chain
        options.decodeOptionsV1();
        uint256 verificationFee = IInterchainDB(INTERCHAIN_DB).getInterchainFee(dstChainId, srcModules);
        if (msg.value < verificationFee) {
            revert InterchainClientV1__FeeAmountBelowMin(msg.value, verificationFee);
        }
        desc.dbNonce = IInterchainDB(INTERCHAIN_DB).getDBNonce();
        InterchainTransaction memory icTx = InterchainTransactionLib.constructLocalTransaction({
            srcSender: msg.sender,
            dstReceiver: receiver,
            dstChainId: dstChainId,
            dbNonce: desc.dbNonce,
            options: options,
            message: message
        });
        desc.transactionId = keccak256(encodeTransaction(icTx));
        // Sanity check: nonce returned from DB should match the nonce used to construct the transaction
        {
            uint64 dbNonce = IInterchainDB(INTERCHAIN_DB).writeEntryRequestVerification{value: verificationFee}(
                icTx.dstChainId, desc.transactionId, srcModules
            );
            assert(dbNonce == desc.dbNonce);
        }
        uint256 executionFee;
        unchecked {
            executionFee = msg.value - verificationFee;
        }
        IExecutionService(srcExecutionService).requestTxExecution{value: executionFee}({
            dstChainId: icTx.dstChainId,
            txPayloadSize: InterchainTransactionLib.payloadSize(options.length, message.length),
            transactionId: desc.transactionId,
            options: options
        });
        emit InterchainTransactionSent({
            transactionId: desc.transactionId,
            dbNonce: desc.dbNonce,
            dstChainId: icTx.dstChainId,
            srcSender: icTx.srcSender,
            dstReceiver: icTx.dstReceiver,
            verificationFee: verificationFee,
            executionFee: executionFee,
            options: icTx.options,
            message: icTx.message
        });
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Asserts that the transaction is executable.
    function _assertExecutable(InterchainTransaction memory icTx, bytes32 transactionId) internal view {
        bytes32 linkedClient = _assertLinkedClient(icTx.srcChainId);
        if (_txExecutor[transactionId] != address(0)) {
            revert InterchainClientV1__TxAlreadyExecuted(transactionId);
        }
        // Construct expected entry based on interchain transaction data
        InterchainEntry memory entry = InterchainEntry({
            srcChainId: icTx.srcChainId,
            dbNonce: icTx.dbNonce,
            entryValue: InterchainEntryLib.getEntryValue({srcWriter: linkedClient, digest: transactionId})
        });
        address receiver = icTx.dstReceiver.bytes32ToAddress();
        (AppConfigV1 memory appConfig, address[] memory approvedModules) = getAppReceivingConfigV1(receiver);
        // Note: appConfig.requiredResponses is never zero at this point, see fallbacks in `getAppReceivingConfigV1`
        // Verify against the Guard if the app opts in to use it
        address guard = _getGuard(appConfig);
        _assertNoGuardConflict(guard, entry);
        // Optimistic period is not used if there's no Guard configured
        uint256 optimisticPeriod = guard == address(0) ? 0 : appConfig.optimisticPeriod;
        uint256 finalizedResponses = _getFinalizedResponsesCount(approvedModules, entry, optimisticPeriod);
        if (finalizedResponses < appConfig.requiredResponses) {
            revert InterchainClientV1__ResponsesAmountBelowMin(finalizedResponses, appConfig.requiredResponses);
        }
    }

    /// @dev Asserts that the chain is linked and returns the linked client address.
    function _assertLinkedClient(uint64 chainId) internal view returns (bytes32 linkedClient) {
        if (chainId == block.chainid) {
            revert InterchainClientV1__ChainIdNotRemote(chainId);
        }
        linkedClient = _linkedClient[chainId];
        if (linkedClient == 0) {
            revert InterchainClientV1__ChainIdNotLinked(chainId);
        }
    }

    /// @dev Asserts that the Guard has not submitted a conflicting entry.
    function _assertNoGuardConflict(address guard, InterchainEntry memory entry) internal view {
        if (guard != address(0)) {
            uint256 confirmedAt = IInterchainDB(INTERCHAIN_DB).checkEntryVerification(guard, entry);
            if (confirmedAt == ENTRY_CONFLICT) {
                revert InterchainClientV1__EntryConflict(guard);
            }
        }
    }

    /// @dev Returns the Guard address to use for the given app config.
    function _getGuard(AppConfigV1 memory appConfig) internal view returns (address) {
        if (appConfig.guardFlag == APP_CONFIG_GUARD_DISABLED) {
            return address(0);
        }
        if (appConfig.guardFlag == APP_CONFIG_GUARD_DEFAULT) {
            return defaultGuard;
        }
        return appConfig.guard;
    }

    /// @dev Counts the number of finalized responses for the given entry.
    /// Note: Reverts if a conflicting entry has been verified by any of the approved modules.
    function _getFinalizedResponsesCount(
        address[] memory approvedModules,
        InterchainEntry memory entry,
        uint256 optimisticPeriod
    )
        internal
        view
        returns (uint256 finalizedResponses)
    {
        for (uint256 i = 0; i < approvedModules.length; ++i) {
            address module = approvedModules[i];
            uint256 confirmedAt = IInterchainDB(INTERCHAIN_DB).checkEntryVerification(module, entry);
            // No-op if the module has not verified anything with the same entry key
            if (confirmedAt == ENTRY_UNVERIFIED) {
                continue;
            }
            // Revert if the module has verified a conflicting entry with the same entry key
            if (confirmedAt == ENTRY_CONFLICT) {
                revert InterchainClientV1__EntryConflict(module);
            }
            // The module has verified this exact entry, check if optimistic period has passed
            if (confirmedAt + optimisticPeriod < block.timestamp) {
                unchecked {
                    ++finalizedResponses;
                }
            }
        }
    }

    /// @dev Asserts that the transaction version is correct and that the transaction is for the current chain.
    /// Note: returns the decoded transaction for chaining purposes.
    function _assertCorrectTransaction(bytes calldata versionedTx)
        internal
        view
        returns (InterchainTransaction memory icTx)
    {
        uint16 version = versionedTx.getVersion();
        if (version != CLIENT_VERSION) {
            revert InterchainClientV1__TxVersionMismatch(version, CLIENT_VERSION);
        }
        icTx = InterchainTransactionLib.decodeTransaction(versionedTx.getPayload());
        if (icTx.dstChainId != block.chainid) {
            revert InterchainClientV1__DstChainIdNotLocal(icTx.dstChainId);
        }
    }

    // solhint-disable no-inline-assembly
    /// @dev Decodes the revert data into a selector and two arguments.
    /// Zero values are returned if the revert data is not long enough.
    /// Note: this is only used in `getTxReadinessV1` to decode the revert data,
    /// so usage of assembly is not a security risk.
    function _decodeRevertData(bytes memory revertData)
        internal
        pure
        returns (bytes4 selector, bytes32 firstArg, bytes32 secondArg)
    {
        // The easiest way to load the bytes chunks onto the stack is to use assembly.
        // Each time we try to load a value, we check if the revert data is long enough.
        // We add 0x20 to skip the length field of the revert data.
        if (revertData.length >= 4) {
            // Load the first 32 bytes, then apply the mask that has only the 4 highest bytes set.
            // There is no need to shift, as `bytesN` variables are right-aligned.
            // https://github.com/ProjectOpenSea/seaport/blob/2ff6ea37/contracts/helpers/SeaportRouter.sol#L161-L175
            selector = bytes4(0xFFFFFFFF);
            assembly {
                selector := and(mload(add(revertData, 0x20)), selector)
            }
        }
        if (revertData.length >= 36) {
            // Skip the length field + selector to get the 32 bytes of the first argument.
            assembly {
                firstArg := mload(add(revertData, 0x24))
            }
        }
        if (revertData.length >= 68) {
            // Skip the length field + selector + first argument to get the 32 bytes of the second argument.
            assembly {
                secondArg := mload(add(revertData, 0x44))
            }
        }
    }
}
