// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { ByteString } from "../../../contracts/libs/ByteString.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes ByteString methods for testing against golang.
 */
contract ByteStringHarness {
    using ByteString for bytes;
    using ByteString for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToRawBytes(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = ByteString.castToRawBytes(_payload);
        return (_view.typeOf(), _view.clone());
    }

    function castToSignature(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.castToSignature();
        return (_view.typeOf(), _view.clone());
    }

    function castToCallPayload(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.castToCallPayload();
        return (_view.typeOf(), _view.clone());
    }

    function argumentsPayload(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).argumentsPayload();
        return (_view.typeOf(), _view.clone());
    }

    function callSelector(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).callSelector();
        return (_view.typeOf(), _view.clone());
    }

    function toRSV(uint40 _type, bytes memory _payload)
        public
        pure
        returns (
            bytes32,
            bytes32,
            uint8
        )
    {
        return _payload.ref(_type).toRSV();
    }

    function argumentWords(uint40 _type, bytes memory _payload) public pure returns (uint256) {
        return _payload.ref(_type).argumentWords();
    }

    function isSignature(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isSignature();
    }

    function isCallPayload(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isCallPayload();
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
