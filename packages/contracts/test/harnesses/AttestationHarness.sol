// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Attestation } from "../../contracts/libs/Attestation.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

contract AttestationHarness {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    function formatAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root,
        bytes memory _signature
    ) public pure returns (bytes memory) {
        return
            Attestation.formatAttestation(
                formatAttestationData(_domain, _nonce, _root),
                _signature
            );
    }

    function formatAttestationData(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) public pure returns (bytes memory) {
        return Attestation.formatAttestationData(_domain, _nonce, _root);
    }

    function isAttestation(bytes memory _attestation) public pure returns (bool) {
        return _attestation.castToAttestation().isAttestation();
    }

    function domain(bytes memory _attestation) public pure returns (uint32) {
        return _attestation.castToAttestation().attestedDomain();
    }

    function nonce(bytes memory _attestation) public pure returns (uint32) {
        return _attestation.castToAttestation().attestedNonce();
    }

    function root(bytes memory _attestation) public pure returns (bytes32) {
        return _attestation.castToAttestation().attestedRoot();
    }

    function data(bytes memory _attestation) public view returns (bytes memory) {
        return _attestation.castToAttestation().attestationData().clone();
    }

    function signature(bytes memory _attestation) public view returns (bytes memory) {
        return _attestation.castToAttestation().attestationSignature().clone();
    }
}
