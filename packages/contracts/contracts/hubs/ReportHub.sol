// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

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
     *      - Report signer is not a Guard.
     *      - Reported notary is not a Notary.
     * @param _report	Payload with Report data and signature
     * @return TRUE if Report was handled correctly.
     */
    function submitReport(bytes memory _report) external returns (bool) {
        // Check if real Guard & signature
        (address _guard, bytes29 _reportView) = _checkGuardAuth(_report);
        // Check if fraud flag is supported,
        // e.g. some contracts might want to accept only Fraud Reports
        _checkFraudFlag(_reportView.reportedFraud());
        bytes29 _attestationView = _reportView.reportedAttestation();
        // Check if real Notary & signature
        address _notary = _checkNotaryAuth(_attestationView);
        return _handleReport(_guard, _notary, _attestationView, _report);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Implement logic for checking the fraud flag in the child contracts.
     * Note: contract is supposed to revert if the fraud flag is not supported.
     */
    function _checkFraudFlag(bool _flag) internal virtual;

    /**
     * @dev Implement logic for handling the Report in the child contracts.
     * @param _guard            Guard address (signature&role already verified)
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _report           Payload with Report data and signature
     * @return TRUE if Report was handled correctly.
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes memory _report
    ) internal virtual returns (bool);
}
