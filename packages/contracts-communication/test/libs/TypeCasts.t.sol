// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {TypeCastsHarness} from "../harnesses/TypeCastsHarness.sol";

import {Test} from "forge-std/Test.sol";

contract TypeCastsLibraryTest is Test {
    TypeCastsHarness public libHarness;

    function setUp() public {
        libHarness = new TypeCastsHarness();
    }

    function test_addressToBytes32() public {
        assertEq(
            libHarness.addressToBytes32(address(0x01)),
            0x0000000000000000000000000000000000000000000000000000000000000001
        );
        assertEq(
            libHarness.addressToBytes32(0xfEdcBA9876543210FedCBa9876543210fEdCBa98),
            0x000000000000000000000000FEDCBA9876543210FEDCBA9876543210FEDCBA98
        );
    }

    function test_addressToBytes32(address addr) public {
        bytes32 expected = bytes32(abi.encode(addr));
        assertEq(libHarness.addressToBytes32(addr), expected);
    }

    function test_bytes32ToAddress() public {
        assertEq(
            libHarness.bytes32ToAddress(0x0000000000000000000000000000000000000000000000000000000000000001),
            address(0x01)
        );
        assertEq(
            libHarness.bytes32ToAddress(0xFFFFFFFFFFFFFFFFFFFFFFFF0000000000000000000000000000000000000001),
            address(0x01)
        );
        assertEq(
            libHarness.bytes32ToAddress(0x000000000000000000000000FEDCBA9876543210FEDCBA9876543210FEDCBA98),
            0xfEdcBA9876543210FedCBa9876543210fEdCBa98
        );
        assertEq(
            libHarness.bytes32ToAddress(0xFFFFFFFFFFFFFFFFFFFFFFFFFEDCBA9876543210FEDCBA9876543210FEDCBA98),
            0xfEdcBA9876543210FedCBa9876543210fEdCBa98
        );
    }

    function test_bytes32ToAddress(bytes32 b) public {
        // Clear the first 96 bits
        bytes32 cleared = b & bytes32(uint256(type(uint160).max));
        address expected = abi.decode(bytes.concat(cleared), (address));
        assertEq(libHarness.bytes32ToAddress(b), expected);
    }

    function test_roundtrip(address addr) public {
        bytes32 b = libHarness.addressToBytes32(addr);
        assertEq(libHarness.bytes32ToAddress(b), addr);
    }
}
