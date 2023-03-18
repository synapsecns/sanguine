// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { AgentInfo } from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager } from "./AgentManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { Versioned } from "../Version.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is Versioned, AgentManager {
    constructor(uint32 _domain) DomainContext(_domain) Versioned("0.0.3") {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    function initialize() external initializer {
        __SystemContract_initialize();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           OWNER ONLY (MVP)                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: remove these MVP functions once Agent Merkle Tree is implemented

    function addAgent(uint32 _domain, address _account) external onlyOwner {
        // Add an Agent, break execution if they are already active
        if (!_addAgent(_domain, _account)) return;
        // bonded = true
        _syncAgentLocalRegistries(AgentInfo(_domain, _account, true));
    }

    function removeAgent(uint32 _domain, address _account) external onlyOwner {
        // Remove an Agent, break execution if they are not currently active
        if (!_removeAgent(_domain, _account)) return;
        // bonded = false
        _syncAgentLocalRegistries(AgentInfo(_domain, _account, false));
    }
}
