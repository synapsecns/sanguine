// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Tips, TipsLib} from "../../../../contracts/libs/stack/Tips.sol";

/**
 * @notice Exposes TipsLib methods for testing against golang.
 */
contract TipsHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    /// @notice Returns summitTip field
    function summitTip(uint256 paddedTips) public pure returns (uint64) {
        return TipsLib.wrapPadded(paddedTips).summitTip();
    }

    /// @notice Returns attestationTip field
    function attestationTip(uint256 paddedTips) public pure returns (uint64) {
        return TipsLib.wrapPadded(paddedTips).attestationTip();
    }

    /// @notice Returns executionTip field
    function executionTip(uint256 paddedTips) public pure returns (uint64) {
        return TipsLib.wrapPadded(paddedTips).executionTip();
    }

    /// @notice Returns deliveryTip field
    function deliveryTip(uint256 paddedTips) public pure returns (uint64) {
        return TipsLib.wrapPadded(paddedTips).deliveryTip();
    }

    /// @notice Returns total tip amount.
    function value(uint256 paddedTips) public pure returns (uint256) {
        return TipsLib.wrapPadded(paddedTips).value();
    }

    /// @notice Increases the delivery tip to match the new value.
    function matchValue(Tips tips, uint256 newValue) public pure returns (Tips newTips) {
        return tips.matchValue(newValue);
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function encodeTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_)
        public
        pure
        returns (uint256)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Tips tips = TipsLib.encodeTips(summitTip_, attestationTip_, executionTip_, deliveryTip_);
        return Tips.unwrap(tips);
    }

    function encodeTips256(uint256 summitTip_, uint256 attestationTip_, uint256 executionTip_, uint256 deliveryTip_)
        public
        pure
        returns (uint256)
    {
        Tips tips = TipsLib.encodeTips256(summitTip_, attestationTip_, executionTip_, deliveryTip_);
        return Tips.unwrap(tips);
    }

    function wrapPadded(uint256 paddedTips) public pure returns (uint256) {
        return Tips.unwrap(TipsLib.wrapPadded(paddedTips));
    }

    function leaf(uint256 paddedTips) public pure returns (bytes32) {
        return TipsLib.wrapPadded(paddedTips).leaf();
    }

    function emptyTips() public pure returns (uint256) {
        // TODO: figure out why this leaves `TipsLib.emptyTips()` uncovered
        Tips tips = TipsLib.emptyTips();
        return Tips.unwrap(tips);
    }
}
