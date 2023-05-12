// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IndexedTooMuch, OccupiedMemory, PrecompileOutOfGas, UnallocatedMemory, ViewOverrun} from "../Errors.sol";

/// @dev MemView is an untyped view over a portion of memory to be used instead of `bytes memory`
type MemView is uint256;

/// @dev Attach library functions to MemView
using MemViewLib for MemView global;

/// @notice Library for operations with the memory views.
/// Forked from https://github.com/summa-tx/memview-sol with several breaking changes:
/// - The codebase is ported to Solidity 0.8
/// - Custom errors are added
/// - The runtime type checking is replaced with compile-time check provided by User-Defined Value Types
///   https://docs.soliditylang.org/en/latest/types.html#user-defined-value-types
/// - uint256 is used as the underlying type for the "memory view" instead of bytes29.
///   It is wrapped into MemView custom type in order not to be confused with actual integers.
/// - Therefore the "type" field is discarded, allowing to allocate 16 bytes for both view location and length
/// - The documentation is expanded
/// - Library functions unused by the rest of the codebase are removed
//  - Very pretty code separators are added :)
library MemViewLib {
    /// @notice Stack layout for uint256 (from highest bits to lowest)
    /// (32 .. 16]      loc     16 bytes    Memory address of underlying bytes
    /// (16 .. 00]      len     16 bytes    Length of underlying bytes

    // ═══════════════════════════════════════════ BUILDING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Instantiate a new untyped memory view. This should generally not be called directly.
     * Prefer `ref` wherever possible.
     * @param loc_          The memory address
     * @param len_          The length
     * @return The new view with the specified location and length
     */
    function build(uint256 loc_, uint256 len_) internal pure returns (MemView) {
        uint256 end_ = loc_ + len_;
        // Make sure that a view is not constructed that points to unallocated memory
        // as this could be indicative of a buffer overflow attack
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            if gt(end_, mload(0x40)) { end_ := 0 }
        }
        if (end_ == 0) {
            revert UnallocatedMemory();
        }
        return _unsafeBuildUnchecked(loc_, len_);
    }

    /**
     * @notice Instantiate a memory view from a byte array.
     * @dev Note that due to Solidity memory representation, it is not possible to
     * implement a deref, as the `bytes` type stores its len in memory.
     * @param arr           The byte array
     * @return The memory view over the provided byte array
     */
    function ref(bytes memory arr) internal pure returns (MemView) {
        uint256 len_ = arr.length;
        // `bytes arr` is stored in memory in the following way
        // 1. First, uint256 arr.length is stored. That requires 32 bytes (0x20).
        // 2. Then, the array data is stored.
        uint256 loc_;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // We add 0x20, so that the view starts exactly where the array data starts
            loc_ := add(arr, 0x20)
        }
        return build(loc_, len_);
    }

    // ════════════════════════════════════════════ CLONING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Copies the referenced memory to a new loc in memory, returning a `bytes` pointing to the new memory.
     * @param memView       The memory view
     * @return arr          The cloned byte array
     */
    function clone(MemView memView) internal view returns (bytes memory arr) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
            // This is where the byte array will be stored
            arr := ptr
        }
        unchecked {
            _unsafeCopyTo(memView, ptr + 0x20);
        }
        // `bytes arr` is stored in memory in the following way
        // 1. First, uint256 arr.length is stored. That requires 32 bytes (0x20).
        // 2. Then, the array data is stored.
        uint256 len_ = memView.len();
        uint256 footprint_ = memView.footprint();
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Write new unused pointer: the old value + array footprint + 32 bytes to store the length
            mstore(0x40, add(add(ptr, footprint_), 0x20))
            // Write len of new array (in bytes)
            mstore(ptr, len_)
        }
    }

    /**
     * @notice Copies all views, joins them into a new bytearray.
     * @param memViews      The memory views
     * @return arr          The new byte array with joined data behind the given views
     */
    function join(MemView[] memory memViews) internal view returns (bytes memory arr) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
            // This is where the byte array will be stored
            arr := ptr
        }
        MemView newView;
        unchecked {
            newView = _unsafeJoin(memViews, ptr + 0x20);
        }
        uint256 len_ = newView.len();
        uint256 footprint_ = newView.footprint();
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Write new unused pointer: the old value + array footprint + 32 bytes to store the length
            mstore(0x40, add(add(ptr, footprint_), 0x20))
            // Write len of new array (in bytes)
            mstore(ptr, len_)
        }
    }

    // ══════════════════════════════════════════ INSPECTING MEMORY VIEW ═══════════════════════════════════════════════

    /**
     * @notice Returns the memory address of the underlying bytes.
     * @param memView       The memory view
     * @return loc_         The memory address
     */
    function loc(MemView memView) internal pure returns (uint256 loc_) {
        // loc is stored in the highest 16 bytes of the underlying uint256
        return MemView.unwrap(memView) >> 128;
    }

    /**
     * @notice Returns the number of bytes of the view.
     * @param memView       The memory view
     * @return len_         The length of the view
     */
    function len(MemView memView) internal pure returns (uint256 len_) {
        // len is stored in the lowest 16 bytes of the underlying uint256
        return MemView.unwrap(memView) & type(uint128).max;
    }

    /**
     * @notice Returns the endpoint of `memView`.
     * @param memView       The memory view
     * @return end_         The endpoint of `memView`
     */
    function end(MemView memView) internal pure returns (uint256 end_) {
        // The endpoint never overflows uint128, let alone uint256, so we could use unchecked math here
        unchecked {
            return memView.loc() + memView.len();
        }
    }

    /**
     * @notice Returns the number of memory words this memory view occupies, rounded up.
     * @param memView       The memory view
     * @return words_       The number of memory words
     */
    function words(MemView memView) internal pure returns (uint256 words_) {
        // returning ceil(length / 32.0)
        unchecked {
            return (memView.len() + 31) >> 5;
        }
    }

    /**
     * @notice Returns the in-memory footprint of a fresh copy of the view.
     * @param memView       The memory view
     * @return footprint_   The in-memory footprint of a fresh copy of the view.
     */
    function footprint(MemView memView) internal pure returns (uint256 footprint_) {
        // words() * 32
        return memView.words() << 5;
    }

    // ════════════════════════════════════════════ HASHING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Returns the keccak256 hash of the underlying memory
     * @param memView       The memory view
     * @return digest       The keccak256 hash of the underlying memory
     */
    function keccak(MemView memView) internal pure returns (bytes32 digest) {
        uint256 loc_ = memView.loc();
        uint256 len_ = memView.len();
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            digest := keccak256(loc_, len_)
        }
    }

    /**
     * @notice Adds a salt to the keccak256 hash of the underlying data and returns the keccak256 hash of the
     * resulting data.
     * @param memView       The memory view
     * @return digestSalted keccak256(salt, keccak256(memView))
     */
    function keccakSalted(MemView memView, bytes32 salt) internal pure returns (bytes32 digestSalted) {
        return keccak256(bytes.concat(salt, memView.keccak()));
    }

    // ════════════════════════════════════════════ SLICING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Safe slicing without memory modification.
     * @param memView       The memory view
     * @param index_        The start index
     * @param len_          The length
     * @return The new view for the slice of the given length starting from the given index
     */
    function slice(MemView memView, uint256 index_, uint256 len_) internal pure returns (MemView) {
        uint256 loc_ = memView.loc();
        // Ensure it doesn't overrun the view
        if (loc_ + index_ + len_ > memView.end()) {
            revert ViewOverrun();
        }
        // Build a view starting from index with the given length
        unchecked {
            // loc_ + index_ <= memView.end()
            return build({loc_: loc_ + index_, len_: len_});
        }
    }

    /**
     * @notice Shortcut to `slice`. Gets a view representing bytes from `index` to end(memView).
     * @param memView       The memory view
     * @param index_        The start index
     * @return The new view for the slice starting from the given index until the initial view endpoint
     */
    function sliceFrom(MemView memView, uint256 index_) internal pure returns (MemView) {
        uint256 len_ = memView.len();
        // Ensure it doesn't overrun the view
        if (index_ > len_) {
            revert ViewOverrun();
        }
        // Build a view starting from index with the given length
        unchecked {
            // index_ <= len_ => memView.loc() + index_ <= memView.loc() + memView.len() == memView.end()
            return build({loc_: memView.loc() + index_, len_: len_ - index_});
        }
    }

    /**
     * @notice Shortcut to `slice`. Gets a view representing the first `len` bytes.
     * @param memView       The memory view
     * @param len_          The length
     * @return The new view for the slice of the given length starting from the initial view beginning
     */
    function prefix(MemView memView, uint256 len_) internal pure returns (MemView) {
        return memView.slice({index_: 0, len_: len_});
    }

    /**
     * @notice Shortcut to `slice`. Gets a view representing the last `len` byte.
     * @param memView       The memory view
     * @param len_          The length
     * @return The new view for the slice of the given length until the initial view endpoint
     */
    function postfix(MemView memView, uint256 len_) internal pure returns (MemView) {
        uint256 viewLen = memView.len();
        // Ensure it doesn't overrun the view
        if (len_ > viewLen) {
            revert ViewOverrun();
        }
        // Could do the unchecked math due to the check above
        uint256 index_;
        unchecked {
            index_ = viewLen - len_;
        }
        // Build a view starting from index with the given length
        unchecked {
            // len_ <= memView.len() => memView.loc() <= loc_ <= memView.end()
            return build({loc_: memView.loc() + viewLen - len_, len_: len_});
        }
    }

    // ═══════════════════════════════════════════ INDEXING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Load up to 32 bytes from the view onto the stack.
     * @dev Returns a bytes32 with only the `bytes_` HIGHEST bytes set.
     * This can be immediately cast to a smaller fixed-length byte array.
     * To automatically cast to an integer, use `indexUint`.
     * @param memView       The memory view
     * @param index_        The index
     * @param bytes_        The amount of bytes to load onto the stack
     * @return result       The 32 byte result having only `bytes_` highest bytes set
     */
    function index(MemView memView, uint256 index_, uint256 bytes_) internal pure returns (bytes32 result) {
        if (bytes_ == 0) {
            return bytes32(0);
        }
        // Can't load more than 32 bytes to the stack in one go
        if (bytes_ > 32) {
            revert IndexedTooMuch();
        }
        // The last indexed byte should be within view boundaries
        if (index_ + bytes_ > memView.len()) {
            revert ViewOverrun();
        }
        uint256 bitLength = bytes_ << 3; // bytes_ * 8
        uint256 loc_ = memView.loc();
        // Get a mask with `bitLength` highest bits set
        uint256 mask;
        // 0x800...00 binary representation is 100...00
        // sar stands for "signed arithmetic shift": https://en.wikipedia.org/wiki/Arithmetic_shift
        // sar(N-1, 100...00) = 11...100..00, with exactly N highest bits set to 1
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            mask := sar(sub(bitLength, 1), 0x8000000000000000000000000000000000000000000000000000000000000000)
        }
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load a full word using index offset, and apply mask to ignore non-relevant bytes
            result := and(mload(add(loc_, index_)), mask)
        }
    }

    /**
     * @notice Parse an unsigned integer from the view at `index`.
     * @dev Requires that the view have >= `bytes_` bytes following that index.
     * @param memView       The memory view
     * @param index_        The index
     * @param bytes_        The amount of bytes to load onto the stack
     * @return The unsigned integer
     */
    function indexUint(MemView memView, uint256 index_, uint256 bytes_) internal pure returns (uint256) {
        bytes32 indexedBytes = memView.index(index_, bytes_);
        // `index()` returns left-aligned `bytes_`, while integers are right-aligned
        // Shifting here to right-align with the full 32 bytes word: need to shift right `(32 - bytes_)` bytes
        unchecked {
            // memView.index() reverts when bytes_ > 32, thus unchecked math
            return uint256(indexedBytes) >> ((32 - bytes_) << 3);
        }
    }

    /**
     * @notice Parse an address from the view at `index`.
     * @dev Requires that the view have >= 20 bytes following that index.
     * @param memView       The memory view
     * @param index_        The index
     * @return The address
     */
    function indexAddress(MemView memView, uint256 index_) internal pure returns (address) {
        // index 20 bytes as `uint160`, and then cast to `address`
        return address(uint160(memView.indexUint(index_, 20)));
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns a memory view over the specified memory location
    /// without checking if it points to unallocated memory.
    function _unsafeBuildUnchecked(uint256 loc_, uint256 len_) private pure returns (MemView) {
        // There is no scenario where loc or len would overflow uint128, so we omit this check.
        // We use the highest 128 bits to encode the location and the lowest 128 bits to encode the length.
        return MemView.wrap((loc_ << 128) | len_);
    }

    /**
     * @notice Copy the view to a location, return an unsafe memory reference
     * @dev Super Dangerous direct memory access.
     * This reference can be overwritten if anything else modifies memory (!!!).
     * As such it MUST be consumed IMMEDIATELY. Update the free memory pointer to ensure the copied data
     * is not overwritten. This function is private to prevent unsafe usage by callers.
     * @param memView       The memory view
     * @param newLoc        The new location to copy the underlying view data
     * @return The memory view over the unsafe memory with the copied underlying data
     */
    function _unsafeCopyTo(MemView memView, uint256 newLoc) private view returns (MemView) {
        uint256 len_ = memView.len();
        uint256 oldLoc = memView.loc();

        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
        }
        // Revert if we're writing in occupied memory
        if (newLoc < ptr) {
            revert OccupiedMemory();
        }
        bool res;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // use the identity precompile (0x04) to copy
            res := staticcall(gas(), 0x04, oldLoc, len_, newLoc, len_)
        }
        if (!res) revert PrecompileOutOfGas();
        return _unsafeBuildUnchecked({loc_: newLoc, len_: len_});
    }

    /**
     * @notice Join the views in memory, return an unsafe reference to the memory.
     * @dev Super Dangerous direct memory access.
     * This reference can be overwritten if anything else modifies memory (!!!).
     * As such it MUST be consumed IMMEDIATELY. Update the free memory pointer to ensure the copied data
     * is not overwritten. This function is private to prevent unsafe usage by callers.
     * @param memViews      The memory views
     * @return The conjoined view pointing to the new memory
     */
    function _unsafeJoin(MemView[] memory memViews, uint256 location) private view returns (MemView) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
        }
        // Revert if we're writing in occupied memory
        if (location < ptr) {
            revert OccupiedMemory();
        }
        // Copy the views to the specified location one by one, by tracking the amount of copied bytes so far
        uint256 offset = 0;
        for (uint256 i = 0; i < memViews.length;) {
            MemView memView = memViews[i];
            // We can use the unchecked math here as location + sum(view.length) will never overflow uint256
            unchecked {
                _unsafeCopyTo(memView, location + offset);
                offset += memView.len();
                ++i;
            }
        }
        return _unsafeBuildUnchecked({loc_: location, len_: offset});
    }
}
