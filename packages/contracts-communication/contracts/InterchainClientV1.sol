pragma solidity 0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";

import {IInterchainApp} from "./interfaces/IInterchainApp.sol";

import {InterchainEntry} from "./libs/InterchainEntry.sol";

import {IInterchainClientV1} from "./interfaces/IInterchainClientV1.sol";

import {TypeCasts} from "./libs/TypeCasts.sol";

import { OptionsLib } from "./libs/Options.sol";

/**
 * @title InterchainClientV1
 * @dev Implements the operations of the Interchain Execution Layer.
 */
contract InterchainClientV1 is Ownable, IInterchainClientV1 {
    using OptionsLib for bytes;

    uint64 public clientNonce;
    address public interchainDB;
    mapping(bytes32 => bool) public executedTransactions;

    // Chain ID => Bytes32 Address of src clients
    mapping(uint256 => bytes32) public linkedClients;

    // TODO: Add permissioning
    // @inheritdoc IInterchainClientV1
    function setLinkedClient(uint256 chainId, bytes32 client) public onlyOwner {
        linkedClients[chainId] = client;
    }

    constructor() Ownable(msg.sender) {}

    // @inheritdoc IInterchainClientV1
    function setInterchainDB(address _interchainDB) public onlyOwner {
        interchainDB = _interchainDB;
    }

    /**
     * @notice Emitted when an interchain transaction is sent.
     */
    event InterchainTransactionSent(
        bytes32 srcSender,
        uint256 srcChainId,
        bytes32 indexed dstReceiver,
        uint256 indexed dstChainId,
        bytes message,
        uint64 nonce,
        bytes options,
        bytes32 indexed transactionId,
        uint256 dbWriterNonce
    );

    // @notice Emitted when an interchain transaction is executed.
    // TODO: Indexing
    event InterchainTransactionExecuted(
        bytes32 indexed srcSender,
        uint256 indexed srcChainId,
        bytes32 dstReceiver,
        uint256 dstChainId,
        bytes message,
        uint64 nonce,
        bytes options,
        bytes32 indexed transactionId,
        uint256 dbWriterNonce
    );

    /**
     * @dev Represents an interchain transaction.
     */
    struct InterchainTransaction {
        bytes32 srcSender;
        uint256 srcChainId;
        bytes32 dstReceiver;
        uint256 dstChainId;
        bytes message;
        uint64 nonce;
        bytes options;
        bytes32 transactionId;
        uint256 dbWriterNonce;
    }

    function _generateTransactionId(
        bytes32 srcSender,
        uint256 srcChainId,
        bytes32 dstReceiver,
        uint256 dstChainId,
        bytes memory message,
        uint64 nonce,
        bytes memory options
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(srcSender, srcChainId, dstReceiver, dstChainId, message, nonce, options));
    }

    // TODO: Calculate Gas Pricing per module and charge fees
    // @inheritdoc IInterchainClientV1
    function interchainSend(
        bytes32 receiver,
        uint256 dstChainId,
        bytes calldata message,
        bytes calldata options,
        address[] calldata srcModules
    )
        public
        payable
    {
        uint256 totalModuleFees = msg.value;

        InterchainTransaction memory icTx = InterchainTransaction({
            srcSender: TypeCasts.addressToBytes32(msg.sender),
            srcChainId: block.chainid,
            dstReceiver: receiver,
            dstChainId: dstChainId,
            message: message,
            nonce: clientNonce,
            options: options,
            transactionId: 0,
            dbWriterNonce: 0
        });

        bytes32 transactionId = _generateTransactionId(icTx.srcSender, icTx.srcChainId, icTx.dstReceiver, icTx.dstChainId, icTx.message, icTx.nonce, icTx.options);
        icTx.transactionId = transactionId;

        uint256 dbWriterNonce = IInterchainDB(interchainDB).writeEntryWithVerification{value: totalModuleFees}(
            icTx.dstChainId, icTx.transactionId, srcModules
        );
        icTx.dbWriterNonce = dbWriterNonce;

        emit InterchainTransactionSent(
            icTx.srcSender, icTx.srcChainId, icTx.dstReceiver, icTx.dstChainId, icTx.message, icTx.nonce, icTx.options, icTx.transactionId, icTx.dbWriterNonce
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
        requiredResponses = IInterchainApp(receiverApp).getRequiredResponses();
        optimisticTimePeriod = IInterchainApp(receiverApp).getOptimisticTimePeriod();
        approvedDstModules = IInterchainApp(receiverApp).getReceivingModules();
    }

    // @inheritdoc IInterchainClientV1
    function isExecutable(bytes calldata transaction) public view returns (bool) {
        InterchainTransaction memory icTx = abi.decode(transaction, (InterchainTransaction));
        require(executedTransactions[icTx.transactionId] == false, "Transaction already executed");
        // Construct expected entry based on icTransaction data
        InterchainEntry memory icEntry = InterchainEntry({
            srcChainId: icTx.srcChainId,
            srcWriter: linkedClients[icTx.srcChainId],
            writerNonce: icTx.dbWriterNonce,
            dataHash: icTx.transactionId
        });

        bytes32 reconstructedID = _generateTransactionId(icTx.srcSender, icTx.srcChainId, icTx.dstReceiver, icTx.dstChainId, icTx.message, icTx.nonce, icTx.options);

        require(icTx.transactionId == reconstructedID, "Invalid transaction ID");

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
            if (approvedResponses[i] + optimisticTimePeriod >= block.timestamp) {
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

    // TODO: Gas Fee Consideration that is paid to executor
    // @inheritdoc IInterchainClientV1
    function interchainExecute(bytes calldata transaction) public {
        require(isExecutable(transaction), "Transaction is not executable");
        InterchainTransaction memory icTx = abi.decode(transaction, (InterchainTransaction));
        executedTransactions[icTx.transactionId] = true;

        OptionsLib.Options memory decodedOptions = icTx.options.decodeOptions();

        IInterchainApp(TypeCasts.bytes32ToAddress(icTx.dstReceiver)).appReceive{gas: decodedOptions.gasLimit}();
        emit InterchainTransactionExecuted(
            icTx.srcSender,
            icTx.srcChainId,
            icTx.dstReceiver,
            icTx.dstChainId,
            icTx.message,
            icTx.nonce,
            icTx.options,
            icTx.transactionId,
            icTx.dbWriterNonce
        );
    }

}
