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

    // agent => (bool isSlashed, address prover)
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
    function agentStatus(address _agent) external view returns (AgentStatus memory status) {
        status = _agentStatus(_agent);
        // If agent was proven to commit fraud, but their slashing wasn't completed,
        // return the Fraudulent flag instead
        if (slashStatus[_agent].isSlashed && status.flag != AgentFlag.Slashed) {
            status.flag = AgentFlag.Fraudulent;
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Checks and initiates the slashing of an agent.
    /// Should be called, after one of registries confirmed fraud committed by the agent.
    function _initiateSlashing(
        uint32 _domain,
        address _agent,
        address _prover
    ) internal {
        // Check that Agent hasn't been already slashed
        require(!slashStatus[_agent].isSlashed, "Already slashed");
        // Check that agent is Active/Unstaking and that the domains match
        AgentStatus memory status = _agentStatus(_agent);
        require(
            (status.flag == AgentFlag.Active || status.flag == AgentFlag.Unstaking) &&
                status.domain == _domain,
            "Slashing could not be initiated"
        );
        slashStatus[_agent] = SlashStatus({ isSlashed: true, prover: _prover });
    }

    /// @dev Notifies a given set of local registries about the slashed agent.
    /// Set is defined by a bitmask, eg: DESTINATION | ORIGIN
    function _notifySlashing(
        uint256 _registryMask,
        uint32 _domain,
        address _agent,
        address _prover
    ) internal {
        // Notify Destination, if requested
        if (_registryMask & DESTINATION != 0) destination.managerSlash(_domain, _agent, _prover);
        // Notify Origin, if requested
        if (_registryMask & ORIGIN != 0) origin.managerSlash(_domain, _agent, _prover);
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

    /// @dev Returns the last known status for the agent from the Agent Merkle Tree.
    function _agentStatus(address _agent) internal view virtual returns (AgentStatus memory);
}
