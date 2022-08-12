// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { Tips } from "../../contracts/libs/Tips.sol";

/**
 * @notice exposes tips methods for testing against golang
 */
contract TipsHarness {
    using Tips for bytes;
    using Tips for bytes29;

    function tipsVersion() public pure returns (uint16) {
        return Tips.TIPS_VERSION;
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

    function formatTips(
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip
    ) public pure returns (bytes memory) {
        return Tips.formatTips(_notaryTip, _broadcasterTip, _proverTip, _executorTip);
    }

    function emptyTips() external pure returns (bytes memory) {
        return Tips.emptyTips();
    }

    function castToTips(bytes memory _tips) external pure returns (bytes29) {
        return Tips.castToTips(_tips);
    }

    function tipsVersion(bytes29 _tips) external pure returns (uint16) {
        return Tips.tipsVersion(_tips);
    }

    /// @notice Returns notaryTip field
    function notaryTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.notaryTip(_tips);
    }

    /// @notice Returns broadcasterTip field
    function broadcasterTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.broadcasterTip(_tips);
    }

    /// @notice Returns proverTip field
    function proverTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.proverTip(_tips);
    }

    /// @notice Returns executorTip field
    function executorTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.executorTip(_tips);
    }

    function totalTips(bytes29 _tips) external pure returns (uint96) {
        return Tips.totalTips(_tips);
    }
}
