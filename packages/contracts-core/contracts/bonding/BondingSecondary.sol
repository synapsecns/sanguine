// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";
import { LocalDomainContext } from "../context/LocalDomainContext.sol";
import { BondingManager } from "./BondingManager.sol";
import { GlobalNotaryRegistry } from "../registry/GlobalNotaryRegistry.sol";
import { GuardRegistry } from "../registry/GuardRegistry.sol";

contract BondingSecondary is
    LocalDomainContext,
    GlobalNotaryRegistry,
    GuardRegistry,
    BondingManager
{
    constructor(uint32 _domain) LocalDomainContext(_domain) {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating that a new Notary staked a bond.
     * @dev Must be called from the BondingManager on Synapse Chain.
     * @param _domain           Domain where the new Notary will be active
     * @param _notary           New Notary that staked a bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function bondNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    )
        external
        override
        onlySystemRouter
        onlySynapseChainBondingManager(_callOrigin, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // Pass information that Notary staked their bond to relevant local contracts
        // Report back to Synapse Chain that a request has been handled
        _localUpdateNotary({
            _selector: SystemContract.bondNotary.selector,
            _domain: _domain,
            _notary: _notary,
            _callOrigin: _callOrigin,
            _forwardUpdate: true
        });
    }

    /**
     * @notice Receive a system call indicating that an active Notary unstaked their bond.
     * @dev Must be called from the BondingManager on Synapse Chain.
     * @param _domain           Domain where the Notary was active
     * @param _notary           Active Notary that unstaked their bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function unbondNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    )
        external
        override
        onlySystemRouter
        onlySynapseChainBondingManager(_callOrigin, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // Pass information that Notary unstaked their bond to relevant local contracts
        // Report back to Synapse Chain that a request has been handled
        _localUpdateNotary({
            _selector: SystemContract.unbondNotary.selector,
            _domain: _domain,
            _notary: _notary,
            _callOrigin: _callOrigin,
            _forwardUpdate: true
        });
    }

    /**
     * @notice Receive a system call indicating that a new Guard staked a bond.
     * @dev Must be called from the BondingManager on Synapse Chain.
     * @param _guard            New Guard that staked a bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function bondGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    )
        external
        override
        onlySystemRouter
        onlySynapseChainBondingManager(_callOrigin, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // Pass information that Guard staked their bond to relevant local contracts
        // Report back to Synapse Chain that a request has been handled
        _localUpdateGuard({
            _selector: SystemContract.bondGuard.selector,
            _guard: _guard,
            _callOrigin: _callOrigin,
            _forwardUpdate: true
        });
    }

    /**
     * @notice Receive a system call indicating that an active Guard unstaked their bond.
     * @dev Must be called from the BondingManager on Synapse Chain.
     * @param _guard            Active Guard that unstaked their bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function unbondGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    )
        external
        override
        onlySystemRouter
        onlySynapseChainBondingManager(_callOrigin, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // Pass information that Guard unstaked their bond to relevant local contracts
        // Report back to Synapse Chain that a request has been handled
        _localUpdateGuard({
            _selector: SystemContract.unbondGuard.selector,
            _guard: _guard,
            _callOrigin: _callOrigin,
            _forwardUpdate: true
        });
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
            _recipient: ISystemRouter.SystemEntity.BondingManager,
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
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) internal view override {
        // Slashing system call has to originate on Synapse Chain
        _assertSynapseChain(_callOrigin);
        // Slashing system call has to be done by Bonding Manager
        _assertEntityAllowed(BONDING_MANAGER, _caller);
        // Optimistic period should be over
        _assertOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD);
    }
}
