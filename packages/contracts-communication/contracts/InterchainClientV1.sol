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
    mapping(bytes32 => bool) public executedTransactions;

    // Chain ID => Bytes32 Address of src clients
    mapping(uint256 => bytes32) public linkedClients;

    constructor() Ownable(msg.sender) {}

    // @inheritdoc IInterchainClientV1
    function setExecutionFees(address executionFees_) public onlyOwner {
        executionFees = executionFees_;
    }

    // @inheritdoc IInterchainClientV1
    function setInterchainDB(address _interchainDB) public onlyOwner {
        interchainDB = _interchainDB;
    }

    // @inheritdoc IInterchainClientV1
    function setLinkedClient(uint256 chainId, bytes32 client) public onlyOwner {
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
        public
        payable
    {
        // TODO: should check options for being correctly formatted
        uint256 verificationFee = IInterchainDB(interchainDB).getInterchainFee(dstChainId, srcModules);
        // TODO: should check msg.value >= verificationFee
        uint256 executionFee = msg.value - verificationFee;

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
        if (srcExecutionService != address(0)) {
            IExecutionService(srcExecutionService).requestExecution({
                dstChainId: dstChainId,
                // TODO: there should be a way to calculate the payload size without encoding the transaction
                txPayloadSize: abi.encode(icTx).length,
                transactionId: transactionId,
                executionFee: executionFee,
                options: options
            });
        }
        IExecutionFees(executionFees).addExecutionFee{value: executionFee}(icTx.dstChainId, transactionId);
        emit InterchainTransactionSent(
            transactionId,
            icTx.dbNonce,
            icTx.dstChainId,
            icTx.srcSender,
            icTx.dstReceiver,
            verificationFee,
            executionFee,
            options,
            message
        );
        // Increment nonce for next message
        clientNonce++;
    }

    // TODO: App Config Versioning
    // TODO: What if receiver is not a contract / doesn't conform to interface?
    /**
     * @dev Retrieves the application configuration for a given receiver application.
     * @param receiverApp The address of the receiver application.
     * @return requiredResponses The number of required responses from the receiving modules.
     * @return optimisticTimePeriod The time period within which responses are considered valid.
     * @return approvedDstModules An array of addresses of the approved destination modules.
     */
    function _getAppConfig(address receiverApp)
        internal
        view
        returns (uint256 requiredResponses, uint256 optimisticTimePeriod, address[] memory approvedDstModules)
    {
        bytes memory appConfig;
        (appConfig, approvedDstModules) = IInterchainApp(receiverApp).getReceivingConfig();
        AppConfigV1 memory decodedAppConfig = appConfig.decodeAppConfigV1();
        requiredResponses = decodedAppConfig.requiredResponses;
        optimisticTimePeriod = decodedAppConfig.optimisticPeriod;
    }

    // @inheritdoc IInterchainClientV1
    function isExecutable(bytes calldata encodedTx) public view returns (bool) {
        InterchainTransaction memory icTx = InterchainTransactionLib.decodeTransaction(encodedTx);
        bytes32 transactionId = icTx.transactionId();
        return _isExecutable(transactionId, icTx);
    }

    function _isExecutable(bytes32 transactionId, InterchainTransaction memory icTx) internal view returns (bool) {
        require(executedTransactions[transactionId] == false, "Transaction already executed");
        // Construct expected entry based on icTransaction data
        InterchainEntry memory icEntry = InterchainEntry({
            srcChainId: icTx.srcChainId,
            dbNonce: icTx.dbNonce,
            srcWriter: linkedClients[icTx.srcChainId],
            dataHash: transactionId
        });

        (uint256 requiredResponses, uint256 optimisticTimePeriod, address[] memory approvedDstModules) =
            _getAppConfig(TypeCasts.bytes32ToAddress(icTx.dstReceiver));

        uint256[] memory approvedResponses = _getApprovedResponses(approvedDstModules, icEntry);

        uint256 finalizedResponses = _getFinalizedResponsesCount(approvedResponses, optimisticTimePeriod);
        require(finalizedResponses >= requiredResponses, "Not enough valid responses to meet the threshold");
        return true;
    }

    /**
     * @dev Calculates the number of responses that are considered finalized within the optimistic time period.
     * @param approvedResponses An array of timestamps when each approved response was recorded.
     * @param optimisticTimePeriod The time period in seconds within which a response is considered valid.
     * @return finalizedResponses The count of responses that are finalized within the optimistic time period.
     */
    function _getFinalizedResponsesCount(
        uint256[] memory approvedResponses,
        uint256 optimisticTimePeriod
    )
        internal
        view
        returns (uint256)
    {
        uint256 finalizedResponses = 0;
        for (uint256 i = 0; i < approvedResponses.length; i++) {
            if (approvedResponses[i] + optimisticTimePeriod <= block.timestamp) {
                finalizedResponses++;
            }
        }
        return finalizedResponses;
    }
    /**
     * @dev Retrieves the responses from approved modules for a given InterchainEntry.
     * This function iterates over all approved modules, querying the InterchainDB for each module's response
     * to the provided InterchainEntry. It compiles these responses into an array of uint256, where each
     * element represents the timestamp of a module's response.
     *
     * @param approvedModules An array of addresses representing the approved modules that can write responses.
     * @param icEntry The InterchainEntry for which responses are being retrieved.
     * @return approvedResponses An array of uint256 representing the timestamps of responses from approved modules.
     */

    function _getApprovedResponses(
        address[] memory approvedModules,
        InterchainEntry memory icEntry
    )
        internal
        view
        returns (uint256[] memory)
    {
        uint256[] memory approvedResponses = new uint256[](approvedModules.length);
        for (uint256 i = 0; i < approvedModules.length; i++) {
            approvedResponses[i] = IInterchainDB(interchainDB).readEntry(approvedModules[i], icEntry);
        }
        return approvedResponses;
    }

    function encodeTransaction(InterchainTransaction memory icTx) public view returns (bytes memory) {
        return abi.encode(icTx);
    }

    function encodeOptionsV1(OptionsV1 memory options) public view returns (bytes memory) {
        return options.encodeOptionsV1();
    }

    // TODO: Gas Fee Consideration that is paid to executor
    // @inheritdoc IInterchainClientV1
    function interchainExecute(uint256 gasLimit, bytes calldata transaction) external payable {
        InterchainTransaction memory icTx = InterchainTransactionLib.decodeTransaction(transaction);
        bytes32 transactionId = icTx.transactionId();
        require(_isExecutable(transactionId, icTx), "Transaction is not executable");
        executedTransactions[transactionId] = true;

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
}
