// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {REQUEST_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {RequestHarness} from "../../harnesses/libs/RequestHarness.t.sol";

import {RawRequest} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract RequestLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    RequestHarness internal libHarness;

    function setUp() public {
        libHarness = new RequestHarness();
    }

    function test_encodeRequest(RawRequest memory rr) public {
        uint160 encoded = libHarness.encodeRequest(rr.gasDrop, rr.gasLimit);
        uint256 expected = rr.gasLimit + uint256(rr.gasDrop) * 2 ** 64;
        assertEq(encoded, expected, "!encodeRequest");
        assertEq(libHarness.wrapPadded(encoded), expected, "!wrapPadded");
        assertEq(libHarness.gasLimit(encoded), rr.gasLimit, "!gasLimit");
        assertEq(libHarness.gasDrop(encoded), rr.gasDrop, "!gasDrop");
    }

    function test_requestLength(RawRequest memory rr) public {
        assertEq(abi.encodePacked(libHarness.encodeRequest(rr.gasDrop, rr.gasLimit)).length, REQUEST_LENGTH);
    }
}
