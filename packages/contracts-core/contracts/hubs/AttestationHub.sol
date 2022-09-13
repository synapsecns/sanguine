// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { AbstractNotaryRegistry } from "../registry/AbstractNotaryRegistry.sol";

abstract contract AttestationHub is AbstractNotaryRegistry {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Called by the external agent. Submits the signed attestation for handling.
     * @dev Reverts if either of this is true:
     *      - Attestation payload is not properly formatted.
     *      - Attestation signer is not a Notary.
     * @param _attestation  Payload with Attestation data and signature (see Attestation.sol)
     * @return TRUE if Attestation was handled correctly.
     */
    function submitAttestation(bytes memory _attestation) external returns (bool) {
        // Check if real Notary & signature.
        // This also checks if Attestation payload is properly formatted.
        (address _notary, bytes29 _attestationView) = _checkNotaryAuth(_attestation);
        // Pass _attestationView as the existing bytes29 pointer to attestation payload
        // Pass _attestation to avoid extra memory copy when emitting attestation payload
        return _handleAttestation(_notary, _attestationView, _attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Child contract should implement logic for handling the Attestation.
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _attestation      Payload with Attestation data and signature
     * @return TRUE if Attestation was handled correctly.
     */
    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal virtual returns (bool);
}
