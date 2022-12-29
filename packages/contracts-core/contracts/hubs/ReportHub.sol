// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../libs/Report.sol";
import { Auth } from "../libs/Auth.sol";
import { AttestationHub } from "./AttestationHub.sol";

/**
 * @notice Keeps track of the agents and verifies signed reports.
 */
abstract contract ReportHub is AttestationHub {
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
        /* TODO(Chi): enable reports once co-signed Attestation is implemented
        // Verify the report signature and recover an active guard address
        bytes29 reportView = _report.castToReport();
        address guard = _verifyReport(reportView);
        // Verify the attestation signature and recover an active notary address
        bytes29 attestationView = reportView.reportedAttestation();
        address notary = _verifyAttestation(attestationView);
        return _handleReport(guard, notary, attestationView, reportView, _report);
        */
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

    /**
     * @notice Checks if report signer is authorized.
     * @dev Signer needs to be an active Guard.
     * @param _reportView  Memory view over the Report to check
     * @return guard Address of the report signer
     */
    function _verifyReport(bytes29 _reportView) internal view returns (address guard) {
        /* TODO(Chi): enable reports once co-signed Attestation is implemented
        // Check if Report payload is properly formatted.
        require(_reportView.isReport(), "Not a report");
        bytes32 digest = Auth.toEthSignedMessageHash(_reportView.reportData());
        // Check if Guard signature is valid.
        guard = _checkAgentAuth({
            _domain: 0,
            _digest: digest,
            _signatureView: _reportView.guardSignature()
        });
        */
    }
}
