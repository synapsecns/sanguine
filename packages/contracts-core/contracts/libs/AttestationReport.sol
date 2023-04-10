// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Attestation, AttestationLib} from "./Attestation.sol";
import {ByteString} from "./ByteString.sol";
import {ATTESTATION_REPORT_SALT} from "./Constants.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev AttestationReport is a memory view over a formatted Guard report for invalid Attestation
type AttestationReport is bytes29;

/// @dev Possible flags for the AttestationReport
/// Currently has only one possible value, but enables different types of reports in the future
enum AttestationFlag {Invalid}

/// @dev Attach library functions to AttestationFlag
using AttestationReportLib for AttestationFlag global;
/// @dev Attach library functions to AttestationReport
using AttestationReportLib for AttestationReport global;

library AttestationReportLib {
    using AttestationLib for bytes29;
    using ByteString for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev AttestationReport structure represents a Guard statement
     * that a given Attestation is invalid. Attestation is considered invalid, if it doesn't match
     * the saved attestation in Summit contract with the same nonce (or if nonce doesn't exist).
     *
     * Signed AttestationReport together with Notary signature for the reported Attestation
     * could be used on Destination to initiate a Dispute between the Guard and the Notary.
     *
     * @dev Memory layout of AttestationReport fields:
     * [000 .. 001): flag           uint8    1 byte     AttestationFlag for the report
     * [001 .. 048): attestation    bytes   47 bytes    Raw payload with attestation
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_FLAG = 0;
    uint256 private constant OFFSET_ATTESTATION = 1;

    // ════════════════════════════════════════════ ATTESTATION REPORT ═════════════════════════════════════════════════

    /// @notice Returns a formatted AttestationReport payload with provided fields.
    /// @param flag_        Flag signalling type of Attestation Report
    /// @param attPayload   Raw payload with reported attestation
    /// @return Formatted attestation report
    function formatAttestationReport(AttestationFlag flag_, bytes memory attPayload)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(flag_, attPayload);
    }

    /// @notice Returns an AttestationReport view over the given payload
    /// @dev Will revert if the payload is not an attestation report.
    function castToAttestationReport(bytes memory payload) internal pure returns (AttestationReport) {
        return castToAttestationReport(payload.castToRawBytes());
    }

    /// @notice Casts a memory view to an AttestationReport view.
    /// @dev Will revert if if the memory view is not over an AttestationReport payload.
    function castToAttestationReport(bytes29 view_) internal pure returns (AttestationReport) {
        require(isAttestationReport(view_), "Not an attestation report");
        return AttestationReport.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted AttestationReport.
    function isAttestationReport(bytes29 view_) internal pure returns (bool) {
        // Flag needs to exist
        if (view_.len() < OFFSET_ATTESTATION) return false;
        // Flag should fit into AttestationFlag enum
        if (_arFlag(view_) > uint8(type(AttestationFlag).max)) return false;
        // Attestation should be properly formatted
        return _arAttestation(view_).isAttestation();
    }

    function hash(AttestationReport attReport) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 view_ = attReport.unwrap();
        // The final hash to sign is keccak(attestationReportSalt, keccak(attestationReport))
        return keccak256(bytes.concat(ATTESTATION_REPORT_SALT, view_.keccak()));
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(AttestationReport attReport) internal pure returns (bytes29) {
        return AttestationReport.unwrap(attReport);
    }

    // ════════════════════════════════════════ ATTESTATION REPORT SLICING ═════════════════════════════════════════════

    /// @notice Returns AttestationFlag used in the report.
    function flag(AttestationReport attReport) internal pure returns (AttestationFlag) {
        bytes29 view_ = attReport.unwrap();
        // We check that flag fits into enum, when payload is wrapped
        // into AttestationReport, so this never reverts
        return AttestationFlag(_arFlag(view_));
    }

    /// @notice Returns typed memory view over attestation used in the report.
    function attestation(AttestationReport attReport) internal pure returns (Attestation) {
        bytes29 view_ = attReport.unwrap();
        // We check that attestation is properly formatted, when payload is wrapped
        // into AttestationReport, so this never reverts.
        return _arAttestation(view_).castToAttestation();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns AttestationReport without checking that it fits into AttestationFlag enum.
    function _arFlag(bytes29 view_) internal pure returns (uint8) {
        return uint8(view_.indexUint({index_: OFFSET_FLAG, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over Report's attestation
    /// without checking if it is properly formatted.
    function _arAttestation(bytes29 view_) internal pure returns (bytes29) {
        return view_.sliceFrom({index_: OFFSET_ATTESTATION, newType: 0});
    }
}
