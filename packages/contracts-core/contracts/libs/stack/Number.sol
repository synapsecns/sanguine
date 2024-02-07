// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// Number is a compact representation of uint256, that is fit into 16 bits
/// with the maximum relative error under 0.4%.
type Number is uint16;

using NumberLib for Number global;

/// # Number
/// Library for compact representation of uint256 numbers.
/// - Number is stored using mantissa and exponent, each occupying 8 bits.
/// - Numbers under 2**8 are stored as `mantissa` with `exponent = 0xFF`.
/// - Numbers at least 2**8 are approximated as `(256 + mantissa) << exponent`
/// > - `0 <= mantissa < 256`
/// > - `0 <= exponent <= 247` (`256 * 2**248` doesn't fit into uint256)
/// # Number stack layout (from highest bits to lowest)
///
/// | Position   | Field    | Type  | Bytes |
/// | ---------- | -------- | ----- | ----- |
/// | (002..001] | mantissa | uint8 | 1     |
/// | (001..000] | exponent | uint8 | 1     |

library NumberLib {
    /// @dev Amount of bits to shift to mantissa field
    uint16 private constant SHIFT_MANTISSA = 8;

    /// @notice For bwad math (binary wad) we use 2**64 as "wad" unit.
    /// @dev We are using not using 10**18 as wad, because it is not stored precisely in NumberLib.
    uint256 internal constant BWAD_SHIFT = 64;
    uint256 internal constant BWAD = 1 << BWAD_SHIFT;
    /// @notice ~0.1% in bwad units.
    uint256 internal constant PER_MILLE_SHIFT = BWAD_SHIFT - 10;
    uint256 internal constant PER_MILLE = 1 << PER_MILLE_SHIFT;

    /// @notice Compresses uint256 number into 16 bits.
    function compress(uint256 value) internal pure returns (Number) {
        // Find `msb` such as `2**msb <= value < 2**(msb + 1)`
        uint256 msb = mostSignificantBit(value);
        // We want to preserve 9 bits of precision.
        // The highest bit is always 1, so we can skip it.
        // The remaining 8 highest bits are stored as mantissa.
        if (msb < 8) {
            // Value is less than 2**8, so we can use value as mantissa with "-1" exponent.
            return _encode(uint8(value), 0xFF);
        } else {
            // We use `msb - 8` as exponent otherwise. Note that `exponent >= 0`.
            unchecked {
                uint256 exponent = msb - 8;
                // Shifting right by `msb-8` bits will shift the "remaining 8 highest bits" into the 8 lowest bits.
                // uint8() will truncate the highest bit.
                return _encode(uint8(value >> exponent), uint8(exponent));
            }
        }
    }

    /// @notice Decompresses 16 bits number into uint256.
    /// @dev The outcome is an approximation of the original number: `(value - value / 256) < number <= value`.
    function decompress(Number number) internal pure returns (uint256 value) {
        // Isolate 8 highest bits as the mantissa.
        uint256 mantissa = Number.unwrap(number) >> SHIFT_MANTISSA;
        // This will truncate the highest bits, leaving only the exponent.
        uint256 exponent = uint8(Number.unwrap(number));
        if (exponent == 0xFF) {
            return mantissa;
        } else {
            unchecked {
                return (256 + mantissa) << (exponent);
            }
        }
    }

    /// @dev Returns the most significant bit of `x`
    /// https://solidity-by-example.org/bitwise/
    function mostSignificantBit(uint256 x) internal pure returns (uint256 msb) {
        // To find `msb` we determine it bit by bit, starting from the highest one.
        // `0 <= msb <= 255`, so we start from the highest bit, 1<<7 == 128.
        // If `x` is at least 2**128, then the highest bit of `x` is at least 128.
        // solhint-disable no-inline-assembly
        assembly {
            // `f` is set to 1<<7 if `x >= 2**128` and to 0 otherwise.
            let f := shl(7, gt(x, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF))
            // If `x >= 2**128` then set `msb` highest bit to 1 and shift `x` right by 128.
            // Otherwise, `msb` remains 0 and `x` remains unchanged.
            x := shr(f, x)
            msb := or(msb, f)
        }
        // `x` is now at most 2**128 - 1. Continue the same way, the next highest bit is 1<<6 == 64.
        assembly {
            // `f` is set to 1<<6 if `x >= 2**64` and to 0 otherwise.
            let f := shl(6, gt(x, 0xFFFFFFFFFFFFFFFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<5 if `x >= 2**32` and to 0 otherwise.
            let f := shl(5, gt(x, 0xFFFFFFFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<4 if `x >= 2**16` and to 0 otherwise.
            let f := shl(4, gt(x, 0xFFFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<3 if `x >= 2**8` and to 0 otherwise.
            let f := shl(3, gt(x, 0xFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<2 if `x >= 2**4` and to 0 otherwise.
            let f := shl(2, gt(x, 0xF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<1 if `x >= 2**2` and to 0 otherwise.
            let f := shl(1, gt(x, 0x3))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1 if `x >= 2**1` and to 0 otherwise.
            let f := gt(x, 0x1)
            msb := or(msb, f)
        }
    }

    /// @dev Wraps (mantissa, exponent) pair into Number.
    function _encode(uint8 mantissa, uint8 exponent) private pure returns (Number) {
        // Casts below are upcasts, so they are safe.
        return Number.wrap(uint16(mantissa) << SHIFT_MANTISSA | uint16(exponent));
    }
}
