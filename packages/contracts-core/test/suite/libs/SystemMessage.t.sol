// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/SystemMessageHarness.t.sol";

import "../../../contracts/libs/SystemMessage.sol";

// solhint-disable func-name-mixedcase
contract SystemMessageLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    // Mock payload for tests: a selector and two values
    bytes internal constant TEST_MESSAGE_PAYLOAD =
        abi.encodeWithSelector(this.setUp.selector, 1, 2);

    SystemMessageHarness internal libHarness;

    function setUp() public override {
        super.setUp();
        libHarness = new SystemMessageHarness();
    }

    function test_formattedCorrectly(uint8 recipient) public {
        SystemMessage.MessageFlag flag = SystemMessage.MessageFlag.Call;
        // Test formatting
        bytes memory payload = libHarness.formatSystemCall(recipient, TEST_MESSAGE_PAYLOAD);
        assertEq(
            payload,
            abi.encodePacked(flag, recipient, TEST_MESSAGE_PAYLOAD),
            "!formatSystemCall"
        );
        assertEq(
            payload,
            libHarness.formatSystemMessage(flag, abi.encodePacked(recipient, TEST_MESSAGE_PAYLOAD)),
            "!formatSystemMessage"
        );
        // Test formatting checker
        assertTrue(libHarness.isSystemMessage(payload));
        (, bytes memory payloadSystemCall) = libHarness.systemMessageBody(
            SynapseTypes.SYSTEM_MESSAGE,
            payload
        );
        assertTrue(libHarness.isSystemCall(payloadSystemCall));
        // Test getters
        assertEq(
            uint8(libHarness.systemMessageFlag(SynapseTypes.SYSTEM_MESSAGE, payload)),
            uint8(flag),
            "!systemMessageFlag"
        );
        assertEq(
            libHarness.callRecipient(SynapseTypes.SYSTEM_MESSAGE_CALL, payloadSystemCall),
            recipient,
            "!callRecipient"
        );
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToSystemMessage,
            payloadType: SynapseTypes.SYSTEM_MESSAGE,
            payload: payload,
            expectedType: SynapseTypes.SYSTEM_MESSAGE,
            expectedData: payload,
            revertMessage: "!castToSystemMessage"
        });
        checkBytes29Getter({
            getter: libHarness.systemMessageBody,
            payloadType: SynapseTypes.SYSTEM_MESSAGE,
            payload: payload,
            expectedType: SynapseTypes.SYSTEM_MESSAGE_CALL,
            expectedData: abi.encodePacked(recipient, TEST_MESSAGE_PAYLOAD),
            revertMessage: "!systemMessageBody"
        });
        checkBytes29Getter({
            getter: libHarness.callPayload,
            payloadType: SynapseTypes.SYSTEM_MESSAGE_CALL,
            payload: payloadSystemCall,
            expectedType: 0, // TODO: introduce PAYLOAD type
            expectedData: TEST_MESSAGE_PAYLOAD,
            revertMessage: "!systemMessageBody"
        });
    }

    function test_isSystemMessage_invalidMessageFlag(uint8 flag) public {
        // Wrong flag value means payload is not a formatted Report
        vm.assume(flag != uint8(SystemMessage.MessageFlag.Call));
        assertFalse(
            libHarness.isSystemMessage(abi.encodePacked(flag, uint8(0), TEST_MESSAGE_PAYLOAD)),
            "!isSystemMessage: wrong flag"
        );
    }

    function test_isSystemMessage_shortMessagePayload() public {
        // Payloads having not enough length should be considered
        // as unformatted without throwing a revert
        assertFalse(libHarness.isSystemMessage(""), "!isSystemMessage: 0 bytes");
        assertFalse(libHarness.isSystemMessage(new bytes(1)), "!isSystemMessage: 1 byte");
        assertFalse(libHarness.isSystemMessage(new bytes(2)), "!isSystemMessage: 2 bytes");
        // TODO: this shouldn't really be considered a system message (see TODO below)
        assertTrue(libHarness.isSystemMessage(new bytes(3)), "!isSystemMessage: 3 bytes");
    }

    function test_isSystemCall_shortCallPayload() public {
        // Payloads having not enough length should be considered
        // as unformatted without throwing a revert
        assertFalse(libHarness.isSystemCall(""), "!isSystemCall: 0 bytes");
        assertFalse(libHarness.isSystemCall(new bytes(1)), "!isSystemCall: 1 byte");
        // TODO: this shouldn't really be considered a system call
        // A valid system call should have the 4 bytes for selector,
        // and have at least two words (32 bytes chunks) following it
        assertTrue(libHarness.isSystemCall(new bytes(2)), "!isSystemCall: 2 bytes");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_wrongTypeRevert_systemMessageBody(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.SYSTEM_MESSAGE });
        libHarness.systemMessageBody(wrongType, payload);
    }

    function test_wrongTypeRevert_systemMessageFlag(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.SYSTEM_MESSAGE });
        libHarness.systemMessageFlag(wrongType, payload);
    }

    function test_wrongTypeRevert_callPayload(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({
            wrongType: wrongType,
            correctType: SynapseTypes.SYSTEM_MESSAGE_CALL
        });
        libHarness.callPayload(wrongType, payload);
    }

    function test_wrongTypeRevert_callRecipient(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({
            wrongType: wrongType,
            correctType: SynapseTypes.SYSTEM_MESSAGE_CALL
        });
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
