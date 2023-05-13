// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus} from "../libs/Structures.sol";

interface InterfaceSummit {
    // ══════════════════════════════════════════ ACCEPT AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Accepts a receipt, which local `AgentManager` verified to have been signed by an active Notary.
     * > Receipt is a statement about message execution status on the remote chain.
     * - This will distribute the message tips across the off-chain actors once the receipt optimistic period is over.
     * - Notary who signed the receipt is referenced as the "Receipt Notary".
     * - Notary who signed the attestation on destination chain is referenced as the "Attestation Notary".
     * > Will revert if any of these is true:
     * > - Called by anyone other than local `AgentManager`.
     * > - Receipt body payload is not properly formatted.
     * > - Receipt signer is in Dispute.
     * > - Receipt's snapshot root is unknown.
     * @param rcptNotaryIndex   Index of Receipt Notary in Agent Merkle Tree
     * @param attNotaryIndex    Index of Attestation Notary in Agent Merkle Tree
     * @param sigIndex          Index of stored Notary signature
     * @param attNonce          Nonce of the attestation used for proving the executed message
     * @param paddedTips        Padded encoded paid tips information
     * @param rcptPayload       Raw payload with message execution receipt
     * @return wasAccepted      Whether the receipt was accepted
     */
    function acceptReceipt(
        uint32 rcptNotaryIndex,
        uint32 attNotaryIndex,
        uint256 sigIndex,
        uint32 attNonce,
        uint256 paddedTips,
        bytes memory rcptPayload
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a snapshot, which local `AgentManager` verified to have been signed by an active Guard.
     * > Snapshot is a list of states for a set of Origin contracts residing on any of the chains.
     * All the states in the Guard-signed snapshot become available for Notary signing.
     * > Will revert if any of these is true:
     * > - Called by anyone other than local `AgentManager`.
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot contains a state older then the Guard has previously submitted.
     * @param guardIndex        Index of Guard in Agent Merkle Tree
     * @param sigIndex          Index of stored Agent signature
     * @param snapPayload       Raw payload with snapshot data
     */
    function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes memory snapPayload) external;

    /**
     * @notice Accepts a snapshot, which local `AgentManager` verified to have been signed by an active Notary.
     * > Snapshot is a list of states for a set of Origin contracts residing on any of the chains.
     * Snapshot Merkle Root is calculated and saved for valid snapshots, i.e.
     * snapshots which are only using states previously submitted by any of the Guards.
     * - Notary could use states singed by the same of different Guards in their snapshot.
     * - Notary could then proceed to sign the attestation for their submitted snapshot.
     * > Will revert if any of these is true:
     * > - Called by anyone other than local `AgentManager`.
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot contains a state older then the Notary has previously submitted.
     * > - Snapshot contains a state that no Guard has previously submitted.
     * @param notaryIndex       Index of Notary in Agent Merkle Tree
     * @param sigIndex          Index of stored Agent signature
     * @param agentRoot         Current root of the Agent Merkle Tree
     * @param snapPayload       Raw payload with snapshot data
     * @return attPayload       Raw payload with data for attestation derived from Notary snapshot.
     */
    function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes memory snapPayload)
        external
        returns (bytes memory attPayload);

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /**
     * @notice Distributes tips using the first Receipt from the "receipt quarantine queue".
     * Possible scenarios:
     *  - Receipt queue is empty => does nothing
     *  - Receipt optimistic period is not over => does nothing
     *  - Either of Notaries present in Receipt was slashed => receipt is deleted from the queue
     *  - Either of Notaries present in Receipt in Dispute => receipt is moved to the end of queue
     *  - None of the above => receipt tips are distributed
     * @dev Returned value makes it possible to do the following: `while (distributeTips()) {}`
     * @return queuePopped      Whether the first element was popped from the queue
     */
    function distributeTips() external returns (bool queuePopped);

    /**
     * @notice Withdraws locked base message tips from requested domain Origin to the recipient.
     * This is done by a call to a local Origin contract, or by a manager message to the remote chain.
     * @dev This will revert, if the pending balance of origin tips (earned-claimed) is lower than requested.
     * @param origin    Domain of chain to withdraw tips on
     * @param amount    Amount of tips to withdraw
     */
    function withdrawTips(uint32 origin, uint256 amount) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns earned and claimed tips for the actor.
     * Note: Tips for address(0) belong to the Treasury.
     * @param actor     Address of the actor
     * @param origin    Domain where the tips were initially paid
     * @return earned   Total amount of origin tips the actor has earned so far
     * @return claimed  Total amount of origin tips the actor has claimed so far
     */
    function actorTips(address actor, uint32 origin) external view returns (uint128 earned, uint128 claimed);

    /**
     * @notice Returns the amount of receipts in the "Receipt Quarantine Queue".
     */
    function receiptQueueLength() external view returns (uint256);

    /**
     * @notice Returns the state with the highest known nonce
     * submitted by any of the currently active Guards.
     * @param origin        Domain of origin chain
     * @return statePayload Raw payload with latest active Guard state for origin
     */
    function getLatestState(uint32 origin) external view returns (bytes memory statePayload);
}
