// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MulticallFailed} from "../../../contracts/libs/Errors.sol";
import {MultiCallable, MultiCallableHarness} from "../../harnesses/base/MultiCallableHarness.t.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable code-complexity
// solhint-disable func-name-mixedcase
contract MultiCallableTest is Test {
    MultiCallableHarness public mcHarness;

    uint256 public constant AMOUNT = 5;

    function setUp() public {
        mcHarness = new MultiCallableHarness();
    }

    function test_multicall(uint256 mask) public {
        MultiCallableHarness.Call[] memory calls = new MultiCallableHarness.Call[](AMOUNT);
        bool[] memory willRevert = new bool[](AMOUNT);
        bool success = true;
        uint256 successfulCalls = 0;
        for (uint256 i = 0; i < AMOUNT; ++i) {
            uint256 value = i + 1;
            calls[i].allowFailure = (mask & (1 << (2 * i))) != 0;
            calls[i].callData = abi.encodeWithSelector(mcHarness.addUint.selector, value);
            willRevert[i] = (mask & (1 << (2 * i + 1))) != 0;
            if (willRevert[i]) {
                mcHarness.toggleRevert(value, true);
                if (!calls[i].allowFailure) {
                    success = false;
                }
            } else {
                ++successfulCalls;
            }
        }
        if (!success) vm.expectRevert(MulticallFailed.selector);
        MultiCallable.Result[] memory results = mcHarness.multicall(calls);
        if (success) {
            require(results.length == AMOUNT, "results.length != AMOUNT");
            uint256[] memory buffer = new uint256[](successfulCalls);
            successfulCalls = 0;
            for (uint256 i = 0; i < AMOUNT; ++i) {
                uint256 value = i + 1;
                if (willRevert[i]) {
                    assertFalse(results[i].success);
                    assertEq(results[i].returnData, abi.encodePacked(MultiCallableHarness.GmError.selector));
                } else {
                    assertTrue(results[i].success);
                    assertEq(results[i].returnData, abi.encode(2 * value));
                    buffer[successfulCalls++] = value;
                }
            }
            // Check that order of calls is preserved
            assertEq(abi.encode(mcHarness.getBuffer()), abi.encode(buffer));
        }
    }

    function test_multicall_allSuccess() public {
        test_multicall(0);
    }

    function test_multicall_firstCallFails() public {
        test_multicall(2);
    }
}
