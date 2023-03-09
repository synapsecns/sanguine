// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    Attestation,
    AttestationLib,
    DestinationAttestation,
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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToAttestation(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Attestation _attestation = AttestationLib.castToAttestation(_payload);
        return _attestation.unwrap().clone();
    }

    function root(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToAttestation().root();
    }

    function height(bytes memory _payload) public pure returns (uint8) {
        return _payload.castToAttestation().height();
    }

    function nonce(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToAttestation().nonce();
    }

    function blockNumber(bytes memory _payload) public pure returns (uint40) {
        return _payload.castToAttestation().blockNumber();
    }

    function timestamp(bytes memory _payload) public pure returns (uint40) {
        return _payload.castToAttestation().timestamp();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       DESTINATION ATTESTATION                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function toDestinationAttestation(bytes memory _payload, address _notary)
        public
        view
        returns (DestinationAttestation memory)
    {
        return _payload.castToAttestation().toDestinationAttestation(_notary);
    }

    function isEmpty(DestinationAttestation memory _destAtt) public pure returns (bool) {
        return _destAtt.isEmpty();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SUMMIT ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSummitAttestation(SummitAttestation memory _summitAtt, uint32 _nonce)
        public
        pure
        returns (bytes memory)
    {
        return _summitAtt.formatSummitAttestation(_nonce);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        ATTESTATION FORMATTERS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatAttestation(
        bytes32 _root,
        uint8 _depth,
        uint32 _nonce,
        uint40 _blockNumber,
        uint40 _timestamp
    ) public pure returns (bytes memory) {
        return AttestationLib.formatAttestation(_root, _depth, _nonce, _blockNumber, _timestamp);
    }

    function isAttestation(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isAttestation();
    }
}
