// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ExecutionAttestation } from "../libs/Attestation.sol";

interface InterfaceDestination {
    /**
     * @notice Attempts to pass a quarantined Agent Merkle Root to a local Light Manager.
     * @dev Will do nothing, if root optimistic period is not over.
     * Note: both returned values can not be true.
     * @return rootPassed   Whether the agent merkle root was passed to LightManager
     * @return rootPending  Whether there is a pending agent merkle root left
     */
    function passAgentRoot() external returns (bool rootPassed, bool rootPending);

    /**
     * @notice Submit an Attestation signed by a Notary.
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary for local domain.
     *  - Attestation's snapshot root has been previously submitted.
     * @param _attPayload       Raw payload with Attestation data
     * @param _attSignature     Notary signature for the reported attestation
     * @return wasAccepted      Whether the Attestation was accepted (resulting in Dispute between the agents)
     */
    function submitAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool wasAccepted);

    /**
     * @notice Submit an AttestationReport signed by a Guard, as well as Notary signature
     * for the reported Attestation.
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     *  - Attestation signer is not an active Notary for local domain.
     * @param _arPayload        Raw payload with AttestationReport data
     * @param _arSignature      Guard signature for the report
     * @param _attSignature     Notary signature for the reported attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitAttestationReport(
        bytes memory _arPayload,
        bytes memory _arSignature,
        bytes memory _attSignature
    ) external returns (bool wasAccepted);

    // ═════════════════════════════════ VIEWS ═════════════════════════════════

    /**
     * @notice Returns the total amount of Notaries attestations that have been accepted.
     */
    function attestationsAmount() external view returns (uint256);

    /**
     * @notice Returns an attestation from the list of all accepted Notary attestations.
     * @dev Index refers to attestation's snapshot root position in `roots` array.
     * @param _index   Attestation index
     * @return root    Snapshot root for the attestation
     * @return execAtt Rest of attestation data that Destination keeps track of
     */
    function getAttestation(uint256 _index)
        external
        view
        returns (bytes32 root, ExecutionAttestation memory execAtt);

    /**
     * Returns status of Destination contract as far as snapshot/agent roots are concerned
     * @return snapRootTime     Timestamp when latest snapshot root was accepted
     * @return agentRootTime    Timestamp when latest agent root was accepted
     * @return notary           Notary who signed the latest agent root
     */
    function destStatus()
        external
        view
        returns (
            uint48 snapRootTime,
            uint48 agentRootTime,
            address notary
        );

    /**
     * Returns Agent Merkle Root to be passed to LightManager once its optimistic period is over.
     */
    function nextAgentRoot() external view returns (bytes32);
}
