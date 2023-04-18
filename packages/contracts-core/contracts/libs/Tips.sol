// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {TIPS_GRANULARITY, TIPS_LENGTH} from "./Constants.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// @dev Tips is a memory over over a formatted message tips payload.
type Tips is uint256;

/// @dev Attach library functions to Tips
using TipsLib for Tips global;

/**
 * @notice Library for versioned formatting [the tips part]
 * of [the messages used by Origin and Destination].
 */
library TipsLib {
    using MemViewLib for bytes;

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
     * @dev The actual tip values should be determined by multiplying stored values by divided by TIPS_MULTIPLIER=2**32.
     * Tips are packed into a single word of storage, while allowing real values up to ~8*10**28 for every tip category.
     * The only downside is that the "real tip values" are now multiplies of ~4*10**9, which should be fine even for
     * the chains with the most expensive gas currency.
     * @dev Tips memory layout
     * [000 .. 008): summitTip          uint64	 8 bytes    Tip for agents interacting with Summit contract
     * [008 .. 016): attestationTip     uint64	 8 bytes    Tip for Notary posting attestation to Destination contract
     * [016 .. 024): executionTip       uint64	 8 bytes    Tip for valid execution attempt on destination chain
     * [024 .. 032): deliveryTip        uint64	 8 bytes    Tip for successful message delivery on destination chain
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_SUMMIT_TIP = 0;
    uint256 private constant OFFSET_ATTESTATION_TIP = 8;
    uint256 private constant OFFSET_EXECUTION_TIP = 16;
    uint256 private constant OFFSET_DELIVERY_TIP = 24;

    // ═══════════════════════════════════════════════════ TIPS ════════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Tips payload with provided fields
     * @param summitTip_        Tip for agents interacting with Summit contract, divided by TIPS_MULTIPLIER
     * @param attestationTip_   Tip for Notary posting attestation to Destination contract, divided by TIPS_MULTIPLIER
     * @param executionTip_     Tip for valid execution attempt on destination chain, divided by TIPS_MULTIPLIER
     * @param deliveryTip_      Tip for successful message delivery on destination chain, divided by TIPS_MULTIPLIER
     * @return Formatted tips
     */
    function formatTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_)
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
        return castToTips(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Tips view.
     * @dev Will revert if the memory view is not over a tips payload.
     */
    function castToTips(MemView memView) internal pure returns (Tips) {
        require(isTips(memView), "Not a tips payload");
        return Tips.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Tips payload.
    function isTips(MemView memView) internal pure returns (bool) {
        return memView.len() == TIPS_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Tips tips) internal pure returns (MemView) {
        return MemView.wrap(Tips.unwrap(tips));
    }

    // ═══════════════════════════════════════════════ TIPS SLICING ════════════════════════════════════════════════════

    /// @notice Returns summitTip field
    function summitTip(Tips tips) internal pure returns (uint64) {
        MemView memView = tips.unwrap();
        return uint64(_summitTip(memView));
    }

    /// @notice Returns attestationTip field
    function attestationTip(Tips tips) internal pure returns (uint64) {
        MemView memView = tips.unwrap();
        return uint64(_attestationTip(memView));
    }

    /// @notice Returns executionTip field
    function executionTip(Tips tips) internal pure returns (uint64) {
        MemView memView = tips.unwrap();
        return uint64(_executionTip(memView));
    }

    /// @notice Returns deliveryTip field
    function deliveryTip(Tips tips) internal pure returns (uint64) {
        MemView memView = tips.unwrap();
        return uint64(_deliveryTip(memView));
    }

    /// @notice Returns total value of the tips payload.
    /// This is the sum of the encoded values, scaled up by TIPS_MULTIPLIER
    function value(Tips tips) internal pure returns (uint256 value_) {
        MemView memView = tips.unwrap();
        value_ = _summitTip(memView) + _attestationTip(memView) + _executionTip(memView) + _deliveryTip(memView);
        value_ <<= TIPS_GRANULARITY;
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @notice Returns summitTip field as uint256
    function _summitTip(MemView memView) internal pure returns (uint256) {
        return memView.indexUint({index_: OFFSET_SUMMIT_TIP, bytes_: 8});
    }

    /// @notice Returns attestationTip field as uint256
    function _attestationTip(MemView memView) internal pure returns (uint256) {
        return memView.indexUint({index_: OFFSET_ATTESTATION_TIP, bytes_: 8});
    }

    /// @notice Returns executionTip field as uint256
    function _executionTip(MemView memView) internal pure returns (uint256) {
        return memView.indexUint({index_: OFFSET_EXECUTION_TIP, bytes_: 8});
    }

    /// @notice Returns deliveryTip field as uint256
    function _deliveryTip(MemView memView) internal pure returns (uint256) {
        return memView.indexUint({index_: OFFSET_DELIVERY_TIP, bytes_: 8});
    }
}
