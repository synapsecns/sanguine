// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {REQUEST_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {RequestHarness} from "../../harnesses/libs/RequestHarness.t.sol";

import {RawRequest} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract RequestLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    RequestHarness internal libHarness;

    function setUp() public {
        libHarness = new RequestHarness();
    }

    function test_formatRequest(RawRequest memory rs) public {
        bytes memory payload = libHarness.formatRequest(rs.gasLimit);
        // Test formatting of request
        assertEq(payload, abi.encodePacked(rs.gasLimit), "!formatRequest");
        checkCastToRequest({payload: payload, isRequest: true});
        // Test getters
        assertEq(libHarness.gasLimit(payload), rs.gasLimit, "!gasLimit");
    }

    function test_isRequest(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToRequest({payload: payload, isRequest: length == REQUEST_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToRequest(bytes memory payload, bool isRequest) public {
        if (isRequest) {
            assertTrue(libHarness.isRequest(payload), "!isRequest: when valid");
            assertEq(libHarness.castToRequest(payload), payload, "!castToRequest: when valid");
        } else {
            assertFalse(libHarness.isRequest(payload), "!isRequest: when valid");
            vm.expectRevert("Not a request");
            libHarness.castToRequest(payload);
        }
    }
}
