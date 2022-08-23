// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Header } from "../../../contracts/libs/Header.sol";

/**
 * @notice Exposes Header methods for testing against golang.
 */
contract HeaderHarness {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatHeader(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds
    ) public pure returns (bytes memory) {
        return
            Header.formatHeader(
                _origin,
                _sender,
                _nonce,
                _destination,
                _recipient,
                _optimisticSeconds
            );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function headerLength() public pure returns (uint256) {
        return Header.HEADER_LENGTH;
    }

    function headerVersion() public pure returns (uint16) {
        return Header.HEADER_VERSION;
    }

    function offsetVersion() public pure returns (uint256) {
        return Header.OFFSET_VERSION;
    }

    function offsetOrigin() public pure returns (uint256) {
        return Header.OFFSET_ORIGIN;
    }

    function offsetSender() public pure returns (uint256) {
        return Header.OFFSET_SENDER;
    }

    function offsetNonce() public pure returns (uint256) {
        return Header.OFFSET_NONCE;
    }

    function offsetDestination() public pure returns (uint256) {
        return Header.OFFSET_DESTINATION;
    }

    function offsetRecipient() public pure returns (uint256) {
        return Header.OFFSET_RECIPIENT;
    }

    function offsetOptimisticSeconds() public pure returns (uint256) {
        return Header.OFFSET_OPTIMISTIC_SECONDS;
    }
}
