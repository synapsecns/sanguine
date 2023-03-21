// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { SystemContract } from "./SystemContract.sol";
import { SystemRegistryEvents } from "../events/SystemRegistryEvents.sol";
import { IAgentManager } from "../interfaces/IAgentManager.sol";
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
    function managerSlash(uint32 _domain, address _agent) external onlyAgentManager {
        _processSlashed(_domain, _agent);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISystemRegistry
    function isActiveAgent(address _account) external view returns (bool isActive, uint32 domain) {
        return _isActiveAgent(_account);
    }

    /// @inheritdoc ISystemRegistry
    function isActiveAgent(uint32 _domain, address _account) external view returns (bool) {
        return _isActiveAgent(_domain, _account);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Child contract could define custom logic for processing the slashed Agent.
    /// This will be called when the slashing was initiated in this contract or elsewhere.
    function _processSlashed(uint32 _domain, address _agent) internal virtual {}

    /// @dev This function should be called when the agent is proven to commit fraud in this contract.
    function _slashAgent(uint32 _domain, address _agent) internal {
        _processSlashed(_domain, _agent);
        agentManager.registrySlash(_domain, _agent);
        emit AgentSlashed(_domain, _agent);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Checks if the account is an active Agent on any of the domains.
    function _isActiveAgent(address _account) internal view returns (bool isActive, uint32 domain) {
        return agentManager.isActiveAgent(_account);
    }

    /// @dev Checks if the account is an active Agent on the given domain.
    function _isActiveAgent(uint32 _domain, address _account) internal view returns (bool) {
        return agentManager.isActiveAgent(_domain, _account);
    }
}
