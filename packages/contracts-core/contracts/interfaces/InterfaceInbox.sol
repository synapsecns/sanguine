// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceInbox {
    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Accepts a snapshot signed by a Guard or a Notary and passes it to Summit contract to save.
     * > Snapshot is a list of states for a set of Origin contracts residing on any of the chains.
     * - Guard-signed snapshots: all the states in the snapshot become available for Notary signing.
     * - Notary-signed snapshots: Snapshot Merkle Root is saved for valid snapshots, i.e.
     * snapshots which are only using states previously submitted by any of the Guards.
     * - Notary doesn't have to use states submitted by a single Guard in their snapshot.
     * - Notary could then proceed to sign the attestation for their submitted snapshot.
     * > Will revert if any of these is true:
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot signer is not an active Agent.
     * > - Agent snapshot contains a state with a nonce smaller or equal then they have previously submitted.
     * > - Notary snapshot contains a state that hasn't been previously submitted by any of the Guards.
     * > - Note: Agent will NOT be slashed for submitting such a snapshot.
     * @dev Notary will need to provide both `agentRoot` and `snapGas` when submitting an attestation on
     * the remote chain (the attestation contains only their merged hash). These are returned by this function,
     * and could be also obtained by calling `getAttestation(nonce)` or `getLatestNotaryAttestation(notary)`.
     * @param snapPayload       Raw payload with snapshot data
     * @param snapSignature     Agent signature for the snapshot
     * @return attPayload       Raw payload with data for attestation derived from Notary snapshot.
     *                          Empty payload, if a Guard snapshot was submitted.
     * @return agentRoot        Current root of the Agent Merkle Tree (zero, if a Guard snapshot was submitted)
     * @return snapGas          Gas data for each chain in the snapshot
     *                          Empty list, if a Guard snapshot was submitted.
     */
    function submitSnapshot(bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas);

    /**
     * @notice Accepts a receipt signed by a Notary and passes it to Summit contract to save.
     * > Receipt is a statement about message execution status on the remote chain.
     * - This will distribute the message tips across the off-chain actors once the receipt optimistic period is over.
     * > Will revert if any of these is true:
     * > - Receipt payload is not properly formatted.
     * > - Receipt signer is not an active Notary.
     * > - Receipt signer is in Dispute.
     * > - Receipt's snapshot root is unknown.
     * > - Provided tips could not be proven against the message hash.
     * @param rcptPayload       Raw payload with receipt data
     * @param rcptSignature     Notary signature for the receipt
     * @param paddedTips        Tips for the message execution
     * @param headerHash        Hash of the message header
     * @param bodyHash          Hash of the message body excluding the tips
     * @return wasAccepted      Whether the receipt was accepted
     */
    function submitReceipt(
        bytes memory rcptPayload,
        bytes memory rcptSignature,
        uint256 paddedTips,
        bytes32 headerHash,
        bytes32 bodyHash
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a Guard's receipt report signature, as well as Notary signature
     * for the reported Receipt.
     * > ReceiptReport is a Guard statement saying "Reported receipt is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a ReceiptReport and use receipt signature from
     * `verifyReceipt()` successful call that led to Notary being slashed in Summit on Synapse Chain.
     * > Will revert if any of these is true:
     * > - Receipt payload is not properly formatted.
     * > - Receipt Report signer is not an active Guard.
     * > - Receipt signer is not an active Notary.
     * @param rcptPayload       Raw payload with Receipt data that Guard reports as invalid
     * @param rcptSignature     Notary signature for the reported receipt
     * @param rrSignature       Guard signature for the report
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitReceiptReport(bytes memory rcptPayload, bytes memory rcptSignature, bytes memory rrSignature)
        external
        returns (bool wasAccepted);

    /**
     * @notice Passes the message execution receipt from Destination to the Summit contract to save.
     * > Will revert if any of these is true:
     * > - Called by anyone other than Destination.
     * @dev If a receipt is not accepted, any of the Notaries can submit it later using `submitReceipt`.
     * @param attNotaryIndex    Index of the Notary who signed the attestation
     * @param attNonce          Nonce of the attestation used for proving the executed message
     * @param paddedTips        Tips for the message execution
     * @param rcptPayload       Raw payload with message execution receipt
     * @return wasAccepted      Whether the receipt was accepted
     */
    function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes memory rcptPayload)
        external
        returns (bool wasAccepted);

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Verifies an attestation signed by a Notary.
     *  - Does nothing, if the attestation is valid (was submitted by this Notary as a snapshot).
     *  - Slashes the Notary, if the attestation is invalid.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValidAttestation   Whether the provided attestation is valid.
     *                              Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(bytes memory attPayload, bytes memory attSignature)
        external
        returns (bool isValidAttestation);

    /**
     * @notice Verifies a Guard's attestation report signature.
     *  - Does nothing, if the report is valid (if the reported attestation is invalid).
     *  - Slashes the Guard, if the report is invalid (if the reported attestation is valid).
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation Report signer is not an active Guard.
     * @param attPayload        Raw payload with Attestation data that Guard reports as invalid
     * @param arSignature       Guard signature for the report
     * @return isValidReport    Whether the provided report is valid.
     *                          Guard is slashed, if return value is FALSE.
     */
    function verifyAttestationReport(bytes memory attPayload, bytes memory arSignature)
        external
        returns (bool isValidReport);
}
