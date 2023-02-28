// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ByteString, TypedMemView } from "./ByteString.sol";
import { MerkleList } from "./MerkleList.sol";
import { SummitAttestation } from "./SnapAttestation.sol";
import { State, StateLib } from "./State.sol";
import { SNAPSHOT_MAX_STATES, STATE_LENGTH } from "./Structures.sol";

/// @dev Snapshot is a memory view over a formatted snapshot payload: a list of states.
type Snapshot is bytes29;
/// @dev Attach library functions to Snapshot
using {
    SnapshotLib.unwrap,
    SnapshotLib.hash,
    SnapshotLib.state,
    SnapshotLib.statesAmount,
    SnapshotLib.height,
    SnapshotLib.root,
    SnapshotLib.toSummitAttestation
} for Snapshot global;

/// @dev Struct representing Snapshot, as it is stored in the Summit contract.
/// Summit contract is supposed to store states. Snapshot is a list of states,
/// so we are storing a list of references to already stored states.
struct SummitSnapshot {
    // TODO: compress this - indexes might as well be uint32/uint64
    uint256[] statePtrs;
}
/// @dev Attach library functions to SummitSnapshot
using { SnapshotLib.getStatesAmount, SnapshotLib.getStatePtr } for SummitSnapshot global;

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
    ▏*║                           SUMMIT SNAPSHOT                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function toSummitSnapshot(uint256[] memory _statePtrs)
        internal
        pure
        returns (SummitSnapshot memory snapshot)
    {
        snapshot.statePtrs = _statePtrs;
    }

    function getStatesAmount(SummitSnapshot memory _snapshot) internal pure returns (uint256) {
        return _snapshot.statePtrs.length;
    }

    function getStatePtr(SummitSnapshot memory _snapshot, uint256 _index)
        internal
        pure
        returns (uint256)
    {
        require(_index < getStatesAmount(_snapshot), "Out of range");
        return _snapshot.statePtrs[_index];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SNAPSHOT HASHING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns the hash of a Snapshot, that could be later signed by an Agent.
    function hash(Snapshot _snapshot) internal pure returns (bytes32 hashedSnapshot) {
        // Get the underlying memory view
        bytes29 _view = _snapshot.unwrap();
        // TODO: include Snapshot-unique salt in the hash
        return _view.keccak();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SNAPSHOT SLICING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a state with a given index from the snapshot.
    function state(Snapshot _snapshot, uint256 _stateIndex) internal pure returns (State) {
        bytes29 _view = _snapshot.unwrap();
        uint256 indexFrom = _stateIndex * STATE_LENGTH;
        require(indexFrom < _view.len(), "Out of range");
        return _view.slice({ _index: indexFrom, _len: STATE_LENGTH, newType: 0 }).castToState();
    }

    /// @notice Returns the amount of states in the snapshot.
    function statesAmount(Snapshot _snapshot) internal pure returns (uint256) {
        bytes29 _view = _snapshot.unwrap();
        return _view.len() / STATE_LENGTH;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SNAPSHOT ROOT                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns the height of the extended "Snapshot Merkle Tree":
    /// every "state leaf" is in fact a node with two sub-leafs.
    /// @dev snapshot.height() is the length of the "extended merkle proof" for (root, origin) leaf:
    /// keccak256(metadata) is the first item in the "extended proof" list,
    /// followed by the remainder of the "merkle proof" from the "Snapshot Merkle Tree"
    function height(Snapshot _snapshot) internal pure returns (uint8 treeHeight) {
        // Account for the fact that every "state leaf" is a node with two sub-leafs
        treeHeight = 1;
        uint256 _statesAmount = _snapshot.statesAmount();
        for (uint256 amount = 1; amount < _statesAmount; amount <<= 1) {
            ++treeHeight;
        }
    }

    /// @notice Returns the root for the "Snapshot Merkle Tree" composed of state leafs from the snapshot.
    function root(Snapshot _snapshot) internal pure returns (bytes32) {
        uint256 _statesAmount = _snapshot.statesAmount();
        bytes32[] memory hashes = new bytes32[](_statesAmount);
        for (uint256 i = 0; i < _statesAmount; ++i) {
            // Each State has two sub-leafs, their hash is used as "leaf" in "Snapshot Merkle Tree"
            hashes[i] = _snapshot.state(i).hash();
        }
        MerkleList.calculateRoot(hashes);
        // hashes[0] now stores the value for the Merkle Root of the list
        return hashes[0];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SUMMIT ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns an Attestation struct to save in the Summit contract.
    /// Current block number and timestamp are used.
    function toSummitAttestation(Snapshot _snapshot)
        internal
        view
        returns (SummitAttestation memory attestation)
    {
        attestation.root = _snapshot.root();
        attestation.height = _snapshot.height();
        attestation.blockNumber = uint40(block.number);
        attestation.timestamp = uint40(block.timestamp);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Checks if snapshot's states amount is valid.
    function _isValidAmount(uint256 _statesAmount) internal pure returns (bool) {
        // Need to have at least one state in a snapshot.
        // Also need to have no more than `SNAPSHOT_MAX_STATES` states in a snapshot.
        return _statesAmount > 0 && _statesAmount <= SNAPSHOT_MAX_STATES;
    }
}
