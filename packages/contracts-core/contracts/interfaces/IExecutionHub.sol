// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MessageStatus} from "../libs/Structures.sol";

interface IExecutionHub {
    /**
     * @notice Attempts to prove inclusion of message into one of Snapshot Merkle Trees,
     * previously submitted to this contract in a form of a signed Attestation.
     * Proven message is immediately executed by passing its contents to the specified recipient.
     * @dev Will revert if any of these is true:
     *  - Message is not meant to be executed on this chain
     *  - Message was sent from this chain
     *  - Message payload is not properly formatted.
     *  - Snapshot root (reconstructed from message hash and proofs) is unknown
     *  - Snapshot root is known, but was submitted by an inactive Notary
     *  - Snapshot root is known, but optimistic period for a message hasn't passed
     *  - Provided gas limit is lower than the one requested in the message
     *  - Recipient doesn't implement a `handle` method (refer to IMessageRecipient.sol)
     *  - Recipient reverted upon receiving a message
     * Note: refer to libs/memory/State.sol for details about Origin State's sub-leafs.
     * @param msgPayload    Raw payload with a formatted message to execute
     * @param originProof   Proof of inclusion of message in the Origin Merkle Tree
     * @param snapProof     Proof of inclusion of Origin State's Left Leaf into Snapshot Merkle Tree
     * @param stateIndex    Index of Origin State in the Snapshot
     * @param gasLimit      Gas limit for message execution
     */
    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint8 stateIndex,
        uint64 gasLimit
    ) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns attestation nonce for a given snapshot root.
     * @dev Will return 0 if the root is unknown.
     */
    function getAttestationNonce(bytes32 snapRoot) external view returns (uint32 attNonce);

    /**
     * @notice Checks the validity of the unsigned message receipt.
     * @dev Will revert if any of these is true:
     *  - Receipt payload is not properly formatted.
     *  - Receipt signer is not an active Notary.
     *  - Receipt destination chain does not refer to this chain.
     * @param rcptPayload       Raw payload with Receipt data
     * @return isValid          Whether the requested receipt is valid.
     */
    function isValidReceipt(bytes memory rcptPayload) external view returns (bool isValid);

    /**
     * @notice Returns message execution status: None/Failed/Success.
     * @param messageHash       Hash of the message payload
     * @return status           Message execution status
     */
    function messageStatus(bytes32 messageHash) external view returns (MessageStatus status);

    /**
     * @notice Returns a formatted payload with the message receipt.
     * @dev Notaries could derive the tips, and the tips proof using the message payload, and submit
     * the signed receipt with the proof of tips to `Summit` in order to initiate tips distribution.
     * @param messageHash       Hash of the message payload
     * @return data             Formatted payload with the message execution receipt
     */
    function messageReceipt(bytes32 messageHash) external view returns (bytes memory data);
}
