// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Mock interface with all available "default" actions as per libs/Structs.sol.
/// Note: this interface is only included to generate the common ABi to enable the action calldata
/// decoding using the generated binding. There is in fact not a single implementation of this interface,
/// as it includes woth WETH and Pool functions.
interface IDefaultActions {
    // ═══════════════════════════════════════════════════ WETH ════════════════════════════════════════════════════════
    function deposit(uint256) external payable;

    function withdraw(uint256 amount) external;

    // ═══════════════════════════════════════════════════ POOL ════════════════════════════════════════════════════════

    function addLiquidity(uint256[] calldata amounts, uint256 minToMint, uint256 deadline) external returns (uint256);

    function removeLiquidityOneToken(
        uint256 tokenAmount,
        uint8 tokenIndex,
        uint256 minAmount,
        uint256 deadline
    )
        external
        returns (uint256);

    function swap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx,
        uint256 minDy,
        uint256 deadline
    )
        external
        returns (uint256 amountOut);
}
