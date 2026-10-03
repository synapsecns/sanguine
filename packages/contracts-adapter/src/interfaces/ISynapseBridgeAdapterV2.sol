// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {ISynapseBridgeAdapter} from "./ISynapseBridgeAdapter.sol";

interface ISynapseBridgeAdapterV2 is ISynapseBridgeAdapter {
    /// @notice Configures the 18-decimal SYN to 8-decimal HyperCore route.
    /// @dev Reads the local ERC20 decimals and rejects any value other than the fixed SYN decimal pair.
    function setHyperCoreDecimals(uint32 dstEid, address token, uint8 coreDecimals) external;

    /// @notice Configures the destination composer for a local token. The composer binding is immutable.
    function setHyperCoreComposer(address token, address composer) external;

    /// @notice Bridges an ERC20 token to a recipient on HyperCore through the destination composer.
    function bridgeERC20ToHyperCore(
        uint32 dstEid,
        address to,
        address token,
        uint256 amount,
        uint64 gasLimit
    )
        external
        payable;

    /// @notice Returns the native fee for a HyperCore delivery message.
    function getNativeFeeToHyperCore(
        uint32 dstEid,
        address token,
        uint64 gasLimit
    )
        external
        view
        returns (uint256 nativeFee);

    /// @notice Returns the amount scale for a configured HyperCore token route.
    function getHyperCoreAmountScale(uint32 dstEid, address token) external view returns (uint256 amountScale);

    /// @notice Returns the destination HyperCore composer for a local token.
    function getHyperCoreComposer(address token) external view returns (address composer);
}
