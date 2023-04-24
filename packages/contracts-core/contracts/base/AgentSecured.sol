// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecuredEvents} from "../events/AgentSecuredEvents.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {AgentStatus, IAgentSecured} from "../interfaces/IAgentSecured.sol";
import {MessagingBase} from "./MessagingBase.sol";

abstract contract AgentSecured is MessagingBase, AgentSecuredEvents, IAgentSecured {
    // ════════════════════════════════════════════════ IMMUTABLES ═════════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    address public immutable agentManager;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev gap for upgrade safety
    uint256[50] private __GAP; // solhint-disable-line var-name-mixedcase

    modifier onlyAgentManager() {
        require(msg.sender == agentManager, "!agentManager");
        _;
    }

    constructor(string memory version_, uint32 localDomain_, address agentManager_)
        MessagingBase(version_, localDomain_)
    {
        agentManager = agentManager_;
    }

    // ════════════════════════════════════════════ ONLY AGENT MANAGER ═════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    function managerSlash(uint32 domain, address agent, address prover) external onlyAgentManager {
        _processSlashed(domain, agent, prover);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    function agentStatus(address agent) external view returns (AgentStatus memory) {
        return _agentStatus(agent);
    }

    /// @inheritdoc IAgentSecured
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status) {
        return _getAgent(index);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Child contract could define custom logic for processing the slashed Agent.
    /// This will be called when the slashing was initiated in this contract or elsewhere.
    function _processSlashed(uint32 domain, address agent, address prover) internal virtual {
        emit AgentSlashed(domain, agent, prover);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns status of the given agent: (flag, domain, index).
    function _agentStatus(address agent) internal view returns (AgentStatus memory) {
        return IAgentManager(agentManager).agentStatus(agent);
    }

    /// @dev Returns agent and their status for a given agent index. Returns zero values for non existing indexes.
    function _getAgent(uint256 index) internal view returns (address agent, AgentStatus memory status) {
        return IAgentManager(agentManager).getAgent(index);
    }
}
