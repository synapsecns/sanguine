// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {STATE_REPORT_SALT} from "./Constants.sol";
import {State, StateLib} from "./State.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// @dev StateReport is a memory view over a formatted Guard report for invalid State
type StateReport is uint256;

/// @dev Possible flags for the StateReport
/// Currently has only one possible value, but enables different types of reports in the future
enum StateFlag {Invalid}

/// @dev Attach library functions to StateFlag
using StateReportLib for StateFlag global;
/// @dev Attach library functions to StateReport
using StateReportLib for StateReport global;

library StateReportLib {
    using MemViewLib for bytes;
    using StateLib for MemView;

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

    // ═══════════════════════════════════════════════ STATE REPORT ════════════════════════════════════════════════════

    /// @notice Returns a formatted StateReport payload with provided fields.
    /// @param flag_            Flag signalling type of State Report
    /// @param statePayload     Raw payload with reported state
    /// @return Formatted state report
    function formatStateReport(StateFlag flag_, bytes memory statePayload) internal pure returns (bytes memory) {
        return abi.encodePacked(flag_, statePayload);
    }

    /// @notice Returns a StateReport view over the given payload
    /// @dev Will revert if the payload is not a state report.
    function castToStateReport(bytes memory payload) internal pure returns (StateReport) {
        return castToStateReport(payload.ref());
    }

    /// @notice Casts a memory view to a StateReport view.
    /// @dev Will revert if if the memory view is not over a StateReport payload.
    function castToStateReport(MemView memView) internal pure returns (StateReport) {
        require(isStateReport(memView), "Not a state report");
        return StateReport.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted StateReport.
    function isStateReport(MemView memView) internal pure returns (bool) {
        // Flag needs to exist
        if (memView.len() < OFFSET_STATE) return false;
        // Flag should fit into StateFlag enum
        if (_srFlag(memView) > uint8(type(StateFlag).max)) return false;
        // State should be properly formatted
        return _srState(memView).isState();
    }

    function hash(StateReport stateReport) internal pure returns (bytes32) {
        // Get the underlying memory view
        MemView memView = stateReport.unwrap();
        // The final hash to sign is keccak(stateReportSalt, keccak(stateReport))
        return keccak256(bytes.concat(STATE_REPORT_SALT, memView.keccak()));
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(StateReport stateReport) internal pure returns (MemView) {
        return MemView.wrap(StateReport.unwrap(stateReport));
    }

    // ═══════════════════════════════════════════ STATE REPORT SLICING ════════════════════════════════════════════════

    /// @notice Returns StateFlag used in the report.
    function flag(StateReport stateReport) internal pure returns (StateFlag) {
        MemView memView = stateReport.unwrap();
        // We check that flag fits into enum, when payload is wrapped
        // into StateReport, so this never reverts
        return StateFlag(_srFlag(memView));
    }

    /// @notice Returns typed memory view over state used in the report.
    function state(StateReport stateReport) internal pure returns (State) {
        MemView memView = stateReport.unwrap();
        // We check that state is properly formatted, when payload is wrapped
        // into StateReport, so this never reverts.
        return _srState(memView).castToState();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns StateReport flag without checking that it fits into StateFlag enum.
    function _srFlag(MemView memView) internal pure returns (uint8) {
        return uint8(memView.indexUint({index_: OFFSET_FLAG, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over Report's state without checking if it is properly formatted.
    function _srState(MemView memView) internal pure returns (MemView) {
        return memView.sliceFrom({index_: OFFSET_STATE});
    }
}
