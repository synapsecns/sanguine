// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Auth } from "../../../contracts/libs/Auth.sol";
import { ByteString } from "../../../contracts/libs/ByteString.sol";

/**
 * @notice Exposes Auth methods for testing against golang.
 */
contract AuthHarness {
    using ByteString for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    function toEthSignedMessageHash(bytes memory data) external pure returns (bytes32) {
        return Auth.toEthSignedMessageHash(data.castToRawBytes());
    }

    function recoverSigner(bytes32 digest, bytes memory signature) external pure returns (address) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        address signer = Auth.recoverSigner(digest, signature.castToSignature());
        return signer;
    }
}
