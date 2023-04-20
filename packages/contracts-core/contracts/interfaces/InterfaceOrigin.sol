// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceOrigin {
    // ═══════════════════════════════════════════════ SEND MESSAGES ═══════════════════════════════════════════════════

    /**
     * @notice Send a message to the recipient located on destination domain.
     * @dev Recipient has to conform to IMessageRecipient interface, otherwise message won't be delivered.
     * @param destination           Domain of destination chain
     * @param recipient             Address of recipient on destination chain as bytes32
     * @param optimisticPeriod      Optimistic period for message execution on destination chain
     * @param paddedTips            Padded encoded paid tips information
     * @param paddedRequest         Padded encoded message execution request on destination chain
     * @param content               Raw bytes content of message
     * @return messageNonce         Nonce of the sent message
     * @return messageHash          Hash of the sent message
     */
    function sendBaseMessage(
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticPeriod,
        uint256 paddedTips,
        uint256 paddedRequest,
        bytes memory content
    ) external payable returns (uint32 messageNonce, bytes32 messageHash);

    /**
     * @notice Send a system message to the destination domain.
     * @dev This could only be called by SystemRouter, which takes care of encoding/decoding the message body.
     * The message body includes the sender and the recipient of the system message.
     * Note: function is not payable, as no tips are required for sending a system message.
     * @param destination           Domain of destination chain
     * @param optimisticPeriod      Optimistic period for message execution on destination chain
     * @param body                  Body of the system message
     */
    function sendSystemMessage(uint32 destination, uint32 optimisticPeriod, bytes memory body)
        external
        returns (uint32 messageNonce, bytes32 messageHash);

    /**
     * @notice Withdraws locked base message tips to the recipient.
     * @dev Could only be called by a local AgentManager.
     * @param recipient     Address to withdraw tips to
     * @param amount        Tips value to withdraw
     */
    function withdrawTips(address recipient, uint256 amount) external;

    // ═════════════════════════════════════════════ VERIFY STATEMENTS ═════════════════════════════════════════════════

    /**
     * @notice Verifies a state from the snapshot, that was used for the Notary-signed attestation.
     * Does nothing, if the state is valid (matches the historical state of this contract).
     * Slashes the attestation signer, if the state is invalid.
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     *  - Attestation root is not equal to Snapshot Merkle Root.
     *  - Snapshot payload is not properly formatted.
     *  - State index is out of range.
     *  - State does not refer to this chain.
     * @param stateIndex        State index to check
     * @param snapPayload       Raw payload with snapshot data
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValid          Whether the requested state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(
        uint256 stateIndex,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValid);

    /**
     * @notice Verifies a state from the snapshot, that was used for the Notary-signed attestation.
     * Does nothing, if the state is valid (matches the historical state of this contract).
     * Slashes the attestation signer, if the state is invalid.
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     *  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
     *  - Snapshot Proof's first element does not match the State metadata.
     *  - Snapshot Proof length exceeds Snapshot tree Height.
     *  - State payload is not properly formatted.
     *  - State index is out of range.
     *  - State does not refer to this chain.
     * @param stateIndex        Index of state in the snapshot
     * @param statePayload      Raw payload with State data to check
     * @param snapProof         Proof of inclusion of provided State's Left Leaf into Snapshot Merkle Tree
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValid          Whether the requested state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyAttestationWithProof(
        uint256 stateIndex,
        bytes memory statePayload,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValid);

    /**
     * @notice Verifies a state from the snapshot (a list of states) signed by a Guard or a Notary.
     * Does nothing, if the state is valid (matches the historical state of this contract).
     * Slashes the snapshot signer, if the state is invalid.
     * @dev Will revert if any of these is true:
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Agent.
     *  - State index is out of range.
     *  - Snapshot state does not refer to this chain.
     * @param stateIndex        State index to check
     * @param snapPayload       Raw payload with snapshot data
     * @param snapSignature     Agent signature for the snapshot
     * @return isValid          Whether the requested state is valid.
     *                          Agent is slashed, if return value is FALSE.
     */
    function verifySnapshot(uint256 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValid);

    /**
     * @notice Verifies a snapshot report signed by a Guard.
     *  - Does nothing, if the report is valid (if the reported snapshot is invalid).
     *  - Slashes the Guard otherwise (meaning the reported snapshot is valid, making the report invalid).
     * @dev Will revert if any of these is true:
     *  - Report payload is not properly formatted.
     *  - Report signer is not an active Guard.
     * @param srPayload         Raw payload with StateReport data
     * @param srSignature       Guard signature for the report
     * @return isValid          Whether the provided report is valid.
     *                          Guard is slashed, if return value is FALSE.
     */
    function verifyStateReport(bytes memory srPayload, bytes memory srSignature) external returns (bool isValid);
}
