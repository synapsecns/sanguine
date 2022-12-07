// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseTypes } from "./SynapseTypes.sol";
import { TypedMemView } from "./TypedMemView.sol";
import { ByteString } from "./ByteString.sol";

library Attestation {
    using ByteString for bytes;

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
     * [044 .. 045): G = guardSigs  uint8    1 byte
     * [045 .. 046): N = notarySigs uint8    1 byte
     * [046 .. 111): guardSig[0]    bytes   65 bytes
     *      ..
     * [AAA .. BBB): guardSig[G-1]  bytes   65 bytes
     * [BBB .. CCC): notarySig[0]   bytes   65 bytes
     *      ..
     * [DDD .. END): notarySig[N-1] bytes   65 bytes
     */

    uint256 internal constant OFFSET_ORIGIN = 0;
    uint256 internal constant OFFSET_DESTINATION = 4;
    uint256 internal constant OFFSET_NONCE = 8;
    uint256 internal constant OFFSET_ROOT = 12;
    uint256 internal constant ATTESTATION_DATA_LENGTH = 44;

    uint256 internal constant OFFSET_AGENT_SIGS = ATTESTATION_DATA_LENGTH;
    uint256 internal constant OFFSET_FIRST_SIGNATURE = OFFSET_AGENT_SIGS + 2;

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
     * @notice Returns a formatted Attestation payload with provided fields
     * @param _data                 Attestation Data (see above)
     * @param _guardSignatures      Guard signatures on `_data`
     * @param _notarySignatures     Notary signatures on `_data`
     * @return Formatted attestation
     **/
    function formatAttestation(
        bytes memory _data,
        bytes[] memory _guardSignatures,
        bytes[] memory _notarySignatures
    ) internal view returns (bytes memory) {
        uint8 guardSigs = uint8(_guardSignatures.length);
        uint8 notarySigs = uint8(_notarySignatures.length);
        // Pack (guardSigs, notarySigs) into a single 16-byte value
        uint16 agentSigs = (uint16(guardSigs) << 8) | notarySigs;
        bytes29[] memory guardSigViews = ByteString.castToSignatures(_guardSignatures);
        bytes29[] memory notarySigViews = ByteString.castToSignatures(_notarySignatures);
        // We need to join: `_data`, `agentSigs`, `guardSigViews`, `notarySigViews`
        bytes29[] memory allViews = new bytes29[](2 + guardSigs + notarySigs);
        allViews[0] = _data.castToRawBytes();
        allViews[1] = abi.encodePacked(agentSigs).castToRawBytes();
        for (uint256 i = 0; i < guardSigs; ++i) {
            allViews[2 + i] = guardSigViews[i];
        }
        for (uint256 i = 0; i < notarySigs; ++i) {
            allViews[2 + guardSigs + i] = notarySigViews[i];
        }
        return TypedMemView.join(allViews);
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
     * @notice Returns a properly typed bytes29 pointer for an attestation payload.
     */
    function castToAttestation(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.ATTESTATION);
    }

    /**
     * @notice Checks that a payload is a formatted Attestation payload.
     */
    function isAttestation(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // (attData, guardSigs, notarySigs) need to exist
        if (length < OFFSET_FIRST_SIGNATURE) return false;
        (uint256 guardSigs, uint256 notarySigs) = _agentSignatures(_view);
        uint256 totalSigs = guardSigs + notarySigs;
        // There should be at least one signature
        if (totalSigs == 0) return false;
        // Every signature has length of exactly `ByteString.SIGNATURE_LENGTH`
        return length == OFFSET_FIRST_SIGNATURE + totalSigs * ByteString.SIGNATURE_LENGTH;
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
     * @notice Returns the amount of guard and notary signatures present in the Attestation.
     */
    function agentSignatures(bytes29 _view)
        internal
        pure
        onlyAttestation(_view)
        returns (uint8 guardSigs, uint8 notarySigs)
    {
        (guardSigs, notarySigs) = _agentSignatures(_view);
    }

    /**
     * @notice Returns the amount of guard signatures present in the Attestation.
     */
    function guardSignatures(bytes29 _view)
        internal
        pure
        onlyAttestation(_view)
        returns (uint8 guardSigs)
    {
        (guardSigs, ) = _agentSignatures(_view);
    }

    /**
     * @notice Returns the amount of notary signatures present in the Attestation.
     */
    function notarySignatures(bytes29 _view)
        internal
        pure
        onlyAttestation(_view)
        returns (uint8 notarySigs)
    {
        (, notarySigs) = _agentSignatures(_view);
    }

    /**
     * @notice Returns signature of the i-th Guard on AttestationData,
     * @dev Will revert if index is out of range.
     */
    function guardSignature(bytes29 _view, uint256 _guardIndex)
        internal
        pure
        onlyAttestation(_view)
        returns (bytes29)
    {
        (uint8 guardSigs, ) = _agentSignatures(_view);
        require(_guardIndex < guardSigs, "Out of range");
        return
            _view.slice({
                _index: OFFSET_FIRST_SIGNATURE + _guardIndex * ByteString.SIGNATURE_LENGTH,
                _len: ByteString.SIGNATURE_LENGTH,
                newType: SynapseTypes.SIGNATURE
            });
    }

    /**
     * @notice Returns signature of the i-th Notary on AttestationData,
     * @dev Will revert if index is out of range.
     */
    function notarySignature(bytes29 _view, uint256 _notaryIndex)
        internal
        pure
        onlyAttestation(_view)
        returns (bytes29)
    {
        (uint8 guardSigs, uint8 notarySigs) = _agentSignatures(_view);
        require(_notaryIndex < notarySigs, "Out of range");
        return
            _view.slice({
                _index: OFFSET_FIRST_SIGNATURE +
                    (_notaryIndex + guardSigs) *
                    ByteString.SIGNATURE_LENGTH,
                _len: ByteString.SIGNATURE_LENGTH,
                newType: SynapseTypes.SIGNATURE
            });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PRIVATE HELPERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Returns the amount of guard and notary signatures present in the Attestation.
     * Doesn't check the pointer type - to be used in functions that perform the typecheck.
     */
    function _agentSignatures(bytes29 _view)
        private
        pure
        returns (uint8 guardSigs, uint8 notarySigs)
    {
        // Read both amounts at once
        uint16 combinedAmounts = uint16(_view.indexUint({ _index: OFFSET_AGENT_SIGS, _bytes: 2 }));
        // First 8 bits is the amount of guard signatures
        guardSigs = uint8(combinedAmounts >> 8);
        // Last 8 bits is the amount of notary signatures
        notarySigs = uint8(combinedAmounts & 0xFF);
    }
}
