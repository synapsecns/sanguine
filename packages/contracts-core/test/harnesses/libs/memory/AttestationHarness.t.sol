// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {Attestation, AttestationLib, MemView, MemViewLib} from "../../../../contracts/libs/memory/Attestation.sol";

// solhint-disable ordering
/// @notice Exposes Attestation methods for testing against golang.
contract AttestationHarness {
    using AttestationLib for bytes;
    using AttestationLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToAttestation(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Attestation attestation = AttestationLib.castToAttestation(payload);
        return attestation.unwrap().clone();
    }

    function snapRoot(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestation().snapRoot();
    }

    function dataHash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestation().dataHash();
    }

    function dataHash(bytes32 agentRoot_, bytes32 snapGasHash_) public pure returns (bytes32) {
        return AttestationLib.dataHash(agentRoot_, snapGasHash_);
    }

    function nonce(bytes memory payload) public pure returns (uint32) {
        return payload.castToAttestation().nonce();
    }

    function blockNumber(bytes memory payload) public pure returns (uint40) {
        return payload.castToAttestation().blockNumber();
    }

    function timestamp(bytes memory payload) public pure returns (uint40) {
        return payload.castToAttestation().timestamp();
    }

    function hashValid(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestation().hashValid();
    }

    function hashInvalid(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestation().hashInvalid();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatAttestation(
        bytes32 snapRoot_,
        bytes32 dataHash_,
        uint32 nonce_,
        uint40 blockNumber_,
        uint40 timestamp_
    ) public pure returns (bytes memory) {
        return AttestationLib.formatAttestation(snapRoot_, dataHash_, nonce_, blockNumber_, timestamp_);
    }

    function isAttestation(bytes memory payload) public pure returns (bool) {
        return payload.ref().isAttestation();
    }
}
