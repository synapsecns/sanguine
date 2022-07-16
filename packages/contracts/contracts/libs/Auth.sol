// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

library Auth {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @notice Checks signer is authorized and that their signature is valid.
     * @param _signer       Who signed the message
     * @param _data         Data that was signed
     * @param _signature    `_data` signed by `_signer`, reverts if invalid
     */
    function checkSignature(
        address _signer,
        bytes29 _data,
        bytes memory _signature
    ) internal pure {
        bytes32 digest = _data.keccak();
        digest = ECDSA.toEthSignedMessageHash(digest);
        require((ECDSA.recover(digest, _signature) == _signer), "Invalid signature");
    }
}
