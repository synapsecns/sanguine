// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ByteString, TypedMemView } from "./ByteString.sol";
import { STATE_LENGTH } from "./Structures.sol";

/// @dev State is a memory view over a formatted state payload.
type State is bytes29;
/// @dev Attach library functions to State
using {
    StateLib.unwrap,
    StateLib.leaf,
    StateLib.root,
    StateLib.origin,
    StateLib.nonce,
    StateLib.blockNumber,
    StateLib.timestamp
} for State global;

library StateLib {
    using ByteString for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev State structure represents the state of Origin contract at some point of time.
     * State is structured in a way to track the updates of the Origin Merkle Tree. State includes
     * root of the Origin Merkle Tree, origin domain and some additional metadata.
     *
     * Hash of every dispatched message is inserted in the Origin Merkle Tree, which changes the
     * value of Origin Merkle Root (which is the root for the mentioned tree).
     * Origin has a single Merkle Tree for all messages, regardless of their destination domain.
     * This leads to Origin state being updated if and only if a message was dispatched in a block.
     *
     * Origin contract is a "source of truth" for states: a state is considered "valid" in its Origin,
     * if it matches the state of the Origin contract after the N-th (nonce) message was dispatched.
     *
     * @dev Memory layout of State fields
     * [000 .. 032): root           bytes32 32 bytes    Root of the Origin Merkle Tree
     * [032 .. 036): origin         uint32   4 bytes    Domain where Origin is located
     * [036 .. 040): nonce          uint32   4 bytes    Amount of dispatched messages
     * [040 .. 045): blockNumber    uint40   5 bytes    Block of last dispatched message
     * [045 .. 050): timestamp      uint40   5 bytes    Time of last dispatched message
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
     * @param _root         New merkle root
     * @param _origin       Domain of Origin's chain
     * @param _nonce        Nonce of the merkle root
     * @param _blockNumber  Block number when root was saved in Origin
     * @param _timestamp    Block timestamp when root was saved in Origin
     * @return Formatted state
     **/
    function formatState(
        bytes32 _root,
        uint32 _origin,
        uint32 _nonce,
        uint40 _blockNumber,
        uint40 _timestamp
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_root, _origin, _nonce, _blockNumber, _timestamp);
    }

    /**
     * @notice Returns a State view over the given payload.
     * @dev Will revert if the payload is not a state.
     */
    function castToState(bytes memory _payload) internal pure returns (State) {
        return castToState(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a State view.
     * @dev Will revert if the memory view is not over a state.
     */
    function castToState(bytes29 _view) internal pure returns (State) {
        require(isState(_view), "Not a state");
        return State.wrap(_view);
    }

    /// @notice Checks that a payload is a formatted State.
    function isState(bytes29 _view) internal pure returns (bool) {
        return _view.len() == STATE_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(State _state) internal pure returns (bytes29) {
        return State.unwrap(_state);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            STATE HASHING                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a "state leaf": a unique hash for every unique state, that is
    /// going to be used as a leaf in the "Snapshot Merkle Tree".
    /// @dev We use the hashing technique similar to a Merkle tree here and return hash for a node
    /// having two children: (root, origin) and (nonce, blockNumber, timestamp).
    function leaf(State _state) internal pure returns (bytes32) {
        bytes29 _view = unwrap(_state);
        // Derive hash for (root, origin) bytestring
        bytes32 rootOriginHash = _view.prefix({ _len: OFFSET_NONCE, newType: 0 }).keccak();
        // Derive hash for (nonce, blockNumber, timestamp) bytestring
        bytes32 metadataHash = _view.sliceFrom({ _index: OFFSET_NONCE, newType: 0 }).keccak();
        // Final hash is two hashes concatenated, and then hashed
        return keccak256(bytes.concat(rootOriginHash, metadataHash));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            STATE SLICING                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a historical Merkle root from the Origin contract.
    function root(State _state) internal pure returns (bytes32) {
        bytes29 _view = unwrap(_state);
        return _view.index({ _index: OFFSET_ROOT, _bytes: 32 });
    }

    /// @notice Returns domain of chain where the Origin contract is deployed.
    function origin(State _state) internal pure returns (uint32) {
        bytes29 _view = unwrap(_state);
        return uint32(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 4 }));
    }

    /// @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
    function nonce(State _state) internal pure returns (uint32) {
        bytes29 _view = unwrap(_state);
        return uint32(_view.indexUint({ _index: OFFSET_NONCE, _bytes: 4 }));
    }

    /// @notice Returns a block number when `root` was saved in Origin.
    function blockNumber(State _state) internal pure returns (uint40) {
        bytes29 _view = unwrap(_state);
        return uint40(_view.indexUint({ _index: OFFSET_BLOCK_NUMBER, _bytes: 5 }));
    }

    /// @notice Returns a block timestamp when `root` was saved in Origin.
    /// @dev This is the timestamp according to the origin chain.
    function timestamp(State _state) internal pure returns (uint40) {
        bytes29 _view = unwrap(_state);
        return uint40(_view.indexUint({ _index: OFFSET_TIMESTAMP, _bytes: 5 }));
    }
}
