// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Snapshot } from "../libs/Snapshot.sol";
import { State, StateLib, SummitState } from "../libs/State.sol";

/**
 * @notice Hub to accept and save snapshots, as well as verify attestations.
 */
abstract contract SnapshotHub {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev All States submitted by any of the Guards
    SummitState[] private guardStates;

    /// @dev Pointer for the given State Leaf of the origin
    /// with ZERO as a sentinel value for "state not submitted yet".
    // (origin => (stateLeaf => {state index in guardStates PLUS 1}))
    mapping(uint32 => mapping(bytes32 => uint256)) private leafPtr;

    /// @dev Pointer for the latest Guard State of a given origin
    /// with ZERO as a sentinel value for "no states submitted yet".
    // (origin => (guard => {latest state index in guardStates PLUS 1}))
    mapping(uint32 => mapping(address => uint256)) private latestStatePtr;

    /// @dev gap for upgrade safety
    uint256[47] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns the state with the highest known nonce submitted by a given Guard.
     * @param _origin       Domain of origin chain
     * @param _guard        Guard address
     * @return stateData    Raw payload with guard latest state for origin
     */
    function getLatestState(uint32 _origin, address _guard)
        public
        view
        returns (bytes memory stateData)
    {
        SummitState memory latestState = _latestState(_origin, _guard);
        if (latestState.nonce == 0) return bytes("");
        return StateLib.formatState(_origin, latestState);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ACCEPT DATA                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Accepts a Snapshot signed by a Guard.
    /// It is assumed that the Guard signature has been checked outside of this contract.
    function _acceptGuardSnapshot(Snapshot _snapshot, address _guard) internal {
        // Snapshot Signer is a Guard: save the states for later use.
        uint256 statesAmount = _snapshot.statesAmount();
        for (uint256 i = 0; i < statesAmount; ++i) {
            _saveState(_snapshot.state(i), _guard);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         SAVE STATEMENT DATA                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Saves the state signed by a Guard.
    function _saveState(State _state, address _guard) internal returns (uint256 stateRef) {
        uint32 origin = _state.origin();
        // Check that Guard hasn't submitted a fresher State before
        require(_state.nonce() > _latestState(origin, _guard).nonce, "Outdated nonce");
        bytes32 leaf = _state.leaf();
        stateRef = leafPtr[origin][leaf];
        // Save state only if it wasn't previously submitted
        if (stateRef == 0) {
            // Extract data that needs to be saved
            SummitState memory state = _state.toSummitState();
            guardStates.push(state);
            // State is stored at (length - 1), but we are tracking "index PLUS 1" as "pointer"
            stateRef = guardStates.length;
            leafPtr[origin][leaf] = stateRef;
        }
        // Update latest guard state for origin
        latestStatePtr[origin][_guard] = stateRef;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CHECK STATEMENT DATA                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns the pointer for a matching Guard State, if it exists.
    function _stateRef(State _state) internal view returns (uint256) {
        return leafPtr[_state.origin()][_state.leaf()];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          LATEST STATE VIEWS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _latestState(uint32 _origin, address _guard)
        internal
        view
        returns (SummitState memory state)
    {
        // Get value for "index in guardStates PLUS 1"
        uint256 latestPtr = latestStatePtr[_origin][_guard];
        // Check if the Guard has submitted at least one State for origin
        if (latestPtr != 0) {
            state = guardStates[latestPtr - 1];
        }
        // An empty struct is returned if the Guard hasn't submitted a single State for origin yet.
    }
}
