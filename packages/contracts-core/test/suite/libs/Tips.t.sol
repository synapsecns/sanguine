// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {TipsHarness} from "../../harnesses/libs/TipsHarness.t.sol";

import {TIPS_MULTIPLIER, TIPS_LENGTH} from "../../../contracts/libs/Constants.sol";

// solhint-disable func-name-mixedcase
contract TipsLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    TipsHarness internal libHarness;

    function setUp() public {
        libHarness = new TipsHarness();
    }

    function test_encodeTips(uint64 summitTip, uint64 attestationTip, uint64 executionTip, uint64 deliveryTip) public {
        uint256 totalTips = uint256(summitTip) + attestationTip + executionTip + deliveryTip;
        vm.assume(totalTips <= type(uint64).max);
        // Test encoding
        uint256 encodedTips = libHarness.encodeTips(summitTip, attestationTip, executionTip, deliveryTip);
        uint256 expected = uint256(summitTip) * 2 ** 192 + uint256(attestationTip) * 2 ** 128
            + uint256(executionTip) * 2 ** 64 + uint256(deliveryTip);
        assertEq(encodedTips, expected, "!encodeTips");
        assertEq(libHarness.wrapPadded(encodedTips), expected, "!wrapPadded");
        // Test getters
        assertEq(libHarness.summitTip(encodedTips), summitTip, "!summitTip");
        assertEq(libHarness.attestationTip(encodedTips), attestationTip, "!attestationTip");
        assertEq(libHarness.executionTip(encodedTips), executionTip, "!executionTip");
        assertEq(libHarness.deliveryTip(encodedTips), deliveryTip, "!deliveryTip");
        assertEq(libHarness.value(encodedTips), totalTips * TIPS_MULTIPLIER, "!totalTips");
    }

    function test_emptyTips() public {
        test_encodeTips(0, 0, 0, 0);
        assertEq(libHarness.emptyTips(), 0, "!emptyTips");
    }
}
