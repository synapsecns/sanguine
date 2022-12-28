// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/ByteStringHarness.t.sol";
import "../../tools/libs/ByteStringTools.t.sol";

import "../../../contracts/libs/ByteString.sol";

// solhint-disable func-name-mixedcase
contract ByteStringLibraryTest is ByteStringTools, SynapseLibraryTest {
    using ByteString for bytes;

    ByteStringHarness internal libHarness;
    uint256 internal extraBytes = 0;
    bytes4 internal selector = this.setUp.selector;

    // First element is (bytes4 selector)
    uint8 internal constant FIRST_ELEMENT_BYTES = 4;

    function setUp() public override {
        super.setUp();
        libHarness = new ByteStringHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formattedCorrectly_callData(uint8 words) public {
        // Set a sensible limit for the total payload length
        vm.assume(uint256(words) * 32 <= MAX_MESSAGE_BODY_BYTES);
        bytes memory arguments = createTestArguments(words, "seed");
        bytes memory callData = abi.encodePacked(selector, arguments);
        require(callData.length == selector.length + 32 * uint256(words), "!length");
        // Test related constants
        assertEq(libHarness.selectorLength(), selector.length, "!selectorLength");
        // Test formatting checker
        assertTrue(libHarness.isCallData(callData), "!isCallData");
        assertEq(libHarness.castToCallData(callData), callData, "!castToCallData");
        // Test CallData getters
        assertEq(libHarness.argumentWords(callData), words, "!argumentWords");
        assertEq(libHarness.callSelector(callData), bytes.concat(selector), "!callSelector");
        assertEq(libHarness.arguments(callData), arguments, "!arguments");
    }

    function test_formattedCorrectly_callData_adjusted(uint8 wordsPrefix, uint8 wordsFollowing)
        public
    {
        // Set a sensible limit for the total payload length
        vm.assume((uint256(wordsPrefix) + wordsFollowing) * 32 <= MAX_MESSAGE_BODY_BYTES);
        // Create "random" arguments and new/old prefix with different random seeds
        bytes memory prefixOld = createTestArguments(wordsPrefix, "prefixOld");
        bytes memory following = createTestArguments(wordsFollowing, "following");
        bytes memory prefixNew = createTestArguments(wordsPrefix, "prefixNew");
        bytes memory callData = bytes.concat(selector, prefixOld, following);
        bytes memory adjustedCallData = SystemMessageLib.formatAdjustedCallData(
            callData.castToCallData(),
            prefixNew.castToRawBytes()
        );
        // Correct formatting is checked in SystemMessage.t.sol
        // Test formatting checker
        assertTrue(libHarness.isCallData(adjustedCallData), "!isCallData");
        assertEq(libHarness.castToCallData(adjustedCallData), adjustedCallData, "!castToCallData");
        // Test CallData getters
        assertEq(
            libHarness.argumentWords(adjustedCallData),
            uint256(wordsPrefix) + wordsFollowing,
            "!argumentWords"
        );
        assertEq(
            libHarness.callSelector(adjustedCallData),
            bytes.concat(selector),
            "!callSelector"
        );
        assertEq(
            libHarness.arguments(adjustedCallData),
            bytes.concat(prefixNew, following),
            "!arguments"
        );
    }

    function test_formattedCorrectly_signature() public {
        bytes memory signature = signMessage({ privKey: 1, message: "" });
        // Test related constants
        assertEq(libHarness.signatureLength(), signature.length, "!signatureLength");
        // Test formatting checker
        checkCastToSignature({ payload: signature, isSignature: true });
        (bytes32 r, bytes32 s, uint8 v) = libHarness.toRSV(signature);
        assertEq(abi.encodePacked(r, s, v), signature, "!toRSV");
        assertEq(libHarness.formatSignature(r, s, v), signature, "!formatSignature");
    }

    function test_formattedCorrectly_rawBytes() public {
        bytes memory payload = "test payload";
        assertEq(libHarness.castToRawBytes(payload), payload, "!castToRawBytes");
    }

    function test_castToSignature_incorrectLength(uint16 length) public {
        vm.assume(length != libHarness.signatureLength());
        bytes memory payload = new bytes(length);
        checkCastToSignature({ payload: payload, isSignature: false });
    }

    function test_isCallData_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than CallData's first element (bytes4 selector)
        // should be correctly treated as unformatted (i.e. with no reverts)
        bytes memory payload = createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data);
        assertFalse(libHarness.isCallData(payload), "!isCallData: short payload");
        vm.expectRevert("Not a calldata");
        libHarness.castToCallData(payload);
    }

    function test_isCallData_noArgs() public {
        checkCastToCallData(abi.encodeWithSelector(selector), "!isCallData: no arguments");
    }

    function test_isCallData_withArgs() public {
        checkCastToCallData(
            abi.encodeWithSelector(selector, uint16(42), uint128(4815162342)),
            "!isCallData: with arguments"
        );
    }

    function test_isCallData_dynamicArgs() public {
        checkCastToCallData(
            abi.encodeWithSelector(selector, new bytes(13)),
            "!isCallData: bytes(13)"
        );
        checkCastToCallData(
            abi.encodeWithSelector(selector, new bytes(13), new bytes(42)),
            "!isCallData: bytes(13), bytes(42)"
        );
        checkCastToCallData(
            abi.encodeWithSelector(selector, new uint8[](2)),
            "!isCallData: uint8[](2)"
        );
        uint8[2] memory arg;
        checkCastToCallData(abi.encodeWithSelector(selector, arg), "!isCallData: uint8[2]");
    }

    function test_isCallData_extraBytes(uint8 _extraBytes) public {
        vm.assume(_extraBytes != 0);
        extraBytes = _extraBytes;
        test_isCallData_noArgs();
        test_isCallData_withArgs();
        test_isCallData_dynamicArgs();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToSignature(bytes memory payload, bool isSignature) public {
        if (isSignature) {
            assertTrue(libHarness.isSignature(payload), "!isSignature: when valid");
            assertEq(libHarness.castToSignature(payload), payload, "!castToSignature: when valid");
        } else {
            assertFalse(libHarness.isSignature(payload), "!isSignature: when valid");
            vm.expectRevert("Not a signature");
            libHarness.castToSignature(payload);
        }
    }

    function checkCastToCallData(bytes memory payload, string memory revertMessage) public {
        // Add extra bytes to the call payload
        payload = bytes.concat(payload, new bytes(extraBytes));
        bool isCallData = extraBytes % 32 == 0;
        if (isCallData) {
            assertTrue(libHarness.isCallData(payload), "!isCallData: when valid");
            assertEq(libHarness.castToCallData(payload), payload, "!castToCallData: when valid");
        } else {
            assertFalse(libHarness.isCallData(payload), revertMessage);
            vm.expectRevert("Not a calldata");
            libHarness.castToCallData(payload);
        }
    }
}
