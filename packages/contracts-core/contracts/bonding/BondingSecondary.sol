// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { BondingManager } from "./BondingManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { InterfaceSystemRouter } from "../interfaces/InterfaceSystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";

contract BondingSecondary is BondingManager {
    constructor(uint32 _domain) DomainContext(_domain) {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating the list of off-chain agents needs to be synced.
     * @dev Must be called from the BondingManager on Synapse Chain.
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _requestID        Unique ID of the sync request
     * @param _removeExisting   Whether the existing agents need to be removed first
     * @param _infos            Information about a list of agents to sync
     */
    function syncAgents(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        uint256 _requestID,
        bool _removeExisting,
        AgentInfo[] memory _infos
    )
        external
        onlySystemRouter
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
        onlySynapseChainBondingManager(_callOrigin, _caller)
    {
        // Pass the list of agents to all local registries
        // Don't forward the same array back to Synapse Chain
        _updateLocalRegistries({
            _data: _dataSyncAgents(_requestID, _removeExisting, _infos),
            _forwardUpdate: false,
            _callOrigin: _callOrigin
        });
        // Report back to Synapse Chain that a request has been handled
        // Request ID will be used for identifying, so we could pass an empty array here
        _forwardUpdateData(
            _dataSyncAgents(_requestID, _removeExisting, new AgentInfo[](0)),
            _callOrigin
        );
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
    function _forwardUpdateData(bytes memory _data, uint32) internal override {
        systemRouter.systemCall({
            _destination: SYNAPSE_DOMAIN,
            _optimisticSeconds: BONDING_OPTIMISTIC_PERIOD,
            _recipient: SystemEntity.BondingManager,
            _data: _data
        });
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
        // Slashing system call has to originate on Synapse Chain
        _assertSynapseChain(_callOrigin);
        // Slashing system call has to be done by Bonding Manager
        _assertEntityAllowed(BONDING_MANAGER, _caller);
    }
}
