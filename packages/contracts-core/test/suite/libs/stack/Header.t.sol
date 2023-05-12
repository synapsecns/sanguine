// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {HEADER_LENGTH} from "../../../../contracts/libs/Constants.sol";
import {FlagOutOfRange} from "../../../../contracts/libs/Errors.sol";

import {SynapseLibraryTest} from "../../../utils/SynapseLibraryTest.t.sol";
import {HeaderHarness, MessageFlag} from "../../../harnesses/libs/stack/HeaderHarness.t.sol";

import {RawHeader} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract HeaderLibraryTest is SynapseLibraryTest {
    HeaderHarness internal libHarness;

    function setUp() public {
        libHarness = new HeaderHarness();
    }

    function test_encodeHeader(RawHeader memory rh) public {
        // Make sure flag is within bounds
        rh.boundFlag();
        // Test encoding
        uint136 encoded = libHarness.encodeHeader(rh.flag, rh.origin, rh.nonce, rh.destination, rh.optimisticPeriod);
        uint256 expected = uint256(rh.flag) * 2 ** 128 + uint256(rh.origin) * 2 ** 96 + uint256(rh.nonce) * 2 ** 64
            + uint256(rh.destination) * 2 ** 32 + uint256(rh.optimisticPeriod);
        assertEq(encoded, expected, "!encodeHeader");
        assertEq(libHarness.wrapPadded(encoded), expected, "!wrapPadded");
        assertTrue(libHarness.isHeader(encoded), "!isHeader");
        // Test getters
        assertEq(uint8(libHarness.flag(encoded)), rh.flag, "!flag");
        assertEq(libHarness.origin(encoded), rh.origin, "!origin");
        assertEq(libHarness.nonce(encoded), rh.nonce, "!nonce");
        assertEq(libHarness.destination(encoded), rh.destination, "!destination");
        assertEq(libHarness.optimisticPeriod(encoded), rh.optimisticPeriod, "!optimisticPeriod");
        // Test hashing
        assertEq(libHarness.leaf(encoded), keccak256(abi.encode(expected)), "!leaf");
    }

    function test_headerLength(RawHeader memory rh) public {
        rh.boundFlag();
        bytes memory packedHeader =
            abi.encodePacked(libHarness.encodeHeader(rh.flag, rh.origin, rh.nonce, rh.destination, rh.optimisticPeriod));
        assertEq(packedHeader.length, HEADER_LENGTH);
    }

    function test_wrapPadded_revert_flagOutOfRange(uint8 flag, uint128 remainder) public {
        flag = uint8(bound(flag, uint8(type(MessageFlag).max) + 1, 255));
        uint256 encodedBadFlag = uint256(flag) * 2 ** 128 + remainder;
        assertFalse(libHarness.isHeader(encodedBadFlag));
        vm.expectRevert(FlagOutOfRange.selector);
        libHarness.wrapPadded(encodedBadFlag);
    }
}
