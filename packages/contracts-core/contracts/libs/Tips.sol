// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TypedMemView } from "./TypedMemView.sol";
import { TypeCasts } from "./TypeCasts.sol";
import { SynapseTypes } from "./SynapseTypes.sol";

/**
 * @notice Library for versioned formatting [the tips part]
 * of [the messages used by Origin and Destination].
 */
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
