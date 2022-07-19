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

    function test_checkSignature() public {
        bytes memory signature = signMessage(updaterPK, message);
        harness.checkSignature(updater, message, signature);
    }

    function test_checkSignature_wrongSigner() public {
        bytes memory signature = signMessage(fakeUpdaterPK, message);
        vm.expectRevert("Invalid signature");
        harness.checkSignature(updater, message, signature);
    }
}
