// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/Attestation.sol";

/**
 * @notice Exposes Attestation methods for testing against golang.
 */
contract AttestationHarness {
    using AttestationLib for bytes;
    using AttestationLib for bytes29;
    using AttestationLib for Attestation;
    using AttestationLib for AttestationData;
    using ByteString for Signature;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        ATTESTATION FORMATTERS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatAttestation(
        bytes memory _data,
        bytes memory _guardSignatures,
        bytes memory _notarySignatures
    ) public view returns (bytes memory) {
        return AttestationLib.formatAttestation(_data, _guardSignatures, _notarySignatures);
    }

    function formatAttestationFromViews(
        bytes memory _data,
        bytes memory _guardSignatures,
        bytes memory _notarySignatures
    ) public view returns (bytes memory) {
        return
            AttestationLib.formatAttestation({
                _data: _data.castToAttestationData(),
                _guardSigs: _guardSignatures.ref(0),
                _notarySigs: _notarySignatures.ref(0)
            });
    }

    function castToAttestation(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Attestation att = AttestationLib.castToAttestation(_payload);
        return att.unwrap().clone();
    }

    function isAttestation(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isAttestation();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION GETTERS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line ordering
    function data(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToAttestation().data().unwrap().clone();
    }

    function guardSignature(bytes memory _payload, uint256 _index)
        public
        view
        returns (bytes memory)
    {
        return _payload.castToAttestation().guardSignature(_index).unwrap().clone();
    }

    function notarySignature(bytes memory _payload, uint256 _index)
        public
        view
        returns (bytes memory)
    {
        return _payload.castToAttestation().notarySignature(_index).unwrap().clone();
    }

    function agentsAmount(bytes memory _payload) public pure returns (uint8, uint8) {
        return _payload.castToAttestation().agentsAmount();
    }

    function guardsAmount(bytes memory _payload) public pure returns (uint8) {
        return _payload.castToAttestation().guardsAmount();
    }

    function notariesAmount(bytes memory _payload) public pure returns (uint8) {
        return _payload.castToAttestation().notariesAmount();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     ATTESTATION DATA FORMATTERS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatAttestationData(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root,
        uint40 _blockNumber,
        uint40 _timestamp
    ) public pure returns (bytes memory) {
        return
            AttestationLib.formatAttestationData(
                _origin,
                _destination,
                _nonce,
                _root,
                _blockNumber,
                _timestamp
            );
    }

    function castToAttestationData(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToAttestationData().unwrap().clone();
    }

    function isAttestationData(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isAttestationData();
    }

    function packDomains(uint32 _origin, uint32 _destination) public pure returns (uint64) {
        return AttestationLib.packDomains(_origin, _destination);
    }

    function packKey(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce
    ) public pure returns (uint96) {
        return AttestationLib.packKey(_origin, _destination, _nonce);
    }

    function unpackDomains(uint64 _attestationDomains) public pure returns (uint32, uint32) {
        return AttestationLib.unpackDomains(_attestationDomains);
    }

    function unpackKey(uint96 _attestationKey) public pure returns (uint64, uint32) {
        return AttestationLib.unpackKey(_attestationKey);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       ATTESTATION DATA GETTERS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function origin(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToAttestationData().origin();
    }

    function destination(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToAttestationData().destination();
    }

    function nonce(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToAttestationData().nonce();
    }

    function domains(bytes memory _payload) public pure returns (uint64) {
        return _payload.castToAttestationData().domains();
    }

    function key(bytes memory _payload) public pure returns (uint96) {
        return _payload.castToAttestationData().key();
    }

    function root(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToAttestationData().root();
    }

    function blockNumber(bytes memory _payload) public pure returns (uint40) {
        return _payload.castToAttestationData().blockNumber();
    }

    function timestamp(bytes memory _payload) public pure returns (uint40) {
        return _payload.castToAttestationData().timestamp();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function attestationDataLength() public pure returns (uint256) {
        return AttestationLib.ATTESTATION_DATA_LENGTH;
    }

    function offsetOrigin() public pure returns (uint256) {
        return AttestationLib.OFFSET_ORIGIN;
    }

    function offsetDestination() public pure returns (uint256) {
        return AttestationLib.OFFSET_DESTINATION;
    }

    function offsetNonce() public pure returns (uint256) {
        return AttestationLib.OFFSET_NONCE;
    }

    function offsetRoot() public pure returns (uint256) {
        return AttestationLib.OFFSET_ROOT;
    }

    function offsetBlockNumber() public pure returns (uint256) {
        return AttestationLib.OFFSET_BLOCK_NUMBER;
    }

    function offsetTimestamp() public pure returns (uint256) {
        return AttestationLib.OFFSET_TIMESTAMP;
    }

    function offsetAgentSignatures() public pure returns (uint256) {
        return AttestationLib.OFFSET_AGENT_SIGS;
    }

    function offsetFirstSignature() public pure returns (uint256) {
        return AttestationLib.OFFSET_FIRST_SIGNATURE;
    }
}
