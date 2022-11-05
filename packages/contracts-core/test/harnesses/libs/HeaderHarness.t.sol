// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Header } from "../../../contracts/libs/Header.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes Header methods for testing against golang.
 */
contract HeaderHarness {
    using Header for bytes;
    using Header for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToHeader(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = Header.castToHeader(_payload);
        return (_view.typeOf(), _view.clone());
    }

    /// @notice Returns header's version field.
    function headerVersion(uint40 _type, bytes memory _payload) public pure returns (uint16) {
        return _payload.ref(_type).headerVersion();
    }

    /// @notice Returns header's origin field
    function origin(uint40 _type, bytes memory _payload) public pure returns (uint32) {
        return _payload.ref(_type).origin();
    }

    /// @notice Returns header's sender field
    function sender(uint40 _type, bytes memory _payload) public pure returns (bytes32) {
        return _payload.ref(_type).sender();
    }

    /// @notice Returns header's nonce field
    function nonce(uint40 _type, bytes memory _payload) public pure returns (uint32) {
        return _payload.ref(_type).nonce();
    }

    /// @notice Returns header's destination field
    function destination(uint40 _type, bytes memory _payload) public pure returns (uint32) {
        return _payload.ref(_type).destination();
    }

    /// @notice Returns header's recipient field as bytes32
    function recipient(uint40 _type, bytes memory _payload) public pure returns (bytes32) {
        return _payload.ref(_type).recipient();
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticSeconds(uint40 _type, bytes memory _payload) public pure returns (uint32) {
        return _payload.ref(_type).optimisticSeconds();
    }

    /// @notice Returns header's recipient field as an address
    function recipientAddress(uint40 _type, bytes memory _payload) public pure returns (address) {
        return _payload.ref(_type).recipientAddress();
    }

    function isHeader(bytes memory _payload) public pure returns (bool) {
        return _payload.castToHeader().isHeader();
    }

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
