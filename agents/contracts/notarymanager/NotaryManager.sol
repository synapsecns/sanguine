pragma solidity 0.8.17;


interface INotaryManager {
    function slashNotary(address payable _reporter) external;

    function notary() external view returns (address);
}

abstract contract DomainContext {
    /**
     * @notice Ensures that a domain matches the local domain.
     */
    modifier onlyLocalDomain(uint32 _domain) {
        require(_domain == _localDomain(), "!localDomain");
        _;
    }

    function localDomain() external view returns (uint32) {
        return _localDomain();
    }

    function _localDomain() internal view virtual returns (uint32);
}

contract LocalDomainContext is DomainContext {
    uint32 private immutable __localDomain;

    constructor(uint32 localDomain_) {
        __localDomain = localDomain_;
    }

    function _localDomain() internal view override returns (uint32) {
        return __localDomain;
    }
}

abstract contract OriginEvents {
    /**
     * @notice Emitted when a new message is dispatched
     * @param messageHash Hash of message; the leaf inserted to the Merkle tree
     *        for the message
     * @param nonce Nonce of sent message (starts from 1)
     * @param destination Destination domain
     * @param tips Tips paid for the remote off-chain agents
     * @param message Raw bytes of message
     */
    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes tips,
        bytes message
    );

    /**
     * @notice Emitted when the Guard is slashed
     * (should be paired with IncorrectReport event)
     * @param guard     The address of the guard that signed the incorrect report
     * @param reporter  The address of the entity that reported the guard misbehavior
     */
    event GuardSlashed(address indexed guard, address indexed reporter);

    /**
     * @notice Emitted when the Notary is slashed
     * (should be paired with FraudAttestation event)
     * @param notary    The address of the notary
     * @param guard     The address of the guard that signed the fraud report
     * @param reporter  The address of the entity that reported the notary misbehavior
     */
    event NotarySlashed(address indexed notary, address indexed guard, address indexed reporter);

    /**
     * @notice Emitted when the NotaryManager contract is changed
     * @param notaryManager The address of the new notaryManager
     */
    event NewNotaryManager(address notaryManager);
}

contract Version0 {
    uint8 public constant VERSION = 0;
}

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
    // - - ` modifer onlyMyType(bytes29 myView) { myView.assertType(MY_TYPE); }`
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
    uint256 public constant LOW_12_MASK = 0xffffffffffffffffffffffff;
    uint8 public constant TWELVE_BYTES = 96;

    /**
     * @notice      Returns the encoded hex character that represents
     *              the lower 4 bits of the argument.
     * @param _b    The byte
     * @return      char - The encoded hex character
     */
    // solhint-disable-next-line code-complexity
    function nibbleHex(uint8 _b) internal pure returns (uint8 char) {
        // This can probably be done more efficiently, but it's only in error
        // paths, so we don't really care :)
        uint8 _nibble = _b | 0xf0; // set top 4, keep bottom 4
        if (_nibble == 0xf0) {
            return 0x30;
        } // 0
        if (_nibble == 0xf1) {
            return 0x31;
        } // 1
        if (_nibble == 0xf2) {
            return 0x32;
        } // 2
        if (_nibble == 0xf3) {
            return 0x33;
        } // 3
        if (_nibble == 0xf4) {
            return 0x34;
        } // 4
        if (_nibble == 0xf5) {
            return 0x35;
        } // 5
        if (_nibble == 0xf6) {
            return 0x36;
        } // 6
        if (_nibble == 0xf7) {
            return 0x37;
        } // 7
        if (_nibble == 0xf8) {
            return 0x38;
        } // 8
        if (_nibble == 0xf9) {
            return 0x39;
        } // 9
        if (_nibble == 0xfa) {
            return 0x61;
        } // a
        if (_nibble == 0xfb) {
            return 0x62;
        } // b
        if (_nibble == 0xfc) {
            return 0x63;
        } // c
        if (_nibble == 0xfd) {
            return 0x64;
        } // d
        if (_nibble == 0xfe) {
            return 0x65;
        } // e
        if (_nibble == 0xff) {
            return 0x66;
        } // f
    }

    /**
     * @notice      Returns a uint16 containing the hex-encoded byte.
     * @param _b    The byte
     * @return      encoded - The hex-encoded byte
     */
    function byteHex(uint8 _b) internal pure returns (uint16 encoded) {
        encoded |= nibbleHex(_b >> 4); // top 4 bits
        encoded <<= 8;
        encoded |= nibbleHex(_b); // lower 4 bits
    }

    /**
     * @notice      Encodes the uint256 to hex. `first` contains the encoded top 16 bytes.
     *              `second` contains the encoded lower 16 bytes.
     *
     * @param _b    The 32 bytes as uint256
     * @return      first - The top 16 bytes
     * @return      second - The bottom 16 bytes
     */
    function encodeHex(uint256 _b) internal pure returns (uint256 first, uint256 second) {
        for (uint8 i = 31; i > 15; ) {
            uint8 _byte = uint8(_b >> (i * 8));
            first |= byteHex(_byte);
            if (i != 16) {
                first <<= 16;
            }
            unchecked {
                i -= 1;
            }
        }

        // abusing underflow here =_=
        for (uint8 i = 15; i < 255; ) {
            uint8 _byte = uint8(_b >> (i * 8));
            second |= byteHex(_byte);
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
     * @param _b        The unsigned integer to reverse
     * @return          v - The reversed value
     */
    function reverseUint256(uint256 _b) internal pure returns (uint256 v) {
        v = _b;

        // swap bytes
        v =
            ((v >> 8) & 0x00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF) |
            ((v & 0x00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF) << 8);
        // swap 2-byte long pairs
        v =
            ((v >> 16) & 0x0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF) |
            ((v & 0x0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF) << 16);
        // swap 4-byte long pairs
        v =
            ((v >> 32) & 0x00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF) |
            ((v & 0x00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF) << 32);
        // swap 8-byte long pairs
        v =
            ((v >> 64) & 0x0000000000000000FFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF) |
            ((v & 0x0000000000000000FFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF) << 64);
        // swap 16-byte long pairs
        v = (v >> 128) | (v << 128);
    }

    /**
     * @notice      Create a mask with the highest `_len` bits set.
     * @param _len  The length
     * @return      mask - The mask
     */
    function leftMask(uint8 _len) private pure returns (uint256 mask) {
        // 0x800...00 binary representation is 100...00
        // sar stands for "signed arithmetic shift": https://en.wikipedia.org/wiki/Arithmetic_shift
        // sar(N-1, 100...00) = 11...100..00, with exactly N highest bits set to 1
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            mask := sar(
                sub(_len, 1),
                0x8000000000000000000000000000000000000000000000000000000000000000
            )
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
    function isNull(bytes29 memView) internal pure returns (bool) {
        return memView == NULL;
    }

    /**
     * @notice      Check if the view is not null.
     * @return      bool - True if the view is not null
     */
    function notNull(bytes29 memView) internal pure returns (bool) {
        return !isNull(memView);
    }

    /**
     * @notice          Check if the view is of a valid type and points to a valid location
     *                  in memory.
     * @dev             We perform this check by examining solidity's unallocated memory
     *                  pointer and ensuring that the view's upper bound is less than that.
     * @param memView   The view
     * @return          ret - True if the view is valid
     */
    function isValid(bytes29 memView) internal pure returns (bool ret) {
        if (typeOf(memView) == 0xffffffffff) {
            return false;
        }
        uint256 _end = end(memView);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ret := not(gt(_end, mload(0x40)))
        }
    }

    /**
     * @notice          Require that a typed memory view be valid.
     * @dev             Returns the view for easy chaining.
     * @param memView   The view
     * @return          bytes29 - The validated view
     */
    function assertValid(bytes29 memView) internal pure returns (bytes29) {
        require(isValid(memView), "Validity assertion failed");
        return memView;
    }

    /**
     * @notice          Return true if the memview is of the expected type. Otherwise false.
     * @param memView   The view
     * @param _expected The expected type
     * @return          bool - True if the memview is of the expected type
     */
    function isType(bytes29 memView, uint40 _expected) internal pure returns (bool) {
        return typeOf(memView) == _expected;
    }

    /**
     * @notice          Require that a typed memory view has a specific type.
     * @dev             Returns the view for easy chaining.
     * @param memView   The view
     * @param _expected The expected type
     * @return          bytes29 - The view with validated type
     */
    function assertType(bytes29 memView, uint40 _expected) internal pure returns (bytes29) {
        if (!isType(memView, _expected)) {
            (, uint256 g) = encodeHex(uint256(typeOf(memView)));
            (, uint256 e) = encodeHex(uint256(_expected));
            string memory err = string(
                abi.encodePacked(
                    "Type assertion failed. Got 0x",
                    uint80(g),
                    ". Expected 0x",
                    uint80(e)
                )
            );
            revert(err);
        }
        return memView;
    }

    /**
     * @notice          Return an identical view with a different type.
     * @param memView   The view
     * @param _newType  The new type
     * @return          newView - The new view with the specified type
     */
    function castTo(bytes29 memView, uint40 _newType) internal pure returns (bytes29 newView) {
        // then | in the new type
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // shift off the top 5 bytes
            newView := or(newView, shr(40, shl(40, memView)))
            newView := or(newView, shl(216, _newType))
        }
    }

    /**
     * @notice          Unsafe raw pointer construction. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @dev             Unsafe raw pointer construction. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @param _type     The type
     * @param _loc      The memory address
     * @param _len      The length
     * @return          newView - The new view with the specified type, location and length
     */
    function unsafeBuildUnchecked(
        uint256 _type,
        uint256 _loc,
        uint256 _len
    ) private pure returns (bytes29 newView) {
        /// @dev Ref memory layout
        /// [000..005) 5 bytes of type
        /// [005..017) 12 bytes of location
        /// [017..029) 12 bytes of length
        /// last 3 bits are blank and dropped in typecast
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            newView := shl(96, or(newView, _type)) // insert type
            newView := shl(96, or(newView, _loc)) // insert loc
            newView := shl(24, or(newView, _len)) // empty bottom 3 bytes
        }
    }

    /**
     * @notice          Instantiate a new memory view. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @dev             Instantiate a new memory view. This should generally not be called
     *                  directly. Prefer `ref` wherever possible.
     * @param _type     The type
     * @param _loc      The memory address
     * @param _len      The length
     * @return          newView - The new view with the specified type, location and length
     */
    function build(
        uint256 _type,
        uint256 _loc,
        uint256 _len
    ) internal pure returns (bytes29 newView) {
        uint256 _end = _loc + _len;
        // Make sure that a view is not constructed that points to unallocated memory
        // as this could be indicative of a buffer overflow attack
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            if gt(_end, mload(0x40)) {
                _end := 0
            }
        }
        if (_end == 0) {
            return NULL;
        }
        newView = unsafeBuildUnchecked(_type, _loc, _len);
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
        uint256 _len = arr.length;

        uint256 _loc;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            _loc := add(arr, 0x20) // our view is of the data, not the struct
        }

        return build(newType, _loc, _len);
    }

    /**
     * @notice          Return the associated type information.
     * @param memView   The memory view
     * @return          _type - The type associated with the view
     */
    function typeOf(bytes29 memView) internal pure returns (uint40 _type) {
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // 216 == 256 - 40
            _type := shr(216, memView) // shift out lower 24 bytes
        }
    }

    /**
     * @notice          Optimized type comparison. Checks that the 5-byte type flag is equal.
     * @param left      The first view
     * @param right     The second view
     * @return          bool - True if the 5-byte type flag is equal
     */
    function sameType(bytes29 left, bytes29 right) internal pure returns (bool) {
        return (left ^ right) >> (2 * TWELVE_BYTES) == 0;
    }

    /**
     * @notice          Return the memory address of the underlying bytes.
     * @param memView   The view
     * @return          _loc - The memory address
     */
    function loc(bytes29 memView) internal pure returns (uint96 _loc) {
        uint256 _mask = LOW_12_MASK; // assembly can't use globals
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // 120 bits = 12 bytes (the encoded loc) + 3 bytes (empty low space)
            _loc := and(shr(120, memView), _mask)
        }
    }

    /**
     * @notice          The number of memory words this memory view occupies, rounded up.
     * @param memView   The view
     * @return          uint256 - The number of memory words
     */
    function words(bytes29 memView) internal pure returns (uint256) {
        return (uint256(len(memView)) + 32) / 32;
    }

    /**
     * @notice          The in-memory footprint of a fresh copy of the view.
     * @param memView   The view
     * @return          uint256 - The in-memory footprint of a fresh copy of the view.
     */
    function footprint(bytes29 memView) internal pure returns (uint256) {
        return words(memView) * 32;
    }

    /**
     * @notice          The number of bytes of the view.
     * @param memView   The view
     * @return          _len - The length of the view
     */
    function len(bytes29 memView) internal pure returns (uint96 _len) {
        uint256 _mask = LOW_12_MASK; // assembly can't use globals
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            _len := and(shr(24, memView), _mask)
        }
    }

    /**
     * @notice          Returns the endpoint of `memView`.
     * @param memView   The view
     * @return          uint256 - The endpoint of `memView`
     */
    function end(bytes29 memView) internal pure returns (uint256) {
        unchecked {
            return loc(memView) + len(memView);
        }
    }

    /**
     * @notice          Safe slicing without memory modification.
     * @param memView   The view
     * @param _index    The start index
     * @param _len      The length
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function slice(
        bytes29 memView,
        uint256 _index,
        uint256 _len,
        uint40 newType
    ) internal pure returns (bytes29) {
        uint256 _loc = loc(memView);

        // Ensure it doesn't overrun the view
        if (_loc + _index + _len > end(memView)) {
            return NULL;
        }

        _loc = _loc + _index;
        return build(newType, _loc, _len);
    }

    /**
     * @notice          Shortcut to `slice`. Gets a view representing
     *                  bytes from `_index` to end(memView).
     * @param memView   The view
     * @param _index    The start index
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function sliceFrom(
        bytes29 memView,
        uint256 _index,
        uint40 newType
    ) internal pure returns (bytes29) {
        return slice(memView, _index, len(memView) - _index, newType);
    }

    /**
     * @notice          Shortcut to `slice`. Gets a view representing the first `_len` bytes.
     * @param memView   The view
     * @param _len      The length
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function prefix(
        bytes29 memView,
        uint256 _len,
        uint40 newType
    ) internal pure returns (bytes29) {
        return slice(memView, 0, _len, newType);
    }

    /**
     * @notice          Shortcut to `slice`. Gets a view representing the last `_len` byte.
     * @param memView   The view
     * @param _len      The length
     * @param newType   The new type
     * @return          bytes29 - The new view
     */
    function postfix(
        bytes29 memView,
        uint256 _len,
        uint40 newType
    ) internal pure returns (bytes29) {
        return slice(memView, uint256(len(memView)) - _len, _len, newType);
    }

    /**
     * @notice          Construct an error message for an indexing overrun.
     * @param _loc      The memory address
     * @param _len      The length
     * @param _index    The index
     * @param _slice    The slice where the overrun occurred
     * @return          err - The err
     */
    function indexErrOverrun(
        uint256 _loc,
        uint256 _len,
        uint256 _index,
        uint256 _slice
    ) internal pure returns (string memory err) {
        (, uint256 a) = encodeHex(_loc);
        (, uint256 b) = encodeHex(_len);
        (, uint256 c) = encodeHex(_index);
        (, uint256 d) = encodeHex(_slice);
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
     * @dev             Returns a bytes32 with only the `_bytes` highest bytes set.
     *                  This can be immediately cast to a smaller fixed-length byte array.
     *                  To automatically cast to an integer, use `indexUint`.
     * @param memView   The view
     * @param _index    The index
     * @param _bytes    The bytes
     * @return          result - The 32 byte result
     */
    function index(
        bytes29 memView,
        uint256 _index,
        uint8 _bytes
    ) internal pure returns (bytes32 result) {
        if (_bytes == 0) {
            return bytes32(0);
        }
        if (_index + _bytes > len(memView)) {
            revert(indexErrOverrun(loc(memView), len(memView), _index, uint256(_bytes)));
        }
        require(_bytes <= 32, "Index: more than 32 bytes");

        uint8 bitLength;
        unchecked {
            bitLength = _bytes * 8;
        }
        uint256 _loc = loc(memView);
        uint256 _mask = leftMask(bitLength);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            result := and(mload(add(_loc, _index)), _mask)
        }
    }

    /**
     * @notice          Parse an unsigned integer from the view at `_index`.
     * @dev             Requires that the view have >= `_bytes` bytes following that index.
     * @param memView   The view
     * @param _index    The index
     * @param _bytes    The bytes
     * @return          result - The unsigned integer
     */
    function indexUint(
        bytes29 memView,
        uint256 _index,
        uint8 _bytes
    ) internal pure returns (uint256 result) {
        return uint256(index(memView, _index, _bytes)) >> ((32 - _bytes) * 8);
    }

    /**
     * @notice          Parse an unsigned integer from LE bytes.
     * @param memView   The view
     * @param _index    The index
     * @param _bytes    The bytes
     * @return          result - The unsigned integer
     */
    function indexLEUint(
        bytes29 memView,
        uint256 _index,
        uint8 _bytes
    ) internal pure returns (uint256 result) {
        return reverseUint256(uint256(index(memView, _index, _bytes)));
    }

    /**
     * @notice          Parse an address from the view at `_index`.
     *                  Requires that the view have >= 20 bytes following that index.
     * @param memView   The view
     * @param _index    The index
     * @return          address - The address
     */
    function indexAddress(bytes29 memView, uint256 _index) internal pure returns (address) {
        return address(uint160(indexUint(memView, _index, 20)));
    }

    /**
     * @notice          Return the keccak256 hash of the underlying memory
     * @param memView   The view
     * @return          digest - The keccak256 hash of the underlying memory
     */
    function keccak(bytes29 memView) internal pure returns (bytes32 digest) {
        uint256 _loc = loc(memView);
        uint256 _len = len(memView);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            digest := keccak256(_loc, _len)
        }
    }

    /**
     * @notice          Return the sha2 digest of the underlying memory.
     * @dev             We explicitly deallocate memory afterwards.
     * @param memView   The view
     * @return          digest - The sha2 hash of the underlying memory
     */
    function sha2(bytes29 memView) internal view returns (bytes32 digest) {
        uint256 _loc = loc(memView);
        uint256 _len = len(memView);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            pop(staticcall(gas(), 2, _loc, _len, ptr, 0x20)) // sha2 #1
            digest := mload(ptr)
        }
    }

    /**
     * @notice          Implements bitcoin's hash160 (rmd160(sha2()))
     * @param memView   The pre-image
     * @return          digest - the Digest
     */
    function hash160(bytes29 memView) internal view returns (bytes20 digest) {
        uint256 _loc = loc(memView);
        uint256 _len = len(memView);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            pop(staticcall(gas(), 2, _loc, _len, ptr, 0x20)) // sha2
            pop(staticcall(gas(), 3, ptr, 0x20, ptr, 0x20)) // rmd160
            digest := mload(add(ptr, 0xc)) // return value is 0-prefixed.
        }
    }

    /**
     * @notice          Implements bitcoin's hash256 (double sha2)
     * @param memView   A view of the preimage
     * @return          digest - the Digest
     */
    function hash256(bytes29 memView) internal view returns (bytes32 digest) {
        uint256 _loc = loc(memView);
        uint256 _len = len(memView);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            pop(staticcall(gas(), 2, _loc, _len, ptr, 0x20)) // sha2 #1
            pop(staticcall(gas(), 2, ptr, 0x20, ptr, 0x20)) // sha2 #2
            digest := mload(ptr)
        }
    }

    /**
     * @notice          Return true if the underlying memory is equal. Else false.
     * @param left      The first view
     * @param right     The second view
     * @return          bool - True if the underlying memory is equal
     */
    function untypedEqual(bytes29 left, bytes29 right) internal pure returns (bool) {
        return
            (loc(left) == loc(right) && len(left) == len(right)) || keccak(left) == keccak(right);
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
     * @param memView   The view
     * @param _newLoc   The new location
     * @return          written - the unsafe memory reference
     */
    function unsafeCopyTo(bytes29 memView, uint256 _newLoc) private view returns (bytes29 written) {
        require(notNull(memView), "copyTo: Null pointer deref");
        require(isValid(memView), "copyTo: Invalid pointer deref");
        uint256 _len = len(memView);
        uint256 _oldLoc = loc(memView);

        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ptr := mload(0x40)
            // revert if we're writing in occupied memory
            if gt(ptr, _newLoc) {
                revert(0x60, 0x20) // empty revert message
            }

            // use the identity precompile to copy
            // guaranteed not to fail, so pop the success
            pop(staticcall(gas(), 4, _oldLoc, _len, _newLoc, _len))
        }

        written = unsafeBuildUnchecked(typeOf(memView), _newLoc, _len);
    }

    /**
     * @notice          Copies the referenced memory to a new loc in memory,
     *                  returning a `bytes` pointing to the new memory.
     * @dev             Shortcuts if the pointers are identical, otherwise compares type and digest.
     * @param memView   The view
     * @return          ret - The view pointing to the new memory
     */
    function clone(bytes29 memView) internal view returns (bytes memory ret) {
        uint256 ptr;
        uint256 _len = len(memView);
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            ptr := mload(0x40) // load unused memory pointer
            ret := ptr
        }
        unchecked {
            unsafeCopyTo(memView, ptr + 0x20);
        }
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            mstore(0x40, add(add(ptr, _len), 0x20)) // write new unused pointer
            mstore(ptr, _len) // write len of new array (in bytes)
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
    function unsafeJoin(bytes29[] memory memViews, uint256 _location)
        private
        view
        returns (bytes29 unsafeView)
    {
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            let ptr := mload(0x40)
            // revert if we're writing in occupied memory
            if gt(ptr, _location) {
                revert(0x60, 0x20) // empty revert message
            }
        }

        uint256 _offset = 0;
        for (uint256 i = 0; i < memViews.length; i++) {
            bytes29 memView = memViews[i];
            unchecked {
                unsafeCopyTo(memView, _location + _offset);
                _offset += len(memView);
            }
        }
        unsafeView = unsafeBuildUnchecked(0, _location, _offset);
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

        bytes29 _newView;
        unchecked {
            _newView = unsafeJoin(memViews, ptr + 0x20);
        }
        uint256 _written = len(_newView);
        uint256 _footprint = footprint(_newView);

        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // store the legnth
            mstore(ptr, _written)
            // new pointer is old + 0x20 + the footprint of the body
            mstore(0x40, add(add(ptr, _footprint), 0x20))
            ret := ptr
        }
    }
}

library SynapseTypes {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          0X00: BYTE STRINGS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * 1. RAW_BYTES refers to a generic byte string, that is not supposed to be parsed
     * by the messaging contracts. RAW_BYTES is set to uint40(0) so that
     * the "default zero" type would represent a generic byte string.
     * 2. SIGNATURE refers to 65 bytes string that is an off-chain agent signature for some data.
     * 3. CALL_PAYLOAD refers to the payload, that is supposed to be used for an external call, i.e.
     * recipient.call(CALL_PAYLOAD). Its length is always (4 + 32 * N) bytes:
     *      - First 4 bytes represent the function selector.
     *      - 32 * N bytes represent N function arguments.
     */
    // prettier-ignore
    uint40 internal constant RAW_BYTES                  = 0x00_00_00_00_00;
    // prettier-ignore
    uint40 internal constant SIGNATURE                  = 0x00_01_00_00_00;
    // prettier-ignore
    uint40 internal constant CALL_PAYLOAD               = 0x00_02_00_00_00;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X01: ATTESTATION                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant ATTESTATION                = 0x01_01_00_00_00;
    // prettier-ignore
    uint40 internal constant ATTESTATION_DATA           = 0x01_01_01_00_00;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X02: REPORT                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant REPORT                     = 0x02_01_00_00_00;
    // prettier-ignore
    uint40 internal constant REPORT_DATA                = 0x02_01_01_00_00;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X03: MESSAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant MESSAGE                    = 0x03_01_00_00_00;
    // prettier-ignore
    uint40 internal constant MESSAGE_HEADER             = 0x03_01_01_00_00;
    // prettier-ignore
    uint40 internal constant MESSAGE_TIPS               = 0x03_01_02_00_00;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             0X04: SYSTEM                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant SYSTEM_CALL                = 0x04_00_00_00_00;
}

library ByteString {
    using TypedMemView for bytes29;

    // @dev non-compact ECDSA signatures are enforced as of OZ 4.7.3
    uint256 internal constant SIGNATURE_LENGTH = 65;

    /**
     * @dev Call payload memory layout
     * [000 .. 004) selector    bytes4  4 bytes
     *      Optional: N function arguments
     * [004 .. 036) arg1        bytes32 32 bytes
     *      ..
     * [AAA .. END) argN        bytes32 32 bytes
     */
    uint256 internal constant SELECTOR_LENGTH = 4;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyType(bytes29 _view, uint40 _type) {
        _view.assertType(_type);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Checks that a byte string is a signature
     */
    function isSignature(bytes29 _view) internal pure returns (bool) {
        return _view.len() == SIGNATURE_LENGTH;
    }

    /**
     * @notice Checks that a byte string is a call payload, i.e.
     * a function selector, followed by arbitrary amount of arguments.
     */
    function isCallPayload(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Call payload should at least have a function selector
        if (length < SELECTOR_LENGTH) return false;
        // The remainder of the payload should be exactly N words (N >= 0), i.e.
        // (length - SELECTOR_LENGTH) % 32 == 0
        // We're using logical AND here to speed it up a bit
        return (length - SELECTOR_LENGTH) & 31 == 0;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CALL PAYLOAD SLICING                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns amount of memory words (32 byte chunks) the function arguments
     * occupy in the call payload.
     * @dev This might differ from amount of arguments supplied, if any of the arguments
     * occupies more than one memory slot. It is true, however, that argument part of the payload
     * occupies exactly N words, even for dynamic types like `bytes`
     */
    function argumentWords(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.CALL_PAYLOAD)
        returns (uint256)
    {
        // Equivalent of (length - SELECTOR_LENGTH) / 32
        return (_view.len() - SELECTOR_LENGTH) >> 5;
    }
}

library Attestation {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev AttestationData memory layout
     * [000 .. 004): origin         uint32   4 bytes
     * [004 .. 008): destination    uint32   4 bytes
     * [008 .. 012): nonce          uint32   4 bytes
     * [012 .. 044): root           bytes32 32 bytes
     *
     *      Attestation memory layout
     * [000 .. 044): attData        bytes   44 bytes (see above)
     * [044 .. 109): signature      bytes   65 bytes (65 bytes)
     */

    uint256 internal constant OFFSET_ORIGIN = 0;
    uint256 internal constant OFFSET_DESTINATION = 4;
    uint256 internal constant OFFSET_NONCE = 8;
    uint256 internal constant OFFSET_ROOT = 12;
    uint256 internal constant ATTESTATION_DATA_LENGTH = 44;
    uint256 internal constant OFFSET_SIGNATURE = ATTESTATION_DATA_LENGTH;
    uint256 internal constant ATTESTATION_LENGTH = ATTESTATION_DATA_LENGTH + 65;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyAttestation(bytes29 _view) {
        _view.assertType(SynapseTypes.ATTESTATION);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for an attestation payload.
     */
    function castToAttestation(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.ATTESTATION);
    }

    /**
     * @notice Returns a formatted Attestation payload with provided fields
     * @param _data         Attestation Data (see above)
     * @param _signature    Notary's signature on `_data`
     * @return Formatted attestation
     **/
    function formatAttestation(bytes memory _data, bytes memory _signature)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_data, _signature);
    }

    /**
     * @notice Returns a formatted AttestationData payload with provided fields
     * @param _origin       Domain of Origin's chain
     * @param _destination  Domain of Destination's chain
     * @param _root         New merkle root
     * @param _nonce        Nonce of the merkle root
     * @return Formatted attestation data
     **/
    function formatAttestationData(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_origin, _destination, _nonce, _root);
    }

    /**
     * @notice Checks that a payload is a formatted Attestation payload.
     */
    function isAttestation(bytes29 _view) internal pure returns (bool) {
        return _view.len() == ATTESTATION_LENGTH;
    }

    /**
     * @notice Combines origin and destination domains into `attestationDomains`,
     * a unique ID for every (origin, destination) pair. Could be used to identify
     * Merkle trees on Origin, or Mirrors on Destination.
     */
    function attestationDomains(uint32 _origin, uint32 _destination)
        internal
        pure
        returns (uint64)
    {
        return (uint64(_origin) << 32) | _destination;
    }

    /**
     * @notice Combines origin, destination domains and message nonce into `attestationKey`,
     * a unique key for every (origin, destination, nonce) tuple. Could be used to identify
     * any dispatched message.
     */
    function attestationKey(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce
    ) internal pure returns (uint96) {
        return (uint96(_origin) << 64) | (uint96(_destination) << 32) | _nonce;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns domain of chain where the Origin contract is deployed
     */
    function attestedOrigin(bytes29 _view) internal pure onlyAttestation(_view) returns (uint32) {
        return uint32(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 4 }));
    }

    /**
     * @notice Returns domain of chain where the Destination contract is deployed
     */
    function attestedDestination(bytes29 _view)
        internal
        pure
        onlyAttestation(_view)
        returns (uint32)
    {
        return uint32(_view.indexUint({ _index: OFFSET_DESTINATION, _bytes: 4 }));
    }

    /**
     * @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
     */
    function attestedNonce(bytes29 _view) internal pure onlyAttestation(_view) returns (uint32) {
        return uint32(_view.indexUint({ _index: OFFSET_NONCE, _bytes: 4 }));
    }

    /**
     * @notice Returns a combined field for (origin, destination). See `attestationDomains()`.
     */
    function attestedDomains(bytes29 _view) internal pure onlyAttestation(_view) returns (uint64) {
        return uint64(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 8 }));
    }

    /**
     * @notice Returns a combined field for (origin, destination, nonce). See `attestationKey()`.
     */
    function attestedKey(bytes29 _view) internal pure onlyAttestation(_view) returns (uint96) {
        return uint96(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 12 }));
    }

    /**
     * @notice Returns a historical Merkle root from the Origin contract
     */
    function attestedRoot(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes32) {
        return _view.index({ _index: OFFSET_ROOT, _bytes: 32 });
    }

    /**
     * @notice Returns Attestation's Data, that is going to be signed by the Notary
     */
    function attestationData(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes29) {
        return
            _view.slice({
                _index: OFFSET_ORIGIN,
                _len: ATTESTATION_DATA_LENGTH,
                newType: SynapseTypes.ATTESTATION_DATA
            });
    }

    /**
     * @notice Returns Notary's signature on AttestationData
     */
    function notarySignature(bytes29 _view) internal pure onlyAttestation(_view) returns (bytes29) {
        return
            _view.slice({
                _index: OFFSET_SIGNATURE,
                _len: ByteString.SIGNATURE_LENGTH,
                newType: SynapseTypes.SIGNATURE
            });
    }
}

library Report {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev More flag values could be added in the future,
     *      e.g. flag indicating "type" of fraud.
     *      Going forward, Flag.Valid is guaranteed to be
     *      the only Flag specifying a valid attestation.
     *
     *      Flag.Valid indicates a reported valid Attestation.
     *      Flag.Fraud indicates a reported fraud Attestation.
     */
    enum Flag {
        Valid,
        Fraud
    }

    /**
     * @dev ReportData memory layout
     * [000 .. 001): flag           uint8    1 bytes
     * [001 .. 045): attData        bytes   44 bytes
     *
     * guardSig is Guard's signature on ReportData
     *
     *      Report memory layout
     * [000 .. 001): flag           uint8    1 bytes
     * [001 .. 110): attestation    bytes   109 bytes (44 + 65 bytes)
     * [110 .. 175): guardSig       bytes   65 bytes
     *
     *      Unpack attestation field (see Attestation.sol)
     * [000 .. 001): flag           uint8    1 bytes
     * [001 .. 045): attData        bytes   44 bytes
     * [045 .. 110): notarySig      bytes   65 bytes
     * [110 .. 175): guardSig       bytes   65 bytes
     *
     * notarySig is Notary's signature on AttestationData
     *
     * flag + attData = reportData (see above), so
     *
     *      Report memory layout (sliced alternatively)
     * [000 .. 045): reportData     bytes   45 bytes
     * [045 .. 110): notarySig      bytes   65 bytes
     * [110 .. 171): guardSig       bytes   65 bytes
     */

    uint256 internal constant OFFSET_FLAG = 0;
    uint256 internal constant OFFSET_ATTESTATION = 1;

    uint256 internal constant ATTESTATION_DATA_LENGTH = 44;
    uint256 internal constant REPORT_DATA_LENGTH = 1 + ATTESTATION_DATA_LENGTH;
    uint256 internal constant REPORT_LENGTH = REPORT_DATA_LENGTH + 2 * 65;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyReport(bytes29 _view) {
        _view.assertType(SynapseTypes.REPORT);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       FORMATTERS: REPORT DATA                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns formatted report data with provided fields
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatReportData(Flag _flag, bytes memory _attestation)
        internal
        view
        returns (bytes memory)
    {
        // Extract attestation data from payload
        bytes memory attestationData = _attestation.castToAttestation().attestationData().clone();
        // Construct report data
        return abi.encodePacked(uint8(_flag), attestationData);
    }

    /**
     * @notice Returns formatted report data on valid attestation with provided fields
     * @param _validAttestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatValidReportData(bytes memory _validAttestation)
        internal
        view
        returns (bytes memory)
    {
        return formatReportData(Flag.Valid, _validAttestation);
    }

    /**
     * @notice Returns formatted report data on fraud attestation with provided fields
     * @param _fraudAttestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatFraudReportData(bytes memory _fraudAttestation)
        internal
        view
        returns (bytes memory)
    {
        return formatReportData(Flag.Fraud, _fraudAttestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          FORMATTERS: REPORT                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a report payload.
     */
    function castToReport(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.REPORT);
    }

    /**
     * @notice Returns formatted report payload with provided fields.
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @param _guardSig     Guard signature on reportData (see formatReportData below)
     * @return Formatted report
     **/
    function formatReport(
        Flag _flag,
        bytes memory _attestation,
        bytes memory _guardSig
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(uint8(_flag), _attestation, _guardSig);
    }

    /**
     * @notice Returns formatted report payload on a valid attestation with provided fields.
     * @param _validAttestation Formatted attestation (see Attestation.sol)
     * @param _guardSig         Guard signature on reportData (see ReportData section above)
     * @return Formatted report
     **/
    function formatValidReport(bytes memory _validAttestation, bytes memory _guardSig)
        internal
        pure
        returns (bytes memory)
    {
        return formatReport(Flag.Valid, _validAttestation, _guardSig);
    }

    /**
     * @notice Returns formatted report payload on a fraud attestation with provided fields.
     * @param _fraudAttestation Formatted attestation (see Attestation.sol)
     * @param _guardSig         Guard signature on reportData (see ReportData section above)
     * @return Formatted report
     **/
    function formatFraudReport(bytes memory _fraudAttestation, bytes memory _guardSig)
        internal
        pure
        returns (bytes memory)
    {
        return formatReport(Flag.Fraud, _fraudAttestation, _guardSig);
    }

    /**
     * @notice Checks that a payload is a formatted Report payload.
     */
    function isReport(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Report should be the correct length
        if (length != REPORT_LENGTH) return false;
        // Flag needs to match an existing enum value
        if (_flagIntValue(_view) > uint8(type(Flag).max)) return false;
        // Attestation needs to be formatted as well
        return reportedAttestation(_view).isAttestation();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            REPORT SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns whether Report's Flag is Fraud (indicating fraudulent attestation).
     */
    function reportedFraud(bytes29 _view) internal pure onlyReport(_view) returns (bool) {
        return _flagIntValue(_view) != uint8(Flag.Valid);
    }

    /**
     * @notice Returns Report's Attestation (which is supposed to be signed by the Notary already).
     */
    function reportedAttestation(bytes29 _view) internal pure onlyReport(_view) returns (bytes29) {
        return
            _view.slice({
                _index: OFFSET_ATTESTATION,
                _len: Attestation.ATTESTATION_LENGTH,
                newType: SynapseTypes.ATTESTATION
            });
    }

    /**
     * @notice Returns Report's Data, that is going to be signed by the Guard.
     */
    function reportData(bytes29 _view) internal pure onlyReport(_view) returns (bytes29) {
        // reportData starts from Flag
        return
            _view.slice({
                _index: OFFSET_FLAG,
                _len: REPORT_DATA_LENGTH,
                newType: SynapseTypes.REPORT_DATA
            });
    }

    /**
     * @notice Returns Guard's signature on ReportData.
     */
    function guardSignature(bytes29 _view) internal pure onlyReport(_view) returns (bytes29) {
        uint256 offsetSignature = OFFSET_ATTESTATION + Attestation.ATTESTATION_LENGTH;
        return
            _view.slice({
                _index: offsetSignature,
                _len: ByteString.SIGNATURE_LENGTH,
                newType: SynapseTypes.SIGNATURE
            });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Returns int value of Report flag.
     *      Needed to prevent overflow when casting to Flag.
     */
    function _flagIntValue(bytes29 _view) private pure returns (uint8 flagIntValue) {
        flagIntValue = uint8(_view.indexUint({ _index: OFFSET_FLAG, _bytes: 1 }));
    }
}

// 
// OpenZeppelin Contracts (last updated v4.7.0) (utils/Strings.sol)
/**
 * @dev String operations.
 */
library Strings {
    bytes16 private constant _HEX_SYMBOLS = "0123456789abcdef";
    uint8 private constant _ADDRESS_LENGTH = 20;

    /**
     * @dev Converts a `uint256` to its ASCII `string` decimal representation.
     */
    function toString(uint256 value) internal pure returns (string memory) {
        // Inspired by OraclizeAPI's implementation - MIT licence
        // https://github.com/oraclize/ethereum-api/blob/b42146b063c7d6ee1358846c198246239e9360e8/oraclizeAPI_0.4.25.sol

        if (value == 0) {
            return "0";
        }
        uint256 temp = value;
        uint256 digits;
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
        bytes memory buffer = new bytes(digits);
        while (value != 0) {
            digits -= 1;
            buffer[digits] = bytes1(uint8(48 + uint256(value % 10)));
            value /= 10;
        }
        return string(buffer);
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation.
     */
    function toHexString(uint256 value) internal pure returns (string memory) {
        if (value == 0) {
            return "0x00";
        }
        uint256 temp = value;
        uint256 length = 0;
        while (temp != 0) {
            length++;
            temp >>= 8;
        }
        return toHexString(value, length);
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation with fixed length.
     */
    function toHexString(uint256 value, uint256 length) internal pure returns (string memory) {
        bytes memory buffer = new bytes(2 * length + 2);
        buffer[0] = "0";
        buffer[1] = "x";
        for (uint256 i = 2 * length + 1; i > 1; --i) {
            buffer[i] = _HEX_SYMBOLS[value & 0xf];
            value >>= 4;
        }
        require(value == 0, "Strings: hex length insufficient");
        return string(buffer);
    }

    /**
     * @dev Converts an `address` with fixed length of 20 bytes to its not checksummed ASCII `string` hexadecimal representation.
     */
    function toHexString(address addr) internal pure returns (string memory) {
        return toHexString(uint256(uint160(addr)), _ADDRESS_LENGTH);
    }
}

library ECDSA {
    enum RecoverError {
        NoError,
        InvalidSignature,
        InvalidSignatureLength,
        InvalidSignatureS,
        InvalidSignatureV
    }

    function _throwError(RecoverError error) private pure {
        if (error == RecoverError.NoError) {
            return; // no error: do nothing
        } else if (error == RecoverError.InvalidSignature) {
            revert("ECDSA: invalid signature");
        } else if (error == RecoverError.InvalidSignatureLength) {
            revert("ECDSA: invalid signature length");
        } else if (error == RecoverError.InvalidSignatureS) {
            revert("ECDSA: invalid signature 's' value");
        } else if (error == RecoverError.InvalidSignatureV) {
            revert("ECDSA: invalid signature 'v' value");
        }
    }

    /**
     * @dev Returns the address that signed a hashed message (`hash`) with
     * `signature` or error string. This address can then be used for verification purposes.
     *
     * The `ecrecover` EVM opcode allows for malleable (non-unique) signatures:
     * this function rejects them by requiring the `s` value to be in the lower
     * half order, and the `v` value to be either 27 or 28.
     *
     * IMPORTANT: `hash` _must_ be the result of a hash operation for the
     * verification to be secure: it is possible to craft signatures that
     * recover to arbitrary addresses for non-hashed data. A safe way to ensure
     * this is by receiving a hash of the original message (which may otherwise
     * be too long), and then calling {toEthSignedMessageHash} on it.
     *
     * Documentation for signature generation:
     * - with https://web3js.readthedocs.io/en/v1.3.4/web3-eth-accounts.html#sign[Web3.js]
     * - with https://docs.ethers.io/v5/api/signer/#Signer-signMessage[ethers]
     *
     * _Available since v4.3._
     */
    function tryRecover(bytes32 hash, bytes memory signature) internal pure returns (address, RecoverError) {
        if (signature.length == 65) {
            bytes32 r;
            bytes32 s;
            uint8 v;
            // ecrecover takes the signature parameters, and the only way to get them
            // currently is to use assembly.
            /// @solidity memory-safe-assembly
            assembly {
                r := mload(add(signature, 0x20))
                s := mload(add(signature, 0x40))
                v := byte(0, mload(add(signature, 0x60)))
            }
            return tryRecover(hash, v, r, s);
        } else {
            return (address(0), RecoverError.InvalidSignatureLength);
        }
    }

    /**
     * @dev Returns the address that signed a hashed message (`hash`) with
     * `signature`. This address can then be used for verification purposes.
     *
     * The `ecrecover` EVM opcode allows for malleable (non-unique) signatures:
     * this function rejects them by requiring the `s` value to be in the lower
     * half order, and the `v` value to be either 27 or 28.
     *
     * IMPORTANT: `hash` _must_ be the result of a hash operation for the
     * verification to be secure: it is possible to craft signatures that
     * recover to arbitrary addresses for non-hashed data. A safe way to ensure
     * this is by receiving a hash of the original message (which may otherwise
     * be too long), and then calling {toEthSignedMessageHash} on it.
     */
    function recover(bytes32 hash, bytes memory signature) internal pure returns (address) {
        (address recovered, RecoverError error) = tryRecover(hash, signature);
        _throwError(error);
        return recovered;
    }

    /**
     * @dev Overload of {ECDSA-tryRecover} that receives the `r` and `vs` short-signature fields separately.
     *
     * See https://eips.ethereum.org/EIPS/eip-2098[EIP-2098 short signatures]
     *
     * _Available since v4.3._
     */
    function tryRecover(
        bytes32 hash,
        bytes32 r,
        bytes32 vs
    ) internal pure returns (address, RecoverError) {
        bytes32 s = vs & bytes32(0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff);
        uint8 v = uint8((uint256(vs) >> 255) + 27);
        return tryRecover(hash, v, r, s);
    }

    /**
     * @dev Overload of {ECDSA-recover} that receives the `r and `vs` short-signature fields separately.
     *
     * _Available since v4.2._
     */
    function recover(
        bytes32 hash,
        bytes32 r,
        bytes32 vs
    ) internal pure returns (address) {
        (address recovered, RecoverError error) = tryRecover(hash, r, vs);
        _throwError(error);
        return recovered;
    }

    /**
     * @dev Overload of {ECDSA-tryRecover} that receives the `v`,
     * `r` and `s` signature fields separately.
     *
     * _Available since v4.3._
     */
    function tryRecover(
        bytes32 hash,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal pure returns (address, RecoverError) {
        // EIP-2 still allows signature malleability for ecrecover(). Remove this possibility and make the signature
        // unique. Appendix F in the Ethereum Yellow paper (https://ethereum.github.io/yellowpaper/paper.pdf), defines
        // the valid range for s in (301): 0 < s < secp256k1n ÷ 2 + 1, and for v in (302): v ∈ {27, 28}. Most
        // signatures from current libraries generate a unique signature with an s-value in the lower half order.
        //
        // If your library generates malleable signatures, such as s-values in the upper range, calculate a new s-value
        // with 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141 - s1 and flip v from 27 to 28 or
        // vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept
        // these malleable signatures as well.
        if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {
            return (address(0), RecoverError.InvalidSignatureS);
        }
        if (v != 27 && v != 28) {
            return (address(0), RecoverError.InvalidSignatureV);
        }

        // If the signature is valid (and not malleable), return the signer address
        address signer = ecrecover(hash, v, r, s);
        if (signer == address(0)) {
            return (address(0), RecoverError.InvalidSignature);
        }

        return (signer, RecoverError.NoError);
    }

    /**
     * @dev Overload of {ECDSA-recover} that receives the `v`,
     * `r` and `s` signature fields separately.
     */
    function recover(
        bytes32 hash,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal pure returns (address) {
        (address recovered, RecoverError error) = tryRecover(hash, v, r, s);
        _throwError(error);
        return recovered;
    }

    /**
     * @dev Returns an Ethereum Signed Message, created from a `hash`. This
     * produces hash corresponding to the one signed with the
     * https://eth.wiki/json-rpc/API#eth_sign[`eth_sign`]
     * JSON-RPC method as part of EIP-191.
     *
     * See {recover}.
     */
    function toEthSignedMessageHash(bytes32 hash) internal pure returns (bytes32) {
        // 32 is the length in bytes of hash,
        // enforced by the type signature above
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }

    /**
     * @dev Returns an Ethereum Signed Message, created from `s`. This
     * produces hash corresponding to the one signed with the
     * https://eth.wiki/json-rpc/API#eth_sign[`eth_sign`]
     * JSON-RPC method as part of EIP-191.
     *
     * See {recover}.
     */
    function toEthSignedMessageHash(bytes memory s) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n", Strings.toString(s.length), s));
    }

    /**
     * @dev Returns an Ethereum Signed Typed Data, created from a
     * `domainSeparator` and a `structHash`. This produces hash corresponding
     * to the one signed with the
     * https://eips.ethereum.org/EIPS/eip-712[`eth_signTypedData`]
     * JSON-RPC method as part of EIP-712.
     *
     * See {recover}.
     */
    function toTypedDataHash(bytes32 domainSeparator, bytes32 structHash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));
    }
}

library Auth {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @notice Recovers signer from data and signature.
     * @param _data         Data that was signed
     * @param _signature    `_data` signed by `signer`
     * @return signer       Address that signed the data
     */
    function recoverSigner(bytes29 _data, bytes memory _signature)
        internal
        pure
        returns (address signer)
    {
        bytes32 digest = _data.keccak();
        digest = ECDSA.toEthSignedMessageHash(digest);
        signer = ECDSA.recover(digest, _signature);
    }
}

abstract contract NotaryRegistryEvents {
    /*
     * @notice Emitted when a new Notary is added.
     * @param domain    Domain where a Notary was added
     * @param notary    Address of the added notary
     */
    event NotaryAdded(uint32 indexed domain, address notary);

    /**
     * @notice Emitted when a new Notary is removed.
     * @param domain    Domain where a Notary was removed
     * @param notary    Address of the removed notary
     */
    event NotaryRemoved(uint32 indexed domain, address notary);
}

abstract contract AbstractNotaryRegistry is NotaryRegistryEvents {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Adds a new Notary to Registry.
     * @dev Child contracts should implement this depending on how Notaries are stored.
     * @param _origin   Origin domain where Notary is added
     * @param _notary   New Notary to add
     * @return TRUE if a notary was added
     */
    function _addNotary(uint32 _origin, address _notary) internal virtual returns (bool);

    /**
     * @notice Removes a Notary from Registry.
     * @dev Child contracts should implement this depending on how Notaries are stored.
     * @param _origin   Origin domain where Notary is removed
     * @param _notary   Notary to remove
     * @return TRUE if a notary was removed
     */
    function _removeNotary(uint32 _origin, address _notary) internal virtual returns (bool);

    /**
     * @notice  Checks all following statements are true:
     *          - `_attestation` is a formatted Attestation payload
     *          - `_attestation` contains a signature
     *          - such signature belongs to an authorized Notary
     * @param _attestation  Attestation of Origin merkle root
     * @return _notary      Notary that signed the Attestation
     * @return _view        Memory view on attestation
     */
    function _checkNotaryAuth(bytes memory _attestation)
        internal
        view
        returns (address _notary, bytes29 _view)
    {
        _view = _attestation.castToAttestation();
        _notary = _checkNotaryAuth(_view);
    }

    /**
     * @notice  Checks all following statements are true:
     *          - `_view` is a memory view on a formatted Attestation payload
     *          - `_view` contains a signature
     *          - such signature belongs to an authorized Notary
     * @param _view     Memory view on Attestation of Origin merkle root
     * @return _notary  Notary that signed the Attestation
     */
    function _checkNotaryAuth(bytes29 _view) internal view returns (address _notary) {
        require(_view.isAttestation(), "Not an attestation");
        _notary = Auth.recoverSigner(_view.attestationData(), _view.notarySignature().clone());
        require(_isNotary(_view.attestedOrigin(), _notary), "Signer is not a notary");
    }

    /**
     * @notice Checks whether a given account in an authorized Notary.
     * @dev Child contracts should implement this depending on how Notaries are stored.
     * @param _origin   Origin domain to check
     * @param _account  Address to check for being a Notary
     * @return TRUE if the account is an authorized Notary.
     */
    function _isNotary(uint32 _origin, address _account) internal view virtual returns (bool);
}

library EnumerableSet {
    // To implement this library for multiple types with as little code
    // repetition as possible, we write it in terms of a generic Set type with
    // bytes32 values.
    // The Set implementation uses private functions, and user-facing
    // implementations (such as AddressSet) are just wrappers around the
    // underlying Set.
    // This means that we can only create new EnumerableSets for types that fit
    // in bytes32.

    struct Set {
        // Storage of set values
        bytes32[] _values;
        // Position of the value in the `values` array, plus 1 because index 0
        // means a value is not in the set.
        mapping(bytes32 => uint256) _indexes;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function _add(Set storage set, bytes32 value) private returns (bool) {
        if (!_contains(set, value)) {
            set._values.push(value);
            // The value is stored at length-1, but we add 1 to all indexes
            // and use 0 as a sentinel value
            set._indexes[value] = set._values.length;
            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function _remove(Set storage set, bytes32 value) private returns (bool) {
        // We read and store the value's index to prevent multiple reads from the same storage slot
        uint256 valueIndex = set._indexes[value];

        if (valueIndex != 0) {
            // Equivalent to contains(set, value)
            // To delete an element from the _values array in O(1), we swap the element to delete with the last one in
            // the array, and then remove the last element (sometimes called as 'swap and pop').
            // This modifies the order of the array, as noted in {at}.

            uint256 toDeleteIndex = valueIndex - 1;
            uint256 lastIndex = set._values.length - 1;

            if (lastIndex != toDeleteIndex) {
                bytes32 lastValue = set._values[lastIndex];

                // Move the last value to the index where the value to delete is
                set._values[toDeleteIndex] = lastValue;
                // Update the index for the moved value
                set._indexes[lastValue] = valueIndex; // Replace lastValue's index to valueIndex
            }

            // Delete the slot where the moved value was stored
            set._values.pop();

            // Delete the index for the deleted slot
            delete set._indexes[value];

            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function _contains(Set storage set, bytes32 value) private view returns (bool) {
        return set._indexes[value] != 0;
    }

    /**
     * @dev Returns the number of values on the set. O(1).
     */
    function _length(Set storage set) private view returns (uint256) {
        return set._values.length;
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function _at(Set storage set, uint256 index) private view returns (bytes32) {
        return set._values[index];
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function _values(Set storage set) private view returns (bytes32[] memory) {
        return set._values;
    }

    // Bytes32Set

    struct Bytes32Set {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(Bytes32Set storage set, bytes32 value) internal returns (bool) {
        return _add(set._inner, value);
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(Bytes32Set storage set, bytes32 value) internal returns (bool) {
        return _remove(set._inner, value);
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(Bytes32Set storage set, bytes32 value) internal view returns (bool) {
        return _contains(set._inner, value);
    }

    /**
     * @dev Returns the number of values in the set. O(1).
     */
    function length(Bytes32Set storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(Bytes32Set storage set, uint256 index) internal view returns (bytes32) {
        return _at(set._inner, index);
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function values(Bytes32Set storage set) internal view returns (bytes32[] memory) {
        return _values(set._inner);
    }

    // AddressSet

    struct AddressSet {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(AddressSet storage set, address value) internal returns (bool) {
        return _add(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(AddressSet storage set, address value) internal returns (bool) {
        return _remove(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(AddressSet storage set, address value) internal view returns (bool) {
        return _contains(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Returns the number of values in the set. O(1).
     */
    function length(AddressSet storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(AddressSet storage set, uint256 index) internal view returns (address) {
        return address(uint160(uint256(_at(set._inner, index))));
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function values(AddressSet storage set) internal view returns (address[] memory) {
        bytes32[] memory store = _values(set._inner);
        address[] memory result;

        /// @solidity memory-safe-assembly
        assembly {
            result := store
        }

        return result;
    }

    // UintSet

    struct UintSet {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(UintSet storage set, uint256 value) internal returns (bool) {
        return _add(set._inner, bytes32(value));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(UintSet storage set, uint256 value) internal returns (bool) {
        return _remove(set._inner, bytes32(value));
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(UintSet storage set, uint256 value) internal view returns (bool) {
        return _contains(set._inner, bytes32(value));
    }

    /**
     * @dev Returns the number of values on the set. O(1).
     */
    function length(UintSet storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(UintSet storage set, uint256 index) internal view returns (uint256) {
        return uint256(_at(set._inner, index));
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function values(UintSet storage set) internal view returns (uint256[] memory) {
        bytes32[] memory store = _values(set._inner);
        uint256[] memory result;

        /// @solidity memory-safe-assembly
        assembly {
            result := store
        }

        return result;
    }
}

abstract contract DomainNotaryRegistry is AbstractNotaryRegistry, DomainContext {
    using EnumerableSet for EnumerableSet.AddressSet;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // All active notaries for the tracked chain
    EnumerableSet.AddressSet internal notaries;

    // gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Ensures that there is at least one active Notary.
     */
    modifier haveActiveNotary() {
        require(notariesAmount() != 0, "!notaries");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns addresses of all Notaries.
     * @dev This copies storage into memory, so can consume a lof of gas, if
     * amount of notaries is large (see EnumerableSet.values())
     */
    function allNotaries() external view returns (address[] memory) {
        return notaries.values();
    }

    /**
     * @notice Returns i-th Notary. O(1)
     * @dev Will revert if index is out of range
     */
    function getNotary(uint256 _index) public view returns (address) {
        return notaries.at(_index);
    }

    /**
     * @notice Returns amount of active notaries. O(1)
     */
    function notariesAmount() public view returns (uint256) {
        return notaries.length();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Tries to add a new notary, emits an event only if notary was added.
     * @dev Reverts if domain doesn't match the tracked domain.
     */
    function _addNotary(uint32 _domain, address _notary)
        internal
        override
        onlyLocalDomain(_domain)
        returns (bool)
    {
        return _addNotary(_notary);
    }

    /**
     * @notice Tries to add a new notary, emits an event only if notary was added.
     */
    function _addNotary(address _notary) internal returns (bool notaryAdded) {
        notaryAdded = notaries.add(_notary);
        if (notaryAdded) {
            emit NotaryAdded(_localDomain(), _notary);
        }
    }

    /**
     * @notice Tries to remove a notary, emits an event only if notary was removed.
     * @dev Reverts if domain doesn't match the tracked domain.
     */
    function _removeNotary(uint32 _domain, address _notary)
        internal
        override
        onlyLocalDomain(_domain)
        returns (bool)
    {
        return _removeNotary(_notary);
    }

    /**
     * @notice Tries to remove a notary, emits an event only if notary was removed.
     */
    function _removeNotary(address _notary) internal returns (bool notaryRemoved) {
        notaryRemoved = notaries.remove(_notary);
        if (notaryRemoved) {
            emit NotaryRemoved(_localDomain(), _notary);
        }
    }

    /**
     * @notice Returns whether given address is a notary for the tracked domain.
     * @dev Reverts if domain doesn't match the tracked domain.
     */
    function _isNotary(uint32 _domain, address _account)
        internal
        view
        override
        onlyLocalDomain(_domain)
        returns (bool)
    {
        return _isNotary(_account);
    }

    /**
     * @notice Returns whether given address is a notary for the tracked domain.
     */
    function _isNotary(address _account) internal view returns (bool) {
        return notaries.contains(_account);
    }
}

abstract contract GuardRegistryEvents {
    /**
     * @notice Emitted when a new Guard is added.
     * @param guard    Address of the added guard
     */
    event GuardAdded(address guard);

    /**
     * @notice Emitted when a Guard is removed.
     * @param guard    Address of the removed guard
     */
    event GuardRemoved(address guard);
}

abstract contract AbstractGuardRegistry is GuardRegistryEvents {
    using Report for bytes;
    using Report for bytes29;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Adds a new Guard to Registry.
     * @dev Child contracts should implement this depending on how Guards are stored.
     * @param _guard    New Guard to add
     * @return TRUE if a guard was added
     */
    function _addGuard(address _guard) internal virtual returns (bool);

    /**
     * @notice Removes a Guard from Registry.
     * @dev Child contracts should implement this depending on how Guards are stored.
     * @param _guard    Guard to remove
     * @return TRUE if a guard was removed
     */
    function _removeGuard(address _guard) internal virtual returns (bool);

    /**
     * @notice  Checks all following statements are true:
     *          - `_report` is a formatted Report payload
     *          - `_report` contains a signature
     *          - such signature belongs to an authorized Guard
     * @param _report   Report on a Attestation of Origin merkle root
     * @return _guard   Notary that signed the Attestation
     * @return _view    Memory view on report
     */
    function _checkGuardAuth(bytes memory _report)
        internal
        view
        returns (address _guard, bytes29 _view)
    {
        _view = _report.castToReport();
        require(_view.isReport(), "Not a report");
        _guard = Auth.recoverSigner(_view.reportData(), _view.guardSignature().clone());
        require(_isGuard(_guard), "Signer is not a guard");
    }

    /**
     * @notice Checks whether a given account in an authorized Guard.
     * @dev Child contracts should implement this depending on how Guards are stored.
     * @param _account  Address to check for being a Guard
     * @return TRUE if the account is an authorized Guard.
     */
    function _isGuard(address _account) internal view virtual returns (bool);
}

contract GuardRegistry is AbstractGuardRegistry {
    using EnumerableSet for EnumerableSet.AddressSet;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // All active guards
    EnumerableSet.AddressSet internal guards;

    // gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns addresses of all Guards.
     * @dev This copies storage into memory, so can consume a lof of gas, if
     * amount of notaries is large (see EnumerableSet.values())
     */
    function allGuards() external view returns (address[] memory) {
        return guards.values();
    }

    /**
     * @notice Returns i-th Guard. O(1)
     * @dev Will revert if index is out of range
     */
    function getGuard(uint256 _index) external view returns (address) {
        return guards.at(_index);
    }

    /**
     * @notice Returns amount of active guards. O(1)
     */
    function guardsAmount() external view returns (uint256) {
        return guards.length();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Tries to add a new guard, emits an event only if guard was added.
     */
    function _addGuard(address _guard) internal override returns (bool guardAdded) {
        guardAdded = guards.add(_guard);
        if (guardAdded) {
            emit GuardAdded(_guard);
        }
    }

    /**
     * @notice Tries to remove a guard, emits an event only if guard was removed.
     */
    function _removeGuard(address _guard) internal override returns (bool guardRemoved) {
        guardRemoved = guards.remove(_guard);
        if (guardRemoved) {
            emit GuardRemoved(_guard);
        }
    }

    /**
     * @notice Returns whether given address is a guard.
     */
    function _isGuard(address _account) internal view override returns (bool) {
        return guards.contains(_account);
    }
}

abstract contract OriginHubEvents {
    /**
     * @notice Emitted when a correct report on a fraud attestation is submitted.
     * @param guard     Guard who signed the fraud report
     * @param report    Report data and signature
     */
    event CorrectFraudReport(address indexed guard, bytes report);

    /**
     * @notice Emitted when proof of an incorrect report is submitted.
     * @param guard     Guard who signed the incorrect report
     * @param report    Report data and signature
     */
    event IncorrectReport(address indexed guard, bytes report);

    /**
     * @notice Emitted when proof of an fraud attestation is submitted.
     * @param notary        Notary who signed fraud attestation
     * @param attestation   Attestation data and signature
     */
    event FraudAttestation(address indexed notary, bytes attestation);
}

abstract contract AttestationHubEvents {
    /**
     * @notice Emitted when an attestation is submitted to AttestationHub.
     * @param notary        Notary who signed the attestation
     * @param attestation   Raw payload with attestation data and notary signature
     */
    event AttestationAccepted(address indexed notary, bytes attestation);
}

abstract contract AttestationHub is AttestationHubEvents, AbstractNotaryRegistry {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Called by the external agent. Submits the signed attestation for handling.
     * @dev Reverts if either of this is true:
     *      - Attestation payload is not properly formatted.
     *      - Attestation signer is not a Notary.
     * @param _attestation  Payload with Attestation data and signature (see Attestation.sol)
     * @return TRUE if Attestation was handled correctly.
     */
    function submitAttestation(bytes memory _attestation) external returns (bool) {
        // Check if real Notary & signature.
        // This also checks if Attestation payload is properly formatted.
        (address _notary, bytes29 _attestationView) = _checkNotaryAuth(_attestation);
        // Pass _attestationView as the existing bytes29 pointer to attestation payload
        // Pass _attestation to avoid extra memory copy when emitting attestation payload
        return _handleAttestation(_notary, _attestationView, _attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Child contract should implement logic for handling the Attestation.
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _attestation      Payload with Attestation data and signature
     * @return TRUE if Attestation was handled correctly.
     */
    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal virtual returns (bool);
}

abstract contract ReportHub is AbstractGuardRegistry, AbstractNotaryRegistry {
    using Report for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Called by the external agent. Submits the signed report for handling.
     * @dev Reverts if either of this is true:
     *      - Report payload is not properly formatted.
     *      - Report signer is not a Guard.
     *      - Reported notary is not a Notary.
     * @param _report	Payload with Report data and signature
     * @return TRUE if Report was handled correctly.
     */
    function submitReport(bytes memory _report) external returns (bool) {
        // Check if real Guard & signature.
        // This also checks if Report payload is properly formatted.
        (address _guard, bytes29 _reportView) = _checkGuardAuth(_report);
        bytes29 _attestationView = _reportView.reportedAttestation();
        // Check if real Notary & signature.
        // This also checks if Attestation payload is properly formatted,
        // though it's already been checked in _checkGuardAuth(_report) [see Report.sol].
        address _notary = _checkNotaryAuth(_attestationView);
        // Pass _reportView as the existing bytes29 pointer to report payload.
        // Pass _report to avoid extra memory copy when emitting report payload.
        return _handleReport(_guard, _notary, _attestationView, _reportView, _report);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Implement logic for handling the Report in the child contracts.
     * Note: Report can have either Valid or Fraud flag, make sure to check that.
     * @param _guard            Guard address (signature&role already verified)
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _reportView       Memory view over Report for convenience
     * @param _report           Payload with Report data and signature
     * @return TRUE if Report was handled correctly.
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal virtual returns (bool);
}

library MerkleLib {
    uint256 internal constant TREE_DEPTH = 32;
    uint256 internal constant MAX_LEAVES = 2**TREE_DEPTH - 1;

    /**
     * @notice Struct representing incremental merkle tree. Contains the current branch,
     * while the number of inserted leaves are stored externally.
     **/
    // solhint-disable-next-line ordering
    struct Tree {
        bytes32[TREE_DEPTH] branch;
    }

    /**
     * @notice Inserts `_node` into merkle tree
     * @dev Reverts if tree is full
     * @param _newCount Amount of inserted leaves in the tree after the insertion (i.e. current + 1)
     * @param _node     Element to insert into tree
     **/
    function insert(
        Tree storage _tree,
        uint256 _newCount,
        bytes32 _node
    ) internal {
        require(_newCount <= MAX_LEAVES, "merkle tree full");
        // No need to increase _newCount here,
        // as it is already the amount of leaves after the insertion.
        for (uint256 i = 0; i < TREE_DEPTH; ) {
            if ((_newCount & 1) == 1) {
                _tree.branch[i] = _node;
                return;
            }
            _node = keccak256(abi.encodePacked(_tree.branch[i], _node));
            _newCount >>= 1;
            unchecked {
                ++i;
            }
        }
        // As the loop should always end prematurely with the `return` statement,
        // this code should be unreachable. We assert `false` just to be safe.
        assert(false);
    }

    /**
     * @notice Calculates and returns`_tree`'s current root given array of zero
     * hashes
     * @param _count    Current amount of inserted leaves in the tree
     * @param _zeroes   Array of zero hashes
     * @return _current Calculated root of `_tree`
     **/
    function rootWithCtx(
        Tree storage _tree,
        uint256 _count,
        bytes32[TREE_DEPTH] memory _zeroes
    ) internal view returns (bytes32 _current) {
        for (uint256 i = 0; i < TREE_DEPTH; ) {
            uint256 _ithBit = (_count >> i) & 0x01;
            if (_ithBit == 1) {
                _current = keccak256(abi.encodePacked(_tree.branch[i], _current));
            } else {
                _current = keccak256(abi.encodePacked(_current, _zeroes[i]));
            }
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Calculates and returns`_tree`'s current root
     * @param _count    Current amount of inserted leaves in the tree
     * @return Calculated root of `_tree`
     **/
    function root(Tree storage _tree, uint256 _count) internal view returns (bytes32) {
        return rootWithCtx(_tree, _count, zeroHashes());
    }

    /// @notice Returns array of TREE_DEPTH zero hashes
    /// @return _zeroes Array of TREE_DEPTH zero hashes
    function zeroHashes() internal pure returns (bytes32[TREE_DEPTH] memory _zeroes) {
        _zeroes[0] = Z_0;
        _zeroes[1] = Z_1;
        _zeroes[2] = Z_2;
        _zeroes[3] = Z_3;
        _zeroes[4] = Z_4;
        _zeroes[5] = Z_5;
        _zeroes[6] = Z_6;
        _zeroes[7] = Z_7;
        _zeroes[8] = Z_8;
        _zeroes[9] = Z_9;
        _zeroes[10] = Z_10;
        _zeroes[11] = Z_11;
        _zeroes[12] = Z_12;
        _zeroes[13] = Z_13;
        _zeroes[14] = Z_14;
        _zeroes[15] = Z_15;
        _zeroes[16] = Z_16;
        _zeroes[17] = Z_17;
        _zeroes[18] = Z_18;
        _zeroes[19] = Z_19;
        _zeroes[20] = Z_20;
        _zeroes[21] = Z_21;
        _zeroes[22] = Z_22;
        _zeroes[23] = Z_23;
        _zeroes[24] = Z_24;
        _zeroes[25] = Z_25;
        _zeroes[26] = Z_26;
        _zeroes[27] = Z_27;
        _zeroes[28] = Z_28;
        _zeroes[29] = Z_29;
        _zeroes[30] = Z_30;
        _zeroes[31] = Z_31;
    }

    /**
     * @notice Calculates and returns the merkle root for the given leaf
     * `_item`, a merkle branch, and the index of `_item` in the tree.
     * @param _item Merkle leaf
     * @param _branch Merkle proof
     * @param _index Index of `_item` in tree
     * @return _current Calculated merkle root
     **/
    function branchRoot(
        bytes32 _item,
        bytes32[TREE_DEPTH] memory _branch,
        uint256 _index
    ) internal pure returns (bytes32 _current) {
        _current = _item;

        for (uint256 i = 0; i < TREE_DEPTH; ) {
            uint256 _ithBit = (_index >> i) & 0x01;
            bytes32 _next = _branch[i];
            if (_ithBit == 1) {
                _current = keccak256(abi.encodePacked(_next, _current));
            } else {
                _current = keccak256(abi.encodePacked(_current, _next));
            }
            unchecked {
                ++i;
            }
        }
    }

    // keccak256 zero hashes
    bytes32 internal constant Z_0 =
        hex"0000000000000000000000000000000000000000000000000000000000000000";
    bytes32 internal constant Z_1 =
        hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";
    bytes32 internal constant Z_2 =
        hex"b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30";
    bytes32 internal constant Z_3 =
        hex"21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85";
    bytes32 internal constant Z_4 =
        hex"e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a19344";
    bytes32 internal constant Z_5 =
        hex"0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d";
    bytes32 internal constant Z_6 =
        hex"887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968";
    bytes32 internal constant Z_7 =
        hex"ffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f83";
    bytes32 internal constant Z_8 =
        hex"9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af";
    bytes32 internal constant Z_9 =
        hex"cefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0";
    bytes32 internal constant Z_10 =
        hex"f9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5";
    bytes32 internal constant Z_11 =
        hex"f8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf892";
    bytes32 internal constant Z_12 =
        hex"3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c";
    bytes32 internal constant Z_13 =
        hex"c1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb";
    bytes32 internal constant Z_14 =
        hex"5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc";
    bytes32 internal constant Z_15 =
        hex"da7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d2";
    bytes32 internal constant Z_16 =
        hex"2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f";
    bytes32 internal constant Z_17 =
        hex"e1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a";
    bytes32 internal constant Z_18 =
        hex"5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0";
    bytes32 internal constant Z_19 =
        hex"b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0";
    bytes32 internal constant Z_20 =
        hex"c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2";
    bytes32 internal constant Z_21 =
        hex"f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd9";
    bytes32 internal constant Z_22 =
        hex"5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e377";
    bytes32 internal constant Z_23 =
        hex"4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652";
    bytes32 internal constant Z_24 =
        hex"cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef";
    bytes32 internal constant Z_25 =
        hex"0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d";
    bytes32 internal constant Z_26 =
        hex"b8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0";
    bytes32 internal constant Z_27 =
        hex"838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e";
    bytes32 internal constant Z_28 =
        hex"662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e";
    bytes32 internal constant Z_29 =
        hex"388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea322";
    bytes32 internal constant Z_30 =
        hex"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735";
    bytes32 internal constant Z_31 =
        hex"8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9";
}

abstract contract OriginHub is
    OriginHubEvents,
    AttestationHub,
    ReportHub,
    DomainNotaryRegistry,
    GuardRegistry
{
    using Attestation for bytes29;
    using Report for bytes29;
    using TypedMemView for bytes29;

    using MerkleLib for MerkleLib.Tree;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // [destination domain] => [Merkle Tree containing all hashes of sent messages to that domain]
    mapping(uint32 => MerkleLib.Tree) internal trees;
    // [destination domain] => [Merkle tree roots after inserting a sent message to that domain]
    mapping(uint32 => bytes32[]) internal historicalRoots;

    // gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

    // Merkle root for an empty merkle tree.
    bytes32 internal constant EMPTY_TREE_ROOT =
        hex"27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757";

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Suggest attestation for the off-chain actors to sign for a specific destination.
     * Note: signing the suggested attestation will will never lead
     * to slashing of the actor, assuming they have confirmed that the block, where the merkle root
     * was updated, is not subject to reorganization (which is different for every observed chain).
     * @dev If no messages have been sent, following values are returned:
     * - nonce = 0
     * - root = 0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757
     * Which is the merkle root for an empty merkle tree.
     * @return latestNonce Current nonce
     * @return latestRoot  Current merkle root
     */
    function suggestAttestation(uint32 _destination)
        external
        view
        returns (uint32 latestNonce, bytes32 latestRoot)
    {
        latestNonce = nonce(_destination);
        latestRoot = getHistoricalRoot(_destination, latestNonce);
    }

    // TODO: add suggestAttestations() once OriginHub inherits from GlobalNotaryRegistry

    /**
     * @notice Returns a historical merkle root for the given destination.
     * Note: signing the attestation with the given historical root will never lead
     * to slashing of the actor, assuming they have confirmed that the block, where the merkle root
     * was updated, is not subject to reorganization (which is different for every observed chain).
     * @param _destination  Destination domain
     * @param _nonce        Historical nonce
     * @return Root for destination's merkle tree right after message to `_destination`
     * with `nonce = _nonce` was dispatched.
     */
    function getHistoricalRoot(uint32 _destination, uint32 _nonce) public view returns (bytes32) {
        // Check if destination is known
        if (historicalRoots[_destination].length > 0) {
            // Check if nonce exists
            require(_nonce < historicalRoots[_destination].length, "!nonce: existing destination");
            return historicalRoots[_destination][_nonce];
        } else {
            // If destination is unknown, we have the root of an empty merkle tree
            require(_nonce == 0, "!nonce: unknown destination");
            return EMPTY_TREE_ROOT;
        }
    }

    /**
     * @notice Returns nonce of the last inserted Merkle root for the given destination,
     * which is also the number of inserted leaves in the destination merkle tree (current index).
     */
    function nonce(uint32 _destination) public view returns (uint32 latestNonce) {
        latestNonce = uint32(_getTreeCount(_destination));
    }

    /**
     * @notice Calculates and returns tree's current root for the given destination.
     */
    function root(uint32 _destination) public view returns (bytes32) {
        return trees[_destination].root(_getTreeCount(_destination));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Checks if a submitted Attestation is a valid Attestation.
     * Attestation Flag can be either "Fraud" or "Valid".
     * A "Fraud" Attestation is a (_destination, _nonce, _root) attestation that doesn't correspond
     * with the historical state of Origin contract. Either of these needs to be true:
     * - _nonce is higher than current nonce for _destination (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce for _destination
     * This would mean that message(s) that were not truly
     * dispatched on Origin were falsely included in the signed root.
     *
     * A Fraud Attestation will only be accepted as valid by the Mirror.
     * If a Fraud Attestation is submitted to the Mirror, a Guard should
     * submit a Fraud Report using Origin.submitReport()
     * in order to slash the Notary with a Fraud Attestation.
     *
     * @dev Notary signature and role have been checked in AttestationHub,
     * hence `_notary` is an active Notary at this point.
     *
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _attestation      Payload with Attestation data and signature
     * @return isValid          TRUE if Attestation was valid (implying Notary was not slashed).
     */
    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal override returns (bool isValid) {
        uint32 attestedDestination = _attestationView.attestedDestination();
        uint32 attestedNonce = _attestationView.attestedNonce();
        bytes32 attestedRoot = _attestationView.attestedRoot();
        isValid = _isValidAttestation(attestedDestination, attestedNonce, attestedRoot);
        if (!isValid) {
            emit FraudAttestation(_notary, _attestation);
            // Guard doesn't receive anything, as Notary wasn't slashed using the Fraud Report
            _slashNotary({ _domain: _localDomain(), _notary: _notary, _guard: address(0) });
            /**
             * TODO: design incentives for the reporter in a way, where they get less
             * by reporting directly instead of using a correct Fraud Report.
             * That will allow Guards to focus on Report signing and don't worry
             * about submitReport (whether their own or outsourced) txs being frontrun.
             */
        }
    }

    /**
     * @notice Checks if a submitted Report is a correct Report. Reported Attestation
     * can be either valid or fraud. Report flag can also be either Valid or Fraud.
     * Report is correct if its flag matches the Attestation validity.
     * 1. Attestation: valid, Flag: Fraud.
     *      Report is deemed incorrect, Guard is slashed (if they haven't been already).
     * 2. Attestation: valid, Flag: Valid.
     *      Report is deemed correct, no action is done.
     * 3. Attestation: Fraud, Flag: Fraud.
     *      Report is deemed correct, Notary is slashed (if they haven't been already).
     * 4. Attestation: Fraud, Flag: Valid.
     *      Report is deemed incorrect, Guard is slashed (if they haven't been already).
     *      Notary is slashed (if they haven't been already), but Guard doesn't receive
     *      any rewards (as their report indicated that the attestation was valid).
     *
     * A "Fraud" Attestation is a (_destination, _nonce, _root) attestation that doesn't correspond
     * with the historical state of Origin contract. Either of these needs to be true:
     * - _nonce is higher than current nonce for _destination (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce for _destination
     * This would mean that message(s) that were not truly
     * dispatched on Origin were falsely included in the signed root.
     *
     * A Fraud Attestation will only be accepted as valid by the Mirror.
     * If a Fraud Attestation is submitted to the Mirror, a Guard should
     * submit a Fraud Report using Origin.submitReport()
     * in order to slash the Notary with a Fraud Attestation.
     *
     * @dev Both Notary and Guard signatures and roles have been checked in ReportHub,
     * hence `_notary` is an active Notary, `_guard` is an active Guard at this point.
     *
     * @param _guard            Guard address (signature&role already verified)
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation
     * @param _reportView       Memory view over Report
     * @param _report           Payload with Report data and signature
     * @return TRUE if Report was correct (implying Guard was not slashed)
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal override returns (bool) {
        uint32 attestedDestination = _attestationView.attestedDestination();
        uint32 attestedNonce = _attestationView.attestedNonce();
        bytes32 attestedRoot = _attestationView.attestedRoot();
        if (_isValidAttestation(attestedDestination, attestedNonce, attestedRoot)) {
            // Attestation: Valid
            if (_reportView.reportedFraud()) {
                // Flag: Fraud
                // Report is incorrect, slash the Guard
                emit IncorrectReport(_guard, _report);
                _slashGuard(_guard);
                return false;
            } else {
                // Flag: Valid
                // Report is correct, no action needed
                return true;
            }
        } else {
            // Attestation: Fraud
            if (_reportView.reportedFraud()) {
                // Flag: Fraud
                // Report is correct, slash the Notary
                emit CorrectFraudReport(_guard, _report);
                emit FraudAttestation(_notary, _attestationView.clone());
                _slashNotary({ _domain: _localDomain(), _notary: _notary, _guard: _guard });
                return true;
            } else {
                // Flag: Valid
                // Report is incorrect, slash the Guard
                emit IncorrectReport(_guard, _report);
                _slashGuard(_guard);
                emit FraudAttestation(_notary, _attestationView.clone());
                // Guard doesn't receive anything due to Valid flag on the Report
                _slashNotary({ _domain: _localDomain(), _notary: _notary, _guard: address(0) });
                return false;
            }
        }
    }

    /**
     * @notice Inserts a merkle root for an empty merkle tree into the historical roots array
     * for the given destination.
     * @dev This enables:
     * - Counting nonces from 1 (nonce=0 meaning no messages have been sent).
     * - Not slashing the Notaries for signing an attestation for an empty tree
     * (assuming they sign the correct root outlined below).
     */
    function _initializeHistoricalRoots(uint32 _destination) internal {
        // This function should only be called only if the array is empty
        assert(historicalRoots[_destination].length == 0);
        // Insert a historical root so nonces start at 1 rather then 0.
        // Here we insert the root of an empty merkle tree
        historicalRoots[_destination].push(EMPTY_TREE_ROOT);
    }

    /**
     * @notice Inserts new message into the Merkle tree for the given destination
     * and stores the new merkle root.
     * @param _destination  Destination domain of the dispatched message
     * @param _messageNonce Nonce of the dispatched message
     * @param _messageHash  Hash of the dispatched message
     */
    function _insertMessage(
        uint32 _destination,
        uint32 _messageNonce,
        bytes32 _messageHash
    ) internal {
        // TODO: when Notary is active on Destination, initialize historical roots
        // upon adding a first Notary for given destination
        if (historicalRoots[_destination].length == 0) _initializeHistoricalRoots(_destination);
        /// @dev _messageNonce == tree.count() + 1
        // tree.insert() requires amount of leaves AFTER the leaf insertion (i.e. tree.count() + 1)
        trees[_destination].insert(_messageNonce, _messageHash);
        /// @dev leaf is inserted => _messageNonce == tree.count()
        // tree.root() requires current amount of leaves (i.e. tree.count())
        historicalRoots[_destination].push(trees[_destination].root(_messageNonce));
    }

    /**
     * @notice Child contract should implement the slashing logic for Notaries
     * with all the required system calls.
     * @dev Called when fraud is proven (Fraud Attestation).
     * @param _domain   Domain where the reported Notary is active
     * @param _notary   Notary to slash
     * @param _guard    Guard who reported fraudulent Notary [address(0) if not a Guard report]
     */
    function _slashNotary(
        uint32 _domain,
        address _notary,
        address _guard
    ) internal virtual;

    /**
     * @notice Child contract should implement the slashing logic for Guards
     * with all the required system calls.
     * @dev Called when guard misbehavior is proven (Incorrect Report).
     * @param _guard    Guard to slash
     */
    function _slashGuard(address _guard) internal virtual;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns whether (_destination, _nonce, _root) matches the historical state
     * of the Merkle Tree for that destination.
     * @dev For `_nonce == 0`: root has to match `EMPTY_TREE_ROOT` (root of an empty merkle tree)
     * For `_nonce != 0`:
     * - There has to be at least `_nonce` messages sent to `_destination`
     * - Merkle root after sending message with `nonce == _nonce` should match `_root`
     */
    function _isValidAttestation(
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bool) {
        if (_nonce < historicalRoots[_destination].length) {
            // If a nonce exists for a given destination,
            // a root should match the historical root
            return _root == historicalRoots[_destination][_nonce];
        }
        // If a nonce doesn't exist for a given destination,
        // it should be a zero nonce with a root of an empty merkle tree
        return _nonce == 0 && _root == EMPTY_TREE_ROOT;
    }

    /**
     * @notice Returns amount of leaves in the merkle tree for the given destination.
     * @dev Every inserted leaf leads to adding a historical root,
     * removing the necessity to store amount of leaves separately.
     * Historical roots array is initialized with a root of an empty Merkle tree,
     * thus actual amount of leaves is lower by one.
     */
    function _getTreeCount(uint32 _destination) internal view returns (uint256) {
        // if no historical roots are saved, destination is unknown, and there were
        // no dispatched messages to that destination
        if (historicalRoots[_destination].length == 0) return 0;
        // We subtract 1, as the very first inserted root is EMPTY_TREE_ROOT
        return historicalRoots[_destination].length - 1;
    }
}

//

library TypeCasts {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    function coerceBytes32(string memory _s) internal pure returns (bytes32 _b) {
        _b = bytes(_s).ref(0).index(0, uint8(bytes(_s).length));
    }

    // treat it as a null-terminated string of max 32 bytes
    function coerceString(bytes32 _buf) internal pure returns (string memory _newStr) {
        uint8 _slen = 0;
        while (_slen < 32 && _buf[_slen] != 0) {
            _slen++;
        }

        // solhint-disable-next-line no-inline-assembly
        assembly {
            _newStr := mload(0x40)
            mstore(0x40, add(_newStr, 0x40)) // may end up with extra
            mstore(_newStr, _slen)
            mstore(add(_newStr, 0x20), _buf)
        }
    }

    // alignment preserving cast
    function addressToBytes32(address _addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(_addr)));
    }

    // alignment preserving cast
    function bytes32ToAddress(bytes32 _buf) internal pure returns (address) {
        return address(uint160(uint256(_buf)));
    }
}

library Header {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint16 internal constant HEADER_VERSION = 1;

    /**
     * @dev Header memory layout
     * [000 .. 002): version            uint16   2 bytes
     * [002 .. 006): origin             uint32   4 bytes
     * [006 .. 038): sender             bytes32 32 bytes
     * [038 .. 042): nonce              uint32   4 bytes
     * [042 .. 046): destination        uint32   4 bytes
     * [046 .. 078): recipient          bytes32 32 bytes
     * [078 .. 082): optimisticSeconds  uint32   4 bytes
     */

    uint256 internal constant OFFSET_VERSION = 0;
    uint256 internal constant OFFSET_ORIGIN = 2;
    uint256 internal constant OFFSET_SENDER = 6;
    uint256 internal constant OFFSET_NONCE = 38;
    uint256 internal constant OFFSET_DESTINATION = 42;
    uint256 internal constant OFFSET_RECIPIENT = 46;
    uint256 internal constant OFFSET_OPTIMISTIC_SECONDS = 78;

    uint256 internal constant HEADER_LENGTH = 82;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyHeader(bytes29 _view) {
        _view.assertType(SynapseTypes.MESSAGE_HEADER);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a header payload.
     */
    function castToHeader(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.MESSAGE_HEADER);
    }

    /**
     * @notice Returns a formatted Header payload with provided fields
     * @param _origin               Domain of origin chain
     * @param _sender               Address that sent the message
     * @param _nonce                Message nonce on origin chain
     * @param _destination          Domain of destination chain
     * @param _recipient            Address that will receive the message
     * @param _optimisticSeconds    Optimistic period for message execution
     * @return Formatted header
     **/
    function formatHeader(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked(
                HEADER_VERSION,
                _origin,
                _sender,
                _nonce,
                _destination,
                _recipient,
                _optimisticSeconds
            );
    }

    /**
     * @notice Checks that a payload is a formatted Header.
     */
    function isHeader(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Check if version exists in the payload
        if (length < 2) return false;
        // Check that header version and its length matches
        return headerVersion(_view) == HEADER_VERSION && length == HEADER_LENGTH;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            HEADER SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns header's version field.
    function headerVersion(bytes29 _header) internal pure onlyHeader(_header) returns (uint16) {
        return uint16(_header.indexUint(OFFSET_VERSION, 2));
    }

    /// @notice Returns header's origin field
    function origin(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_ORIGIN, 4));
    }

    /// @notice Returns header's sender field
    function sender(bytes29 _header) internal pure onlyHeader(_header) returns (bytes32) {
        return _header.index(OFFSET_SENDER, 32);
    }

    /// @notice Returns header's nonce field
    function nonce(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_NONCE, 4));
    }

    /// @notice Returns header's destination field
    function destination(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_DESTINATION, 4));
    }

    /// @notice Returns header's recipient field as bytes32
    function recipient(bytes29 _header) internal pure onlyHeader(_header) returns (bytes32) {
        return _header.index(OFFSET_RECIPIENT, 32);
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticSeconds(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_OPTIMISTIC_SECONDS, 4));
    }

    /// @notice Returns header's recipient field as an address
    function recipientAddress(bytes29 _header) internal pure returns (address) {
        return TypeCasts.bytes32ToAddress(recipient(_header));
    }
}

library Tips {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint16 internal constant TIPS_VERSION = 1;

    // TODO: determine if we need to pack the tips values,
    // or if using uint256 instead will suffice.

    /**
     * @dev Tips memory layout
     * [000 .. 002): version            uint16	 2 bytes
     * [002 .. 014): notaryTip          uint96	12 bytes
     * [014 .. 026): broadcasterTip     uint96	12 bytes
     * [026 .. 038): proverTip          uint96	12 bytes
     * [038 .. 050): executorTip        uint96	12 bytes
     */

    uint256 internal constant OFFSET_VERSION = 0;
    uint256 internal constant OFFSET_NOTARY = 2;
    uint256 internal constant OFFSET_BROADCASTER = 14;
    uint256 internal constant OFFSET_PROVER = 26;
    uint256 internal constant OFFSET_EXECUTOR = 38;

    uint256 internal constant TIPS_LENGTH = 50;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyTips(bytes29 _view) {
        _view.assertType(SynapseTypes.MESSAGE_TIPS);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a tips payload.
     */
    function castToTips(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.MESSAGE_TIPS);
    }

    /**
     * @notice Returns a formatted Tips payload with provided fields
     * @param _notaryTip        Tip for the Notary
     * @param _broadcasterTip   Tip for the Broadcaster
     * @param _proverTip        Tip for the Prover
     * @param _executorTip      Tip for the Executor
     * @return Formatted tips
     **/
    function formatTips(
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked(TIPS_VERSION, _notaryTip, _broadcasterTip, _proverTip, _executorTip);
    }

    /**
     * @notice Returns a formatted Tips payload specifying empty tips.
     * @return Formatted tips
     **/
    function emptyTips() internal pure returns (bytes memory) {
        return formatTips(0, 0, 0, 0);
    }

    /**
     * @notice Checks that a payload is a formatted Tips payload.
     */
    function isTips(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Check if version exists in the payload
        if (length < 2) return false;
        // Check that header version and its length matches
        return tipsVersion(_view) == TIPS_VERSION && length == TIPS_LENGTH;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TIPS SLICING                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns version of formatted tips
    function tipsVersion(bytes29 _tips) internal pure onlyTips(_tips) returns (uint16) {
        return uint16(_tips.indexUint(OFFSET_VERSION, 2));
    }

    /// @notice Returns notaryTip field
    function notaryTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint96(_tips.indexUint(OFFSET_NOTARY, 12));
    }

    /// @notice Returns broadcasterTip field
    function broadcasterTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint96(_tips.indexUint(OFFSET_BROADCASTER, 12));
    }

    /// @notice Returns proverTip field
    function proverTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint96(_tips.indexUint(OFFSET_PROVER, 12));
    }

    /// @notice Returns executorTip field
    function executorTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint96(_tips.indexUint(OFFSET_EXECUTOR, 12));
    }

    /// @notice Returns total tip amount.
    function totalTips(bytes29 _tips) internal pure returns (uint96) {
        // In practice there's no chance that the total tips value would not fit into uint96.
        // TODO: determine if we want to use uint256 here instead anyway.
        return notaryTip(_tips) + broadcasterTip(_tips) + proverTip(_tips) + executorTip(_tips);
    }
}

library Message {
    using Header for bytes29;
    using Tips for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    enum Parts {
        Version,
        Header,
        Tips,
        Body
    }

    /**
     * @dev This is only updated if the whole message structure is changed,
     *      i.e. if a new part is added.
     *      If already existing part is changed, the message version does not get bumped.
     */
    uint16 internal constant MESSAGE_VERSION = 1;

    /**
     * @dev Message memory layout
     * [000 .. 002): version            uint16  2 bytes
     * [002 .. 004): header length      uint16  2 bytes (length == AAA - 6)
     * [004 .. 006): tips length        uint16  2 bytes (length == BBB - AAA)
     * [006 .. AAA): header             bytes   ? bytes
     * [AAA .. BBB): tips               bytes   ? bytes
     * [BBB .. CCC): body               bytes   ? bytes (length could be zero)
     */

    uint256 internal constant OFFSET_VERSION = 0;

    /// @dev How much bytes is used for storing the version, or a single offset value
    uint8 internal constant TWO_BYTES = 2;
    /// @dev This value reflects the header offset in the latest message version
    uint16 internal constant OFFSET_HEADER = TWO_BYTES * uint8(type(Parts).max);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyMessage(bytes29 _view) {
        _view.assertType(SynapseTypes.MESSAGE);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a message payload.
     */
    function castToMessage(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.MESSAGE);
    }

    /**
     * @notice Returns formatted message with provided fields
     * @param _header       Formatted header payload
     * @param _tips         Formatted tips payload
     * @param _messageBody  Raw bytes of message body
     * @return Formatted message
     **/
    function formatMessage(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes memory) {
        // Header and Tips are supposed to fit within 65535 bytes
        return
            abi.encodePacked(
                MESSAGE_VERSION,
                uint16(_header.length),
                uint16(_tips.length),
                _header,
                _tips,
                _messageBody
            );
    }

    /**
     * @notice Returns formatted message with provided fields
     * @param _origin               Domain of origin chain
     * @param _sender               Address that sent the message
     * @param _nonce                Message nonce on origin chain
     * @param _destination          Domain of destination chain
     * @param _recipient            Address that will receive the message
     * @param _optimisticSeconds    Optimistic period for message execution
     * @param _tips                 Formatted tips payload
     * @param _messageBody          Raw bytes of message body
     * @return Formatted message
     **/
    function formatMessage(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes memory) {
        return
            formatMessage(
                Header.formatHeader(
                    _origin,
                    _sender,
                    _nonce,
                    _destination,
                    _recipient,
                    _optimisticSeconds
                ),
                _tips,
                _messageBody
            );
    }

    /**
     * @notice Checks that a payload is a formatted Message.
     */
    function isMessage(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Check if version and lengths exist in the payload
        if (length < OFFSET_HEADER) return false;
        // Check message version
        if (messageVersion(_view) != MESSAGE_VERSION) return false;

        uint256 headerLength = _loadLength(_view, Parts.Header);
        uint256 tipsLength = _loadLength(_view, Parts.Tips);
        // Header and Tips need to exist
        // Body could be empty, thus >
        if (OFFSET_HEADER + headerLength + tipsLength > length) return false;

        // Check header for being a formatted header payload
        // Check tips for being a formatted tips payload
        if (!header(_view).isHeader() || !tips(_view).isTips()) return false;
        return true;
    }

    /**
     * @notice Returns leaf of formatted message with provided fields.
     * @param _header       Formatted header payload
     * @param _tips         Formatted tips payload
     * @param _messageBody  Raw bytes of message body
     * @return Leaf (hash) of formatted message
     **/
    function messageHash(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes32) {
        return keccak256(formatMessage(_header, _tips, _messageBody));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           MESSAGE SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns message's version field.
    function messageVersion(bytes29 _view) internal pure onlyMessage(_view) returns (uint16) {
        return uint16(_view.indexUint(OFFSET_VERSION, 2));
    }

    /// @notice Returns message's header field as bytes29 pointer.
    function header(bytes29 _view) internal pure onlyMessage(_view) returns (bytes29) {
        return
            _view.slice(
                OFFSET_HEADER,
                _loadLength(_view, Parts.Header),
                SynapseTypes.MESSAGE_HEADER
            );
    }

    /// @notice Returns message's tips field as bytes29 pointer.
    function tips(bytes29 _view) internal pure onlyMessage(_view) returns (bytes29) {
        return
            _view.slice(
                OFFSET_HEADER + _loadLength(_view, Parts.Header),
                _loadLength(_view, Parts.Tips),
                SynapseTypes.MESSAGE_TIPS
            );
    }

    /// @notice Returns message's body field as bytes29 pointer.
    function body(bytes29 _view) internal pure onlyMessage(_view) returns (bytes29) {
        return
            _view.sliceFrom(
                OFFSET_HEADER + _loadLength(_view, Parts.Header) + _loadLength(_view, Parts.Tips),
                SynapseTypes.RAW_BYTES
            );
    }

    /// @notice Loads length for a given part of the message
    function _loadLength(bytes29 _view, Parts _part) private pure returns (uint256) {
        return _view.indexUint(uint256(_part) * TWO_BYTES, TWO_BYTES);
    }
}

library SystemCall {
    using ByteString for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev Custom address, used for sending and receiving system messages.
     *      Origin is supposed to dispatch messages from SystemRouter
     *      as if they were sent by this address.
     *      Destination is supposed to reroute messages for this address to SystemRouter.
     *
     *      Note: all bits except for lower 20 bytes are set to 1.
     *      Note: TypeCasts.bytes32ToAddress(SYSTEM_ROUTER) == address(0)
     */
    bytes32 internal constant SYSTEM_ROUTER = bytes32(type(uint256).max << 160);

    /**
     * @dev SystemCall memory layout
     * [000 .. 001): recipient      uint8   1 bytes
     * [001 .. END]: payload        bytes   ? bytes
     */

    uint256 internal constant OFFSET_CALL_RECIPIENT = 0;
    uint256 internal constant OFFSET_CALL_PAYLOAD = 1;

    /**
     * @dev System Router is supposed to append (origin, caller) to the given payload,
     * meaning for a valid system call payload there have to exist at least two arguments,
     * occupying at least two words in total.
     */
    uint256 internal constant PAYLOAD_MIN_ARGUMENT_WORDS = 2;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyType(bytes29 _view, uint40 _type) {
        _view.assertType(_type);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a system call payload.
     */
    function castToSystemCall(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.SYSTEM_CALL);
    }

    /**
     * @notice Returns a formatted System Call payload with provided fields
     * @param _systemRecipient  System Contract to receive message
     *                          (see ISystemRouter.SystemEntity)
     * @param _payload          Payload for call on destination chain
     * @return Formatted System Call
     **/
    function formatSystemCall(uint8 _systemRecipient, bytes memory _payload)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_systemRecipient, _payload);
    }

    /**
     * @notice Checks that a payload is a formatted System Call.
     */
    function isSystemCall(bytes29 _view) internal pure returns (bool) {
        // Payload needs to exist (system calls are never done via fallback function)
        if (_view.len() < OFFSET_CALL_PAYLOAD) return false;
        bytes29 payload = _callPayload(_view);
        // Payload needs to be a proper call payload
        if (!payload.isCallPayload()) return false;
        // Payload needs to have at least this amount of argument words
        return payload.argumentWords() >= PAYLOAD_MIN_ARGUMENT_WORDS;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         SYSTEM CALL SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns int value of System Call's recipient (see ISystemRouter.SystemEntity).
     */
    function callRecipient(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SYSTEM_CALL)
        returns (uint8)
    {
        return uint8(_view.indexUint({ _index: OFFSET_CALL_RECIPIENT, _bytes: 1 }));
    }

    /**
     * @notice Returns System Call's payload.
     */
    function callPayload(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SYSTEM_CALL)
        returns (bytes29)
    {
        return _callPayload(_view);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns System Call's payload WITHOUT checking the view type.
     * To be used in `isSystemCall`, where type check is not necessary.
     */
    function _callPayload(bytes29 _view) private pure returns (bytes29) {
        return _view.sliceFrom({ _index: OFFSET_CALL_PAYLOAD, newType: SynapseTypes.CALL_PAYLOAD });
    }
}

interface ISystemRouter {
    /// @dev Potential senders/recipients of a system message
    enum SystemEntity {
        Origin,
        Destination
    }

    /**
     * @notice Call a System Contract on the destination chain with a given data payload.
     * Note: for system calls on the local chain
     * - use `destination = localDomain`
     * - `_optimisticSeconds` value will be ignored
     *
     * @dev Only System contracts are allowed to call this function.
     * Note: knowledge of recipient address is not required, routing will be done by SystemRouter
     * on the destination chain. Following call will be made on destination chain:
     * - recipient.call(_data, callOrigin, systemCaller, rootSubmittedAt)
     * This allows recipient to check:
     * - callOrigin: domain where a system call originated (local domain in this case)
     * - systemCaller: system entity who initiated the call (msg.sender on local chain)
     * - rootSubmittedAt:
     *   - For cross-chain calls: timestamp when merkle root (used for executing the system call)
     *     was submitted to destination and its optimistic timer started ticking
     *   - For on-chain calls: timestamp of the current block
     *
     * @param _destination          Domain of destination chain
     * @param _optimisticSeconds    Optimistic period for the message
     * @param _recipient            System entity to receive the call on destination chain
     * @param _data                 Data for calling recipient on destination chain
     */
    function systemCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity _recipient,
        bytes memory _data
    ) external;

    /**
     * @notice Calls a few system contracts using the given calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes[] memory _dataArray
    ) external;

    /**
     * @notice Calls a few system contracts using the same calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes memory _data
    ) external;

    /**
     * @notice Calls a single system contract a few times using the given calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity _recipient,
        bytes[] memory _dataArray
    ) external;
}

// 
// OpenZeppelin Contracts (last updated v4.5.0) (utils/Address.sol)
/**
 * @dev Collection of functions related to the address type
 */
library AddressUpgradeable {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     * ====
     *
     * [IMPORTANT]
     * ====
     * You shouldn't rely on `isContract` to protect against flash loan attacks!
     *
     * Preventing calls from contracts is highly discouraged. It breaks composability, breaks support for smart wallets
     * like Gnosis Safe, and does not provide security since it can be circumvented by calling from a contract
     * constructor.
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://diligence.consensys.net/posts/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.5.11/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain `call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCall(target, data, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        require(isContract(target), "Address: call to non-contract");

        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        require(isContract(target), "Address: static call to non-contract");

        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verifies that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason using the provided one.
     *
     * _Available since v4.3._
     */
    function verifyCallResult(
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        if (success) {
            return returndata;
        } else {
            // Look for revert reason and bubble it up if present
            if (returndata.length > 0) {
                // The easiest way to bubble the revert reason is using memory via assembly

                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert(errorMessage);
            }
        }
    }
}

// 
// OpenZeppelin Contracts (last updated v4.6.0) (proxy/utils/Initializable.sol)
/**
 * @dev This is a base contract to aid in writing upgradeable contracts, or any kind of contract that will be deployed
 * behind a proxy. Since proxied contracts do not make use of a constructor, it's common to move constructor logic to an
 * external initializer function, usually called `initialize`. It then becomes necessary to protect this initializer
 * function so it can only be called once. The {initializer} modifier provided by this contract will have this effect.
 *
 * The initialization functions use a version number. Once a version number is used, it is consumed and cannot be
 * reused. This mechanism prevents re-execution of each "step" but allows the creation of new initialization steps in
 * case an upgrade adds a module that needs to be initialized.
 *
 * For example:
 *
 * [.hljs-theme-light.nopadding]
 * ```
 * contract MyToken is ERC20Upgradeable {
 *     function initialize() initializer public {
 *         __ERC20_init("MyToken", "MTK");
 *     }
 * }
 * contract MyTokenV2 is MyToken, ERC20PermitUpgradeable {
 *     function initializeV2() reinitializer(2) public {
 *         __ERC20Permit_init("MyToken");
 *     }
 * }
 * ```
 *
 * TIP: To avoid leaving the proxy in an uninitialized state, the initializer function should be called as early as
 * possible by providing the encoded function call as the `_data` argument to {ERC1967Proxy-constructor}.
 *
 * CAUTION: When used with inheritance, manual care must be taken to not invoke a parent initializer twice, or to ensure
 * that all initializers are idempotent. This is not verified automatically as constructors are by Solidity.
 *
 * [CAUTION]
 * ====
 * Avoid leaving a contract uninitialized.
 *
 * An uninitialized contract can be taken over by an attacker. This applies to both a proxy and its implementation
 * contract, which may impact the proxy. To prevent the implementation contract from being used, you should invoke
 * the {_disableInitializers} function in the constructor to automatically lock it when it is deployed:
 *
 * [.hljs-theme-light.nopadding]
 * ```
 * /// @custom:oz-upgrades-unsafe-allow constructor
 * constructor() {
 *     _disableInitializers();
 * }
 * ```
 * ====
 */
abstract contract Initializable {
    /**
     * @dev Indicates that the contract has been initialized.
     * @custom:oz-retyped-from bool
     */
    uint8 private _initialized;

    /**
     * @dev Indicates that the contract is in the process of being initialized.
     */
    bool private _initializing;

    /**
     * @dev Triggered when the contract has been initialized or reinitialized.
     */
    event Initialized(uint8 version);

    /**
     * @dev A modifier that defines a protected initializer function that can be invoked at most once. In its scope,
     * `onlyInitializing` functions can be used to initialize parent contracts. Equivalent to `reinitializer(1)`.
     */
    modifier initializer() {
        bool isTopLevelCall = _setInitializedVersion(1);
        if (isTopLevelCall) {
            _initializing = true;
        }
        _;
        if (isTopLevelCall) {
            _initializing = false;
            emit Initialized(1);
        }
    }

    /**
     * @dev A modifier that defines a protected reinitializer function that can be invoked at most once, and only if the
     * contract hasn't been initialized to a greater version before. In its scope, `onlyInitializing` functions can be
     * used to initialize parent contracts.
     *
     * `initializer` is equivalent to `reinitializer(1)`, so a reinitializer may be used after the original
     * initialization step. This is essential to configure modules that are added through upgrades and that require
     * initialization.
     *
     * Note that versions can jump in increments greater than 1; this implies that if multiple reinitializers coexist in
     * a contract, executing them in the right order is up to the developer or operator.
     */
    modifier reinitializer(uint8 version) {
        bool isTopLevelCall = _setInitializedVersion(version);
        if (isTopLevelCall) {
            _initializing = true;
        }
        _;
        if (isTopLevelCall) {
            _initializing = false;
            emit Initialized(version);
        }
    }

    /**
     * @dev Modifier to protect an initialization function so that it can only be invoked by functions with the
     * {initializer} and {reinitializer} modifiers, directly or indirectly.
     */
    modifier onlyInitializing() {
        require(_initializing, "Initializable: contract is not initializing");
        _;
    }

    /**
     * @dev Locks the contract, preventing any future reinitialization. This cannot be part of an initializer call.
     * Calling this in the constructor of a contract will prevent that contract from being initialized or reinitialized
     * to any version. It is recommended to use this to lock implementation contracts that are designed to be called
     * through proxies.
     */
    function _disableInitializers() internal virtual {
        _setInitializedVersion(type(uint8).max);
    }

    function _setInitializedVersion(uint8 version) private returns (bool) {
        // If the contract is initializing we ignore whether _initialized is set in order to support multiple
        // inheritance patterns, but we only do this in the context of a constructor, and for the lowest level
        // of initializers, because in other contexts the contract may have been reentered.
        if (_initializing) {
            require(
                version == 1 && !AddressUpgradeable.isContract(address(this)),
                "Initializable: contract is already initialized"
            );
            return false;
        } else {
            require(_initialized < version, "Initializable: contract is already initialized");
            _initialized = version;
            return true;
        }
    }
}

// 
// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)
/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract ContextUpgradeable is Initializable {
    function __Context_init() internal onlyInitializing {
    }

    function __Context_init_unchained() internal onlyInitializing {
    }
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}

abstract contract OwnableUpgradeable is Initializable, ContextUpgradeable {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    function __Ownable_init() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    function __Ownable_init_unchained() internal onlyInitializing {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}

abstract contract SystemContract is DomainContext, OwnableUpgradeable {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // domain of the Synapse Chain
    // Answer to the Ultimate Question of Life, the Universe, and Everything
    // And answer to less important questions wink wink
    uint32 public constant SYNAPSE_DOMAIN = 4269;
    // TODO: replace the placeholder with actual value

    uint256 internal constant ORIGIN = 1 << uint8(ISystemRouter.SystemEntity.Origin);
    uint256 internal constant DESTINATION = 1 << uint8(ISystemRouter.SystemEntity.Destination);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    ISystemRouter public systemRouter;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * System Contracts on all chains (either local or remote).
     * Note: any function protected by this modifier should have last three params:
     * - uint32 _callOrigin
     * - SystemEntity _systemCaller
     * - uint256 _rootSubmittedAt
     * Make sure to check domain/caller, if a function should be only called
     * from a given domain / by a given caller.
     * Make sure to check that a needed amount of time has passed since
     * root submission for the cross-chain calls.
     */
    modifier onlySystemRouter() {
        _assertSystemRouter();
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * System Contracts on Synapse chain.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     */
    modifier onlySynapseChain(uint32 _originDomain) {
        require(_originDomain == SYNAPSE_DOMAIN, "!synapseDomain");
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * a set of System Contracts on any chain.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     * Note: check constants section for existing mask constants
     * E.g. to restrict the set of callers to three allowed system callers:
     *  onlyCallers(MASK_0 | MASK_1 | MASK_2, _systemCaller)
     */
    modifier onlyCallers(uint256 _allowedMask, ISystemRouter.SystemEntity _systemCaller) {
        require(_entityAllowed(_allowedMask, _systemCaller), "!allowedCaller");
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * System Contracts on remote chain with a defined minimum optimistic period.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     * Note: message could be sent with a period lower than that, but will be executed
     * only when `_optimisticSeconds` have passed.
     * Note: _optimisticSeconds=0 will allow calls from a local chain as well
     */
    modifier onlyOptimisticPeriodOver(uint256 _rootSubmittedAt, uint256 _optimisticSeconds) {
        _assertOptimisticPeriodOver(_rootSubmittedAt, _optimisticSeconds);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line func-name-mixedcase
    function __SystemContract_initialize() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line ordering
    function setSystemRouter(ISystemRouter _systemRouter) external onlyOwner {
        systemRouter = _systemRouter;
    }

    /**
     * @dev Should be impossible to renounce ownership;
     * we override OpenZeppelin OwnableUpgradeable's
     * implementation of renounceOwnership to make it a no-op
     */
    function renounceOwnership() public override onlyOwner {} //solhint-disable-line no-empty-blocks

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _assertSystemRouter() internal view {
        require(msg.sender == address(systemRouter), "!systemRouter");
    }

    function _assertOptimisticPeriodOver(uint256 _rootSubmittedAt, uint256 _optimisticSeconds)
        internal
        view
    {
        require(block.timestamp >= _rootSubmittedAt + _optimisticSeconds, "!optimisticPeriod");
    }

    /**
     * @notice Checks if a given entity is allowed to call a function using a _systemMask
     * @param _systemMask a mask of allowed entities
     * @param _entity a system entity to check
     * @return true if _entity is allowed to call a function
     *
     * @dev this function works by converting the enum value to a non-zero bit mask
     * we then use a bitwise AND operation to check if permission bits allow the entity
     * to perform this operation, more details can be found here:
     * https://en.wikipedia.org/wiki/Bitwise_operation#AND
     */
    function _entityAllowed(uint256 _systemMask, ISystemRouter.SystemEntity _entity)
        internal
        pure
        returns (bool)
    {
        return _systemMask & _getSystemMask(_entity) != 0;
    }

    /**
     * @notice Returns a mask for a given system entity
     * @param _entity System entity
     * @return a non-zero mask for a given system entity
     *
     * Converts an enum value into a non-zero bit mask used for a bitwise AND check
     * E.g. for Origin (0) returns 1, for Destination (1) returns 2
     */
    function _getSystemMask(ISystemRouter.SystemEntity _entity) internal pure returns (uint256) {
        return 1 << uint8(_entity);
    }
}

library Address {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     * ====
     *
     * [IMPORTANT]
     * ====
     * You shouldn't rely on `isContract` to protect against flash loan attacks!
     *
     * Preventing calls from contracts is highly discouraged. It breaks composability, breaks support for smart wallets
     * like Gnosis Safe, and does not provide security since it can be circumvented by calling from a contract
     * constructor.
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://diligence.consensys.net/posts/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.5.11/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain `call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCall(target, data, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        require(isContract(target), "Address: call to non-contract");

        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        require(isContract(target), "Address: static call to non-contract");

        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionDelegateCall(target, data, "Address: low-level delegate call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(isContract(target), "Address: delegate call to non-contract");

        (bool success, bytes memory returndata) = target.delegatecall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verifies that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason using the provided one.
     *
     * _Available since v4.3._
     */
    function verifyCallResult(
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        if (success) {
            return returndata;
        } else {
            // Look for revert reason and bubble it up if present
            if (returndata.length > 0) {
                // The easiest way to bubble the revert reason is using memory via assembly
                /// @solidity memory-safe-assembly
                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert(errorMessage);
            }
        }
    }
}

contract Origin is Version0, OriginEvents, SystemContract, LocalDomainContext, OriginHub {
    using Tips for bytes;
    using Tips for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Maximum bytes per message = 2 KiB
    // (somewhat arbitrarily set to begin)
    uint256 public constant MAX_MESSAGE_BODY_BYTES = 2 * 2**10;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // contract responsible for Notary bonding, slashing and rotation
    // TODO: use "bonding manager" instead when implemented
    INotaryManager public notaryManager;

    // gap for upgrade safety
    uint256[49] private __GAP; //solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Ensures that function is called by the NotaryManager contract
     */
    modifier onlyNotaryManager() {
        require(msg.sender == address(notaryManager), "!notaryManager");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) LocalDomainContext(_domain) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize(INotaryManager _notaryManager) external initializer {
        __SystemContract_initialize();
        _setNotaryManager(_notaryManager);
        _addNotary(notaryManager.notary());
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    EXTERNAL FUNCTIONS: RESTRICTED                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Set a new Notary
     * @dev To be set when rotating Notary after Fraud
     * @param _notary the new Notary
     */
    function setNotary(address _notary) external onlyNotaryManager {
        /**
         * TODO: do this properly
         * @dev 1. New Notaries should be added to all System Contracts
         *      from "secondary" Bonding contracts (global Notary/Guard registry)
         *      1a. onlyNotaryManager -> onlyBondingManager (or w/e the name would be)
         *      2. There is supposed to be more than one active Notary
         *      2a. setNotary() -> addNotary()
         */
        _addNotary(_notary);
    }

    /**
     * @notice Set a new NotaryManager contract
     * @dev Origin(s) will initially be initialized using a trusted NotaryManager contract;
     * we will progressively decentralize by swapping the trusted contract with a new implementation
     * that implements Notary bonding & slashing, and rules for Notary selection & rotation
     * @param _notaryManager the new NotaryManager contract
     */
    function setNotaryManager(address _notaryManager) external onlyOwner {
        _setNotaryManager(INotaryManager(_notaryManager));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Dispatch the message to the destination domain & recipient
     * @dev Format the message, insert its hash into Merkle tree,
     * enqueue the new Merkle root, and emit `Dispatch` event with message information.
     * @param _destination      Domain of destination chain
     * @param _recipient        Address of recipient on destination chain as bytes32
     * @param _messageBody      Raw bytes content of message
     */
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable haveActiveNotary returns (uint32 messageNonce, bytes32 messageHash) {
        // TODO: add unit tests covering return values
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        bytes29 tips = _tips.castToTips();
        // Check: tips payload is correctly formatted
        require(tips.isTips(), "!tips: formatting");
        // Check: total tips value matches msg.value
        require(tips.totalTips() == msg.value, "!tips: totalTips");
        // Latest nonce (i.e. "last message" nonce) is current amount of leaves in the tree.
        // Message nonce is the amount of leaves after the new leaf insertion
        messageNonce = nonce(_destination) + 1;
        // format the message into packed bytes
        bytes memory message = Message.formatMessage({
            _origin: _localDomain(),
            _sender: _checkForSystemRouter(_recipient),
            _nonce: messageNonce,
            _destination: _destination,
            _recipient: _recipient,
            _optimisticSeconds: _optimisticSeconds,
            _tips: _tips,
            _messageBody: _messageBody
        });
        messageHash = keccak256(message);
        // insert the hashed message into the Merkle tree
        _insertMessage(_destination, messageNonce, messageHash);
        // Emit Dispatch event with message information
        // note: leaf index in the tree is messageNonce - 1, meaning we don't need to emit that
        emit Dispatch(messageHash, messageNonce, _destination, _tips, message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Set the NotaryManager
     * @param _notaryManager Address of the NotaryManager
     */
    function _setNotaryManager(INotaryManager _notaryManager) internal {
        require(Address.isContract(address(_notaryManager)), "!contract notaryManager");
        notaryManager = INotaryManager(_notaryManager);
        emit NewNotaryManager(address(_notaryManager));
    }

    /**
     * @notice Slash the Notary.
     * @dev Called when fraud is proven (Fraud Attestation).
     * @param _notary   Notary to slash
     * @param _guard    Guard who reported fraudulent Notary [address(0) if not a Guard report]
     */
    function _slashNotary(
        uint32 _domain,
        address _notary,
        address _guard
    ) internal override {
        // _notary is always an active Notary at this point
        _removeNotary(_domain, _notary);
        notaryManager.slashNotary(payable(msg.sender));
        // TODO: add domain to the event (decide what fields need to be indexed)
        emit NotarySlashed(_notary, _guard, msg.sender);
    }

    /**
     * @notice Slash the Guard.
     * @dev Called when guard misbehavior is proven (Incorrect Report).
     * @param _guard    Guard to slash
     */
    function _slashGuard(address _guard) internal override {
        // _guard is always an active Guard at this point
        _removeGuard(_guard);
        emit GuardSlashed(_guard, msg.sender);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns adjusted "sender" field.
     * @dev By default, "sender" field is msg.sender address casted to bytes32.
     * However, if SYSTEM_ROUTER is used for "recipient" field, and msg.sender is SystemRouter,
     * SYSTEM_ROUTER is also used as "sender" field.
     * Note: tx will revert if anyone but SystemRouter uses SYSTEM_ROUTER as the recipient.
     */
    function _checkForSystemRouter(bytes32 _recipient) internal view returns (bytes32 sender) {
        if (_recipient != SystemCall.SYSTEM_ROUTER) {
            sender = TypeCasts.addressToBytes32(msg.sender);
            /**
             * @dev Note: SYSTEM_ROUTER has only the highest 12 bytes set,
             * whereas TypeCasts.addressToBytes32 sets only the lowest 20 bytes.
             * Thus, in this branch: sender != SystemCall.SYSTEM_ROUTER
             */
        } else {
            // Check that SystemRouter specified SYSTEM_ROUTER as recipient, revert otherwise.
            _assertSystemRouter();
            // Adjust "sender" field for correct processing on remote chain.
            sender = SystemCall.SYSTEM_ROUTER;
        }
    }
}

abstract contract NotaryManagerEvents {
    /**
     * @notice Emitted when a new origin is set
     * @param origin The address of the new origin contract
     */
    event NewOrigin(address origin);

    /**
     * @notice Emitted when a new notary is set
     * @param notary The address of the new notary
     */
    event NewNotary(address notary);

    /**
     * @notice Emitted when slashNotary is called
     */
    event FakeSlashed(address reporter);
}

// 
// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)
/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}

// 
// ============ Internal Imports ============
// ============ External Imports ============
/**
 * @title NotaryManager
 * @author Illusory Systems Inc.
 * @notice MVP / centralized version of contract
 * that will manage Notary bonding, slashing,
 * selection and rotation
 */
contract NotaryManager is NotaryManagerEvents, INotaryManager, Ownable {
    // ============ Public Storage ============

    // address of origin contract
    address public origin;

    // ============ Private Storage ============

    // address of the current notary
    address private _notary;

    // ============ Modifiers ============

    /**
     * @notice Require that the function is called
     * by the Origin contract
     */
    modifier onlyOrigin() {
        require(msg.sender == origin, "!origin");
        _;
    }

    // ============ Constructor ============

    constructor(address _notaryAddress) payable Ownable() {
        _notary = _notaryAddress;
    }

    // ============ External Functions ============

    /**
     * @notice Set the address of the a new origin contract
     * @dev only callable by trusted owner
     * @param _origin The address of the new origin contract
     */
    function setOrigin(address _origin) external onlyOwner {
        require(Address.isContract(_origin), "!contract origin");
        origin = _origin;

        emit NewOrigin(_origin);
    }

    /**
     * @notice Set the address of a new notary
     * @dev only callable by trusted owner
     * @param _notaryAddress The address of the new notary
     */
    function setNotary(address _notaryAddress) external onlyOwner {
        _notary = _notaryAddress;
        Origin(origin).setNotary(_notaryAddress);
        emit NewNotary(_notaryAddress);
    }

    /**
     * @notice Slashes the notary
     * @dev Currently does nothing, functionality will be implemented later
     * when notary bonding and rotation are also implemented
     * @param _reporter The address of the entity that reported the notary fraud
     */
    function slashNotary(address payable _reporter) external override onlyOrigin {
        emit FakeSlashed(_reporter);
    }

    /**
     * @notice Get address of current notary
     * @return the notary address
     */
    function notary() external view override returns (address) {
        return _notary;
    }

    /**
     * @dev should be impossible to renounce ownership;
     * we override OpenZeppelin Ownable implementation
     * of renounceOwnership to make it a no-op
     */
    // solhint-disable-next-line no-empty-blocks
    function renounceOwnership() public override onlyOwner {
        // do nothing
    }
}