// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./ByteString.sol";

library SystemCall {
    using ByteString for bytes29;
    using ByteString for CallData;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev Custom address, used for sending and receiving system messages.
     *      Origin is supposed to dispatch messages from SystemRouter
     *      as if they were sent by this address.
     *      Destination is supposed to reroute messages for this address to SystemRouter.
     *
     *      Note: all bits except for lower 20 bytes are set to 1.
     *      Note: TypeCasts.bytes32ToAddress(SYSTEM_ROUTER) == address(0)
     */
    bytes32 internal constant SYSTEM_ROUTER = bytes32(type(uint256).max << 160);

    /**
     * @dev SystemCall memory layout
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
     * @notice Returns a formatted SystemCall payload with provided fields.
     * See: formatAdjustedCallData() for more details.
     * @param _systemRecipient  System Contract to receive message (see ISystemRouter.SystemEntity)
     * @param _callData         Calldata where the first arguments need to be replaced
     * @param _prefix           ABI-encoded arguments to use as the first arguments in the calldata
     * @return Formatted SystemCall payload.
     */
    function formatSystemCall(
        uint8 _systemRecipient,
        CallData _callData,
        bytes29 _prefix
    ) internal view returns (bytes memory) {
        bytes29 arguments = _callData.arguments();
        // Arguments payload should be at least as long as the replacement prefix
        require(arguments.len() >= _prefix.len(), "Payload too short");
        bytes29[] memory views = new bytes29[](4);
        // First byte is encoded system recipient
        views[0] = abi.encodePacked(_systemRecipient).ref(SynapseTypes.RAW_BYTES);
        // Use payload's function selector
        views[1] = _callData.callSelector();
        // Use prefix as the first arguments
        views[2] = _prefix;
        // Use payload's remaining arguments (following prefix)
        views[3] = arguments.sliceFrom({ _index: _prefix.len(), newType: SynapseTypes.RAW_BYTES });
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
        views[2] = arguments.sliceFrom({ _index: _prefix.len(), newType: SynapseTypes.RAW_BYTES });
        return TypedMemView.join(views);
    }

    /**
     * @notice Returns a properly typed bytes29 pointer for a system call payload.
     */
    function castToSystemCall(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.SYSTEM_CALL);
    }

    /**
     * @notice Checks that a payload is a formatted System Call.
     */
    function isSystemCall(bytes29 _view) internal pure returns (bool) {
        // Payload needs to exist (system calls are never done via fallback function)
        if (_view.len() < OFFSET_CALLDATA) return false;
        bytes29 _callData = _getCallData(_view);
        // Payload needs to be a proper calldata
        if (!_callData.isCallData()) return false;
        // Payload needs to have at least this amount of argument words
        return _callData.castToCallData().argumentWords() >= CALLDATA_MIN_ARGUMENT_WORDS;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         SYSTEM CALL SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns int value of System Call's recipient (see ISystemRouter.SystemEntity).
     */
    function callRecipient(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SYSTEM_CALL)
        returns (uint8)
    {
        return uint8(_view.indexUint({ _index: OFFSET_RECIPIENT, _bytes: 1 }));
    }

    /**
     * @notice Returns System Call's calldata.
     */
    function callData(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SYSTEM_CALL)
        returns (CallData)
    {
        return _getCallData(_view).castToCallData();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns System Call's calldata WITHOUT checking the view type and
     * without verifying that this is a valid calldata.
     * To be used in `isSystemCall`, where type check is not necessary.
     */
    function _getCallData(bytes29 _view) private pure returns (bytes29) {
        return _view.sliceFrom({ _index: OFFSET_CALLDATA, newType: 0 });
    }
}
