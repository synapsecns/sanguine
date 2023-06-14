// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {REQUEST_LENGTH} from "../../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {RequestHarness} from "../../../harnesses/libs/stack/RequestHarness.t.sol";

import {RawRequest} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract RequestLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    RequestHarness internal libHarness;

    function setUp() public {
        libHarness = new RequestHarness();
    }

    function test_encodeRequest(RawRequest memory rr) public {
        uint192 encoded = libHarness.encodeRequest(rr.gasDrop, rr.gasLimit, rr.version);
        uint256 expected = uint256(rr.gasDrop) * 2 ** 96 + uint256(rr.gasLimit) * 2 ** 32 + rr.version;
        assertEq(encoded, expected, "!encodeRequest");
        assertEq(libHarness.wrapPadded(encoded), expected, "!wrapPadded");
        assertEq(libHarness.gasLimit(encoded), rr.gasLimit, "!gasLimit");
        assertEq(libHarness.gasDrop(encoded), rr.gasDrop, "!gasDrop");
        assertEq(libHarness.version(encoded), rr.version, "!version");
    }

    function test_requestLength(RawRequest memory rr) public {
        bytes memory packedRequest = abi.encodePacked(libHarness.encodeRequest(rr.gasDrop, rr.gasLimit, rr.version));
        assertEq(packedRequest.length, REQUEST_LENGTH);
    }
}
