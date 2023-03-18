// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";
import { SynapseTestSuite } from "../utils/SynapseTestSuite.t.sol";

abstract contract SynapseLibraryTest is SynapseTestSuite {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

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
