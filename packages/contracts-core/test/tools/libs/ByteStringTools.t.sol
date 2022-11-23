// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract ByteStringTools {
    function createTestArguments(uint8 words, bytes memory seed)
        public
        pure
        returns (bytes memory)
    {
        bytes32[] memory arguments = new bytes32[](words);
        bytes32 randomData = keccak256(seed);
        for (uint256 i = 0; i < words; ++i) {
            arguments[i] = randomData;
            randomData = keccak256(abi.encode(randomData));
        }
        return abi.encodePacked(arguments);
    }
}
