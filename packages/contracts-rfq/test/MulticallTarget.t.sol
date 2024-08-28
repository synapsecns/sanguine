// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IMulticallTarget} from "../contracts/interfaces/IMulticallTarget.sol";
import {MulticallTargetHarness} from "./harnesses/MulticallTargetHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract MulticallTargetTest is Test {
    MulticallTargetHarness public harness;

    function setUp() public {
        harness = new MulticallTargetHarness();
        harness.setAddressField(address(1));
        harness.setUintField(2);
    }

    function getEncodedStringRevertMessage() internal view returns (bytes memory) {
        return abi.encodeWithSignature("Error(string)", harness.REVERT_MESSAGE());
    }

    function getNoRevertsData() internal view returns (bytes[] memory) {
        return toArray(
            abi.encodeCall(harness.addressField, ()),
            abi.encodeCall(harness.setAddressField, (address(1234))),
            abi.encodeCall(harness.setUintField, (42)),
            abi.encodeCall(harness.setAddressField, (address(0xDEADBEAF)))
        );
    }

    function getNoRevertsResults() internal pure returns (IMulticallTarget.Result[] memory) {
        return toArray(
            IMulticallTarget.Result(true, abi.encode(address(1))),
            IMulticallTarget.Result(true, abi.encode(address(1234))),
            IMulticallTarget.Result(true, abi.encode(42)),
            IMulticallTarget.Result(true, abi.encode(address(0xDEADBEAF)))
        );
    }

    function getCustomErrorRevertData() internal view returns (bytes[] memory) {
        return toArray(
            abi.encodeCall(harness.setAddressField, (address(1234))),
            abi.encodeCall(harness.setUintField, (42)),
            abi.encodeCall(harness.customErrorRevert, ()),
            abi.encodeCall(harness.setAddressField, (address(0xDEADBEAF)))
        );
    }

    function getCustomErrorRevertResults() internal pure returns (IMulticallTarget.Result[] memory) {
        return toArray(
            IMulticallTarget.Result(true, abi.encode(address(1234))),
            IMulticallTarget.Result(true, abi.encode(42)),
            IMulticallTarget.Result(false, abi.encodeWithSelector(MulticallTargetHarness.CustomError.selector)),
            IMulticallTarget.Result(true, abi.encode(address(0xDEADBEAF)))
        );
    }

    function getStringRevertData() internal view returns (bytes[] memory) {
        return toArray(
            abi.encodeCall(harness.setAddressField, (address(1234))),
            abi.encodeCall(harness.setUintField, (42)),
            abi.encodeCall(harness.revertingFunction, ()),
            abi.encodeCall(harness.setAddressField, (address(0xDEADBEAF)))
        );
    }

    function getStringRevertResults() internal view returns (IMulticallTarget.Result[] memory) {
        return toArray(
            IMulticallTarget.Result(true, abi.encode(address(1234))),
            IMulticallTarget.Result(true, abi.encode(42)),
            IMulticallTarget.Result(false, abi.encodeWithSignature("Error(string)", harness.REVERT_MESSAGE())),
            IMulticallTarget.Result(true, abi.encode(address(0xDEADBEAF)))
        );
    }

    // ══════════════════════════════════════════ MULTICALL (NO RESULTS) ═══════════════════════════════════════════════

    function test_multicallNoResults_ignoreReverts_noReverts() public {
        bytes[] memory data = getNoRevertsData();
        harness.multicallNoResults({data: data, ignoreReverts: true});

        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallNoResults_ignoreReverts_withCustomErrorRevert() public {
        bytes[] memory data = getCustomErrorRevertData();
        harness.multicallNoResults({data: data, ignoreReverts: true});

        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallNoResults_ignoreReverts_withStringRevert() public {
        bytes[] memory data = getStringRevertData();
        harness.multicallNoResults({data: data, ignoreReverts: true});

        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallNoResults_dontIgnoreReverts_noReverts() public {
        bytes[] memory data = getNoRevertsData();
        harness.multicallNoResults({data: data, ignoreReverts: false});

        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallNoResults_dontIgnoreReverts_withCustomErrorRevert() public {
        bytes[] memory data = getCustomErrorRevertData();
        vm.expectRevert(MulticallTargetHarness.CustomError.selector);
        harness.multicallNoResults({data: data, ignoreReverts: false});
    }

    function test_multicallNoResults_dontIgnoreReverts_withStringRevert() public {
        bytes[] memory data = getStringRevertData();
        string memory revertMessage = harness.REVERT_MESSAGE();
        vm.expectRevert(bytes(revertMessage));
        harness.multicallNoResults({data: data, ignoreReverts: false});
    }

    // ═════════════════════════════════════════ MULTICALL (WITH RESULTS) ══════════════════════════════════════════════

    function test_multicallWithResults_ignoreReverts_noReverts() public {
        bytes[] memory data = getNoRevertsData();
        IMulticallTarget.Result[] memory results = harness.multicallWithResults({data: data, ignoreReverts: true});

        assertEq(results, getNoRevertsResults());
        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallWithResults_ignoreReverts_withCustomErrorRevert() public {
        bytes[] memory data = getCustomErrorRevertData();
        IMulticallTarget.Result[] memory results = harness.multicallWithResults({data: data, ignoreReverts: true});

        assertEq(results, getCustomErrorRevertResults());
        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallWithResults_ignoreReverts_withStringRevert() public {
        bytes[] memory data = getStringRevertData();
        IMulticallTarget.Result[] memory results = harness.multicallWithResults({data: data, ignoreReverts: true});

        assertEq(results, getStringRevertResults());
        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallWithResults_dontIgnoreReverts_noReverts() public {
        bytes[] memory data = getNoRevertsData();
        IMulticallTarget.Result[] memory results = harness.multicallWithResults({data: data, ignoreReverts: false});

        assertEq(results, getNoRevertsResults());
        assertEq(harness.addressField(), address(0xDEADBEAF));
        assertEq(harness.uintField(), 42);
    }

    function test_multicallWithResults_dontIgnoreReverts_withCustomErrorRevert() public {
        bytes[] memory data = getCustomErrorRevertData();
        vm.expectRevert(MulticallTargetHarness.CustomError.selector);
        harness.multicallWithResults({data: data, ignoreReverts: false});
    }

    function test_multicallWithResults_dontIgnoreReverts_withStringRevert() public {
        bytes[] memory data = getStringRevertData();
        string memory revertMessage = harness.REVERT_MESSAGE();
        vm.expectRevert(bytes(revertMessage));
        harness.multicallWithResults({data: data, ignoreReverts: false});
    }

    // ══════════════════════════════════════════════════ VIEW ════════════════════════════════════════════════════

    function assertEq(IMulticallTarget.Result memory a, IMulticallTarget.Result memory b) internal pure {
        assertEq(a.success, b.success);
        assertEq(a.returnData, b.returnData);
    }

    function assertEq(IMulticallTarget.Result[] memory a, IMulticallTarget.Result[] memory b) internal pure {
        assertEq(a.length, b.length);
        for (uint256 i = 0; i < a.length; i++) {
            assertEq(a[i], b[i]);
        }
    }

    function toArray(
        bytes memory a,
        bytes memory b,
        bytes memory c,
        bytes memory d
    )
        internal
        pure
        returns (bytes[] memory arr)
    {
        arr = new bytes[](4);
        arr[0] = a;
        arr[1] = b;
        arr[2] = c;
        arr[3] = d;
    }

    function toArray(
        IMulticallTarget.Result memory a,
        IMulticallTarget.Result memory b,
        IMulticallTarget.Result memory c,
        IMulticallTarget.Result memory d
    )
        internal
        pure
        returns (IMulticallTarget.Result[] memory arr)
    {
        arr = new IMulticallTarget.Result[](4);
        arr[0] = a;
        arr[1] = b;
        arr[2] = c;
        arr[3] = d;
    }
}
