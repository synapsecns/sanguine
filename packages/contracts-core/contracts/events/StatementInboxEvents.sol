// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract StatementInboxEvents {
    // ════════════════════════════════════════════ STATEMENT ACCEPTED ═════════════════════════════════════════════════

    /**
     * @notice Emitted when a snapshot is accepted by the Destination contract.
     * @param domain        Domain where the signed Notary is active
     * @param notary        Notary who signed the attestation
     * @param attPayload    Raw payload with attestation data
     * @param attSignature  Notary signature for the attestation
     */
    event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature);

    // ═════════════════════════════════════════ INVALID STATEMENT PROVED ══════════════════════════════════════════════

    /**
     * @notice Emitted when a proof of invalid receipt statement is submitted.
     * @param rcptPayload   Raw payload with the receipt statement
     * @param rcptSignature Notary signature for the receipt statement
     */
    event InvalidReceipt(bytes rcptPayload, bytes rcptSignature);

    /**
     * @notice Emitted when a proof of invalid receipt report is submitted.
     * @param rrPayload     Raw payload with report data
     * @param rrSignature   Guard signature for the report
     */
    event InvalidReceiptReport(bytes rrPayload, bytes rrSignature);

    /**
     * @notice Emitted when a proof of invalid state in the signed attestation is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param statePayload  Raw payload with state data
     * @param attPayload    Raw payload with Attestation data for snapshot
     * @param attSignature  Notary signature for the attestation
     */
    event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature);

    /**
     * @notice Emitted when a proof of invalid state in the signed snapshot is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param snapPayload   Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature);

    /**
     * @notice Emitted when a proof of invalid state report is submitted.
     * @param srPayload     Raw payload with report data
     * @param srSignature   Guard signature for the report
     */
    event InvalidStateReport(bytes srPayload, bytes srSignature);
}
