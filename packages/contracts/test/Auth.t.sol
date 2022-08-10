// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { AuthHarness } from "./harnesses/AuthHarness.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract AuthTest is SynapseTest {
    using TypedMemView for bytes29;

    AuthHarness internal harness;
    bytes internal message = "Nothing to see here, please disperse";

    function setUp() public override {
        super.setUp();
        harness = new AuthHarness();
    }

    function test_recoverSigner() public {
        bytes memory signature = signMessage(updaterPK, message);
        assertEq(harness.recoverSigner(message, signature), updater);
    }
}
