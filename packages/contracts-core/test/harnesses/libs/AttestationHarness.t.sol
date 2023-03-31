// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    Attestation,
    AttestationLib,
    ExecutionAttestation,
    SummitAttestation,
    TypedMemView
} from "../../../contracts/libs/Attestation.sol";

/**
 * @notice Exposes Attestation methods for testing against golang.
 */
contract AttestationHarness {
    using AttestationLib for bytes;
    using AttestationLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToAttestation(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Attestation _attestation = AttestationLib.castToAttestation(payload);
        return _attestation.unwrap().clone();
    }

    function snapRoot(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestation().snapRoot();
    }

    function agentRoot(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestation().agentRoot();
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

    function hash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestation().hash();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       DESTINATION ATTESTATION                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function toExecutionAttestation(bytes memory payload, address notary)
        public
        view
        returns (ExecutionAttestation memory)
    {
        return payload.castToAttestation().toExecutionAttestation(notary);
    }

    function isEmpty(ExecutionAttestation memory execAtt) public pure returns (bool) {
        return execAtt.isEmpty();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SUMMIT ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSummitAttestation(SummitAttestation memory summitAtt, uint32 nonce)
        public
        pure
        returns (bytes memory)
    {
        return summitAtt.formatSummitAttestation(nonce);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        ATTESTATION FORMATTERS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatAttestation(
        bytes32 snapRoot,
        bytes32 agentRoot,
        uint32 nonce,
        uint40 blockNumber,
        uint40 timestamp
    ) public pure returns (bytes memory) {
        return AttestationLib.formatAttestation(snapRoot, agentRoot, nonce, blockNumber, timestamp);
    }

    function isAttestation(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isAttestation();
    }
}
