// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemViewLib} from "../../contracts/libs/memory/MemView.sol";
import {SynapseUtilities} from "./SynapseUtilities.t.sol";

// solhint-disable no-empty-blocks
abstract contract SynapseLibraryTest is SynapseUtilities {
    using MemViewLib for bytes;

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseLibraryTest() external {}

    function createShortPayload(uint8 payloadLength, uint8 firstElementLength, bytes32 data)
        public
        view
        returns (bytes memory)
    {
        payloadLength = payloadLength % firstElementLength;
        // 8 bytes should be enough
        bytes memory payload = abi.encodePacked(data, data, data, data, data, data, data, data);
        // Use first `payloadLength` bytes
        return payload.ref().slice({index_: 0, len_: payloadLength}).clone();
    }

    function cutLastByte(bytes memory payload) public view returns (bytes memory) {
        return payload.ref().slice({index_: 0, len_: payload.length - 1}).clone();
    }

    function addLastByte(bytes memory payload) public pure returns (bytes memory) {
        return bytes.concat(payload, bytes1(0));
    }
}
