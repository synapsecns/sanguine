// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract AttestationHubEvents {
    /**
     * @notice Emitted when an attestation is submitted to AttestationHub.
     * @param guards        Guards who signed the attestation
     * @param notaries      Notaries who signed the attestation
     * @param attestation   Raw payload with attestation data and notary signature
     */
    event AttestationAccepted(address[] guards, address[] notaries, bytes attestation);
}
