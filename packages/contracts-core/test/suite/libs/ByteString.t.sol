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

    function setUp() public override {
        super.setUp();
        libHarness = new ByteStringHarness();
    }

    function test_selectorLength() public {
        assertEq(libHarness.selectorLength(), selector.length, "!selectorLength");
    }

    function test_signatureLength() public {
        bytes memory signature = signMessage({ privKey: 1, message: "" });
        assertEq(libHarness.signatureLength(), signature.length, "!signatureLength");
    }

    function test_isSignature() public {
        bytes memory signatureMock = new bytes(libHarness.signatureLength());
        assertTrue(libHarness.isSignature(signatureMock), "!isSignature: correct length");
    }

    function test_isSignature_incorrectLength(uint16 length) public {
        vm.assume(length != libHarness.signatureLength());
        assertFalse(libHarness.isSignature(new bytes(length)), "!isSignature: wrong length");
    }

    function test_isCallPayload_noArgs() public {
        _checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector),
            _revertMessage: "!isCallPayload: no arguments"
        });
    }

    function test_isCallPayload_withArgs() public {
        _checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, uint16(42), uint128(4815162342)),
            _revertMessage: "!isCallPayload: with arguments"
        });
    }

    function test_isCallPayload_dynamicArgs() public {
        _checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, new bytes(13)),
            _revertMessage: "!isCallPayload: bytes(13)"
        });
        _checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, new bytes(13), new bytes(42)),
            _revertMessage: "!isCallPayload: bytes(13), bytes(42)"
        });
        _checkIsCallPayload({
            _payload: abi.encodeWithSelector(selector, new uint8[](2)),
            _revertMessage: "!isCallPayload: uint8[](2)"
        });
        uint8[2] memory arg;
        _checkIsCallPayload({
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

    function _checkIsCallPayload(bytes memory _payload, string memory _revertMessage) internal {
        // Add extra bytes to the call payload
        bytes memory payloadMock = bytes.concat(_payload, new bytes(extraBytes));
        bool isCallPayload = extraBytes % 32 == 0;
        assertEq(libHarness.isCallPayload(payloadMock), isCallPayload, _revertMessage);
    }
}
