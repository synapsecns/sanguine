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
     * @param systemRecipient   System Contract to receive message (see SystemEntity)
     * @param callData_         Calldata where the first arguments need to be replaced
     * @param prefix            ABI-encoded arguments to use as the first arguments in the calldata
     * @return Formatted SystemMessage payload.
     */
    function formatSystemMessage(
        uint8 systemRecipient,
        CallData callData_,
        bytes29 prefix
    ) internal view returns (bytes memory) {
        bytes29 arguments = callData_.arguments();
        // Arguments payload should be at least as long as the replacement prefix
        require(arguments.len() >= prefix.len(), "Payload too short");
        bytes29[] memory views = new bytes29[](4);
        // First byte is encoded system recipient
        views[0] = abi.encodePacked(systemRecipient).castToRawBytes();
        // Use payload's function selector
        views[1] = callData_.callSelector();
        // Use prefix as the first arguments
        views[2] = prefix;
        // Use payload's remaining arguments (following prefix)
        views[3] = arguments.sliceFrom({ index_: prefix.len(), newType: 0 });
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
     * @param callData_ Calldata where the first arguments need to be replaced
     * @param prefix    ABI-encoded arguments to use as the first arguments in the calldata
     * @return Adjusted calldata with replaced first arguments
     */
    function formatAdjustedCallData(CallData callData_, bytes29 prefix)
        internal
        view
        returns (bytes memory)
    {
        bytes29 arguments = callData_.arguments();
        // Arguments payload should be at least as long as the replacement prefix
        require(arguments.len() >= prefix.len(), "Payload too short");
        bytes29[] memory views = new bytes29[](3);
        // Use payload's function selector
        views[0] = callData_.callSelector();
        // Use prefix as the first arguments
        views[1] = prefix;
        // Use payload's remaining arguments (following prefix)
        views[2] = arguments.sliceFrom({ index_: prefix.len(), newType: 0 });
        return TypedMemView.join(views);
    }

    /**
     * @notice Returns a SystemMessage view over for the given payload.
     * @dev Will revert if the payload is not a system message.
     */
    function castToSystemMessage(bytes memory payload) internal pure returns (SystemMessage) {
        return castToSystemMessage(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a SystemMessage view.
     * @dev Will revert if the payload is not a system message.
     */
    function castToSystemMessage(bytes29 view_) internal pure returns (SystemMessage) {
        require(isSystemMessage(view_), "Not a system message");
        return SystemMessage.wrap(view_);
    }

    /**
     * @notice Checks that a payload is a formatted System Message.
     */
    function isSystemMessage(bytes29 view_) internal pure returns (bool) {
        // Payload needs to exist (system calls are never done via fallback function)
        if (view_.len() < OFFSET_CALLDATA) return false;
        bytes29 callDataView = _getCallData(view_);
        // Payload needs to be a proper calldata
        if (!callDataView.isCallData()) return false;
        // Payload needs to have at least this amount of argument words
        return callDataView.castToCallData().argumentWords() >= CALLDATA_MIN_ARGUMENT_WORDS;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(SystemMessage systemMessage) internal pure returns (bytes29) {
        return SystemMessage.unwrap(systemMessage);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        SYSTEM MESSAGE SLICING                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns int value of System Message recipient (see SystemEntity).
     */
    function callRecipient(SystemMessage systemMessage) internal pure returns (uint8) {
        // Get the underlying memory view
        bytes29 view_ = unwrap(systemMessage);
        return uint8(view_.indexUint({ index_: OFFSET_RECIPIENT, bytes_: 1 }));
    }

    /**
     * @notice Returns System Message calldata.
     */
    function callData(SystemMessage systemMessage) internal pure returns (CallData) {
        // Get the underlying memory view
        bytes29 view_ = unwrap(systemMessage);
        return _getCallData(view_).castToCallData();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a generic memory view over System Message calldata,
     * without verifying that this is a valid calldata.
     */
    function _getCallData(bytes29 view_) private pure returns (bytes29) {
        return view_.sliceFrom({ index_: OFFSET_CALLDATA, newType: 0 });
    }
}
