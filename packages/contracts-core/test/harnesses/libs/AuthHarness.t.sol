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
        external
        pure
        returns (address)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        address signer = Auth.recoverSigner(_data.ref(0), _signature);
        return signer;
    }
}
