// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, Dispute} from "../libs/Structures.sol";

abstract contract AgentManagerEvents {
    // ═══════════════════════════════════════════════ DATA UPDATED ════════════════════════════════════════════════════

    /**
     * @notice Emitted whenever the root of the Agent Merkle Tree is updated.
     * @param newRoot   New agent merkle root
     */
    event RootUpdated(bytes32 newRoot);

    /**
     * @notice Emitted whenever a status of the agent is updated.
     * @dev Only Active/Unstaking/Resting/Slashed flags could be stored in the Agent Merkle Tree.
     * Unknown flag is the default (zero) value and is used to represent agents that never
     * interacted with the BondingManager contract.
     * Fraudulent flag is the value for the agent who has been proven to commit fraud, but their
     * status hasn't been updated to Slashed in the Agent Merkle Tree. This is due to the fact
     * that the update of the status requires a merkle proof of the old status, and happens
     * in a separate transaction because of that.
     * @param flag      Flag defining agent status:
     * @param domain    Domain assigned to the agent (ZERO for Guards)
     * @param agent     Agent address
     */
    event StatusUpdated(AgentFlag flag, uint32 indexed domain, address indexed agent);

    /**
     * @notice Emitted whenever a Dispute status of an the agent is updated.
     * See Structures.sol for details.
     * @param agent     Agent address
     * @param dispute   Dispute status: (flag, rivalIndex, fraudProver)
     */
    event DisputeUpdated(address agent, Dispute dispute);
}
