// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemView, MemView, MemViewLib} from "../../../../contracts/libs/memory/MemView.sol";

// solhint-disable ordering
contract MemViewHarness {
    using MemViewLib for bytes;

    /// @notice Tries to build a memory view that touches the unallocated memory
    function buildUnallocated(uint256 offset, uint256 words_)
        external
        view
        returns (bytes memory, bytes memory, bytes memory)
    {
        uint256 loc_;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            loc_ := mload(0x40)
        }
        bytes memory allocated = new bytes(32 * words_);
        bytes memory result = MemViewLib.build(loc_ + offset, 32 * (words_ + 1)).clone();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirty = hex"FedFed";
        return (allocated, result, dirty);
    }

    // ════════════════════════════════════════════ CLONING MEMORY VIEW ════════════════════════════════════════════════

    function clone(bytes memory arr) external view returns (bytes memory, bytes memory, bytes memory) {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        bytes memory result = arr.ref().clone();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function join(bytes[] memory arrays) external view returns (bytes memory, bytes memory, bytes memory) {
        MemView[] memory views = new MemView[](arrays.length);
        for (uint256 i = 0; i < arrays.length; ++i) {
            views[i] = arrays[i].ref();
        }
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        bytes memory result = MemViewLib.join(views);
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    // ══════════════════════════════════════════ INSPECTING MEMORY VIEW ═══════════════════════════════════════════════

    function unsafeBuildUnchecked(uint256 loc_, uint256 len_) external pure returns (MemView) {
        return MemView.wrap((loc_ << 128) | len_);
    }

    function loc(MemView memView) external pure returns (uint256 loc_) {
        return memView.loc();
    }

    function len(MemView memView) external pure returns (uint256 len_) {
        return memView.len();
    }

    function end(MemView memView) external pure returns (uint256 end_) {
        return memView.end();
    }

    function words(MemView memView) external pure returns (uint256 words_) {
        return memView.words();
    }

    function footprint(MemView memView) external pure returns (uint256 footprint_) {
        return memView.footprint();
    }

    // ════════════════════════════════════════════ HASHING MEMORY VIEW ════════════════════════════════════════════════

    function keccak(bytes memory arr) external pure returns (bytes memory, bytes32, bytes memory) {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes32 result = memView.keccak();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function keccakSalted(bytes memory arr, bytes32 salt) external pure returns (bytes memory, bytes32, bytes memory) {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes32 result = memView.keccakSalted(salt);
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    // ═══════════════════════════════════════════ INDEXING MEMORY VIEW ════════════════════════════════════════════════

    function index(bytes memory arr, uint256 index_, uint256 bytes_)
        external
        pure
        returns (bytes memory, bytes32, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        bytes32 result = memView.index(index_, bytes_);
        return (dirtyBefore, result, dirtyAfter);
    }

    function indexUint(bytes memory arr, uint256 index_, uint256 bytes_)
        external
        pure
        returns (bytes memory, uint256, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        uint256 result = memView.indexUint(index_, bytes_);
        return (dirtyBefore, result, dirtyAfter);
    }

    function indexAddress(bytes memory arr, uint256 index_)
        external
        pure
        returns (bytes memory, address, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        address result = memView.indexAddress(index_);
        return (dirtyBefore, result, dirtyAfter);
    }

    // ════════════════════════════════════════════ SLICING MEMORY VIEW ════════════════════════════════════════════════

    function slice(bytes memory arr, uint256 index_, uint256 len_)
        external
        view
        returns (bytes memory, bytes memory, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes memory result = memView.slice(index_, len_).clone();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function sliceTwice(bytes memory arr, uint256 indexFirst, uint256 lenFirst, uint256 indexSecond, uint256 lenSecond)
        external
        view
        returns (bytes memory, bytes memory, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes memory result = memView.slice(indexFirst, lenFirst).slice(indexSecond, lenSecond).clone();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function sliceFrom(bytes memory arr, uint256 index_)
        external
        view
        returns (bytes memory, bytes memory, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes memory result = memView.sliceFrom(index_).clone();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function prefix(bytes memory arr, uint256 len_) external view returns (bytes memory, bytes memory, bytes memory) {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes memory result = memView.prefix(len_).clone();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function postfix(bytes memory arr, uint256 len_) external view returns (bytes memory, bytes memory, bytes memory) {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes memory result = memView.postfix(len_).clone();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    // ═══════════════════════════════════════════════ SLICE & HASH ════════════════════════════════════════════════════

    function sliceKeccak(bytes memory arr, uint256 index_, uint256 len_)
        external
        pure
        returns (bytes memory, bytes32, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes32 result = memView.slice(index_, len_).keccak();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function sliceFromKeccak(bytes memory arr, uint256 index_)
        external
        pure
        returns (bytes memory, bytes32, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes32 result = memView.sliceFrom(index_).keccak();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function prefixKeccak(bytes memory arr, uint256 len_) external pure returns (bytes memory, bytes32, bytes memory) {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes32 result = memView.prefix(len_).keccak();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }

    function postfixKeccak(bytes memory arr, uint256 len_)
        external
        pure
        returns (bytes memory, bytes32, bytes memory)
    {
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyBefore = hex"DeadDead";
        MemView memView = arr.ref();
        bytes32 result = memView.postfix(len_).keccak();
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        return (dirtyBefore, result, dirtyAfter);
    }
}
