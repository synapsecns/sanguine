// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Auth } from "../../../contracts/libs/Auth.sol";
import { ByteString } from "../../../contracts/libs/ByteString.sol";

/**
 * @notice Exposes Auth methods for testing against golang.
 */
contract AuthHarness {
    using ByteString for bytes;

    function toEthSignedMessageHash(bytes memory _data) external pure returns (bytes32) {
        return Auth.toEthSignedMessageHash(_data.castToRawBytes());
    }

    function recoverSigner(bytes32 _digest, bytes memory _signature)
        external
        pure
        returns (address)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        address signer = Auth.recoverSigner(_digest, _signature.castToSignature());
        return signer;
    }
}
