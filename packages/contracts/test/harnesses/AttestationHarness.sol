// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Attestation } from "../../contracts/libs/Attestation.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

contract AttestationHarness {
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    function formatAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) public pure returns (bytes memory) {
        return Attestation.formatAttestation(_domain, _nonce, _root);
    }

    function isValid(bytes memory _attestation) public pure returns (bool) {
        return _attestation.ref(0).isValidAttestation();
    }

    function domain(bytes memory _attestation) public pure returns (uint32) {
        return _attestation.ref(0).attestationDomain();
    }

    function nonce(bytes memory _attestation) public pure returns (uint32) {
        return _attestation.ref(0).attestationNonce();
    }

    function root(bytes memory _attestation) public pure returns (bytes32) {
        return _attestation.ref(0).attestationRoot();
    }
}
