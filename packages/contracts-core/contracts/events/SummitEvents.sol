// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the Summit contract
abstract contract SummitEvents {
    /**
     * @notice Emitted when a proof of invalid attestation is submitted.
     * @param attestation   Raw payload with attestation data
     * @param attSignature  Notary signature for the attestation
     */
    event InvalidAttestation(bytes attestation, bytes attSignature);

    /**
     * @notice Emitted when a proof of invalid attestation report is submitted.
     * @param arPayload     Raw payload with report data
     * @param arSignature   Guard signature for the report
     */
    event InvalidAttestationReport(bytes arPayload, bytes arSignature);

    /**
     * @notice Emitted when a snapshot is accepted by the Summit contract.
     * @param domain        Domain where the signed Notary is active
     * @param notary        Notary who signed the attestation
     * @param rcptPayload   Raw payload with receipt data
     * @param rcptSignature Notary signature for the receipt
     */
    event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature);

    /**
     * @notice Emitted when a snapshot is accepted by the Summit contract.
     * @param domain        Domain where the signed Agent is active (ZERO for Guards)
     * @param agent         Agent who signed the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature);

    /**
     * @notice Emitted when a tip is awarded to the actor, whether they are bonded or unbonded actor.
     * @param actor     Actor address
     * @param origin    Domain where tips were originally paid
     * @param tip       Tip value, scaled down by TIPS_MULTIPLIER
     */
    event TipAwarded(address actor, uint32 origin, uint256 tip);
}
