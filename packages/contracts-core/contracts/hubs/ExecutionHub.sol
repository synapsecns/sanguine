// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation} from "../libs/Attestation.sol";
import {BaseMessage, BaseMessageLib} from "../libs/BaseMessage.sol";
import {SYSTEM_ROUTER, ORIGIN_TREE_HEIGHT, SNAPSHOT_TREE_HEIGHT} from "../libs/Constants.sol";
import {MerkleLib} from "../libs/Merkle.sol";
import {Header, Message, MessageFlag, MessageLib} from "../libs/Message.sol";
import {Receipt, ReceiptLib} from "../libs/Receipt.sol";
import {MessageStatus} from "../libs/Structures.sol";
import {SystemMessage, SystemMessageLib} from "../libs/SystemMessage.sol";
import {Tips} from "../libs/Tips.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";
import {TypedMemView} from "../libs/TypedMemView.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentStatus, DisputeHub} from "./DisputeHub.sol";
import {ExecutionHubEvents} from "../events/ExecutionHubEvents.sol";
import {IExecutionHub} from "../interfaces/IExecutionHub.sol";
import {IMessageRecipient} from "../interfaces/IMessageRecipient.sol";

/**
 * @notice ExecutionHub is responsible for executing the messages that are
 * proven against the Snapshot Merkle Roots.
 * The Snapshot Merkle Roots themselves are supposed to be dealt with in the child contracts.
 * On the Synapse Chain Notaries are submitting the snapshots that are later used for proving.
 * On the other chains Notaries are submitting the attestations that are later used for proving.
 */
abstract contract ExecutionHub is DisputeHub, ExecutionHubEvents, IExecutionHub {
    using BaseMessageLib for bytes29;
    using MessageLib for bytes;
    using TypeCasts for bytes32;
    using TypedMemView for bytes29;

    /// @notice Struct representing stored data for the snapshot root
    /// @param notaryIndex  Index of Notary who submitted the statement with the snapshot root
    /// @param attNonce     Nonce of the attestation for this snapshot root
    /// @param index        Index of snapshot root in `_roots`
    /// @param submittedAt  Timestamp when the statement with the snapshot root was submitted
    struct SnapRootData {
        uint32 notaryIndex;
        uint32 attNonce;
        uint32 index;
        uint40 submittedAt;
    }
    // 120 bits left for tight packing

    /// @notice Struct representing stored receipt data for the message in Execution Hub.
    /// @param origin       Domain where message originated
    /// @param rootIndex    Index of snapshot root used for proving the message
    /// @param stateIndex   Index of state used for the snapshot proof
    /// @param executor     Executor who successfully executed the message
    struct ReceiptData {
        uint32 origin;
        uint32 rootIndex;
        uint8 stateIndex;
        address executor;
    }
    // TODO: include nonce?
    // 24 bits available for tight packing

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @notice (messageHash => status)
    /// @dev Messages coming from different origins will always have a different hash
    /// as origin domain is encoded into the formatted message.
    /// Thus we can use hash as a key instead of an (origin, hash) tuple.
    mapping(bytes32 => ReceiptData) private _receiptData;

    /// @notice First executor who made a valid attempt of executing a message.
    /// Note: stored only for messages that had Failed status at some point of time
    mapping(bytes32 => address) private _firstExecutor;

    /// @dev All saved snapshot roots
    bytes32[] internal _roots;

    /// @dev Tracks data for all saved snapshot roots
    mapping(bytes32 => SnapRootData) internal _rootData;

    /// @dev gap for upgrade safety
    uint256[46] private __GAP; // solhint-disable-line var-name-mixedcase

    // ═════════════════════════════════════════════ MESSAGE EXECUTION ═════════════════════════════════════════════════

    /// @inheritdoc IExecutionHub
    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint256 stateIndex,
        uint64 gasLimit
    ) external {
        // TODO: add reentrancy check
        // This will revert if payload is not a formatted message payload
        Message message = msgPayload.castToMessage();
        Header header = message.header();
        bytes32 msgLeaf = message.leaf();
        // Ensure message was meant for this domain
        require(header.destination() == localDomain, "!destination");
        // Check that message has not been executed before
        ReceiptData memory rcptData = _receiptData[msgLeaf];
        require(rcptData.executor == address(0), "Already executed");
        // Check proofs validity
        SnapRootData memory rootData = _proveAttestation(header, msgLeaf, originProof, snapProof, stateIndex);
        // Check if optimistic period has passed
        uint256 proofMaturity = block.timestamp - rootData.submittedAt;
        require(proofMaturity >= header.optimisticPeriod(), "!optimisticPeriod");
        bool success;
        // Only System/Base message flags exist
        if (message.flag() == MessageFlag.System) {
            // gasLimit is ignored when executing system messages
            success = _executeSystemMessage(header, proofMaturity, message.body());
        } else {
            // This will revert if message body is not a formatted BaseMessage payload
            BaseMessage baseMessage = message.body().castToBaseMessage();
            success = _executeBaseMessage(header, proofMaturity, gasLimit, baseMessage);
            emit TipsRecorded(msgLeaf, baseMessage.tips().unwrap().clone());
        }
        if (rcptData.origin == 0) {
            // This is the first valid attempt to execute the message => save origin and snapshot proof
            rcptData.origin = header.origin();
            rcptData.rootIndex = rootData.index;
            rcptData.stateIndex = uint8(stateIndex);
            if (success) {
                // This is the successful attempt to execute the message => save the executor
                rcptData.executor = msg.sender;
            } else {
                // Save as the "first executor", if execution failed
                _firstExecutor[msgLeaf] = msg.sender;
            }
            _receiptData[msgLeaf] = rcptData;
        } else if (success) {
            // There has been a failed attempt to execute the message before => don't touch origin and snapshot root
            // This is the successful attempt to execute the message => save the executor
            rcptData.executor = msg.sender;
            _receiptData[msgLeaf] = rcptData;
        }
        emit Executed(header.origin(), msgLeaf);
    }

    /// @inheritdoc IExecutionHub
    function verifyReceipt(bytes memory rcptPayload, bytes memory rcptSignature) external returns (bool isValid) {
        // This will revert if payload is not an receipt
        Receipt rcpt = _wrapReceipt(rcptPayload);
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyReceipt(rcpt, rcptSignature);
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // This will revert if receipt refers to another domain
        isValid = _isValidReceipt(rcpt);
        if (!isValid) {
            emit InvalidReceipt(rcptPayload, rcptSignature);
            // Slash Notary and notify local AgentManager
            _slashAgent(status.domain, notary);
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IExecutionHub
    function isValidReceipt(bytes memory rcptPayload) external view returns (bool isValid) {
        // This will revert if payload is not an receipt
        Receipt rcpt = _wrapReceipt(rcptPayload);
        // This will revert if receipt refers to another domain
        return _isValidReceipt(rcpt);
    }

    /// @inheritdoc IExecutionHub
    function messageStatus(bytes32 messageHash) external view returns (MessageStatus status) {
        ReceiptData memory rcptData = _receiptData[messageHash];
        if (rcptData.executor != address(0)) {
            return MessageStatus.Success;
        } else if (_firstExecutor[messageHash] != address(0)) {
            return MessageStatus.Failed;
        } else {
            return MessageStatus.None;
        }
    }

    /// @inheritdoc IExecutionHub
    function receiptData(bytes32 messageHash) external view returns (bytes memory data) {
        ReceiptData memory rcptData = _receiptData[messageHash];
        // Return empty payload if there has been no attempt to execute the message
        if (rcptData.origin == 0) return "";
        // Determine the first executor who tried to execute the message
        address firstExecutor = _firstExecutor[messageHash];
        if (firstExecutor == address(0)) firstExecutor = rcptData.executor;
        // Determine the snapshot root that was used for proving the message
        bytes32 snapRoot = _roots[rcptData.rootIndex];
        (address attNotary,) = _getAgent(_rootData[snapRoot].notaryIndex);
        // ExecutionHub does not store the tips, the Notary will have to append the tips payload
        return ReceiptLib.formatReceipt({
            origin_: rcptData.origin,
            destination_: localDomain,
            messageHash_: messageHash,
            snapshotRoot_: snapRoot,
            stateIndex_: rcptData.stateIndex,
            attNotary_: attNotary,
            firstExecutor_: firstExecutor,
            finalExecutor_: rcptData.executor,
            tipsPayload: ""
        });
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Passes message content to recipient that conforms to IMessageRecipient interface.
    function _executeBaseMessage(Header header, uint256 proofMaturity, uint64 gasLimit, BaseMessage baseMessage)
        internal
        returns (bool)
    {
        // Check that gas limit covers the one requested by the sender.
        // We let the executor specify gas limit higher than requested to guarantee the execution of
        // messages with gas limit set too low.
        require(gasLimit >= baseMessage.request().gasLimit(), "Gas limit too low");
        // TODO: check that the discarded bits are empty
        address recipient = baseMessage.recipient().bytes32ToAddress();
        // Forward message content to the recipient, and limit the amount of forwarded gas
        require(gasleft() > gasLimit, "Not enough gas supplied");
        try IMessageRecipient(recipient).receiveBaseMessage{gas: gasLimit}(
            header.origin(), header.nonce(), baseMessage.sender(), proofMaturity, baseMessage.content().clone()
        ) {
            return true;
        } catch {
            return false;
        }
    }

    function _executeSystemMessage(Header header, uint256 proofMaturity, bytes29 body) internal returns (bool) {
        // TODO: introduce incentives for executing System Messages?
        // Forward system message to System Router
        systemRouter.receiveSystemMessage(header.origin(), header.nonce(), proofMaturity, body.clone());
        return true;
    }

    /// @dev Saves a snapshot root with the attestation data provided by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    function _saveAttestation(Attestation att, uint32 notaryIndex) internal {
        bytes32 root = att.snapRoot();
        require(_rootData[root].submittedAt == 0, "Root already exists");
        _rootData[root] = SnapRootData(notaryIndex, att.nonce(), uint32(_roots.length), uint40(block.timestamp));
        _roots.push(root);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Checks if receipt matches the saved data for the referenced message.
    /// Reverts if destination domain doesn't match the local domain.
    function _isValidReceipt(Receipt rcpt) internal view returns (bool) {
        // Check if receipt refers to this contract
        require(rcpt.destination() == localDomain, "Wrong destination");
        bytes32 messageHash = rcpt.messageHash();
        ReceiptData memory rcptData = _receiptData[messageHash];
        // Check if there has been a single attempt to execute the message
        if (rcptData.origin == 0) return false;
        // Check that origin and state index fields match
        if (rcpt.origin() != rcptData.origin || rcpt.stateIndex() != rcptData.stateIndex) return false;
        // Check that snapshot root and notary who submitted it match in the Receipt
        bytes32 snapRoot = rcpt.snapshotRoot();
        (address attNotary,) = _getAgent(_rootData[snapRoot].notaryIndex);
        if (snapRoot != _roots[rcptData.rootIndex] || rcpt.attNotary() != attNotary) return false;
        // Check if message was executed from the first attempt
        address firstExecutor = _firstExecutor[messageHash];
        if (firstExecutor == address(0)) {
            // Both first and final executors are saved in receipt data
            return rcpt.firstExecutor() == rcptData.executor && rcpt.finalExecutor() == rcptData.executor;
        } else {
            // Message was Failed at some point of time, so both receipts are valid:
            // "Failed": finalExecutor is ZERO
            // "Success": finalExecutor matches executor from saved receipt data
            address finalExecutor = rcpt.finalExecutor();
            return rcpt.firstExecutor() == firstExecutor
                && (finalExecutor == address(0) || finalExecutor == rcptData.executor);
        }
    }

    /**
     * @notice Attempts to prove the validity of the cross-chain message.
     * First, the origin Merkle Root is reconstructed using the origin proof.
     * Then the origin state's "left leaf" is reconstructed using the origin domain.
     * After that the snapshot Merkle Root is reconstructed using the snapshot proof.
     * The snapshot root needs to have been submitted by an undisputed Notary.
     * @dev Reverts if any of the checks fail.
     * @param header        Memory view over the message header
     * @param msgLeaf       Message Leaf that was inserted in the Origin Merkle Tree
     * @param originProof   Proof of inclusion of Message Leaf in the Origin Merkle Tree
     * @param snapProof     Proof of inclusion of Origin State Left Leaf into Snapshot Merkle Tree
     * @param stateIndex    Index of Origin State in the Snapshot
     * @return rootData     Data for the derived snapshot root
     */
    function _proveAttestation(
        Header header,
        bytes32 msgLeaf,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint256 stateIndex
    ) internal view returns (SnapRootData memory rootData) {
        // Reconstruct Origin Merkle Root using the origin proof
        // Message index in the tree is (nonce - 1), as nonce starts from 1
        // This will revert if origin proof length exceeds Origin Tree height
        bytes32 originRoot = MerkleLib.proofRoot(header.nonce() - 1, msgLeaf, originProof, ORIGIN_TREE_HEIGHT);
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if:
        //  - State index is out of range.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        bytes32 snapshotRoot = _snapshotRoot(originRoot, header.origin(), snapProof, stateIndex);
        // Fetch the attestation data for the snapshot root
        rootData = _rootData[snapshotRoot];
        // Check if snapshot root has been submitted
        require(rootData.submittedAt != 0, "Invalid snapshot root");
        // Check if Notary who submitted the attestation is still active
        (address attNotary, AgentStatus memory attNotaryStatus) = _getAgent(rootData.notaryIndex);
        _verifyActive(attNotaryStatus);
        // Check that Notary who submitted the attestation is not in dispute
        require(!_inDispute(attNotary), "Notary is in dispute");
    }
}
