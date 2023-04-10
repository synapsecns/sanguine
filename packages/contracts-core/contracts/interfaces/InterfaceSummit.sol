// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceSummit {
    /**
     * @notice Submit a message execution receipt. This will distribute the message tips
     * across the off-chain actors once the receipt optimistic period is over.
     * @dev Will revert if any of these is true:
     *  - Receipt payload is not properly formatted.
     *  - Receipt signer is not an active Notary.
     *  - Receipt's snapshot root is unknown.
     * @param rcptPayload       Raw payload with receipt data
     * @param rcptSignature     Notary signature for the receipt
     * @return wasAccepted      Whether the receipt was accepted
     */
    function submitReceipt(bytes memory rcptPayload, bytes memory rcptSignature) external returns (bool wasAccepted);

    /**
     * @notice Submit a snapshot (list of states) signed by a Guard or a Notary.
     * Guard-signed snapshots: all the states in the snapshot become available for Notary signing.
     * Notary-signed snapshots: Snapshot Merkle Root is saved for valid snapshots, i.e.
     * snapshots which are only using states previously submitted by any of the Guards.
     * Notary doesn't have to use states submitted by a single Guard in their snapshot.
     * Notary could then proceed to sign the attestation for their submitted snapshot.
     * @dev Will revert if any of these is true:
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Agent.
     *  - Guard snapshot contains a state older then they have previously submitted
     *  - Notary snapshot contains a state that hasn't been previously submitted by a Guard.
     * Note that Notary will NOT be slashed for submitting such a snapshot.
     * @param snapPayload       Raw payload with snapshot data
     * @param snapSignature     Agent signature for the snapshot
     * @return attPayload       Raw payload with data for attestation derived from Notary snapshot.
     *                          Empty payload, if a Guard snapshot was submitted.
     */
    function submitSnapshot(bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bytes memory attPayload);

    /**
     * @notice Verifies an attestation signed by a Notary.
     *  - Does nothing, if the attestation is valid (was submitted by this Notary as a snapshot).
     *  - Slashes the Notary otherwise (meaning the attestation is invalid).
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValid          Whether the provided attestation is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(bytes memory attPayload, bytes memory attSignature) external returns (bool isValid);

    /**
     * @notice Verifies an attestation report signed by a Guard.
     *  - Does nothing, if the report is valid (if the reported attestation is invalid).
     *  - Slashes the Guard otherwise (meaning the reported attestation is valid, making the report invalid).
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     * @param arPayload         Raw payload with AttestationReport data
     * @param arSignature       Guard signature for the report
     * @return isValid          Whether the provided report is valid.
     *                          Guard is slashed, if return value is FALSE.
     */
    function verifyAttestationReport(bytes memory arPayload, bytes memory arSignature)
        external
        returns (bool isValid);

    /**
     * @notice Returns the state with the highest known nonce
     * submitted by any of the currently active Guards.
     * @param origin        Domain of origin chain
     * @return statePayload Raw payload with latest active Guard state for origin
     */
    function getLatestState(uint32 origin) external view returns (bytes memory statePayload);
}
