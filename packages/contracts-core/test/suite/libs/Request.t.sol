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

    function test_encodeRequest(RawRequest memory rs) public {
        uint256 encoded = libHarness.encodeRequest(rs.gasDrop, rs.gasLimit);
        uint256 expected = rs.gasLimit + uint256(rs.gasDrop) * 2 ** 64;
        assertEq(encoded, expected, "!encodeRequest");
        assertEq(libHarness.wrapPadded(encoded), expected, "!wrapPadded");
        assertEq(libHarness.gasLimit(encoded), rs.gasLimit, "!gasLimit");
        assertEq(libHarness.gasDrop(encoded), rs.gasDrop, "!gasDrop");
    }
}
