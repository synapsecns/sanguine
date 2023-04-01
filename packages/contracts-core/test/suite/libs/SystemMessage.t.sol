// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {SystemMessageHarness} from "../../harnesses/libs/SystemMessageHarness.t.sol";
import {ByteStringTools} from "../../tools/libs/ByteStringTools.t.sol";

// solhint-disable func-name-mixedcase
contract SystemMessageLibraryTest is ByteStringTools, SynapseLibraryTest {
    using TypedMemView for bytes;

    // Mock payload for tests: a selector and values for three security arguments
    bytes internal constant TEST_MESSAGE_PAYLOAD = abi.encodeWithSelector(this.setUp.selector, 1, 2, 3);

    // First element is (uint8 recipient)
    uint8 internal constant FIRST_ELEMENT_BYTES = 8 / 8;

    uint8 internal constant MIN_ARGUMENT_WORDS = 3;
    uint8 internal constant MIN_SYSTEM_MESSAGE_LENGTH = 1 + 4 + 32 * MIN_ARGUMENT_WORDS;

    SystemMessageHarness internal libHarness;

    function setUp() public {
        libHarness = new SystemMessageHarness();
    }

    function test_formattedCorrectly(uint8 recipient, uint8 wordsPrefix, uint8 wordsFollowing) public {
        // Set a sensible limit for the total payload length
        vm.assume(uint256(wordsPrefix) + wordsFollowing >= MIN_ARGUMENT_WORDS);
        vm.assume((uint256(wordsPrefix) + wordsFollowing) * 32 <= MAX_CONTENT_BYTES);
        bytes4 selector = this.setUp.selector;
        // Create "random" arguments and new/old prefix with different random seeds
        bytes memory prefixOld = createTestArguments(wordsPrefix, "prefixOld");
        bytes memory following = createTestArguments(wordsFollowing, "following");
        bytes memory prefixNew = createTestArguments(wordsPrefix, "prefixNew");
        bytes memory callData = bytes.concat(selector, prefixOld, following);
        // Format the calldata
        bytes memory adjustedCallData = libHarness.formatAdjustedCallData({callData_: callData, prefix: prefixNew});
        // Test formatter against manually constructed payload
        assertEq(adjustedCallData, bytes.concat(selector, prefixNew, following), "!formatAdjustedCallData");
        // Format the system message
        bytes memory systemMessage =
            libHarness.formatSystemMessage({systemRecipient: recipient, callData_: callData, prefix: prefixNew});
        checkCastToSystemMessage({payload: systemMessage, isSystemMessage: true});
        // Test formatter against manually constructed payload
        assertEq(systemMessage, abi.encodePacked(recipient, selector, prefixNew, following), "!formatSystemMessage");
        // Test getters
        assertEq(libHarness.callRecipient(systemMessage), recipient, "!callRecipient");
        assertEq(libHarness.callData(systemMessage), adjustedCallData, "!callData");
    }

    function test_formatAdjusted_revert_shortPayload(uint8 recipient, uint8 wordsPayload, uint8 bytesExtra) public {
        uint256 length = uint256(wordsPayload) * 32;
        // Set a sensible limit for the total payload length
        vm.assume(length <= MAX_CONTENT_BYTES);
        vm.assume(bytesExtra != 0);
        // Let "arguments" be shorter than the prefix
        bytes memory callData = bytes.concat(this.setUp.selector, new bytes(length));
        bytes memory prefix = new bytes(length + bytesExtra);
        vm.expectRevert("Payload too short");
        libHarness.formatAdjustedCallData({callData_: callData, prefix: prefix});
        vm.expectRevert("Payload too short");
        libHarness.formatSystemMessage({systemRecipient: recipient, callData_: callData, prefix: prefix});
    }

    function test_isSystemMessage_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than SystemMessage's first element (uint8 recipient)
        // should be correctly treated as unformatted (i.e. with no reverts)
        bytes memory payload = createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data);
        checkCastToSystemMessage({payload: payload, isSystemMessage: false});
    }

    function test_isSystemMessage_shortCallData(uint8 length) public {
        vm.assume(length != 0);
        vm.assume(length < MIN_SYSTEM_MESSAGE_LENGTH);
        // Payloads having not enough length should be considered
        // as unformatted without throwing a revert
        bytes memory payload = new bytes(length);
        checkCastToSystemMessage({payload: payload, isSystemMessage: false});
    }

    function test_isSystemMessage_incorrectPayloadLength(uint16 length) public {
        vm.assume(length >= MIN_SYSTEM_MESSAGE_LENGTH);
        // System call payload is
        // - recipient (1 byte)
        // - selector (4 bytes)
        // - arguments (unknown amount of 32-byte words, at least two words)
        // Thus, payload should have a length of 1 + 4 + 32 * words
        vm.assume((length - 5) % 32 != 0);
        bytes memory payload = new bytes(length);
        checkCastToSystemMessage({payload: payload, isSystemMessage: false});
    }

    function test_isSystemMessage_correctPayloadLength(uint8 argumentWords) public {
        vm.assume(argumentWords >= MIN_ARGUMENT_WORDS);
        // System call payload is
        // - recipient (1 byte)
        // - selector (4 bytes)
        // - arguments (unknown amount of 32-byte words, at least two words)
        // Thus, payload should have a length of 1 + 4 + 32 * words
        uint256 length = 5 + 32 * uint256(argumentWords);
        bytes memory payload = new bytes(length);
        checkCastToSystemMessage({payload: payload, isSystemMessage: true});
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
        assertEq(bytes32ToAddress(libHarness.systemRouter()), address(0), "!SYSTEM_ROUTER: cast to address");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToSystemMessage(bytes memory payload, bool isSystemMessage) public {
        if (isSystemMessage) {
            assertTrue(libHarness.isSystemMessage(payload), "!isSystemMessage: when valid");
            assertEq(libHarness.castToSystemMessage(payload), payload, "!castToSystemMessage: when valid");
        } else {
            assertFalse(libHarness.isSystemMessage(payload), "!isSystemMessage: when valid");
            vm.expectRevert("Not a system message");
            libHarness.castToSystemMessage(payload);
        }
    }

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatSystemMessage({systemRecipient: 0, callData_: TEST_MESSAGE_PAYLOAD, prefix: ""});
    }
}
