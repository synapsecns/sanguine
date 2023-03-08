// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { BondingManager } from "./BondingManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { InterfaceSystemRouter } from "../interfaces/InterfaceSystemRouter.sol";
import { AgentRegistry } from "../system/AgentRegistry.sol";
import { SystemContract } from "../system/SystemContract.sol";

contract BondingPrimary is AgentRegistry, BondingManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain) DomainContext(_domain) {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Mocks for staking and unstaking of agents. Token locking/unlocking is omitted,
    // instead adding and removing agents are done by the contract owner.

    function addAgent(uint32 _domain, address _account) external onlyOwner {
        // Add an Agent, break execution if they are already active
        if (!_addAgent(_domain, _account)) return;
        _updateAgentStatus({ _domain: _domain, _agent: _account, _bonded: true });
    }

    function removeAgent(uint32 _domain, address _account) external onlyOwner {
        // Remove an Agent, break execution if they are not currently active
        if (!_removeAgent(_domain, _account)) return;
        _updateAgentStatus({ _domain: _domain, _agent: _account, _bonded: false });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║          INTERNAL HELPERS: UPDATE AGENT (BOND/UNBOND/SLASH)          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _updateAgentStatus(
        uint32 _domain,
        address _agent,
        bool _bonded
    ) internal {
        // Pass information about the new agent status to the local registries
        // Forward information about the new agent status to the remote chains
        // We optimistically expect the system message to be delivered,
        // and don't require sending a PONG back in the MVP.
        // This will be reworked once Agent Merkle Tree is implemented
        _updateLocalRegistries({
            _data: _dataSyncAgent(AgentInfo(_domain, _agent, _bonded)),
            _forwardUpdate: true,
            _callOrigin: 0 // there was no system call that initiated the bonding
        });
    }

    /**
     * @notice Forward data with an agent status update (due to
     * a system call from `_callOrigin`).
     * @dev If BondingManager is deployed on Synapse Chain, all other chains should be notified.
     * Otherwise, only Synapse Chain should be notified.
     */
    function _forwardUpdateData(bytes memory _data, uint32 _callOrigin) internal override {
        // TODO: forward update data to all chains except `_callOrigin`
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Perform all required security checks for a cross-chain
     * system call for slashing an agent.
     */
    function _assertCrossChainSlashing(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller
    ) internal view override {
        // Optimistic period should be over
        _assertOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD);
        // Slashing system call can originate on any chain
        _callOrigin;
        // Slashing system call has to be done by Bonding Manager
        _assertEntityAllowed(BONDING_MANAGER, _caller);
    }

    function _isIgnoredAgent(uint32, address) internal pure override returns (bool) {
        // BondingPrimary doesn't ignore anything
        return false;
    }
}
