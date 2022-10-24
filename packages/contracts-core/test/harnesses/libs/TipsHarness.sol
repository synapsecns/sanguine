// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Tips } from "../../../contracts/libs/Tips.sol";

/**
 * @notice Exposes Tips methods for testing against golang.
 */
contract TipsHarness {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatTips(
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip
    ) public pure returns (bytes memory) {
        return Tips.formatTips(_notaryTip, _broadcasterTip, _proverTip, _executorTip);
    }

    function emptyTips() public pure returns (bytes memory) {
        return Tips.emptyTips();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function tipsLength() public pure returns (uint256) {
        return Tips.TIPS_LENGTH;
    }

    function tipsVersion() public pure returns (uint16) {
        return Tips.TIPS_VERSION;
    }

    function offsetVersion() public pure returns (uint256) {
        return Tips.OFFSET_VERSION;
    }

    function offsetNotary() public pure returns (uint256) {
        return Tips.OFFSET_NOTARY;
    }

    function offsetBroadcaster() public pure returns (uint256) {
        return Tips.OFFSET_BROADCASTER;
    }

    function offsetProver() public pure returns (uint256) {
        return Tips.OFFSET_PROVER;
    }

    function offsetExecutor() public pure returns (uint256) {
        return Tips.OFFSET_EXECUTOR;
    }
}
