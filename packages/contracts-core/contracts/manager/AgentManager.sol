// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAgentManager } from "../interfaces/IAgentManager.sol";
import { ISystemRegistry } from "../interfaces/ISystemRegistry.sol";
import { SystemContract } from "../system/SystemContract.sol";

// TODO: adjust when Agent Merkle Tree is implemented
abstract contract AgentManager is SystemContract, IAgentManager {
    struct AgentInfo {
        bool isActive;
        uint32 domain;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    mapping(address => AgentInfo) private agentInfo;

    ISystemRegistry public origin;

    ISystemRegistry public destination;

    /// @dev gap for upgrade safety
    uint256[47] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function __AgentManager_init(ISystemRegistry _origin, ISystemRegistry _destination)
        internal
        onlyInitializing
    {
        origin = _origin;
        destination = _destination;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          AGENTS LOGIC (MVP)                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: remove these MVP functions once Agent Merkle Tree is implemented

    function addAgent(uint32 _domain, address _account) external onlyOwner returns (bool isAdded) {
        return _addAgent(_domain, _account);
    }

    function removeAgent(uint32 _domain, address _account)
        external
        onlyOwner
        returns (bool isRemoved)
    {
        return _removeAgent(_domain, _account);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function isActiveAgent(address _account) external view returns (bool isActive, uint32 domain) {
        return _isActiveAgent(_account);
    }

    /// @inheritdoc IAgentManager
    function isActiveAgent(uint32 _domain, address _account) external view returns (bool) {
        return _isActiveAgent(_domain, _account);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Adds agent to the domain, if they are not currently active on any of the domains.
    function _addAgent(uint32 _domain, address _agent) internal returns (bool wasAdded) {
        // Agent could only be active on one chain
        (bool isActive, ) = _isActiveAgent(_agent);
        if (!isActive) {
            agentInfo[_agent] = AgentInfo(true, _domain);
            wasAdded = true;
        }
    }

    /// @dev Removes agent from the domain, if they are currently active on this domain.
    function _removeAgent(uint32 _domain, address _agent) internal returns (bool wasRemoved) {
        // Agent needs to be active on exactly this domain
        bool isActive = _isActiveAgent(_domain, _agent);
        if (isActive) {
            delete agentInfo[_agent];
            wasRemoved = true;
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Checks if the account is an active Agent on any of the domains.
    function _isActiveAgent(address _account) internal view returns (bool isActive, uint32 domain) {
        AgentInfo memory info = agentInfo[_account];
        if (info.isActive) {
            isActive = true;
            domain = info.domain;
        }
    }

    /// @dev Checks if the account is an active Agent on the given domain.
    function _isActiveAgent(uint32 _domain, address _account) internal view returns (bool) {
        AgentInfo memory info = agentInfo[_account];
        return info.isActive && info.domain == _domain;
    }
}
