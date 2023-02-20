// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./State.sol";

/// @dev Snapshot is a memory view over a formatted snapshot payload: a list of states.
type Snapshot is bytes29;
/// @dev Attach library functions to Snapshot
using {
    SnapshotLib.unwrap,
    SnapshotLib.hash,
    SnapshotLib.state,
    SnapshotLib.statesAmount
} for Snapshot global;

library SnapshotLib {
    using ByteString for bytes;
    using StateLib for bytes29;
    using TypedMemView for bytes29;

    /**
     * @dev Snapshot structure represent the historical state of multiple Origin contracts deployed on multiple chains.
     * In short, snapshot is a list of "State" structs.
     *
     * @dev Snapshot memory layout
     * [000 .. 050) states[0]   bytes   50 bytes
     * [050 .. 100) states[1]   bytes   50 bytes
     *      ..
     * [AAA .. BBB) states[N-1] bytes   50 bytes
     */

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               SNAPSHOT                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted Snapshot payload using a list of States.
     * @param _states   Arrays of State-typed memory views over Origin states
     * @return Formatted snapshot
     */
    function formatSnapshot(State[] memory _states) internal view returns (bytes memory) {
        require(_isValidAmount(_states.length), "Invalid states amount");
        // First we unwrap State-typed views into generic views
        uint256 length = _states.length;
        bytes29[] memory views = new bytes29[](length);
        for (uint256 i = 0; i < length; ++i) {
            views[i] = _states[i].unwrap();
        }
        // Finally, we join them in a single payload. This avoids doing unnecessary copies in the process.
        return TypedMemView.join(views);
    }

    /**
     * @notice Returns a Snapshot view over for the given payload.
     * @dev Will revert if the payload is not a snapshot payload.
     */
    function castToSnapshot(bytes memory _payload) internal pure returns (Snapshot) {
        return castToSnapshot(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Snapshot view.
     * @dev Will revert if the memory view is not over a snapshot payload.
     */
    function castToSnapshot(bytes29 _view) internal pure returns (Snapshot) {
        require(isSnapshot(_view), "Not a snapshot");
        return Snapshot.wrap(_view);
    }

    /**
     * @notice Checks that a payload is a formatted Snapshot.
     */
    function isSnapshot(bytes29 _view) internal pure returns (bool) {
        // Snapshot needs to have exactly N * STATE_LENGTH bytes length
        // N needs to be in [1 .. SNAPSHOT_MAX_STATES] range
        uint256 length = _view.len();
        uint256 _statesAmount = length / STATE_LENGTH;
        return _statesAmount * STATE_LENGTH == length && _isValidAmount(_statesAmount);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Snapshot _snapshot) internal pure returns (bytes29) {
        return Snapshot.unwrap(_snapshot);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SNAPSHOT HASHING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns the hash of a snapshot.
    function hash(Snapshot _snapshot) internal pure returns (bytes32 hashedSnapshot) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_snapshot);
        // TODO: include Snapshot-unique salt in the hash
        return _view.keccak();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SNAPSHOT SLICING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a state with a given index from the snapshot.
    function state(Snapshot _snapshot, uint256 _stateIndex) internal pure returns (State) {
        bytes29 _view = unwrap(_snapshot);
        uint256 indexFrom = _stateIndex * STATE_LENGTH;
        require(indexFrom < _view.len(), "Out of range");
        return _view.slice({ _index: indexFrom, _len: STATE_LENGTH, newType: 0 }).castToState();
    }

    /// @notice Returns the amount of states in the snapshot.
    function statesAmount(Snapshot _snapshot) internal pure returns (uint256) {
        bytes29 _view = unwrap(_snapshot);
        return _view.len() / STATE_LENGTH;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns if snapshot's states amount is valid.
    function _isValidAmount(uint256 _statesAmount) internal pure returns (bool) {
        // Need to have at least one state in a snapshot.
        // Also need to have no more than `SNAPSHOT_MAX_STATES` states in a snapshot.
        return _statesAmount > 0 && _statesAmount <= SNAPSHOT_MAX_STATES;
    }
}
