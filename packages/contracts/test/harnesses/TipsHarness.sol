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

    function offsetUpdater() public pure returns (uint256) {
        return Tips.OFFSET_UPDATER;
    }

    function offsetRelayer() public pure returns (uint256) {
        return Tips.OFFSET_RELAYER;
    }

    function offsetProver() public pure returns (uint256) {
        return Tips.OFFSET_PROVER;
    }

    function offsetProcessor() public pure returns (uint256) {
        return Tips.OFFSET_PROCESSOR;
    }

    function formatTips(
        uint96 _updaterTip,
        uint96 _relayerTip,
        uint96 _proverTip,
        uint96 _processorTip
    ) public pure returns (bytes memory) {
        return Tips.formatTips(_updaterTip, _relayerTip, _proverTip, _processorTip);
    }

    function emptyTips() external pure returns (bytes memory) {
        return Tips.emptyTips();
    }

    function tipsView(bytes memory _tips) external pure returns (bytes29) {
        return Tips.tipsView(_tips);
    }

    function tipsVersion(bytes29 _tips) external pure returns (uint16) {
        return Tips.tipsVersion(_tips);
    }

    /// @notice Returns updaterTip field
    function updaterTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.updaterTip(_tips);
    }

    /// @notice Returns relayerTip field
    function relayerTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.relayerTip(_tips);
    }

    /// @notice Returns proverTip field
    function proverTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.proverTip(_tips);
    }

    /// @notice Returns processorTip field
    function processorTip(bytes29 _tips) external pure returns (uint96) {
        return Tips.processorTip(_tips);
    }

    function totalTips(bytes29 _tips) external pure returns (uint96) {
        return Tips.totalTips(_tips);
    }
}
