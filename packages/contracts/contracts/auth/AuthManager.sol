// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { Attestation } from "../libs/Attestation.sol";
import { Auth } from "../libs/Auth.sol";

abstract contract AuthManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              LIBRARIES                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    using Attestation for bytes29;
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
     * @param _attestation  Attestation of Home merkle root. Needs to be valid, revert otherwise.
     * @return _updater     Updater that signed the Attestation
     * @return _view        Memory view on attestation
     */
    function _checkUpdaterAuth(bytes memory _attestation)
        internal
        view
        returns (address _updater, bytes29 _view)
    {
        _view = _attestation.ref(0);
        require(_view.isAttestation(), "Not an attestation");
        _updater = Auth.recoverSigner(
            _view.attestationData(),
            _view.attestationSignature().clone()
        );
        require(_isUpdater(_view.attestationDomain(), _updater), "Signer is not an updater");
    }

    function _checkWatchtowerAuth(bytes memory _report)
        internal
        view
        returns (address _watchtower, bytes29 _data)
    {
        // TODO: check if _report is valid, once watchtower message standard is finalized
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isUpdater(uint32 _homeDomain, address _updater) internal view virtual returns (bool);

    function _isWatchtower(address _watchtower) internal view virtual returns (bool);
}
