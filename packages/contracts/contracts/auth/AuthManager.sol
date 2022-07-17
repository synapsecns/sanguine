// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { Attestation } from "../libs/Attestation.sol";
import { Report } from "../libs/Report.sol";
import { Auth } from "../libs/Auth.sol";

abstract contract AuthManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              LIBRARIES                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    using Attestation for bytes29;
    using Report for bytes;
    using Report for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[50] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Checks if the passed payload is a valid Attestation message,
     *          if the signature is valid and if the signer is an authorized updater.
     * @param _updater      Signer of the message, needs to be authorized as updater, revert otherwise.
     * @param _attestation  Attestation of Home merkle root. Needs to be valid, revert otherwise.
     * @return _view        Memory view on attestation
     */
    function _checkUpdaterAuth(address _updater, bytes memory _attestation)
        internal
        view
        returns (bytes29 _view)
    {
        _view = _attestation.ref(0);
        _checkUpdaterAuth(_updater, _view);
    }

    function _checkUpdaterAuth(address _updater, bytes29 _view) internal view {
        require(_view.isAttestation(), "Not an attestation");
        // This will revert if signature is invalid
        Auth.checkSignature(
            _updater,
            _view.attestationData(),
            _view.attestationSignature().clone()
        );
        require(_isUpdater(_view.attestationDomain(), _updater), "Signer is not an updater");
    }

    function _checkWatchtowerAuth(address _watchtower, bytes memory _report)
        internal
        view
        returns (bytes29 _view)
    {
        _view = _report.castToReport();
        _checkWatchtowerAuth(_watchtower, _view);
    }

    function _checkWatchtowerAuth(address _watchtower, bytes29 _view) internal view {
        require(_view.isReport(), "Not a report");
        Auth.checkSignature(_watchtower, _view.reportData(), _view.reportSignature().clone());
        require(_isWatchtower(_watchtower), "Signer is not a watchtower");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isUpdater(uint32 _homeDomain, address _updater) internal view virtual returns (bool);

    function _isWatchtower(address _watchtower) internal view virtual returns (bool);
}
