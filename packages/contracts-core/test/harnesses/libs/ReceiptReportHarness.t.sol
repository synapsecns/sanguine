// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    ReceiptFlag,
    ReceiptReport,
    ReceiptReportLib,
    MemView,
    MemViewLib
} from "../../../contracts/libs/ReceiptReport.sol";

// solhint-disable ordering
/// @notice Exposes Report methods for testing against golang.
contract ReceiptReportHarness {
    using ReceiptReportLib for bytes;
    using ReceiptReportLib for MemView;
    using MemViewLib for bytes;

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToReceiptReport(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        ReceiptReport receiptReport = ReceiptReportLib.castToReceiptReport(payload);
        return receiptReport.unwrap().clone();
    }

    function hash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToReceiptReport().hash();
    }

    function flag(bytes memory payload) public pure returns (ReceiptFlag) {
        return payload.castToReceiptReport().flag();
    }

    function receiptBody(bytes memory payload) public view returns (bytes memory) {
        return payload.castToReceiptReport().receiptBody().unwrap().clone();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatReceiptReport(ReceiptFlag flag_, bytes memory rcptBodyPayload) public pure returns (bytes memory) {
        return ReceiptReportLib.formatReceiptReport(flag_, rcptBodyPayload);
    }

    function isReceiptReport(bytes memory payload) public pure returns (bool) {
        return payload.ref().isReceiptReport();
    }
}
