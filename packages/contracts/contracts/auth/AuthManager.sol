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
     *          if the signature is valid and if the signer is an authorized notary.
     * @param _notary       Signer of the message, needs to be authorized as notary, revert otherwise.
     * @param _attestation  Attestation of Home merkle root. Needs to be valid, revert otherwise.
     * @return _view        Memory view on attestation
     */
    function _checkNotaryAuth(address _notary, bytes memory _attestation)
        internal
        view
        returns (bytes29 _view)
    {
        _view = _attestation.ref(0);
        require(_view.isAttestation(), "Not an attestation");
        // This will revert if signature is invalid
        Auth.checkSignature(_notary, _view.attestationData(), _view.attestationSignature().clone());
        require(_isNotary(_view.attestationDomain(), _notary), "Signer is not an notary");
    }

    function _checkWatchtowerAuth(address _watchtower, bytes memory _report)
        internal
        view
        returns (bytes29 _data)
    {
        // TODO: check if _report is valid, once watchtower message standard is finalized
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isNotary(uint32 _homeDomain, address _notary) internal view virtual returns (bool);

    function _isWatchtower(address _watchtower) internal view virtual returns (bool);
}
