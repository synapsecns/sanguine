// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceOrigin {
    /**
     * @notice Dispatch the message to the recipient located on destination domain.
     * @param _destination          Domain of destination chain
     * @param _recipient            Address of recipient on destination chain as bytes32
     * @param _optimisticSeconds    Optimistic period for message execution on destination chain
     * @param _tips                 Payload with information about paid tips
     * @param _messageBody          Raw bytes content of message
     * @return messageNonce         Nonce of the dispatched message
     * @return messageHash          Hash of the dispatched message
     */
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable returns (uint32 messageNonce, bytes32 messageHash);

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
     * @param _stateIndex       State index to check
     * @param _snapPayload      Raw payload with snapshot data
     * @param _attPayload       Raw payload with Attestation data
     * @param _attSignature     Notary signature for the attestation
     * @return isValid          Whether the requested state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(
        uint256 _stateIndex,
        bytes memory _snapPayload,
        bytes memory _attPayload,
        bytes memory _attSignature
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
     * @param _stateIndex       State index to check
     * @param _snapPayload      Raw payload with snapshot data
     * @param _snapSignature    Agent signature for the snapshot
     * @return isValid          Whether the requested state is valid.
     *                          Agent is slashed, if return value is FALSE.
     */
    function verifySnapshot(
        uint256 _stateIndex,
        bytes memory _snapPayload,
        bytes memory _snapSignature
    ) external returns (bool isValid);
}
