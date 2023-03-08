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
