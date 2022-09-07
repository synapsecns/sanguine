// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Bytes29Test } from "../utils/Bytes29Test.sol";
import { Header } from "../../contracts/libs/Header.sol";
import { Message } from "../../contracts/libs/Message.sol";
import { Tips } from "../../contracts/libs/Tips.sol";

import { SynapseTypes } from "../../contracts/libs/SynapseTypes.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase

contract MessageTest is Bytes29Test {
    using Header for bytes;
    using Header for bytes29;
    using Message for bytes;
    using Message for bytes29;
    using Tips for bytes;
    using Tips for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint32 internal constant ORIGIN = 1000;
    bytes32 internal constant SENDER = bytes32("AAAA THE SENDOOOOOR");
    uint32 internal constant NONCE = 42;
    uint32 internal constant DESTINATION = 2000;
    uint32 internal constant OPTIMISTIC_SECONDS = 4;
    bytes32 internal constant RECIPIENT = bytes32("AAAA THE RECEIVOOOR");

    uint96 internal constant UPDATER_TIP = 1234;
    uint96 internal constant RELAYER_TIP = 3456;
    uint96 internal constant PROVER_TIP = 5678;
    uint96 internal constant PROCESSOR_TIP = 7890;

    bytes internal header;
    bytes internal tips;
    bytes internal body = bytes("Messagoooor");

    function setUp() public {
        header = Header.formatHeader(
            ORIGIN,
            SENDER,
            NONCE,
            DESTINATION,
            RECIPIENT,
            OPTIMISTIC_SECONDS
        );
        tips = Tips.formatTips(UPDATER_TIP, RELAYER_TIP, PROVER_TIP, PROCESSOR_TIP);
    }

    function test_formattedCorrectly() public {
        bytes29 _view = _createTestView();
        _checkFormattedCorrectly(_view);
    }

    function test_formattedCorrectly_emptyBody() public {
        body = bytes("");
        bytes29 _view = _createTestView();
        _checkFormattedCorrectly(_view);
    }

    function test_messageHash() public {
        bytes29 _view = _createTestView();
        assertEq(_view.keccak(), keccak256(_recreateTestMessage()));
    }

    function test_incorrectType_messageVersion() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE).messageVersion();
    }

    function test_incorrectType_header() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE).header();
    }

    function test_incorrectType_tips() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE).tips();
    }

    function test_incorrectType_body() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE).body();
    }

    function test_isMessage_incorrectVersion() public {
        // set version to 0
        bytes memory message = _modifyTestView(Message.OFFSET_VERSION, abi.encodePacked(uint16(0)));
        assertFalse(message.castToMessage().isMessage());
    }

    function test_isMessage_emptyEverything() public {
        header = bytes("");
        tips = bytes("");
        body = bytes("");
        bytes memory message = _recreateTestMessage();
        assertFalse(message.castToMessage().isMessage());
    }

    function test_isMessage_incorrectHeaderLength() public {
        body = bytes("");
        bytes memory message = _modifyTestView(
            Message.TWO_BYTES * uint8(Message.Parts.Header),
            abi.encodePacked(uint16(header.length + 1))
        );
        // With an empty body, header.length + tips.length will overrun the memory view
        assertFalse(message.castToMessage().isMessage());
    }

    function test_isMessage_incorrectTipsLength() public {
        body = bytes("");
        bytes memory message = _modifyTestView(
            Message.TWO_BYTES * uint8(Message.Parts.Tips),
            abi.encodePacked(uint16(tips.length + 1))
        );
        // With an empty body, header.length + tips.length will overrun the memory view
        assertFalse(message.castToMessage().isMessage());
    }

    function test_isMessage_incorrectHeader() public {
        // include 1 empty byte before the header
        header = abi.encodePacked(new bytes(1), header);
        // should not be a formatted Header payload
        assertFalse(header.castToHeader().isHeader());
        bytes memory message = _recreateTestMessage();
        // Message could not have unformatted Header
        assertFalse(message.castToMessage().isMessage());
    }

    function test_isMessage_emptyHeaderPayload() public {
        header = bytes("");
        bytes memory message = _recreateTestMessage();
        // Header payload could not be empty
        assertFalse(message.castToMessage().isMessage());
    }

    function test_isMessage_incorrectTips() public {
        // include 1 empty byte before the tips
        tips = abi.encodePacked(new bytes(1), tips);
        // should not be a formatted Tips payload
        assertFalse(tips.castToTips().isTips());
        bytes memory message = _recreateTestMessage();
        // Message could not have unformatted Tips
        assertFalse(message.castToMessage().isMessage());
    }

    function test_isMessage_emptyTipsPayload() public {
        tips = bytes("");
        bytes memory message = _recreateTestMessage();
        // Tips payload could not be empty
        assertFalse(message.castToMessage().isMessage());
    }

    function _checkFormattedCorrectly(bytes29 _view) internal {
        assertTrue(_view.isMessage(), "!isMessage");

        assertEq(_view.messageVersion(), Message.MESSAGE_VERSION, "!version");
        assertEq(_view.header().clone(), header, "!header");
        assertEq(_view.tips().clone(), tips, "!tips");
        assertEq(_view.body().clone(), body, "!body");
    }

    function _createTestView() internal view override returns (bytes29) {
        bytes memory message = Message.formatMessage(header, tips, body);
        return message.castToMessage();
    }

    function _recreateTestMessage() internal view returns (bytes memory) {
        return
            abi.encodePacked(
                uint16(Message.MESSAGE_VERSION),
                uint16(header.length),
                uint16(tips.length),
                header,
                tips,
                body
            );
    }
}
