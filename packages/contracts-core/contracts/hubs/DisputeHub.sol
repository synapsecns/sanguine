// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { Snapshot, StatementHub, StateReport } from "./StatementHub.sol";
import { IDisputeHub } from "../interfaces/IDisputeHub.sol";

abstract contract DisputeHub is StatementHub, IDisputeHub {
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
}
