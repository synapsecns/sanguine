// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ByteString } from "./ByteString.sol";
import { STATE_REPORT_SALT } from "./Constants.sol";
import { State, StateLib } from "./State.sol";
import { TypedMemView } from "./TypedMemView.sol";
/// @dev StateReport is a memory view over a formatted Guard report for invalid State
type StateReport is bytes29;
/// @dev Possible flags for the StateReport
/// Currently has only one possible value, but enables different types of reports in the future
enum StateFlag {
    Invalid
}
/// @dev Attach library functions to StateFlag
using { StateReportLib.formatStateReport } for StateFlag global;
/// @dev Attach library functions to StateReport
using {
    StateReportLib.hash,
    StateReportLib.unwrap,
    StateReportLib.flag,
    StateReportLib.state
} for StateReport global;

library StateReportLib {
    using ByteString for bytes;
    using StateLib for bytes29;
    using TypedMemView for bytes29;

    /**
     * @dev StateReport structure represents a Guard statement that a State is invalid.
     * State is considered invalid, if it doesn't match the saved state in Origin contract
     *  with the same nonce (or if nonce doesn't exist).
     *
     * Signed StateReport together with a proof that Notary used the reported State for their signed statement,
     * could be used on Destination and Summit to initiate a Dispute between the Guard and the Notary.
     *
     * @dev Memory layout of StateReport fields:
     * [000 .. 001): flag       uint8   1 byte      StateFlag for the report
     * [001 .. 051): state      uint8   1 byte      Raw payload for the reported State
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_FLAG = 0;
    uint256 private constant OFFSET_STATE = 1;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             STATE REPORT                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a formatted StateReport payload with provided fields.
    /// @param _flag            Flag signalling type of State Report
    /// @param _statePayload    Raw payload with reported state
    /// @return Formatted state report
    function formatStateReport(StateFlag _flag, bytes memory _statePayload)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_flag, _statePayload);
    }

    /// @notice Returns a StateReport view over the given payload
    /// @dev Will revert if the payload is not a state report.
    function castToStateReport(bytes memory _payload) internal pure returns (StateReport) {
        return castToStateReport(_payload.castToRawBytes());
    }

    /// @notice Casts a memory view to a StateReport view.
    /// @dev Will revert if if the memory view is not over a StateReport payload.
    function castToStateReport(bytes29 _view) internal pure returns (StateReport) {
        require(isStateReport(_view), "Not a state report");
        return StateReport.wrap(_view);
    }

    /// @notice Checks that a payload is a formatted StateReport.
    function isStateReport(bytes29 _view) internal pure returns (bool) {
        // Flag needs to exist
        if (_view.len() < OFFSET_STATE) return false;
        // Flag should fit into StateFlag enum
        if (_srFlag(_view) > uint8(type(StateFlag).max)) return false;
        // State should be properly formatted
        return _srState(_view).isState();
    }

    function hash(StateReport _stateReport) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 _view = _stateReport.unwrap();
        // The final hash to sign is keccak(stateReportSalt, keccak(stateReport))
        return keccak256(bytes.concat(STATE_REPORT_SALT, _view.keccak()));
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(StateReport _stateReport) internal pure returns (bytes29) {
        return StateReport.unwrap(_stateReport);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         STATE REPORT SLICING                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns StateFlag used in the report.
    function flag(StateReport _stateReport) internal pure returns (StateFlag) {
        bytes29 _view = _stateReport.unwrap();
        // We check that flag fits into enum, when payload is wrapped
        // into StateReport, so this never reverts
        return StateFlag(_srFlag(_view));
    }

    /// @notice Returns typed memory view over state used in the report.
    function state(StateReport _stateReport) internal pure returns (State) {
        bytes29 _view = _stateReport.unwrap();
        // We check that state is properly formatted, when payload is wrapped
        // into StateReport, so this never reverts.
        return _srState(_view).castToState();
    }

    /// @dev Returns StateReport flag without checking that it fits into StateFlag enum.
    function _srFlag(bytes29 _view) internal pure returns (uint8) {
        return uint8(_view.indexUint({ _index: OFFSET_FLAG, _bytes: 1 }));
    }

    /// @dev Returns an untyped memory view over Report's state without checking if it is properly formatted.
    function _srState(bytes29 _view) internal pure returns (bytes29) {
        return _view.sliceFrom({ _index: OFFSET_STATE, newType: 0 });
    }
}
