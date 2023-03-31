// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TypedMemView } from "./TypedMemView.sol";

library TypeCasts {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    function coerceBytes32(string memory s) internal pure returns (bytes32 b) {
        b = bytes(s).ref(0).index(0, uint8(bytes(s).length));
    }

    // treat it as a null-terminated string of max 32 bytes
    function coerceString(bytes32 buf) internal pure returns (string memory newStr) {
        uint8 sLen = 0;
        while (sLen < 32 && buf[sLen] != 0) {
            sLen++;
        }

        // solhint-disable-next-line no-inline-assembly
        assembly {
            newStr := mload(0x40)
            mstore(0x40, add(newStr, 0x40)) // may end up with extra
            mstore(newStr, sLen)
            mstore(add(newStr, 0x20), buf)
        }
    }

    // alignment preserving cast
    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    // alignment preserving cast
    function bytes32ToAddress(bytes32 buf) internal pure returns (address) {
        return address(uint160(uint256(buf)));
    }
}
