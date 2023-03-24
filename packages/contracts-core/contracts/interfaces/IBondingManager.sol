// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

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
}
