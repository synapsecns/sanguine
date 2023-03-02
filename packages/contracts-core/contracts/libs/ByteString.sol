// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./TypedMemView.sol";

/// @dev CallData is a memory view over the payload to be used for an external call, i.e.
/// recipient.call(callData). Its length is always (4 + 32 * N) bytes:
/// - First 4 bytes represent the function selector.
/// - 32 * N bytes represent N words that function arguments occupy.
type CallData is bytes29;
/// @dev Signature is a memory view over a "65 bytes" array representing a ECDSA signature.
type Signature is bytes29;

library ByteString {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev non-compact ECDSA signatures are enforced as of OZ 4.7.3
     *
     *      Signature payload memory layout
     * [000 .. 032) r   bytes32 32 bytes
     * [032 .. 064) s   bytes32 32 bytes
     * [064 .. 065) v   uint8    1 byte
     */
    uint256 internal constant SIGNATURE_LENGTH = 65;
    uint256 internal constant OFFSET_R = 0;
    uint256 internal constant OFFSET_S = 32;
    uint256 internal constant OFFSET_V = 64;

    /**
     * @dev Calldata memory layout
     * [000 .. 004) selector    bytes4  4 bytes
     *      Optional: N function arguments
     * [004 .. 036) arg1        bytes32 32 bytes
     *      ..
     * [AAA .. END) argN        bytes32 32 bytes
     */
    uint256 internal constant SELECTOR_LENGTH = 4;
    uint256 internal constant OFFSET_SELECTOR = 0;
    uint256 internal constant OFFSET_ARGUMENTS = SELECTOR_LENGTH;

    /**
     * @notice Returns a memory view over the given payload, treating it as raw bytes.
     * @dev Shortcut for .ref(0) - to be deprecated once "uint40 type" is removed from bytes29.
     */
    function castToRawBytes(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref({ newType: 0 });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              SIGNATURE                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Constructs the signature payload from the given values.
     * @dev Using ByteString.formatSignature({r: r, s: s, v: v}) will make sure
     * that params are given in the right order.
     */
    function formatSignature(
        bytes32 r,
        bytes32 s,
        uint8 v
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(r, s, v);
    }

    /**
     * @notice Returns a Signature view over for the given payload.
     * @dev Will revert if the payload is not a signature.
     */
    function castToSignature(bytes memory _payload) internal pure returns (Signature) {
        return castToSignature(castToRawBytes(_payload));
    }

    /**
     * @notice Casts a memory view to a Signature view.
     * @dev Will revert if the memory view is not over a signature.
     */
    function castToSignature(bytes29 _view) internal pure returns (Signature) {
        require(isSignature(_view), "Not a signature");
        return Signature.wrap(_view);
    }

    /**
     * @notice Checks that a byte string is a signature
     */
    function isSignature(bytes29 _view) internal pure returns (bool) {
        return _view.len() == SIGNATURE_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Signature _signature) internal pure returns (bytes29) {
        return Signature.unwrap(_signature);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SIGNATURE SLICING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Unpacks signature payload into (r, s, v) parameters.
    /// @dev Make sure to verify signature length with isSignature() beforehand.
    function toRSV(Signature _signature)
        internal
        pure
        returns (
            bytes32 r,
            bytes32 s,
            uint8 v
        )
    {
        // Get the underlying memory view
        bytes29 _view = unwrap(_signature);
        r = _view.index({ _index: OFFSET_R, _bytes: 32 });
        s = _view.index({ _index: OFFSET_S, _bytes: 32 });
        v = uint8(_view.indexUint({ _index: OFFSET_V, _bytes: 1 }));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               CALLDATA                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a CallData view over for the given payload.
     * @dev Will revert if the memory view is not over a calldata.
     */
    function castToCallData(bytes memory _payload) internal pure returns (CallData) {
        return castToCallData(castToRawBytes(_payload));
    }

    /**
     * @notice Casts a memory view to a CallData view.
     * @dev Will revert if the memory view is not over a calldata.
     */
    function castToCallData(bytes29 _view) internal pure returns (CallData) {
        require(isCallData(_view), "Not a calldata");
        return CallData.wrap(_view);
    }

    /**
     * @notice Checks that a byte string is a valid calldata, i.e.
     * a function selector, followed by arbitrary amount of arguments.
     */
    function isCallData(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Calldata should at least have a function selector
        if (length < SELECTOR_LENGTH) return false;
        // The remainder of the calldata should be exactly N words (N >= 0), i.e.
        // (length - SELECTOR_LENGTH) % 32 == 0
        // We're using logical AND here to speed it up a bit
        return (length - SELECTOR_LENGTH) & 31 == 0;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(CallData _callData) internal pure returns (bytes29) {
        return CallData.unwrap(_callData);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CALLDATA SLICING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns amount of memory words (32 byte chunks) the function arguments
     * occupy in the calldata.
     * @dev This might differ from amount of arguments supplied, if any of the arguments
     * occupies more than one memory slot. It is true, however, that argument part of the payload
     * occupies exactly N words, even for dynamic types like `bytes`
     */
    function argumentWords(CallData _callData) internal pure returns (uint256) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_callData);
        // Equivalent of (length - SELECTOR_LENGTH) / 32
        return (_view.len() - SELECTOR_LENGTH) >> 5;
    }

    /// @notice Returns selector for the provided calldata.
    function callSelector(CallData _callData) internal pure returns (bytes29) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_callData);
        return _view.slice({ _index: OFFSET_SELECTOR, _len: SELECTOR_LENGTH, newType: 0 });
    }

    /// @notice Returns abi encoded arguments for the provided calldata.
    function arguments(CallData _callData) internal pure returns (bytes29) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_callData);
        return _view.sliceFrom({ _index: OFFSET_ARGUMENTS, newType: 0 });
    }
}
