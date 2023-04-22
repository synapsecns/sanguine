// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// solhint-disable func-visibility
// Collection of free functions for the tests

function addressToBytes32(address addr) pure returns (bytes32) {
    return bytes32(uint256(uint160(addr)));
}

function bytes32ToAddress(bytes32 buf) pure returns (address) {
    return address(uint160(uint256(buf)));
}
