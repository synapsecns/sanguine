// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract AttestationHubHarnessEvents {
    event LogAttestation(
        address[] guards,
        address[] notaries,
        bytes attestationView,
        bytes attestation
    );
}
