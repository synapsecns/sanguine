// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the Destination contract
abstract contract DestinationEvents {
    /**
     * @notice Emitted when a snapshot is accepted by the Destination contract.
     * @param domain        Domain where the signed Notary is active
     * @param notary        Notary who signed the attestation
     * @param attestation   Raw payload with attestation data
     * @param attSignature  Notary signature for the attestation
     */
    event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature);

    /**
     * @notice Emitted when message is executed.
     * @param remoteDomain  Remote domain where message originated
     * @param messageHash   The keccak256 hash of the message that was executed
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash);

    /**
     * @notice Emitted when tips are stored.
     * @param notary        Notary who signed the Snapshot Root used for proving the message
     * @param tips          Raw payload with tips paid for the off-chain agents
     */
    event TipsStored(address indexed notary, bytes tips);
}
