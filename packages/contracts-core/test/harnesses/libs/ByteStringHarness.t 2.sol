// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    ByteString,
    CallData,
    Signature,
    TypedMemView
} from "../../../contracts/libs/ByteString.sol";

/**
 * @notice Exposes ByteString methods for testing against golang.
 */
contract ByteStringHarness {
    using ByteString for bytes;
    using ByteString for bytes29;
    using ByteString for CallData;
    using ByteString for Signature;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToRawBytes(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = ByteString.castToRawBytes(_payload);
        return _view.clone();
    }

    function castToSignature(bytes memory _payload) public view returns (bytes memory) {
        Signature _sig = _payload.castToSignature();
        return _sig.unwrap().clone();
    }

    function castToCallData(bytes memory _payload) public view returns (bytes memory) {
        CallData _callData = _payload.castToCallData();
        return _callData.unwrap().clone();
    }

    function arguments(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToCallData().arguments().clone();
    }

    function callSelector(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToCallData().callSelector().clone();
    }

    function argumentWords(bytes memory _payload) public pure returns (uint256) {
        return _payload.castToCallData().argumentWords();
    }

    function toRSV(bytes memory _payload)
        public
        pure
        returns (
            bytes32,
            bytes32,
            uint8
        )
    {
        return _payload.castToSignature().toRSV();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTING                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSignature(
        bytes32 r,
        bytes32 s,
        uint8 v
    ) public pure returns (bytes memory) {
        return ByteString.formatSignature({ r: r, s: s, v: v });
    }

    function isSignature(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isSignature();
    }

    function isCallData(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isCallData();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function signatureLength() public pure returns (uint256) {
        return ByteString.SIGNATURE_LENGTH;
    }

    function selectorLength() public pure returns (uint256) {
        return ByteString.SELECTOR_LENGTH;
    }
}
