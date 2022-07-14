// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

library Attestation {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev Attestation memory layout
     * [000 .. 004): homeDomain     uint32   4 bytes
     * [004 .. 008): nonce          uint32   4 bytes
     * [008 .. 040): root           bytes32 32 bytes
     */

    uint256 internal constant OFFSET_HOME_DOMAIN = 0;
    uint256 internal constant OFFSET_NONCE = 4;
    uint256 internal constant OFFSET_ROOT = 8;
    uint256 internal constant ATTESTATION_LENGTH = 40;

    /**
     * @notice Returns formatted Attestation message with provided fields
     * @param _domain   Domain of Home's chain
     * @param _root     New merkle root
     * @param _nonce    Nonce of the merkle root
     * @return Formatted message
     **/
    function formatAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_domain, _nonce, _root);
    }

    /**
     * @notice Checks that message is a valid Attestation, by checking its length
     */
    function isValidAttestation(bytes29 _view) internal pure returns (bool) {
        return _view.len() == ATTESTATION_LENGTH;
    }

    /**
     * @notice Returns domain of chain where the Home contract is deployed
     */
    function attestationDomain(bytes29 _view) internal pure returns (uint32) {
        return uint32(_view.indexUint(OFFSET_HOME_DOMAIN, 4));
    }

    /**
     * @notice Returns nonce of Home contract at the time, when `root` was the Merkle root.
     */
    function attestationNonce(bytes29 _view) internal pure returns (uint32) {
        return uint32(_view.indexUint(OFFSET_NONCE, 4));
    }

    /**
     * @notice Returns a historical Merkle root from the Home contract
     */
    function attestationRoot(bytes29 _view) internal pure returns (bytes32) {
        return _view.index(OFFSET_ROOT, 32);
    }
}
