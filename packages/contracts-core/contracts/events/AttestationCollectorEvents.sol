// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract AttestationCollectorEvents {
    event AttestationSubmitted(address indexed notary, bytes attestation);
}
