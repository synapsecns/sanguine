// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {State, StateLib} from "./State.sol";
import {MemView, MemViewLib} from "./MemView.sol";
import {SNAPSHOT_MAX_STATES, SNAPSHOT_VALID_SALT, SNAPSHOT_TREE_HEIGHT, STATE_LENGTH} from "../Constants.sol";
import {IncorrectStatesAmount, IndexOutOfRange, UnformattedSnapshot} from "../Errors.sol";
import {MerkleMath} from "../merkle/MerkleMath.sol";
import {GasDataLib, ChainGas} from "../stack/GasData.sol";

/// Snapshot is a memory view over a formatted snapshot payload: a list of states.
type Snapshot is uint256;

using SnapshotLib for Snapshot global;

/// # Snapshot
/// Snapshot structure represents the state of multiple Origin contracts deployed on multiple chains.
/// In short, snapshot is a list of "State" structs. See State.sol for details about the "State" structs.
///
/// ## Snapshot usage
/// - Both Guards and Notaries are supposed to form snapshots and sign `snapshot.hash()` to verify its validity.
/// - Each Guard should be monitoring a set of Origin contracts chosen as they see fit.
///   - They are expected to form snapshots with Origin states for this set of chains,
///   sign and submit them to Summit contract.
/// - Notaries are expected to monitor the Summit contract for new snapshots submitted by the Guards.
///   - They should be forming their own snapshots using states from snapshots of any of the Guards.
///   - The states for the Notary snapshots don't have to come from the same Guard snapshot,
///   or don't even have to be submitted by the same Guard.
/// - With their signature, Notary effectively "notarizes" the work that some Guards have done in Summit contract.
///   - Notary signature on a snapshot doesn't only verify the validity of the Origins, but also serves as
///   a proof of liveliness for Guards monitoring these Origins.
///
/// ## Snapshot validity
/// - Snapshot is considered "valid" in Origin, if every state referring to that Origin is valid there.
/// - Snapshot is considered "globally valid", if it is "valid" in every Origin contract.
///
/// # Snapshot memory layout
///
/// | Position   | Field       | Type  | Bytes | Description                  |
/// | ---------- | ----------- | ----- | ----- | ---------------------------- |
/// | [000..050) | states[0]   | bytes | 50    | Origin State with index==0   |
/// | [050..100) | states[1]   | bytes | 50    | Origin State with index==1   |
/// | ...        | ...         | ...   | 50    | ...                          |
/// | [AAA..BBB) | states[N-1] | bytes | 50    | Origin State with index==N-1 |
///
/// @dev Snapshot could be signed by both Guards and Notaries and submitted to `Summit` in order to produce Attestations
/// that could be used in ExecutionHub for proving the messages coming from origin chains that the snapshot refers to.
library SnapshotLib {
    using MemViewLib for bytes;
    using StateLib for MemView;

    // ═════════════════════════════════════════════════ SNAPSHOT ══════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Snapshot payload using a list of States.
     * @param states    Arrays of State-typed memory views over Origin states
     * @return Formatted snapshot
     */
    function formatSnapshot(State[] memory states) internal view returns (bytes memory) {
        if (!_isValidAmount(states.length)) revert IncorrectStatesAmount();
        // First we unwrap State-typed views into untyped memory views
        uint256 length = states.length;
        MemView[] memory views = new MemView[](length);
        for (uint256 i = 0; i < length; ++i) {
            views[i] = states[i].unwrap();
        }
        // Finally, we join them in a single payload. This avoids doing unnecessary copies in the process.
        return MemViewLib.join(views);
    }

    /**
     * @notice Returns a Snapshot view over for the given payload.
     * @dev Will revert if the payload is not a snapshot payload.
     */
    function castToSnapshot(bytes memory payload) internal pure returns (Snapshot) {
        return castToSnapshot(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Snapshot view.
     * @dev Will revert if the memory view is not over a snapshot payload.
     */
    function castToSnapshot(MemView memView) internal pure returns (Snapshot) {
        if (!isSnapshot(memView)) revert UnformattedSnapshot();
        return Snapshot.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a payload is a formatted Snapshot.
     */
    function isSnapshot(MemView memView) internal pure returns (bool) {
        // Snapshot needs to have exactly N * STATE_LENGTH bytes length
        // N needs to be in [1 .. SNAPSHOT_MAX_STATES] range
        uint256 length = memView.len();
        uint256 statesAmount_ = length / STATE_LENGTH;
        return statesAmount_ * STATE_LENGTH == length && _isValidAmount(statesAmount_);
    }

    /// @notice Returns the hash of a Snapshot, that could be later signed by an Agent  to signal
    /// that the snapshot is valid.
    function hashValid(Snapshot snapshot) internal pure returns (bytes32 hashedSnapshot) {
        // The final hash to sign is keccak(snapshotSalt, keccak(snapshot))
        return snapshot.unwrap().keccakSalted(SNAPSHOT_VALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Snapshot snapshot) internal pure returns (MemView) {
        return MemView.wrap(Snapshot.unwrap(snapshot));
    }

    // ═════════════════════════════════════════════ SNAPSHOT SLICING ══════════════════════════════════════════════════

    /// @notice Returns a state with a given index from the snapshot.
    function state(Snapshot snapshot, uint256 stateIndex) internal pure returns (State) {
        MemView memView = snapshot.unwrap();
        uint256 indexFrom = stateIndex * STATE_LENGTH;
        if (indexFrom >= memView.len()) revert IndexOutOfRange();
        return memView.slice({index_: indexFrom, len_: STATE_LENGTH}).castToState();
    }

    /// @notice Returns the amount of states in the snapshot.
    function statesAmount(Snapshot snapshot) internal pure returns (uint256) {
        // Each state occupies exactly `STATE_LENGTH` bytes
        return snapshot.unwrap().len() / STATE_LENGTH;
    }

    /// @notice Extracts the list of ChainGas structs from the snapshot.
    function snapGas(Snapshot snapshot) internal pure returns (ChainGas[] memory snapGas_) {
        uint256 statesAmount_ = snapshot.statesAmount();
        snapGas_ = new ChainGas[](statesAmount_);
        for (uint256 i = 0; i < statesAmount_; ++i) {
            State state_ = snapshot.state(i);
            snapGas_[i] = GasDataLib.encodeChainGas(state_.gasData(), state_.origin());
        }
    }

    // ═════════════════════════════════════════ SNAPSHOT ROOT CALCULATION ═════════════════════════════════════════════

    /// @notice Returns the root for the "Snapshot Merkle Tree" composed of state leafs from the snapshot.
    function calculateRoot(Snapshot snapshot) internal pure returns (bytes32) {
        uint256 statesAmount_ = snapshot.statesAmount();
        bytes32[] memory hashes = new bytes32[](statesAmount_);
        for (uint256 i = 0; i < statesAmount_; ++i) {
            // Each State has two sub-leafs, which are used as the "leafs" in "Snapshot Merkle Tree"
            // We save their parent in order to calculate the root for the whole tree later
            hashes[i] = snapshot.state(i).leaf();
        }
        // We are subtracting one here, as we already calculated the hashes
        // for the tree level above the "leaf level".
        MerkleMath.calculateRoot(hashes, SNAPSHOT_TREE_HEIGHT - 1);
        // hashes[0] now stores the value for the Merkle Root of the list
        return hashes[0];
    }

    /// @notice Reconstructs Snapshot merkle Root from State Merkle Data (root + origin domain)
    /// and proof of inclusion of State Merkle Data (aka State "left sub-leaf") in Snapshot Merkle Tree.
    /// > Reverts if any of these is true:
    /// > - State index is out of range.
    /// > - Snapshot Proof length exceeds Snapshot tree Height.
    /// @param originRoot    Root of Origin Merkle Tree
    /// @param domain        Domain of Origin chain
    /// @param snapProof     Proof of inclusion of State Merkle Data into Snapshot Merkle Tree
    /// @param stateIndex    Index of Origin State in the Snapshot
    function proofSnapRoot(bytes32 originRoot, uint32 domain, bytes32[] memory snapProof, uint8 stateIndex)
        internal
        pure
        returns (bytes32)
    {
        // Index of "leftLeaf" is twice the state position in the snapshot
        // This is because each state is represented by two leaves in the Snapshot Merkle Tree:
        // - leftLeaf is a hash of (originRoot, originDomain)
        // - rightLeaf is a hash of (nonce, blockNumber, timestamp, gasData)
        uint256 leftLeafIndex = uint256(stateIndex) << 1;
        // Check that "leftLeaf" index fits into Snapshot Merkle Tree
        if (leftLeafIndex >= (1 << SNAPSHOT_TREE_HEIGHT)) revert IndexOutOfRange();
        bytes32 leftLeaf = StateLib.leftLeaf(originRoot, domain);
        // Reconstruct snapshot root using proof of inclusion
        // This will revert if snapshot proof length exceeds Snapshot Tree Height
        return MerkleMath.proofRoot(leftLeafIndex, leftLeaf, snapProof, SNAPSHOT_TREE_HEIGHT);
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Checks if snapshot's states amount is valid.
    function _isValidAmount(uint256 statesAmount_) internal pure returns (bool) {
        // Need to have at least one state in a snapshot.
        // Also need to have no more than `SNAPSHOT_MAX_STATES` states in a snapshot.
        return statesAmount_ != 0 && statesAmount_ <= SNAPSHOT_MAX_STATES;
    }
}
