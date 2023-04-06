// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {BaseMessageHarness} from "../../harnesses/libs/BaseMessageHarness.t.sol";

import {RawBaseMessage} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract BaseMessageLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    BaseMessageHarness internal libHarness;

    function setUp() public {
        libHarness = new BaseMessageHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatBaseMessage(RawBaseMessage memory rbm) public {
        bytes memory tipsPayload = rbm.tips.formatTips();
        bytes memory requestPayload = rbm.request.formatRequest();
        // Test formatting
        bytes memory payload =
            libHarness.formatBaseMessage(rbm.sender, rbm.recipient, tipsPayload, requestPayload, rbm.content);
        assertEq(
            payload,
            abi.encodePacked(rbm.sender, rbm.recipient, tipsPayload, requestPayload, rbm.content),
            "!formatBaseMessage"
        );
        // Test formatting checker
        checkCastToBaseMessage({payload: payload, isBaseMessage: true});
        // Test getters
        assertEq(libHarness.sender(payload), rbm.sender, "!sender");
        assertEq(libHarness.recipient(payload), rbm.recipient, "!recipient");
        assertEq(libHarness.tips(payload), tipsPayload, "!tips");
        assertEq(libHarness.request(payload), requestPayload, "!request");
        assertEq(libHarness.content(payload), rbm.content, "!content");
    }

    function test_isBaseMessage(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToBaseMessage({payload: payload, isBaseMessage: length >= MIN_BASE_MESSAGE_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToBaseMessage(bytes memory payload, bool isBaseMessage) public {
        if (isBaseMessage) {
            assertTrue(libHarness.isBaseMessage(payload), "!isBaseMessage: when valid");
            assertEq(libHarness.castToBaseMessage(payload), payload, "!castToBaseMessage: when valid");
        } else {
            assertFalse(libHarness.isBaseMessage(payload), "!isBaseMessage: when valid");
            vm.expectRevert("Not a base message");
            libHarness.castToBaseMessage(payload);
        }
    }
}
