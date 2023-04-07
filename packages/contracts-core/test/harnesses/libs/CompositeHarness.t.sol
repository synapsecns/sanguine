// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Composite} from "../../../contracts/libs/Composite.sol";

contract CompositeHarness {
    function mergeUint32(uint32 first, uint32 second) public pure returns (uint64 combined) {
        return Composite.mergeUint32(first, second);
    }

    function splitUint32(uint64 combined) public pure returns (uint32 first, uint32 second) {
        return Composite.splitUint32(combined);
    }
}
