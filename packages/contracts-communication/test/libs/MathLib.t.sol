// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MathLibHarness} from "../harnesses/MathLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract MathLibraryTest is Test {
    MathLibHarness public libHarness;

    function setUp() public {
        libHarness = new MathLibHarness();
    }

    function test_roundUpToWord_zero() public {
        assertEq(libHarness.roundUpToWord(0), 0);
    }

    function test_roundUpToWord_nonZero_fitsInUint256(uint256 x) public {
        x = bound(x, 1, type(uint256).max - 31);
        uint256 rounded = libHarness.roundUpToWord(x);
        // rounded is a multiple of 32
        assertEq(rounded % 32, 0);
        // rounded - 32 < x <= rounded
        assertGe(rounded, x);
        assertLt(rounded - 32, x);
    }

    function test_roundUpToWord_nonZero_overflows(uint256 x) public {
        x = bound(x, type(uint256).max - 30, type(uint256).max);
        assertEq(libHarness.roundUpToWord(x), 0);
    }
}
