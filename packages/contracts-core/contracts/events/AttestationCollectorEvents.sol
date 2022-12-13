// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract AttestationCollectorEvents {
    event AttestationSaved(uint256 indexed attestationIndex, bytes attestation);
}
