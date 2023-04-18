// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {MessageHarness} from "../../harnesses/libs/MessageHarness.t.sol";

import {HeaderLib, HEADER_LENGTH} from "../../../contracts/libs/Header.sol";
import {MessageLib} from "../../../contracts/libs/Message.sol";
import {TipsLib} from "../../../contracts/libs/Tips.sol";

import {MessageFlag, RawHeader, RawBaseMessage, RawMessage} from "../../utils/libs/SynapseStructs.t.sol";
import {BaseMessageLibraryTest} from "./BaseMessage.t.sol";

// solhint-disable func-name-mixedcase
contract MessageLibraryTest is SynapseLibraryTest {
    MessageHarness internal libHarness;

    function setUp() public {
        libHarness = new MessageHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatMessage_base(RawHeader memory rh, RawBaseMessage memory rbm) public {
        // Construct message parts: this has been tested in the dedicated unit tests
        MessageFlag flag = MessageFlag.Base;
        bytes memory header = rh.formatHeader();
        bytes memory body = rbm.formatBaseMessage();
        check_formatMessage(flag, header, body);
    }

    function check_formatMessage(MessageFlag flag, bytes memory header, bytes memory body) public {
        // Prepare message
        bytes memory message = libHarness.formatMessage(flag, header, body);
        // Test formatter
        assertEq(message, abi.encodePacked(flag, header, body), "!formatMessage");
        // Test formatting checker
        checkCastToMessage({payload: message, isMessage: true});
        // Test getters
        assertEq(uint8(libHarness.flag(message)), uint8(flag), "!flag");
        assertEq(libHarness.header(message), header, "!header");
        assertEq(libHarness.body(message), body, "!body");
        // Test hashing
        assertEq(libHarness.leaf(message), keccak256(message), "!leaf");
    }

    function test_isMessage_flagOutOfRange(uint8 flag, RawHeader memory rh) public {
        // Make sure flag does NOT fit into MessageFlag enum
        flag = uint8(bound(flag, uint8(type(MessageFlag).max) + 1, type(uint8).max));
        // Use incorrect flag and empty body
        bytes memory payload = abi.encodePacked(flag, rh.formatHeader());
        checkCastToMessage({payload: payload, isMessage: false});
    }

    function test_isMessage_tooShort(uint256 length) public {
        // Make sure length is shorter than (flag + header)
        length = length % (1 + HEADER_LENGTH);
        bytes memory payload = new bytes(length);
        checkCastToMessage({payload: payload, isMessage: false});
    }

    function test_isMessage_base(RawMessage memory rm, uint8 lengthBM) public {
        rm.flag = uint8(MessageFlag.Base);
        rm.body = new bytes(lengthBM);
        bytes memory payload = rm.formatMessage();
        checkCastToMessage({payload: payload, isMessage: lengthBM >= MIN_BASE_MESSAGE_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToMessage(bytes memory payload, bool isMessage) public {
        if (isMessage) {
            assertTrue(libHarness.isMessage(payload), "!isMessage: when valid");
            assertEq(libHarness.castToMessage(payload), payload, "!castToMessage: when valid");
        } else {
            assertFalse(libHarness.isMessage(payload), "!isMessage: when valid");
            vm.expectRevert("Not a message payload");
            libHarness.castToMessage(payload);
        }
    }
}
