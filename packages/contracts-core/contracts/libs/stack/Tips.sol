// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {TIPS_GRANULARITY} from "../Constants.sol";
import {TipsOverflow, TipsValueTooLow} from "../Errors.sol";
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

/// Tips is encoded data with "tips paid for sending a base message".
/// Note: even though uint256 is also an underlying type for MemView, Tips is stored ON STACK.
type Tips is uint256;

using TipsLib for Tips global;

/// # Tips
/// Library for formatting _the tips part_ of _the base messages_.
///
/// ## How the tips are awarded
/// Tips are paid for sending a base message, and are split across all the agents that
/// made the message execution on destination chain possible.
/// ### Summit tips
/// Split between:
///     - Guard posting a snapshot with state ST_G for the origin chain.
///     - Notary posting a snapshot SN_N using ST_G. This creates attestation A.
///     - Notary posting a message receipt after it is executed on destination chain.
/// ### Attestation tips
/// Paid to:
///     - Notary posting attestation A to destination chain.
/// ### Execution tips
/// Paid to:
///     - First executor performing a valid execution attempt (correct proofs, optimistic period over),
///      using attestation A to prove message inclusion on origin chain, whether the recipient reverted or not.
/// ### Delivery tips.
/// Paid to:
///     - Executor who successfully executed the message on destination chain.
///
/// ## Tips encoding
/// - Tips occupy a single storage word, and thus are stored on stack instead of being stored in memory.
/// - The actual tip values should be determined by multiplying stored values by divided by TIPS_MULTIPLIER=2**32.
/// - Tips are packed into a single word of storage, while allowing real values up to ~8*10**28 for every tip category.
/// > The only downside is that the "real tip values" are now multiplies of ~4*10**9, which should be fine even for
/// the chains with the most expensive gas currency.
/// # Tips stack layout (from highest bits to lowest)
///
/// | Position   | Field          | Type   | Bytes | Description                                                |
/// | ---------- | -------------- | ------ | ----- | ---------------------------------------------------------- |
/// | (032..024] | summitTip      | uint64 | 8     | Tip for agents interacting with Summit contract            |
/// | (024..016] | attestationTip | uint64 | 8     | Tip for Notary posting attestation to Destination contract |
/// | (016..008] | executionTip   | uint64 | 8     | Tip for valid execution attempt on destination chain       |
/// | (008..000] | deliveryTip    | uint64 | 8     | Tip for successful message delivery on destination chain   |

library TipsLib {
    using SafeCast for uint256;

    /// @dev Amount of bits to shift to summitTip field
    uint256 private constant SHIFT_SUMMIT_TIP = 24 * 8;
    /// @dev Amount of bits to shift to attestationTip field
    uint256 private constant SHIFT_ATTESTATION_TIP = 16 * 8;
    /// @dev Amount of bits to shift to executionTip field
    uint256 private constant SHIFT_EXECUTION_TIP = 8 * 8;

    // ═══════════════════════════════════════════════════ TIPS ════════════════════════════════════════════════════════

    /// @notice Returns encoded tips with the given fields
    /// @param summitTip_        Tip for agents interacting with Summit contract, divided by TIPS_MULTIPLIER
    /// @param attestationTip_   Tip for Notary posting attestation to Destination contract, divided by TIPS_MULTIPLIER
    /// @param executionTip_     Tip for valid execution attempt on destination chain, divided by TIPS_MULTIPLIER
    /// @param deliveryTip_      Tip for successful message delivery on destination chain, divided by TIPS_MULTIPLIER
    function encodeTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_)
        internal
        pure
        returns (Tips)
    {
        // forgefmt: disable-next-item
        return Tips.wrap(
            uint256(summitTip_) << SHIFT_SUMMIT_TIP |
            uint256(attestationTip_) << SHIFT_ATTESTATION_TIP |
            uint256(executionTip_) << SHIFT_EXECUTION_TIP |
            uint256(deliveryTip_)
        );
    }

    /// @notice Convenience function to encode tips with uint256 values.
    function encodeTips256(uint256 summitTip_, uint256 attestationTip_, uint256 executionTip_, uint256 deliveryTip_)
        internal
        pure
        returns (Tips)
    {
        // In practice, the tips amounts are not supposed to be higher than 2**96, and with 32 bits of granularity
        // using uint64 is enough to store the values. However, we still check for overflow just in case.
        // TODO: consider using Number type to store the tips values.
        return encodeTips({
            summitTip_: (summitTip_ >> TIPS_GRANULARITY).toUint64(),
            attestationTip_: (attestationTip_ >> TIPS_GRANULARITY).toUint64(),
            executionTip_: (executionTip_ >> TIPS_GRANULARITY).toUint64(),
            deliveryTip_: (deliveryTip_ >> TIPS_GRANULARITY).toUint64()
        });
    }

    /// @notice Wraps the padded encoded tips into a Tips-typed value.
    /// @dev There is no actual padding here, as the underlying type is already uint256,
    /// but we include this function for consistency and to be future-proof, if tips will eventually use anything
    /// smaller than uint256.
    function wrapPadded(uint256 paddedTips) internal pure returns (Tips) {
        return Tips.wrap(paddedTips);
    }

    /**
     * @notice Returns a formatted Tips payload specifying empty tips.
     * @return Formatted tips
     */
    function emptyTips() internal pure returns (Tips) {
        return Tips.wrap(0);
    }

    /// @notice Returns tips's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(Tips tips) internal pure returns (bytes32 hashedTips) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Store tips in scratch space
            mstore(0, tips)
            // Compute hash of tips padded to 32 bytes
            hashedTips := keccak256(0, 32)
        }
    }

    // ═══════════════════════════════════════════════ TIPS SLICING ════════════════════════════════════════════════════

    /// @notice Returns summitTip field
    function summitTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips) >> SHIFT_SUMMIT_TIP);
    }

    /// @notice Returns attestationTip field
    function attestationTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips) >> SHIFT_ATTESTATION_TIP);
    }

    /// @notice Returns executionTip field
    function executionTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips) >> SHIFT_EXECUTION_TIP);
    }

    /// @notice Returns deliveryTip field
    function deliveryTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips));
    }

    // ════════════════════════════════════════════════ TIPS VALUE ═════════════════════════════════════════════════════

    /// @notice Returns total value of the tips payload.
    /// This is the sum of the encoded values, scaled up by TIPS_MULTIPLIER
    function value(Tips tips) internal pure returns (uint256 value_) {
        value_ = uint256(tips.summitTip()) + tips.attestationTip() + tips.executionTip() + tips.deliveryTip();
        value_ <<= TIPS_GRANULARITY;
    }

    /// @notice Increases the delivery tip to match the new value.
    function matchValue(Tips tips, uint256 newValue) internal pure returns (Tips newTips) {
        uint256 oldValue = tips.value();
        if (newValue < oldValue) revert TipsValueTooLow();
        // We want to increase the delivery tip, while keeping the other tips the same
        unchecked {
            uint256 delta = (newValue - oldValue) >> TIPS_GRANULARITY;
            // `delta` fits into uint224, as TIPS_GRANULARITY is 32, so this never overflows uint256.
            // In practice, this will never overflow uint64 as well, but we still check it just in case.
            if (delta + tips.deliveryTip() > type(uint64).max) revert TipsOverflow();
            // Delivery tips occupy lowest 8 bytes, so we can just add delta to the tips value
            // to effectively increase the delivery tip (knowing that delta fits into uint64).
            newTips = Tips.wrap(Tips.unwrap(tips) + delta);
        }
    }
}
