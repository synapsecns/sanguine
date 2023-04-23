// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DisputeStatus} from "../libs/Structures.sol";

/// @notice Common functions for contracts relying on Agent-signed Statements
interface IDisputeHub {
    /**
     * @notice Opens a Dispute between the Guard and the Notary. Could be only called by a local `AgentManager` after
     * the Guard submits a report for a Notary-signed statement, indicating that either the Notary committed fraud,
     * or the Guard supplied an invalid report.
     * - Any agent could be only in a single Dispute.
     * - All Notary-signed statements (snapshots, attestations, receipts) could not be used if a Notary is in Dispute.
     * - New reports from the Guard are rejected, if a Guard is in Dispute.
     * - Guard could continue supplying signed snapshots to `Summit`, while in Dispute.
     * > Will revert if either Guard or a Notary is already in Dispute.
     * @param guard     Address of the Guard to be put into Dispute
     * @param domain    Domain where the Notary is active
     * @param notary    Address of the Notary to be put into Dispute
     */
    function openDispute(address guard, uint32 domain, address notary) external;

    /**
     * @notice Submit an StateReport signed by a Guard, a Snapshot containing the reported State,
     * as well as Notary signature for the Snapshot.
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Notary.
     *  - State index is out of range.
     *  - Snapshot's state and reported state don't match.
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param srPayload         Raw payload with StateReport data
     * @param srSignature       Guard signature for the report
     * @param snapPayload       Raw payload with Snapshot data
     * @param snapSignature     Notary signature for the Snapshot
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReport(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bool wasAccepted);

    /**
     * @notice Submit an StateReport signed by a Guard, a proof of inclusion
     * of the reported State in an Attestation, as well as Notary signature for the Attestation.
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     *  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
     *  - Snapshot Proof's first element does not match the State metadata.
     *  - Snapshot Proof length exceeds Snapshot tree Height.
     *  - State index is out of range.
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param srPayload         Raw payload with StateReport data
     * @param srSignature       Guard signature for the report
     * @param snapProof         Proof of inclusion of reported State's Left Leaf into Snapshot Merkle Tree
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the Attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithProof(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the Dispute status of the given agent.
     * @param agent     Agent address
     * @return status   Struct with the Dispute status of the agent:
     *                  - flag: None/Pending/Slashed (see Structures.sol)
     *                  - counterpart: who the agent is in dispute with
     */
    function disputeStatus(address agent) external view returns (DisputeStatus memory status);
}
