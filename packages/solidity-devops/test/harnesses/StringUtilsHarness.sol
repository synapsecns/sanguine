// SPDX-License-Identifier: MIT
pragma solidity >=0.6.2 <0.9.0;

import {StringUtils} from "../../src/libs/StringUtils.sol";

contract StringUtilsHarness {
    function length(string memory str) external pure returns (uint256) {
        uint256 result = StringUtils.length(str);
        return result;
    }

    // ══════════════════════════════════════════════════ SLICING ══════════════════════════════════════════════════════

    function substring(string memory str, uint256 startIndex, uint256 endIndex) external pure returns (string memory) {
        string memory result = StringUtils.substring(str, startIndex, endIndex);
        return result;
    }

    function suffix(string memory str, uint256 startIndex) external pure returns (string memory) {
        string memory result = StringUtils.suffix(str, startIndex);
        return result;
    }

    function prefix(string memory str, uint256 endIndex) external pure returns (string memory) {
        string memory result = StringUtils.prefix(str, endIndex);
        return result;
    }

    // ═══════════════════════════════════════════════ CONCATENATION ═══════════════════════════════════════════════════

    function concat(string memory a, string memory b) external pure returns (string memory) {
        string memory result = StringUtils.concat(a, b);
        return result;
    }

    function concat(string memory a, string memory b, string memory c) external pure returns (string memory) {
        string memory result = StringUtils.concat(a, b, c);
        return result;
    }

    function concat(
        string memory a,
        string memory b,
        string memory c,
        string memory d
    )
        external
        pure
        returns (string memory)
    {
        string memory result = StringUtils.concat(a, b, c, d);
        return result;
    }

    function concat(
        string memory a,
        string memory b,
        string memory c,
        string memory d,
        string memory e
    )
        external
        pure
        returns (string memory)
    {
        string memory result = StringUtils.concat(a, b, c, d, e);
        return result;
    }

    function concat(
        string memory a,
        string memory b,
        string memory c,
        string memory d,
        string memory e,
        string memory f
    )
        external
        pure
        returns (string memory)
    {
        string memory result = StringUtils.concat(a, b, c, d, e, f);
        return result;
    }

    function duplicate(string memory str, uint256 times) external pure returns (string memory) {
        string memory result = StringUtils.duplicate(str, times);
        return result;
    }

    // ════════════════════════════════════════════════ COMPARISON ═════════════════════════════════════════════════════

    function equals(string memory a, string memory b) external pure returns (bool) {
        bool result = StringUtils.equals(a, b);
        return result;
    }

    function indexOf(string memory str, string memory subStr) external pure returns (uint256) {
        uint256 result = StringUtils.indexOf(str, subStr);
        return result;
    }

    function lastIndexOf(string memory str, string memory subStr) external pure returns (uint256) {
        uint256 result = StringUtils.lastIndexOf(str, subStr);
        return result;
    }

    // ════════════════════════════════════════════ INTEGER CONVERSION ═════════════════════════════════════════════════

    function toUint(string memory str) external pure returns (uint256) {
        uint256 result = StringUtils.toUint(str);
        return result;
    }

    function fromUint(uint256 val) external pure returns (string memory) {
        string memory result = StringUtils.fromUint(val);
        return result;
    }

    // ═════════════════════════════════════════════ FLOAT CONVERSION ══════════════════════════════════════════════════

    function fromFloat(uint256 val, uint256 decimals) external pure returns (string memory) {
        string memory result = StringUtils.fromFloat(val, decimals);
        return result;
    }

    function fromWei(uint256 val) external pure returns (string memory) {
        string memory result = StringUtils.fromWei(val);
        return result;
    }
}
