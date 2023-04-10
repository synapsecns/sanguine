// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {AgentStatus} from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager} from "./manager/AgentManager.sol";
import {DomainContext} from "./context/DomainContext.sol";
import {SummitEvents} from "./events/SummitEvents.sol";
import {IAgentManager} from "./interfaces/IAgentManager.sol";
import {InterfaceSummit} from "./interfaces/InterfaceSummit.sol";
import {DisputeHub, ExecutionHub, Receipt} from "./hubs/ExecutionHub.sol";
import {SnapshotHub, SummitAttestation, SummitState} from "./hubs/SnapshotHub.sol";
import {Attestation, AttestationLib, AttestationReport, Snapshot} from "./hubs/StatementHub.sol";
import {DomainContext, Versioned} from "./system/SystemContract.sol";
import {SystemRegistry} from "./system/SystemRegistry.sol";

contract Summit is ExecutionHub, SnapshotHub, SummitEvents, InterfaceSummit {
    using AttestationLib for bytes;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain, IAgentManager agentManager_)
        DomainContext(domain)
        SystemRegistry(agentManager_)
        Versioned("0.0.3")
    {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        _initializeAttestations();
    }

    // ═════════════════════════════════════════════ ACCEPT STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceSummit
    function submitReceipt(bytes memory rcptPayload, bytes memory rcptSignature) external returns (bool wasAccepted) {
        // Call the hook and check if we can accept the statement
        if (!_beforeStatement()) return false;
        // This will revert if payload is not an receipt
        Receipt rcpt = _wrapReceipt(rcptPayload);
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyReceipt(rcpt, rcptSignature);
        // Notary needs to be Active and not in Dispute
        _verifyActive(status);
        require(!_inDispute(notary), "Notary is in dispute");
        // Receipt needs to be signed by a destination chain Notary
        require(rcpt.destination() == status.domain, "Wrong Notary domain");
        _saveReceipt(rcpt);
        emit ReceiptAccepted(status.domain, notary, rcptPayload, rcptSignature);
    }

    /// @inheritdoc InterfaceSummit
    function submitSnapshot(bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bytes memory attPayload)
    {
        // Call the hook and check if we can accept the statement
        if (!_beforeStatement()) return "";
        // This will revert if payload is not a snapshot
        Snapshot snapshot = _wrapSnapshot(snapPayload);
        // This will revert if the signer is not a known Agent
        (AgentStatus memory status, address agent) = _verifySnapshot(snapshot, snapSignature);
        // Check that Agent is active
        _verifyActive(status);
        if (status.domain == 0) {
            /// @dev We don't check if Guard is in dispute for accepting the snapshots.
            /// Guard could only be in Dispute, if they submitted a Report on a Notary.
            /// This should not strip away their ability to post snapshots, as they require
            /// a Notary signature in order to be used / gain tips anyway.

            // This will revert if Guard has previously submitted
            // a fresher state than one in the snapshot.
            _acceptGuardSnapshot(snapshot, agent);
        } else {
            // Check that Notary who submitted the snapshot is not in dispute
            require(!_inDispute(agent), "Notary is in dispute");
            // Fetch current Agent Root from BondingManager
            bytes32 agentRoot = agentManager.agentRoot();
            // This will revert if any of the states from the Notary snapshot
            // haven't been submitted by any of the Guards before.
            attPayload = _acceptNotarySnapshot(snapshot, agentRoot, agent);
            // Save attestation derived from Notary snapshot
            _saveAttestation(attPayload.castToAttestation(), agent);
        }
        emit SnapshotAccepted(status.domain, agent, snapPayload, snapSignature);
    }

    // ═════════════════════════════════════════════ VERIFY STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceSummit
    function verifyAttestation(bytes memory attPayload, bytes memory attSignature) external returns (bool isValid) {
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(attPayload);
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        isValid = _isValidAttestation(att);
        if (!isValid) {
            emit InvalidAttestation(attPayload, attSignature);
            // Slash Notary and notify local AgentManager
            _slashAgent(status.domain, notary);
        }
    }

    /// @inheritdoc InterfaceSummit
    function verifyAttestationReport(bytes memory arPayload, bytes memory arSignature)
        external
        returns (bool isValid)
    {
        // This will revert if payload is not an attestation report
        AttestationReport report = _wrapAttestationReport(arPayload);
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyAttestationReport(report, arSignature);
        // Guard needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // Report is valid, if the reported attestation is invalid
        isValid = !_isValidAttestation(report.attestation());
        if (!isValid) {
            emit InvalidAttestationReport(arPayload, arSignature);
            // Slash Guard and notify local AgentManager
            _slashAgent(0, guard);
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceSummit
    function getLatestState(uint32 origin) external view returns (bytes memory statePayload) {
        // TODO: implement once Agent Merkle Tree is done
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Saves the message from the receipt into the "quarantine queue". Once message leaves the queue,
    /// tips associated with the message are distributed across off-chain actors.
    function _saveReceipt(Receipt receipt) internal {
        bytes32 snapRoot = receipt.snapshotRoot();
        SnapRootData memory rootData = _rootData[snapRoot];
        require(rootData.submittedAt != 0, "Unknown snapshot root");
        // TODO: implement the quarantine queue
    }

    /// @inheritdoc DisputeHub
    function _beforeStatement() internal pure override returns (bool acceptNext) {
        // Summit is always open for new Guard/Notary statements
        return true;
    }
}
