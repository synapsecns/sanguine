// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest} from "../../utils/SynapseLibraryTest.t.sol";

import {MemView, MemViewHarness} from "../../harnesses/libs/MemViewHarness.t.sol";

// solhint-disable func-name-mixedcase
contract MemViewLibraryTest is SynapseLibraryTest {
    MemViewHarness public libHarness;

    function setUp() public {
        libHarness = new MemViewHarness();
    }

    // ════════════════════════════════════════════ CLONING MEMORY VIEW ════════════════════════════════════════════════

    function test_clone(bytes memory arr) public {
        (, bytes memory cloned,) = libHarness.clone(arr);
        checkEqual(cloned, arr);
    }

    function test_join(bytes[] memory arrays) public {
        (, bytes memory cloned,) = libHarness.join(arrays);
        bytes memory expected = "";
        for (uint256 i = 0; i < arrays.length; ++i) {
            expected = bytes.concat(expected, arrays[i]);
        }
        checkEqual(cloned, expected);
    }

    function checkEqual(bytes memory arr, bytes memory expected) public {
        assertEq(arr.length, expected.length, "!length");
        assertEq(arr, expected, "!data");
    }

    // ══════════════════════════════════════════ INSPECTING MEMORY VIEW ═══════════════════════════════════════════════

    function test_memView(uint256 loc, uint256 len) public {
        // Make sure loc and len fit in uint128
        loc = bound(loc, 0, type(uint128).max);
        len = bound(len, 0, type(uint128).max);
        MemView memView = libHarness.unsafeBuildUnchecked(loc, len);
        assertEq(libHarness.loc(memView), loc, "!loc");
        assertEq(libHarness.len(memView), len, "!len");
        assertEq(libHarness.end(memView), loc + len, "!end");
        uint256 words = len / 32;
        if (len % 32 != 0) ++words;
        assertEq(libHarness.words(memView), words, "!words");
        assertEq(libHarness.footprint(memView), 32 * words, "!footprint");
    }

    // ════════════════════════════════════════════ HASHING MEMORY VIEW ════════════════════════════════════════════════

    function test_keccak(bytes memory arr) public {
        (, bytes32 hash,) = libHarness.keccak(arr);
        assertEq(hash, keccak256(arr));
    }
}
