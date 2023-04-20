// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {TipsHarness} from "../../harnesses/libs/TipsHarness.t.sol";

import {TIPS_MULTIPLIER, TIPS_LENGTH} from "../../../contracts/libs/Constants.sol";

import {RawTips} from "../../utils/libs/SynapseStructs.t.sol";

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
