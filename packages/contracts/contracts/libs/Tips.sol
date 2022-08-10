// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";
import { TypeCasts } from "./TypeCasts.sol";
import { Message } from "./Message.sol";

/**
 * @notice Library for versioned formatting [the tips part] of [the messages used by Home and Replicas].
 */
library Tips {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint16 internal constant TIPS_VERSION = 1;

    /**
     * @dev Tips memory layout
     * [000 .. 002): version            uint16	 2 bytes
     * [002 .. 014): notaryTip          uint96	12 bytes
     * [014 .. 026): broadcasterTip     uint96	12 bytes
     * [026 .. 038): proverTip          uint96	12 bytes
     * [038 .. 050): executorTip        uint96	12 bytes
     */

    uint256 internal constant OFFSET_NOTARY = 2;
    uint256 internal constant OFFSET_BROADCASTER = 14;
    uint256 internal constant OFFSET_PROVER = 26;
    uint256 internal constant OFFSET_EXECUTOR = 38;

    modifier onlyTips(bytes29 _view) {
        _view.assertType(Message.TIPS_TYPE);
        _;
    }

    /**
     * @notice Returns formatted (packed) tips with provided fields
     * @param _notaryTip Tip for the Notary
     * @param _broadcasterTip Tip for the Broadcaster
     * @param _proverTip Tip for the Prover
     * @param _executorTip Tip for the Executor
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
     * @notice Returns formatted empty tips
     * @return Formatted tips
     **/
    function emptyTips() internal pure returns (bytes memory) {
        return formatTips(0, 0, 0, 0);
    }

    /// @notice Returns view for the formatted tips
    /// @dev Providing anything other than formatted tips will lead to unexpected behavior
    function tipsView(bytes memory _tips) internal pure returns (bytes29) {
        return _tips.ref(Message.TIPS_TYPE);
    }

    /// @notice Returns version of formatted tips
    function tipsVersion(bytes29 _tips) internal pure onlyTips(_tips) returns (uint16) {
        return uint16(_tips.indexUint(0, 2));
    }

    /// @notice Returns notaryTip field
    function notaryTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint32(_tips.indexUint(OFFSET_NOTARY, 12));
    }

    /// @notice Returns broadcasterTip field
    function broadcasterTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint32(_tips.indexUint(OFFSET_BROADCASTER, 12));
    }

    /// @notice Returns proverTip field
    function proverTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint32(_tips.indexUint(OFFSET_PROVER, 12));
    }

    /// @notice Returns executorTip field
    function executorTip(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return uint32(_tips.indexUint(OFFSET_EXECUTOR, 12));
    }

    function totalTips(bytes29 _tips) internal pure onlyTips(_tips) returns (uint96) {
        return notaryTip(_tips) + broadcasterTip(_tips) + proverTip(_tips) + executorTip(_tips);
    }
}
