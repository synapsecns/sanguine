// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice Minimal interface for Origin contract, required for sending messages.
interface IOrigin {
    /**
     * @notice Emitted when a new message is dispatched.
     * @param messageHash   Hash of message; the leaf inserted to the Merkle tree for the message
     * @param nonce         Nonce of sent message (starts from 1)
     * @param destination   Destination domain
     * @param root          Merkle tree root after the new leaf was inserted
     * @param message       Raw bytes of message
     */
    event Dispatched(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes32 root,
        bytes message
    );

    /**
     * @notice Emitted when a proof of incorrect snapshot is submitted.
     * @param stateIndex    Index of incorrect state in the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param signature     Agent signature for the snapshot
     */
    event IncorrectSnapshot(uint256 stateIndex, bytes snapshot, bytes signature);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               EXTERNAL                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Dispatch the message to the destination domain & recipient
     * @dev Format the message, insert its hash into Merkle tree,
     * enqueue the new Merkle root, and emit `Dispatch` event with message information.
     * @param _destination          Domain of destination chain
     * @param _recipient            Address of recipient on destination chain as bytes32
     * @param _optimisticSeconds    Optimistic period for message execution on destination chain
     * @param _tips                 Payload with information about paid tips
     * @param _messageBody          Raw bytes content of message
     * @return messageNonce         Nonce of the dispatched message
     * @return messageHash          Hash of the formatted dispatched message
     */
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable returns (uint32 messageNonce, bytes32 messageHash);

    /**
     * @notice Verifies a state from the snapshot (a list of states) signed by a Guard or a Notary.
     * Does nothing, if the state is valid (matches the historical state of this contract).
     * Slashes the snapshot signer, if the state is invalid.
     * @dev Will revert if either of these is true:
     * - Provided payload is not a formatted Snapshot.
     * - Snapshot signer is not an active Guard or Notary.
     * - State index is out of range.
     * - Snapshot state does not refer to this chain.
     * @param _payload      Raw payload with snapshot data
     * @param _signature    Agent signature for the snapshot
     * @param _stateIndex   State index to check
     * @return isValid      Whether the requested state is valid.
     *                      Agent is slashed, if return value is FALSE.
     */
    function verifySnapshot(
        bytes memory _payload,
        bytes memory _signature,
        uint256 _stateIndex
    ) external returns (bool isValid);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Checks if the snapshot's states match the historical state of Origin contract.
     * Only the states referring to this chain are checked for validity.
     * @dev Will revert if provided payload is not a formatted Snapshot.
     * Will return ZERO, if no states from the snapshot refer to this chain.
     * Will return a non-zero value, if at least one snapshots's states is invalid.
     * @param _payload          Raw payload with snapshot data
     * @return invalidState     ZERO, if all states are valid, as far as this chain is concerned.
     *                          Index of the first invalid state PLUS 1 otherwise.
     */
    function isValidSnapshot(bytes memory _payload) external view returns (uint256 invalidState);

    /**
     * @notice Checks if a state matches the historical state of Origin contract.
     * @dev Will revert if provided payload is not a formatted State.
     * Will return TRUE, if state does not refer to this chain.
     * @param _payload  Raw payload with state data
     * @return isValid  Whether the provided state is valid, as far as this chain is concerned
     */
    function isValidState(bytes memory _payload) external view returns (bool isValid);

    /**
     * @notice Suggest a historical state data for an off-chain agent to sign.
     * Note: signing the suggested state data will will never lead to slashing of the actor,
     * assuming they have confirmed that the block, which number is included in the data,
     * is not subject to reorganization (which is different for every observed chain).
     * @dev Will revert if message with given nonce hasn't been dispatched yet
     * @param _nonce        Nonce of a historical dispatched message
     * @return stateData    Formatted State payload containing data about state of the Origin
     *                      contract after message with provided nonce was dispatched.
     */
    function suggestState(uint32 _nonce) external view returns (bytes memory stateData);
}
