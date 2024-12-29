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
        address finalToken,
        address forwardTo,
        uint256 minFwdAmount,
        address target,
        bytes memory payload
    )
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(version, amountPosition, finalToken, forwardTo, minFwdAmount, target, payload);
    }

    function test_roundtrip_withAmount(
        address finalToken,
        address forwardTo,
        uint256 minFwdAmount,
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
        vm.assume(finalToken != address(0) || forwardTo == address(0));
        vm.assume(forwardTo != address(0) || minFwdAmount == 0);

        // We don't know the amount at the time of encoding, so we provide a placeholder.
        uint16 amountPosition = uint16(prefix.length);
        bytes memory encodedPayload = abi.encodePacked(prefix, uint256(0), postfix);
        // We expect the correct amount to be substituted in the payload at the time of Zap.
        bytes memory finalPayload = abi.encodePacked(prefix, amount, postfix);

        bytes memory zapData =
            harness.encodeV1(amountPosition, finalToken, forwardTo, minFwdAmount, target, encodedPayload);

        harness.validateV1(zapData);
        assertEq(harness.version(zapData), 1);
        assertEq(harness.finalToken(zapData), finalToken);
        assertEq(harness.forwardTo(zapData), forwardTo);
        assertEq(harness.minFwdAmount(zapData), minFwdAmount);
        assertEq(harness.target(zapData), target);
        assertEq(harness.payload(zapData, amount), finalPayload);
        // Check against manually encoded ZapData.
        assertEq(
            zapData,
            encodeZapData(EXPECTED_VERSION, amountPosition, finalToken, forwardTo, minFwdAmount, target, encodedPayload)
        );
    }

    function test_roundtrip_noAmount(
        address finalToken,
        address forwardTo,
        uint256 minFwdAmount,
        address target,
        uint256 amount,
        bytes memory payload
    )
        public
        view
    {
        vm.assume(payload.length < type(uint16).max);
        vm.assume(target != address(0));
        vm.assume(finalToken != address(0) || forwardTo == address(0));
        vm.assume(forwardTo != address(0) || minFwdAmount == 0);

        uint16 amountPosition = type(uint16).max;
        bytes memory zapData = harness.encodeV1(amountPosition, finalToken, forwardTo, minFwdAmount, target, payload);

        harness.validateV1(zapData);
        assertEq(harness.version(zapData), 1);
        assertEq(harness.finalToken(zapData), finalToken);
        assertEq(harness.forwardTo(zapData), forwardTo);
        assertEq(harness.minFwdAmount(zapData), minFwdAmount);
        assertEq(harness.target(zapData), target);
        assertEq(harness.payload(zapData, amount), payload);
        // Check against manually encoded ZapData.
        assertEq(
            zapData,
            encodeZapData(EXPECTED_VERSION, amountPosition, finalToken, forwardTo, minFwdAmount, target, payload)
        );
    }

    function test_encodeV1_revert_targetZeroAddress() public {
        vm.expectRevert(ZapDataV1.ZapDataV1__TargetZeroAddress.selector);
        harness.encodeV1(type(uint16).max, address(0), address(0), 0, address(0), "");
    }

    function test_encodeV1_revert_forwardToWithoutFinalToken() public {
        vm.expectRevert(ZapDataV1.ZapDataV1__ForwardParamsIncorrect.selector);
        harness.encodeV1(type(uint16).max, address(0), address(1), 0, address(2), "");
    }

    function test_encodeV1_revert_minFwdAmountWithoutForwardTo() public {
        vm.expectRevert(ZapDataV1.ZapDataV1__ForwardParamsIncorrect.selector);
        harness.encodeV1(type(uint16).max, address(1), address(0), 1, address(2), "");
    }

    function test_encodeV1_revert_payloadLengthAboveMax() public {
        vm.expectRevert(ZapDataV1.ZapDataV1__PayloadLengthAboveMax.selector);
        harness.encodeV1(type(uint16).max, address(0), address(0), 0, address(1), new bytes(2 ** 16));
    }

    function test_encodeDecodeV1_revert_invalidAmountPosition(
        address finalToken,
        address forwardTo,
        uint256 minFwdAmount,
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
        bytes memory invalidEncodedZapData =
            encodeZapData(uint16(1), amountPosition, finalToken, forwardTo, minFwdAmount, target, payload);

        vm.expectRevert(ZapDataV1.ZapDataV1__InvalidEncoding.selector);
        harness.encodeV1(amountPosition, finalToken, forwardTo, minFwdAmount, target, payload);

        // Validation should pass
        harness.validateV1(invalidEncodedZapData);
        harness.finalToken(invalidEncodedZapData);
        harness.forwardTo(invalidEncodedZapData);
        harness.target(invalidEncodedZapData);
        // But payload extraction should revert
        vm.expectRevert(ZapDataV1.ZapDataV1__InvalidEncoding.selector);
        harness.payload(invalidEncodedZapData, amount);
    }

    function test_validateV1_revert_unsupportedVersion_withAmount(
        uint16 version,
        address finalToken,
        address forwardTo,
        uint256 minFwdAmount,
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

        bytes memory invalidEncodedZapData =
            encodeZapData(version, amountPosition, finalToken, forwardTo, minFwdAmount, target, encodedPayload);

        vm.expectRevert(abi.encodeWithSelector(ZapDataV1.ZapDataV1__UnsupportedVersion.selector, version));
        harness.validateV1(invalidEncodedZapData);
    }

    function test_validateV1_revert_unsupportedVersion_noAmount(
        uint16 version,
        address finalToken,
        address forwardTo,
        uint256 minFwdAmount,
        address target,
        bytes memory payload
    )
        public
    {
        vm.assume(version != 1);
        vm.assume(payload.length < type(uint16).max);

        uint16 amountPosition = type(uint16).max;
        bytes memory invalidEncodedZapData =
            encodeZapData(version, amountPosition, finalToken, forwardTo, minFwdAmount, target, payload);

        vm.expectRevert(abi.encodeWithSelector(ZapDataV1.ZapDataV1__UnsupportedVersion.selector, version));
        harness.validateV1(invalidEncodedZapData);
    }

    function test_validateV1_revert_invalidLength(bytes calldata fuzzData) public {
        bytes memory minimumValidZapData =
            encodeZapData(EXPECTED_VERSION, type(uint16).max, address(0), address(0), 0, address(0), "");
        uint256 invalidLength = fuzzData.length % minimumValidZapData.length;
        bytes calldata invalidEncodedZapData = fuzzData[:invalidLength];

        vm.expectRevert(ZapDataV1.ZapDataV1__InvalidEncoding.selector);
        harness.validateV1(invalidEncodedZapData);
    }
}
