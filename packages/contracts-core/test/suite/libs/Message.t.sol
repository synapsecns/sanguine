// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseLibraryTest, TypedMemView } from "../../utils/SynapseLibraryTest.t.sol";
import { MessageHarness } from "../../harnesses/libs/MessageHarness.t.sol";

import { HeaderLib } from "../../../contracts/libs/Header.sol";
import { MessageLib } from "../../../contracts/libs/Message.sol";
import { TipsLib } from "../../../contracts/libs/Tips.sol";

// solhint-disable func-name-mixedcase
contract MessageLibraryTest is SynapseLibraryTest {
    struct RawHeader {
        uint32 origin;
        bytes32 sender;
        uint32 nonce;
        uint32 destination;
        bytes32 recipient;
        uint32 optimisticSeconds;
    }

    struct RawTips {
        uint96 notaryTip;
        uint96 broadcasterTip;
        uint96 proverTip;
        uint96 executorTip;
    }

    MessageHarness internal libHarness;
    // First element is (uint16 messageVersion)
    uint8 internal constant FIRST_ELEMENT_BYTES = 16 / 8;
    bytes internal constant TEST_MESSAGE_BODY = "This is a test message body";

    function setUp() public {
        libHarness = new MessageHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formattedCorrectly(RawHeader memory rh, RawTips memory rt) public {
        // Construct message parts: this has been tested in the dedicated unit tests
        bytes memory tipsPayload = TipsLib.formatTips(
            rt.notaryTip,
            rt.broadcasterTip,
            rt.proverTip,
            rt.executorTip
        );
        bytes memory header = HeaderLib.formatHeader(
            rh.origin,
            rh.sender,
            rh.nonce,
            rh.destination,
            rh.recipient,
            rh.optimisticSeconds
        );
        // Prepare message
        bytes memory message = libHarness.formatMessage(
            rh.origin,
            rh.sender,
            rh.nonce,
            rh.destination,
            rh.recipient,
            rh.optimisticSeconds,
            rt.notaryTip,
            rt.broadcasterTip,
            rt.proverTip,
            rt.executorTip,
            TEST_MESSAGE_BODY
        );
        // Test formatter against manually constructed payload
        assertEq(
            message,
            constructPayload(MessageLib.MESSAGE_VERSION, header, tipsPayload),
            "!formatMessage"
        );
        // All formatters should return the same results
        assertEq(
            message,
            libHarness.formatMessage(header, tipsPayload, TEST_MESSAGE_BODY),
            "!formatMessage: 3 args variant"
        );
        assertEq(
            message,
            libHarness.formatMessage(
                rh.origin,
                rh.sender,
                rh.nonce,
                rh.destination,
                rh.recipient,
                rh.optimisticSeconds,
                tipsPayload,
                TEST_MESSAGE_BODY
            ),
            "!formatMessage: unpacked header variant"
        );
        // Test formatting checker
        checkCastToMessage({ payload: message, isMessage: true });
        // Test getters (most getters are tested in Header, Tips tests)
        assertEq(libHarness.version(message), MessageLib.MESSAGE_VERSION, "!messageVersion");
        assertEq(libHarness.header(message), header, "!header");
        assertEq(libHarness.tips(message), tipsPayload, "!tips");
        assertEq(libHarness.body(message), TEST_MESSAGE_BODY, "!body");
        assertEq(libHarness.leaf(message), keccak256(message), "!leaf");
    }

    function test_constants() public {
        assertEq(libHarness.messageVersion(), 1);
        // TODO: figure out why this doesn't mark offsetVersion as covered
        assertEq(libHarness.offsetVersion(), 0);
        // 2 + 2 + 2
        assertEq(libHarness.offsetHeader(), 6);
    }

    function test_isMessage_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Message's first element (uint16 messageVersion)
        // should be correctly treated as unformatted (i.e. with no reverts)
        bytes memory payload = createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data);
        checkCastToMessage({ payload: payload, isMessage: false });
    }

    function test_isMessage_wrongMessageVersion(uint16 messageVersion) public {
        // Wrong message version value means payload is not a formatted Message
        vm.assume(messageVersion != MessageLib.MESSAGE_VERSION);
        bytes memory payload = constructPayload(
            messageVersion,
            createTestHeader(),
            createTestTips()
        );
        checkCastToMessage({ payload: payload, isMessage: false });
    }

    function test_isMessage_emptyBody() public {
        // A formatted Message could have an empty body
        bytes memory payload = libHarness.formatMessage(createTestHeader(), createTestTips(), "");
        checkCastToMessage({ payload: payload, isMessage: true });
    }

    function test_isMessage_emptyEverything() public {
        // Empty header or tips means payload is not a formatted Message
        // empty header and tips
        bytes memory payload = libHarness.formatMessage("", "", "");
        checkCastToMessage({ payload: payload, isMessage: false });
        payload = libHarness.formatMessage("", "", TEST_MESSAGE_BODY);
        checkCastToMessage({ payload: payload, isMessage: false });
        // empty header
        payload = libHarness.formatMessage("", createTestTips(), "");
        checkCastToMessage({ payload: payload, isMessage: false });
        payload = libHarness.formatMessage("", createTestTips(), TEST_MESSAGE_BODY);
        checkCastToMessage({ payload: payload, isMessage: false });
        // empty tips
        payload = libHarness.formatMessage(createTestHeader(), "", "");
        checkCastToMessage({ payload: payload, isMessage: false });
        payload = libHarness.formatMessage(createTestHeader(), "", TEST_MESSAGE_BODY);
        checkCastToMessage({ payload: payload, isMessage: false });
    }

    function test_isMessage_incorrectLengths() public {
        uint16 version = MessageLib.MESSAGE_VERSION;
        bytes memory header = createTestHeader();
        bytes memory tipsPayload = createTestTips();
        // With an empty body, specifying a longer length leads
        // to a memory view overrun. Should be treated without a revert/panic.
        bytes memory payload = abi.encodePacked(
            version,
            uint16(header.length + 1),
            uint16(tipsPayload.length),
            header,
            tipsPayload
        );
        checkCastToMessage({ payload: payload, isMessage: false });
        payload = abi.encodePacked(
            version,
            uint16(header.length),
            uint16(tipsPayload.length + 1),
            header,
            tipsPayload
        );
        checkCastToMessage({ payload: payload, isMessage: false });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatMessage(createTestHeader(), createTestTips(), TEST_MESSAGE_BODY);
    }

    function createTestHeader() public pure returns (bytes memory) {
        return HeaderLib.formatHeader(0, bytes32(0), 0, 0, bytes32(0), 0);
    }

    function createTestTips() public pure returns (bytes memory) {
        return TipsLib.emptyTips();
    }

    function constructPayload(
        uint16 messageVersion,
        bytes memory header,
        bytes memory tipsPayload
    ) public pure returns (bytes memory) {
        return
            abi.encodePacked(
                messageVersion,
                uint16(header.length),
                uint16(tipsPayload.length),
                header,
                tipsPayload,
                TEST_MESSAGE_BODY
            );
    }
}
