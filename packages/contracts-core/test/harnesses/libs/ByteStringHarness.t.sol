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

    function arguments(uint40 _type, bytes memory _payload) public pure returns (uint256) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        return ByteString.argumentWords(_payload.ref(_type));
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
