// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract DestinationHubEvents {
    // TODO: emit full attestation payload instead of signature?
    event AttestationAccepted(
        uint32 indexed origin,
        uint32 indexed nonce,
        bytes32 indexed root,
        bytes signature
    );
}
