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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(IAgentManager _agentManager) {
        agentManager = _agentManager;
    }

    modifier onlyAgentManager() {
        require(msg.sender == address(agentManager), "!agentManager");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ONLY AGENT MANAGER                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISystemRegistry
    function managerSlash(
        uint32 _domain,
        address _agent,
        address _prover
    ) external onlyAgentManager {
        _processSlashed(_domain, _agent, _prover);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISystemRegistry
    function agentStatus(address _agent) external view returns (AgentStatus memory) {
        return _agentStatus(_agent);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Child contract could define custom logic for processing the slashed Agent.
    /// This will be called when the slashing was initiated in this contract or elsewhere.
    function _processSlashed(
        uint32 _domain,
        address _agent,
        address _prover
    ) internal virtual {}

    /// @dev This function should be called when the agent is proven to commit fraud in this contract.
    function _slashAgent(uint32 _domain, address _agent) internal {
        // Prover is msg.sender
        _processSlashed(_domain, _agent, msg.sender);
        agentManager.registrySlash(_domain, _agent, msg.sender);
        emit AgentSlashed(_domain, _agent, msg.sender);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns status of the given agent: (flag, domain, index).
    function _agentStatus(address _agent) internal view returns (AgentStatus memory) {
        return agentManager.agentStatus(_agent);
    }
}
