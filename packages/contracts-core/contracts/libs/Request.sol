// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {REQUEST_LENGTH} from "./Constants.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev Request is a memory view over a formatted "message execution request" payload.
type Request is bytes29;

/// @dev Attach library functions to Snapshot
using RequestLib for Request global;

library RequestLib {
    using ByteString for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev Request structure represents a message sender requirements for
     * the message execution on the destination chain.
     *
     * @dev Memory layout of Request fields
     * TODO: figure out the fields packing (uint64 is too much for gas limit)
     * [000 .. 008): gasLimit       uint64   8 bytes    Amount of gas units to supply on destination chain
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_GAS_LIMIT = 0;

    // ══════════════════════════════════════════════════ REQUEST ══════════════════════════════════════════════════════

    function formatRequest(uint64 gasLimit_) internal pure returns (bytes memory) {
        return abi.encodePacked(gasLimit_);
    }

    /**
     * @notice Returns a Request view over the given payload.
     * @dev Will revert if the payload is not a request.
     */
    function castToRequest(bytes memory payload) internal pure returns (Request) {
        return castToRequest(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Request view.
     * @dev Will revert if the memory view is not over a request.
     */
    function castToRequest(bytes29 view_) internal pure returns (Request) {
        require(isRequest(view_), "Not a request");
        return Request.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted Request.
    function isRequest(bytes29 view_) internal pure returns (bool) {
        return view_.len() == REQUEST_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Request request) internal pure returns (bytes29) {
        return Request.unwrap(request);
    }

    // ══════════════════════════════════════════════ REQUEST SLICING ══════════════════════════════════════════════════

    function gasLimit(Request request) internal pure returns (uint64) {
        bytes29 view_ = unwrap(request);
        return uint64(view_.indexUint({index_: OFFSET_GAS_LIMIT, bytes_: 8}));
    }
}
