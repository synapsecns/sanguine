// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {Tips, TipsHarness} from "../../../harnesses/libs/stack/TipsHarness.t.sol";

import {TipsOverflow, TipsValueTooLow} from "../../../../contracts/libs/Errors.sol";
import {TIPS_MULTIPLIER, TIPS_LENGTH} from "../../../../contracts/libs/Constants.sol";

import {RawTips} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract TipsLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    TipsHarness internal libHarness;

    function setUp() public {
        libHarness = new TipsHarness();
    }

    function test_encodeTips(RawTips memory rt) public {
        uint256 totalTips = uint256(rt.summitTip) + rt.attestationTip + rt.executionTip + rt.deliveryTip;
        vm.assume(totalTips <= type(uint64).max);
        // Test encoding
        uint256 encodedTips = libHarness.encodeTips(rt.summitTip, rt.attestationTip, rt.executionTip, rt.deliveryTip);
        uint256 expected = uint256(rt.summitTip) * 2 ** 192 + uint256(rt.attestationTip) * 2 ** 128
            + uint256(rt.executionTip) * 2 ** 64 + uint256(rt.deliveryTip);
        assertEq(encodedTips, expected, "!encodeTips");
        assertEq(libHarness.wrapPadded(encodedTips), expected, "!wrapPadded");
        // Test getters
        assertEq(libHarness.summitTip(encodedTips), rt.summitTip, "!summitTip");
        assertEq(libHarness.attestationTip(encodedTips), rt.attestationTip, "!attestationTip");
        assertEq(libHarness.executionTip(encodedTips), rt.executionTip, "!executionTip");
        assertEq(libHarness.deliveryTip(encodedTips), rt.deliveryTip, "!deliveryTip");
        assertEq(libHarness.value(encodedTips), totalTips * TIPS_MULTIPLIER, "!totalTips");
        // Test hashing
        assertEq(libHarness.leaf(encodedTips), keccak256(abi.encode(expected)), "!leaf");
    }

    function test_encodeTips256(uint96 summitTip, uint96 attestationTip, uint96 executionTip, uint96 deliveryTip)
        public
    {
        uint256 encodedTips256 = libHarness.encodeTips256(summitTip, attestationTip, executionTip, deliveryTip);
        uint256 encodedTips = libHarness.encodeTips(
            uint64(summitTip / TIPS_MULTIPLIER),
            uint64(attestationTip / TIPS_MULTIPLIER),
            uint64(executionTip / TIPS_MULTIPLIER),
            uint64(deliveryTip / TIPS_MULTIPLIER)
        );
        assertEq(encodedTips256, encodedTips, "!encodeTips256");
    }

    function test_matchTips(RawTips memory rt, uint256 newValue) public {
        uint256 totalTips = uint256(rt.summitTip) + rt.attestationTip + rt.executionTip + rt.deliveryTip;
        vm.assume(totalTips <= type(uint64).max);
        uint256 maxNewTotalTips = uint256(rt.summitTip) + rt.attestationTip + rt.executionTip + type(uint64).max;
        newValue = bound(newValue, totalTips * TIPS_MULTIPLIER, maxNewTotalTips * TIPS_MULTIPLIER);
        Tips newTips = libHarness.matchValue(rt.castToTips(), newValue);
        // Should not exceed newValue
        assertLe(newTips.value(), newValue);
        // Increasing the delivery tip by one should exceed newValue
        if (newTips.deliveryTip() < type(uint64).max) {
            newTips = Tips.wrap(Tips.unwrap(newTips) + 1);
            assertGt(newTips.value(), newValue);
        }
    }

    function test_matchTips_revert_newValueTooLow(RawTips memory rt, uint256 newValue) public {
        uint256 totalTips = uint256(rt.summitTip) + rt.attestationTip + rt.executionTip + rt.deliveryTip;
        vm.assume(totalTips != 0 && totalTips <= type(uint64).max);
        newValue = bound(newValue, 0, totalTips * TIPS_MULTIPLIER - 1);
        vm.expectRevert(TipsValueTooLow.selector);
        libHarness.matchValue(rt.castToTips(), newValue);
    }

    function test_matchTips_revert_tipsOverflow(RawTips memory rt, uint256 newValue) public {
        uint256 totalTips = uint256(rt.summitTip) + rt.attestationTip + rt.executionTip + rt.deliveryTip;
        vm.assume(totalTips <= type(uint64).max);
        uint256 overflowTotalTips = uint256(rt.summitTip) + rt.attestationTip + rt.executionTip + 1 << 64;
        newValue = bound(newValue, overflowTotalTips * TIPS_MULTIPLIER, type(uint256).max);
        vm.expectRevert(TipsOverflow.selector);
        libHarness.matchValue(rt.castToTips(), newValue);
    }

    function test_emptyTips() public {
        test_encodeTips(RawTips(0, 0, 0, 0));
        assertEq(libHarness.emptyTips(), 0, "!emptyTips");
    }

    function test_tipsLength(RawTips memory rt) public {
        assertEq(
            abi.encodePacked(libHarness.encodeTips(rt.summitTip, rt.attestationTip, rt.executionTip, rt.deliveryTip))
                .length,
            TIPS_LENGTH
        );
    }
}
