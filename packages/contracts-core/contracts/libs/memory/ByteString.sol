// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemView, MemViewLib} from "./MemView.sol";
import {UnformattedCallData, UnformattedCallDataPrefix, UnformattedSignature} from "../Errors.sol";

/// @dev CallData is a memory view over the payload to be used for an external call, i.e.
/// recipient.call(callData). Its length is always (4 + 32 * N) bytes:
/// - First 4 bytes represent the function selector.
/// - 32 * N bytes represent N words that function arguments occupy.
type CallData is uint256;

/// @dev Attach library functions to CallData
using ByteString for CallData global;

/// @dev Signature is a memory view over a "65 bytes" array representing a ECDSA signature.
type Signature is uint256;

/// @dev Attach library functions to Signature
using ByteString for Signature global;

library ByteString {
    using MemViewLib for bytes;

    /**
     * @dev non-compact ECDSA signatures are enforced as of OZ 4.7.3
     *
     *      Signature payload memory layout
     * [000 .. 032) r   bytes32 32 bytes
     * [032 .. 064) s   bytes32 32 bytes
     * [064 .. 065) v   uint8    1 byte
     */
    uint256 internal constant SIGNATURE_LENGTH = 65;
    uint256 private constant OFFSET_R = 0;
    uint256 private constant OFFSET_S = 32;
    uint256 private constant OFFSET_V = 64;

    /**
     * @dev Calldata memory layout
     * [000 .. 004) selector    bytes4  4 bytes
     *      Optional: N function arguments
     * [004 .. 036) arg1        bytes32 32 bytes
     *      ..
     * [AAA .. END) argN        bytes32 32 bytes
     */
    uint256 internal constant SELECTOR_LENGTH = 4;
    uint256 private constant OFFSET_SELECTOR = 0;
    uint256 private constant OFFSET_ARGUMENTS = SELECTOR_LENGTH;

    // ═════════════════════════════════════════════════ SIGNATURE ═════════════════════════════════════════════════════

    /**
     * @notice Constructs the signature payload from the given values.
     * @dev Using ByteString.formatSignature({r: r, s: s, v: v}) will make sure
     * that params are given in the right order.
     */
    function formatSignature(bytes32 r, bytes32 s, uint8 v) internal pure returns (bytes memory) {
        return abi.encodePacked(r, s, v);
    }

    /**
     * @notice Returns a Signature view over for the given payload.
     * @dev Will revert if the payload is not a signature.
     */
    function castToSignature(bytes memory payload) internal pure returns (Signature) {
        return castToSignature(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Signature view.
     * @dev Will revert if the memory view is not over a signature.
     */
    function castToSignature(MemView memView) internal pure returns (Signature) {
        if (!isSignature(memView)) revert UnformattedSignature();
        return Signature.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a byte string is a signature
     */
    function isSignature(MemView memView) internal pure returns (bool) {
        return memView.len() == SIGNATURE_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Signature signature) internal pure returns (MemView) {
        return MemView.wrap(Signature.unwrap(signature));
    }

    // ═════════════════════════════════════════════ SIGNATURE SLICING ═════════════════════════════════════════════════

    /// @notice Unpacks signature payload into (r, s, v) parameters.
    /// @dev Make sure to verify signature length with isSignature() beforehand.
    function toRSV(Signature signature) internal pure returns (bytes32 r, bytes32 s, uint8 v) {
        // Get the underlying memory view
        MemView memView = unwrap(signature);
        r = memView.index({index_: OFFSET_R, bytes_: 32});
        s = memView.index({index_: OFFSET_S, bytes_: 32});
        // Can be safely casted to uint8, since we index a single byte
        v = uint8(memView.indexUint({index_: OFFSET_V, bytes_: 1}));
    }

    // ═════════════════════════════════════════════════ CALLDATA ══════════════════════════════════════════════════════

    /**
     * @notice Constructs the calldata with the modified arguments:
     * the existing arguments are prepended with the arguments from the prefix.
     * @dev Given:
     *  - `calldata = abi.encodeWithSelector(foo.selector, d, e);`
     *  - `prefix = abi.encode(a, b, c);`
     *  - `a`, `b`, `c` are arguments of static type (i.e. not dynamically sized ones)
     *      Then:
     *  - Function will return abi.encodeWithSelector(foo.selector, a, c, c, d, e)
     *  - Returned calldata will trigger `foo(a, b, c, d, e)` when used for a contract call.
     * Note: for clarification as to what types are considered static, see
     * https://docs.soliditylang.org/en/latest/abi-spec.html#formal-specification-of-the-encoding
     * @param callData  Calldata that needs to be modified
     * @param prefix    ABI-encoded arguments to use as the first arguments in the new calldata
     * @return Modified calldata having prefix as the first arguments.
     */
    function addPrefix(CallData callData, bytes memory prefix) internal view returns (bytes memory) {
        // Prefix should occupy a whole amount of words in memory
        if (!_fullWords(prefix.length)) revert UnformattedCallDataPrefix();
        MemView[] memory views = new MemView[](3);
        // Use payload's function selector
        views[0] = abi.encodePacked(callData.callSelector()).ref();
        // Use prefix as the first arguments
        views[1] = prefix.ref();
        // Use payload's remaining arguments
        views[2] = callData.arguments();
        return MemViewLib.join(views);
    }

    /**
     * @notice Returns a CallData view over for the given payload.
     * @dev Will revert if the memory view is not over a calldata.
     */
    function castToCallData(bytes memory payload) internal pure returns (CallData) {
        return castToCallData(payload.ref());
    }

    /**
     * @notice Casts a memory view to a CallData view.
     * @dev Will revert if the memory view is not over a calldata.
     */
    function castToCallData(MemView memView) internal pure returns (CallData) {
        if (!isCallData(memView)) revert UnformattedCallData();
        return CallData.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a byte string is a valid calldata, i.e.
     * a function selector, followed by arbitrary amount of arguments.
     */
    function isCallData(MemView memView) internal pure returns (bool) {
        uint256 length = memView.len();
        // Calldata should at least have a function selector
        if (length < SELECTOR_LENGTH) return false;
        // The remainder of the calldata should be exactly N memory words (N >= 0)
        return _fullWords(length - SELECTOR_LENGTH);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(CallData callData) internal pure returns (MemView) {
        return MemView.wrap(CallData.unwrap(callData));
    }

    /// @notice Returns callData's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(CallData callData) internal pure returns (bytes32) {
        return callData.unwrap().keccak();
    }

    // ═════════════════════════════════════════════ CALLDATA SLICING ══════════════════════════════════════════════════

    /**
     * @notice Returns amount of memory words (32 byte chunks) the function arguments
     * occupy in the calldata.
     * @dev This might differ from amount of arguments supplied, if any of the arguments
     * occupies more than one memory slot. It is true, however, that argument part of the payload
     * occupies exactly N words, even for dynamic types like `bytes`
     */
    function argumentWords(CallData callData) internal pure returns (uint256) {
        // Get the underlying memory view
        MemView memView = unwrap(callData);
        // Equivalent of (length - SELECTOR_LENGTH) / 32
        return (memView.len() - SELECTOR_LENGTH) >> 5;
    }

    /// @notice Returns selector for the provided calldata.
    function callSelector(CallData callData) internal pure returns (bytes4) {
        // Get the underlying memory view
        MemView memView = unwrap(callData);
        return bytes4(memView.index({index_: OFFSET_SELECTOR, bytes_: SELECTOR_LENGTH}));
    }

    /// @notice Returns abi encoded arguments for the provided calldata.
    function arguments(CallData callData) internal pure returns (MemView) {
        // Get the underlying memory view
        MemView memView = unwrap(callData);
        return memView.sliceFrom({index_: OFFSET_ARGUMENTS});
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Checks if length is full amount of memory words (32 bytes).
    function _fullWords(uint256 length) internal pure returns (bool) {
        // The equivalent of length % 32 == 0
        return length & 31 == 0;
    }
}
