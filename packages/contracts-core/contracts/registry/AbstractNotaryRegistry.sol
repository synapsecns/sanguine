// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { Attestation } from "../libs/Attestation.sol";
import { Auth } from "../libs/Auth.sol";
import { NotaryRegistryEvents } from "../events/NotaryRegistryEvents.sol";

/**
 * @notice Registry used for verifying Attestations signed by Notaries.
 * This is done agnostic of how the Notaries are actually stored.
 * The child contract is responsible for implementing the Notaries storage.
 * @dev It is assumed that the Notary signature is only valid for a subset of origins.
 */
abstract contract AbstractNotaryRegistry is NotaryRegistryEvents {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier haveActiveNotary(uint32 _domain) virtual;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Adds a new Notary to Registry.
     * @dev Child contracts should implement this depending on how Notaries are stored.
     * @param _origin   Origin domain where Notary is added
     * @param _notary   New Notary to add
     * @return TRUE if a notary was added
     */
    function _addNotary(uint32 _origin, address _notary) internal virtual returns (bool);

    /**
     * @notice Removes a Notary from Registry.
     * @dev Child contracts should implement this depending on how Notaries are stored.
     * @param _origin   Origin domain where Notary is removed
     * @param _notary   Notary to remove
     * @return TRUE if a notary was removed
     */
    function _removeNotary(uint32 _origin, address _notary) internal virtual returns (bool);

    // solhint-disable no-empty-blocks
    /**
     * @notice Hook that is called just before a Notary is added for specified domain.
     */
    function _beforeNotaryAdded(uint32 _domain, address _notary) internal virtual {}

    /**
     * @notice Hook that is called right after a Notary is added for specified domain.
     */
    function _afterNotaryAdded(uint32 _domain, address _notary) internal virtual {}

    /**
     * @notice Hook that is called just before a Notary is removed from specified domain.
     */
    function _beforeNotaryRemoved(uint32 _domain, address _notary) internal virtual {}

    /**
     * @notice Hook that is called right after a Notary is removed from specified domain.
     */
    function _afterNotaryRemoved(uint32 _domain, address _notary) internal virtual {}

    // solhint-enable no-empty-blocks
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
        require(_isNotary(_view.attestedDestination(), _notary), "Signer is not a notary");
    }

    /**
     * @notice Checks whether a given account in an authorized Notary.
     * @dev Child contracts should implement this depending on how Notaries are stored.
     * @param _origin   Origin domain to check
     * @param _account  Address to check for being a Notary
     * @return TRUE if the account is an authorized Notary.
     */
    function _isNotary(uint32 _origin, address _account) internal view virtual returns (bool);
}
