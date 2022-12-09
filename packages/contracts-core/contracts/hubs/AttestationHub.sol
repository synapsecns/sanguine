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

    // TODO: implement a way to store the submitted Attestations, so that
    // the off-chain actors don't need to rely on eth_getLogs in order to query the latest ones.

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
        (address[] memory guards, address[] memory notaries) = _verifyAttestation(attestationView);
        return _handleAttestation(guards, notaries, attestationView, _attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Child contract should implement logic for handling the Attestation.
     * @param _guards           Guard addresses (signatures&roles already verified)
     * @param _notaries         Notary addresses (signatures&roles already verified)
     * @param _attestationView  Memory view over the Attestation for convenience
     * @param _attestation      Payload with Attestation data and signature
     * @return TRUE if Attestation was handled correctly.
     */
    function _handleAttestation(
        address[] memory _guards,
        address[] memory _notaries,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal virtual returns (bool);

    /**
     * @notice Checks if attestation signer is authorized.
     * @dev Guard signers need to be active globally.
     * Notary signers need to be active on destination domain.
     * @param _attestationView  Memory view over the Attestation to check
     * @return guards   Addresses of the Guards who signed the Attestation
     * @return notaries Addresses of the Notaries who signed the Attestation
     */
    function _verifyAttestation(bytes29 _attestationView)
        internal
        view
        returns (address[] memory guards, address[] memory notaries)
    {
        // Check if Attestation payload is properly formatted, i.e that it
        // contains attestation data and at least one agent signature for that data
        require(_attestationView.isAttestation(), "Not an attestation");
        bytes32 digest = Auth.toEthSignedMessageHash(_attestationView.attestationData());
        // Get amount of signatures, and initiate the returned arrays
        (uint256 guardSigs, uint256 notarySigs) = _attestationView.agentSignatures();
        guards = new address[](guardSigs);
        notaries = new address[](notarySigs);
        // Check if all Guard signatures are valid. Guards are stored with `_domain == 0`.
        for (uint256 i = 0; i < guardSigs; ++i) {
            guards[i] = _checkAgentAuth({
                _domain: 0,
                _digest: digest,
                _signatureView: _attestationView.guardSignature(i)
            });
        }
        // Check if all Notary signatures are valid. Should be active on destination domain.
        uint32 destination = _attestationView.attestedDestination();
        for (uint256 i = 0; i < notarySigs; ++i) {
            notaries[i] = _checkAgentAuth({
                _domain: destination,
                _digest: digest,
                _signatureView: _attestationView.notarySignature(i)
            });
        }
    }
}
