// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { SystemContract } from "./SystemContract.sol";
import { SystemRegistryEvents } from "../events/SystemRegistryEvents.sol";
import { AgentStatus, IAgentManager } from "../interfaces/IAgentManager.sol";
import { ISystemRegistry } from "../interfaces/ISystemRegistry.sol";

/// @notice Shared utilities for Origin, Destination/Summit contracts.
/// This abstract contract is responsible for all interactions with the local AgentManager,
/// where all agent are being tracked.
abstract contract SystemRegistry is SystemContract, SystemRegistryEvents, ISystemRegistry {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              IMMUTABLES                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    IAgentManager public immutable agentManager;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev gap for upgrade safety
    uint256[50] private __GAP; // solhint-disable-line var-name-mixedcase

    modifier onlyAgentManager() {
        require(msg.sender == address(agentManager), "!agentManager");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(IAgentManager agentManager_) {
        agentManager = agentManager_;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ONLY AGENT MANAGER                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISystemRegistry
    function managerSlash(
        uint32 domain,
        address agent,
        address prover
    ) external onlyAgentManager {
        _processSlashed(domain, agent, prover);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISystemRegistry
    function agentStatus(address agent) external view returns (AgentStatus memory) {
        return _agentStatus(agent);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Child contract could define custom logic for processing the slashed Agent.
    /// This will be called when the slashing was initiated in this contract or elsewhere.
    function _processSlashed(
        uint32 domain,
        address agent,
        address prover
    ) internal virtual {
        emit AgentSlashed(domain, agent, prover);
    }

    /// @dev This function should be called when the agent is proven to commit fraud in this contract.
    function _slashAgent(uint32 domain, address agent) internal {
        // Prover is msg.sender
        _processSlashed(domain, agent, msg.sender);
        agentManager.registrySlash(domain, agent, msg.sender);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns status of the given agent: (flag, domain, index).
    function _agentStatus(address agent) internal view returns (AgentStatus memory) {
        return agentManager.agentStatus(agent);
    }
}
