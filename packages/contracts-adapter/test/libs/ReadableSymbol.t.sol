// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ReadableSymbol, ReadableSymbolHarness} from "../harnesses/ReadableSymbolHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract ReadableSymbolTest is Test {
    ReadableSymbolHarness internal harness;

    function setUp() public {
        harness = new ReadableSymbolHarness();
    }

    function test_toBytes31_nonEmpty() public {
        assertEq(harness.toBytes31("Hello, World!"), bytes31("Hello, World!"));
    }

    function test_toBytes31_exactly31() public {
        assertEq(harness.toBytes31("This is exactly 31 chars long!!"), bytes31("This is exactly 31 chars long!!"));
    }

    function test_toBytes31_empty() public {
        assertEq(harness.toBytes31(""), bytes31(0));
    }

    function test_toBytes31_tooLong() public {
        vm.expectRevert(ReadableSymbol.ReadableSymbol__StringTooLong.selector);
        harness.toBytes31(string(new bytes(32)));
    }

    function test_toString_nonEmpty() public {
        assertEq(harness.toString(bytes31("Hello, World!")), "Hello, World!");
    }

    function test_toString_empty() public {
        assertEq(harness.toString(bytes31(0)), "");
    }

    function test_toString_exactly31() public {
        assertEq(harness.toString(bytes31("This is exactly 31 chars long!!")), "This is exactly 31 chars long!!");
    }
}
