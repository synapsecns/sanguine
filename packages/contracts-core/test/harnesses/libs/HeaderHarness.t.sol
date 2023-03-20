// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Header, HeaderLib, TypedMemView } from "../../../contracts/libs/Header.sol";

/**
 * @notice Exposes Header methods for testing against golang.
 */
contract HeaderHarness {
    using HeaderLib for bytes;
    using HeaderLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToHeader(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Header _header = HeaderLib.castToHeader(_payload);
        return _header.unwrap().clone();
    }

    /// @notice Returns header's version field.
    function version(bytes memory _payload) public pure returns (uint16) {
        return _payload.castToHeader().version();
    }

    /// @notice Returns header's origin field
    function origin(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToHeader().origin();
    }

    /// @notice Returns header's sender field
    function sender(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToHeader().sender();
    }

    /// @notice Returns header's nonce field
    function nonce(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToHeader().nonce();
    }

    /// @notice Returns header's destination field
    function destination(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToHeader().destination();
    }

    /// @notice Returns header's recipient field as bytes32
    function recipient(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToHeader().recipient();
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticSeconds(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToHeader().optimisticSeconds();
    }

    /// @notice Returns header's recipient field as an address
    function recipientAddress(bytes memory _payload) public pure returns (address) {
        return _payload.castToHeader().recipientAddress();
    }

    function isHeader(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isHeader();
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
            HeaderLib.formatHeader(
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
        return HeaderLib.HEADER_LENGTH;
    }

    function headerVersion() public pure returns (uint16) {
        return HeaderLib.HEADER_VERSION;
    }

    function offsetVersion() public pure returns (uint256) {
        return HeaderLib.OFFSET_VERSION;
    }

    function offsetOrigin() public pure returns (uint256) {
        return HeaderLib.OFFSET_ORIGIN;
    }

    function offsetSender() public pure returns (uint256) {
        return HeaderLib.OFFSET_SENDER;
    }

    function offsetNonce() public pure returns (uint256) {
        return HeaderLib.OFFSET_NONCE;
    }

    function offsetDestination() public pure returns (uint256) {
        return HeaderLib.OFFSET_DESTINATION;
    }

    function offsetRecipient() public pure returns (uint256) {
        return HeaderLib.OFFSET_RECIPIENT;
    }

    function offsetOptimisticSeconds() public pure returns (uint256) {
        return HeaderLib.OFFSET_OPTIMISTIC_SECONDS;
    }
}
