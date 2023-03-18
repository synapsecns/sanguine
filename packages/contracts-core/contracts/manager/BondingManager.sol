// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { AgentInfo } from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager } from "./AgentManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { Versioned } from "../Version.sol";

/// @notice BondingManager keeps track of all existing agents.
/// Used on the Synapse Chain, serves as the "source of truth" for LightManagers on remote chains.
contract BondingManager is Versioned, AgentManager {
    constructor(uint32 _domain) DomainContext(_domain) Versioned("0.0.3") {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    function initialize() external initializer {
        __SystemContract_initialize();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            ADDING AGENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: replace with actual token staking/unstaking

    function addAgent(uint32 _domain, address _account) external onlyOwner returns (bool isAdded) {
        isAdded = _addAgent(_domain, _account);
        if (isAdded) {
            _syncAgentLocalRegistries(AgentInfo(_domain, _account, true));
        }
    }

    function removeAgent(uint32 _domain, address _account)
        external
        onlyOwner
        returns (bool isRemoved)
    {
        isRemoved = _removeAgent(_domain, _account);
        if (isRemoved) {
            _syncAgentLocalRegistries(AgentInfo(_domain, _account, false));
        }
    }
}
