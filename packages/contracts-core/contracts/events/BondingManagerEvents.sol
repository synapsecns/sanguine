// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract BondingManagerEvents {
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
