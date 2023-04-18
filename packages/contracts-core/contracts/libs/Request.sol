// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {REQUEST_LENGTH} from "./Constants.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// @dev Request is a memory view over a formatted "message execution request" payload.
type Request is uint256;

/// @dev Attach library functions to Request
using RequestLib for Request global;

library RequestLib {
    using MemViewLib for bytes;

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
        return castToRequest(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Request view.
     * @dev Will revert if the memory view is not over a request.
     */
    function castToRequest(MemView memView) internal pure returns (Request) {
        require(isRequest(memView), "Not a request");
        return Request.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Request.
    function isRequest(MemView memView) internal pure returns (bool) {
        return memView.len() == REQUEST_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Request request) internal pure returns (MemView) {
        return MemView.wrap(Request.unwrap(request));
    }

    // ══════════════════════════════════════════════ REQUEST SLICING ══════════════════════════════════════════════════

    function gasLimit(Request request) internal pure returns (uint64) {
        MemView memView = unwrap(request);
        return uint64(memView.indexUint({index_: OFFSET_GAS_LIMIT, bytes_: 8}));
    }
}
