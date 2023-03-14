// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { Attestation, Snapshot, StatementHub, StateReport } from "./StatementHub.sol";
import { DisputeHubEvents } from "../events/DisputeHubEvents.sol";
import { IDisputeHub } from "../interfaces/IDisputeHub.sol";

abstract contract DisputeHub is StatementHub, DisputeHubEvents, IDisputeHub {
    /// @inheritdoc IDisputeHub
    function submitStateReport(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes memory _snapPayload,
        bytes memory _snapSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a state report
        StateReport report = _wrapStateReport(_srPayload);
        // This will revert if the report signer is not an active Guard
        address guard = _verifyStateReport(report, _srSignature);
        // This will revert if payload is not a snapshot
        Snapshot snapshot = _wrapSnapshot(_snapPayload);
        // This will revert if the snapshot signer is not an active Agent
        (uint32 domain, address notary) = _verifySnapshot(snapshot, _snapSignature);
        // Snapshot signer needs to be a Notary, not a Guard
        require(domain != 0, "Snapshot signer is not a Notary");
        // Snapshot state and reported state need to be the same
        // This will revert if state index is out of range
        require(snapshot.state(_stateIndex).equals(report.state()), "States don't match");
        // Reported State was used by the Notary for their signed snapshot => open dispute
        _openDispute(guard, domain, notary);
        return true;
    }

    /// @inheritdoc IDisputeHub
    function submitStateReportWithProof(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes32[] memory _snapProof,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a state report
        StateReport report = _wrapStateReport(_srPayload);
        // This will revert if the report signer is not an active Guard
        address guard = _verifyStateReport(report, _srSignature);
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(_attPayload);
        // This will revert if signer is not an active Notary
        (uint32 domain, address notary) = _verifyAttestation(att, _attSignature);
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof has length different to Attestation height.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, _stateIndex, report.state(), _snapProof);
        // Reported State was used by the Notary for their signed attestation => open dispute
        _openDispute(guard, domain, notary);
        return true;
    }

    /// @dev Opens a Dispute between a Guard and a Notary.
    /// This should be called, when the Guard submits a Report on a statement signed by the Notary.
    function _openDispute(
        address _guard,
        uint32 _domain,
        address _notary
    ) internal {
        // TODO: implement this
        emit Dispute(_guard, _domain, _notary);
    }
}
