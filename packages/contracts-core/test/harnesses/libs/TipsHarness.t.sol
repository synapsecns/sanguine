// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Tips, TipsLib, TIPS_LENGTH, MemView, MemViewLib} from "../../../contracts/libs/Tips.sol";

/**
 * @notice Exposes TipsLib methods for testing against golang.
 */
contract TipsHarness {
    using TipsLib for bytes;
    using TipsLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToTips(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Tips tips = TipsLib.castToTips(payload);
        return tips.unwrap().clone();
    }

    /// @notice Returns summitTip field
    function summitTip(bytes memory payload) public pure returns (uint64) {
        return payload.castToTips().summitTip();
    }

    /// @notice Returns attestationTip field
    function attestationTip(bytes memory payload) public pure returns (uint64) {
        return payload.castToTips().attestationTip();
    }

    /// @notice Returns executionTip field
    function executionTip(bytes memory payload) public pure returns (uint64) {
        return payload.castToTips().executionTip();
    }

    /// @notice Returns deliveryTip field
    function deliveryTip(bytes memory payload) public pure returns (uint64) {
        return payload.castToTips().deliveryTip();
    }

    /// @notice Returns total tip amount.
    function value(bytes memory payload) public pure returns (uint256) {
        return payload.castToTips().value();
    }

    function isTips(bytes memory payload) public pure returns (bool) {
        return payload.ref().isTips();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_)
        public
        pure
        returns (bytes memory)
    {
        return TipsLib.formatTips(summitTip_, attestationTip_, executionTip_, deliveryTip_);
    }

    function emptyTips() public pure returns (bytes memory) {
        return TipsLib.emptyTips();
    }
}
