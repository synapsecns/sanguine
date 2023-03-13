// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { Snapshot, StatementHub, StateReport } from "./StatementHub.sol";
import { DisputeHubEvents } from "../events/DisputeHubEvents.sol";
import { IDisputeHub } from "../interfaces/IDisputeHub.sol";

abstract contract DisputeHub is StatementHub, DisputeHubEvents, IDisputeHub {
    /// @inheritdoc IDisputeHub
    function submitStateReport(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes memory _snapPayload,
        bytes memory _snapSignature
    ) external returns (bool wasAccepted) {
        // TODO: implement
    }

    /// @inheritdoc IDisputeHub
    function submitStateReportWithProof(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes32[] memory _snapProof,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool wasAccepted) {
        // TODO: implement
    }

    /// @dev Opens a Dispute between a Guard and a Notary.
    /// This should be called, when the Guard submits a Report on a statement signed by the Notary.
    function _openDispute(
        address _guard,
        uint32 _domain,
        address _notary
    ) internal {
        // TODO: implement this
        emit Dispute(_guard, _domain, _notary);
    }
}
