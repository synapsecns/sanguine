// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./ByteString.sol";
import "./Structures.sol";

/// @dev State is a memory view over a formatted state payload.
type State is bytes29;
/// @dev Attach library functions to State
using {
    StateLib.unwrap,
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
     * @dev State structure represents the state of Origin contract at some point of time,
     * aka the "historical state".
     *
     * @dev State memory layout
     * [000 .. 032): root           bytes32 32 bytes
     * [032 .. 036): origin         uint32   4 bytes
     * [036 .. 040): nonce          uint32   4 bytes
     * [040 .. 045): blockNumber    uint40   5 bytes
     * [045 .. 050): timestamp      uint40   5 bytes
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
