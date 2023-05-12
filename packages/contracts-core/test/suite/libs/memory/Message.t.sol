// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {UnformattedMessage} from "../../../../contracts/libs/Errors.sol";
import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {MessageHarness} from "../../../harnesses/libs/memory/MessageHarness.t.sol";

import {HEADER_LENGTH} from "../../../../contracts/libs/Constants.sol";
import {Header} from "../../../../contracts/libs/stack/Header.sol";
import {MessageLib} from "../../../../contracts/libs/memory/Message.sol";
import {TipsLib} from "../../../../contracts/libs/stack/Tips.sol";

import {MessageFlag, RawHeader, RawBaseMessage, RawMessage} from "../../../utils/libs/SynapseStructs.t.sol";
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
        rh.flag = uint8(MessageFlag.Base);
        bytes memory body = rbm.formatBaseMessage();
        check_formatMessage(rh, body, rbm.castToBaseMessage().leaf());
    }

    function check_formatMessage(RawHeader memory rh, bytes memory body, bytes32 bodyLeaf) public {
        Header header = rh.castToHeader();
        uint136 encodedHeader = rh.encodeHeader();
        // Prepare message
        bytes memory message = libHarness.formatMessage(header, body);
        // Test formatter
        assertEq(message, abi.encodePacked(header, body), "!formatMessage");
        // Test formatting checker
        checkCastToMessage({payload: message, isMessage: true});
        // Test getters
        assertEq(Header.unwrap(libHarness.header(message)), encodedHeader, "!header");
        assertEq(libHarness.body(message), body, "!body");
        // Test hashing
        assertEq(libHarness.leaf(message), keccak256(bytes.concat(header.leaf(), bodyLeaf)), "!leaf");
    }

    function test_isMessage_flagOutOfRange(uint8 flag, uint128 remainder) public {
        // Make sure flag does NOT fit into MessageFlag enum
        flag = uint8(bound(flag, uint8(type(MessageFlag).max) + 1, type(uint8).max));
        // Use incorrect flag and empty body
        bytes memory payload = abi.encodePacked(flag, remainder);
        checkCastToMessage({payload: payload, isMessage: false});
    }

    function test_isMessage_tooShort(uint256 length) public {
        // Make sure length is shorter than (flag + header)
        length = length % (1 + HEADER_LENGTH);
        bytes memory payload = new bytes(length);
        checkCastToMessage({payload: payload, isMessage: false});
    }

    function test_isMessage_base(RawMessage memory rm, uint8 lengthBM) public {
        rm.header.flag = uint8(MessageFlag.Base);
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
            vm.expectRevert(UnformattedMessage.selector);
            libHarness.castToMessage(payload);
        }
    }
}
