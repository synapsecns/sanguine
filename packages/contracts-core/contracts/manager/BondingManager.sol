// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentStatus } from "../libs/Structures.sol";
import { DynamicTree } from "../libs/Merkle.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager, IAgentManager, ISystemRegistry } from "./AgentManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { Versioned } from "../Version.sol";

/// @notice BondingManager keeps track of all existing agents.
/// Used on the Synapse Chain, serves as the "source of truth" for LightManagers on remote chains.
contract BondingManager is Versioned, AgentManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // (agent => agent information)
    mapping(address => AgentStatus) public agentStatus;

    // A list of all agent accounts. First entry is address(0) to make agent indexes start from 1.
    address[] private agents;

    // Merkle Tree for Agents.
    // leafs[0] = 0
    // leafs[index > 0] = keccak(agentFlag, domain, agents[index])
    DynamicTree private agentTree;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain) DomainContext(_domain) Versioned("0.0.3") {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    function initialize(ISystemRegistry _origin, ISystemRegistry _destination)
        external
        initializer
    {
        __AgentManager_init(_origin, _destination);
        __Ownable_init();
        // Insert a zero address to make indexes for Agents start from 1.
        // Zeroed index is supposed to be used as a sentinel value meaning "no agent".
        agents.push(address(0));
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
