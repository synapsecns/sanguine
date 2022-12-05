// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Auth } from "../libs/Auth.sol";
import { Attestation } from "../libs/Attestation.sol";
import { AttestationHubEvents } from "../events/AttestationHubEvents.sol";
import { AgentRegistry } from "../system/AgentRegistry.sol";

/**
 * @notice Keeps track of the agents and verifies signed attestations.
 */
abstract contract AttestationHub is AttestationHubEvents, AgentRegistry {
    using Attestation for bytes;
    using Attestation for bytes29;

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
        bytes29 attestationView = _attestation.castToAttestation();
        // Verify the attestation signature and recover an active notary address
        address notary = _verifyAttestation(attestationView);
        return _handleAttestation(notary, attestationView, _attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Child contract should implement logic for handling the Attestation.
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over the Attestation for convenience
     * @param _attestation      Payload with Attestation data and signature
     * @return TRUE if Attestation was handled correctly.
     */
    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal virtual returns (bool);

    /**
     * @notice Checks if attestation signer is authorized.
     * @dev Signer needs to be an active Notary on destination domain.
     * @param _attestationView  Memory view over the Attestation to check
     * @return notary Address of the attestation signer
     */
    function _verifyAttestation(bytes29 _attestationView) internal view returns (address notary) {
        // Check if Attestation payload is properly formatted.
        require(_attestationView.isAttestation(), "Not an attestation");
        bytes32 digest = Auth.toEthSignedMessageHash(_attestationView.attestationData());
        // Check if Notary signature is valid. Should be active on destination domain
        notary = _checkAgentAuth({
            _domain: _attestationView.attestedDestination(),
            _digest: digest,
            _signatureView: _attestationView.notarySignature()
        });
    }
}
