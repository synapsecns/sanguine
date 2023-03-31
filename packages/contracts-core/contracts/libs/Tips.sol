// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ByteString } from "./ByteString.sol";
import { TypedMemView } from "./TypedMemView.sol";

/// @dev Tips is a memory over over a formatted message tips payload.
type Tips is bytes29;
/// @dev Attach library functions to Tips
using {
    TipsLib.unwrap,
    TipsLib.version,
    TipsLib.notaryTip,
    TipsLib.broadcasterTip,
    TipsLib.proverTip,
    TipsLib.executorTip,
    TipsLib.totalTips
} for Tips global;

/**
 * @notice Library for versioned formatting [the tips part]
 * of [the messages used by Origin and Destination].
 */
library TipsLib {
    using ByteString for bytes;
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
    ▏*║                                 TIPS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted Tips payload with provided fields
     * @param notaryTip_        Tip for the Notary
     * @param broadcasterTip_   Tip for the Broadcaster
     * @param proverTip_        Tip for the Prover
     * @param executorTip_      Tip for the Executor
     * @return Formatted tips
     **/
    function formatTips(
        uint96 notaryTip_,
        uint96 broadcasterTip_,
        uint96 proverTip_,
        uint96 executorTip_
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked(TIPS_VERSION, notaryTip_, broadcasterTip_, proverTip_, executorTip_);
    }

    /**
     * @notice Returns a formatted Tips payload specifying empty tips.
     * @return Formatted tips
     **/
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
        uint256 length = view_.len();
        // Check if version exists in the payload
        if (length < OFFSET_NOTARY) return false;
        // Check that tips version and its length matches
        return _getVersion(view_) == TIPS_VERSION && length == TIPS_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Tips tips) internal pure returns (bytes29) {
        return Tips.unwrap(tips);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TIPS SLICING                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns version of formatted tips
    function version(Tips tips) internal pure returns (uint16) {
        // Get the underlying memory view
        bytes29 view_ = tips.unwrap();
        return _getVersion(view_);
    }

    /// @notice Returns notaryTip field
    function notaryTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_NOTARY, 12));
    }

    /// @notice Returns broadcasterTip field
    function broadcasterTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_BROADCASTER, 12));
    }

    /// @notice Returns proverTip field
    function proverTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_PROVER, 12));
    }

    /// @notice Returns executorTip field
    function executorTip(Tips tips) internal pure returns (uint96) {
        bytes29 view_ = tips.unwrap();
        return uint96(view_.indexUint(OFFSET_EXECUTOR, 12));
    }

    /// @notice Returns total tip amount.
    function totalTips(Tips tips) internal pure returns (uint96) {
        // In practice there's no chance that the total tips value would not fit into uint96.
        // TODO: determine if we want to use uint256 here instead anyway.
        return notaryTip(tips) + broadcasterTip(tips) + proverTip(tips) + executorTip(tips);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PRIVATE HELPERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a version field without checking if payload is properly formatted.
    function _getVersion(bytes29 view_) private pure returns (uint16) {
        return uint16(view_.indexUint(OFFSET_VERSION, 2));
    }
}
