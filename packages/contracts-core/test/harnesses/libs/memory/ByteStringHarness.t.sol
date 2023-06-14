// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {ByteString, CallData, Signature, MemView, MemViewLib} from "../../../../contracts/libs/memory/ByteString.sol";

// solhint-disable ordering
/**
 * @notice Exposes ByteString methods for testing against golang.
 */
contract ByteStringHarness {
    using ByteString for bytes;
    using ByteString for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ═════════════════════════════════════════════════ SIGNATURE ═════════════════════════════════════════════════════

    function formatSignature(bytes32 r, bytes32 s, uint8 v) public pure returns (bytes memory) {
        return ByteString.formatSignature({r: r, s: s, v: v});
    }

    function castToSignature(bytes memory payload) public view returns (bytes memory) {
        Signature signature = payload.castToSignature();
        return signature.unwrap().clone();
    }

    function isSignature(bytes memory payload) public pure returns (bool) {
        return payload.ref().isSignature();
    }

    function toRSV(bytes memory payload) public pure returns (bytes32, bytes32, uint8) {
        return payload.castToSignature().toRSV();
    }

    // ═════════════════════════════════════════════════ CALLDATA ══════════════════════════════════════════════════════

    function addPrefix(bytes memory payload, bytes memory prefix) public view returns (bytes memory) {
        return payload.castToCallData().addPrefix(prefix);
    }

    function castToCallData(bytes memory payload) public view returns (bytes memory) {
        CallData callData = payload.castToCallData();
        return callData.unwrap().clone();
    }

    function isCallData(bytes memory payload) public pure returns (bool) {
        return payload.ref().isCallData();
    }

    function leaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToCallData().leaf();
    }

    function arguments(bytes memory payload) public view returns (bytes memory) {
        return payload.castToCallData().arguments().clone();
    }

    function callSelector(bytes memory payload) public pure returns (bytes4) {
        return payload.castToCallData().callSelector();
    }

    function argumentWords(bytes memory payload) public pure returns (uint256) {
        return payload.castToCallData().argumentWords();
    }

    // ═════════════════════════════════════════════ CONSTANT GETTERS ══════════════════════════════════════════════════

    function signatureLength() public pure returns (uint256) {
        return ByteString.SIGNATURE_LENGTH;
    }

    function selectorLength() public pure returns (uint256) {
        return ByteString.SELECTOR_LENGTH;
    }
}
