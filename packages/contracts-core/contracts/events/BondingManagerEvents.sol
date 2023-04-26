// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract BondingManagerEvents {
    /**
     * @notice Emitted when a snapshot is accepted by the Summit contract.
     * @param domain        Domain where the signed Agent is active (ZERO for Guards)
     * @param agent         Agent who signed the snapshot
     * @param snapPayload   Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapPayload, bytes snapSignature);

    /**
     * @notice Emitted when a proof of invalid attestation is submitted.
     * @param attPayload    Raw payload with Attestation data
     * @param attSignature  Notary signature for the attestation
     */
    event InvalidAttestation(bytes attPayload, bytes attSignature);

    /**
     * @notice Emitted when a proof of invalid attestation report is submitted.
     * @param arPayload     Raw payload with report data
     * @param arSignature   Guard signature for the report
     */
    event InvalidAttestationReport(bytes arPayload, bytes arSignature);
}
