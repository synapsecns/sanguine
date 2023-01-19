// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../libs/Attestation.sol";
import { Auth } from "../libs/Auth.sol";
import { AttestationHubEvents } from "../events/AttestationHubEvents.sol";
import { AgentRegistry } from "../system/AgentRegistry.sol";

/**
 * @notice Keeps track of the agents and verifies signed attestations.
 */
abstract contract AttestationHub is AttestationHubEvents, AgentRegistry {
    using AttestationLib for bytes;
    using AttestationLib for Attestation;
    using AttestationLib for AttestationData;

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
     * @param _attPayload   Payload with Attestation data and signature (see Attestation.sol)
     * @return TRUE if Attestation was handled correctly.
     */
    function submitAttestation(bytes memory _attPayload) external returns (bool) {
        // Check if Attestation payload is properly formatted, i.e that it
        // contains attestation data and at least one agent signature for that data
        Attestation att = _attPayload.castToAttestation();
        // Verify the attestation signature and recover an active notary address
        (address[] memory guards, address[] memory notaries) = _verifyAttestation(att);
        return _handleAttestation(guards, notaries, att, _attPayload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Child contract should implement logic for handling the Attestation.
     * @param _guards       Guard addresses (signatures&roles already verified)
     * @param _notaries     Notary addresses (signatures&roles already verified)
     * @param _att          Memory view over the Attestation for convenience
     * @param _attPayload   Payload with Attestation data and signature
     * @return TRUE if Attestation was handled correctly.
     */
    function _handleAttestation(
        address[] memory _guards,
        address[] memory _notaries,
        Attestation _att,
        bytes memory _attPayload
    ) internal virtual returns (bool);

    /**
     * @notice Checks if attestation signer is authorized.
     * @dev Guard signers need to be active globally.
     * Notary signers need to be active on destination domain.
     * @param _att      Memory view over the Attestation to check
     * @return guards   Addresses of the Guards who signed the Attestation
     * @return notaries Addresses of the Notaries who signed the Attestation
     */
    function _verifyAttestation(Attestation _att)
        internal
        view
        returns (address[] memory guards, address[] memory notaries)
    {
        AttestationData attData = _att.data();
        bytes32 digest = Auth.toEthSignedMessageHash(attData.unwrap());
        // Get amount of signatures, and initiate the returned arrays
        (uint256 guardsAmount, uint256 notariesAmount) = _att.agentsAmount();
        guards = new address[](guardsAmount);
        notaries = new address[](notariesAmount);
        // Check if all Guard signatures are valid. Guards are stored with `_domain == 0`.
        for (uint256 i = 0; i < guardsAmount; ++i) {
            guards[i] = _checkAgentAuth({
                _domain: 0,
                _digest: digest,
                _signature: _att.guardSignature(i)
            });
        }
        // Check if all Notary signatures are valid. Should be active on destination domain.
        uint32 destination = attData.destination();
        for (uint256 i = 0; i < notariesAmount; ++i) {
            notaries[i] = _checkAgentAuth({
                _domain: destination,
                _digest: digest,
                _signature: _att.notarySignature(i)
            });
        }
    }
}
