// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Bytes29Test } from "../utils/Bytes29Test.sol";

import { SystemMessage } from "../../contracts/libs/SystemMessage.sol";
import { SynapseTypes } from "../../contracts/libs/SynapseTypes.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";
import { TypeCasts } from "../../contracts/libs/TypeCasts.sol";

contract SystemMessageTest is Bytes29Test {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using SystemMessage for bytes;
    using SystemMessage for bytes29;

    // TODO: tests for Adjust once its structure is finalized
    uint8 internal messageFlag = 0;
    uint8 internal recipient = 42;
    bytes internal payload = "payload";

    function test_formattedCorrectly() public {
        bytes29 _view = _createTestView();

        assertTrue(_view.isSystemMessage());
        (SystemMessage.MessageFlag _messageFlag, bytes29 bodyView) = _view.unpackMessage();
        assertEq(uint8(_messageFlag), messageFlag);

        assertTrue(bodyView.isSystemCall());
        assertEq(bodyView.callRecipient(), recipient);
        assertEq(bodyView.callPayload().clone(), payload);
    }

    function test_incorrectType_unpackMessage() public {
        _prepareMistypedTest(SynapseTypes.SYSTEM_MESSAGE).unpackMessage();
    }

    function test_incorrectType_callRecipient() public {
        _prepareMistypedTest(SynapseTypes.SYSTEM_MESSAGE_CALL).callRecipient();
    }

    function test_incorrectType_callPayload() public {
        _prepareMistypedTest(SynapseTypes.SYSTEM_MESSAGE_CALL).callPayload();
    }

    function test_isSystemMessage_emptyPayload() public {
        bytes memory message = bytes("");
        assert(message.length == 0);
        assertFalse(message.castToSystemMessage().isSystemMessage());
    }

    function test_isSystemMessage_tooShort() public {
        bytes memory message = new bytes(1);
        assert(message.length == SystemMessage.OFFSET_BODY);
        assertFalse(message.castToSystemMessage().isSystemMessage());
    }

    function test_isSystemMessage_invalidMessageFlag() public {
        bytes memory message = abi.encodePacked(uint8(42), new bytes(42));
        assertFalse(message.castToSystemMessage().isSystemMessage());
    }

    function test_isSystemMessage_invalidCall() public {
        bytes memory message = SystemMessage.formatSystemMessage(
            SystemMessage.MessageFlag.Call,
            new bytes(1)
        );
        bytes29 _view = message.castToSystemMessage();
        (, bytes29 _bodyView) = _view.unpackMessage();
        assert(!_bodyView.isSystemCall());
        assertFalse(_view.isSystemMessage());
    }

    function test_isSystemMessage_invalidAdjust() public {
        bytes memory message = SystemMessage.formatSystemMessage(
            SystemMessage.MessageFlag.Adjust,
            bytes("")
        );
        bytes29 _view = message.castToSystemMessage();
        (, bytes29 _bodyView) = _view.unpackMessage();
        assert(!_bodyView.isSystemAdjust());
        assertFalse(_view.isSystemMessage());
    }

    function test_systemRouter() public {
        bytes32 systemRouter = SystemMessage.SYSTEM_ROUTER;
        emit log_bytes32(systemRouter);
        // Check last 20 bytes
        assertEq(TypeCasts.bytes32ToAddress(systemRouter), address(0));
        // Shift 20 bytes left, should be 0xFFFF_FFFF_FFFF
        assertEq(uint256(systemRouter) >> (20 * 8), 2**96 - 1);
    }

    function _createTestView() internal view override returns (bytes29) {
        bytes memory systemCall = SystemMessage.formatSystemCall(recipient, payload);
        return systemCall.castToSystemMessage();
    }
}
