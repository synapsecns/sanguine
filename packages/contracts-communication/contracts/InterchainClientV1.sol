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

    uint64 public clientNonce;
    address public interchainDB;
    address public executionFees;

    // Chain ID => Bytes32 Address of src clients
    mapping(uint256 => bytes32) public linkedClients;

    mapping(bytes32 transactionId => address executor) internal _txExecutor;

    constructor() Ownable(msg.sender) {}

    // @inheritdoc IInterchainClientV1
    function setExecutionFees(address executionFees_) external onlyOwner {
        executionFees = executionFees_;
    }

    // @inheritdoc IInterchainClientV1
    function setInterchainDB(address _interchainDB) external onlyOwner {
        interchainDB = _interchainDB;
    }

    // @inheritdoc IInterchainClientV1
    function setLinkedClient(uint256 chainId, bytes32 client) external onlyOwner {
        linkedClients[chainId] = client;
    }

    // TODO: Calculate Gas Pricing per module and charge fees
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
    {
        // TODO: should check options for being correctly formatted
        uint256 verificationFee = IInterchainDB(interchainDB).getInterchainFee(dstChainId, srcModules);
        if (msg.value < verificationFee) {
            revert InterchainClientV1__FeeAmountTooLow(msg.value, verificationFee);
        }
        uint256 executionFee;
        unchecked {
            executionFee = msg.value - verificationFee;
        }
        InterchainTransaction memory icTx = InterchainTransactionLib.constructLocalTransaction({
            srcSender: msg.sender,
            dstReceiver: receiver,
            dstChainId: dstChainId,
            nonce: clientNonce,
            dbNonce: IInterchainDB(interchainDB).getDBNonce(),
            options: options,
            message: message
        });

        bytes32 transactionId = icTx.transactionId();
        // Sanity check: nonce returned from DB should match the nonce used to construct the transaction
        assert(
            icTx.dbNonce
                == IInterchainDB(interchainDB).writeEntryWithVerification{value: verificationFee}(
                    icTx.dstChainId, transactionId, srcModules
                )
        );
        IExecutionFees(executionFees).addExecutionFee{value: executionFee}(icTx.dstChainId, transactionId);
        // TODO: Should this be moved into a seperate configurable contract, for easier upgradeability later?
        if (srcExecutionService != address(0)) {
            IExecutionService(srcExecutionService).requestExecution({
                dstChainId: dstChainId,
                // TODO: there should be a way to calculate the payload size without encoding the transaction
                txPayloadSize: abi.encode(icTx).length,
                transactionId: transactionId,
                executionFee: executionFee,
                options: options
            });
            address srcExecutorEOA = IExecutionService(srcExecutionService).executorEOA();
            IExecutionFees(executionFees).recordExecutor(dstChainId, transactionId, srcExecutorEOA);
        }
        emit InterchainTransactionSent(
            transactionId,
            icTx.dbNonce,
            icTx.nonce,
            icTx.dstChainId,
            icTx.srcSender,
            icTx.dstReceiver,
            verificationFee,
            executionFee,
            icTx.options,
            icTx.message
        );
        // Increment nonce for next message
        clientNonce++;
    }

    // TODO: Gas Fee Consideration that is paid to executor
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
            nonce: icTx.nonce,
            message: icTx.message
        });
        emit InterchainTransactionReceived(
            transactionId, icTx.dbNonce, icTx.srcChainId, icTx.srcSender, icTx.dstReceiver
        );
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

    function encodeTransaction(InterchainTransaction memory icTx) external pure returns (bytes memory) {
        return abi.encode(icTx);
    }

    function decodeOptions(bytes memory encodedOptions) external pure returns (OptionsV1 memory) {
        return encodedOptions.decodeOptionsV1();
    }

    /// @dev Asserts that the transaction is executable. Returns the transactionId for chaining purposes.
    function _assertExecutable(InterchainTransaction memory icTx) internal view returns (bytes32 transactionId) {
        transactionId = icTx.transactionId();
        if (_txExecutor[transactionId] != address(0)) {
            revert InterchainClientV1__AlreadyExecuted(transactionId);
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
        if (responses < appConfig.requiredResponses) {
            revert InterchainClientV1__NotEnoughResponses(responses, appConfig.requiredResponses);
        }
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
            uint256 confirmedAt = IInterchainDB(interchainDB).readEntry(approvedModules[i], icEntry);
            // readEntry() returns 0 if entry hasn't been confirmed by the module, so we check for that as well
            if (confirmedAt != 0 && confirmedAt + optimisticPeriod <= block.timestamp) {
                ++finalizedResponses;
            }
        }
    }
}
