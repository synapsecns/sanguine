// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { HomeUpdateHarness } from "./harnesses/HomeUpdateHarness.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract HomeUpdateTest is Test {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    HomeUpdateHarness internal harness;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");

    function setUp() public {
        harness = new HomeUpdateHarness();
    }

    function test_formatHomeUpdate() public {
        bytes memory homeUpdate = harness.formatHomeUpdate(domain, nonce, root);
        assertEq(harness.domain(homeUpdate), domain);
        assertEq(harness.nonce(homeUpdate), nonce);
        assertEq(harness.root(homeUpdate), root);
        assertTrue(harness.isValid(homeUpdate));
    }

    function test_invalidUpdate_tooShort() public {
        // 1 byte shorter than standard
        bytes memory homeUpdate = abi.encodePacked(uint24(domain), nonce, root);
        assertFalse(harness.isValid(homeUpdate));
    }

    function test_invalidUpdate_tooLong() public {
        // 1 byte longer than standard
        bytes memory homeUpdate = abi.encodePacked(domain, nonce, root, uint8(69));
        assertFalse(harness.isValid(homeUpdate));
    }
}
