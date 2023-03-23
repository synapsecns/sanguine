// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { AgentInfo, SystemEntity } from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentRegistry } from "../system/AgentRegistry.sol";
import { ISystemContract, SystemContract } from "../system/SystemContract.sol";

/// @notice BondingManager keeps track of all agents.
abstract contract BondingManager is AgentRegistry, SystemContract {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISystemContract
    function slashAgent(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external onlySystemRouter {
        bool forwardUpdate;
        if (_callOrigin == localDomain) {
            // Forward information about slashed agent to remote chains
            forwardUpdate = true;
            // Only Origin can slash agents on local domain.
            // Summit is BondingManager on SynChain, so
            // Summit Notary slashing will not require a local slashAgent call.
            _assertEntityAllowed(ORIGIN, _caller);
        } else {
            // Forward information about slashed agent to remote chains
            // only if BondingManager is deployed on Synapse Chain
            forwardUpdate = _onSynapseChain();
            // Validate security params for cross-chain slashing
            _assertCrossChainSlashing(_rootSubmittedAt, _callOrigin, _caller);
        }
        // Forward information about the slashed agent to local Registries
        // Forward information about slashed agent to remote chains if needed
        _updateLocalRegistries(_dataSlashAgent(_info), forwardUpdate, _callOrigin);
    }

    /// @inheritdoc ISystemContract
    function syncAgent(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external onlySystemRouter {
        // BondingPrimary doesn't receive any valid syncAgent calls
        if (_onSynapseChain()) revert("Disabled for BondingPrimary");
        // Validate security params for cross-chain synching
        _assertCrossChainSynching(_rootSubmittedAt, _callOrigin, _caller);
        // Forward information about the synced agent to local Registries
        // Don't forward any information back to Synapse Chain
        _syncAgentLocalRegistries(_info);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║          INTERNAL HELPERS: UPDATE AGENT (BOND/UNBOND/SLASH)          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: generalize this further when Agent Merkle Tree is implemented

    /// @dev Passes an "update status" message to local Registries:
    /// that an Agent has been added / removed
    function _syncAgentLocalRegistries(AgentInfo memory _info) internal {
        // TODO: rework once Agent Merkle Tree is implemented
        // In the MVP version we don't do any forwarding for agents added/removed
        // Instead, BondingSecondary exposes owner-only addAgent() and removeAgent()
        _updateLocalRegistries(_dataSyncAgent(_info), false, 0);
    }

    /// @dev Passes an "update status" message to local Registries:
    /// that an Agent has been added / removed / slashed
    function _updateLocalRegistries(
        bytes memory _data,
        bool _forwardUpdate,
        uint32 _callOrigin
    ) internal {
        // Pass data to all System Registries. This could lead to duplicated data, meaning that
        // every Registry is responsible for ignoring the data it already has. This makes Registries
        // a bit more complex, but greatly reduces the complexity of BondingManager.
        systemRouter.systemMultiCall({
            _destination: localDomain,
            _optimisticSeconds: 0,
            _recipients: _localSystemRegistries(),
            _data: _data
        });
        // Forward data cross-chain, if requested
        if (_forwardUpdate) {
            _forwardUpdateData(_data, _callOrigin);
        }
    }

    /**
     * @notice Forward data with an agent status update (due to a system call from `_callOrigin`).
     * @dev If BondingManager is deployed on Synapse Chain, all chains should be notified,
     * excluding `_callOrigin` and Synapse Chain.
     * If BondingManager is not deployed on Synapse CHain, only Synapse Chain should be notified.
     */
    function _forwardUpdateData(bytes memory _data, uint32 _callOrigin) internal {
        if (_onSynapseChain()) {
            // SynapseChain: forward data to all OTHER chains except for callOrigin
            uint256 amount = amountDomains();
            for (uint256 i = 0; i < amount; ++i) {
                uint32 domain = getDomain(i);
                if (domain != _callOrigin && domain != SYNAPSE_DOMAIN) {
                    _callBondingManager(domain, BONDING_OPTIMISTIC_PERIOD, _data);
                }
            }
        } else {
            // Not Synapse Chain: forward data to Synapse Chain
            _callBondingManager(SYNAPSE_DOMAIN, BONDING_OPTIMISTIC_PERIOD, _data);
        }
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
    ) internal view {
        // Optimistic period should be over
        _assertOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD);
        // Either BondingManager is deployed on Synapse Chain, or
        // slashing system call has to originate on Synapse Chain
        if (!_onSynapseChain()) {
            _assertSynapseChain(_callOrigin);
        }
        // Slashing system call has to be done by Bonding Manager
        _assertEntityAllowed(BONDING_MANAGER, _caller);
    }

    /**
     * @notice Perform all required security checks for a cross-chain
     * system call for synching an agent.
     */
    function _assertCrossChainSynching(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller
    ) internal view {
        // Optimistic period should be over
        _assertOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD);
        // Synching system call has to originate on Synapse Chain
        _assertSynapseChain(_callOrigin);
        // Synching system call has to be done by Bonding Manager
        _assertEntityAllowed(BONDING_MANAGER, _caller);
    }

    /**
     * @notice Returns a list of local System Registries: system contracts, keeping track
     * of active Notaries and Guards.
     */
    function _localSystemRegistries() internal pure returns (SystemEntity[] memory recipients) {
        recipients = new SystemEntity[](2);
        recipients[0] = SystemEntity.Origin;
        recipients[1] = SystemEntity.Destination;
    }

    function _isIgnoredAgent(uint32, address) internal pure override returns (bool) {
        // Bonding keeps track of every agent
        return false;
    }
}
