// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceBondingManager {
    // ═══════════════════════════════════════════════ AGENTS LOGIC ════════════════════════════════════════════════════

    /**
     * @notice Adds a new agent for the domain. This is either a fresh address (Inactive),
     * or an agent who used to be active on the same domain before (Resting).
     * @dev Inactive: `proof` should be the proof of inclusion of an empty leaf
     * having index following the last added agent in the tree.
     * @dev Resting: `proof` should be the proof of inclusion of the agent leaf
     * with Resting flag having index previously assigned to the agent.
     * @param domain    Domain where the Agent will be active
     * @param agent     Address of the Agent
     * @param proof     Merkle proof of the Inactive/Resting status for the agent
     */
    function addAgent(uint32 domain, address agent, bytes32[] memory proof) external;

    /**
     * @notice Initiates the unstaking of the agent bond. Agent signature is immediately no longer
     * considered valid on Synapse Chain, and will be invalid on other chains once the Light Manager
     * updates their agent merkle root on these chains.
     * @dev `proof` should be the proof of inclusion of the agent leaf
     * with Active flag having index previously assigned to the agent.
     * @param domain    Domain where the Agent is active
     * @param agent     Address of the Agent
     * @param proof     Merkle proof of the Active status for the agent
     */
    function initiateUnstaking(uint32 domain, address agent, bytes32[] memory proof) external;

    /**
     * @notice Completes the unstaking of the agent bond. Agent signature is no longer considered
     * valid on any of the chains.
     * @dev `proof` should be the proof of inclusion of the agent leaf
     * with Unstaking flag having index previously assigned to the agent.
     * @param domain    Domain where the Agent was active
     * @param agent     Address of the Agent
     * @param proof     Merkle proof of the unstaking status for the agent
     */
    function completeUnstaking(uint32 domain, address agent, bytes32[] memory proof) external;

    /**
     * @notice Completes the slashing of the agent bond. Agent signature is no longer considered
     * valid under the updated Agent Merkle Root.
     * @dev `proof` should be the proof of inclusion of the agent leaf
     * with Active/Unstaking flag having index previously assigned to the agent.
     * @param domain    Domain where the Agent was active
     * @param agent     Address of the Agent
     * @param proof     Merkle proof of the active/unstaking status for the agent
     */
    function completeSlashing(uint32 domain, address agent, bytes32[] memory proof) external;

    /**
     * @notice Remote AgentManager should call this function to indicate that the agent
     * has been proven to commit fraud on the origin chain.
     * @dev This initiates the process of agent slashing. It could be immediately
     * completed by anyone calling completeSlashing() providing a correct merkle proof
     * for the OLD agent status.
     * Note: as an extra security check this function returns its own selector, so that
     * Destination could verify that a "remote" function was called when executing a manager message.
     * Will revert if `msgOrigin` is equal to contract's local domain.
     * @param domain        Domain where the slashed agent was active
     * @param agent         Address of the slashed Agent
     * @param prover        Address that initially provided fraud proof to remote AgentManager
     * @return magicValue   Selector of this function
     */
    function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover)
        external
        returns (bytes4 magicValue);

    /**
     * @notice Withdraws locked base message tips from requested domain Origin to the recipient.
     * Issues a call to a local Origin contract, or sends a manager message to the remote chain.
     * @dev Could only be called by the Summit contract.
     * @param recipient     Address to withdraw tips to
     * @param origin        Domain where tips need to be withdrawn
     * @param amount        Tips value to withdraw
     */
    function withdrawTips(address recipient, uint32 origin, uint256 amount) external;

    /**
     * @notice Allows contract owner to resolve a stuck Dispute.
     * This could only be called if no fresh data has been submitted by the Notaries to the Inbox,
     * which is required for the Dispute to be resolved naturally.
     * > Will revert if any of these is true:
     * > - Caller is not contract owner.
     * > - Domain doesn't match the saved agent domain.
     * > - `slashedAgent` is not in Dispute.
     * > - Less than `FRESH_DATA_TIMEOUT` has passed since the last Notary submission to the Inbox.
     * @param slashedAgent  Agent that is being slashed
     */
    function resolveDisputeWhenStuck(uint32 domain, address slashedAgent) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns all active agents for a given domain.
     * @param domain    Domain to get agents from (ZERO for Guards)
     * @param agents    List of active agents for the domain
     */
    function getActiveAgents(uint32 domain) external view returns (address[] memory agents);

    /**
     * @notice Returns a leaf representing the current status of agent in the Agent Merkle Tree.
     * @dev Will return an empty leaf, if agent is not added to the tree yet.
     * @param agent     Agent address
     * @return leaf     Agent leaf in the Agent Merkle Tree
     */
    function agentLeaf(address agent) external view returns (bytes32 leaf);

    /**
     * @notice Returns a total amount of leafs representing known agents.
     * @dev This includes active, unstaking, resting and slashed agents.
     * This also includes an empty leaf as the very first entry.
     */
    function leafsAmount() external view returns (uint256 amount);

    /**
     * @notice Returns a full list of leafs from the Agent Merkle Tree.
     * @dev This might consume a lot of gas, do not use this on-chain.
     */
    function allLeafs() external view returns (bytes32[] memory leafs);

    /**
     * @notice Returns a list of leafs from the Agent Merkle Tree
     * with indexes [indexFrom .. indexFrom + amount).
     * @dev This might consume a lot of gas, do not use this on-chain.
     * @dev Will return less than `amount` entries, if indexFrom + amount > leafsAmount
     */
    function getLeafs(uint256 indexFrom, uint256 amount) external view returns (bytes32[] memory leafs);

    /**
     * @notice Returns a proof of inclusion of the agent in the Agent Merkle Tree.
     * @dev Will return a proof for an empty leaf, if agent is not added to the tree yet.
     * This proof could be used by ANY next new agent that calls {addAgent}.
     * @dev This WILL consume a lot of gas, do not use this on-chain.
     * @dev The alternative way to create a proof is to fetch the full list of leafs using
     * either {allLeafs} or {getLeafs}, and create a merkle proof from that.
     * @param agent     Agent address
     * @return proof    Merkle proof for the agent
     */
    function getProof(address agent) external view returns (bytes32[] memory proof);
}
