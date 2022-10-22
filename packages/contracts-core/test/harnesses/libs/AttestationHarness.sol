// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Attestation } from "../../../contracts/libs/Attestation.sol";

/**
 * @notice Exposes Attestation methods for testing against golang.
 */
contract AttestationHarness {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root,
        bytes memory _signature
    ) public pure returns (bytes memory) {
        return formatAttestation(formatAttestationData(_domain, _nonce, _root), _signature);
    }

    function formatAttestation(bytes memory _data, bytes memory _signature)
        public
        pure
        returns (bytes memory)
    {
        return Attestation.formatAttestation(_data, _signature);
    }

    function formatAttestationData(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) public pure returns (bytes memory) {
        return Attestation.formatAttestationData(_domain, _nonce, _root);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function attestationDataLength() public pure returns (uint256) {
        return Attestation.ATTESTATION_DATA_LENGTH;
    }

    function offsetOrigin() public pure returns (uint256) {
        return Attestation.OFFSET_ORIGIN_DOMAIN;
    }

    function offsetNonce() public pure returns (uint256) {
        return Attestation.OFFSET_NONCE;
    }

    function offsetRoot() public pure returns (uint256) {
        return Attestation.OFFSET_ROOT;
    }

    function offsetSignature() public pure returns (uint256) {
        return Attestation.OFFSET_SIGNATURE;
    }
}
