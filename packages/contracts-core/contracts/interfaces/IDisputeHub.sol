// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice Common functions for contracts relying on Agent-signed Statements
interface IDisputeHub {
    /**
     * @notice Submit an StateReport signed by a Guard, a Snapshot containing the reported State,
     * as well as Notary signature for the Snapshot.
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Notary.
     *  - State index is out of range.
     *  - Snapshot's state and reported state don't match.
     * @param _stateIndex       Index of the reported State in the Snapshot
     * @param _srPayload        Raw payload with StateReport data
     * @param _srSignature      Guard signature for the report
     * @param _snapPayload      Raw payload with Snapshot data
     * @param _snapSignature    Notary signature for the Snapshot
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReport(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes memory _snapPayload,
        bytes memory _snapSignature
    ) external returns (bool wasAccepted);

    /**
     * @notice Submit an StateReport signed by a Guard, a proof of inclusion
     * of the reported State in an Attestation, as well as Notary signature for the Attestation.
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     *  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
     *  - Snapshot Proof has length different to Attestation height.
     *  - Snapshot Proof's first element does not match the State metadata.
     *  - State index is out of range.
     * @param _stateIndex       Index of the reported State in the Snapshot
     * @param _srPayload        Raw payload with StateReport data
     * @param _srSignature      Guard signature for the report
     * @param _snapProof        Proof of inclusion of reported State's Left Leaf into Snapshot Merkle Tree
     * @param _attPayload       Raw payload with Attestation data
     * @param _attSignature     Notary signature for the Attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithProof(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes32[] memory _snapProof,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool wasAccepted);
}
