// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {TipsHarness} from "../../harnesses/libs/TipsHarness.t.sol";

import {TIPS_MULTIPLIER, TIPS_LENGTH} from "../../../contracts/libs/Constants.sol";

// solhint-disable func-name-mixedcase
contract TipsLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    TipsHarness internal libHarness;

    function setUp() public {
        libHarness = new TipsHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatTips(uint64 summitTip, uint64 attestationTip, uint64 executionTip, uint64 deliveryTip) public {
        uint256 totalTips = uint256(summitTip) + attestationTip + executionTip + deliveryTip;
        vm.assume(totalTips <= type(uint64).max);
        // Test formatting
        bytes memory payload = libHarness.formatTips(summitTip, attestationTip, executionTip, deliveryTip);
        assertEq(payload, abi.encodePacked(summitTip, attestationTip, executionTip, deliveryTip), "!formatTips");
        // Test formatting checker
        checkCastToTips({payload: payload, isTips: true});
        // Test getters
        assertEq(libHarness.summitTip(payload), summitTip, "!summitTip");
        assertEq(libHarness.attestationTip(payload), attestationTip, "!attestationTip");
        assertEq(libHarness.executionTip(payload), executionTip, "!executionTip");
        assertEq(libHarness.deliveryTip(payload), deliveryTip, "!deliveryTip");
        assertEq(libHarness.value(payload), totalTips * TIPS_MULTIPLIER, "!totalTips");
    }

    function test_isTips(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToTips({payload: payload, isTips: length == TIPS_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToTips(bytes memory payload, bool isTips) public {
        if (isTips) {
            assertTrue(libHarness.isTips(payload), "!isTips: when valid");
            assertEq(libHarness.castToTips(payload), payload, "!castToTips: when valid");
        } else {
            assertFalse(libHarness.isTips(payload), "!isTips: when valid");
            vm.expectRevert("Not a tips payload");
            libHarness.castToTips(payload);
        }
    }
}
