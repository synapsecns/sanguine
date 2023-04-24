// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag} from "../libs/Structures.sol";

abstract contract AgentManagerEvents {
    /**
     * @notice Emitted when a proof of invalid receipt statement is submitted.
     * @param rcptPayload   Raw payload with the receipt statement
     * @param rcptSignature Notary signature for the receipt statement
     */
    event InvalidReceipt(bytes rcptPayload, bytes rcptSignature);

    /**
     * @notice Emitted when a proof of invalid state in the signed attestation is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param statePayload  Raw payload with state data
     * @param attPayload    Raw payload with Attestation data for snapshot
     * @param attSignature  Notary signature for the attestation
     */
    event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature);

    /**
     * @notice Emitted when a proof of invalid state in the signed snapshot is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param snapPayload   Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature);

    /**
     * @notice Emitted when a proof of invalid state report is submitted.
     * @param srPayload     Raw payload with report data
     * @param srSignature   Guard signature for the report
     */
    event InvalidStateReport(bytes srPayload, bytes srSignature);

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
}
