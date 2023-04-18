// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @dev MemView is an untyped view over a portion of memory to be used instead of `bytes memory`
type MemView is uint256;

/// @dev Attach library functions to MemView
using MemViewLib for MemView global;

/// @notice Library for operations with the memory views.
/// Forked from https://github.com/summa-tx/memview-sol with several breaking changes:
/// - The codebase is ported to Solidity 0.8
/// - The runtime type checking is replaced with compile-time check provided by User-Defined Value Types
///   https://docs.soliditylang.org/en/latest/types.html#user-defined-value-types
/// - Therefore the "type" field is discarded, allowing to allocate 16 bytes for both view location and length
/// - The documentation is expanded
//  - Very pretty code separators are added :)
library MemViewLib {
    error IndexedTooMuch();
    error ViewOverrun();
    error UnallocatedMemory();

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
        return (memView.len() + 31) / 32;
    }

    /**
     * @notice Returns the in-memory footprint of a fresh copy of the view.
     * @param memView       The memory view
     * @return footprint_   The in-memory footprint of a fresh copy of the view.
     */
    function footprint(MemView memView) internal pure returns (uint256 footprint_) {
        return memView.words() * 32;
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
        return build({loc_: loc_ + index_, len_: len_});
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
        // Could do the unchecked math due to the check above
        unchecked {
            len_ = len_ - index_;
        }
        // Build a view starting from index with the given length
        return build({loc_: memView.loc() + index_, len_: len_});
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
        return build({loc_: memView.loc() + index_, len_: len_});
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
        if (bytes_ > 32) {
            revert IndexedTooMuch();
        }
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
        return uint256(indexedBytes) >> ((32 - bytes_) << 3);
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
}
