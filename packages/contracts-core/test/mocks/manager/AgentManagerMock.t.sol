// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, AgentStatus, IAgentManager, SlashStatus} from "../../../contracts/manager/AgentManager.sol";

contract AgentManagerMock is IAgentManager {
    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // (agent => their status)
    mapping(address => AgentStatus) private _agentMap;

    /// @inheritdoc IAgentManager
    mapping(address => SlashStatus) public slashStatus;

    /// @inheritdoc IAgentManager
    bytes32 public agentRoot;

    /// @notice Total amount of added agents
    uint256 public totalAgents;

    // ════════════════════════════════════════════ EXTERNAL FUNCTIONS ═════════════════════════════════════════════════

    /// @notice Adds agent for testing.
    /// @dev Reverts if agent is already active.
    function addAgent(uint32 domain, address agent) external {
        AgentStatus memory status = _agentMap[agent];
        require(status.flag == AgentFlag.Unknown, "Agent already active");
        status.flag = AgentFlag.Active;
        status.domain = domain;
        // index starts from 1
        status.index = uint32(++totalAgents);
        _agentMap[agent] = status;
    }

    function removeAgent(uint32 domain, address agent) external {
        AgentStatus memory status = _agentMap[agent];
        require(status.flag == AgentFlag.Active, "Agent not active");
        require(status.domain == domain, "Incorrect domain");
        delete _agentMap[agent];
    }

    /// @notice Sets agent root.
    function setAgentRoot(bytes32 agentRoot_) external {
        agentRoot = agentRoot_;
    }

    /// @inheritdoc IAgentManager
    function registrySlash(uint32 domain, address agent, address prover) external {
        AgentStatus memory status = _agentMap[agent];
        require(status.flag == AgentFlag.Active, "Agent not active");
        require(status.domain == domain, "Incorrect domain");
        status.flag = AgentFlag.Slashed;
        _agentMap[agent] = status;
        slashStatus[agent] = SlashStatus(true, prover);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    function agentStatus(address agent) external view returns (AgentStatus memory) {
        return _agentMap[agent];
    }
}
