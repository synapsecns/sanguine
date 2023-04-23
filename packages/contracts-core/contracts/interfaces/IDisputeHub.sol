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
