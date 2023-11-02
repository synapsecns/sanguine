// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation} from "../libs/memory/Attestation.sol";
import {BaseMessage, BaseMessageLib, MemView} from "../libs/memory/BaseMessage.sol";
import {ByteString, CallData} from "../libs/memory/ByteString.sol";
import {ORIGIN_TREE_HEIGHT, SNAPSHOT_TREE_HEIGHT} from "../libs/Constants.sol";
import {
    AlreadyExecuted,
    AlreadyFailed,
    DisputeTimeoutNotOver,
    DuplicatedSnapshotRoot,
    IncorrectDestinationDomain,
    IncorrectMagicValue,
    IncorrectOriginDomain,
    IncorrectSnapshotRoot,
    GasLimitTooLow,
    GasSuppliedTooLow,
    MessageOptimisticPeriod,
    NotaryInDispute
} from "../libs/Errors.sol";
import {SafeCall} from "../libs/SafeCall.sol";
import {MerkleMath} from "../libs/merkle/MerkleMath.sol";
import {Header, Message, MessageFlag, MessageLib} from "../libs/memory/Message.sol";
import {Receipt, ReceiptLib} from "../libs/memory/Receipt.sol";
import {Request} from "../libs/stack/Request.sol";
import {SnapshotLib} from "../libs/memory/Snapshot.sol";
import {AgentFlag, AgentStatus, MessageStatus} from "../libs/Structures.sol";
import {Tips} from "../libs/stack/Tips.sol";
import {ChainContext} from "../libs/ChainContext.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "../base/AgentSecured.sol";
import {ExecutionHubEvents} from "../events/ExecutionHubEvents.sol";
import {InterfaceInbox} from "../interfaces/InterfaceInbox.sol";
import {IExecutionHub} from "../interfaces/IExecutionHub.sol";
import {IMessageRecipient} from "../interfaces/IMessageRecipient.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

/// @notice `ExecutionHub` is a parent contract for `Destination`. It is responsible for the following:
/// - Executing the messages that are proven against the saved Snapshot Merkle Roots.
/// - Base messages are forwarded to the specified message recipient, ensuring that the original
///   execution request is fulfilled correctly.
/// - Manager messages are forwarded to the local `AgentManager` contract.
/// - Keeping track of the saved Snapshot Merkle Roots (which are accepted in `Destination`).
/// - Keeping track of message execution Receipts, as well as verify their validity.
abstract contract ExecutionHub is AgentSecured, ReentrancyGuardUpgradeable, ExecutionHubEvents, IExecutionHub {
    using Address for address;
    using BaseMessageLib for MemView;
    using ByteString for MemView;
    using MessageLib for bytes;
    using ReceiptLib for bytes;
    using SafeCall for address;
    using SafeCast for uint256;
    using TypeCasts for bytes32;

    /// @notice Struct representing stored data for the snapshot root
    /// @param notaryIndex  Index of Notary who submitted the statement with the snapshot root
    /// @param attNonce     Nonce of the attestation for this snapshot root
    /// @param attBN        Summit block number of the attestation for this snapshot root
    /// @param attTS        Summit timestamp of the attestation for this snapshot root
    /// @param index        Index of snapshot root in `_roots`
    /// @param submittedAt  Timestamp when the statement with the snapshot root was submitted
    /// @param notaryV      V-value from the Notary signature for the attestation
    // TODO: tight pack this
    struct SnapRootData {
        uint32 notaryIndex;
        uint32 attNonce;
        uint40 attBN;
        uint40 attTS;
        uint32 index;
        uint40 submittedAt;
        uint256 sigIndex;
    }

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
        uint8 stateIndex,
        uint64 gasLimit
    ) external nonReentrant {
        // This will revert if payload is not a formatted message payload
        Message message = msgPayload.castToMessage();
        Header header = message.header();
        bytes32 msgLeaf = message.leaf();
        // Ensure message was meant for this domain
        if (header.destination() != localDomain) revert IncorrectDestinationDomain();
        // Ensure message was not sent from this domain
        if (header.origin() == localDomain) revert IncorrectOriginDomain();
        // Check that message has not been executed before
        ReceiptData memory rcptData = _receiptData[msgLeaf];
        if (rcptData.executor != address(0)) revert AlreadyExecuted();
        // Check proofs validity
        SnapRootData memory rootData = _proveAttestation(header, msgLeaf, originProof, snapProof, stateIndex);
        // Check if optimistic period has passed
        uint256 proofMaturity = block.timestamp - rootData.submittedAt;
        if (proofMaturity < header.optimisticPeriod()) revert MessageOptimisticPeriod();
        uint256 paddedTips;
        bool success;
        // Only Base/Manager message flags exist
        if (header.flag() == MessageFlag.Base) {
            // This will revert if message body is not a formatted BaseMessage payload
            BaseMessage baseMessage = message.body().castToBaseMessage();
            success = _executeBaseMessage(header, proofMaturity, gasLimit, baseMessage);
            paddedTips = Tips.unwrap(baseMessage.tips());
        } else {
            // gasLimit is ignored when executing manager messages
            success = _executeManagerMessage(header, proofMaturity, message.body());
        }
        if (rcptData.origin == 0) {
            // This is the first valid attempt to execute the message => save origin and snapshot proof
            rcptData.origin = header.origin();
            rcptData.rootIndex = rootData.index;
            rcptData.stateIndex = stateIndex;
            if (success) {
                // This is the successful attempt to execute the message => save the executor
                rcptData.executor = msg.sender;
            } else {
                // Save as the "first executor", if execution failed
                _firstExecutor[msgLeaf] = msg.sender;
            }
            _receiptData[msgLeaf] = rcptData;
        } else {
            if (!success) revert AlreadyFailed();
            // There has been a failed attempt to execute the message before => don't touch origin and snapshot root
            // This is the successful attempt to execute the message => save the executor
            rcptData.executor = msg.sender;
            _receiptData[msgLeaf] = rcptData;
        }
        emit Executed(header.origin(), msgLeaf, success);
        if (!_passReceipt(rootData.notaryIndex, rootData.attNonce, msgLeaf, paddedTips, rcptData)) {
            // Emit event with the recorded tips so that Notaries could form a receipt to submit to Summit
            emit TipsRecorded(msgLeaf, paddedTips);
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IExecutionHub
    function getAttestationNonce(bytes32 snapRoot) external view returns (uint32 attNonce) {
        return _rootData[snapRoot].attNonce;
    }

    /// @inheritdoc IExecutionHub
    function isValidReceipt(bytes memory rcptPayload) external view returns (bool isValid) {
        // This will revert if payload is not a receipt
        // This will revert if receipt refers to another domain
        return _isValidReceipt(rcptPayload.castToReceipt());
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
    function messageReceipt(bytes32 messageHash) external view returns (bytes memory rcptPayload) {
        ReceiptData memory rcptData = _receiptData[messageHash];
        // Return empty payload if there has been no attempt to execute the message
        if (rcptData.origin == 0) return "";
        return _messageReceipt(messageHash, rcptData);
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
        Request request = baseMessage.request();
        if (gasLimit < request.gasLimit()) revert GasLimitTooLow();
        // TODO: check that the discarded bits are empty
        address recipient = baseMessage.recipient().bytes32ToAddress();
        // Forward message content to the recipient, and limit the amount of forwarded gas
        if (gasleft() <= gasLimit) revert GasSuppliedTooLow();
        // receiveBaseMessage(origin, nonce, sender, proofMaturity, version, content)
        bytes memory payload = abi.encodeCall(
            IMessageRecipient.receiveBaseMessage,
            (
                header.origin(),
                header.nonce(),
                baseMessage.sender(),
                proofMaturity,
                request.version(),
                baseMessage.content().clone()
            )
        );
        // Pass the base message to the recipient, return the success status of the call
        return recipient.safeCall({gasLimit: gasLimit, msgValue: 0, payload: payload});
    }

    /// @dev Uses message body for a call to AgentManager, and checks the returned magic value to ensure that
    /// only "remoteX" functions could be called this way.
    function _executeManagerMessage(Header header, uint256 proofMaturity, MemView body) internal returns (bool) {
        // TODO: introduce incentives for executing Manager Messages?
        CallData callData = body.castToCallData();
        // Add the (origin, proofMaturity) values to the calldata
        bytes memory payload = callData.addPrefix(abi.encode(header.origin(), proofMaturity));
        // functionCall() calls AgentManager and bubbles the revert from the external call
        bytes memory magicValue = address(agentManager).functionCall(payload);
        // We check the returned value here to ensure that only "remoteX" functions could be called this way.
        // This is done to prevent an attack by a malicious Notary trying to force Destination to call an arbitrary
        // function in a local AgentManager. Any other function will not return the required selector,
        // while the "remoteX" functions will perform the proofMaturity check that will make impossible to
        // submit an attestation and execute a malicious Manager Message immediately, preventing this attack vector.
        if (magicValue.length != 32 || bytes32(magicValue) != callData.callSelector()) revert IncorrectMagicValue();
        return true;
    }

    /// @dev Passes the message receipt to the Inbox contract, if it is deployed on Synapse Chain.
    /// This ensures that the message receipts for the messages executed on Synapse Chain are passed to Summit
    /// without a Notary having to sign them.
    function _passReceipt(
        uint32 attNotaryIndex,
        uint32 attNonce,
        bytes32 messageHash,
        uint256 paddedTips,
        ReceiptData memory rcptData
    ) internal returns (bool) {
        // Do nothing if contract is not deployed on Synapse Chain
        if (localDomain != synapseDomain) return false;
        // Do nothing for messages with no tips (TODO: introduce incentives for manager messages?)
        if (paddedTips == 0) return false;
        return InterfaceInbox(inbox).passReceipt({
            attNotaryIndex: attNotaryIndex,
            attNonce: attNonce,
            paddedTips: paddedTips,
            rcptPayload: _messageReceipt(messageHash, rcptData)
        });
    }

    /// @dev Saves a snapshot root with the attestation data provided by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    function _saveAttestation(Attestation att, uint32 notaryIndex, uint256 sigIndex) internal {
        bytes32 root = att.snapRoot();
        if (_rootData[root].submittedAt != 0) revert DuplicatedSnapshotRoot();
        // TODO: consider using more than 32 bits for the root index
        _rootData[root] = SnapRootData({
            notaryIndex: notaryIndex,
            attNonce: att.nonce(),
            attBN: att.blockNumber(),
            attTS: att.timestamp(),
            index: _roots.length.toUint32(),
            submittedAt: ChainContext.blockTimestamp(),
            sigIndex: sigIndex
        });
        _roots.push(root);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Checks if receipt body matches the saved data for the referenced message.
    /// Reverts if destination domain doesn't match the local domain.
    function _isValidReceipt(Receipt rcpt) internal view returns (bool) {
        // Check if receipt refers to this chain
        if (rcpt.destination() != localDomain) revert IncorrectDestinationDomain();
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
        uint8 stateIndex
    ) internal view returns (SnapRootData memory rootData) {
        // Reconstruct Origin Merkle Root using the origin proof
        // Message index in the tree is (nonce - 1), as nonce starts from 1
        // This will revert if origin proof length exceeds Origin Tree height
        bytes32 originRoot = MerkleMath.proofRoot(header.nonce() - 1, msgLeaf, originProof, ORIGIN_TREE_HEIGHT);
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if:
        //  - State index is out of range.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        bytes32 snapshotRoot = SnapshotLib.proofSnapRoot(originRoot, header.origin(), snapProof, stateIndex);
        // Fetch the attestation data for the snapshot root
        rootData = _rootData[snapshotRoot];
        // Check if snapshot root has been submitted
        if (rootData.submittedAt == 0) revert IncorrectSnapshotRoot();
        // Check that Notary who submitted the attestation is not in dispute
        if (_notaryDisputeExists(rootData.notaryIndex)) revert NotaryInDispute();
        // Check that Notary who submitted the attestation isn't in post-dispute timeout
        if (_notaryDisputeTimeout(rootData.notaryIndex)) revert DisputeTimeoutNotOver();
    }

    /// @dev Formats the message execution receipt payload for the given hash and receipt data.
    function _messageReceipt(bytes32 messageHash, ReceiptData memory rcptData)
        internal
        view
        returns (bytes memory rcptPayload)
    {
        // Determine the first executor who tried to execute the message
        address firstExecutor = _firstExecutor[messageHash];
        if (firstExecutor == address(0)) firstExecutor = rcptData.executor;
        // Determine the snapshot root that was used for proving the message
        bytes32 snapRoot = _roots[rcptData.rootIndex];
        (address attNotary,) = _getAgent(_rootData[snapRoot].notaryIndex);
        // ExecutionHub does not store the tips,
        // the Notary will have to derive the proof of tips from the message payload.
        return ReceiptLib.formatReceipt({
            origin_: rcptData.origin,
            destination_: localDomain,
            messageHash_: messageHash,
            snapshotRoot_: snapRoot,
            stateIndex_: rcptData.stateIndex,
            attNotary_: attNotary,
            firstExecutor_: firstExecutor,
            finalExecutor_: rcptData.executor
        });
    }
}
