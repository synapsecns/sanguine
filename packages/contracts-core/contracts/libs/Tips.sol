// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {TIPS_GRANULARITY, TIPS_LENGTH} from "./Constants.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev Tips is a memory over over a formatted message tips payload.
type Tips is bytes29;

/// @dev Attach library functions to Tips
using TipsLib for Tips global;

/**
 * @notice Library for versioned formatting [the tips part]
 * of [the messages used by Origin and Destination].
 */
library TipsLib {
    using ByteString for bytes;
    using TypedMemView for bytes29;

    // TODO: determine if we need to pack the tips values,
    // or if using uint256 instead will suffice.

    /**
     * @dev Tips are paid for sending a base message, and are split across all the agents that
     * made the message execution on destination chain possible.
     *  1. Summit tips. Split between:
     *      a. Guard posting a snapshot with state ST_G for the origin chain.
     *      b. Notary posting a snapshot SN_N using ST_G. This creates attestation A.
     *      c. Notary posting a message receipt after it is executed on destination chain.
     *  2. Attestation tips. Paid to:
     *      a. Notary posting attestation A to destination chain.
     *  3. Execution tips. Paid to:
     *      a. First executor performing a valid execution attempt (correct proofs, optimistic period over),
     *      using attestation A to prove message inclusion on origin chain, whether the recipient reverted or not.
     *  4. Delivery tips. Paid to:
     *      a. Executor who successfully executed the message on destination chain.
     * @dev Tips memory layout
     * [000 .. 012): summitTip          uint96	12 bytes    Tip for agents interacting with Summit contract
     * [012 .. 024): attestationTip     uint96	12 bytes    Tip for Notary posting attestation to Destination contract
     * [024 .. 036): executionTip       uint96	12 bytes    Tip for valid execution attempt on destination chain
     * [036 .. 048): deliveryTip        uint96	12 bytes    Tip for successful message delivery on destination chain
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_SUMMIT_TIP = 0;
    uint256 private constant OFFSET_ATTESTATION_TIP = 12;
    uint256 private constant OFFSET_EXECUTION_TIP = 24;
    uint256 private constant OFFSET_DELIVERY_TIP = 36;

    // ═══════════════════════════════════════════════════ TIPS ════════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Tips payload with provided fields
     * @param summitTip_        Tip for agents interacting with Summit contract
     * @param attestationTip_   Tip for Notary posting attestation to Destination contract
     * @param executionTip_     Tip for valid execution attempt on destination chain
     * @param deliveryTip_      Tip for successful message delivery on destination chain
     * @return Formatted tips
     */
    function formatTips(uint96 summitTip_, uint96 attestationTip_, uint96 executionTip_, uint96 deliveryTip_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(summitTip_, attestationTip_, executionTip_, deliveryTip_);
    }

    /**
     * @notice Returns a formatted Tips payload specifying empty tips.
     * @return Formatted tips
     */
    function emptyTips() internal pure returns (bytes memory) {
        return formatTips(0, 0, 0, 0);
    }

    /**
     * @notice Returns a Tips view over for the given payload.
     * @dev Will revert if the payload is not a tips payload.
     */
    function castToTips(bytes memory payload) internal pure returns (Tips) {
        return castToTips(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Tips view.
     * @dev Will revert if the memory view is not over a tips payload.
     */
    function castToTips(bytes29 view_) internal pure returns (Tips) {
        require(isTips(view_), "Not a tips payload");
        return Tips.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted Tips payload.
    function isTips(bytes29 view_) internal pure returns (bool) {
        return view_.len() == TIPS_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Tips tips) internal pure returns (bytes29) {
        return Tips.unwrap(tips);
    }

    // ═══════════════════════════════════════════════ TIPS SLICING ════════════════════════════════════════════════════

    /// @notice Returns summitTip field
    function summitTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_SUMMIT_TIP, 12));
    }

    /// @notice Returns attestationTip field
    function attestationTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_ATTESTATION_TIP, 12));
    }

    /// @notice Returns executionTip field
    function executionTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_EXECUTION_TIP, 12));
    }

    /// @notice Returns deliveryTip field
    function deliveryTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_DELIVERY_TIP, 12));
    }

    /// @notice Returns total value of the tips payload.
    /// This is the sum of the encoded values, scaled up by TIPS_MULTIPLIER
    function value(Tips tips) internal pure returns (uint256 value_) {
        value_ = summitTip(tips) + attestationTip(tips) + executionTip(tips) + deliveryTip(tips);
        value_ <<= TIPS_GRANULARITY;
    }
}
