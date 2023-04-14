// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {BaseMessageLib} from "./libs/BaseMessage.sol";
import {ByteString} from "./libs/ByteString.sol";
import {MAX_CONTENT_BYTES} from "./libs/Constants.sol";
import {HeaderLib, MessageFlag} from "./libs/Message.sol";
import {StateReport} from "./libs/StateReport.sol";
import {State, TypedMemView} from "./libs/State.sol";
import {SystemMessageLib} from "./libs/SystemMessage.sol";
import {Tips, TipsLib} from "./libs/Tips.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {OriginEvents} from "./events/OriginEvents.sol";
import {IAgentManager} from "./interfaces/IAgentManager.sol";
import {InterfaceOrigin} from "./interfaces/InterfaceOrigin.sol";
import {StateHub} from "./hubs/StateHub.sol";
import {AgentStatus, Attestation, Snapshot, StatementHub} from "./hubs/StatementHub.sol";
import {DomainContext, Versioned} from "./system/SystemContract.sol";
import {SystemRegistry} from "./system/SystemRegistry.sol";

contract Origin is StatementHub, StateHub, OriginEvents, InterfaceOrigin {
    using ByteString for bytes;
    using SystemMessageLib for bytes29;
    using TipsLib for bytes;
    using TypeCasts for address;
    using TypedMemView for bytes29;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain, IAgentManager agentManager_)
        DomainContext(domain)
        SystemRegistry(agentManager_)
        Versioned("0.0.3")
    {} // solhint-disable-line no-empty-blocks

    /// @notice Initializes Origin contract:
    /// - msg.sender is set as contract owner
    /// - State of "empty merkle tree" is saved
    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        // Initialize "states": state of an "empty merkle tree" is saved
        _initializeStates();
    }

    // ═════════════════════════════════════════════ VERIFY STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceOrigin
    function verifyAttestation(
        uint256 stateIndex,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(attPayload);
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // This will revert if payload is not a snapshot
        Snapshot snapshot = _wrapSnapshot(snapPayload);
        // This will revert if snapshot/attestation Merkle data doesn't match
        _verifySnapshotMerkle(att, snapshot);
        // This will revert if state index is out of range
        State state = snapshot.state(stateIndex);
        // This will revert if  state refers to another domain
        isValid = _isValidState(state);
        if (!isValid) {
            emit InvalidAttestationState(stateIndex, state.unwrap().clone(), attPayload, attSignature);
            // Slash Notary and notify local AgentManager
            _slashAgent(status.domain, notary);
        }
    }

    /// @inheritdoc InterfaceOrigin
    function verifyAttestationWithProof(
        uint256 stateIndex,
        bytes memory statePayload,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(attPayload);
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // This will revert if payload is not a state
        State state = _wrapState(statePayload);
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, stateIndex, state, snapProof);
        // This will revert, if state refers to another domain
        isValid = _isValidState(state);
        if (!isValid) {
            emit InvalidAttestationState(stateIndex, statePayload, attPayload, attSignature);
            // Slash Notary and notify local AgentManager
            _slashAgent(status.domain, notary);
        }
    }

    /// @inheritdoc InterfaceOrigin
    function verifySnapshot(uint256 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValid)
    {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = _wrapSnapshot(snapPayload);
        // This will revert if the snapshot signer is not a known Agent
        (AgentStatus memory status, address agent) = _verifySnapshot(snapshot, snapSignature);
        // Agent needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // This will revert, if state index is out of range, or state refers to another domain
        isValid = _isValidState(snapshot.state(stateIndex));
        if (!isValid) {
            emit InvalidSnapshotState(stateIndex, snapPayload, snapSignature);
            // Slash Agent and notify local AgentManager
            _slashAgent(status.domain, agent);
        }
    }

    /// @inheritdoc InterfaceOrigin
    function verifyStateReport(bytes memory srPayload, bytes memory srSignature) external returns (bool isValid) {
        // This will revert if payload is not a snapshot report
        StateReport report = _wrapStateReport(srPayload);
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyStateReport(report, srSignature);
        // Guard needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // Report is valid, if the reported state is invalid
        isValid = !_isValidState(report.state());
        if (!isValid) {
            emit InvalidStateReport(srPayload, srSignature);
            // Slash Guard and notify local AgentManager
            _slashAgent(0, guard);
        }
    }

    // ═══════════════════════════════════════════════ SEND MESSAGES ═══════════════════════════════════════════════════

    /// @inheritdoc InterfaceOrigin
    function sendBaseMessage(
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticPeriod,
        bytes memory tipsPayload,
        bytes memory requestPayload,
        bytes memory content
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {
        // Check that content is not too large
        require(content.length <= MAX_CONTENT_BYTES, "content too long");
        // This will revert if payload is not a formatted tips payload
        Tips tips = tipsPayload.castToTips();
        // Tips value must exactly match msg.value
        require(tips.value() == msg.value, "!tips: value");
        // Format the BaseMessage body
        bytes memory body = BaseMessageLib.formatBaseMessage({
            sender_: msg.sender.addressToBytes32(),
            recipient_: recipient,
            tipsPayload: tipsPayload,
            requestPayload: requestPayload,
            content_: content
        });
        // Send the message
        return _sendMessage(destination, optimisticPeriod, MessageFlag.Base, body);
    }

    /// @inheritdoc InterfaceOrigin
    function sendSystemMessage(uint32 destination, uint32 optimisticPeriod, bytes memory body)
        external
        onlySystemRouter
        returns (uint32 messageNonce, bytes32 messageHash)
    {
        // SystemRouter (checked via modifier) is responsible for constructing the body correctly.
        return _sendMessage(destination, optimisticPeriod, MessageFlag.System, body);
    }

    /// @inheritdoc InterfaceOrigin
    function withdrawTips(address recipient, uint256 amount) external onlyAgentManager {
        require(address(this).balance >= amount, "Insufficient balance");
        (bool success,) = recipient.call{value: amount}("");
        require(success, "Recipient reverted");
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Sends the given message to the specified destination. Message hash is inserted
    /// into the Origin Merkle Tree, which will enable message execution on destination chain.
    function _sendMessage(uint32 destination, uint32 optimisticPeriod, MessageFlag flag, bytes memory body)
        internal
        returns (uint32 messageNonce, bytes32 messageHash)
    {
        // Format the message header
        messageNonce = _nextNonce();
        bytes memory headerPayload = HeaderLib.formatHeader({
            origin_: localDomain,
            nonce_: messageNonce,
            destination_: destination,
            optimisticPeriod_: optimisticPeriod
        });
        // Format the full message payload
        bytes memory msgPayload = flag.formatMessage(headerPayload, body);
        // Insert new leaf into the Origin Merkle Tree and save the updated state
        messageHash = keccak256(msgPayload);
        _insertAndSave(messageHash);
        // Emit event with message information
        emit Sent(messageHash, messageNonce, destination, msgPayload);
    }
}
