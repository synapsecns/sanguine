// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager, IAgentManager, ISystemRegistry } from "./AgentManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { Versioned } from "../Version.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is Versioned, AgentManager {
    constructor(uint32 _domain) DomainContext(_domain) Versioned("0.0.3") {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    function initialize(ISystemRegistry _origin, ISystemRegistry _destination)
        external
        initializer
    {
        __AgentManager_init(_origin, _destination);
        __Ownable_init();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SLASHING LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function registrySlash(uint32 _domain, address _agent) external {
        // On chains other than Synapse Chain only Origin could slash Agents
        // TODO: add slashing logic
        if (msg.sender == address(origin)) {
            _removeAgent(_domain, _agent);
            destination.managerSlash(_domain, _agent);
            // TODO: issue a system call to BondingManager on SynChain
        } else {
            revert("Unauthorized caller");
        }
    }
}
