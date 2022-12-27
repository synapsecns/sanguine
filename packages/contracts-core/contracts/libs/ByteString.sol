// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseTypes } from "./SynapseTypes.sol";
import { TypedMemView } from "./TypedMemView.sol";

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
     * @dev Call payload memory layout
     * [000 .. 004) selector    bytes4  4 bytes
     *      Optional: N function arguments
     * [004 .. 036) arg1        bytes32 32 bytes
     *      ..
     * [AAA .. END) argN        bytes32 32 bytes
     */
    uint256 internal constant SELECTOR_LENGTH = 4;
    uint256 internal constant OFFSET_SELECTOR = 0;
    uint256 internal constant OFFSET_ARGUMENTS = SELECTOR_LENGTH;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyType(bytes29 _view, uint40 _type) {
        _view.assertType(_type);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
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
     * @notice Returns a properly typed bytes29 pointer for a raw bytes payload.
     */
    function castToRawBytes(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.RAW_BYTES);
    }

    /**
     * @notice Returns a properly typed bytes29 pointer for a signature payload.
     */
    function castToSignature(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.SIGNATURE);
    }

    /**
     * @notice Checks that a byte string is a signature
     */
    function isSignature(bytes29 _view) internal pure returns (bool) {
        return _view.len() == SIGNATURE_LENGTH;
    }

    /**
     * @notice Returns a properly typed bytes29 pointer for a call payload.
     */
    function castToCallPayload(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.CALL_PAYLOAD);
    }

    /**
     * @notice Checks that a byte string is a call payload, i.e.
     * a function selector, followed by arbitrary amount of arguments.
     */
    function isCallPayload(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Call payload should at least have a function selector
        if (length < SELECTOR_LENGTH) return false;
        // The remainder of the payload should be exactly N words (N >= 0), i.e.
        // (length - SELECTOR_LENGTH) % 32 == 0
        // We're using logical AND here to speed it up a bit
        return (length - SELECTOR_LENGTH) & 31 == 0;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CALL PAYLOAD SLICING                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns amount of memory words (32 byte chunks) the function arguments
     * occupy in the call payload.
     * @dev This might differ from amount of arguments supplied, if any of the arguments
     * occupies more than one memory slot. It is true, however, that argument part of the payload
     * occupies exactly N words, even for dynamic types like `bytes`
     */
    function argumentWords(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.CALL_PAYLOAD)
        returns (uint256)
    {
        // Equivalent of (length - SELECTOR_LENGTH) / 32
        return (_view.len() - SELECTOR_LENGTH) >> 5;
    }

    /// @notice Returns selector for the provided call payload.
    function callSelector(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.CALL_PAYLOAD)
        returns (bytes29)
    {
        return
            _view.slice({
                _index: OFFSET_SELECTOR,
                _len: SELECTOR_LENGTH,
                newType: SynapseTypes.RAW_BYTES
            });
    }

    /// @notice Returns abi encoded arguments for the provided call payload.
    function argumentsPayload(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.CALL_PAYLOAD)
        returns (bytes29)
    {
        return _view.sliceFrom({ _index: OFFSET_ARGUMENTS, newType: SynapseTypes.RAW_BYTES });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SIGNATURE SLICING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Unpacks signature payload into (r, s, v) parameters.
    /// @dev Make sure to verify signature length with isSignature() beforehand.
    function toRSV(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SIGNATURE)
        returns (
            bytes32 r,
            bytes32 s,
            uint8 v
        )
    {
        r = _view.index({ _index: OFFSET_R, _bytes: 32 });
        s = _view.index({ _index: OFFSET_S, _bytes: 32 });
        v = uint8(_view.indexUint({ _index: OFFSET_V, _bytes: 1 }));
    }
}
