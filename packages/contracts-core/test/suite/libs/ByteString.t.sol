// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/ByteStringHarness.t.sol";

import "../../../contracts/libs/ByteString.sol";

// solhint-disable func-name-mixedcase
contract ByteStringLibraryTest is SynapseLibraryTest {
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
    ▏*║                           TESTS: CONSTANTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_selectorLength() public {
        assertEq(libHarness.selectorLength(), selector.length, "!selectorLength");
    }

    function test_signatureLength() public {
        bytes memory signature = signMessage({ privKey: 1, message: "" });
        assertEq(libHarness.signatureLength(), signature.length, "!signatureLength");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_callPayloadFormattedCorrectly(uint8 words) public {
        bytes memory data = createTestPayload(words);
        require(data.length == selector.length + 32 * uint256(words), "!length");
        // Test formatting checker
        assertTrue(libHarness.isCallPayload(data), "!isCallPayload");
        // Test getters
        assertEq(
            libHarness.argumentWords({ _type: SynapseTypes.CALL_PAYLOAD, _payload: data }),
            words,
            "!argumentWords"
        );
    }

    function test_isSignature() public {
        bytes memory signatureMock = new bytes(libHarness.signatureLength());
        assertTrue(libHarness.isSignature(signatureMock), "!isSignature: correct length");
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
        bytes memory payload = createTestPayload(0);
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.CALL_PAYLOAD });
        libHarness.argumentWords(wrongType, payload);
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

    function createTestPayload(uint8 words) public view returns (bytes memory) {
        return abi.encodePacked(selector, createTestArguments(words));
    }

    function createTestArguments(uint8 words) public pure returns (bytes memory) {
        bytes32[] memory arguments = new bytes32[](words);
        bytes32 randomData = keccak256("very random seed");
        for (uint256 i = 0; i < words; ++i) {
            arguments[i] = randomData;
            randomData = keccak256(abi.encode(randomData));
        }
        return abi.encodePacked(arguments);
    }
}
