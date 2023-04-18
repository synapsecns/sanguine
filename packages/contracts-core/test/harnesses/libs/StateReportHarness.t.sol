// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {StateFlag, StateReport, StateReportLib, MemView, MemViewLib} from "../../../contracts/libs/StateReport.sol";

// solhint-disable ordering
/// @notice Exposes Report methods for testing against golang.
contract StateReportHarness {
    using StateReportLib for bytes;
    using StateReportLib for MemView;
    using MemViewLib for bytes;

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToStateReport(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        StateReport stateReport = StateReportLib.castToStateReport(payload);
        return stateReport.unwrap().clone();
    }

    function hash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToStateReport().hash();
    }

    function flag(bytes memory payload) public pure returns (StateFlag) {
        return payload.castToStateReport().flag();
    }

    function state(bytes memory payload) public view returns (bytes memory) {
        return payload.castToStateReport().state().unwrap().clone();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatStateReport(StateFlag flag_, bytes memory statePayload) public pure returns (bytes memory) {
        return StateReportLib.formatStateReport(flag_, statePayload);
    }

    function isStateReport(bytes memory payload) public pure returns (bool) {
        return payload.ref().isStateReport();
    }
}
