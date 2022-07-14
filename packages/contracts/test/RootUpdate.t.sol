// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { RootUpdateHarness } from "./harnesses/RootUpdateHarness.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract RootUpdateTest is Test {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    RootUpdateHarness internal harness;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");

    function setUp() public {
        harness = new RootUpdateHarness();
    }

    function test_formatRootUpdate() public {
        bytes memory rootUpdate = harness.formatRootUpdate(domain, nonce, root);
        assertEq(harness.domain(rootUpdate), domain);
        assertEq(harness.nonce(rootUpdate), nonce);
        assertEq(harness.root(rootUpdate), root);
        assertTrue(harness.isValid(rootUpdate));
    }

    function test_invalidUpdate_tooShort() public {
        // 1 byte shorter than standard
        bytes memory rootUpdate = abi.encodePacked(uint24(domain), nonce, root);
        assertFalse(harness.isValid(rootUpdate));
    }

    function test_invalidUpdate_tooLong() public {
        // 1 byte longer than standard
        bytes memory rootUpdate = abi.encodePacked(domain, nonce, root, uint8(69));
        assertFalse(harness.isValid(rootUpdate));
    }
}
