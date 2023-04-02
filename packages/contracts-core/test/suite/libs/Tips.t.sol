// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {TipsHarness} from "../../harnesses/libs/TipsHarness.t.sol";

import {TIPS_LENGTH} from "../../../contracts/libs/Tips.sol";

// solhint-disable func-name-mixedcase
contract TipsLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    TipsHarness internal libHarness;

    function setUp() public {
        libHarness = new TipsHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip) public {
        // TODO: Determine if we actually need uint96 for storing tips / totalTips
        uint256 totalTips = uint256(notaryTip) + broadcasterTip + proverTip + executorTip;
        vm.assume(totalTips <= type(uint96).max);
        // Test formatting
        bytes memory payload = libHarness.formatTips(notaryTip, broadcasterTip, proverTip, executorTip);
        assertEq(payload, abi.encodePacked(notaryTip, broadcasterTip, proverTip, executorTip), "!formatTips");
        // Test formatting checker
        checkCastToTips({payload: payload, isTips: true});
        // Test getters
        assertEq(libHarness.notaryTip(payload), notaryTip, "!notaryTip");
        assertEq(libHarness.broadcasterTip(payload), broadcasterTip, "!broadcasterTip");
        assertEq(libHarness.proverTip(payload), proverTip, "!proverTip");
        assertEq(libHarness.executorTip(payload), executorTip, "!executorTip");
        assertEq(libHarness.totalTips(payload), totalTips, "!totalTips");
    }

    function test_constants() public {
        // TODO: figure out why this doesn't mark offsetNotary as covered
        assertEq(libHarness.offsetNotary(), 0);
        assertEq(libHarness.offsetBroadcaster(), 12);
        // 12 + 12
        assertEq(libHarness.offsetProver(), 24);
        // 12 + 12 + 12
        assertEq(libHarness.offsetExecutor(), 36);
        // 12 + 12 + 12 + 12
        assertEq(libHarness.tipsLength(), 48);
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
