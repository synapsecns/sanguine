// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";
import { AbstractGuardRegistry } from "../registry/AbstractGuardRegistry.sol";
import { AbstractNotaryRegistry } from "../registry/AbstractNotaryRegistry.sol";

abstract contract BondingManager is AbstractGuardRegistry, AbstractNotaryRegistry, SystemContract {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
        // TODO: pass information that Notary staked their bond to relevant local contracts
    }

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
        // TODO: pass information that Notary unstaked their bond to relevant local contracts
    }

    function slashNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external override onlySystemRouter {
        if (_callOrigin == _localDomain()) {
            // Only Origin can slash agents on local domain
            _assertEntityAllowed(ORIGIN, _caller);
            // TODO: forward information about slashed Notary to SynChain
        } else {
            // Cross-chain slashing: have to originate on Synapse Chain
            _assertSynapseChain(_callOrigin);
            // Cross-chain slashing: have to be done by Bonding Manager
            _assertEntityAllowed(BONDING_MANAGER, _caller);
            // TODO: slash Notary on relevant local contracts
        }
    }

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
        // TODO: pass information that Guard staked their bond to relevant local contracts
    }

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
        // TODO: pass information that Guard unstaked their bond to relevant local contracts
    }

    function slashGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external override onlySystemRouter {
        if (_callOrigin == _localDomain()) {
            // Only Origin can slash agents on local domain
            _assertEntityAllowed(ORIGIN, _caller);
            // TODO: forward information about slashed Guard to SynChain
        } else {
            // Cross-chain slashing: have to originate on Synapse Chain
            _assertSynapseChain(_callOrigin);
            // Cross-chain slashing: have to be done by Bonding Manager
            _assertEntityAllowed(BONDING_MANAGER, _caller);
            // TODO: slash Guard on relevant local contracts
        }
    }
}
