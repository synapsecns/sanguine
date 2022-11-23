// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseTypes } from "./SynapseTypes.sol";
import { TypedMemView } from "./TypedMemView.sol";
import { ByteString } from "./ByteString.sol";

library Attestation {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev AttestationData memory layout
     * [000 .. 004): origin         uint32   4 bytes
     * [004 .. 008): destination    uint32   4 bytes
     * [008 .. 012): nonce          uint32   4 bytes
     * [012 .. 044): root           bytes32 32 bytes
     *
     *      Attestation memory layout
     * [000 .. 044): attData        bytes   44 bytes (see above)
     * [044 .. 109): signature      bytes   65 bytes (65 bytes)
     */

    uint256 internal constant OFFSET_ORIGIN = 0;
    uint256 internal constant OFFSET_DESTINATION = 4;
    uint256 internal constant OFFSET_NONCE = 8;
    uint256 internal constant OFFSET_ROOT = 12;
    uint256 internal constant ATTESTATION_DATA_LENGTH = 44;
    uint256 internal constant OFFSET_SIGNATURE = ATTESTATION_DATA_LENGTH;
    uint256 internal constant ATTESTATION_LENGTH = ATTESTATION_DATA_LENGTH + 65;

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
     * @param _origin       Domain of Origin's chain
     * @param _destination  Domain of Destination's chain
     * @param _root         New merkle root
     * @param _nonce        Nonce of the merkle root
     * @return Formatted attestation data
     **/
    function formatAttestationData(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_origin, _destination, _nonce, _root);
    }

    /**
     * @notice Checks that a payload is a formatted Attestation payload.
     */
    function isAttestation(bytes29 _view) internal pure returns (bool) {
        return _view.len() == ATTESTATION_LENGTH;
    }

    /**
     * @notice Combines origin and destination domains into `attestationDomains`,
     * a unique ID for every (origin, destination) pair. Could be used to identify
     * Merkle trees on Origin, or Mirrors on Destination.
     */
    function attestationDomains(uint32 _origin, uint32 _destination)
        internal
        pure
        returns (uint64)
    {
        return (uint64(_origin) << 32) | _destination;
    }

    /**
     * @notice Combines origin, destination domains and message nonce into `attestationKey`,
     * a unique key for every (origin, destination, nonce) tuple. Could be used to identify
     * any dispatched message.
     */
    function attestationKey(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce
    ) internal pure returns (uint96) {
        return (uint96(_origin) << 64) | (uint96(_destination) << 32) | _nonce;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns domain of chain where the Origin contract is deployed
     */
    function attestedOrigin(bytes29 _view) internal pure onlyAttestation(_view) returns (uint32) {
        return uint32(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 4 }));
    }

    /**
     * @notice Returns domain of chain where the Destination contract is deployed
     */
    function attestedDestination(bytes29 _view)
        internal
        pure
        onlyAttestation(_view)
        returns (uint32)
    {
        return uint32(_view.indexUint({ _index: OFFSET_DESTINATION, _bytes: 4 }));
    }

    /**
     * @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
     */
    function attestedNonce(bytes29 _view) internal pure onlyAttestation(_view) returns (uint32) {
        return uint32(_view.indexUint({ _index: OFFSET_NONCE, _bytes: 4 }));
    }

    /**
     * @notice Returns a combined field for (origin, destination). See `attestationDomains()`.
     */
    function attestedDomains(bytes29 _view) internal pure onlyAttestation(_view) returns (uint64) {
        return uint64(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 8 }));
    }

    /**
     * @notice Returns a combined field for (origin, destination, nonce). See `attestationKey()`.
     */
    function attestedKey(bytes29 _view) internal pure onlyAttestation(_view) returns (uint96) {
        return uint96(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 12 }));
    }

    /**
     * @notice Returns a historical Merkle root from the Origin contract
     */
    function attestedRoot(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes32) {
        return _view.index({ _index: OFFSET_ROOT, _bytes: 32 });
    }

    /**
     * @notice Returns Attestation's Data, that is going to be signed by the Notary
     */
    function attestationData(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes29) {
        return
            _view.slice({
                _index: OFFSET_ORIGIN,
                _len: ATTESTATION_DATA_LENGTH,
                newType: SynapseTypes.ATTESTATION_DATA
            });
    }

    /**
     * @notice Returns Notary's signature on AttestationData
     */
    function notarySignature(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes29) {
        return
            _view.slice({
                _index: OFFSET_SIGNATURE,
                _len: ByteString.SIGNATURE_LENGTH,
                newType: SynapseTypes.SIGNATURE
            });
    }
}
