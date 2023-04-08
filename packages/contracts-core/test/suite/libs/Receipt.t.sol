// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RECEIPT_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {ReceiptHarness} from "../../harnesses/libs/ReceiptHarness.t.sol";

import {RawExecReceipt} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract ReceiptLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    ReceiptHarness internal libHarness;

    function setUp() public {
        libHarness = new ReceiptHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatReceipt(RawExecReceipt memory re) public {
        bytes memory tipsPayload = re.tips.formatTips();
        // Test formatting
        bytes memory payload = libHarness.formatReceipt(
            re.origin, re.destination, re.messageHash, re.snapshotRoot, re.firstExecutor, re.finalExecutor, tipsPayload
        );
        assertEq(
            payload,
            abi.encodePacked(
                re.origin,
                re.destination,
                re.messageHash,
                re.snapshotRoot,
                re.firstExecutor,
                re.finalExecutor,
                tipsPayload
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
        assertEq(libHarness.firstExecutor(payload), re.firstExecutor, "!firstExecutor");
        assertEq(libHarness.finalExecutor(payload), re.finalExecutor, "!finalExecutor");
        assertEq(libHarness.tips(payload), tipsPayload, "!tipsPayload");
    }

    function test_isReceipt(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToReceipt({payload: payload, isReceipt: length == RECEIPT_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

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
