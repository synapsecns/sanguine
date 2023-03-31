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
     * @param domain        Domain where the signed Agent is active (ZERO for Guards)
     * @param agent         Agent who signed the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature);
}
