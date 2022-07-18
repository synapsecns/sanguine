// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { SystemMessage } from "../contracts/system/SystemMessage.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";
import { TypeCasts } from "../contracts/libs/TypeCasts.sol";

contract SystemMessageTest is Test {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using SystemMessage for bytes29;

    bytes internal payload = "payload";

    function test_formatSystemMessage() public {
        for (uint8 t = 0; t <= uint8(SystemMessage.SystemMessageType.Adjust); ++t) {
            bytes memory formatted = SystemMessage.formatSystemMessage(
                SystemMessage.SystemMessageType(t),
                payload
            );
            (SystemMessage.SystemMessageType _type, bytes29 _view) = formatted
                .ref(0)
                .systemMessage();
            assertEq(uint8(_type), t);
            assertEq(_view.clone(), payload);
        }
    }

    function test_formatCall() public {
        for (uint8 r = 0; r < 4; ++r) {
            bytes memory formatted = SystemMessage.formatCall(r, payload);
            (SystemMessage.SystemMessageType _type, bytes29 _call) = formatted
                .ref(0)
                .systemMessage();
            assertEq(uint8(_type), uint8(SystemMessage.SystemMessageType.Call));
            assertEq(_call.callRecipient(), r);
            assertEq(_call.callPayload().clone(), payload);
        }
    }

    function test_systemSender() public {
        bytes32 systemSender = SystemMessage.SYSTEM_SENDER;
        emit log_bytes32(systemSender);
        // Check last 20 bytes
        assertEq(TypeCasts.bytes32ToAddress(systemSender), address(0));
        // Shift 20 bytes left, should be 0xFFFF_FFFF_FFFF
        assertEq(uint256(systemSender) >> (20 * 8), 2**96 - 1);
    }
}
