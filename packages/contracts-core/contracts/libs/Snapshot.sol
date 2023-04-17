// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {SNAPSHOT_MAX_STATES, SNAPSHOT_SALT, SNAPSHOT_TREE_HEIGHT, STATE_LENGTH} from "./Constants.sol";
import {MerkleList} from "./MerkleList.sol";
import {State, StateLib} from "./State.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev Snapshot is a memory view over a formatted snapshot payload: a list of states.
type Snapshot is bytes29;

/// @dev Attach library functions to Snapshot
using SnapshotLib for Snapshot global;

library SnapshotLib {
    using ByteString for bytes;
    using StateLib for bytes29;
    using TypedMemView for bytes29;

    /**
     * @dev Snapshot structure represents the state of multiple Origin contracts deployed on multiple chains.
     * In short, snapshot is a list of "State" structs. See State.sol for details about the "State" structs.
     *
     * Snapshot is considered "valid" in Origin, if every state referring to that Origin is valid there.
     * Snapshot is considered "globally valid", if it is "valid" in every Origin contract.
     *
     * Both Guards and Notaries are supposed to form snapshots and sign snapshot.hash() to verify its validity.
     * Each Guard should be monitoring a set of Origin contracts chosen as they see fit. They are expected
     * to form snapshots with Origin states for this set of chains, sign and submit them to Summit contract.
     *
     * Notaries are expected to monitor the Summit contract for new snapshots submitted by the Guards.
     * They should be forming their own snapshots using states from snapshots of any of the Guards.
     * The states for the Notary snapshots don't have to come from the same Guard snapshot,
     * or don't even have to be submitted by the same Guard.
     *
     * With their signature, Notary effectively "notarizes" the work that some Guards have done in Summit contract.
     * Notary signature on a snapshot doesn't only verify the validity of the Origins, but also serves as
     * a proof of liveliness for Guards monitoring these Origins.
     *
     * @dev Snapshot memory layout
     * [000 .. 050) states[0]   bytes   50 bytes
     * [050 .. 100) states[1]   bytes   50 bytes
     *      ..
     * [AAA .. BBB) states[N-1] bytes   50 bytes
     */

    // ═════════════════════════════════════════════════ SNAPSHOT ══════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Snapshot payload using a list of States.
     * @param states    Arrays of State-typed memory views over Origin states
     * @return Formatted snapshot
     */
    function formatSnapshot(State[] memory states) internal view returns (bytes memory) {
        require(_isValidAmount(states.length), "Invalid states amount");
        // First we unwrap State-typed views into generic views
        uint256 length = states.length;
        bytes29[] memory views = new bytes29[](length);
        for (uint256 i = 0; i < length; ++i) {
            views[i] = states[i].unwrap();
        }
        // Finally, we join them in a single payload. This avoids doing unnecessary copies in the process.
        return TypedMemView.join(views);
    }

    /**
     * @notice Returns a Snapshot view over for the given payload.
     * @dev Will revert if the payload is not a snapshot payload.
     */
    function castToSnapshot(bytes memory payload) internal pure returns (Snapshot) {
        return castToSnapshot(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Snapshot view.
     * @dev Will revert if the memory view is not over a snapshot payload.
     */
    function castToSnapshot(bytes29 view_) internal pure returns (Snapshot) {
        require(isSnapshot(view_), "Not a snapshot");
        return Snapshot.wrap(view_);
    }

    /**
     * @notice Checks that a payload is a formatted Snapshot.
     */
    function isSnapshot(bytes29 view_) internal pure returns (bool) {
        // Snapshot needs to have exactly N * STATE_LENGTH bytes length
        // N needs to be in [1 .. SNAPSHOT_MAX_STATES] range
        uint256 length = view_.len();
        uint256 statesAmount_ = length / STATE_LENGTH;
        return statesAmount_ * STATE_LENGTH == length && _isValidAmount(statesAmount_);
    }

    /// @notice Returns the hash of a Snapshot, that could be later signed by an Agent.
    function hash(Snapshot snapshot) internal pure returns (bytes32 hashedSnapshot) {
        // Get the underlying memory view
        bytes29 view_ = snapshot.unwrap();
        // The final hash to sign is keccak(attestationSalt, keccak(attestation))
        return keccak256(bytes.concat(SNAPSHOT_SALT, view_.keccak()));
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Snapshot snapshot) internal pure returns (bytes29) {
        return Snapshot.unwrap(snapshot);
    }

    // ═════════════════════════════════════════════ SNAPSHOT SLICING ══════════════════════════════════════════════════

    /// @notice Returns a state with a given index from the snapshot.
    function state(Snapshot snapshot, uint256 stateIndex) internal pure returns (State) {
        bytes29 view_ = snapshot.unwrap();
        uint256 indexFrom = stateIndex * STATE_LENGTH;
        require(indexFrom < view_.len(), "State index out of range");
        return view_.slice({index_: indexFrom, len_: STATE_LENGTH, newType: 0}).castToState();
    }

    /// @notice Returns the amount of states in the snapshot.
    function statesAmount(Snapshot snapshot) internal pure returns (uint256) {
        bytes29 view_ = snapshot.unwrap();
        return view_.len() / STATE_LENGTH;
    }

    /// @notice Returns the root for the "Snapshot Merkle Tree" composed of state leafs from the snapshot.
    function root(Snapshot snapshot) internal pure returns (bytes32) {
        uint256 statesAmount_ = snapshot.statesAmount();
        bytes32[] memory hashes = new bytes32[](statesAmount_);
        for (uint256 i = 0; i < statesAmount_; ++i) {
            // Each State has two sub-leafs, which are used as the "leafs" in "Snapshot Merkle Tree"
            // We save their parent in order to calculate the root for the whole tree later
            hashes[i] = snapshot.state(i).leaf();
        }
        // We are subtracting one here, as we already calculated the hashes
        // for the tree level above the "leaf level".
        MerkleList.calculateRoot(hashes, SNAPSHOT_TREE_HEIGHT - 1);
        // hashes[0] now stores the value for the Merkle Root of the list
        return hashes[0];
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Checks if snapshot's states amount is valid.
    function _isValidAmount(uint256 statesAmount_) internal pure returns (bool) {
        // Need to have at least one state in a snapshot.
        // Also need to have no more than `SNAPSHOT_MAX_STATES` states in a snapshot.
        return statesAmount_ > 0 && statesAmount_ <= SNAPSHOT_MAX_STATES;
    }
}
