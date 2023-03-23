// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager, IAgentManager, ISystemRegistry } from "./AgentManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { Versioned } from "../Version.sol";

/// @notice BondingManager keeps track of all existing agents.
/// Used on the Synapse Chain, serves as the "source of truth" for LightManagers on remote chains.
contract BondingManager is Versioned, AgentManager {
    constructor(uint32 _domain) DomainContext(_domain) Versioned("0.0.3") {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    function initialize(ISystemRegistry _origin, ISystemRegistry _destination)
        external
        initializer
    {
        __AgentManager_init(_origin, _destination);
        __Ownable_init();
    }

    // TODO: move addAgent into BondingManager and introduce Events,
    // when Agent Merkle Tree is implemented.

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SLASHING LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function registrySlash(uint32 _domain, address _agent) external {
        // On SynChain both Origin and Destination (Summit) could slash agents
        // TODO: add slashing logic
        if (msg.sender == address(origin)) {
            _removeAgent(_domain, _agent);
            destination.managerSlash(_domain, _agent);
        } else if (msg.sender == address(destination)) {
            _removeAgent(_domain, _agent);
            origin.managerSlash(_domain, _agent);
        } else {
            revert("Unauthorized caller");
        }
    }
}
