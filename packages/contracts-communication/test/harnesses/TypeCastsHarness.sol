// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {TypeCasts} from "../../contracts/libs/TypeCasts.sol";

contract TypeCastsHarness {
    function addressToBytes32(address addr) external pure returns (bytes32) {
        return TypeCasts.addressToBytes32(addr);
    }

    function bytes32ToAddress(bytes32 b) external pure returns (address) {
        return TypeCasts.bytes32ToAddress(b);
    }
}
