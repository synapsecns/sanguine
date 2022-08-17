// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Header } from "../../../contracts/libs/Header.sol";

contract HeaderHarness {
    using Header for bytes29;

    function headerVersion() public pure returns (uint16) {
        return Header.HEADER_VERSION;
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

    function headerVersion(bytes29 _header) public pure returns (uint16) {
        return Header.headerVersion(_header);
    }

    function origin(bytes29 _header) public pure returns (uint32) {
        return Header.origin(_header);
    }

    function sender(bytes29 _header) public pure returns (bytes32) {
        return Header.sender(_header);
    }

    function nonce(bytes29 _header) public pure returns (uint32) {
        return Header.nonce(_header);
    }

    function destination(bytes29 _header) public pure returns (uint32) {
        return Header.destination(_header);
    }

    function recipient(bytes29 _header) public pure returns (bytes32) {
        return Header.recipient(_header);
    }

    function optimisticSeconds(bytes29 _header) public pure returns (uint32) {
        return Header.optimisticSeconds(_header);
    }
}
