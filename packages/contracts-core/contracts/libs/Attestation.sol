// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseTypes } from "./SynapseTypes.sol";
import { TypedMemView } from "./TypedMemView.sol";
import { Auth } from "./Auth.sol";

library Attestation {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev AttestationData memory layout
     * [000 .. 004): origin         uint32   4 bytes
     * [000 .. 008): destination    uint32   4 bytes
     * [008 .. 012): nonce          uint32   4 bytes
     * [012 .. 044): root           bytes32 32 bytes
     *
     *      Attestation memory layout
     * [000 .. 044): data           bytes   44 bytes (see above)
     * [044 .. 109): signature      bytes   65 bytes (65 bytes)
     */

    uint256 internal constant OFFSET_ORIGIN_DOMAIN = 0;
    uint256 internal constant OFFSET_DESTINATION_DOMAIN = 4;
    uint256 internal constant OFFSET_NONCE = 8;
    uint256 internal constant OFFSET_ROOT = 12;
    uint256 internal constant ATTESTATION_DATA_LENGTH = 44;
    uint256 internal constant OFFSET_SIGNATURE = ATTESTATION_DATA_LENGTH;
    uint256 internal constant ATTESTATION_LENGTH = 109;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyAttestation(bytes29 _view) {
        _view.assertType(SynapseTypes.ATTESTATION);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for an attestation payload.
     */
    function castToAttestation(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.ATTESTATION);
    }

    /**
     * @notice Returns a formatted Attestation payload with provided fields
     * @param _data         Attestation Data (see above)
     * @param _signature    Notary's signature on `_data`
     * @return Formatted attestation
     **/
    function formatAttestation(bytes memory _data, bytes memory _signature)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_data, _signature);
    }

    /**
     * @notice Returns a formatted AttestationData payload with provided fields
     * @param _origin_domain        Domain of Origin's chain
     * @param _destination_domain   Domain of Destination's chain
     * @param _root                 New merkle root
     * @param _nonce                Nonce of the merkle root
     * @return Formatted attestation data
     **/
    function formatAttestationData(
        uint32 _origin_domain,
        uint32 _destination_domain,
        uint32 _nonce,
        bytes32 _root
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_origin_domain, _destination_domain, _nonce, _root);
    }

    /**
     * @notice Checks that a payload is a formatted Attestation payload.
     */
    function isAttestation(bytes29 _view) internal pure returns (bool) {
        return _view.len() == ATTESTATION_LENGTH;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns domain of chain where the Origin contract is deployed
     */
    function attestedOriginDomain(bytes29 _view)
        internal
        pure
        onlyAttestation(_view)
        returns (uint32)
    {
        return uint32(_view.indexUint(OFFSET_ORIGIN_DOMAIN, 4));
    }

    /**
     * @notice Returns domain of chain where the Destination contract is deployed
     */
    function attestedDestinationDomain(bytes29 _view)
        internal
        pure
        onlyAttestation(_view)
        returns (uint32)
    {
        return uint32(_view.indexUint(OFFSET_DESTINATION_DOMAIN, 4));
    }

    /**
     * @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
     */
    function attestedNonce(bytes29 _view) internal pure onlyAttestation(_view) returns (uint32) {
        return uint32(_view.indexUint(OFFSET_NONCE, 4));
    }

    /**
     * @notice Returns a historical Merkle root from the Origin contract
     */
    function attestedRoot(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes32) {
        return _view.index(OFFSET_ROOT, 32);
    }

    /**
     * @notice Returns Attestation's Data, that is going to be signed by the Notary
     */
    function attestationData(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes29) {
        return
            _view.slice(
                OFFSET_ORIGIN_DOMAIN,
                ATTESTATION_DATA_LENGTH,
                SynapseTypes.ATTESTATION_DATA
            );
    }

    /**
     * @notice Returns Notary's signature on AttestationData
     */
    function notarySignature(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes29) {
        return _view.slice(OFFSET_SIGNATURE, Auth.SIGNATURE_LENGTH, SynapseTypes.SIGNATURE);
    }
}
