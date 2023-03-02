// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAttestationHub } from "./IAttestationHub.sol";
import { ORIGIN_TREE_DEPTH } from "../libs/Structures.sol";

interface InterfaceDestination is IAttestationHub {
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
     * @param _snapIndex    Index of Origin State's Left Leaf in the Snapshot Merkle Tree
     */
    function execute(
        bytes memory _message,
        bytes32[ORIGIN_TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _snapIndex
    ) external;

    /**
     * @notice Submit an Attestation signed by a Notary.
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary for local domain.
     *  - Attestation's snapshot root has been previously submitted.
     */
    function submitAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool wasAccepted);
}
