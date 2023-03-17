// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TREE_DEPTH } from "../libs/Constants.sol";

interface InterfaceDestination {
    /**
     * @notice Attempts to prove inclusion of message into one of Snapshot Merkle Trees,
     * previously submitted to this contract in a form of a signed Attestation.
     * Proven message is immediately executed by passing its contents to the specified recipient.
     * @dev Will revert if any of these is true:
     *  - Message payload is not properly formatted.
     *  - Snapshot root (reconstructed from message hash and proofs) is unknown
     *  - Snapshot root is known, but was submitted by an inactive Notary
     *  - Snapshot root is known, but optimistic period for a message hasn't passed
     *  - Recipient doesn't implement a `handle` method (refer to IMessageRecipient.sol)
     *  - Recipient reverted upon receiving a message
     * Note: refer to libs/State.sol for details about Origin State's sub-leafs.
     * @param _message      Raw payload with a formatted message to execute
     * @param _originProof  Proof of inclusion of message in the Origin Merkle Tree
     * @param _snapProof    Proof of inclusion of Origin State's Left Leaf into Snapshot Merkle Tree
     * @param _stateIndex   Index of Origin State in the Snapshot
     */
    function execute(
        bytes memory _message,
        bytes32[TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) external;

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
}
