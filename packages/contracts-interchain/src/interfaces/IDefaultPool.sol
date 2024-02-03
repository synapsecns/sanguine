// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Minimal interface for the DefaultPool, supported by the SynapseBridge contract.
interface IDefaultPool {
    /**
     * @notice Swap two tokens using this pool
     * @param tokenIndexFrom    The index of the token to swap from
     * @param tokenIndexTo      The index of the token to swap to
     * @param dx                The amount of tokens to swap from
     * @param minDy             The minimum amount of tokens to receive
     * @param deadline          The deadline for the swap
     */
    function swap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx,
        uint256 minDy,
        uint256 deadline
    )
        external
        returns (uint256 amountOut);

    /**
     * @notice Calculate the amount of tokens to receive for a given amount of tokens to swap
     * @param tokenIndexFrom    The index of the token to swap from
     * @param tokenIndexTo      The index of the token to swap to
     * @param dx                The amount of tokens to swap from
     * @return amountOut        The amount of tokens to receive
     */
    function calculateSwap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx
    )
        external
        view
        returns (uint256 amountOut);

    /**
     * @notice Get the address of the token at the given index
     * @param index             The index of the token
     * @return token            The address of the token
     */
    function getToken(uint8 index) external view returns (address token);
}
