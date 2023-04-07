// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {CompositeHarness} from "../../harnesses/libs/CompositeHarness.t.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract CompositeLibraryTest is Test {
    CompositeHarness internal libHarness;

    function setUp() public {
        libHarness = new CompositeHarness();
    }

    function test_mergeUint32(uint32 first, uint32 second) public {
        uint64 combined = libHarness.mergeUint32(first, second);
        assertEq(combined, mergeDumb(first, second));
        (uint32 first_, uint32 second_) = libHarness.splitUint32(combined);
        assertEq(first_, first);
        assertEq(second_, second);
    }

    function test_splitUint32(uint64 combined) public {
        (uint32 first, uint32 second) = libHarness.splitUint32(combined);
        assertEq(mergeDumb(first, second), combined);
        assertEq(libHarness.mergeUint32(first, second), combined);
    }

    function mergeDumb(uint32 first, uint32 second) public pure returns (uint256) {
        return uint256(first) * 2 ** 32 + second;
    }
}
