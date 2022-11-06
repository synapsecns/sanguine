// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/SystemCallHarness.t.sol";

import "../../../contracts/libs/SystemCall.sol";

// solhint-disable func-name-mixedcase
contract SystemCallLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    // Mock payload for tests: a selector and two values
    bytes internal constant TEST_MESSAGE_PAYLOAD =
        abi.encodeWithSelector(this.setUp.selector, 1, 2);

    // First element is (uint8 recipient)
    uint8 internal constant FIRST_ELEMENT_BYTES = 8 / 8;

    uint8 internal constant MIN_ARGUMENT_WORDS = 2;
    uint8 internal constant MIN_SYSTEM_CALL_LENGTH = 1 + 4 + 32 * MIN_ARGUMENT_WORDS;

    SystemCallHarness internal libHarness;

    function setUp() public override {
        super.setUp();
        libHarness = new SystemCallHarness();
    }

    function test_formattedCorrectly(uint8 recipient) public {
        // Test formatting
        bytes memory payload = libHarness.formatSystemCall(recipient, TEST_MESSAGE_PAYLOAD);
        assertEq(payload, abi.encodePacked(recipient, TEST_MESSAGE_PAYLOAD), "!formatSystemCall");
        // Test formatting checker
        assertTrue(libHarness.isSystemCall(payload));
        // Test getters
        assertEq(
            libHarness.callRecipient(SynapseTypes.SYSTEM_CALL, payload),
            recipient,
            "!callRecipient"
        );
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToSystemCall,
            payloadType: SynapseTypes.SYSTEM_CALL,
            payload: payload,
            expectedType: SynapseTypes.SYSTEM_CALL,
            expectedData: payload,
            revertMessage: "!castToSystemCall"
        });
        checkBytes29Getter({
            getter: libHarness.callPayload,
            payloadType: SynapseTypes.SYSTEM_CALL,
            payload: payload,
            expectedType: SynapseTypes.CALL_PAYLOAD,
            expectedData: TEST_MESSAGE_PAYLOAD,
            revertMessage: "!systemMessageBody"
        });
    }

    function test_isSystemCall_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than SystemCall's first element (uint8 recipient)
        // should be correctly treated as unformatted (i.e. with no reverts)
        assertFalse(
            libHarness.isSystemCall(createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data)),
            "!isSystemCall: short payload"
        );
    }

    function test_isSystemCall_shortCallPayload(uint8 length) public {
        vm.assume(length != 0);
        vm.assume(length < MIN_SYSTEM_CALL_LENGTH);
        // Payloads having not enough length should be considered
        // as unformatted without throwing a revert
        assertFalse(libHarness.isSystemCall(new bytes(length)), "!isSystemCall: short payload");
    }

    function test_isSystemCall_incorrectPayloadLength(uint16 length) public {
        vm.assume(length >= MIN_SYSTEM_CALL_LENGTH);
        // System call payload is
        // - recipient (1 byte)
        // - selector (4 bytes)
        // - arguments (unknown amount of 32-byte words, at least two words)
        // Thus, payload should have a length of 1 + 4 + 32 * words
        vm.assume((length - 5) % 32 != 0);
        assertFalse(libHarness.isSystemCall(new bytes(length)), "!isSystemCall: incorrect length");
    }

    function test_isSystemCall_correctPayloadLength(uint8 argumentWords) public {
        vm.assume(argumentWords >= MIN_ARGUMENT_WORDS);
        // System call payload is
        // - recipient (1 byte)
        // - selector (4 bytes)
        // - arguments (unknown amount of 32-byte words, at least two words)
        // Thus, payload should have a length of 1 + 4 + 32 * words
        uint256 length = 5 + 32 * uint256(argumentWords);
        assertTrue(libHarness.isSystemCall(new bytes(length)), "!isSystemCall: correct length");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_wrongTypeRevert_callPayload(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.SYSTEM_CALL });
        libHarness.callPayload(wrongType, payload);
    }

    function test_wrongTypeRevert_callRecipient(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.SYSTEM_CALL });
        libHarness.callRecipient(wrongType, payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           TESTS: CONSTANTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_constant_systemRouter() public {
        // SYSTEM_ROUTER constant should have
        // highest 96 bits set
        // lowest 160 bits unset
        uint256 systemRouter = uint256(libHarness.systemRouter());
        // Clear 160 lowest bits => check (256 - 160 = 96) highest bits
        assertEq(systemRouter >> 160, type(uint96).max, "!SYSTEM_ROUTER: highest bits");
        // Clear 96 highest bits => check (256 - 96 = 160) lowest bits.
        assertEq(systemRouter << 96, 0, "!SYSTEM_ROUTER: lowest bits");
        assertEq(
            bytes32ToAddress(libHarness.systemRouter()),
            address(0),
            "!SYSTEM_ROUTER: cast to address"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatSystemCall(0, TEST_MESSAGE_PAYLOAD);
    }
}
