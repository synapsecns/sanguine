// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Report } from "../libs/Report.sol";
import { AbstractGuardRegistry } from "../registry/AbstractGuardRegistry.sol";
import { AbstractNotaryRegistry } from "../registry/AbstractNotaryRegistry.sol";

abstract contract ReportHub is AbstractGuardRegistry, AbstractNotaryRegistry {
    using Report for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Called by the external agent. Submits the signed report for handling.
     * @dev Reverts if either of this is true:
     *      - Report payload is not properly formatted.
     *      - Report signer is not a Guard.
     *      - Reported notary is not a Notary.
     * @param _report	Payload with Report data and signature
     * @return TRUE if Report was handled correctly.
     */
    function submitReport(bytes memory _report) external returns (bool) {
        // Check if real Guard & signature.
        // This also checks if Report payload is properly formatted.
        (address _guard, bytes29 _reportView) = _checkGuardAuth(_report);
        bytes29 _attestationView = _reportView.reportedAttestation();
        // Check if real Notary & signature.
        // This also checks if Attestation payload is properly formatted,
        // though it's already been checked in _checkGuardAuth(_report) [see Report.sol].
        address _notary = _checkNotaryAuth(_attestationView);
        // Pass _reportView as the existing bytes29 pointer to report payload.
        // Pass _report to avoid extra memory copy when emitting report payload.
        return _handleReport(_guard, _notary, _attestationView, _reportView, _report);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Implement logic for handling the Report in the child contracts.
     * Note: Report can have either Valid or Fraud flag, make sure to check that.
     * @param _guard            Guard address (signature&role already verified)
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _reportView       Memory view over Report for convenience
     * @param _report           Payload with Report data and signature
     * @return TRUE if Report was handled correctly.
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal virtual returns (bool);
}
