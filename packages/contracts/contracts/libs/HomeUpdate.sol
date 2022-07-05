// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

library HomeUpdate {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev HomeUpdate memory layout
     * [000 .. 004): homeDomain     uint32   4 bytes
     * [004 .. 008): nonce           uint32   4 bytes
     * [008 .. 040): root          bytes32 32 bytes
     */

    uint256 internal constant OFFSET_HOME_DOMAIN = 0;
    uint256 internal constant OFFSET_NONCE = 4;
    uint256 internal constant OFFSET_ROOT = 8;
    uint256 internal constant HOME_UPDATE_LENGTH = 40;

    /**
     * @notice Returns formatted HomeUpdate message with provided fields
     * @param _domain   Domain of Home's chain
     * @param _root     New merkle root
     * @param _nonce    Nonce of the merkle root
     * @return Formatted message
     **/
    function formatHomeUpdate(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_domain, _nonce, _root);
    }

    /**
     * @notice Checks that message is a valid HopeUpdate, by checking its length
     */
    function isValidUpdate(bytes29 _homeUpdate) internal pure returns (bool) {
        return _homeUpdate.len() == HOME_UPDATE_LENGTH;
    }

    /**
     * @notice Returns domain of chain where the Home contract is deployed
     */
    function homeDomain(bytes29 _homeUpdate) internal pure returns (uint32) {
        return uint32(_homeUpdate.indexUint(OFFSET_HOME_DOMAIN, 4));
    }

    /**
     * @notice Returns nonce of Home contract at the time, when `root` was the Merkle root.
     */
    function nonce(bytes29 _homeUpdate) internal pure returns (uint32) {
        return uint32(_homeUpdate.indexUint(OFFSET_NONCE, 4));
    }

    /**
     * @notice Returns a historical Merkle root from the Home contract
     */
    function root(bytes29 _homeUpdate) internal pure returns (bytes32) {
        return _homeUpdate.index(OFFSET_ROOT, 32);
    }
}
