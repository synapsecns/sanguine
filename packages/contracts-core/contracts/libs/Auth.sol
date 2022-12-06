// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ByteString } from "../libs/ByteString.sol";
import { TypedMemView } from "./TypedMemView.sol";

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

library Auth {
    using ByteString for bytes29;
    using TypedMemView for bytes29;

    /**
     * @notice Returns an Ethereum Signed Message, created from a `_view`.
     * @dev This produces hash corresponding to the one signed with the
     * https://eth.wiki/json-rpc/API#eth_sign[`eth_sign`]
     * JSON-RPC method as part of EIP-191.
     * See {recoverSigner}.
     * @param _dataView Memory view over the data that needs to be signed
     * @return digest   An Ethereum Signed Message for the given data
     */
    function toEthSignedMessageHash(bytes29 _dataView) internal pure returns (bytes32 digest) {
        // Derive hash of the original data and use that for forming an Ethereum Signed Message
        digest = ECDSA.toEthSignedMessageHash(_dataView.keccak());
    }

    /**
     * @notice Recovers signer from digest and signature.
     * @dev IMPORTANT: `_digest` _must_ be the result of a hash operation for the
     * verification to be secure: it is possible to craft signatures that
     * recover to arbitrary addresses for non-hashed data. A safe way to ensure
     * this is by receiving a hash of the original message (which may otherwise
     * be too long), and then calling {toEthSignedMessageHash} on it.
     * @param _digest           Digest that was signed
     * @param _signatureView    Memory view over `signer` signature on `_digest`
     * @return signer           Address that signed the data
     */
    function recoverSigner(bytes32 _digest, bytes29 _signatureView)
        internal
        pure
        returns (address signer)
    {
        require(_signatureView.isSignature(), "Not a signature");
        (bytes32 r, bytes32 s, uint8 v) = _signatureView.toRSV();
        signer = ECDSA.recover({ hash: _digest, r: r, s: s, v: v });
    }
}
