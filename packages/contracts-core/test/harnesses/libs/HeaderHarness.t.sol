// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {Header, HeaderLib, TypedMemView} from "../../../contracts/libs/Header.sol";

/**
 * @notice Exposes Header methods for testing against golang.
 */
contract HeaderHarness {
    using HeaderLib for bytes;
    using HeaderLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToHeader(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Header header = HeaderLib.castToHeader(payload);
        return header.unwrap().clone();
    }

    /// @notice Returns header's version field.
    function version(bytes memory payload) public pure returns (uint16) {
        return payload.castToHeader().version();
    }

    /// @notice Returns header's origin field
    function origin(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().origin();
    }

    /// @notice Returns header's sender field
    function sender(bytes memory payload) public pure returns (bytes32) {
        return payload.castToHeader().sender();
    }

    /// @notice Returns header's nonce field
    function nonce(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().nonce();
    }

    /// @notice Returns header's destination field
    function destination(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().destination();
    }

    /// @notice Returns header's recipient field as bytes32
    function recipient(bytes memory payload) public pure returns (bytes32) {
        return payload.castToHeader().recipient();
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticSeconds(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().optimisticSeconds();
    }

    /// @notice Returns header's recipient field as an address
    function recipientAddress(bytes memory payload) public pure returns (address) {
        return payload.castToHeader().recipientAddress();
    }

    function isHeader(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isHeader();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatHeader(
        uint32 origin_,
        bytes32 sender_,
        uint32 nonce_,
        uint32 destination_,
        bytes32 recipient_,
        uint32 optimisticSeconds_
    ) public pure returns (bytes memory) {
        return HeaderLib.formatHeader(origin_, sender_, nonce_, destination_, recipient_, optimisticSeconds_);
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
