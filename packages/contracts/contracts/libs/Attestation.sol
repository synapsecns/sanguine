// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

library Attestation {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev AttestationData memory layout
     * [000 .. 004): originDomain     uint32   4 bytes
     * [004 .. 008): nonce          uint32   4 bytes
     * [008 .. 040): root           bytes32 32 bytes
     *
     *      Attestation memory layout
     * [000 .. 040): data           bytes   40 bytes (see above)
     * [040 .. END): signature      bytes   ?? bytes (64/65 bytes)
     */

    uint256 internal constant OFFSET_ORIGIN_DOMAIN = 0;
    uint256 internal constant OFFSET_NONCE = 4;
    uint256 internal constant OFFSET_ROOT = 8;
    uint256 internal constant ATTESTATION_DATA_LENGTH = 40;
    uint256 internal constant OFFSET_SIGNATURE = ATTESTATION_DATA_LENGTH;

    /**
     * @notice Returns formatted Attestation with provided fields
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
     * @notice Returns formatted Attestation Data with provided fields
     * @param _domain   Domain of Origin's chain
     * @param _root     New merkle root
     * @param _nonce    Nonce of the merkle root
     * @return Formatted data
     **/
    function formatAttestationData(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_domain, _nonce, _root);
    }

    /**
     * @notice Checks that message is an Attestation, by checking its length
     */
    function isAttestation(bytes29 _view) internal pure returns (bool) {
        // Should have non-zero length for signature. Signature validity is not checked.
        return _view.len() > ATTESTATION_DATA_LENGTH;
    }

    /**
     * @notice Returns domain of chain where the Origin contract is deployed
     */
    function attestationDomain(bytes29 _view) internal pure returns (uint32) {
        return uint32(_view.indexUint(OFFSET_ORIGIN_DOMAIN, 4));
    }

    /**
     * @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
     */
    function attestationNonce(bytes29 _view) internal pure returns (uint32) {
        return uint32(_view.indexUint(OFFSET_NONCE, 4));
    }

    /**
     * @notice Returns a historical Merkle root from the Origin contract
     */
    function attestationRoot(bytes29 _view) internal pure returns (bytes32) {
        return _view.index(OFFSET_ROOT, 32);
    }

    /**
     * @notice Returns Attestation's Data, that is going to be signed by the Notary
     */
    function attestationData(bytes29 _view) internal pure returns (bytes29) {
        return _view.slice(OFFSET_ORIGIN_DOMAIN, ATTESTATION_DATA_LENGTH, 0);
    }

    /**
     * @notice Returns Notary's signature on AttestationData
     */
    function attestationSignature(bytes29 _view) internal pure returns (bytes29) {
        return _view.slice(OFFSET_SIGNATURE, _view.len() - ATTESTATION_DATA_LENGTH, 0);
    }
}
