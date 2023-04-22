// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ByteString, CallData } from "./ByteString.sol";
import { TypedMemView } from "./TypedMemView.sol";

/// @dev SystemMessage is a memory view over the message with instructions for a system call.
type SystemMessage is bytes29;

library SystemMessageLib {
    using ByteString for bytes;
    using ByteString for bytes29;
    using ByteString for CallData;
    using TypedMemView for bytes29;

    /**
     * @dev SystemMessage memory layout
     * [000 .. 001): recipient      uint8   1 bytes
     * [001 .. END]: calldata       bytes   ? bytes
     */

    uint256 internal constant OFFSET_RECIPIENT = 0;
    uint256 internal constant OFFSET_CALLDATA = 1;

    /**
     * @dev System Router is supposed to modify (rootSubmittedAt, origin, caller)
     * in the given calldata, meaning for a valid system calldata
     * there has to exist at least three arguments, occupying at least three words in total.
     */
    uint256 internal constant CALLDATA_MIN_ARGUMENT_WORDS = 3;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SYSTEM MESSAGE                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted SystemMessage payload with provided fields.
     * See: formatAdjustedCallData() for more details.
     * @param _systemRecipient  System Contract to receive message (see SystemEntity)
     * @param _callData         Calldata where the first arguments need to be replaced
     * @param _prefix           ABI-encoded arguments to use as the first arguments in the calldata
     * @return Formatted SystemMessage payload.
     */
    function formatSystemMessage(
        uint8 _systemRecipient,
        CallData _callData,
        bytes29 _prefix
    ) internal view returns (bytes memory) {
        bytes29 arguments = _callData.arguments();
        // Arguments payload should be at least as long as the replacement prefix
        require(arguments.len() >= _prefix.len(), "Payload too short");
        bytes29[] memory views = new bytes29[](4);
        // First byte is encoded system recipient
        views[0] = abi.encodePacked(_systemRecipient).castToRawBytes();
        // Use payload's function selector
        views[1] = _callData.callSelector();
        // Use prefix as the first arguments
        views[2] = _prefix;
        // Use payload's remaining arguments (following prefix)
        views[3] = arguments.sliceFrom({ _index: _prefix.len(), newType: 0 });
        return TypedMemView.join(views);
    }

    /**
     * @notice Constructs the calldata having the first arguments replaced with given prefix.
     * @dev Given:
     * - `payload = abi.encodeWithSelector(foo.selector, a0, b0, c0, d0, e0);`
     * - `prefix = abi.encode(a1, b1, c1);`
     * - `a`, `b`, `c` are static type arguments
     *      Then:
     * - Existing payload will trigger `foo(a0, b0, c0, d0, e0)`
     * - Adjusted payload will trigger `foo(a1, b1, c1, d0, e0)`
     * @param _callData Calldata where the first arguments need to be replaced
     * @param _prefix   ABI-encoded arguments to use as the first arguments in the calldata
     * @return Adjusted calldata with replaced first arguments
     */
    function formatAdjustedCallData(CallData _callData, bytes29 _prefix)
        internal
        view
        returns (bytes memory)
    {
        bytes29 arguments = _callData.arguments();
        // Arguments payload should be at least as long as the replacement prefix
        require(arguments.len() >= _prefix.len(), "Payload too short");
        bytes29[] memory views = new bytes29[](3);
        // Use payload's function selector
        views[0] = _callData.callSelector();
        // Use prefix as the first arguments
        views[1] = _prefix;
        // Use payload's remaining arguments (following prefix)
        views[2] = arguments.sliceFrom({ _index: _prefix.len(), newType: 0 });
        return TypedMemView.join(views);
    }

    /**
     * @notice Returns a SystemMessage view over for the given payload.
     * @dev Will revert if the payload is not a system message.
     */
    function castToSystemMessage(bytes memory _payload) internal pure returns (SystemMessage) {
        return castToSystemMessage(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a SystemMessage view.
     * @dev Will revert if the payload is not a system message.
     */
    function castToSystemMessage(bytes29 _view) internal pure returns (SystemMessage) {
        require(isSystemMessage(_view), "Not a system message");
        return SystemMessage.wrap(_view);
    }

    /**
     * @notice Checks that a payload is a formatted System Message.
     */
    function isSystemMessage(bytes29 _view) internal pure returns (bool) {
        // Payload needs to exist (system calls are never done via fallback function)
        if (_view.len() < OFFSET_CALLDATA) return false;
        bytes29 _callData = _getCallData(_view);
        // Payload needs to be a proper calldata
        if (!_callData.isCallData()) return false;
        // Payload needs to have at least this amount of argument words
        return _callData.castToCallData().argumentWords() >= CALLDATA_MIN_ARGUMENT_WORDS;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(SystemMessage _sm) internal pure returns (bytes29) {
        return SystemMessage.unwrap(_sm);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        SYSTEM MESSAGE SLICING                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns int value of System Message recipient (see SystemEntity).
     */
    function callRecipient(SystemMessage _systemMessage) internal pure returns (uint8) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_systemMessage);
        return uint8(_view.indexUint({ _index: OFFSET_RECIPIENT, _bytes: 1 }));
    }

    /**
     * @notice Returns System Message calldata.
     */
    function callData(SystemMessage _systemMessage) internal pure returns (CallData) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_systemMessage);
        return _getCallData(_view).castToCallData();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a generic memory view over System Message calldata,
     * without verifying that this is a valid calldata.
     */
    function _getCallData(bytes29 _view) private pure returns (bytes29) {
        return _view.sliceFrom({ _index: OFFSET_CALLDATA, newType: 0 });
    }
}
