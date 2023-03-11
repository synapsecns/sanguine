// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Tips, TipsLib, TypedMemView } from "../../../contracts/libs/Tips.sol";

/**
 * @notice Exposes TipsLib methods for testing against golang.
 */
contract TipsHarness {
    using TipsLib for bytes;
    using TipsLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToTips(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Tips tips = TipsLib.castToTips(_payload);
        return tips.unwrap().clone();
    }

    /// @notice Returns version of formatted tips
    function version(bytes memory _payload) public pure returns (uint16) {
        return _payload.castToTips().version();
    }

    /// @notice Returns notaryTip field
    function notaryTip(bytes memory _payload) public pure returns (uint96) {
        return _payload.castToTips().notaryTip();
    }

    /// @notice Returns broadcasterTip field
    function broadcasterTip(bytes memory _payload) public pure returns (uint96) {
        return _payload.castToTips().broadcasterTip();
    }

    /// @notice Returns proverTip field
    function proverTip(bytes memory _payload) public pure returns (uint96) {
        return _payload.castToTips().proverTip();
    }

    /// @notice Returns executorTip field
    function executorTip(bytes memory _payload) public pure returns (uint96) {
        return _payload.castToTips().executorTip();
    }

    /// @notice Returns total tip amount.
    function totalTips(bytes memory _payload) public pure returns (uint96) {
        return _payload.castToTips().totalTips();
    }

    function isTips(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isTips();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatTips(
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip
    ) public pure returns (bytes memory) {
        return TipsLib.formatTips(_notaryTip, _broadcasterTip, _proverTip, _executorTip);
    }

    function emptyTips() public pure returns (bytes memory) {
        return TipsLib.emptyTips();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function tipsLength() public pure returns (uint256) {
        return TipsLib.TIPS_LENGTH;
    }

    function tipsVersion() public pure returns (uint16) {
        return TipsLib.TIPS_VERSION;
    }

    function offsetVersion() public pure returns (uint256) {
        return TipsLib.OFFSET_VERSION;
    }

    function offsetNotary() public pure returns (uint256) {
        return TipsLib.OFFSET_NOTARY;
    }

    function offsetBroadcaster() public pure returns (uint256) {
        return TipsLib.OFFSET_BROADCASTER;
    }

    function offsetProver() public pure returns (uint256) {
        return TipsLib.OFFSET_PROVER;
    }

    function offsetExecutor() public pure returns (uint256) {
        return TipsLib.OFFSET_EXECUTOR;
    }
}
