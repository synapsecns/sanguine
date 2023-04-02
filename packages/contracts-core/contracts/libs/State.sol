// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {STATE_LENGTH} from "./Constants.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev State is a memory view over a formatted state payload.
type State is bytes29;
/// @dev Attach library functions to State

using {
    StateLib.unwrap,
    StateLib.equals,
    StateLib.equalToOrigin,
    StateLib.leaf,
    StateLib.subLeafs,
    StateLib.toSummitState,
    StateLib.root,
    StateLib.origin,
    StateLib.nonce,
    StateLib.blockNumber,
    StateLib.timestamp
} for State global;

/// @dev Struct representing State, as it is stored in the Origin contract.
struct OriginState {
    uint40 blockNumber;
    uint40 timestamp;
}
// 176 bits left for tight packing
/// @dev Attach library functions to OriginState

using {StateLib.formatOriginState} for OriginState global;

/// @dev Struct representing State, as it is stored in the Summit contract.
struct SummitState {
    bytes32 root;
    uint32 origin;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
}
// 112 bits left for tight packing
/// @dev Attach library functions to SummitState

using {StateLib.formatSummitState} for SummitState global;

library StateLib {
    using ByteString for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev State structure represents the state of Origin contract at some point of time.
     * State is structured in a way to track the updates of the Origin Merkle Tree. State includes
     * root of the Origin Merkle Tree, origin domain and some additional metadata.
     *
     * Hash of every sent message is inserted in the Origin Merkle Tree, which changes the
     * value of Origin Merkle Root (which is the root for the mentioned tree).
     * Origin has a single Merkle Tree for all messages, regardless of their destination domain.
     * This leads to Origin state being updated if and only if a message was sent in a block.
     *
     * Origin contract is a "source of truth" for states: a state is considered "valid" in its Origin,
     * if it matches the state of the Origin contract after the N-th (nonce) message was sent.
     *
     * @dev Memory layout of State fields
     * [000 .. 032): root           bytes32 32 bytes    Root of the Origin Merkle Tree
     * [032 .. 036): origin         uint32   4 bytes    Domain where Origin is located
     * [036 .. 040): nonce          uint32   4 bytes    Amount of sent messages
     * [040 .. 045): blockNumber    uint40   5 bytes    Block of last sent message
     * [045 .. 050): timestamp      uint40   5 bytes    Time of last sent message
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_ROOT = 0;
    uint256 private constant OFFSET_ORIGIN = 32;
    uint256 private constant OFFSET_NONCE = 36;
    uint256 private constant OFFSET_BLOCK_NUMBER = 40;
    uint256 private constant OFFSET_TIMESTAMP = 45;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                STATE                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted State payload with provided fields
     * @param root_         New merkle root
     * @param origin_       Domain of Origin's chain
     * @param nonce_        Nonce of the merkle root
     * @param blockNumber_  Block number when root was saved in Origin
     * @param timestamp_    Block timestamp when root was saved in Origin
     * @return Formatted state
     */
    function formatState(bytes32 root_, uint32 origin_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(root_, origin_, nonce_, blockNumber_, timestamp_);
    }

    /**
     * @notice Returns a State view over the given payload.
     * @dev Will revert if the payload is not a state.
     */
    function castToState(bytes memory payload) internal pure returns (State) {
        return castToState(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a State view.
     * @dev Will revert if the memory view is not over a state.
     */
    function castToState(bytes29 view_) internal pure returns (State) {
        require(isState(view_), "Not a state");
        return State.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted State.
    function isState(bytes29 view_) internal pure returns (bool) {
        return view_.len() == STATE_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(State state) internal pure returns (bytes29) {
        return State.unwrap(state);
    }

    /// @notice Compares two State structures.
    function equals(State a, State b) internal pure returns (bool) {
        // Length of a State payload is fixed, so we just need to compare the hashes
        return a.unwrap().keccak() == b.unwrap().keccak();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ORIGIN STATE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted State payload with provided fields.
     * @param origin_       Domain of Origin's chain
     * @param nonce_        Nonce of the merkle root
     * @param originState_  State struct as it is stored in Origin contract
     * @return Formatted state
     */
    function formatOriginState(OriginState memory originState_, bytes32 root_, uint32 origin_, uint32 nonce_)
        internal
        pure
        returns (bytes memory)
    {
        return formatState({
            root_: root_,
            origin_: origin_,
            nonce_: nonce_,
            blockNumber_: originState_.blockNumber,
            timestamp_: originState_.timestamp
        });
    }

    /// @notice Returns a struct to save in the Origin contract.
    /// Current block number and timestamp are used.
    // solhint-disable-next-line ordering
    function originState() internal view returns (OriginState memory state) {
        state.blockNumber = uint40(block.number);
        state.timestamp = uint40(block.timestamp);
    }

    /// @notice Checks that a state and its Origin representation are equal.
    function equalToOrigin(State state, OriginState memory originState_) internal pure returns (bool) {
        return state.blockNumber() == originState_.blockNumber && state.timestamp() == originState_.timestamp;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             SUMMIT STATE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted State payload with provided fields.
     * @param summitState   State struct as it is stored in Summit contract
     * @return Formatted state
     */
    function formatSummitState(SummitState memory summitState) internal pure returns (bytes memory) {
        return formatState({
            root_: summitState.root,
            origin_: summitState.origin,
            nonce_: summitState.nonce,
            blockNumber_: summitState.blockNumber,
            timestamp_: summitState.timestamp
        });
    }

    /// @notice Returns a struct to save in the Summit contract.
    function toSummitState(State state) internal pure returns (SummitState memory state_) {
        state_.root = state.root();
        state_.origin = state.origin();
        state_.nonce = state.nonce();
        state_.blockNumber = state.blockNumber();
        state_.timestamp = state.timestamp();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            STATE HASHING                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
        bytes29 view_ = state.unwrap();
        // Left leaf is (root, origin)
        leftLeaf_ = view_.prefix({len_: OFFSET_NONCE, newType: 0}).keccak();
        // Right leaf is (metadata), or (nonce, blockNumber, timestamp)
        rightLeaf_ = view_.sliceFrom({index_: OFFSET_NONCE, newType: 0}).keccak();
    }

    /// @notice Returns the left "sub-leaf" of the State.
    function leftLeaf(bytes32 root_, uint32 origin_) internal pure returns (bytes32) {
        // We use encodePacked here to simulate the State memory layout
        return keccak256(abi.encodePacked(root_, origin_));
    }

    /// @notice Returns the right "sub-leaf" of the State.
    function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) internal pure returns (bytes32) {
        // We use encodePacked here to simulate the State memory layout
        return keccak256(abi.encodePacked(nonce_, blockNumber_, timestamp_));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            STATE SLICING                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a historical Merkle root from the Origin contract.
    function root(State state) internal pure returns (bytes32) {
        bytes29 view_ = state.unwrap();
        return view_.index({index_: OFFSET_ROOT, bytes_: 32});
    }

    /// @notice Returns domain of chain where the Origin contract is deployed.
    function origin(State state) internal pure returns (uint32) {
        bytes29 view_ = state.unwrap();
        return uint32(view_.indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
    function nonce(State state) internal pure returns (uint32) {
        bytes29 view_ = state.unwrap();
        return uint32(view_.indexUint({index_: OFFSET_NONCE, bytes_: 4}));
    }

    /// @notice Returns a block number when `root` was saved in Origin.
    function blockNumber(State state) internal pure returns (uint40) {
        bytes29 view_ = state.unwrap();
        return uint40(view_.indexUint({index_: OFFSET_BLOCK_NUMBER, bytes_: 5}));
    }

    /// @notice Returns a block timestamp when `root` was saved in Origin.
    /// @dev This is the timestamp according to the origin chain.
    function timestamp(State state) internal pure returns (uint40) {
        bytes29 view_ = state.unwrap();
        return uint40(view_.indexUint({index_: OFFSET_TIMESTAMP, bytes_: 5}));
    }
}
