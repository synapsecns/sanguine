// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { AttestationHarness } from "./harnesses/AttestationHarness.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract AttestationTest is Test {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    AttestationHarness internal harness;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");

    function setUp() public {
        harness = new AttestationHarness();
    }

    function test_formatAttestation() public {
        bytes memory _view = harness.formatAttestation(domain, nonce, root);
        assertEq(harness.domain(_view), domain);
        assertEq(harness.nonce(_view), nonce);
        assertEq(harness.root(_view), root);
        assertTrue(harness.isValid(_view));
    }

    function test_invalidUpdate_tooShort() public {
        // 1 byte shorter than standard
        bytes memory _view = abi.encodePacked(uint24(domain), nonce, root);
        assertFalse(harness.isValid(_view));
    }

    function test_invalidUpdate_tooLong() public {
        // 1 byte longer than standard
        bytes memory _view = abi.encodePacked(domain, nonce, root, uint8(69));
        assertFalse(harness.isValid(_view));
    }
}
