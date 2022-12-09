// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../utils/SynapseTestSuite.t.sol";
import "../../contracts/libs/SynapseTypes.sol";

abstract contract SynapseLibraryTest is SynapseTestSuite {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    function checkBytes29Getter(
        function(uint40, bytes memory) external view returns (uint40, bytes memory) getter,
        uint40 payloadType,
        bytes memory payload,
        uint40 expectedType,
        bytes memory expectedData,
        string memory revertMessage
    ) public {
        (uint40 _type, bytes memory _data) = getter(payloadType, payload);
        assertEq(_data, expectedData, revertMessage);
        assertEq(_type, expectedType, string.concat(revertMessage, ": type"));
    }

    function expectRevertWrongType(uint40 wrongType, uint40 correctType) public {
        vm.assume(wrongType != correctType);
        (, uint256 g) = TypedMemView.encodeHex(wrongType);
        (, uint256 e) = TypedMemView.encodeHex(correctType);
        vm.expectRevert(
            abi.encodePacked("Type assertion failed. Got 0x", uint80(g), ". Expected 0x", uint80(e))
        );
    }

    function createShortPayload(
        uint8 payloadLength,
        uint8 firstElementLength,
        bytes32 data
    ) public view returns (bytes memory) {
        payloadLength = payloadLength % firstElementLength;
        // 8 bytes should be enough
        bytes memory payload = abi.encodePacked(data, data, data, data, data, data, data, data);
        // Use first `payloadLength` bytes
        return payload.ref(0).slice({ _index: 0, _len: payloadLength, newType: 0 }).clone();
    }

    function cutLastByte(bytes memory payload) public view returns (bytes memory) {
        return payload.ref(0).slice({ _index: 0, _len: payload.length - 1, newType: 0 }).clone();
    }

    function addLastByte(bytes memory payload) public pure returns (bytes memory) {
        return bytes.concat(payload, bytes1(0));
    }
}
