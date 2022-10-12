// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract AttestationHubHarnessEvents {
    event LogAttestation(address notary, bytes attestationView, bytes attestation);
}
