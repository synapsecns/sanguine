// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {StringUtilsHarness, StringUtils} from "../harnesses/StringUtilsHarness.sol";

import {Test} from "forge-std/Test.sol";

/// @notice Test suite for the StringUtils library.
contract StringUtilsTest is Test {
    StringUtilsHarness public libHarness;

    bytes public temp;

    function setUp() public {
        libHarness = new StringUtilsHarness();
    }

    function testLength() public {
        assertEq(libHarness.length(""), 0);
        assertEq(libHarness.length("0"), 1);
        assertEq(libHarness.length("0123456789"), 10);
    }

    // ════════════════════════════════════════════════ SLICING ══════════════════════════════════════════════════════

    function testSubstring() public {
        string memory str = "0123456789";
        // Empty string
        assertEq(libHarness.substring(str, 0, 0), "");
        assertEq(libHarness.substring(str, 6, 6), "");
        assertEq(libHarness.substring(str, 9, 9), "");
        // Single character
        assertEq(libHarness.substring(str, 0, 1), "0");
        assertEq(libHarness.substring(str, 5, 6), "5");
        assertEq(libHarness.substring(str, 9, 10), "9");
        // A few characters
        assertEq(libHarness.substring(str, 0, 2), "01");
        assertEq(libHarness.substring(str, 2, 5), "234");
        assertEq(libHarness.substring(str, 5, 10), "56789");
        // Full string
        assertEq(libHarness.substring(str, 0, 10), str);
    }

    function testSubstringRevertsWhenStartIndexOutOfBounds() public {
        string memory str = "0123456789";
        vm.expectRevert("StringUtils: Start index out of bounds");
        libHarness.substring(str, 11, 10);
    }

    function testSubstringRevertsWhenEndIndexOutOfBounds() public {
        string memory str = "0123456789";
        vm.expectRevert("StringUtils: End index out of bounds");
        libHarness.substring(str, 0, 11);
    }

    function testSubstringRevertsWhenStartIndexGreaterThanEndIndex() public {
        string memory str = "0123456789";
        vm.expectRevert("StringUtils: Invalid range");
        libHarness.substring(str, 5, 4);
    }

    function testSuffix() public {
        string memory str = "0123456789";
        assertEq(libHarness.suffix(str, 10), "");
        assertEq(libHarness.suffix(str, 9), "9");
        assertEq(libHarness.suffix(str, 5), "56789");
        assertEq(libHarness.suffix(str, 0), str);
    }

    function testSuffixRevertsWhenStartIndexOutOfBounds() public {
        string memory str = "0123456789";
        vm.expectRevert("StringUtils: Start index out of bounds");
        libHarness.suffix(str, 11);
    }

    function testPrefix() public {
        string memory str = "0123456789";
        assertEq(libHarness.prefix(str, 0), "");
        assertEq(libHarness.prefix(str, 1), "0");
        assertEq(libHarness.prefix(str, 5), "01234");
        assertEq(libHarness.prefix(str, 10), str);
    }

    function testPrefixRevertsWhenEndIndexOutOfBounds() public {
        string memory str = "0123456789";
        vm.expectRevert("StringUtils: End index out of bounds");
        libHarness.prefix(str, 11);
    }

    // ═══════════════════════════════════════════════ CONCATENATION ═══════════════════════════════════════════════════

    function appendToTemp(string memory str) internal {
        // Push all symbols from str to temp
        for (uint256 i = 0; i < bytes(str).length; i++) {
            temp.push(bytes(str)[i]);
        }
    }

    function testConcatTwoStrings(string memory a, string memory b) public {
        appendToTemp(a);
        appendToTemp(b);
        assertEq(libHarness.concat(a, b), string(temp));
    }

    function testConcatThreeStrings(string memory a, string memory b, string memory c) public {
        appendToTemp(a);
        appendToTemp(b);
        appendToTemp(c);
        assertEq(libHarness.concat(a, b, c), string(temp));
    }

    function testConcatFourStrings(string memory a, string memory b, string memory c, string memory d) public {
        appendToTemp(a);
        appendToTemp(b);
        appendToTemp(c);
        appendToTemp(d);
        assertEq(libHarness.concat(a, b, c, d), string(temp));
    }

    function testConcatFiveStrings(
        string memory a,
        string memory b,
        string memory c,
        string memory d,
        string memory e
    )
        public
    {
        appendToTemp(a);
        appendToTemp(b);
        appendToTemp(c);
        appendToTemp(d);
        appendToTemp(e);
        assertEq(libHarness.concat(a, b, c, d, e), string(temp));
    }

    function testConcatSixStrings(
        string memory a,
        string memory b,
        string memory c,
        string memory d,
        string memory e,
        string memory f
    )
        public
    {
        appendToTemp(a);
        appendToTemp(b);
        appendToTemp(c);
        appendToTemp(d);
        appendToTemp(e);
        appendToTemp(f);
        assertEq(libHarness.concat(a, b, c, d, e, f), string(temp));
    }

    function testDuplicate(string memory a, uint256 times) public {
        times = times % 10;
        for (uint256 i = 0; i < times; i++) {
            appendToTemp(a);
        }
        assertEq(libHarness.duplicate(a, times), string(temp));
    }

    // ════════════════════════════════════════════════ COMPARISON ═════════════════════════════════════════════════════

    function testEquals(string memory a, string memory b) public {
        if (libHarness.equals(a, b)) {
            assertEq(a, b);
        } else {
            assertNotEq(a, b);
        }
    }

    function testEqualsToClone(string memory a) public {
        bytes memory clone = new bytes(bytes(a).length);
        for (uint256 i = 0; i < bytes(a).length; i++) {
            clone[i] = bytes(a)[i];
        }
        assertTrue(libHarness.equals(a, string(clone)));
    }

    function testIndexOf() public {
        string memory str = "01201234012345";
        assertEq(libHarness.indexOf("", ""), 0);
        assertEq(libHarness.indexOf("", "0"), type(uint256).max);
        assertEq(libHarness.indexOf(str, ""), 0);
        // Single character
        assertEq(libHarness.indexOf(str, "0"), 0);
        assertEq(libHarness.indexOf(str, "2"), 2);
        assertEq(libHarness.indexOf(str, "4"), 7);
        // Multiple characters
        assertEq(libHarness.indexOf(str, "01"), 0);
        assertEq(libHarness.indexOf(str, "012"), 0);
        assertEq(libHarness.indexOf(str, "123"), 4);
        assertEq(libHarness.indexOf(str, "2345"), 10);
        // Whole string
        assertEq(libHarness.indexOf(str, str), 0);
        // Not found
        assertEq(libHarness.indexOf(str, "6"), type(uint256).max);
        assertEq(libHarness.indexOf(str, "02"), type(uint256).max);
        assertEq(libHarness.indexOf(str, "01230"), type(uint256).max);
    }

    function testLastIndexOf() public {
        string memory str = "01201234012345";
        assertEq(libHarness.lastIndexOf("", ""), 0);
        assertEq(libHarness.lastIndexOf("", "0"), type(uint256).max);
        assertEq(libHarness.lastIndexOf(str, ""), 14);
        // Single character
        assertEq(libHarness.lastIndexOf(str, "0"), 8);
        assertEq(libHarness.lastIndexOf(str, "2"), 10);
        assertEq(libHarness.lastIndexOf(str, "4"), 12);
        // Multiple characters
        assertEq(libHarness.lastIndexOf(str, "01"), 8);
        assertEq(libHarness.lastIndexOf(str, "012"), 8);
        assertEq(libHarness.lastIndexOf(str, "123"), 9);
        assertEq(libHarness.lastIndexOf(str, "2345"), 10);
        // Whole string
        assertEq(libHarness.lastIndexOf(str, str), 0);
        // Not found
        assertEq(libHarness.lastIndexOf(str, "6"), type(uint256).max);
        assertEq(libHarness.lastIndexOf(str, "02"), type(uint256).max);
        assertEq(libHarness.lastIndexOf(str, "01230"), type(uint256).max);
    }

    // ════════════════════════════════════════════ INTEGER CONVERSION ═════════════════════════════════════════════════

    function testToUint() public {
        assertEq(libHarness.toUint("0"), 0);
        assertEq(libHarness.toUint("1"), 1);
        assertEq(libHarness.toUint("42"), 42);
        assertEq(libHarness.toUint("1234567890"), 1_234_567_890);
        assertEq(
            libHarness.toUint("115792089237316195423570985008687907853269984665640564039457584007913129639935"),
            type(uint256).max
        );
    }

    function testToUintRevertsWhenNotADigit() public {
        vm.expectRevert("StringUtils: Not a digit");
        libHarness.toUint("0a");
    }

    function testFromUint() public {
        assertEq(libHarness.fromUint(0), "0");
        assertEq(libHarness.fromUint(1), "1");
        assertEq(libHarness.fromUint(42), "42");
        assertEq(libHarness.fromUint(1_234_567_890), "1234567890");
        assertEq(
            libHarness.fromUint(type(uint256).max),
            "115792089237316195423570985008687907853269984665640564039457584007913129639935"
        );
    }

    function testUintRoundtrip(uint256 val) public {
        assertEq(libHarness.toUint(libHarness.fromUint(val)), val);
    }

    // ═════════════════════════════════════════════ FLOAT CONVERSION ══════════════════════════════════════════════════

    function testFromFloat() public {
        // Zero decimals
        assertEq(libHarness.fromFloat(0, 0), "0.0");
        assertEq(libHarness.fromFloat(1, 0), "1.0");
        assertEq(libHarness.fromFloat(42, 0), "42.0");
        assertEq(libHarness.fromFloat(1_234_567_890, 0), "1234567890.0");
        assertEq(
            libHarness.fromFloat(type(uint256).max, 0),
            "115792089237316195423570985008687907853269984665640564039457584007913129639935.0"
        );
        // Six decimals
        assertEq(libHarness.fromFloat(0, 6), "0.000000");
        assertEq(libHarness.fromFloat(1, 6), "0.000001");
        assertEq(libHarness.fromFloat(42, 6), "0.000042");
        assertEq(libHarness.fromFloat(1_234_567_890, 6), "1234.567890");
        assertEq(
            libHarness.fromFloat(type(uint256).max, 6),
            "115792089237316195423570985008687907853269984665640564039457584007913129.639935"
        );
        // 18 decimals
        assertEq(libHarness.fromFloat(0, 18), "0.000000000000000000");
        assertEq(libHarness.fromFloat(1, 18), "0.000000000000000001");
        assertEq(libHarness.fromFloat(42, 18), "0.000000000000000042");
        assertEq(libHarness.fromFloat(1_234_567_890, 18), "0.000000001234567890");
        assertEq(
            libHarness.fromFloat(type(uint256).max, 18),
            "115792089237316195423570985008687907853269984665640564039457.584007913129639935"
        );
    }

    function testFromWei() public {
        assertEq(libHarness.fromWei(0), "0.000000000000000000");
        assertEq(libHarness.fromWei(1), "0.000000000000000001");
        assertEq(libHarness.fromWei(42), "0.000000000000000042");
        assertEq(libHarness.fromWei(1_234_567_890), "0.000000001234567890");
        assertEq(
            libHarness.fromWei(type(uint256).max),
            "115792089237316195423570985008687907853269984665640564039457.584007913129639935"
        );
    }
}
