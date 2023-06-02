// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceLightInbox {
    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Accepts an attestation signed by a Notary and passes it to Destination contract to save.
     * > Attestation is created whenever a Notary-signed snapshot is saved in Summit on Synapse Chain.
     * - Saved Attestation could be later used to prove the inclusion of message in the Origin Merkle Tree.
     * - Messages coming from chains included in the Attestation's snapshot could be proven.
     * - Proof only exists for messages that were sent prior to when the Attestation's snapshot was taken.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary for local domain.
     * > - Attestation signer is in Dispute.
     * > - Attestation's snapshot root has been previously submitted.
     * > - Attestation's data hash doesn't match the hash of provided agentRoot and snapshot gas data.
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @param agentRoot         Agent Merkle Root from the Attestation
     * @param snapGas           Gas data for each chain in the snapshot
     * @return wasAccepted      Whether the Attestation was accepted
     */
    function submitAttestation(
        bytes memory attPayload,
        bytes memory attSignature,
        bytes32 agentRoot,
        uint256[] memory snapGas
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a Guard's attestation report signature, as well as Notary signature
     * for the reported Attestation.
     * > AttestationReport is a Guard statement saying "Reported attestation is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a AttestationReport and use attestation signature from
     * `verifyAttestation()` successful call that led to Notary being slashed in Summit on Synapse Chain.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation Report signer is not an active Guard.
     * > - Attestation signer is not an active Notary for local domain.
     * @param attPayload        Raw payload with Attestation data that Guard reports as invalid
     * @param arSignature       Guard signature for the report
     * @param attSignature      Notary signature for the reported attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitAttestationReport(bytes memory attPayload, bytes memory arSignature, bytes memory attSignature)
        external
        returns (bool wasAccepted);
}
