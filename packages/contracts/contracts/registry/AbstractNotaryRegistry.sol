// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { Attestation } from "../libs/Attestation.sol";
import { Auth } from "../libs/Auth.sol";

abstract contract AbstractNotaryRegistry {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @notice  Checks all following statements are true:
     *          - `_attestation` is a formatted Attestation payload
     *          - `_attestation` contains a signature
     *          - such signature belongs to an authorized Notary
     * @param _attestation  Attestation of Origin merkle root
     * @return _notary      Notary that signed the Attestation
     * @return _view        Memory view on attestation
     */
    function _checkNotaryAuth(bytes memory _attestation)
        internal
        view
        returns (address _notary, bytes29 _view)
    {
        _view = _attestation.castToAttestation();
        _notary = _checkNotaryAuth(_view);
    }

    /**
     * @notice  Checks all following statements are true:
     *          - `_view` is a memory view on a formatted Attestation payload
     *          - `_view` contains a signature
     *          - such signature belongs to an authorized Notary
     * @param _view     Memory view on Attestation of Origin merkle root
     * @return _notary  Notary that signed the Attestation
     */
    function _checkNotaryAuth(bytes29 _view) internal view returns (address _notary) {
        require(_view.isAttestation(), "Not an attestation");
        _notary = Auth.recoverSigner(_view.attestationData(), _view.notarySignature().clone());
        require(_isNotary(_view.attestedDomain(), _notary), "Signer is not a notary");
    }

    function _isNotary(uint32 _origin, address _notary) internal view virtual returns (bool);
}
