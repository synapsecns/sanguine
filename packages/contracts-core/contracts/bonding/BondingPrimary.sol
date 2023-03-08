// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { BondingManager } from "./BondingManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { InterfaceSystemRouter } from "../interfaces/InterfaceSystemRouter.sol";
import { ISystemContract } from "../interfaces/ISystemContract.sol";
import { AgentRegistry } from "../system/AgentRegistry.sol";

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
        // there was no system call that initiated the bonding => callOrigin == 0
        _syncAgentLocalRegistries({ _info: AgentInfo(_domain, _account, true), _callOrigin: 0 });
    }

    function removeAgent(uint32 _domain, address _account) external onlyOwner {
        // Remove an Agent, break execution if they are not currently active
        if (!_removeAgent(_domain, _account)) return;
        // there was no system call that initiated the unbonding => callOrigin == 0
        _syncAgentLocalRegistries({ _info: AgentInfo(_domain, _account, false), _callOrigin: 0 });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISystemContract
    function syncAgent(
        uint256,
        uint32,
        SystemEntity,
        AgentInfo memory
    ) external view onlySystemRouter {
        revert("Disabled for BondingPrimary");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║          INTERNAL HELPERS: UPDATE AGENT (BOND/UNBOND/SLASH)          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
