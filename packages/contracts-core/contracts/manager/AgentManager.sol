// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAgentManager } from "../interfaces/IAgentManager.sol";
import { ISystemRegistry } from "../interfaces/ISystemRegistry.sol";
import { AgentFlag, AgentStatus, SlashStatus } from "../libs/Structures.sol";
import { SystemContract } from "../system/SystemContract.sol";

// TODO: adjust when Agent Merkle Tree is implemented
abstract contract AgentManager is SystemContract, IAgentManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    ISystemRegistry public origin;

    ISystemRegistry public destination;

    // agent => (bool isSlashed, address slashedBy)
    mapping(address => SlashStatus) public slashStatus;

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
    ▏*║                            EXTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function agentRoot() external view virtual returns (bytes32);

    /// @inheritdoc IAgentManager
    function isActiveAgent(address _account) external view returns (bool isActive, uint32 domain) {
        return _isActiveAgent(_account);
    }

    /// @inheritdoc IAgentManager
    function isActiveAgent(uint32 _domain, address _account) external view returns (bool) {
        return _isActiveAgent(_domain, _account);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Generates leaf to be saved in the Agent Merkle Tree
    function _agentLeaf(
        AgentFlag _flag,
        uint32 _domain,
        address _agent
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_flag, _domain, _agent));
    }

    /// @dev Returns the last known status for the agent.
    function _agentStatus(address _agent) internal view virtual returns (AgentStatus memory);

    /// @dev Checks if the account is an active Agent on any of the domains.
    function _isActiveAgent(address _account)
        internal
        view
        virtual
        returns (bool isActive, uint32 domain)
    {
        AgentStatus memory status = _agentStatus(_account);
        if (status.flag == AgentFlag.Active && !slashStatus[_account].isSlashed) {
            isActive = true;
            domain = status.domain;
        }
    }

    /// @dev Checks if the account is an active Agent on the given domain.
    function _isActiveAgent(uint32 _domain, address _account) internal view virtual returns (bool) {
        AgentStatus memory status = _agentStatus(_account);
        return
            status.flag == AgentFlag.Active &&
            !slashStatus[_account].isSlashed &&
            status.domain == _domain;
    }
}
