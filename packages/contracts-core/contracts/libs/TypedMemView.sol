// SPDX-License-Identifier: MIT OR Apache-2.0
pragma solidity >=0.8.12;

library TypedMemView {
    // Why does this exist?
    // the solidity `bytes memory` type has a few weaknesses.
    // 1. You can't index ranges effectively
    // 2. You can't slice without copying
    // 3. The underlying data may represent any type
    // 4. Solidity never deallocates memory, and memory costs grow
    //    superlinearly

    // By using a memory view instead of a `bytes memory` we get the following
    // advantages:
    // 1. Slices are done on the stack, by manipulating the pointer
    // 2. We can index arbitrary ranges and quickly convert them to stack types
    // 3. We can insert type info into the pointer, and typecheck at runtime

    // This makes `TypedMemView` a useful tool for efficient zero-copy
    // algorithms.

    // Why bytes29?
    // We want to avoid confusion between views, digests, and other common
    // types so we chose a large and uncommonly used odd number of bytes
    //
    // Note that while bytes are left-aligned in a word, integers and addresses
    // are right-aligned. This means when working in assembly we have to
    // account for the 3 unused bytes on the righthand side
    //
    // First 5 bytes are a type flag.
    // - ff_ffff_fffe is reserved for unknown type.
    // - ff_ffff_ffff is reserved for invalid types/errors.
    // next 12 are memory address
    // next 12 are len
    // bottom 3 bytes are empty

    // Assumptions:
    // - non-modification of memory.
    // - No Solidity updates
    // - - wrt free mem point
    // - - wrt bytes representation in memory
    // - - wrt memory addressing in general

    // Usage:
    // - create type constants
    // - use `assertType` for runtime type assertions
    // - - unfortunately we can't do this at compile time yet :(
    // - recommended: implement modifiers that perform type checking
    // - - e.g.
    // - - `uint40 constant MY_TYPE = 3;`
    // - - ` modifier onlyMyType(bytes29 myView) { myView.assertType(MY_TYPE); }`
    // - instantiate a typed view from a bytearray using `ref`
    // - use `index` to inspect the contents of the view
    // - use `slice` to create smaller views into the same memory
    // - - `slice` can increase the offset
    // - - `slice can decrease the length`
    // - - must specify the output type of `slice`
    // - - `slice` will return a null view if you try to overrun
    // - - make sure to explicitly check for this with `notNull` or `assertType`
    // - use `equal` for typed comparisons.

    // The null view
    bytes29 public constant NULL = hex"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";

    /**
     * @dev Memory layout for bytes29
     * TODO (Chi): with the user defined types storing type is no longer necessary.
     * Update the library, transforming bytes29 to bytes24 in the process.
     * [000..005)   type     5 bytes    Type flag for the pointer
     * [005..017)   loc     12 bytes    Memory address of underlying bytes
     * [017..029)   len     12 bytes    Length of underlying bytes
     * [029..032)   empty    3 bytes    Not used
     */
    uint256 public constant BITS_TYPE = 40;
    uint256 public constant BITS_LOC = 96;
    uint256 public constant BITS_LEN = 96;
    uint256 public constant BITS_EMPTY = 24;

    // `SHIFT_X` is how much bits to shift for `X` to be in the very bottom bits
    uint256 public constant SHIFT_LEN = BITS_EMPTY; // 24
    uint256 public constant SHIFT_LOC = SHIFT_LEN + BITS_LEN; // 24 + 96 = 120
    uint256 public constant SHIFT_TYPE = SHIFT_LOC + BITS_LOC; // 24 + 96 + 96 = 216
    // Bitmask for the lowest 96 bits
    uint256 public constant LOW_96_BITS_MASK = type(uint96).max;

    // For nibble encoding
    bytes private constant NIBBLE_LOOKUP = "0123456789abcdef";

    /**
     * @notice Returns the encoded hex character that represents the lower 4 bits of the argument.
     * @param byte_     The byte
     * @return char     The encoded hex character
     */
    function nibbleHex(uint8 byte_) internal pure returns (uint8 char) {
        uint8 nibble = byte_ & 0x0f; // keep bottom 4 bits, zero out top 4 bits
        char = uint8(NIBBLE_LOOKUP[nibble]);
    }

    /**
     * @notice      Returns a uint16 containing the hex-encoded byte.
     * @param b     The byte
     * @return      encoded - The hex-encoded byte
     */
    function byteHex(uint8 b) internal pure returns (uint16 encoded) {
        encoded |= nibbleHex(b >> 4); // top 4 bits
        encoded <<= 8;
        encoded |= nibbleHex(b); // lower 4 bits
    }

    /**
     * @notice      Encodes the uint256 to hex. `first` contains the encoded top 16 bytes.
     *              `second` contains the encoded lower 16 bytes.
     *
     * @param b     The 32 bytes as uint256
     * @return      first - The top 16 bytes
     * @return      second - The bottom 16 bytes
     */
    function encodeHex(uint256 b) internal pure returns (uint256 first, uint256 second) {
        for (uint8 i = 31; i > 15;) {
            uint8 byte_ = uint8(b >> (i * 8));
            first |= byteHex(byte_);
            if (i != 16) {
                first <<= 16;
            }
            unchecked {
                i -= 1;
            }
        }

        // abusing underflow here =_=
        for (uint8 i = 15; i < 255;) {
            uint8 byte_ = uint8(b >> (i * 8));
            second |= byteHex(byte_);
            if (i != 0) {
                second <<= 16;
            }
            unchecked {
                i -= 1;
            }
        }
    }

    /**
     * @notice          Changes the endianness of a uint256.
     * @dev             https://graphics.stanford.edu/~seander/bithacks.html#ReverseParallel
     * @param b         The unsigned integer to reverse
     * @return          v - The reversed value
     */
    function reverseUint256(uint256 b) internal pure returns (uint256 v) {
        v = b;

        // swap bytes
        v = ((v >> 8) & 0x00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF)
            | ((v & 0x00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF) << 8);
        // swap 2-byte long pairs
        v = ((v >> 16) & 0x0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF)
            | ((v & 0x0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF) << 16);
        // swap 4-byte long pairs
        v = ((v >> 32) & 0x00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF)
            | ((v & 0x00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF) << 32);
        // swap 8-byte long pairs
        v = ((v >> 64) & 0x0000000000000000FFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF)
            | ((v & 0x0000000000000000FFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF) << 64);
        // swap 16-byte long pairs
        v = (v >> 128) | (v << 128);
    }

    /**
     * @notice      Create a mask with the highest `len_` bits set.
     * @param len_  The length
     * @return      mask - The mask
     */
    function leftMask(uint8 len_) private pure returns (uint256 mask) {
        // 0x800...00 binary representation is 100...00
        // sar stands for "signed arithmetic shift": https://en.wikipedia.org/wiki/Arithmetic_shift
        // sar(N-1, 100...00) = 11...100..00, with exactly N highest bits set to 1
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            mask := sar(sub(len_, 1), 0x8000000000000000000000000000000000000000000000000000000000000000)
        }
    }

    /**
     * @notice      Return the null view.
     * @return      bytes29 - The null view
     */
    // solhint-disable-next-line ordering
    function nullView() internal pure returns (bytes29) {
        return NULL;
    }

    /**
     * @notice      Check if the view is null.
     * @return      bool - True if the view is null
     */
    function isNull(bytes29 view_) internal pure returns (bool) {
        return view_ == NULL;
    }

    /**
     * @notice      Check if the view is not null.
     * @return      bool - True if the view is not null
     */
    function notNull(bytes29 view_) internal pure returns (bool) {
        return !isNull(view_);
    }

    /**
     * @notice          Check if the view is of a valid type and points to a valid location
     *                  in memory.
     * @dev             We perform this check by examining solidity's unallocated memory
     *                  pointer and ensuring that the view's upper bound is less than that.
     * @param view_     The view
     * @return          ret - True if the view is valid
     */
    function isValid(bytes29 view_) internal pure returns (bool ret) {
        if (typeOf(view_) == 0xffffffffff) {
            return false;
        }
        uint256 end_ = end(view_);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // View is valid if ("upper bound" <= "unallocated memory pointer")
            // Upper bound is exclusive, hence "<="
            ret := not(gt(end_, mload(0x40)))
        }
    }

    /**
     * @notice          Require that a typed memory view be valid.
     * @dev             Returns the view for easy chaining.
     * @param view_     The view
     * @return          bytes29 - The validated view
     */
    function assertValid(bytes29 view_) internal pure returns (bytes29) {
        require(isValid(view_), "Validity assertion failed");
        return view_;
    }

    /**
     * @notice          Return true if the view_ is of the expected type. Otherwise false.
     * @param view_     The view
     * @param expected  The expected type
     * @return          bool - True if the view_ is of the expected type
     */
    function isType(bytes29 view_, uint40 expected) internal pure returns (bool) {
        return typeOf(view_) == expected;
    }

    /**
     * @notice          Require that a typed memory view has a specific type.
     * @dev             Returns the view for easy chaining.
     * @param view_     The view
     * @param expected  The expected type
     * @return          bytes29 - The view with validated type
     */
    function assertType(bytes29 view_, uint40 expected) internal pure returns (bytes29) {
        if (!isType(view_, expected)) {
            (, uint256 g) = encodeHex(uint256(typeOf(view_)));
            (, uint256 e) = encodeHex(uint256(expected));
            string memory err =
                string(abi.encodePacked("Type assertion failed. Got 0x", uint80(g), ". Expected 0x", uint80(e)));
            revert(err);
        }
        return view_;
    }

    /**
     * @notice          Return an identical view with a different type.
     * @param view_     The view
     * @param newType   The new type
     * @return          newView - The new view with the specified type
     */
    function castTo(bytes29 view_, uint40 newType) internal pure returns (bytes29 newView) {
        // How many bits are the "type bits" occupying
        uint256 bitsType = BITS_TYPE;
        // How many bits are the "type bits" shifted from the bottom
        uint256 shiftType = SHIFT_TYPE;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // shift off the "type bits" (shift left, then sift right)
            newView := or(newView, shr(bitsType, shl(bitsType, view_)))
            // set the new "type bits" (shift left, then OR)
            newView := or(newView, shl(shiftType, newType))
        }
    }

    /**
     * @notice          Unsafe raw pointer construction. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @dev             Unsafe raw pointer construction. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @param type_     The type
     * @param loc_      The memory address
     * @param len_      The length
     * @return          newView - The new view with the specified type, location and length
     */
    function unsafeBuildUnchecked(uint256 type_, uint256 loc_, uint256 len_) private pure returns (bytes29 newView) {
        uint256 bitsLoc = BITS_LOC;
        uint256 bitsLen = BITS_LEN;
        uint256 bitsEmpty = BITS_EMPTY;
        // Ref memory layout
        // [000..005) 5 bytes of type
        // [005..017) 12 bytes of location
        // [017..029) 12 bytes of length
        // last 3 bits are blank and dropped in typecast
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // insert `type`, shift to prepare empty bits for `loc`
            newView := shl(bitsLoc, or(newView, type_))
            // insert `loc`, shift to prepare empty bits for `len`
            newView := shl(bitsLen, or(newView, loc_))
            // insert `len`, shift to insert 3 blank lowest bits
            newView := shl(bitsEmpty, or(newView, len_))
        }
    }

    /**
     * @notice          Instantiate a new memory view. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @dev             Instantiate a new memory view. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @param type_     The type
     * @param loc_      The memory address
     * @param len_      The length
     * @return          newView - The new view with the specified type, location and length
     */
    function build(uint256 type_, uint256 loc_, uint256 len_) internal pure returns (bytes29 newView) {
        uint256 end_ = loc_ + len_;
        // Make sure that a view is not constructed that points to unallocated memory
        // as this could be indicative of a buffer overflow attack
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            if gt(end_, mload(0x40)) { end_ := 0 }
        }
        if (end_ == 0) {
            return NULL;
        }
        newView = unsafeBuildUnchecked(type_, loc_, len_);
    }

    /**
     * @notice          Instantiate a memory view from a byte array.
     * @dev             Note that due to Solidity memory representation, it is not possible to
     *                  implement a deref, as the `bytes` type stores its len in memory.
     * @param arr       The byte array
     * @param newType   The type
     * @return          bytes29 - The memory view
     */
    function ref(bytes memory arr, uint40 newType) internal pure returns (bytes29) {
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

        return build(newType, loc_, len_);
    }

    /**
     * @notice          Return the associated type information.
     * @param view_     The memory view
     * @return          type_ - The type associated with the view
     */
    function typeOf(bytes29 view_) internal pure returns (uint40 type_) {
        // How many bits are the "type bits" shifted from the bottom
        uint256 shiftType = SHIFT_TYPE;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Shift out the bottom bits preceding "type bits". "type bits" are occupying
            // the highest bits, so all that's left is "type bits", OR is not required.
            type_ := shr(shiftType, view_)
        }
    }

    /**
     * @notice          Optimized type comparison. Checks that the 5-byte type flag is equal.
     * @param left      The first view
     * @param right     The second view
     * @return          bool - True if the 5-byte type flag is equal
     */
    function sameType(bytes29 left, bytes29 right) internal pure returns (bool) {
        // Check that the highest 5 bytes are equal: xor and shift out lower 27 bytes
        return (left ^ right) >> SHIFT_TYPE == 0;
    }

    /**
     * @notice          Return the memory address of the underlying bytes.
     * @param view_     The view
     * @return          loc_ - The memory address
     */
    function loc(bytes29 view_) internal pure returns (uint96 loc_) {
        // How many bits are the "loc bits" shifted from the bottom
        uint256 shiftLoc = SHIFT_LOC;
        // Mask for the bottom 96 bits
        uint256 uint96Mask = LOW_96_BITS_MASK;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Shift out the bottom bits preceding "loc bits".
            // Then use the lowest 96 bits to determine `loc` by applying the bit-mask.
            loc_ := and(shr(shiftLoc, view_), uint96Mask)
        }
    }

    /**
     * @notice          The number of memory words this memory view occupies, rounded up.
     * @param view_     The view
     * @return          uint256 - The number of memory words
     */
    function words(bytes29 view_) internal pure returns (uint256) {
        // returning ceil(length / 32.0)
        return (uint256(len(view_)) + 31) / 32;
    }

    /**
     * @notice          The in-memory footprint of a fresh copy of the view.
     * @param view_   The view
     * @return          uint256 - The in-memory footprint of a fresh copy of the view.
     */
    function footprint(bytes29 view_) internal pure returns (uint256) {
        return words(view_) * 32;
    }

    /**
     * @notice          The number of bytes of the view.
     * @param view_     The view
     * @return          len_ - The length of the view
     */
    function len(bytes29 view_) internal pure returns (uint96 len_) {
        // How many bits are the "len bits" shifted from the bottom
        uint256 shiftLen = SHIFT_LEN;
        // Mask for the bottom 96 bits
        uint256 uint96Mask = LOW_96_BITS_MASK;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Shift out the bottom bits preceding "len bits".
            // Then use the lowest 96 bits to determine `len` by applying the bit-mask.
            len_ := and(shr(shiftLen, view_), uint96Mask)
        }
    }

    /**
     * @notice          Returns the endpoint of `view_`.
     * @param view_   The view
     * @return          uint256 - The endpoint of `view_`
     */
    function end(bytes29 view_) internal pure returns (uint256) {
        unchecked {
            return loc(view_) + len(view_);
        }
    }

    /**
     * @notice          Safe slicing without memory modification.
     * @param view_     The view
     * @param index_    The start index
     * @param len_      The length
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function slice(bytes29 view_, uint256 index_, uint256 len_, uint40 newType) internal pure returns (bytes29) {
        uint256 loc_ = loc(view_);

        // Ensure it doesn't overrun the view
        if (loc_ + index_ + len_ > end(view_)) {
            return NULL;
        }

        loc_ = loc_ + index_;
        return build(newType, loc_, len_);
    }

    /**
     * @notice          Shortcut to `slice`. Gets a view representing
     *                  bytes from `index` to end(view_).
     * @param view_     The view
     * @param index_    The start index
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function sliceFrom(bytes29 view_, uint256 index_, uint40 newType) internal pure returns (bytes29) {
        return slice(view_, index_, len(view_) - index_, newType);
    }

    /**
     * @notice          Shortcut to `slice`. Gets a view representing the first `len` bytes.
     * @param view_     The view
     * @param len_      The length
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function prefix(bytes29 view_, uint256 len_, uint40 newType) internal pure returns (bytes29) {
        return slice(view_, 0, len_, newType);
    }

    /**
     * @notice          Shortcut to `slice`. Gets a view representing the last `len` byte.
     * @param view_     The view
     * @param len_      The length
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function postfix(bytes29 view_, uint256 len_, uint40 newType) internal pure returns (bytes29) {
        return slice(view_, uint256(len(view_)) - len_, len_, newType);
    }

    /**
     * @notice          Construct an error message for an indexing overrun.
     * @param loc_      The memory address
     * @param len_      The length
     * @param index_    The index
     * @param slice_    The slice where the overrun occurred
     * @return          err - The err
     */
    function indexErrOverrun(uint256 loc_, uint256 len_, uint256 index_, uint256 slice_)
        internal
        pure
        returns (string memory err)
    {
        (, uint256 a) = encodeHex(loc_);
        (, uint256 b) = encodeHex(len_);
        (, uint256 c) = encodeHex(index_);
        (, uint256 d) = encodeHex(slice_);
        err = string(
            abi.encodePacked(
                "TypedMemView/index - Overran the view. Slice is at 0x",
                uint48(a),
                " with length 0x",
                uint48(b),
                ". Attempted to index at offset 0x",
                uint48(c),
                " with length 0x",
                uint48(d),
                "."
            )
        );
    }

    /**
     * @notice          Load up to 32 bytes from the view onto the stack.
     * @dev             Returns a bytes32 with only the `bytes_` highest bytes set.
     *                  This can be immediately cast to a smaller fixed-length byte array.
     *                  To automatically cast to an integer, use `indexUint`.
     * @param view_     The view
     * @param index_    The index
     * @param bytes_    The bytes
     * @return          result - The 32 byte result
     */
    function index(bytes29 view_, uint256 index_, uint8 bytes_) internal pure returns (bytes32 result) {
        if (bytes_ == 0) {
            return bytes32(0);
        }
        if (index_ + bytes_ > len(view_)) {
            revert(indexErrOverrun(loc(view_), len(view_), index_, uint256(bytes_)));
        }
        require(bytes_ <= 32, "Index: more than 32 bytes");

        uint8 bitLength;
        unchecked {
            bitLength = bytes_ * 8;
        }
        uint256 loc_ = loc(view_);
        // Get a mask with `bitLength` highest bits set
        uint256 mask = leftMask(bitLength);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load a full word using index offset, and apply mask to ignore non-relevant bytes
            result := and(mload(add(loc_, index_)), mask)
        }
    }

    /**
     * @notice          Parse an unsigned integer from the view at `index`.
     * @dev             Requires that the view have >= `bytes_` bytes following that index.
     * @param view_     The view
     * @param index_    The index
     * @param bytes_    The bytes
     * @return          result - The unsigned integer
     */
    function indexUint(bytes29 view_, uint256 index_, uint8 bytes_) internal pure returns (uint256 result) {
        // `index()` returns left-aligned `bytes_`, while integers are right-aligned
        // Shifting here to right-align with the full 32 bytes word
        return uint256(index(view_, index_, bytes_)) >> ((32 - bytes_) * 8);
    }

    /**
     * @notice          Parse an unsigned integer from LE bytes.
     * @param view_     The view
     * @param index_    The index
     * @param bytes_    The bytes
     * @return          result - The unsigned integer
     */
    function indexLEUint(bytes29 view_, uint256 index_, uint8 bytes_) internal pure returns (uint256 result) {
        return reverseUint256(uint256(index(view_, index_, bytes_)));
    }

    /**
     * @notice          Parse an address from the view at `index`.
     *                  Requires that the view have >= 20 bytes following that index.
     * @param view_     The view
     * @param index_    The index
     * @return          address - The address
     */
    function indexAddress(bytes29 view_, uint256 index_) internal pure returns (address) {
        // index 20 bytes as `uint160`, and then cast to `address`
        return address(uint160(indexUint(view_, index_, 20)));
    }

    /**
     * @notice          Return the keccak256 hash of the underlying memory
     * @param view_     The view
     * @return          digest - The keccak256 hash of the underlying memory
     */
    function keccak(bytes29 view_) internal pure returns (bytes32 digest) {
        uint256 loc_ = loc(view_);
        uint256 len_ = len(view_);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            digest := keccak256(loc_, len_)
        }
    }

    /**
     * @notice          Return the sha2 digest of the underlying memory.
     * @dev             We explicitly deallocate memory afterwards.
     * @param view_     The view
     * @return          digest - The sha2 hash of the underlying memory
     */
    function sha2(bytes29 view_) internal view returns (bytes32 digest) {
        uint256 loc_ = loc(view_);
        uint256 len_ = len(view_);
        bool res;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            // sha2 precompile is 0x02
            res := staticcall(gas(), 0x02, loc_, len_, ptr, 0x20)
            digest := mload(ptr)
        }
        require(res, "sha2: out of gas");
    }

    /**
     * @notice          Implements bitcoin's hash160 (rmd160(sha2()))
     * @param view_     The pre-image
     * @return          digest - the Digest
     */
    function hash160(bytes29 view_) internal view returns (bytes20 digest) {
        uint256 loc_ = loc(view_);
        uint256 len_ = len(view_);
        bool res;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            // sha2 precompile is 0x02
            res := staticcall(gas(), 0x02, loc_, len_, ptr, 0x20)
            // rmd160 precompile is 0x03
            res := and(res, staticcall(gas(), 0x03, ptr, 0x20, ptr, 0x20))
            digest := mload(add(ptr, 0xc)) // return value is 0-prefixed.
        }
        require(res, "hash160: out of gas");
    }

    /**
     * @notice          Implements bitcoin's hash256 (double sha2)
     * @param view_     A view of the preimage
     * @return          digest - the Digest
     */
    function hash256(bytes29 view_) internal view returns (bytes32 digest) {
        uint256 loc_ = loc(view_);
        uint256 len_ = len(view_);
        bool res;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            // sha2 precompile is 0x02
            res := staticcall(gas(), 0x02, loc_, len_, ptr, 0x20)
            res := and(res, staticcall(gas(), 0x02, ptr, 0x20, ptr, 0x20))
            digest := mload(ptr)
        }
        require(res, "hash256: out of gas");
    }

    /**
     * @notice          Return true if the underlying memory is equal. Else false.
     * @param left      The first view
     * @param right     The second view
     * @return          bool - True if the underlying memory is equal
     */
    function untypedEqual(bytes29 left, bytes29 right) internal pure returns (bool) {
        return (loc(left) == loc(right) && len(left) == len(right)) || keccak(left) == keccak(right);
    }

    /**
     * @notice          Return false if the underlying memory is equal. Else true.
     * @param left      The first view
     * @param right     The second view
     * @return          bool - False if the underlying memory is equal
     */
    function untypedNotEqual(bytes29 left, bytes29 right) internal pure returns (bool) {
        return !untypedEqual(left, right);
    }

    /**
     * @notice          Compares type equality.
     * @dev             Shortcuts if the pointers are identical, otherwise compares type and digest.
     * @param left      The first view
     * @param right     The second view
     * @return          bool - True if the types are the same
     */
    function equal(bytes29 left, bytes29 right) internal pure returns (bool) {
        return left == right || (typeOf(left) == typeOf(right) && keccak(left) == keccak(right));
    }

    /**
     * @notice          Compares type inequality.
     * @dev             Shortcuts if the pointers are identical, otherwise compares type and digest.
     * @param left      The first view
     * @param right     The second view
     * @return          bool - True if the types are not the same
     */
    function notEqual(bytes29 left, bytes29 right) internal pure returns (bool) {
        return !equal(left, right);
    }

    /**
     * @notice          Copy the view to a location, return an unsafe memory reference
     * @dev             Super Dangerous direct memory access.
     *
     *                  This reference can be overwritten if anything else modifies memory (!!!).
     *                  As such it MUST be consumed IMMEDIATELY.
     *                  This function is private to prevent unsafe usage by callers.
     * @param view_     The view
     * @param newLoc    The new location
     * @return          written - the unsafe memory reference
     */
    function unsafeCopyTo(bytes29 view_, uint256 newLoc) private view returns (bytes29 written) {
        require(notNull(view_), "copyTo: Null pointer deref");
        require(isValid(view_), "copyTo: Invalid pointer deref");
        uint256 len_ = len(view_);
        uint256 oldLoc = loc(view_);

        uint256 ptr;
        bool res;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ptr := mload(0x40)
            // revert if we're writing in occupied memory
            if gt(ptr, newLoc) { revert(0x60, 0x20) } // empty revert message

            // use the identity precompile (0x04) to copy
            res := staticcall(gas(), 0x04, oldLoc, len_, newLoc, len_)
        }
        require(res, "identity: out of gas");

        written = unsafeBuildUnchecked(typeOf(view_), newLoc, len_);
    }

    /**
     * @notice          Copies the referenced memory to a new loc in memory,
     *                  returning a `bytes` pointing to the new memory.
     * @dev             Shortcuts if the pointers are identical, otherwise compares type and digest.
     * @param view_     The view
     * @return          ret - The view pointing to the new memory
     */
    function clone(bytes29 view_) internal view returns (bytes memory ret) {
        uint256 ptr;
        uint256 len_ = len(view_);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ptr := mload(0x40) // load unused memory pointer
            ret := ptr
        }
        unchecked {
            unsafeCopyTo(view_, ptr + 0x20);
        }
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            mstore(0x40, add(add(ptr, len_), 0x20)) // write new unused pointer
            mstore(ptr, len_) // write len of new array (in bytes)
        }
    }

    /**
     * @notice          Join the views in memory, return an unsafe reference to the memory.
     * @dev             Super Dangerous direct memory access.
     *
     *                  This reference can be overwritten if anything else modifies memory (!!!).
     *                  As such it MUST be consumed IMMEDIATELY.
     *                  This function is private to prevent unsafe usage by callers.
     * @param memViews  The views
     * @return          unsafeView - The conjoined view pointing to the new memory
     */
    function unsafeJoin(bytes29[] memory memViews, uint256 location) private view returns (bytes29 unsafeView) {
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            // revert if we're writing in occupied memory
            if gt(ptr, location) { revert(0x60, 0x20) } // empty revert message
        }

        uint256 offset = 0;
        for (uint256 i = 0; i < memViews.length; i++) {
            bytes29 view_ = memViews[i];
            unchecked {
                unsafeCopyTo(view_, location + offset);
                offset += len(view_);
            }
        }
        unsafeView = unsafeBuildUnchecked(0, location, offset);
    }

    /**
     * @notice          Produce the keccak256 digest of the concatenated contents of multiple views.
     * @param memViews  The views
     * @return          bytes32 - The keccak256 digest
     */
    function joinKeccak(bytes29[] memory memViews) internal view returns (bytes32) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ptr := mload(0x40) // load unused memory pointer
        }
        return keccak(unsafeJoin(memViews, ptr));
    }

    /**
     * @notice          Produce the sha256 digest of the concatenated contents of multiple views.
     * @param memViews  The views
     * @return          bytes32 - The sha256 digest
     */
    function joinSha2(bytes29[] memory memViews) internal view returns (bytes32) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ptr := mload(0x40) // load unused memory pointer
        }
        return sha2(unsafeJoin(memViews, ptr));
    }

    /**
     * @notice          copies all views, joins them into a new bytearray.
     * @param memViews  The views
     * @return          ret - The new byte array
     */
    function join(bytes29[] memory memViews) internal view returns (bytes memory ret) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ptr := mload(0x40) // load unused memory pointer
        }

        bytes29 newView;
        unchecked {
            newView = unsafeJoin(memViews, ptr + 0x20);
        }
        uint256 written = len(newView);
        uint256 footprint_ = footprint(newView);

        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // store the length
            mstore(ptr, written)
            // new pointer is old + 0x20 + the footprint of the body
            mstore(0x40, add(add(ptr, footprint_), 0x20))
            ret := ptr
        }
    }
}
