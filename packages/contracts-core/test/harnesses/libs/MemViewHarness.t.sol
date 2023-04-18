// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemView, MemViewLib} from "../../../contracts/libs/MemView.sol";

contract MemViewHarness {
    using MemViewLib for bytes;

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
        // Add some dirty data to where the current free memory pointer is pointing
        bytes memory dirtyAfter = hex"FedFed";
        bytes32 result = memView.keccak();
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
}
