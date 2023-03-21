// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { MAX_MESSAGE_BODY_BYTES, SYSTEM_ROUTER } from "./libs/Constants.sol";
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
import { Attestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { DomainContext, Versioned } from "./system/SystemContract.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Origin is StatementHub, StateHub, OriginEvents, InterfaceOrigin {
    using TipsLib for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain, IAgentManager _agentManager)
        DomainContext(_domain)
        SystemRegistry(_agentManager)
        Versioned("0.0.3")
    {}

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
        uint256 _stateIndex,
        bytes memory _snapPayload,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(_attPayload);
        // This will revert if the attestation signer is not an active Notary
        (uint32 domain, address notary) = _verifyAttestation(att, _attSignature);
        // This will revert if payload is not a snapshot
        Snapshot snapshot = _wrapSnapshot(_snapPayload);
        // This will revert if snapshot/attestation Merkle data doesn't match
        _verifySnapshotMerkle(att, snapshot);
        // This will revert if state index is out of range
        State state = snapshot.state(_stateIndex);
        // This will revert if  state refers to another domain
        isValid = _isValidState(state);
        if (!isValid) {
            emit InvalidAttestationState(
                _stateIndex,
                state.unwrap().clone(),
                _attPayload,
                _attSignature
            );
            // Slash Notary and notify local AgentManager
            _slashAgent(domain, notary);
        }
    }

    /// @inheritdoc InterfaceOrigin
    function verifyAttestationWithProof(
        uint256 _stateIndex,
        bytes memory _statePayload,
        bytes32[] memory _snapProof,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(_attPayload);
        // This will revert if the attestation signer is not an active Notary
        (uint32 domain, address notary) = _verifyAttestation(att, _attSignature);
        // This will revert if payload is not a state
        State state = _wrapState(_statePayload);
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof has length different to Attestation height.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, _stateIndex, state, _snapProof);
        // This will revert, if state refers to another domain
        isValid = _isValidState(state);
        if (!isValid) {
            emit InvalidAttestationState(_stateIndex, _statePayload, _attPayload, _attSignature);
            // Slash Notary and notify local AgentManager
            _slashAgent(domain, notary);
        }
    }

    /// @inheritdoc InterfaceOrigin
    function verifySnapshot(
        uint256 _stateIndex,
        bytes memory _snapPayload,
        bytes memory _snapSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = _wrapSnapshot(_snapPayload);
        // This will revert if the snapshot signer is not an active Agent
        (uint32 domain, address agent) = _verifySnapshot(snapshot, _snapSignature);
        // This will revert, if state index is out of range, or state refers to another domain
        isValid = _isValidState(snapshot.state(_stateIndex));
        if (!isValid) {
            emit InvalidSnapshotState(_stateIndex, _snapPayload, _snapSignature);
            // Slash Agent and notify local AgentManager
            _slashAgent(domain, agent);
        }
    }

    /// @inheritdoc InterfaceOrigin
    function verifyStateReport(bytes memory _srPayload, bytes memory _srSignature)
        external
        returns (bool isValid)
    {
        // This will revert if payload is not a snapshot report
        StateReport report = _wrapStateReport(_srPayload);
        // This will revert if the report signer is not an active Guard
        address guard = _verifyStateReport(report, _srSignature);
        // Report is valid, if the reported state is invalid
        isValid = !_isValidState(report.state());
        if (!isValid) {
            emit InvalidStateReport(_srPayload, _srSignature);
            // Slash Guard and notify local AgentManager
            _slashAgent(0, guard);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          DISPATCH MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceOrigin
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {
        // Modifiers are removed because they prevent from slashing the last active Guard/Notary
        // haveActiveGuard
        // haveActiveNotary(_destination)
        // TODO: figure out a way to filter out unknown domains once Agent Merkle Tree is implemented
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        // This will revert if payload is not a formatted tips payload
        Tips tips = _tips.castToTips();
        // Total tips must exactly match msg.value
        require(tips.totalTips() == msg.value, "!tips: totalTips");
        // Format the message header
        messageNonce = _nextNonce();
        bytes memory header = HeaderLib.formatHeader({
            _origin: localDomain,
            _sender: _checkForSystemRouter(_recipient),
            _nonce: messageNonce,
            _destination: _destination,
            _recipient: _recipient,
            _optimisticSeconds: _optimisticSeconds
        });
        // Format the full message payload
        bytes memory message = MessageLib.formatMessage(header, _tips, _messageBody);

        // Insert new leaf into the Origin Merkle Tree and save the updated state
        messageHash = keccak256(message);
        _insertAndSave(messageHash);

        // Emit Dispatched event with message information
        emit Dispatched(messageHash, messageNonce, _destination, message);
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
    function _checkForSystemRouter(bytes32 _recipient) internal view returns (bytes32 sender) {
        if (_recipient != SYSTEM_ROUTER) {
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
