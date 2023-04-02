// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, ExecutionAttestation} from "../libs/Attestation.sol";
import {BaseMessage, BaseMessageLib} from "../libs/BaseMessage.sol";
import {SYSTEM_ROUTER, ORIGIN_TREE_HEIGHT, SNAPSHOT_TREE_HEIGHT} from "../libs/Constants.sol";
import {MerkleLib} from "../libs/Merkle.sol";
import {Header, Message, MessageFlag, MessageLib} from "../libs/Message.sol";
import {SystemMessage, SystemMessageLib} from "../libs/SystemMessage.sol";
import {Tips} from "../libs/Tips.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";
import {TypedMemView} from "../libs/TypedMemView.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DisputeHub} from "./DisputeHub.sol";
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

    bytes32 internal constant _MESSAGE_STATUS_NONE = bytes32(0);

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @notice (messageHash => status)
    /// TODO: Store something else as "status"? Notary/timestamp?
    /// - Message hasn't been executed: _MESSAGE_STATUS_NONE
    /// - Message has been executed: snapshot root used for proving when executed
    /// @dev Messages coming from different origins will always have a different hash
    /// as origin domain is encoded into the formatted message.
    /// Thus we can use hash as a key instead of an (origin, hash) tuple.
    mapping(bytes32 => bytes32) public messageStatus;

    /// @dev Tracks all saved attestations
    // (root => attestation)
    mapping(bytes32 => ExecutionAttestation) private _rootAttestations;

    /// @dev gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

    // ═════════════════════════════════════════════ EXECUTE MESSAGES ══════════════════════════════════════════════════

    /// @inheritdoc IExecutionHub
    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint256 stateIndex
    ) external {
        // This will revert if payload is not a formatted message payload
        Message message = msgPayload.castToMessage();
        Header header = message.header();
        bytes32 msgLeaf = message.leaf();
        // Ensure message was meant for this domain
        require(header.destination() == localDomain, "!destination");
        uint32 origin = header.origin();
        uint32 nonce = header.nonce();
        // Check proofs validity and optimistically mark message as executed
        ExecutionAttestation memory execAtt =
            _proveAttestation(origin, nonce, msgLeaf, originProof, snapProof, stateIndex);
        // Check if optimistic period has passed
        uint256 rootSubmittedAt = execAtt.submittedAt;
        require(block.timestamp >= rootSubmittedAt + header.optimisticPeriod(), "!optimisticPeriod");
        // Only System/Base message flags exist
        if (message.flag() == MessageFlag.System) {
            _executeSystemMessage(origin, nonce, rootSubmittedAt, message.body());
        } else {
            // This will revert if message body is not a formatted BaseMessage payload
            _executeBaseMessage(origin, nonce, rootSubmittedAt, execAtt.notary, message.body().castToBaseMessage());
        }
        emit Executed(origin, msgLeaf);
    }

    // ═══════════════════════════════════════════ INTERNAL LOGIC: TIPS ════════════════════════════════════════════════

    function _storeTips(address notary, Tips tips) internal {
        // TODO: implement tips logic
        emit TipsStored(notary, tips.unwrap().clone());
    }

    // ═════════════════════════════════════ INTERNAL LOGIC: MESSAGE EXECUTION ═════════════════════════════════════════

    /// @dev Passes message content to recipient that conforms to IMessageRecipient interface.
    function _executeBaseMessage(
        uint32 origin,
        uint32 nonce,
        uint256 rootSubmittedAt,
        address notary,
        BaseMessage baseMessage
    ) internal {
        // Store message tips
        _storeTips(notary, baseMessage.tips());
        // TODO: check that the discarded bits are empty
        address recipient = baseMessage.recipient().bytes32ToAddress();
        // Forward message content to the recipient
        // TODO: this should be "receive base message"
        IMessageRecipient(recipient).handle(
            origin, nonce, baseMessage.sender(), rootSubmittedAt, baseMessage.content().clone()
        );
    }

    function _executeSystemMessage(uint32 origin, uint32 nonce, uint256 rootSubmittedAt, bytes29 body) internal {
        // TODO: introduce incentives for executing System Messages?
        // Forward system message to System Router
        // TODO: this should be a separate function to receive system messages
        IMessageRecipient(address(systemRouter)).handle(origin, nonce, SYSTEM_ROUTER, rootSubmittedAt, body.clone());
    }

    // ══════════════════════════════════════ INTERNAL LOGIC: MESSAGE PROVING ══════════════════════════════════════════

    /**
     * @notice Attempts to prove the validity of the cross-chain message.
     * First, the origin Merkle Root is reconstructed using the origin proof.
     * Then the origin state's "left leaf" is reconstructed using the origin domain.
     * After that the snapshot Merkle Root is reconstructed using the snapshot proof.
     * The snapshot root needs to have been submitted by an undisputed Notary.
     * @dev Reverts if any of the checks fail.
     * @param origin        Domain where message originated
     * @param nonce         Message nonce on the origin domain
     * @param msgLeaf       Message Leaf that was inserted in the Origin Merkle Tree
     * @param originProof   Proof of inclusion of Message Leaf in the Origin Merkle Tree
     * @param snapProof     Proof of inclusion of Origin State Left Leaf into Snapshot Merkle Tree
     * @param stateIndex    Index of Origin State in the Snapshot
     * @return execAtt      Attestation data for derived snapshot root
     */
    function _proveAttestation(
        uint32 origin,
        uint32 nonce,
        bytes32 msgLeaf,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint256 stateIndex
    ) internal returns (ExecutionAttestation memory execAtt) {
        // Check that message has not been executed before
        require(messageStatus[msgLeaf] == _MESSAGE_STATUS_NONE, "!MessageStatus.None");
        // Reconstruct Origin Merkle Root using the origin proof
        // Message index in the tree is (nonce - 1), as nonce starts from 1
        // This will revert if origin proof length exceeds Origin Tree height
        bytes32 originRoot = MerkleLib.proofRoot(nonce - 1, msgLeaf, originProof, ORIGIN_TREE_HEIGHT);
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if:
        //  - State index is out of range.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        bytes32 snapshotRoot = _snapshotRoot(originRoot, origin, snapProof, stateIndex);
        // Fetch the attestation data for the snapshot root
        execAtt = _rootAttestations[snapshotRoot];
        // Check if snapshot root has been submitted
        require(!execAtt.isEmpty(), "Invalid snapshot root");
        // Check if Notary who submitted the attestation is still active
        _verifyActive(_agentStatus(execAtt.notary));
        // Check that Notary who submitted the attestation is not in dispute
        require(!_inDispute(execAtt.notary), "Notary is in dispute");
        // Mark message as executed against the snapshot root
        messageStatus[msgLeaf] = snapshotRoot;
    }

    /// @dev Saves a snapshot root with the attestation data provided by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    function _saveAttestation(Attestation att, address notary) internal {
        bytes32 root = att.snapRoot();
        require(_rootAttestations[root].isEmpty(), "Root already exists");
        _rootAttestations[root] = att.toExecutionAttestation(notary);
    }

    /// @dev Gets a saved attestation for the given snapshot root.
    /// Will return an empty struct, if the snapshot root hasn't been previously saved.
    function _getRootAttestation(bytes32 root) internal view returns (ExecutionAttestation memory) {
        return _rootAttestations[root];
    }
}
