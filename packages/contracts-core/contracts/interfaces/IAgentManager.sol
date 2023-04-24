// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus} from "../libs/Structures.sol";

interface IAgentManager {
    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Accepts a state report signed by a Guard, a Snapshot containing the reported State,
     * as well as Notary signature for the Snapshot.
     * > StateReport is a Guard statement saying "Reported state is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a StateReport and use other values from
     * `verifyStateWithSnapshot()` successful call that led to Notary being slashed in remote Origin.
     * > Will revert if any of these is true:
     * > - Report payload is not properly formatted.
     * > - Report signer is not an active Guard.
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot signer is not an active Notary.
     * > - State index is out of range.
     * > - Snapshot's state and reported state don't match.
     * > - The Guard or the Notary are already in a Dispute
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param srPayload         Raw payload with StateReport data
     * @param srSignature       Guard signature for the report
     * @param snapPayload       Raw payload with Snapshot data
     * @param snapSignature     Notary signature for the Snapshot
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithSnapshot(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a state report signed by a Guard, a Snapshot containing the reported State,
     * as well as Notary signature for the Attestation created from this Snapshot.
     * > StateReport is a Guard statement saying "Reported state is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a StateReport and use other values from
     * `verifyStateWithAttestation()` successful call that led to Notary being slashed in remote Origin.
     * > Will revert if any of these is true:
     * > - Report payload is not properly formatted.
     * > - Report signer is not an active Guard.
     * > - Snapshot payload is not properly formatted.
     * > - State index is out of range.
     * > - Snapshot's state and reported state don't match.
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from the Snapshot.
     * > - The Guard or the Notary are already in a Dispute
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param srPayload         Raw payload with StateReport data
     * @param srSignature       Guard signature for the report
     * @param snapPayload       Raw payload with Snapshot data
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the Attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithAttestation(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a state report signed by a Guard, a proof of inclusion of the reported State in an Attestation,
     * as well as Notary signature for the Attestation.
     * > StateReport is a Guard statement saying "Reported state is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a StateReport and use other values from
     * `verifyStateWithSnapshotProof()` successful call that led to Notary being slashed in remote Origin.
     * > Will revert if any of these is true:
     * > - Report payload is not properly formatted.
     * > - Report signer is not an active Guard.
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from State and Snapshot Proof.
     * > - Snapshot Proof's first element does not match the State metadata.
     * > - Snapshot Proof length exceeds Snapshot Tree Height.
     * > - State index is out of range.
     * > - The Guard or the Notary are already in a Dispute
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param srPayload         Raw payload with StateReport data
     * @param srSignature       Guard signature for the report
     * @param snapProof         Proof of inclusion of reported State's Left Leaf into Snapshot Merkle Tree
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the Attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithSnapshotProof(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted);

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Verifies a message receipt signed by the Notary.
     * - Does nothing, if the receipt is valid (matches the saved receipt data for the referenced message).
     * - Slashes the Notary, if the receipt is invalid.
     * > Will revert if any of these is true:
     * > - Receipt payload is not properly formatted.
     * > - Receipt signer is not an active Notary.
     * > - Receipt's destination chain does not refer to this chain.
     * @param rcptPayload       Raw payload with Receipt data
     * @param rcptSignature     Notary signature for the receipt
     * @return isValidReceipt   Whether the provided receipt is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyReceipt(bytes memory rcptPayload, bytes memory rcptSignature)
        external
        returns (bool isValidReceipt);

    /**
     * @notice Verifies a state from the snapshot, that was used for the Notary-signed attestation.
     * - Does nothing, if the state is valid (matches the historical state of this contract).
     * - Slashes the Notary, if the state is invalid.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from the Snapshot.
     * > - Snapshot payload is not properly formatted.
     * > - State index is out of range.
     * > - State does not refer to this chain.
     * @param stateIndex        State index to check
     * @param snapPayload       Raw payload with snapshot data
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValidState     Whether the provided state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyStateWithAttestation(
        uint256 stateIndex,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState);

    /**
     * @notice Verifies a state from the snapshot, that was used for the Notary-signed attestation.
     * - Does nothing, if the state is valid (matches the historical state of this contract).
     * - Slashes the Notary, if the state is invalid.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from State and Snapshot Proof.
     * > - Snapshot Proof's first element does not match the State metadata.
     * > - Snapshot Proof length exceeds Snapshot Tree Height.
     * > - State payload is not properly formatted.
     * > - State index is out of range.
     * > - State does not refer to this chain.
     * @param stateIndex        Index of state in the snapshot
     * @param statePayload      Raw payload with State data to check
     * @param snapProof         Proof of inclusion of provided State's Left Leaf into Snapshot Merkle Tree
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValidState     Whether the provided state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyStateWithSnapshotProof(
        uint256 stateIndex,
        bytes memory statePayload,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState);

    /**
     * @notice Verifies a state from the snapshot (a list of states) signed by a Guard or a Notary.
     * - Does nothing, if the state is valid (matches the historical state of this contract).
     * - Slashes the Agent, if the state is invalid.
     * > Will revert if any of these is true:
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot signer is not an active Agent.
     * > - State index is out of range.
     * > - State does not refer to this chain.
     * @param stateIndex        State index to check
     * @param snapPayload       Raw payload with snapshot data
     * @param snapSignature     Agent signature for the snapshot
     * @return isValidState     Whether the provided state is valid.
     *                          Agent is slashed, if return value is FALSE.
     */
    function verifyStateWithSnapshot(uint256 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValidState);

    /**
     * @notice Verifies a state report signed by a Guard.
     *  - Does nothing, if the report is valid (if the reported state is invalid).
     *  - Slashes the Guard, if the report is invalid (if the reported state is valid).
     * > Will revert if any of these is true:
     * > - Report payload is not properly formatted.
     * > - Report signer is not an active Guard.
     * > - Reported State does not refer to this chain.
     * @param srPayload         Raw payload with StateReport data
     * @param srSignature       Guard signature for the report
     * @return isValidReport    Whether the provided report is valid.
     *                          Guard is slashed, if return value is FALSE.
     */
    function verifyStateReport(bytes memory srPayload, bytes memory srSignature)
        external
        returns (bool isValidReport);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the latest known root of the Agent Merkle Tree.
     */
    function agentRoot() external view returns (bytes32);

    /**
     * @notice Returns (flag, domain, index) for a given agent. See Structures.sol for details.
     * @dev Will return AgentFlag.Fraudulent for agents that have been proven to commit fraud,
     * but their status is not updated to Slashed yet.
     * @param agent     Agent address
     * @return          Status for the given agent: (flag, domain, index).
     */
    function agentStatus(address agent) external view returns (AgentStatus memory);

    /**
     * @notice Returns agent address and their current status for a given agent index.
     * @dev Will return empty values if agent with given index doesn't exist.
     * @param index     Agent index in the Agent Merkle Tree
     * @return agent    Agent address
     * @return status   Status for the given agent: (flag, domain, index)
     */
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status);

    /**
     * @notice Returns whether the agent has been slashed.
     * @param agent         Agent address
     * @return isSlashed    Whether the agent has been slashed
     * @return prover       Address that presented the proof of fraud committed by the agent
     */
    function slashStatus(address agent) external view returns (bool isSlashed, address prover);
}
