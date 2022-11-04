// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";
import { LocalDomainContext } from "../context/LocalDomainContext.sol";
import { BondingManager } from "./BondingManager.sol";
import { GlobalNotaryRegistry } from "../registry/GlobalNotaryRegistry.sol";
import { GuardRegistry } from "../registry/GuardRegistry.sol";

contract BondingPrimary is LocalDomainContext, GlobalNotaryRegistry, GuardRegistry, BondingManager {
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
        if (!_addNotary(_domain, _notary)) return;
        // Forward information that a Notary "staked" their bond to relevant local contracts
        // Forward information about the added Notary to remote chains (PINGs)
        // See: this.bondNotary() for handling PONGs
        _localUpdateNotary({
            _selector: SystemContract.bondNotary.selector,
            _domain: _domain,
            _notary: _notary,
            _callOrigin: 0, // there was no system call that initiated the bonding
            _forwardUpdate: true
        });
    }

    function removeNotary(uint32 _domain, address _notary) external onlyOwner {
        // Remove a Notary, break execution if they are not currently active
        if (!_removeNotary(_domain, _notary)) return;
        // Pass information that a Notary "unstaked" their bond to relevant local contracts
        // Forward information about the removed Notary to remote chains (PINGs)
        // See: this.unbondNotary() for handling PONGs
        _localUpdateNotary({
            _selector: SystemContract.unbondNotary.selector,
            _domain: _domain,
            _notary: _notary,
            _callOrigin: 0, // there was no system call that initiated the bonding
            _forwardUpdate: true
        });
    }

    function addGuard(address _guard) external onlyOwner {
        // Add a Guard, break execution if they are already active
        if (!_addGuard(_guard)) return;
        // Pass information that a Guard "staked" their bond to relevant local contracts
        // Forward information about the added Guard to remote chains (PINGs)
        // See: this.bondGuard() for handling PONGs
        _localUpdateGuard({
            _selector: SystemContract.bondGuard.selector,
            _guard: _guard,
            _callOrigin: 0, // there was no system call that initiated the bonding
            _forwardUpdate: true
        });
    }

    function removeGuard(address _guard) external onlyOwner {
        // Remove a Guard, break execution if they are not currently active
        if (!_removeGuard(_guard)) return;
        // Pass information that Guard "unstaked" their bond to relevant local contracts
        // Forward information about the removed Guard to remote chains (PINGs)
        // See: this.unbondGuard() for handling PONGs
        _localUpdateGuard({
            _selector: SystemContract.unbondGuard.selector,
            _guard: _guard,
            _callOrigin: 0, // there was no system call that initiated the bonding
            _forwardUpdate: true
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating that a new Notary staked a bond.
     * @dev Must be called from the BondingManager on remote chain, indicating
     * BondingSecondary has successfully handled a new Notary bonding.
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
        onlyCallers(BONDING_MANAGER, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // TODO: handle "notary added" PONG
    }

    /**
     * @notice Receive a system call indicating that an active Notary unstaked their bond.
     * @dev Must be called from the BondingManager on remote chain, indicating
     * BondingSecondary has successfully handled an active Notary unbonding.
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
        onlyCallers(BONDING_MANAGER, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // TODO: handle "notary removed" PONG
    }

    /**
     * @notice Receive a system call indicating that a new Guard staked a bond.
     * @dev Must be called from the BondingManager on remote chain, indicating
     * BondingSecondary has successfully handled a new Guard bonding.
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
        onlyCallers(BONDING_MANAGER, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // TODO: handle "guard added" PONG
    }

    /**
     * @notice Receive a system call indicating that an active Guard unstaked their bond.
     * @dev Must be called from the BondingManager on remote chain, indicating
     * BondingSecondary has successfully handled an active Notary unbonding.
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
        onlyCallers(BONDING_MANAGER, _caller)
        onlyOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD)
    {
        // TODO: handle "guard removed" PONG
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
        uint32,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) internal view override {
        // Slashing system call has to be done by Bonding Manager
        _assertEntityAllowed(BONDING_MANAGER, _caller);
        // Optimistic period should be over
        _assertOptimisticPeriodOver(_rootSubmittedAt, BONDING_OPTIMISTIC_PERIOD);
    }
}
