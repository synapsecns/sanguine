// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest} from "../../utils/SynapseLibraryTest.t.sol";

import {MemView, MemViewLib, MemViewHarness} from "../../harnesses/libs/MemViewHarness.t.sol";

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

    // ═══════════════════════════════════════════ INDEXING MEMORY VIEW ════════════════════════════════════════════════

    function test_index(bytes memory prefix, bytes32 data, bytes memory postfix, uint256 bytes_) public {
        bytes_ = bound(bytes_, 0, 32);
        bytes memory arr = abi.encodePacked(prefix, data, postfix);
        bytes32 expected = bytes32(0);
        for (uint256 i = 0; i < bytes_; ++i) {
            // When casting bytes1 to bytes32 it will populate the highest bits, so shift right here
            expected |= bytes32(data[i]) >> (i * 8);
        }
        (, bytes32 result,) = libHarness.index(arr, prefix.length, bytes_);
        assertEq(result, expected);
    }

    function test_indexUint(bytes memory prefix, uint256 data, bytes memory postfix, uint256 bytes_) public {
        bytes_ = bound(bytes_, 0, 32);
        bytes memory encodedUint = new bytes(bytes_);
        for (uint256 i = 0; i < bytes_; ++i) {
            uint256 shiftBytes = i * 8;
            // Copy bytes from lowest to highest
            uint256 byte_ = (data & (0xFF << shiftBytes)) >> shiftBytes;
            encodedUint[bytes_ - 1 - i] = bytes1(uint8(byte_));
        }
        bytes memory arr = abi.encodePacked(prefix, encodedUint, postfix);
        uint256 mask = bytes_ == 32 ? type(uint256).max : ((1 << (bytes_ * 8)) - 1);
        uint256 expected = data & mask;
        (, uint256 result,) = libHarness.indexUint(arr, prefix.length, bytes_);
        assertEq(result, expected);
    }

    function test_indexAddress(bytes memory prefix, address data, bytes memory postfix) public {
        bytes memory arr = abi.encodePacked(prefix, data, postfix);
        (, address result,) = libHarness.indexAddress(arr, prefix.length);
        assertEq(result, data);
    }
}
