// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { MAX_CONTENT_BYTES, SYSTEM_ROUTER } from "./libs/Constants.sol";
import { HeaderLib, MessageLib } from "./libs/Message.sol";
import { StateReport } from "./libs/StateReport.sol";
import { State, StateLib, TypedMemView } from "./libs/State.sol";
import { Tips, TipsLib } from "./libs/Tips.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { OriginEvents } from "./events/OriginEvents.sol";
import { IAgentManager } from "./interfaces/IAgentManager.sol";
import { InterfaceOrigin } from "./interfaces/InterfaceOrigin.sol";
import { StateHub } from "./hubs/StateHub.sol";
import { AgentStatus, Attestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { DomainContext, Versioned } from "./system/SystemContract.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Origin is StatementHub, StateHub, OriginEvents, InterfaceOrigin {
    using TipsLib for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VERIFY STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
            emit InvalidAttestationState(
                stateIndex,
                state.unwrap().clone(),
                attPayload,
                attSignature
            );
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
    function verifySnapshot(
        uint256 stateIndex,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bool isValid) {
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
    function verifyStateReport(bytes memory srPayload, bytes memory srSignature)
        external
        returns (bool isValid)
    {
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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          DISPATCH MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceOrigin
    function dispatch(
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds,
        bytes memory tipsPayload,
        bytes memory content
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {
        // Modifiers are removed because they prevent from slashing the last active Guard/Notary
        // haveActiveGuard
        // haveActiveNotary(destination)
        // TODO: figure out a way to filter out unknown domains once Agent Merkle Tree is implemented
        require(content.length <= MAX_CONTENT_BYTES, "content too long");
        // This will revert if payload is not a formatted tips payload
        Tips tips = tipsPayload.castToTips();
        // Total tips must exactly match msg.value
        require(tips.totalTips() == msg.value, "!tips: totalTips");
        // Format the message header
        messageNonce = _nextNonce();
        bytes memory headerPayload = HeaderLib.formatHeader({
            origin_: localDomain,
            sender_: _checkForSystemRouter(recipient),
            nonce_: messageNonce,
            destination_: destination,
            recipient_: recipient,
            optimisticSeconds_: optimisticSeconds
        });
        // Format the full message payload
        bytes memory msgPayload = MessageLib.formatMessage(headerPayload, tipsPayload, content);

        // Insert new leaf into the Origin Merkle Tree and save the updated state
        messageHash = keccak256(msgPayload);
        _insertAndSave(messageHash);

        // Emit Dispatched event with message information
        emit Dispatched(messageHash, messageNonce, destination, msgPayload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns adjusted "sender" field.
     * @dev By default, "sender" field is msg.sender address casted to bytes32.
     * However, if SYSTEM_ROUTER is used for "recipient" field, and msg.sender is SystemRouter,
     * SYSTEM_ROUTER is also used as "sender" field.
     * Note: tx will revert if anyone but SystemRouter uses SYSTEM_ROUTER as the recipient.
     */
    function _checkForSystemRouter(bytes32 recipient) internal view returns (bytes32 sender) {
        if (recipient != SYSTEM_ROUTER) {
            sender = TypeCasts.addressToBytes32(msg.sender);
            /**
             * @dev Note: SYSTEM_ROUTER has only the highest 12 bytes set,
             * whereas TypeCasts.addressToBytes32 sets only the lowest 20 bytes.
             * Thus, in this branch: sender != SYSTEM_ROUTER
             */
        } else {
            // Check that SystemRouter specified SYSTEM_ROUTER as recipient, revert otherwise.
            _assertSystemRouter();
            // Adjust "sender" field for correct processing on remote chain.
            sender = SYSTEM_ROUTER;
        }
    }
}
