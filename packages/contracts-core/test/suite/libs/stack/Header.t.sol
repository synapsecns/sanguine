// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {HEADER_LENGTH} from "../../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest} from "../../../utils/SynapseLibraryTest.t.sol";
import {HeaderHarness} from "../../../harnesses/libs/stack/HeaderHarness.t.sol";

import {RawHeader} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract HeaderLibraryTest is SynapseLibraryTest {
    HeaderHarness internal libHarness;

    function setUp() public {
        libHarness = new HeaderHarness();
    }

    function test_encodeHeader(RawHeader memory rh) public {
        // Test encoding
        uint128 encoded = libHarness.encodeHeader(rh.origin, rh.nonce, rh.destination, rh.optimisticPeriod);
        uint256 expected = uint256(rh.origin) * 2 ** 96 + uint256(rh.nonce) * 2 ** 64
            + uint256(rh.destination) * 2 ** 32 + uint256(rh.optimisticPeriod);
        assertEq(encoded, expected, "!encodeHeader");
        assertEq(libHarness.wrapPadded(encoded), expected, "!wrapPadded");
        // Test getters
        assertEq(libHarness.origin(encoded), rh.origin, "!origin");
        assertEq(libHarness.nonce(encoded), rh.nonce, "!nonce");
        assertEq(libHarness.destination(encoded), rh.destination, "!destination");
        assertEq(libHarness.optimisticPeriod(encoded), rh.optimisticPeriod, "!optimisticPeriod");
    }

    function test_headerLength(RawHeader memory rh) public {
        assertEq(
            abi.encodePacked(libHarness.encodeHeader(rh.origin, rh.nonce, rh.destination, rh.optimisticPeriod)).length,
            HEADER_LENGTH
        );
    }
}
