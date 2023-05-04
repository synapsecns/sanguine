// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/Attestation.sol";
import {Receipt, ReceiptLib} from "../libs/Receipt.sol";
import {Snapshot, SnapshotLib} from "../libs/Snapshot.sol";
import {State, StateLib} from "../libs/State.sol";
import {StateReport, StateReportLib} from "../libs/StateReport.sol";
import {AgentFlag, AgentStatus, Dispute, DisputeFlag} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessagingBase} from "../base/MessagingBase.sol";
import {AgentManagerEvents} from "../events/AgentManagerEvents.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {IExecutionHub} from "../interfaces/IExecutionHub.sol";
import {IStateHub} from "../interfaces/IStateHub.sol";
import {IAgentSecured} from "../interfaces/IAgentSecured.sol";
import {VerificationManager} from "./VerificationManager.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

abstract contract AgentManager is MessagingBase, VerificationManager, AgentManagerEvents, IAgentManager {
    using AttestationLib for bytes;
    using ReceiptLib for bytes;
    using StateLib for bytes;
    using StateReportLib for bytes;
    using SnapshotLib for bytes;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    address public origin;

    address public destination;

    // (agent => their dispute status)
    mapping(address => Dispute) internal _disputes;

    // TODO: optimize this
    bytes[] internal _storedSignatures;

    /// @dev gap for upgrade safety
    uint256[46] private __GAP; // solhint-disable-line var-name-mixedcase

    // ════════════════════════════════════════════════ INITIALIZER ════════════════════════════════════════════════════

    // solhint-disable-next-line func-name-mixedcase
    function __AgentManager_init(address origin_, address destination_) internal onlyInitializing {
        origin = origin_;
        destination = destination_;
    }

    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    // solhint-disable-next-line ordering
    function submitStateReportWithSnapshot(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a state report
        StateReport report = srPayload.castToStateReport();
        // This will revert if the report signer is not an known Guard
        (AgentStatus memory guardStatus, address guard) = _verifyStateReport(report, srSignature);
        // Check that Guard is active
        guardStatus.verifyActive();
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the snapshot signer is not a known Notary
        (AgentStatus memory notaryStatus, address notary) =
            _verifySnapshot({snapshot: snapshot, snapSignature: snapSignature, verifyNotary: true});
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(notaryStatus.domain);
        // Snapshot state and reported state need to be the same
        // This will revert if state index is out of range
        require(snapshot.state(stateIndex).equals(report.state()), "States don't match");
        // This will revert if either actor is already in dispute
        _openDispute(guard, guardStatus.index, notary, notaryStatus.index);
        return true;
    }

    /// @inheritdoc IAgentManager
    function submitStateReportWithAttestation(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a state report
        StateReport report = srPayload.castToStateReport();
        // This will revert if the report signer is not an known Guard
        (AgentStatus memory guardStatus, address guard) = _verifyStateReport(report, srSignature);
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // Snapshot state and reported state need to be the same
        // This will revert if state index is out of range
        require(snapshot.state(stateIndex).equals(report.state()), "States don't match");
        // Check that Guard is active
        guardStatus.verifyActive();
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not an known Notary
        (AgentStatus memory notaryStatus, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(notaryStatus.domain);
        require(snapshot.calculateRoot() == att.snapRoot(), "Attestation not matches snapshot");
        // This will revert if either actor is already in dispute
        _openDispute(guard, guardStatus.index, notary, notaryStatus.index);
        return true;
    }

    /// @inheritdoc IAgentManager
    function submitStateReportWithSnapshotProof(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a state report
        StateReport report = srPayload.castToStateReport();
        // This will revert if the report signer is not an known Guard
        (AgentStatus memory guardStatus, address guard) = _verifyStateReport(report, srSignature);
        // Check that Guard is active
        guardStatus.verifyActive();
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not a known Notary
        (AgentStatus memory notaryStatus, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(notaryStatus.domain);
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, stateIndex, report.state(), snapProof);
        // This will revert if either actor is already in dispute
        _openDispute(guard, guardStatus.index, notary, notaryStatus.index);
        return true;
    }

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    function verifyReceipt(bytes memory rcptPayload, bytes memory rcptSignature)
        external
        returns (bool isValidReceipt)
    {
        // This will revert if payload is not a receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyReceipt(rcpt, rcptSignature);
        // Notary needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        isValidReceipt = IExecutionHub(destination).isValidReceipt(rcptPayload);
        if (!isValidReceipt) {
            emit InvalidReceipt(rcptPayload, rcptSignature);
            _slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc IAgentManager
    function verifyStateWithAttestation(
        uint256 stateIndex,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState) {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        require(snapshot.calculateRoot() == att.snapRoot(), "Attestation not matches snapshot");
        // This will revert if state does not refer to this chain
        bytes memory statePayload = snapshot.state(stateIndex).unwrap().clone();
        isValidState = IStateHub(origin).isValidState(statePayload);
        if (!isValidState) {
            emit InvalidStateWithAttestation(stateIndex, statePayload, attPayload, attSignature);
            _slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc IAgentManager
    function verifyStateWithSnapshotProof(
        uint256 stateIndex,
        bytes memory statePayload,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState) {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // This will revert if payload is not a state
        State state = statePayload.castToState();
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, stateIndex, state, snapProof);
        // This will revert if state does not refer to this chain
        isValidState = IStateHub(origin).isValidState(statePayload);
        if (!isValidState) {
            emit InvalidStateWithAttestation(stateIndex, statePayload, attPayload, attSignature);
            _slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc IAgentManager
    function verifyStateWithSnapshot(uint256 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValidState)
    {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the snapshot signer is not a known Guard/Notary
        (AgentStatus memory status, address agent) =
            _verifySnapshot({snapshot: snapshot, snapSignature: snapSignature, verifyNotary: false});
        // Agent needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // This will revert if state does not refer to this chain
        isValidState = IStateHub(origin).isValidState(snapshot.state(stateIndex).unwrap().clone());
        if (!isValidState) {
            emit InvalidStateWithSnapshot(stateIndex, snapPayload, snapSignature);
            _slashAgent(status.domain, agent, msg.sender);
        }
    }

    /// @inheritdoc IAgentManager
    function verifyStateReport(bytes memory srPayload, bytes memory srSignature)
        external
        returns (bool isValidReport)
    {
        // This will revert if payload is not a snapshot report
        StateReport report = srPayload.castToStateReport();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyStateReport(report, srSignature);
        // Guard needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // Report is valid IF AND ONLY IF the reported state in invalid
        // This will revert if the reported state does not refer to this chain
        isValidReport = !IStateHub(origin).isValidState(report.state().unwrap().clone());
        if (!isValidReport) {
            emit InvalidStateReport(srPayload, srSignature);
            _slashAgent(status.domain, guard, msg.sender);
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status) {
        agent = _getAgent(index);
        if (agent != address(0)) status = agentStatus(agent);
    }

    /// @inheritdoc IAgentManager
    function agentStatus(address agent) public view returns (AgentStatus memory status) {
        status = _storedAgentStatus(agent);
        // If agent was proven to commit fraud, but their slashing wasn't completed,
        // return the Fraudulent flag instead
        if (_disputes[agent].flag == DisputeFlag.Slashed && status.flag != AgentFlag.Slashed) {
            status.flag = AgentFlag.Fraudulent;
        }
    }

    /// @inheritdoc IAgentManager
    function disputeStatus(address agent) external view returns (Dispute memory) {
        return _disputes[agent];
    }

    /// @inheritdoc IAgentManager
    function getStoredSignature(uint256 index) external view returns (bytes memory) {
        return _storedSignatures[index];
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Hook that is called after agent was slashed in AgentManager and AgentSecured contracts were notified.
    // solhint-disable-next-line no-empty-blocks
    function _afterAgentSlashed(uint32 domain, address agent, address prover) internal virtual {}

    /// @dev Child contract should implement the logic for notifying AgentSecured contracts about the opened dispute.
    function _notifyDisputeOpened(uint32 guardIndex, uint32 notaryIndex) internal virtual;

    /// @dev Child contract should implement the logic for notifying AgentSecured contracts about the resolved dispute.
    function _notifyDisputeResolved(uint32 slashedIndex, uint32 rivalIndex) internal virtual;

    /// @dev Opens a Dispute between a Guard and a Notary, if they are both not in Dispute already.
    function _openDispute(address guard, uint32 guardIndex, address notary, uint32 notaryIndex) internal {
        // Check that both agents are not in Dispute yet
        require(_disputes[guard].flag == DisputeFlag.None, "Guard already in dispute");
        require(_disputes[notary].flag == DisputeFlag.None, "Notary already in dispute");
        _updateDispute(guard, Dispute(DisputeFlag.Pending, notaryIndex, address(0)));
        _updateDispute(notary, Dispute(DisputeFlag.Pending, guardIndex, address(0)));
        _notifyDisputeOpened(guardIndex, notaryIndex);
    }

    /// @dev Slashes the Agent and notifies the local Destination and Origin contracts about the slashed agent.
    /// Should be called when the agent fraud was confirmed.
    function _slashAgent(uint32 domain, address agent, address prover) internal {
        // Check that agent is Active/Unstaking and that the domains match
        AgentStatus memory status = _storedAgentStatus(agent);
        require(
            (status.flag == AgentFlag.Active || status.flag == AgentFlag.Unstaking) && status.domain == domain,
            "Slashing could not be initiated"
        );
        // The "stored" agent status is not updated yet, however agentStatus() will return AgentFlag.Fraudulent
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
        // This will revert if the agent has been slashed earlier
        _resolveDispute(agent, status.index, prover);
        // Call "after slash" hook - this allows Bonding/Light Manager to add custom "after slash" logic
        _afterAgentSlashed(domain, agent, prover);
    }

    /// @dev Resolves a Dispute between a slashed Agent and their Rival (if there was one).
    function _resolveDispute(address slashedAgent, uint32 slashedIndex, address prover) internal {
        Dispute memory dispute = _disputes[slashedAgent];
        require(dispute.flag != DisputeFlag.Slashed, "Dispute already resolved");
        (dispute.flag, dispute.fraudProver) = (DisputeFlag.Slashed, prover);
        _updateDispute(slashedAgent, dispute);
        // Clear Dispute status for the Rival
        if (dispute.rivalIndex != 0) {
            _updateDispute(_getAgent(dispute.rivalIndex), Dispute(DisputeFlag.None, 0, address(0)));
        }
        _notifyDisputeResolved(slashedIndex, dispute.rivalIndex);
    }

    /// @dev Updates a dispute status for the agent and emits an event.
    function _updateDispute(address agent, Dispute memory dispute) internal {
        _disputes[agent] = dispute;
        emit DisputeUpdated(agent, dispute);
    }

    /// @dev Saves the signature and returns its index.
    function _saveSignature(bytes memory signature) internal returns (uint256 sigIndex) {
        sigIndex = _storedSignatures.length;
        _storedSignatures.push(signature);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Generates leaf to be saved in the Agent Merkle Tree
    function _agentLeaf(AgentFlag flag, uint32 domain, address agent) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(flag, domain, agent));
    }

    /// @dev Returns the last known status for the agent from the Agent Merkle Tree.
    /// Note: the actual agent status (returned by `agentStatus()`) may differ, if agent fraud was proven.
    function _storedAgentStatus(address agent) internal view virtual returns (AgentStatus memory);

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes.
    function _getAgent(uint256 index) internal view virtual returns (address);

    /// @inheritdoc VerificationManager
    function _recoverAgent(bytes32 hashedStatement, bytes memory signature)
        internal
        view
        override
        returns (AgentStatus memory status, address agent)
    {
        bytes32 ethSignedMsg = ECDSA.toEthSignedMessageHash(hashedStatement);
        agent = ECDSA.recover(ethSignedMsg, signature);
        status = agentStatus(agent);
        // Discard signature of unknown agents.
        // Further flag checks are supposed to be performed in a caller function.
        require(status.flag != AgentFlag.Unknown, "Unknown agent");
    }

    /// @dev Verifies that Notary signature is active on local domain
    function _verifyNotaryDomain(uint32 notaryDomain) internal view virtual;
}
