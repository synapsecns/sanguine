// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceDestination {
    /**
     * @notice Attempts to pass a quarantined Agent Merkle Root to a local Light Manager.
     * @dev Will do nothing, if root optimistic period is not over.
     * Note: both returned values can not be true.
     * @return rootPassed   Whether the agent merkle root was passed to LightManager
     * @return rootPending  Whether there is a pending agent merkle root left
     */
    function passAgentRoot() external returns (bool rootPassed, bool rootPending);

    /**
     * @notice Submit an Attestation signed by a Notary.
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary for local domain.
     *  - Attestation's snapshot root has been previously submitted.
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the reported attestation
     * @return wasAccepted      Whether the Attestation was accepted (resulting in Dispute between the agents)
     */
    function submitAttestation(bytes memory attPayload, bytes memory attSignature)
        external
        returns (bool wasAccepted);

    /**
     * @notice Submit an AttestationReport signed by a Guard, as well as Notary signature
     * for the reported Attestation.
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     *  - Attestation signer is not an active Notary for local domain.
     * @param arPayload         Raw payload with AttestationReport data
     * @param arSignature       Guard signature for the report
     * @param attSignature      Notary signature for the reported attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitAttestationReport(bytes memory arPayload, bytes memory arSignature, bytes memory attSignature)
        external
        returns (bool wasAccepted);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the total amount of Notaries attestations that have been accepted.
     */
    function attestationsAmount() external view returns (uint256);

    /**
     * @notice Returns a Notary-signed attestation with a given index. Index refers to the list of all attestations
     * accepted by this contract.
     * @param index             Attestation index
     * @return attPayload       Raw payload with Attestation data
     * @return attSignature     Notary signature for the reported attestation
     */
    function getSignedAttestation(uint256 index)
        external
        view
        returns (bytes memory attPayload, bytes memory attSignature);

    /**
     * Returns status of Destination contract as far as snapshot/agent roots are concerned
     * @return snapRootTime     Timestamp when latest snapshot root was accepted
     * @return agentRootTime    Timestamp when latest agent root was accepted
     * @return notary           Notary who signed the latest agent root
     */
    function destStatus() external view returns (uint48 snapRootTime, uint48 agentRootTime, address notary);

    /**
     * Returns Agent Merkle Root to be passed to LightManager once its optimistic period is over.
     */
    function nextAgentRoot() external view returns (bytes32);
}
