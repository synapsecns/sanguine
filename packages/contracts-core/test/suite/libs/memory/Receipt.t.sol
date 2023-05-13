// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {UnformattedReceipt} from "../../../../contracts/libs/Errors.sol";
import {RECEIPT_LENGTH} from "../../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {ReceiptHarness} from "../../../harnesses/libs/memory/ReceiptHarness.t.sol";

import {RawExecReceipt} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract ReceiptLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    ReceiptHarness internal libHarness;

    function setUp() public {
        libHarness = new ReceiptHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatReceipt(RawExecReceipt memory re) public {
        // Test formatting
        bytes memory payload = libHarness.formatReceipt(
            re.origin,
            re.destination,
            re.messageHash,
            re.snapshotRoot,
            re.stateIndex,
            re.attNotary,
            re.firstExecutor,
            re.finalExecutor
        );
        assertEq(
            payload,
            abi.encodePacked(
                re.origin,
                re.destination,
                re.messageHash,
                re.snapshotRoot,
                re.stateIndex,
                re.attNotary,
                re.firstExecutor,
                re.finalExecutor
            ),
            "!formatReceipt"
        );
        // Test formatting checker
        checkCastToReceipt({payload: payload, isReceipt: true});
        // Test getters
        assertEq(libHarness.origin(payload), re.origin, "!origin");
        assertEq(libHarness.destination(payload), re.destination, "!destination");
        assertEq(libHarness.messageHash(payload), re.messageHash, "!messageHash");
        assertEq(libHarness.snapshotRoot(payload), re.snapshotRoot, "!snapshotRoot");
        assertEq(libHarness.stateIndex(payload), re.stateIndex, "!stateIndex");
        assertEq(libHarness.attNotary(payload), re.attNotary, "!attNotary");
        assertEq(libHarness.firstExecutor(payload), re.firstExecutor, "!firstExecutor");
        assertEq(libHarness.finalExecutor(payload), re.finalExecutor, "!finalExecutor");
        // Test hashing of "valid receipt"
        bytes32 receiptSalt = keccak256("RECEIPT_VALID_SALT");
        bytes32 hashedReceipt = keccak256(abi.encodePacked(receiptSalt, keccak256(payload)));
        assertEq(libHarness.hashValid(payload), hashedReceipt, "!hashValid");
        // Test hashing of "invalid receipt body"
        bytes32 receiptBodyInvalidSalt = keccak256("RECEIPT_INVALID_SALT");
        hashedReceipt = keccak256(abi.encodePacked(receiptBodyInvalidSalt, keccak256(payload)));
        assertEq(libHarness.hashInvalid(payload), hashedReceipt, "!hashInvalid");
    }

    function test_isReceipt(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToReceipt({payload: payload, isReceipt: length == RECEIPT_LENGTH});
    }

    function rest_equals_same(RawExecReceipt memory re) public {
        assertTrue(libHarness.equals(re.formatReceipt(), re.formatReceipt()));
    }

    function test_equals_different(RawExecReceipt memory re, uint256 mask) public {
        RawExecReceipt memory mrb = re.modifyReceipt(mask);
        assertFalse(libHarness.equals(re.formatReceipt(), mrb.formatReceipt()));
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToReceipt(bytes memory payload, bool isReceipt) public {
        if (isReceipt) {
            assertTrue(libHarness.isReceipt(payload), "!isReceipt: when valid");
            assertEq(libHarness.castToReceipt(payload), payload, "!castToReceipt: when valid");
        } else {
            assertFalse(libHarness.isReceipt(payload), "!isReceipt: when valid");
            vm.expectRevert(UnformattedReceipt.selector);
            libHarness.castToReceipt(payload);
        }
    }
}
