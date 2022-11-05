// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract OriginEvents {
    /**
     * @notice Emitted when a new message is dispatched
     * @param messageHash Hash of message; the leaf inserted to the Merkle tree
     *        for the message
     * @param nonce Nonce of sent message (starts from 1)
     * @param destination Destination domain
     * @param tips Tips paid for the remote off-chain agents
     * @param message Raw bytes of message
     */
    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes tips,
        bytes message
    );

    /**
     * @notice Emitted when the Guard is slashed
     * (should be paired with IncorrectReport event)
     * @param guard     The address of the guard that signed the incorrect report
     * @param reporter  The address of the entity that reported the guard misbehavior
     */
    event GuardSlashed(address indexed guard, address indexed reporter);

    /**
     * @notice Emitted when the Notary is slashed
     * (should be paired with FraudAttestation event)
     * @param notary    The address of the notary
     * @param guard     The address of the guard that signed the fraud report
     * @param reporter  The address of the entity that reported the notary misbehavior
     */
    event NotarySlashed(address indexed notary, address indexed guard, address indexed reporter);

    /**
     * @notice Emitted when the NotaryManager contract is changed
     * @param notaryManager The address of the new notaryManager
     */
    event NewNotaryManager(address notaryManager);
}
