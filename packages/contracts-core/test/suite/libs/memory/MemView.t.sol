// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    IndexedTooMuch,
    OccupiedMemory,
    PrecompileOutOfGas,
    UnallocatedMemory,
    ViewOverrun
} from "../../../../contracts/libs/Errors.sol";
import {SynapseLibraryTest} from "../../../utils/SynapseLibraryTest.t.sol";

import {MemView, MemViewLib, MemViewHarness} from "../../../harnesses/libs/memory/MemViewHarness.t.sol";

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

    function test_keccakSalted(bytes memory arr, bytes32 salt) public {
        (, bytes32 hash,) = libHarness.keccakSalted(arr, salt);
        assertEq(hash, keccak256(abi.encodePacked(salt, keccak256(arr))));
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

    // ════════════════════════════════════════════ SLICING MEMORY VIEW ════════════════════════════════════════════════

    function test_slice(bytes memory prefix, bytes memory data, bytes memory postfix) public {
        bytes memory arr = abi.encodePacked(prefix, data, postfix);
        (, bytes memory result,) = libHarness.slice(arr, prefix.length, data.length);
        checkEqual(result, data);
    }

    function test_sliceTwice(
        bytes memory prefixF,
        bytes memory prefixS,
        bytes memory data,
        bytes memory postfixF,
        bytes memory postfixS
    ) public {
        bytes memory arr = abi.encodePacked(prefixF, prefixS, data, postfixF, postfixS);
        // First slice (prefixS, data, postfixS) out of the full payload
        uint256 lenFirst = prefixS.length + data.length + postfixF.length;
        (, bytes memory result,) = libHarness.sliceTwice(arr, prefixF.length, lenFirst, prefixS.length, data.length);
        checkEqual(result, data);
    }

    function test_sliceFrom(bytes memory prefix, bytes memory data) public {
        bytes memory arr = abi.encodePacked(prefix, data);
        (, bytes memory result,) = libHarness.sliceFrom(arr, prefix.length);
        checkEqual(result, data);
    }

    function test_prefix(bytes memory prefix, bytes memory data) public {
        bytes memory arr = abi.encodePacked(prefix, data);
        (, bytes memory result,) = libHarness.prefix(arr, prefix.length);
        checkEqual(result, prefix);
    }

    function test_postfix(bytes memory data, bytes memory postfix) public {
        bytes memory arr = abi.encodePacked(data, postfix);
        (, bytes memory result,) = libHarness.postfix(arr, postfix.length);
        checkEqual(result, postfix);
    }

    // ═══════════════════════════════════════════════ SLICE & HASH ════════════════════════════════════════════════════

    function test_sliceKeccak(bytes memory prefix, bytes memory data, bytes memory postfix) public {
        bytes memory arr = abi.encodePacked(prefix, data, postfix);
        (, bytes32 result,) = libHarness.sliceKeccak(arr, prefix.length, data.length);
        assertEq(result, keccak256(data));
    }

    function test_sliceFromKeccak(bytes memory prefix, bytes memory data) public {
        bytes memory arr = abi.encodePacked(prefix, data);
        (, bytes32 result,) = libHarness.sliceFromKeccak(arr, prefix.length);
        assertEq(result, keccak256(data));
    }

    function test_prefixKeccak(bytes memory prefix, bytes memory data) public {
        bytes memory arr = abi.encodePacked(prefix, data);
        (, bytes32 result,) = libHarness.prefixKeccak(arr, prefix.length);
        assertEq(result, keccak256(prefix));
    }

    function test_postfixKeccak(bytes memory data, bytes memory postfix) public {
        bytes memory arr = abi.encodePacked(data, postfix);
        (, bytes32 result,) = libHarness.postfixKeccak(arr, postfix.length);
        assertEq(result, keccak256(postfix));
    }

    // ═══════════════════════════════════════════════ EXPECT ERRORS ═══════════════════════════════════════════════════

    function test_buildUnallocated(uint256 offset, uint256 words) public {
        words = bound(words, 1, 100);
        offset = bound(offset, 1, 10_000);
        // Should pass with zero offset
        libHarness.buildUnallocated(0, words);
        // Non-zero offset will point to unallocated memory
        vm.expectRevert(UnallocatedMemory.selector);
        libHarness.buildUnallocated(offset, words);
    }

    function test_index_revert_indexOutOfRange(bytes memory arr, uint256 index, uint256 bytes_) public {
        index = bound(index, arr.length, type(uint128).max);
        bytes_ = bound(bytes_, 1, 32);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.index(arr, index, bytes_);
    }

    function test_index_revert_endOutOfRange(bytes memory arr, uint256 index, uint256 bytes_) public {
        vm.assume(arr.length != 0);
        index = bound(index, arr.length >= 31 ? arr.length - 31 : 0, arr.length - 1);
        bytes_ = bound(bytes_, arr.length - index + 1, 32);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.index(arr, index, bytes_);
    }

    function test_index_revert_moreThan32Bytes(bytes memory arr, uint256 index, uint256 bytes_) public {
        vm.assume(arr.length != 0);
        index = bound(index, 0, arr.length - 1);
        bytes_ = bound(bytes_, 33, type(uint256).max);
        vm.expectRevert(IndexedTooMuch.selector);
        libHarness.index(arr, index, bytes_);
    }

    function test_indexUint_indexOutOfRange(bytes memory arr, uint256 index, uint256 bytes_) public {
        index = bound(index, arr.length, type(uint128).max);
        bytes_ = bound(bytes_, 1, 32);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.indexUint(arr, index, bytes_);
    }

    function test_indexUint_revert_endOutOfRange(bytes memory arr, uint256 index, uint256 bytes_) public {
        vm.assume(arr.length != 0);
        index = bound(index, arr.length >= 31 ? arr.length - 31 : 0, arr.length - 1);
        bytes_ = bound(bytes_, arr.length - index + 1, 32);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.indexUint(arr, index, bytes_);
    }

    function test_indexAddress_indexOutOfRange(bytes memory arr, uint256 index) public {
        index = bound(index, arr.length, type(uint128).max);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.indexAddress(arr, index);
    }

    function test_slice_revert_indexOutOfRange(bytes memory arr, uint256 index, uint256 len) public {
        index = bound(index, arr.length, type(uint128).max);
        len = bound(len, 1, type(uint128).max);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.slice(arr, index, len);
    }

    function test_slice_revert_endOutOfRange(bytes memory arr, uint256 index, uint256 len) public {
        vm.assume(arr.length != 0);
        index = bound(index, 0, arr.length - 1);
        len = bound(len, arr.length - index + 1, type(uint128).max);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.slice(arr, index, len);
    }

    function test_sliceFrom_revert_indexOutOfRange(bytes memory arr, uint256 index) public {
        index = bound(index, arr.length + 1, type(uint128).max);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.sliceFrom(arr, index);
    }

    function test_prefix_revert_lengthOutOfRange(bytes memory arr, uint256 len) public {
        len = bound(len, arr.length + 1, type(uint128).max);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.prefix(arr, len);
    }

    function test_postfix_revert_lengthOutOfRange(bytes memory arr, uint256 len) public {
        len = bound(len, arr.length + 1, type(uint128).max);
        vm.expectRevert(ViewOverrun.selector);
        libHarness.postfix(arr, len);
    }
}
