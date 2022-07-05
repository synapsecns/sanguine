// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

library Auth {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @notice Checks signer is authorized and that their signature is valid.
     * Returns a view over the passed payload for later slicing.
     * @param _signer       Who signed the message
     * @param _payload      Message to be signed
     * @param _signature    `_payload` signed by `_signer`, reverts if invalid
     */
    function checkSignature(
        address _signer,
        bytes memory _payload,
        bytes memory _signature
    ) internal pure returns (bytes29 _view) {
        _view = _payload.ref(0);
        bytes32 digest = _view.keccak();
        digest = ECDSA.toEthSignedMessageHash(digest);
        require((ECDSA.recover(digest, _signature) == _signer), "Invalid signature");
    }
}
