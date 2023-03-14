// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    StateFlag,
    StateReport,
    StateReportLib,
    TypedMemView
} from "../../../contracts/libs/StateReport.sol";

/// @notice Exposes Report methods for testing against golang.
contract StateReportHarness {
    using StateReportLib for bytes;
    using StateReportLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         STATE REPORT GETTERS                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToStateReport(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        StateReport _stateReport = StateReportLib.castToStateReport(_payload);
        return _stateReport.unwrap().clone();
    }

    function hash(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToStateReport().hash();
    }

    function flag(bytes memory _payload) public pure returns (StateFlag) {
        return _payload.castToStateReport().flag();
    }

    function state(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToStateReport().state().unwrap().clone();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       STATE REPORT FORMATTERS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatStateReport(StateFlag _flag, bytes memory _statePayload)
        public
        pure
        returns (bytes memory)
    {
        return StateReportLib.formatStateReport(_flag, _statePayload);
    }

    function isStateReport(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isStateReport();
    }
}
