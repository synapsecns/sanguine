// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/MessageHarness.t.sol";

import "../../../contracts/libs/Header.sol";
import "../../../contracts/libs/Message.sol";
import "../../../contracts/libs/Tips.sol";

// solhint-disable func-name-mixedcase
contract MessageLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    MessageHarness internal libHarness;
    // First element is (uint16 messageVersion)
    uint8 internal constant FIRST_ELEMENT_BYTES = 16 / 8;
    bytes internal constant TEST_MESSAGE_BODY = "This is a test message body";

    function setUp() public override {
        super.setUp();
        libHarness = new MessageHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formattedCorrectly(
        uint96 notaryTip,
        uint96 broadcasterTip,
        uint96 proverTip,
        uint96 executorTip,
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds
    ) public {
        // Construct message parts: this has been tested in the dedicated unit tests
        bytes memory tips = Tips.formatTips(notaryTip, broadcasterTip, proverTip, executorTip);
        bytes memory header = Header.formatHeader(
            origin,
            sender,
            nonce,
            destination,
            recipient,
            optimisticSeconds
        );
        // Prepare message
        bytes memory message = libHarness.formatMessage(
            origin,
            sender,
            nonce,
            destination,
            recipient,
            optimisticSeconds,
            tips,
            TEST_MESSAGE_BODY
        );
        // Test formatter against manually constructed payload
        assertEq(
            message,
            constructPayload(Message.MESSAGE_VERSION, header, tips),
            "!formatMessage"
        );
        // All formatters should return the same results
        assertEq(
            message,
            libHarness.formatMessage(header, tips, TEST_MESSAGE_BODY),
            "!formatMessage: 3 args variant"
        );
        // Test formatting checker
        assertTrue(libHarness.isMessage(message), "!isMessage");
        // Test getters (most getters are tested in Header, Tips tests)
        assertEq(
            libHarness.messageVersion(SynapseTypes.MESSAGE, message),
            Message.MESSAGE_VERSION,
            "!messageVersion"
        );
        assertEq(
            libHarness.messageHash(header, tips, TEST_MESSAGE_BODY),
            keccak256(message),
            "!messageHash"
        );
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToMessage,
            payloadType: SynapseTypes.MESSAGE,
            payload: message,
            expectedType: SynapseTypes.MESSAGE,
            expectedData: message,
            revertMessage: "!castToMessage"
        });
        checkBytes29Getter({
            getter: libHarness.header,
            payloadType: SynapseTypes.MESSAGE,
            payload: message,
            expectedType: SynapseTypes.MESSAGE_HEADER,
            expectedData: header,
            revertMessage: "!header"
        });
        checkBytes29Getter({
            getter: libHarness.tips,
            payloadType: SynapseTypes.MESSAGE,
            payload: message,
            expectedType: SynapseTypes.MESSAGE_TIPS,
            expectedData: tips,
            revertMessage: "!tips"
        });
        checkBytes29Getter({
            getter: libHarness.body,
            payloadType: SynapseTypes.MESSAGE,
            payload: message,
            expectedType: SynapseTypes.MESSAGE_BODY,
            expectedData: TEST_MESSAGE_BODY,
            revertMessage: "!body"
        });
    }

    function test_isMessage_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Message's first element (uint16 messageVersion)
        // should be correctly treated as unformatted (i.e. with no reverts)
        assertFalse(
            libHarness.isMessage(createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data)),
            "!isHeader: short payload"
        );
    }

    function test_isMessage_wrongMessageVersion(uint16 messageVersion) public {
        // Wrong message version value means payload is not a formatted Message
        vm.assume(messageVersion != Message.MESSAGE_VERSION);
        assertFalse(
            libHarness.isMessage(
                constructPayload(messageVersion, createTestHeader(), createTestTips())
            ),
            "!isMessage: wrong version"
        );
    }

    function test_isMessage_emptyBody() public {
        // A formatted Message could have an empty body
        assertTrue(
            libHarness.isMessage(
                libHarness.formatMessage(createTestHeader(), createTestTips(), "")
            ),
            "!isMessage: empty body"
        );
    }

    function test_isMessage_emptyEverything() public {
        // Empty header or tips means payload is not a formatted Message
        // empty header and tips
        assertFalse(
            libHarness.isMessage(libHarness.formatMessage("", "", "")),
            "!isMessage: header,tips,body empty"
        );
        assertFalse(
            libHarness.isMessage(libHarness.formatMessage("", "", TEST_MESSAGE_BODY)),
            "!isMessage: header,tips empty"
        );
        // empty header
        assertFalse(
            libHarness.isMessage(libHarness.formatMessage("", createTestTips(), "")),
            "!isMessage: header,body empty"
        );
        assertFalse(
            libHarness.isMessage(libHarness.formatMessage("", createTestTips(), TEST_MESSAGE_BODY)),
            "!isMessage: header empty"
        );
        // empty tips
        assertFalse(
            libHarness.isMessage(libHarness.formatMessage(createTestHeader(), "", "")),
            "!isMessage: tips,body empty"
        );
        assertFalse(
            libHarness.isMessage(
                libHarness.formatMessage(createTestHeader(), "", TEST_MESSAGE_BODY)
            ),
            "!isMessage: tips empty"
        );
    }

    function test_isMessage_incorrectLengths() public {
        uint16 version = Message.MESSAGE_VERSION;
        bytes memory header = createTestHeader();
        bytes memory tips = createTestTips();
        // With an empty body, specifying a longer length leads
        // to a memory view overrun. Should be treated without a revert/panic.
        assertFalse(
            libHarness.isMessage(
                abi.encodePacked(
                    version,
                    uint16(header.length + 1),
                    uint16(tips.length),
                    header,
                    tips
                )
            ),
            "!isMessage: bad header length"
        );
        assertFalse(
            libHarness.isMessage(
                abi.encodePacked(
                    version,
                    uint16(header.length),
                    uint16(tips.length + 1),
                    header,
                    tips
                )
            ),
            "!isMessage: bad tips length"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_wrongTypeRevert_messageVersion(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE });
        libHarness.messageVersion(wrongType, payload);
    }

    function test_wrongTypeRevert_header(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE });
        libHarness.header(wrongType, payload);
    }

    function test_wrongTypeRevert_tips(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE });
        libHarness.tips(wrongType, payload);
    }

    function test_wrongTypeRevert_body(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE });
        libHarness.body(wrongType, payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatMessage(createTestHeader(), createTestTips(), TEST_MESSAGE_BODY);
    }

    function createTestHeader() public pure returns (bytes memory) {
        return Header.formatHeader(0, bytes32(0), 0, 0, bytes32(0), 0);
    }

    function createTestTips() public pure returns (bytes memory) {
        return Tips.emptyTips();
    }

    function constructPayload(
        uint16 messageVersion,
        bytes memory header,
        bytes memory tips
    ) public pure returns (bytes memory) {
        return
            abi.encodePacked(
                messageVersion,
                uint16(header.length),
                uint16(tips.length),
                header,
                tips,
                TEST_MESSAGE_BODY
            );
    }
}
