// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {DisputeFlag, DisputeStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DisputeHubEvents} from "../events/DisputeHubEvents.sol";
import {IDisputeHub} from "../interfaces/IDisputeHub.sol";
import {SystemRegistry} from "../system/SystemRegistry.sol";

abstract contract DisputeHub is SystemRegistry, DisputeHubEvents, IDisputeHub {
    // TODO: Merge with ExecutionHub

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // (agent => their dispute status)
    mapping(address => DisputeStatus) internal _disputes;

    /// @dev gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    // ══════════════════════════════════════════ INITIATE DISPUTE LOGIC ═══════════════════════════════════════════════

    /// @inheritdoc IDisputeHub
    function openDispute(address guard, uint32 domain, address notary) external onlyAgentManager {
        _openDispute(guard, domain, notary);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IDisputeHub
    function disputeStatus(address agent) external view returns (DisputeStatus memory status) {
        return _disputes[agent];
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Opens a Dispute between a Guard and a Notary.
    /// This should be called, when the Guard submits a Report on a statement signed by the Notary.
    function _openDispute(address guard, uint32 domain, address notary) internal virtual {
        // Check that both agents are not in Dispute yet
        require(_disputes[guard].flag == DisputeFlag.None, "Guard already in dispute");
        require(_disputes[notary].flag == DisputeFlag.None, "Notary already in dispute");
        _disputes[guard] = DisputeStatus(DisputeFlag.Pending, notary);
        _disputes[notary] = DisputeStatus(DisputeFlag.Pending, guard);
        emit Dispute(guard, domain, notary);
    }

    /// @dev This is called when the slashing was initiated in this contract or elsewhere.
    function _processSlashed(uint32 domain, address agent, address prover) internal virtual override {
        _resolveDispute(domain, agent);
        super._processSlashed(domain, agent, prover);
    }

    /// @dev Resolves a Dispute for a slashed agent, if it hasn't been done already.
    function _resolveDispute(uint32 domain, address slashedAgent) internal virtual {
        DisputeStatus memory status = _disputes[slashedAgent];
        // Do nothing if dispute was already resolved
        if (status.flag == DisputeFlag.Slashed) return;
        // Update flag for the slashed agent
        // Slashed agent might have had no open Dispute, meaning the `counterpart` could be ZERO.
        // We still want to have the DisputeFlag.Slashed assigned in this case.
        _disputes[slashedAgent].flag = DisputeFlag.Slashed;
        // Delete record of dispute for the counterpart. This sets their Dispute Flag to None.
        if (status.counterpart != address(0)) delete _disputes[status.counterpart];
        // TODO: wo we want to use prover address if there was no counterpart?
        emit DisputeResolved(status.counterpart, domain, slashedAgent);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Checks if an agent is currently in Dispute.
    function _inDispute(address agent) internal view returns (bool) {
        return _disputes[agent].flag != DisputeFlag.None;
    }

    /// @dev Checks if an agent has been slashed.
    function _isSlashed(address agent) internal view returns (bool) {
        return _disputes[agent].flag == DisputeFlag.Slashed;
    }
}
