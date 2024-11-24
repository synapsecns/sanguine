// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ZapDataV1, ZapDataV1Harness} from "../harnesses/ZapDataV1Harness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract ZapDataV1Test is Test {
    uint16 internal constant EXPECTED_VERSION = 1;

    ZapDataV1Harness internal harness;

    function setUp() public {
        harness = new ZapDataV1Harness();
    }

    function encodeZapData(
        uint16 version,
        uint16 amountPosition,
        address target,
        bytes memory payload
    )
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(version, amountPosition, target, payload);
    }

    function test_roundtrip_withAmount(
        address target,
        uint256 amount,
        bytes memory prefix,
        bytes memory postfix
    )
        public
        view
    {
        vm.assume(prefix.length + 32 + postfix.length < type(uint16).max);
        vm.assume(target != address(0));

        // We don't know the amount at the time of encoding, so we provide a placeholder.
        uint16 amountPosition = uint16(prefix.length);
        bytes memory encodedPayload = abi.encodePacked(prefix, uint256(0), postfix);
        // We expect the correct amount to be substituted in the payload at the time of Zap.
        bytes memory finalPayload = abi.encodePacked(prefix, amount, postfix);

        bytes memory zapData = harness.encodeV1(amountPosition, target, encodedPayload);

        harness.validateV1(zapData);
        assertEq(harness.version(zapData), 1);
        assertEq(harness.target(zapData), target);
        assertEq(harness.payload(zapData, amount), finalPayload);
        // Check against manually encoded ZapData.
        assertEq(zapData, encodeZapData(EXPECTED_VERSION, amountPosition, target, encodedPayload));
    }

    function test_roundtrip_noAmount(address target, uint256 amount, bytes memory payload) public view {
        vm.assume(payload.length < type(uint16).max);
        vm.assume(target != address(0));

        uint16 amountPosition = type(uint16).max;
        bytes memory zapData = harness.encodeV1(amountPosition, target, payload);

        harness.validateV1(zapData);
        assertEq(harness.version(zapData), 1);
        assertEq(harness.target(zapData), target);
        assertEq(harness.payload(zapData, amount), payload);
        // Check against manually encoded ZapData.
        assertEq(zapData, encodeZapData(EXPECTED_VERSION, amountPosition, target, payload));
    }

    function test_encodeV1_revert_targetZeroAddress() public {
        vm.expectRevert(ZapDataV1.ZapDataV1__TargetZeroAddress.selector);
        harness.encodeV1(type(uint16).max, address(0), "");
    }

    function test_encodeDecodeV1_revert_invalidAmountPosition(
        address target,
        uint16 amountPosition,
        uint256 amount,
        bytes memory payload
    )
        public
    {
        vm.assume(payload.length < type(uint16).max);
        vm.assume(target != address(0));
        // Make sure that (amountPosition + 32) is outside the bounds of the payload.
        uint16 incorrectMin = payload.length > 31 ? uint16(payload.length) - 31 : 0;
        uint16 incorrectMax = type(uint16).max - 1;
        amountPosition = uint16(bound(uint256(amountPosition), incorrectMin, incorrectMax));
        bytes memory invalidEncodedZapData = abi.encodePacked(uint16(1), amountPosition, target, payload);

        vm.expectRevert(ZapDataV1.ZapDataV1__InvalidEncoding.selector);
        harness.encodeV1(amountPosition, target, payload);

        // Validation should pass
        harness.validateV1(invalidEncodedZapData);
        harness.target(invalidEncodedZapData);
        // But payload extraction should revert
        vm.expectRevert(ZapDataV1.ZapDataV1__InvalidEncoding.selector);
        harness.payload(invalidEncodedZapData, amount);
    }

    function test_validateV1_revert_unsupportedVersion_withAmount(
        uint16 version,
        address target,
        bytes memory prefix,
        bytes memory postfix
    )
        public
    {
        vm.assume(version != 1);
        vm.assume(prefix.length + 32 + postfix.length < type(uint16).max);
        // We don't know the amount at the time of encoding, so we provide a placeholder.
        uint16 amountPosition = uint16(prefix.length);
        bytes memory encodedPayload = abi.encodePacked(prefix, uint256(0), postfix);

        bytes memory invalidEncodedZapData = encodeZapData(version, amountPosition, target, encodedPayload);

        vm.expectRevert(abi.encodeWithSelector(ZapDataV1.ZapDataV1__UnsupportedVersion.selector, version));
        harness.validateV1(invalidEncodedZapData);
    }

    function test_validateV1_revert_unsupportedVersion_noAmount(
        uint16 version,
        address target,
        bytes memory payload
    )
        public
    {
        vm.assume(version != 1);
        vm.assume(payload.length < type(uint16).max);

        uint16 amountPosition = type(uint16).max;
        bytes memory invalidEncodedZapData = encodeZapData(version, amountPosition, target, payload);

        vm.expectRevert(abi.encodeWithSelector(ZapDataV1.ZapDataV1__UnsupportedVersion.selector, version));
        harness.validateV1(invalidEncodedZapData);
    }

    function test_validateV1_revert_invalidLength(bytes calldata fuzzData) public {
        bytes memory minimumValidZapData = encodeZapData(EXPECTED_VERSION, type(uint16).max, address(0), "");
        uint256 invalidLength = fuzzData.length % minimumValidZapData.length;
        bytes calldata invalidEncodedZapData = fuzzData[:invalidLength];

        vm.expectRevert(ZapDataV1.ZapDataV1__InvalidEncoding.selector);
        harness.validateV1(invalidEncodedZapData);
    }
}
