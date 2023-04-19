// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RECEIPT_REPORT_SALT} from "./Constants.sol";
import {ReceiptBody, ReceiptLib} from "./Receipt.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// ReceiptReport is a memory view over a formatted Guard report for invalid ReceiptBody.
type ReceiptReport is uint256;

/// Possible flags for the ReceiptReport.
/// - Currently has only one possible value, but enables different types of reports in the future.
enum ReceiptFlag {Invalid}

using ReceiptReportLib for ReceiptFlag global;
using ReceiptReportLib for ReceiptReport global;

/// @notice ReceiptReport structure represents a Guard statement that a Receipt Body is invalid.
/// Receipt Body is considered invalid, if it doesn't match the saved receipt body in ExecutionHub contract
/// with the same message hash (or the message hash doesn't exist).
/// # Memory layout of ReceiptReport
///
/// | Position   | Field       | Type  | Bytes | Description                               |
/// | ---------- | ----------- | ----- | ----- | ----------------------------------------- |
/// | [000..001) | flag        | uint8 | 1     | ReceiptFlag for the report                |
/// | [001..134) | receiptBody | bytes | 133   | Raw payload for the reported Receipt Body |
///
/// @dev Signed ReceiptReport together with a proof that Notary used the reported Receipt Body for their signed statement,
/// could be used on Destination and Summit to initiate a Dispute between the Guard and the Notary.
library ReceiptReportLib {
    using MemViewLib for bytes;
    using ReceiptLib for MemView;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_FLAG = 0;
    uint256 private constant OFFSET_RECEIPT_BODY = 1;

    // ═══════════════════════════════════════════════ RECEIPT REPORT ════════════════════════════════════════════════════

    /// @notice Returns a formatted ReceiptReport payload with provided fields.
    /// @param flag_            Flag signalling type of Receipt Report
    /// @param rcptBodyPayload  Raw payload with reported receipt body
    /// @return Formatted receipt report
    function formatReceiptReport(ReceiptFlag flag_, bytes memory rcptBodyPayload)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(flag_, rcptBodyPayload);
    }

    /// @notice Returns a ReceiptReport view over the given payload
    /// @dev Will revert if the payload is not a receipt report.
    function castToReceiptReport(bytes memory payload) internal pure returns (ReceiptReport) {
        return castToReceiptReport(payload.ref());
    }

    /// @notice Casts a memory view to a ReceiptReport view.
    /// @dev Will revert if if the memory view is not over a ReceiptReport payload.
    function castToReceiptReport(MemView memView) internal pure returns (ReceiptReport) {
        require(isReceiptReport(memView), "Not a receipt report");
        return ReceiptReport.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted ReceiptReport.
    function isReceiptReport(MemView memView) internal pure returns (bool) {
        // Flag needs to exist
        if (memView.len() < OFFSET_RECEIPT_BODY) return false;
        // Flag should fit into ReceiptFlag enum
        if (_rrFlag(memView) > uint8(type(ReceiptFlag).max)) return false;
        // Receipt Body should be properly formatted
        return _rrReceiptBody(memView).isReceiptBody();
    }

    /// @notice Returns the hash of a ReceiptReport, that could be later signed by a Guard.
    function hash(ReceiptReport receiptReport) internal pure returns (bytes32) {
        // The final hash to sign is keccak(receiptReportSalt, keccak(receiptReport))
        return receiptReport.unwrap().keccakSalted(RECEIPT_REPORT_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(ReceiptReport receiptReport) internal pure returns (MemView) {
        return MemView.wrap(ReceiptReport.unwrap(receiptReport));
    }

    // ═══════════════════════════════════════════ RECEIPT REPORT SLICING ════════════════════════════════════════════════

    /// @notice Returns ReceiptFlag used in the report.
    function flag(ReceiptReport receiptReport) internal pure returns (ReceiptFlag) {
        // We check that flag fits into enum, when payload is wrapped
        // into ReceiptReport, so this never reverts
        return ReceiptFlag(_rrFlag(receiptReport.unwrap()));
    }

    /// @notice Returns typed memory view over the receipt body used in the report.
    function receiptBody(ReceiptReport receiptReport) internal pure returns (ReceiptBody) {
        // We check that receipt body is properly formatted, when payload is wrapped
        // into ReceiptReport, so this never reverts.
        return _rrReceiptBody(receiptReport.unwrap()).castToReceiptBody();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns ReceiptReport flag without checking that it fits into ReceiptFlag enum.
    function _rrFlag(MemView memView) internal pure returns (uint8) {
        return uint8(memView.indexUint({index_: OFFSET_FLAG, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over Report's receipt body without checking if it is properly formatted.
    function _rrReceiptBody(MemView memView) internal pure returns (MemView) {
        return memView.sliceFrom({index_: OFFSET_RECEIPT_BODY});
    }
}
