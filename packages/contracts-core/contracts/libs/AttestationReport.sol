// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Attestation, AttestationLib } from "./Attestation.sol";
import { ByteString } from "./ByteString.sol";
import { ATTESTATION_REPORT_SALT } from "./Constants.sol";
import { TypedMemView } from "./TypedMemView.sol";

/// @dev AttestationReport is a memory view over a formatted Guard report for invalid Attestation
type AttestationReport is bytes29;
/// @dev Possible flags for the AttestationReport
/// Currently has only one possible value, but enables different types of reports in the future
enum AttestationFlag {
    Invalid
}
/// @dev Attach library functions to AttestationFlag
using { AttestationReportLib.formatAttestationReport } for AttestationFlag global;
/// @dev Attach library functions to AttestationReport
using {
    AttestationReportLib.hash,
    AttestationReportLib.unwrap,
    AttestationReportLib.flag,
    AttestationReportLib.attestation
} for AttestationReport global;

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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ATTESTATION REPORT                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a formatted AttestationReport payload with provided fields.
    /// @param _flag        Flag signalling type of Attestation Report
    /// @param _attPayload  Raw payload with reported attestation
    /// @return Formatted attestation report
    function formatAttestationReport(AttestationFlag _flag, bytes memory _attPayload)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_flag, _attPayload);
    }

    /// @notice Returns an AttestationReport view over the given payload
    /// @dev Will revert if the payload is not an attestation report.
    function castToAttestationReport(bytes memory _payload)
        internal
        pure
        returns (AttestationReport)
    {
        return castToAttestationReport(_payload.castToRawBytes());
    }

    /// @notice Casts a memory view to an AttestationReport view.
    /// @dev Will revert if if the memory view is not over an AttestationReport payload.
    function castToAttestationReport(bytes29 _view) internal pure returns (AttestationReport) {
        require(isAttestationReport(_view), "Not an attestation report");
        return AttestationReport.wrap(_view);
    }

    /// @notice Checks that a payload is a formatted AttestationReport.
    function isAttestationReport(bytes29 _view) internal pure returns (bool) {
        // Flag needs to exist
        if (_view.len() < OFFSET_ATTESTATION) return false;
        // Flag should fit into AttestationFlag enum
        if (_arFlag(_view) > uint8(type(AttestationFlag).max)) return false;
        // Attestation should be properly formatted
        return _arAttestation(_view).isAttestation();
    }

    function hash(AttestationReport _attReport) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 _view = _attReport.unwrap();
        // The final hash to sign is keccak(attestationReportSalt, keccak(attestationReport))
        return keccak256(bytes.concat(ATTESTATION_REPORT_SALT, _view.keccak()));
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(AttestationReport _attReport) internal pure returns (bytes29) {
        return AttestationReport.unwrap(_attReport);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      ATTESTATION REPORT SLICING                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns AttestationFlag used in the report.
    function flag(AttestationReport _attReport) internal pure returns (AttestationFlag) {
        bytes29 _view = _attReport.unwrap();
        // We check that flag fits into enum, when payload is wrapped
        // into AttestationReport, so this never reverts
        return AttestationFlag(_arFlag(_view));
    }

    /// @notice Returns typed memory view over attestation used in the report.
    function attestation(AttestationReport _attReport) internal pure returns (Attestation) {
        bytes29 _view = _attReport.unwrap();
        // We check that attestation is properly formatted, when payload is wrapped
        // into AttestationReport, so this never reverts.
        return _arAttestation(_view).castToAttestation();
    }

    /// @dev Returns AttestationReport without checking that it fits into AttestationFlag enum.
    function _arFlag(bytes29 _view) internal pure returns (uint8) {
        return uint8(_view.indexUint({ _index: OFFSET_FLAG, _bytes: 1 }));
    }

    /// @dev Returns an untyped memory view over Report's attestation
    /// without checking if it is properly formatted.
    function _arAttestation(bytes29 _view) internal pure returns (bytes29) {
        return _view.sliceFrom({ _index: OFFSET_ATTESTATION, newType: 0 });
    }
}
