// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentFlag, SystemEntity } from "../libs/Structures.sol";

interface IBondingManager {
    /**
     * @notice Adds a new agent for the domain. This is either a fresh address (Inactive),
     * or an agent who used to be active on the same domain before (Resting).
     * @dev Inactive: `_proof` should be the proof of inclusion of an empty leaf
     * having index following the last added agent in the tree.
     * @dev Resting: `_proof` should be the proof of inclusion of the agent leaf
     * with Resting flag having index previously assigned to the agent.
     * @param _domain   Domain where the Agent will be active
     * @param _agent    Address of the Agent
     * @param _proof    Merkle proof of the Inactive/Resting status for the agent
     */
    function addAgent(
        uint32 _domain,
        address _agent,
        bytes32[] memory _proof
    ) external;

    /**
     * @notice Initiates the unstaking of the agent bond. Agent signature is immediately no longer
     * considered valid on Synapse Chain, and will be invalid on other chains once the Light Manager
     * updates their agent merkle root on these chains.
     * @dev `_proof` should be the proof of inclusion of the agent leaf
     * with Active flag having index previously assigned to the agent.
     * @param _domain   Domain where the Agent is active
     * @param _agent    Address of the Agent
     * @param _proof    Merkle proof of the Active status for the agent
     */
    function initiateUnstaking(
        uint32 _domain,
        address _agent,
        bytes32[] memory _proof
    ) external;

    /**
     * @notice Completes the unstaking of the agent bond. Agent signature is no longer considered
     * valid on any of the chains.
     * @dev `_proof` should be the proof of inclusion of the agent leaf
     * with Unstaking flag having index previously assigned to the agent.
     * @param _domain   Domain where the Agent was active
     * @param _agent    Address of the Agent
     * @param _proof    Merkle proof of the unstaking status for the agent
     */
    function completeUnstaking(
        uint32 _domain,
        address _agent,
        bytes32[] memory _proof
    ) external;

    /**
     * @notice Completes the slashing of the agent bond. Agent signature is no longer considered
     * valid under the updated Agent Merkle Root.
     * @dev `_proof` should be the proof of inclusion of the agent leaf
     * with Active/Unstaking flag having index previously assigned to the agent.
     * @param _domain   Domain where the Agent was active
     * @param _agent    Address of the Agent
     * @param _proof    Merkle proof of the active/unstaking status for the agent
     */
    function completeSlashing(
        uint32 _domain,
        address _agent,
        bytes32[] memory _proof
    ) external;

    /**
     * @notice Remote AgentManager should call this function to indicate that the agent
     * has been proven to commit fraud in the SystemRegistry on the origin chain.
     * @dev This initiates the process of agent slashing. It could be immediately
     * completed by anyone calling completeSlashing() providing a correct merkle proof
     * for the OLD agent status.
     * @param _domain   Domain where the slashed agent was active
     * @param _agent    Address of the slashed Agent
     * @param _reporter Address that initially provided fraud proof in SystemRegistry
     */
    function remoteRegistrySlash(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _systemCaller,
        uint32 _domain,
        address _agent,
        address _reporter
    ) external;

    // ═════════════════════════════════ VIEWS ═════════════════════════════════

    /**
     * @notice Returns current status for a given agent.
     * @param _agent    Agent address
     * @return flag     Flag signalling the agent status (see Structures.sol)
     * @return domain   Domain where the Agent is active (0 for Guards)
     * @return index    Index of agent in the Agent Merkle Tree
     */
    function agentStatus(address _agent)
        external
        view
        returns (
            AgentFlag flag,
            uint32 domain,
            uint32 index
        );

    /**
     * @notice Returns a leaf representing the current status of agent in the Agent Merkle Tree.
     * @dev Will return an empty leaf, if agent is not added to the tree yet.
     * @param _agent    Agent address
     * @return leaf     Agent leaf in the Agent Merkle Tree
     */
    function agentLeaf(address _agent) external view returns (bytes32 leaf);

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
    function getLeafs(uint256 _indexFrom, uint256 _amount)
        external
        view
        returns (bytes32[] memory leafs);

    /**
     * @notice Returns a proof of inclusion of the agent in the Agent Merkle Tree.
     * @dev Will return a proof for an empty leaf, if agent is not added to the tree yet.
     * This proof could be used by ANY next new agent that calls {addAgent}.
     * @dev This WILL consume a lot of gas, do not use this on-chain.
     * @dev The alternative way to create a proof is to fetch the full list of leafs using
     * either {allLeafs} or {getLeafs}, and create a merkle proof from that.
     * @param _agent    Agent address
     * @return proof    Merkle proof for the agent
     */
    function getProof(address _agent) external view returns (bytes32[] memory proof);
}
