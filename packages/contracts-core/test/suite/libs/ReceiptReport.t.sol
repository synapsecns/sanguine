// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RECEIPT_BODY_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {RawReceiptReport} from "../../utils/libs/SynapseStructs.t.sol";

import {ReceiptFlag, ReceiptReportHarness} from "../../harnesses/libs/ReceiptReportHarness.t.sol";

// solhint-disable func-name-mixedcase
contract ReceiptReportLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    ReceiptReportHarness internal libHarness;

    function setUp() public {
        libHarness = new ReceiptReportHarness();
    }

    function test_formatReceiptReport(RawReceiptReport memory rawSR) public {
        // Make sure flag fits into enum
        rawSR.flag = uint8(bound(rawSR.flag, 0, uint8(type(ReceiptFlag).max)));
        // This is tested in ReceiptLibraryTest, we assume it's working here
        bytes memory rcptBodyPayload = rawSR.body.formatReceiptBody();
        bytes memory payload = libHarness.formatReceiptReport(ReceiptFlag(rawSR.flag), rcptBodyPayload);
        assertEq(payload, abi.encodePacked(rawSR.flag, rcptBodyPayload), "!formatReceiptReport");
        checkCastToReceiptReport({payload: payload, isReceiptReport: true});
        // Check getters
        assertEq(uint8(libHarness.flag(payload)), rawSR.flag, "!flag");
        assertEq(libHarness.receiptBody(payload), rcptBodyPayload, "!rcptBodyPayload");
        // Test hashing
        bytes32 receiptReportSalt = keccak256("RECEIPT_REPORT_SALT");
        bytes32 hashedReceiptReport = keccak256(abi.encodePacked(receiptReportSalt, keccak256(payload)));
        assertEq(libHarness.hash(payload), hashedReceiptReport, "!hash");
    }

    function test_isReceiptReport(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToReceiptReport({payload: payload, isReceiptReport: length == 1 + RECEIPT_BODY_LENGTH});
    }

    function test_isReceiptReport_flagOutOfRange(uint8 flag) public {
        flag = uint8(bound(flag, uint8(type(ReceiptFlag).max) + 1, type(uint8).max));
        bytes memory payload = abi.encodePacked(flag, new bytes(RECEIPT_BODY_LENGTH));
        checkCastToReceiptReport({payload: payload, isReceiptReport: false});
    }

    function checkCastToReceiptReport(bytes memory payload, bool isReceiptReport) public {
        if (isReceiptReport) {
            assertTrue(libHarness.isReceiptReport(payload), "!isReceiptReport: when valid");
            assertEq(libHarness.castToReceiptReport(payload), payload, "!castToReceiptReport: when valid");
        } else {
            assertFalse(libHarness.isReceiptReport(payload), "!isReceiptReport: when valid");
            vm.expectRevert("Not a receipt report");
            libHarness.castToReceiptReport(payload);
        }
    }
}
