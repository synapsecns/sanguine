// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {ReadableSymbol} from "../../src/libs/ReadableSymbol.sol";

contract ReadableSymbolHarness {
    function toBytes31(string memory str) public pure returns (bytes31) {
        return ReadableSymbol.toBytes31(str);
    }

    function toString(bytes31 symbol) public pure returns (string memory) {
        return ReadableSymbol.toString(symbol);
    }
}
