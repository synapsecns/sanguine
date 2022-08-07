// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { Report } from "../libs/Report.sol";
import { AuthManager } from "./AuthManager.sol";

abstract contract ReportHub is AuthManager {
    using Report for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[50] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Called by the external agent. Submits the signed report for the contract to handle it.
     * @dev Reverts if report signer is not a watchtower, or if reported updater is not an updater.
     * @param _report       Report data and signature
     */
    function submitReport(bytes memory _report) external returns (bool) {
        // Check if real Guard & signature
        (address _guard, bytes29 _reportView) = _checkWatchtowerAuth(_report);
        // Check if this is a fraud report
        require(_reportView.reportIsFraud(), "!fraud");
        bytes29 _attestationView = _reportView.reportAttestation();
        // Check if real Notary & signature
        address _notary = _checkUpdaterAuth(_attestationView);
        return _handleReport(_guard, _notary, _attestationView, _report);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes memory _report
    ) internal virtual returns (bool);
}
