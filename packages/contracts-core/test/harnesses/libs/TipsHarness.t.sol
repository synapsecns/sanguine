// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Tips } from "../../../contracts/libs/Tips.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes Tips methods for testing against golang.
 */
contract TipsHarness {
    using Tips for bytes;
    using Tips for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToTips(uint40, bytes memory _payload) public view returns (uint40, bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = Tips.castToTips(_payload);
        return (_view.typeOf(), _view.clone());
    }

    /// @notice Returns version of formatted tips
    function tipsVersion(uint40 _type, bytes memory _payload) public pure returns (uint16) {
        return _payload.ref(_type).tipsVersion();
    }

    /// @notice Returns notaryTip field
    function notaryTip(uint40 _type, bytes memory _payload) public pure returns (uint96) {
        return _payload.ref(_type).notaryTip();
    }

    /// @notice Returns broadcasterTip field
    function broadcasterTip(uint40 _type, bytes memory _payload) public pure returns (uint96) {
        return _payload.ref(_type).broadcasterTip();
    }

    /// @notice Returns proverTip field
    function proverTip(uint40 _type, bytes memory _payload) public pure returns (uint96) {
        return _payload.ref(_type).proverTip();
    }

    /// @notice Returns executorTip field
    function executorTip(uint40 _type, bytes memory _payload) public pure returns (uint96) {
        return _payload.ref(_type).executorTip();
    }

    /// @notice Returns total tip amount.
    function totalTips(uint40 _type, bytes memory _payload) public pure returns (uint96) {
        return _payload.ref(_type).totalTips();
    }

    function isTips(bytes memory _payload) public pure returns (bool) {
        return _payload.castToTips().isTips();
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
