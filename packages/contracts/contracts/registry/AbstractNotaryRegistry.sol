// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { Attestation } from "../libs/Attestation.sol";
import { Auth } from "../libs/Auth.sol";

abstract contract AbstractNotaryRegistry {
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @notice  Checks if the passed payload is a valid Attestation message,
     *          if the signature is valid and if the signer is an authorized notary.
     * @param _attestation  Attestation of Home merkle root. Needs to be valid, revert otherwise.
     * @return _notary     Notary that signed the Attestation
     * @return _view        Memory view on attestation
     */
    function _checkNotaryAuth(bytes memory _attestation)
        internal
        view
        returns (address _notary, bytes29 _view)
    {
        _view = _attestation.ref(0);
        require(_view.isAttestation(), "Not an attestation");
        _notary = Auth.recoverSigner(_view.attestationData(), _view.attestationSignature().clone());
        require(_isNotary(_view.attestationDomain(), _notary), "Signer is not a notary");
    }

    function _isNotary(uint32 _homeDomain, address _notary) internal view virtual returns (bool);
}
