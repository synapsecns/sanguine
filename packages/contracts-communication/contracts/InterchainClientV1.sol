pragma solidity 0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";

import {IInterchainApp} from "./interfaces/IInterchainApp.sol";

import {InterchainEntry} from "./libs/InterchainEntry.sol";

import {IInterchainClientV1} from "./interfaces/IInterchainClientV1.sol";

/**
 * @title InterchainClientV1
 * @dev Implements the operations of the Interchain Execution Layer.
 */
contract InterchainClientV1 is Ownable, IInterchainClientV1 {
    uint64 public clientNonce;
    address public interchainDB;

    // Chain ID => Bytes32 Address of src clients
    mapping(uint256 => bytes32) linkedClients;

    // TODO: Add permissioning
    // @inheritdoc IInterchainClientV1
    function setLinkedClient(uint256 chainId, bytes32 client) public {
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
        bytes32 transactionId;
        uint256 dbWriterNonce;
    }

    // TODO: Calculate Gas Pricing per module and charge fees
    // TODO: Customizable Gas Limit for Execution
    // @inheritdoc IInterchainClientV1
    function interchainSend(
        bytes32 receiver,
        uint256 dstChainId,
        bytes calldata message,
        address[] calldata srcModules
    )
        public
        payable
    {
        uint256 totalModuleFees = msg.value;
        bytes32 sender = convertAddressToBytes32(msg.sender);
        bytes32 transactionID = keccak256(abi.encode(sender, block.chainid, receiver, dstChainId, message, clientNonce));

        uint256 dbWriterNonce = IInterchainDB(interchainDB).writeEntryWithVerification{value: totalModuleFees}(
            dstChainId, transactionID, srcModules
        );

        emit InterchainTransactionSent(
            sender, block.chainid, receiver, dstChainId, message, clientNonce, transactionID, dbWriterNonce
        );
        // Increment nonce for next message
        clientNonce++;
    }

    function isExecutable(bytes calldata transaction) public view returns (bool) {
        InterchainTransaction memory icTx = abi.decode(transaction, (InterchainTransaction));

        // Construct expected entry based on icTransaction data
        InterchainEntry memory icEntry = InterchainEntry({
            srcChainId: icTx.srcChainId,
            srcWriter: linkedClients[icTx.srcChainId],
            writerNonce: icTx.dbWriterNonce,
            dataHash: icTx.transactionId
        });

        bytes memory reconstructedID =
            abi.encode(icTx.srcSender, icTx.srcChainId, icTx.dstReceiver, icTx.dstChainId, icTx.message, icTx.nonce);

        if (icTx.transactionId == keccak256(reconstructedID)) {
            return false;
        }

        address receivingApp = convertBytes32ToAddress(icTx.dstReceiver);

        address[] memory approvedDstModules = IInterchainApp(receivingApp).getReceivingModules();

        uint256 appRequiredResponses = IInterchainApp(receivingApp).getRequiredResponses();

        uint256 optimisticTimePeriod = IInterchainApp(receivingApp).getOptimisticTimePeriod();

    uint256[] memory moduleResponseTimestamps = new uint256[](
        approvedDstModules.length
        );

        for (uint256 i = 0; i < approvedDstModules.length; i++) {
            moduleResponseTimestamps[i] = IInterchainDB(interchainDB).readEntry(approvedDstModules[i], icEntry);
        }
        // 6. Confirm module threshold is met
        uint256 validResponses = 0;

        for (uint256 i = 0; i < moduleResponseTimestamps.length; i++) {
            if (moduleResponseTimestamps[i] + optimisticTimePeriod >= block.timestamp) {
                validResponses++;
            }
        }

        if (validResponses >= appRequiredResponses) {
            return true;
        } else {
            return false;
        }
    }

    // function _getValidResponses(address[] memory approvedModules, InterchainEntry memory icEntry) internal view returns (uint256) {
    //     uint256 validResponses = 0;
    //     for (uint256 i = 0; i < approvedModules.length; i++) {
    //         uint256 moduleResponseTimestamp = IInterchainDB(interchainDB).readEntry(approvedModules[i], icEntry);
    //         if (moduleResponseTimestamp + optimisticTimePeriod >= block.timestamp) {
    //             validResponses++;
    //         }
    //     }
    //     return validResponses;
    // }

    // TODO: Gas Fee Consideration that is paid to executor
    // @inheritdoc IInterchainClientV1
    function interchainExecute(bytes32 transactionID, bytes calldata transaction) public {
        // Steps to verify:
        // 1. Call icDB.getEntry(linkedClients.srcChainId, transaction.dbWriterNonce)
        // 2. Verify the entry hash vs bytes calldata provided
        // 3. Check receiver's app dstModule configuration
        // 4. Check receiver app's optimistic time period
        // 5. Read module entry's based on receiver app dstModule config
        // 6. Confirm module threshold is met
        // 7. Check optimistic threshold set on app config
        // 8. Execute the transaction, is optimistic period is met.

        InterchainTransaction memory icTx = abi.decode(transaction, (InterchainTransaction));

        // 1. Call icDB.getEntry(linkedClients.srcChainId, transaction.dbWriterNonce)
        InterchainEntry memory icEntry = InterchainEntry({
            srcChainId: icTx.srcChainId,
            srcWriter: linkedClients[icTx.srcChainId],
            writerNonce: icTx.dbWriterNonce,
            dataHash: icTx.transactionId
        });

        bytes memory reconstructedID =
            abi.encode(icTx.srcSender, icTx.srcChainId, icTx.dstReceiver, icTx.dstChainId, icTx.message, icTx.nonce);

        // 2. Verify the provided TX ID == constructed TX data
        require(icTx.transactionId == keccak256(reconstructedID), "Invalid transaction ID");

        address receivingApp = convertBytes32ToAddress(icTx.dstReceiver);
        // 3. Check receiver's app dstModule configuration
        address[] memory approvedDstModules = IInterchainApp(receivingApp).getReceivingModules();

        uint256 appRequiredResponses = IInterchainApp(receivingApp).getRequiredResponses();

        // 4. Check receiver app's optimistic time period
        uint256 optimisticTimePeriod = IInterchainApp(receivingApp).getOptimisticTimePeriod();

        // 5. Read module entry's based on receiver app dstModule config
        uint256[] memory moduleResponseTimestamps = new uint256[](
      approvedDstModules.length
    );

        for (uint256 i = 0; i < approvedDstModules.length; i++) {
            moduleResponseTimestamps[i] = IInterchainDB(interchainDB).readEntry(approvedDstModules[i], icEntry);
        }
        // 6. Confirm module threshold is met
        uint256 validResponses = 0;

        for (uint256 i = 0; i < moduleResponseTimestamps.length; i++) {
            if (moduleResponseTimestamps[i] + optimisticTimePeriod >= block.timestamp) {
                validResponses++;
            }
        }

        require(validResponses >= appRequiredResponses, "Not enough valid responses to meet the threshold");

        // 8. Execute the transaction, is optimistic period & valid responses is met.
        IInterchainApp(receivingApp).appReceive();
    }

    // TODO: Seperate out into utils
    /**
     * @inheritdoc IInterchainClientV1
     */
    function convertBytes32ToAddress(bytes32 _bytes32) public pure returns (address) {
        return address(uint160(uint256(_bytes32)));
    }

    /**
     * @inheritdoc IInterchainClientV1
     */
    function convertAddressToBytes32(address _address) public pure returns (bytes32) {
        return bytes32(uint256(uint160(_address)));
    }
}
