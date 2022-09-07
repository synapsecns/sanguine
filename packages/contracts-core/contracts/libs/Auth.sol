// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

library Auth {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @notice Recovers signer from data and signature.
     * @param _data         Data that was signed
     * @param _signature    `_data` signed by `signer`
     * @return signer       Address that signed the data
     */
    function recoverSigner(bytes29 _data, bytes memory _signature)
        internal
        pure
        returns (address signer)
    {
        bytes32 digest = _data.keccak();
        digest = ECDSA.toEthSignedMessageHash(digest);
        signer = ECDSA.recover(digest, _signature);
    }
}
