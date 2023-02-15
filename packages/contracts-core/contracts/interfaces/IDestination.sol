// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice Minimal interface for Destination contract, required for sending messages.
interface IDestination {
    /**
     * @notice Emitted when an attestation is accepted by the Destination contract.
     * @param notary        Notary who signed the attestation
     * @param snapshot      Raw payload with snapshot data
     * @param signature     Agent signature for the snapshot
     */
    event SnapshotAccepted(address indexed notary, bytes snapshot, bytes signature);

    /**
     * @notice Emitted when a message is successfully executed.
     * @param origin        Domain where executed message originated
     * @param messageHash   Hash of the executed message
     */
    event Executed(uint32 indexed origin, bytes32 indexed messageHash);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               EXTERNAL                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Executes a formatted message, by passing its contents to the recipient.
     * Origin merkle root is derived using:
     * - message hash
     * - origin proof (message nonce is used as index)
     * Attestation root is derived using:
     * - derived origin merkle root
     * - origin domain
     * - attestation proof (provided attestation index is used)
     * Optimistic timeout for the message started when the derived attestation root
     * was submitted to Destination contract.
     * @dev Recipient must implement a `handle` method (refer to IMessageRecipient.sol)
     * Will revert if either of these is true:
     * - Message destination domain is not local domain.
     * - Derived attestation root is invalid
     * - Optimistic period for the message hasn't passed yet.
     * - Recipient reverted upon receiving the message.
     * @param _message              Formatted message payload
     * @param _originProof          Proof of inclusion of message hash into Origin Merkle Tree
     * @param _attestationIndex     Index of origin state in the attestation's snapshot
     * @param _attestationProof     Proof of inclusion of origin merkle root into the attestation
     */
    function execute(
        bytes memory _message,
        bytes32[32] calldata _originProof,
        uint256 _attestationIndex,
        bytes32[] calldata _attestationProof
    ) external;

    /**
     * @notice Submit an attestation signed by an active Notary.
     * Optimistic timer starts ticking for the attestation root, which effectively
     * starts the timer for all origin merkle roots used in the attestation.
     * @dev Will revert if either of these is true:
     * - Provided payload is not a formatted Attestation.
     * - Attestation signer is not an active Notary.
     * @param _payload      Raw payload with attestation data
     * @param _signature    Notary signature for the attestation
     * @return wasAccepted  Whether the attestation was accepted by the Destination contract
     */
    function submitAttestation(bytes memory _payload, bytes memory _signature)
        external
        returns (bool wasAccepted);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: draft some views for the Executor / Guards / Notaries
}
