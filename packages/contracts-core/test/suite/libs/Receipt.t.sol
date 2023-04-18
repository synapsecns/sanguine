// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RECEIPT_BODY_LENGTH, RECEIPT_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {ReceiptHarness} from "../../harnesses/libs/ReceiptHarness.t.sol";

import {RawExecReceipt, RawReceiptBody} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract ReceiptLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    ReceiptHarness internal libHarness;

    function setUp() public {
        libHarness = new ReceiptHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatReceiptBody(RawReceiptBody memory rrb) public {
        // Test formatting
        bytes memory payload = libHarness.formatReceiptBody(
            rrb.origin,
            rrb.destination,
            rrb.messageHash,
            rrb.snapshotRoot,
            rrb.stateIndex,
            rrb.attNotary,
            rrb.firstExecutor,
            rrb.finalExecutor
        );
        assertEq(
            payload,
            abi.encodePacked(
                rrb.origin,
                rrb.destination,
                rrb.messageHash,
                rrb.snapshotRoot,
                rrb.stateIndex,
                rrb.attNotary,
                rrb.firstExecutor,
                rrb.finalExecutor
            ),
            "!formatReceiptBody"
        );
        // Test formatting checker
        checkCastToReceiptBody({payload: payload, isReceiptBody: true});
        // Test getters
        assertEq(libHarness.origin(payload), rrb.origin, "!origin");
        assertEq(libHarness.destination(payload), rrb.destination, "!destination");
        assertEq(libHarness.messageHash(payload), rrb.messageHash, "!messageHash");
        assertEq(libHarness.snapshotRoot(payload), rrb.snapshotRoot, "!snapshotRoot");
        assertEq(libHarness.stateIndex(payload), rrb.stateIndex, "!stateIndex");
        assertEq(libHarness.attNotary(payload), rrb.attNotary, "!attNotary");
        assertEq(libHarness.firstExecutor(payload), rrb.firstExecutor, "!firstExecutor");
        assertEq(libHarness.finalExecutor(payload), rrb.finalExecutor, "!finalExecutor");
    }

    function test_formatReceipt(RawExecReceipt memory re) public {
        bytes memory bodyPayload = re.body.formatReceiptBody();
        bytes memory tipsPayload = re.tips.formatTips();
        // Test formatting
        bytes memory payload = libHarness.formatReceipt(bodyPayload, tipsPayload);
        assertEq(payload, abi.encodePacked(bodyPayload, tipsPayload), "!formatReceipt");
        // Test formatting checker
        checkCastToReceipt({payload: payload, isReceipt: true});
        // Test getters
        assertEq(libHarness.body(payload), bodyPayload, "!bodyPayload");
        assertEq(libHarness.tips(payload), tipsPayload, "!tipsPayload");
    }

    function test_isReceiptBody(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToReceiptBody({payload: payload, isReceiptBody: length == RECEIPT_BODY_LENGTH});
    }

    function test_isReceipt(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToReceipt({payload: payload, isReceipt: length == RECEIPT_LENGTH});
    }

    function rest_equals_same(RawReceiptBody memory rrb) public {
        assertTrue(libHarness.equals(rrb.formatReceiptBody(), rrb.formatReceiptBody()));
    }

    function test_equals_different(RawReceiptBody memory rrb, uint256 mask) public {
        RawReceiptBody memory mrb = rrb.modifyReceiptBody(mask);
        assertFalse(libHarness.equals(rrb.formatReceiptBody(), mrb.formatReceiptBody()));
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToReceiptBody(bytes memory payload, bool isReceiptBody) public {
        if (isReceiptBody) {
            assertTrue(libHarness.isReceiptBody(payload), "!isReceiptBody: when valid");
            assertEq(libHarness.castToReceiptBody(payload), payload, "!castToReceiptBody: when valid");
        } else {
            assertFalse(libHarness.isReceiptBody(payload), "!isReceiptBody: when valid");
            vm.expectRevert("Not a receipt body");
            libHarness.castToReceiptBody(payload);
        }
    }

    function checkCastToReceipt(bytes memory payload, bool isReceipt) public {
        if (isReceipt) {
            assertTrue(libHarness.isReceipt(payload), "!isReceipt: when valid");
            assertEq(libHarness.castToReceipt(payload), payload, "!castToReceipt: when valid");
        } else {
            assertFalse(libHarness.isReceipt(payload), "!isReceipt: when valid");
            vm.expectRevert("Not a receipt");
            libHarness.castToReceipt(payload);
        }
    }
}
