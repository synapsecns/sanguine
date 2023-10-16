// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemView, MemViewLib} from "./MemView.sol";
import {GAS_DATA_LENGTH, STATE_LENGTH, STATE_INVALID_SALT} from "../Constants.sol";
import {UnformattedState} from "../Errors.sol";
import {GasData, GasDataLib} from "../stack/GasData.sol";

/// State is a memory view over a formatted state payload.
type State is uint256;

using StateLib for State global;

/// # State
/// State structure represents the state of Origin contract at some point of time.
/// - State is structured in a way to track the updates of the Origin Merkle Tree.
/// - State includes root of the Origin Merkle Tree, origin domain and some additional metadata.
/// ## Origin Merkle Tree
/// Hash of every sent message is inserted in the Origin Merkle Tree, which changes
/// the value of Origin Merkle Root (which is the root for the mentioned tree).
/// - Origin has a single Merkle Tree for all messages, regardless of their destination domain.
/// - This leads to Origin state being updated if and only if a message was sent in a block.
/// - Origin contract is a "source of truth" for states: a state is considered "valid" in its Origin,
/// if it matches the state of the Origin contract after the N-th (nonce) message was sent.
///
/// # Memory layout of State fields
///
/// | Position   | Field       | Type    | Bytes | Description                    |
/// | ---------- | ----------- | ------- | ----- | ------------------------------ |
/// | [000..032) | root        | bytes32 | 32    | Root of the Origin Merkle Tree |
/// | [032..036) | origin      | uint32  | 4     | Domain where Origin is located |
/// | [036..040) | nonce       | uint32  | 4     | Amount of sent messages        |
/// | [040..045) | blockNumber | uint40  | 5     | Block of last sent message     |
/// | [045..050) | timestamp   | uint40  | 5     | Time of last sent message      |
/// | [050..062) | gasData     | uint96  | 12    | Gas data for the chain         |
///
/// @dev State could be used to form a Snapshot to be signed by a Guard or a Notary.
library StateLib {
    using MemViewLib for bytes;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_ROOT = 0;
    uint256 private constant OFFSET_ORIGIN = 32;
    uint256 private constant OFFSET_NONCE = 36;
    uint256 private constant OFFSET_BLOCK_NUMBER = 40;
    uint256 private constant OFFSET_TIMESTAMP = 45;
    uint256 private constant OFFSET_GAS_DATA = 50;

    // ═══════════════════════════════════════════════════ STATE ═══════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted State payload with provided fields
     * @param root_         New merkle root
     * @param origin_       Domain of Origin's chain
     * @param nonce_        Nonce of the merkle root
     * @param blockNumber_  Block number when root was saved in Origin
     * @param timestamp_    Block timestamp when root was saved in Origin
     * @param gasData_      Gas data for the chain
     * @return Formatted state
     */
    function formatState(
        bytes32 root_,
        uint32 origin_,
        uint32 nonce_,
        uint40 blockNumber_,
        uint40 timestamp_,
        GasData gasData_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(root_, origin_, nonce_, blockNumber_, timestamp_, gasData_);
    }

    /**
     * @notice Returns a State view over the given payload.
     * @dev Will revert if the payload is not a state.
     */
    function castToState(bytes memory payload) internal pure returns (State) {
        return castToState(payload.ref());
    }

    /**
     * @notice Casts a memory view to a State view.
     * @dev Will revert if the memory view is not over a state.
     */
    function castToState(MemView memView) internal pure returns (State) {
        if (!isState(memView)) revert UnformattedState();
        return State.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted State.
    function isState(MemView memView) internal pure returns (bool) {
        return memView.len() == STATE_LENGTH;
    }

    /// @notice Returns the hash of a State, that could be later signed by a Guard to signal
    /// that the state is invalid.
    function hashInvalid(State state) internal pure returns (bytes32) {
        // The final hash to sign is keccak(stateInvalidSalt, keccak(state))
        return state.unwrap().keccakSalted(STATE_INVALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(State state) internal pure returns (MemView) {
        return MemView.wrap(State.unwrap(state));
    }

    /// @notice Compares two State structures.
    function equals(State a, State b) internal pure returns (bool) {
        // Length of a State payload is fixed, so we just need to compare the hashes
        return a.unwrap().keccak() == b.unwrap().keccak();
    }

    // ═══════════════════════════════════════════════ STATE HASHING ═══════════════════════════════════════════════════

    /// @notice Returns the hash of the State.
    /// @dev We are using the Merkle Root of a tree with two leafs (see below) as state hash.
    function leaf(State state) internal pure returns (bytes32) {
        (bytes32 leftLeaf_, bytes32 rightLeaf_) = state.subLeafs();
        // Final hash is the parent of these leafs
        return keccak256(bytes.concat(leftLeaf_, rightLeaf_));
    }

    /// @notice Returns "sub-leafs" of the State. Hash of these "sub leafs" is going to be used
    /// as a "state leaf" in the "Snapshot Merkle Tree".
    /// This enables proving that leftLeaf = (root, origin) was a part of the "Snapshot Merkle Tree",
    /// by combining `rightLeaf` with the remainder of the "Snapshot Merkle Proof".
    function subLeafs(State state) internal pure returns (bytes32 leftLeaf_, bytes32 rightLeaf_) {
        MemView memView = state.unwrap();
        // Left leaf is (root, origin)
        leftLeaf_ = memView.prefix({len_: OFFSET_NONCE}).keccak();
        // Right leaf is (metadata), or (nonce, blockNumber, timestamp)
        rightLeaf_ = memView.sliceFrom({index_: OFFSET_NONCE}).keccak();
    }

    /// @notice Returns the left "sub-leaf" of the State.
    function leftLeaf(bytes32 root_, uint32 origin_) internal pure returns (bytes32) {
        // We use encodePacked here to simulate the State memory layout
        return keccak256(abi.encodePacked(root_, origin_));
    }

    /// @notice Returns the right "sub-leaf" of the State.
    function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, GasData gasData_)
        internal
        pure
        returns (bytes32)
    {
        // We use encodePacked here to simulate the State memory layout
        return keccak256(abi.encodePacked(nonce_, blockNumber_, timestamp_, gasData_));
    }

    // ═══════════════════════════════════════════════ STATE SLICING ═══════════════════════════════════════════════════

    /// @notice Returns a historical Merkle root from the Origin contract.
    function root(State state) internal pure returns (bytes32) {
        return state.unwrap().index({index_: OFFSET_ROOT, bytes_: 32});
    }

    /// @notice Returns domain of chain where the Origin contract is deployed.
    function origin(State state) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(state.unwrap().indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
    function nonce(State state) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(state.unwrap().indexUint({index_: OFFSET_NONCE, bytes_: 4}));
    }

    /// @notice Returns a block number when `root` was saved in Origin.
    function blockNumber(State state) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(state.unwrap().indexUint({index_: OFFSET_BLOCK_NUMBER, bytes_: 5}));
    }

    /// @notice Returns a block timestamp when `root` was saved in Origin.
    /// @dev This is the timestamp according to the origin chain.
    function timestamp(State state) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(state.unwrap().indexUint({index_: OFFSET_TIMESTAMP, bytes_: 5}));
    }

    /// @notice Returns gas data for the chain.
    function gasData(State state) internal pure returns (GasData) {
        return GasDataLib.wrapGasData(state.unwrap().indexUint({index_: OFFSET_GAS_DATA, bytes_: GAS_DATA_LENGTH}));
    }
}
