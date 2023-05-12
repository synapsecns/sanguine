// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {UnformattedCallData, UnformattedSignature} from "../../../../contracts/libs/Errors.sol";
import {SynapseLibraryTest} from "../../../utils/SynapseLibraryTest.t.sol";
import {ByteStringHarness} from "../../../harnesses/libs/memory/ByteStringHarness.t.sol";

import {ByteString} from "../../../../contracts/libs/memory/ByteString.sol";

import {Random} from "../../../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
contract ByteStringLibraryTest is SynapseLibraryTest {
    using ByteString for bytes;

    ByteStringHarness internal libHarness;
    uint256 internal extraBytes = 0;

    function setUp() public {
        libHarness = new ByteStringHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formattedCorrectly_callData(bytes4 selector, uint256 words) public {
        // Set a sensible limit for the total payload length
        words = words % MAX_SYSTEM_CALL_WORDS;
        bytes memory arguments = Random("random").nextBytesWords(words);
        bytes memory callData = abi.encodePacked(selector, arguments);
        require(callData.length == selector.length + 32 * uint256(words), "!length");
        // Test related constants
        assertEq(libHarness.selectorLength(), selector.length, "!selectorLength");
        // Test formatting checker
        assertTrue(libHarness.isCallData(callData), "!isCallData");
        assertEq(libHarness.castToCallData(callData), callData, "!castToCallData");
        // Test CallData getters
        assertEq(libHarness.argumentWords(callData), words, "!argumentWords");
        assertEq(libHarness.callSelector(callData), selector, "!callSelector");
        assertEq(libHarness.arguments(callData), arguments, "!arguments");
        // Test hashing
        assertEq(libHarness.leaf(callData), keccak256(callData), "!leaf");
    }

    function test_formattedCorrectly_callData_added(bytes4 selector, uint256 wordsPrefix, uint256 wordsFollowing)
        public
    {
        // Set a sensible limit for the total payload length
        wordsPrefix = wordsPrefix % MAX_SYSTEM_CALL_WORDS;
        wordsFollowing = bound(wordsFollowing, 0, MAX_SYSTEM_CALL_WORDS - wordsPrefix);
        // Create "random" arguments and new/old prefix with different random seeds
        bytes memory prefix = Random("prefix").nextBytesWords(wordsPrefix);
        bytes memory following = Random("following").nextBytesWords(wordsFollowing);
        // Initial calldata
        bytes memory callData = bytes.concat(selector, following);
        bytes memory finalCallData = libHarness.addPrefix(callData, prefix);
        // Test formatting checker
        assertTrue(libHarness.isCallData(finalCallData), "!isCallData");
        assertEq(libHarness.castToCallData(finalCallData), finalCallData, "!castToCallData");
        // Test CallData getters
        assertEq(libHarness.argumentWords(finalCallData), wordsPrefix + wordsFollowing, "!argumentWords");
        assertEq(libHarness.callSelector(finalCallData), selector, "!callSelector");
        assertEq(libHarness.arguments(finalCallData), bytes.concat(prefix, following), "!arguments");
    }

    function test_formattedCorrectly_signature() public {
        bytes memory signature = signMessage({privKey: 1, message: ""});
        // Test related constants
        assertEq(libHarness.signatureLength(), signature.length, "!signatureLength");
        // Test formatting checker
        checkCastToSignature({payload: signature, isSignature: true});
        (bytes32 r, bytes32 s, uint8 v) = libHarness.toRSV(signature);
        assertEq(abi.encodePacked(r, s, v), signature, "!toRSV");
        assertEq(libHarness.formatSignature(r, s, v), signature, "!formatSignature");
    }

    function test_castToSignature_incorrectLength(uint16 length) public {
        vm.assume(length != libHarness.signatureLength());
        bytes memory payload = new bytes(length);
        checkCastToSignature({payload: payload, isSignature: false});
    }

    function test_isCallData(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToCallData(payload, "Empty payload");
    }

    function test_isCallData_noArgs(bytes4 selector) public {
        checkCastToCallData(abi.encodeWithSelector(selector), "!isCallData: no arguments");
    }

    function test_isCallData_withArgs(bytes4 selector) public {
        checkCastToCallData(
            abi.encodeWithSelector(selector, uint16(42), uint128(4_815_162_342)), "!isCallData: with arguments"
        );
    }

    function test_isCallData_dynamicArgs(bytes4 selector) public {
        checkCastToCallData(abi.encodeWithSelector(selector, new bytes(13)), "!isCallData: bytes(13)");
        checkCastToCallData(
            abi.encodeWithSelector(selector, new bytes(13), new bytes(42)), "!isCallData: bytes(13), bytes(42)"
        );
        checkCastToCallData(abi.encodeWithSelector(selector, new uint8[](2)), "!isCallData: uint8[](2)");
        uint8[2] memory arg;
        checkCastToCallData(abi.encodeWithSelector(selector, arg), "!isCallData: uint8[2]");
    }

    function test_isCallData_extraBytes(bytes4 selector, uint8 extraBytes_) public {
        vm.assume(extraBytes_ != 0);
        extraBytes = extraBytes_;
        test_isCallData_noArgs(selector);
        test_isCallData_withArgs(selector);
        test_isCallData_dynamicArgs(selector);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToSignature(bytes memory payload, bool isSignature) public {
        if (isSignature) {
            assertTrue(libHarness.isSignature(payload), "!isSignature: when valid");
            assertEq(libHarness.castToSignature(payload), payload, "!castToSignature: when valid");
        } else {
            assertFalse(libHarness.isSignature(payload), "!isSignature: when valid");
            vm.expectRevert(UnformattedSignature.selector);
            libHarness.castToSignature(payload);
        }
    }

    function checkCastToCallData(bytes memory payload, string memory revertMessage) public {
        // Add extra bytes to the call payload
        payload = bytes.concat(payload, new bytes(extraBytes));
        bool isCallData = payload.length % 32 == 4;
        if (isCallData) {
            assertTrue(libHarness.isCallData(payload), "!isCallData: when valid");
            assertEq(libHarness.castToCallData(payload), payload, "!castToCallData: when valid");
        } else {
            assertFalse(libHarness.isCallData(payload), revertMessage);
            vm.expectRevert(UnformattedCallData.selector);
            libHarness.castToCallData(payload);
        }
    }
}
