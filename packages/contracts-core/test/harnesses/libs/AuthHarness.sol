// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Auth } from "../../../contracts/libs/Auth.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes Auth methods for testing against golang.
 */
contract AuthHarness {
    using TypedMemView for bytes;

    function recoverSigner(bytes memory _data, bytes memory _signature)
        public
        pure
        returns (address)
    {
        return Auth.recoverSigner(_data.ref(0), _signature);
    }
}
