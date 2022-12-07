// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/ByteStringHarness.t.sol";
import "../../tools/libs/ByteStringTools.t.sol";

import "../../../contracts/libs/ByteString.sol";

// solhint-disable func-name-mixedcase
contract ByteStringLibraryTest is ByteStringTools, SynapseLibraryTest {
    using TypedMemView for bytes;

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

    function test_formattedCorrectly_callPayload(uint8 words) public {
        // Set a sensible limit for the total payload length
        vm.assume(uint256(words) * 32 <= MAX_MESSAGE_BODY_BYTES);
        bytes memory arguments = createTestArguments(words, "seed");
        bytes memory data = abi.encodePacked(selector, arguments);
        require(data.length == selector.length + 32 * uint256(words), "!length");
        // Test related constants
        assertEq(libHarness.selectorLength(), selector.length, "!selectorLength");
        // Test formatting checker
        assertTrue(libHarness.isCallPayload(data), "!isCallPayload");
        // Test getters
        assertEq(
            libHarness.argumentWords({ _type: SynapseTypes.CALL_PAYLOAD, _payload: data }),
            words,
            "!argumentWords"
        );
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToCallPayload,
            payloadType: SynapseTypes.CALL_PAYLOAD,
            payload: data,
            expectedType: SynapseTypes.CALL_PAYLOAD,
            expectedData: data,
            revertMessage: "!castToCallPayload"
        });
        checkBytes29Getter({
            getter: libHarness.callSelector,
            payloadType: SynapseTypes.CALL_PAYLOAD,
            payload: data,
            expectedType: SynapseTypes.RAW_BYTES,
            expectedData: bytes.concat(selector),
            revertMessage: "!callSelector"
        });
        checkBytes29Getter({
            getter: libHarness.argumentsPayload,
            payloadType: SynapseTypes.CALL_PAYLOAD,
            payload: data,
            expectedType: SynapseTypes.RAW_BYTES,
            expectedData: arguments,
            revertMessage: "!argumentsPayload"
        });
    }

    function test_formattedCorrectly_callPayload_adjusted(uint8 wordsPrefix, uint8 wordsFollowing)
        public
    {
        // Set a sensible limit for the total payload length
        vm.assume((uint256(wordsPrefix) + wordsFollowing) * 32 <= MAX_MESSAGE_BODY_BYTES);
        // Create "random" arguments and new/old prefix with different random seeds
        bytes memory prefixOld = createTestArguments(wordsPrefix, "prefixOld");
        bytes memory following = createTestArguments(wordsFollowing, "following");
        bytes memory prefixNew = createTestArguments(wordsPrefix, "prefixNew");
        bytes memory payload = bytes.concat(selector, prefixOld, following);
        bytes memory adjustedPayload = SystemCall.formatAdjustedCallPayload(
            payload.ref(SynapseTypes.CALL_PAYLOAD),
            prefixNew.ref(SynapseTypes.RAW_BYTES)
        );
        // Correct formatting is checked in SystemCall.t.sol
        // Test formatting checker
        assertTrue(libHarness.isCallPayload(adjustedPayload), "!isCallPayload");
        // Test ByteString getters
        // Test getters
        assertEq(
            libHarness.argumentWords({
                _type: SynapseTypes.CALL_PAYLOAD,
                _payload: adjustedPayload
            }),
            uint256(wordsPrefix) + wordsFollowing,
            "!argumentWords"
        );
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToCallPayload,
            payloadType: SynapseTypes.CALL_PAYLOAD,
            payload: adjustedPayload,
            expectedType: SynapseTypes.CALL_PAYLOAD,
            expectedData: adjustedPayload,
            revertMessage: "!castToCallPayload"
        });
        checkBytes29Getter({
            getter: libHarness.callSelector,
            payloadType: SynapseTypes.CALL_PAYLOAD,
            payload: adjustedPayload,
            expectedType: SynapseTypes.RAW_BYTES,
            expectedData: bytes.concat(selector),
            revertMessage: "!callSelector"
        });
        checkBytes29Getter({
            getter: libHarness.argumentsPayload,
            payloadType: SynapseTypes.CALL_PAYLOAD,
            payload: adjustedPayload,
            expectedType: SynapseTypes.RAW_BYTES,
            expectedData: bytes.concat(prefixNew, following),
            revertMessage: "!argumentsPayload"
        });
    }

    function test_formattedCorrectly_signature() public {
        bytes memory signature = signMessage({ privKey: 1, message: "" });
        // Test related constants
        assertEq(libHarness.signatureLength(), signature.length, "!signatureLength");
        // Test formatting checker
        assertTrue(libHarness.isSignature(signature), "!isSignature");
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToSignature,
            payloadType: SynapseTypes.SIGNATURE,
            payload: signature,
            expectedType: SynapseTypes.SIGNATURE,
            expectedData: signature,
            revertMessage: "!castToSignature"
        });
        (bytes32 r, bytes32 s, uint8 v) = libHarness.toRSV({
            _type: SynapseTypes.SIGNATURE,
            _payload: signature
        });
        assertEq(abi.encodePacked(r, s, v), signature, "!toRSV");
    }

    function test_formattedCorrectly_rawBytes() public {
        bytes memory payload = "test payload";
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToRawBytes,
            payloadType: SynapseTypes.RAW_BYTES,
            payload: payload,
            expectedType: SynapseTypes.RAW_BYTES,
            expectedData: payload,
            revertMessage: "!castToRawBytes"
        });
    }

    function test_isSignature_incorrectLength(uint16 length) public {
        vm.assume(length != libHarness.signatureLength());
        assertFalse(libHarness.isSignature(new bytes(length)), "!isSignature: wrong length");
    }

    function test_isCallPayload_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than CallPayload's first element (bytes4 selector)
        // should be correctly treated as unformatted (i.e. with no reverts)
        assertFalse(
            libHarness.isCallPayload(createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data)),
            "!isCallPayload: short payload"
        );
    }

    function test_isCallPayload_noArgs() public {
        checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector),
            _revertMessage: "!isCallPayload: no arguments"
        });
    }

    function test_isCallPayload_withArgs() public {
        checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, uint16(42), uint128(4815162342)),
            _revertMessage: "!isCallPayload: with arguments"
        });
    }

    function test_isCallPayload_dynamicArgs() public {
        checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, new bytes(13)),
            _revertMessage: "!isCallPayload: bytes(13)"
        });
        checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, new bytes(13), new bytes(42)),
            _revertMessage: "!isCallPayload: bytes(13), bytes(42)"
        });
        checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, new uint8[](2)),
            _revertMessage: "!isCallPayload: uint8[](2)"
        });
        uint8[2] memory arg;
        checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, arg),
            _revertMessage: "!isCallPayload: uint8[2]"
        });
    }

    function test_isCallPayload_extraBytes(uint8 _extraBytes) public {
        vm.assume(_extraBytes != 0);
        extraBytes = _extraBytes;
        test_isCallPayload_noArgs();
        test_isCallPayload_withArgs();
        test_isCallPayload_dynamicArgs();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_wrongTypeRevert_argumentWords(uint40 wrongType) public {
        bytes memory payload = bytes.concat(selector);
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.CALL_PAYLOAD });
        libHarness.argumentWords(wrongType, payload);
    }

    function test_wrongTypeRevert_callSelector(uint40 wrongType) public {
        bytes memory payload = bytes.concat(selector);
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.CALL_PAYLOAD });
        libHarness.callSelector(wrongType, payload);
    }

    function test_wrongTypeRevert_argumentsPayload(uint40 wrongType) public {
        bytes memory payload = bytes.concat(selector);
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.CALL_PAYLOAD });
        libHarness.argumentsPayload(wrongType, payload);
    }

    function test_wrongTypeRevert_toRSV(uint40 wrongType) public {
        bytes memory payload = new bytes(ByteString.SIGNATURE_LENGTH);
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.SIGNATURE });
        libHarness.toRSV(wrongType, payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkIsCallPayload(bytes memory _payload, string memory _revertMessage) public {
        // Add extra bytes to the call payload
        bytes memory payloadMock = bytes.concat(_payload, new bytes(extraBytes));
        bool isCallPayload = extraBytes % 32 == 0;
        assertEq(libHarness.isCallPayload(payloadMock), isCallPayload, _revertMessage);
    }
}
