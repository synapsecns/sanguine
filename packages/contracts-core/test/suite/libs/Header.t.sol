// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {HEADER_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {HeaderHarness} from "../../harnesses/libs/HeaderHarness.t.sol";

import {RawHeader} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract HeaderLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    HeaderHarness internal libHarness;

    function setUp() public {
        libHarness = new HeaderHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatHeader(RawHeader memory rh) public {
        // Test formatting
        bytes memory payload = libHarness.formatHeader(rh.origin, rh.nonce, rh.destination, rh.optimisticPeriod);
        assertEq(payload, abi.encodePacked(rh.origin, rh.nonce, rh.destination, rh.optimisticPeriod), "!formatHeader");
        // Test formatting checker
        checkCastToHeader({payload: payload, isHeader: true});
        // Test getters
        assertEq(libHarness.origin(payload), rh.origin, "!origin");
        assertEq(libHarness.nonce(payload), rh.nonce, "!nonce");
        assertEq(libHarness.destination(payload), rh.destination, "!destination");
        assertEq(libHarness.optimisticPeriod(payload), rh.optimisticPeriod, "!optimisticPeriod");
    }

    function test_isHeader(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToHeader({payload: payload, isHeader: length == HEADER_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToHeader(bytes memory payload, bool isHeader) public {
        if (isHeader) {
            assertTrue(libHarness.isHeader(payload), "!isHeader: when valid");
            assertEq(libHarness.castToHeader(payload), payload, "!castToHeader: when valid");
        } else {
            assertFalse(libHarness.isHeader(payload), "!isHeader: when valid");
            vm.expectRevert("Not a header payload");
            libHarness.castToHeader(payload);
        }
    }
}
