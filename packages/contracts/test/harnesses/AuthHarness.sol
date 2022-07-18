// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";
import { Auth } from "../../contracts/libs/Auth.sol";

contract AuthHarness {
    using TypedMemView for bytes;

    function checkSignature(
        address _signer,
        bytes memory _payload,
        bytes memory _signature
    ) public pure {
        Auth.checkSignature(_signer, _payload.ref(0), _signature);
    }
}
