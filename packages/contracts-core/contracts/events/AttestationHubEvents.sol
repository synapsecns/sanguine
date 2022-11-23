// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract AttestationHubEvents {
    /**
     * @notice Emitted when an attestation is submitted to AttestationHub.
     * @param notary        Notary who signed the attestation
     * @param attestation   Raw payload with attestation data and notary signature
     */
    event AttestationAccepted(address indexed notary, bytes attestation);
}
