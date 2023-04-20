// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceSummit {
    /**
     * @notice Distributes tips using the first Receipt from the "receipt quarantine queue".
     * Possible scenarios:
     *  - Receipt queue is empty => does nothing
     *  - Receipt optimistic period is not over => does nothing
     *  - Either of Notaries present in Receipt was slashed => receipt is deleted from the queue
     *  - Either of Notaries present in Receipt in Dispute => receipt is moved to the end of queue
     *  - None of the above => receipt tips are distributed
     * @dev Returned value makes it possible to do the following: `while (distributeTips()) {}`
     * @return queuePopped      Whether the first element was popped from the queue
     */
    function distributeTips() external returns (bool queuePopped);

    /**
     * @notice Withdraws locked base message tips from requested domain Origin to the recipient.
     * This is done by a call to a local Origin contract, or by a manager message to the remote chain.
     * @dev This will revert, if the pending balance of origin tips (earned-claimed) is lower than requested.
     * @param origin    Domain of chain to withdraw tips on
     * @param amount    Amount of tips to withdraw
     */
    function withdrawTips(uint32 origin, uint256 amount) external;

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

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns earned and claimed tips for the actor.
     * Note: Tips for address(0) belong to the Treasury.
     * @param actor     Address of the actor
     * @param origin    Domain where the tips were initially paid
     * @return earned   Total amount of origin tips the actor has earned so far
     * @return claimed  Total amount of origin tips the actor has claimed so far
     */
    function actorTips(address actor, uint32 origin) external view returns (uint128 earned, uint128 claimed);

    /**
     * @notice Returns the amount of receipts in the "Receipt Quarantine Queue".
     */
    function receiptQueueLength() external view returns (uint256);

    /**
     * @notice Returns the state with the highest known nonce
     * submitted by any of the currently active Guards.
     * @param origin        Domain of origin chain
     * @return statePayload Raw payload with latest active Guard state for origin
     */
    function getLatestState(uint32 origin) external view returns (bytes memory statePayload);
}
