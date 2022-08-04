// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";
import { Auth } from "../../contracts/libs/Auth.sol";

contract AuthHarness {
    using TypedMemView for bytes;

    function recoverSigner(bytes memory _payload, bytes memory _signature)
        public
        pure
        returns (address)
    {
        return Auth.recoverSigner(_payload.ref(0), _signature);
    }
}
