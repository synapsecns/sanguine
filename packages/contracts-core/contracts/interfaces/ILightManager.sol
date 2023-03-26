// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface ILightManager {
    /**
     * @notice Registers an agent as active against the latest known Agent Merkle Root.
     * @dev Will revert if the provided proof doesn't match the latest merkle root.
     * @param _domain   Domain where the Agent is active
     * @param _agent    Agent address
     * @param _proof    Merkle proof of Active status for the agent
     * @param _index    Agent index in the merkle tree
     */
    function addAgent(
        uint32 _domain,
        address _agent,
        bytes32[] memory _proof,
        uint256 _index
    ) external;

    /**
     * @notice Updates the root of Agent Merkle Tree that the Light Manager is tracking.
     * Could be only called by a local Destination contract, which is supposed to
     * verify the attested Agent Merkle Roots.
     * @param _agentRoot    New Agent Merkle Root
     */
    function setAgentRoot(bytes32 _agentRoot) external;
}
