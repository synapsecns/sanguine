// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {UnformattedBaseMessage} from "../../../../contracts/libs/Errors.sol";
import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {BaseMessageHarness} from "../../../harnesses/libs/memory/BaseMessageHarness.t.sol";

import {RawBaseMessage, Request, Tips} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract BaseMessageLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    BaseMessageHarness internal libHarness;

    function setUp() public {
        libHarness = new BaseMessageHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatBaseMessage(RawBaseMessage memory rbm) public {
        Tips tips = rbm.tips.castToTips();
        uint256 encodedTips = rbm.tips.encodeTips();
        Request request = rbm.request.castToRequest();
        uint192 encodedRequest = rbm.request.encodeRequest();
        // Test formatting
        bytes memory payload = libHarness.formatBaseMessage(tips, rbm.sender, rbm.recipient, request, rbm.content);
        assertEq(
            payload,
            abi.encodePacked(encodedTips, rbm.sender, rbm.recipient, encodedRequest, rbm.content),
            "!formatBaseMessage"
        );
        // Test formatting checker
        checkCastToBaseMessage({payload: payload, isBaseMessage: true});
        // Test getters
        assertEq(libHarness.sender(payload), rbm.sender, "!sender");
        assertEq(libHarness.recipient(payload), rbm.recipient, "!recipient");
        assertEq(libHarness.tips(payload), encodedTips, "!tips");
        assertEq(libHarness.request(payload), encodedRequest, "!request");
        assertEq(libHarness.content(payload), rbm.content, "!content");
        // Test hashing
        bytes32 leftChild = keccak256(abi.encodePacked(encodedTips));
        bytes32 rightChild = keccak256(abi.encodePacked(rbm.sender, rbm.recipient, encodedRequest, rbm.content));
        assertEq(libHarness.bodyLeaf(payload), rightChild, "!bodyLeaf");
        assertEq(libHarness.leaf(payload), keccak256(abi.encodePacked(leftChild, rightChild)), "!leaf");
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
            vm.expectRevert(UnformattedBaseMessage.selector);
            libHarness.castToBaseMessage(payload);
        }
    }
}
