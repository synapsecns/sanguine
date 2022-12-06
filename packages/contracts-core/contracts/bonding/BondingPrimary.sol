// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";
import { LocalDomainContext } from "../context/LocalDomainContext.sol";
import { BondingManager } from "./BondingManager.sol";
import { AgentRegistry } from "../system/AgentRegistry.sol";

contract BondingPrimary is LocalDomainContext, AgentRegistry, BondingManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice id of the last "sync actors" request
    uint256 internal requestID;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain) LocalDomainContext(_domain) {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Mocks for staking and unstaking of agents. Token locking/unlocking is omitted,
    // instead adding and removing agents are done by the contract owner.

    function addNotary(uint32 _domain, address _notary) external onlyOwner {
        // Add a Notary, break execution if they are already active
        if (!_addAgent(_domain, _notary)) return;
        _updateAgentStatus(_notaryInfo({ _domain: _domain, _notary: _notary, _bonded: true }));
    }

    function removeNotary(uint32 _domain, address _notary) external onlyOwner {
        // Remove a Notary, break execution if they are not currently active
        if (!_removeAgent(_domain, _notary)) return;
        _updateAgentStatus(_notaryInfo({ _domain: _domain, _notary: _notary, _bonded: false }));
    }

    function addGuard(address _guard) external onlyOwner {
        // Add a Guard, break execution if they are already active
        if (!_addAgent({ _domain: 0, _account: _guard })) return;
        _updateAgentStatus(_guardInfo({ _guard: _guard, _bonded: true }));
    }

    function removeGuard(address _guard) external onlyOwner {
        // Remove a Guard, break execution if they are not currently active
        if (!_removeAgent({ _domain: 0, _account: _guard })) return;
        _updateAgentStatus(_guardInfo({ _guard: _guard, _bonded: false }));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating the list of off-chain agents needs to be synced.
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
        ISystemRouter.SystemEntity _caller,
        uint256 _requestID,
        bool _removeExisting,
        AgentInfo[] memory _infos
    )
        external
        override
        onlySystemRouter
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
        onlyCallers(BONDING_MANAGER, _caller)
    {
        // TODO: handle PONGs
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║          INTERNAL HELPERS: UPDATE AGENT (BOND/UNBOND/SLASH)          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _updateAgentStatus(AgentInfo memory _info) internal {
        // Increase the request counter and use it as the new request ID
        uint256 _requestID = ++requestID;
        // Construct the array with the given agent info
        // TODO: bulk bond/unbond requests in a single message
        AgentInfo[] memory infos = new AgentInfo[](1);
        infos[0] = _info;
        // Pass information about the new agent status to the local registries
        // Forward information about the new agent status to the remote chains (PINGs)
        // Existing agents don't need to be removed on remote chains
        // See: this.syncAgents() for handling PONGs
        _updateLocalRegistries({
            _data: _dataSyncAgents({
                _requestID: _requestID,
                _removeExisting: false,
                _infos: infos
            }),
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
        ISystemRouter.SystemEntity _caller
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
