// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/Attestation.sol";
import {Receipt, ReceiptLib} from "../libs/Receipt.sol";
import {Snapshot, SnapshotLib} from "../libs/Snapshot.sol";
import {State, StateLib} from "../libs/State.sol";
import {StateReport, StateReportLib} from "../libs/StateReport.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {IDisputeHub} from "../interfaces/IDisputeHub.sol";
import {IExecutionHub} from "../interfaces/IExecutionHub.sol";
import {IStateHub} from "../interfaces/IStateHub.sol";
import {IStatementManager} from "../interfaces/IStatementManager.sol";
import {AgentManager, AgentStatus} from "./AgentManager.sol";

abstract contract StatementManager is AgentManager, IStatementManager {
    using AttestationLib for bytes;
    using ReceiptLib for bytes;
    using StateLib for bytes;
    using StateReportLib for bytes;
    using SnapshotLib for bytes;

    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc IStatementManager
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
        _verifyActive(guardStatus);
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the snapshot signer is not a known Agent
        (AgentStatus memory notaryStatus, address notary) = _verifySnapshot(snapshot, snapSignature);
        // Snapshot signer needs to be a Notary, not a Guard
        require(notaryStatus.domain != 0, "Snapshot signer is not a Notary");
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(notaryStatus);
        // Snapshot state and reported state need to be the same
        // This will revert if state index is out of range
        require(snapshot.state(stateIndex).equals(report.state()), "States don't match");
        // TODO: open a Dispute - which will revert if either actor is already in dispute
        IDisputeHub(destination);
        return true;
    }

    /// @inheritdoc IStatementManager
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
        _verifyActive(guardStatus);
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not an known Notary
        (AgentStatus memory notaryStatus, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(notaryStatus);
        require(snapshot.root() == att.snapRoot(), "Attestation not matches snapshot");
        // TODO: open a Dispute - which will revert if either actor is already in dispute
        IDisputeHub(destination);
        return true;
    }

    /// @inheritdoc IStatementManager
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
        _verifyActive(guardStatus);
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not a known Notary
        (AgentStatus memory notaryStatus, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(notaryStatus);
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, stateIndex, report.state(), snapProof);
        // TODO: open a Dispute - which will revert if either actor is already in dispute
        IDisputeHub(destination);
        return true;
    }

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc IStatementManager
    function verifyReceipt(bytes memory rcptPayload, bytes memory rcptSignature)
        external
        returns (bool isValidReceipt)
    {
        // This will revert if payload is not an receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyReceipt(rcpt, rcptSignature);
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        isValidReceipt = IExecutionHub(destination).isValidReceipt(rcptPayload);
        if (!isValidReceipt) {
            // TODO: Slash Notary
            notary;
        }
    }

    /// @inheritdoc IStatementManager
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
        _verifyActiveUnstaking(status);
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        require(snapshot.root() == att.snapRoot(), "Attestation not matches snapshot");
        // This will revert if state does not refer to this chain
        isValidState = IStateHub(origin).isValidState(snapshot.state(stateIndex).unwrap().clone());
        if (!isValidState) {
            // TODO: Slash Notary
            notary;
        }
    }

    /// @inheritdoc IStatementManager
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
        _verifyActiveUnstaking(status);
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
            // TODO: Slash Notary
            notary;
        }
    }

    /// @inheritdoc IStatementManager
    function verifyStateWithSnapshot(uint256 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValidState)
    {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the snapshot signer is not a known Agent
        (AgentStatus memory status, address agent) = _verifySnapshot(snapshot, snapSignature);
        // Agent needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // This will revert if state does not refer to this chain
        isValidState = IStateHub(origin).isValidState(snapshot.state(stateIndex).unwrap().clone());
        if (!isValidState) {
            // TODO: Slash Agent
            agent;
        }
    }

    /// @inheritdoc IStatementManager
    function verifyStateReport(bytes memory srPayload, bytes memory srSignature)
        external
        returns (bool isValidReport)
    {
        // This will revert if payload is not a snapshot report
        StateReport report = srPayload.castToStateReport();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyStateReport(report, srSignature);
        // Guard needs to be Active/Unstaking
        _verifyActiveUnstaking(status);
        // Report is valid IF AND ONLY IF the reported state in invalid
        // This will revert if the reported state does not refer to this chain
        isValidReport = !IStateHub(origin).isValidState(report.state().unwrap().clone());
        if (!isValidReport) {
            // TODO: Slash Guard
            guard;
        }
    }
}
