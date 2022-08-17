// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseTest } from "../utils/SynapseTest.sol";

import { Auth } from "../../contracts/libs/Auth.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase

contract AuthTest is SynapseTest {
    using TypedMemView for bytes;

    bytes internal message = "Nothing to see here, please disperse";

    function test_recoverSigner() public {
        bytes memory signature = signMessage(notaryPK, message);
        assertEq(Auth.recoverSigner(message.ref(0), signature), notary);
    }
}
