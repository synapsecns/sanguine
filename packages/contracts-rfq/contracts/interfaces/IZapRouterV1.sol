// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IZapRouterV1 {
    struct ZapParams {
        address token;
        uint256 amount;
        uint256 msgValue;
        bytes zapData;
    }

    /// @notice Perform a series of Zap actions in a single transaction. Verifies that the ZapRecipient balance
    /// for every token in `zapParams` has not increased after the last Zap.
    /// - Each step is verified to be a correct Zap as per `IZapRecipient` specification.
    /// - The amounts used for each Zap can be predetermined or based on the proceeds from the previous Zaps.
    /// - ZapRouter does not perform any checks on the Zap Data, nor the ZapRecipient balance after the Zaps are
    ///   performed.
    /// - The user is responsible for selecting a correct ZapRecipient and for the correct encoding of the Zap Data.
    /// - ZapRecipient must be able to modify the Zap Data to adjust to possible changes in the passed amount value.
    /// @dev Typical workflow involves a series of "prep" Zaps followed by a final Zap, such as
    /// bridging, depositing, or a transfer to the final recipient. ZapRecipient must be set as the funds recipient
    /// for the prep Zaps, while a different recipient must be set for the last Zap.
    /// @dev This function will revert in any of the following cases:
    /// - The deadline has passed.
    /// - The array of ZapParams is empty.
    /// - The amount of tokens to use for the last Zap is below the specified minimum.
    /// - Any Zap fails.
    /// @param zapRecipient         Address of the IZapRecipient contract to use for the series of Zaps
    /// @param amountIn             Initial amount of tokens (zapParams[0].token) to transfer into ZapRecipient
    /// @param minLastZapAmountIn   Minimum amount of tokens (zapParams[N-1].token) to use for the last Zap
    /// @param deadline             Deadline for the series of Zaps
    /// @param zapParams            Parameters for each Zap. Use amount = type(uint256).max for Zaps that
    ///                             should use the full ZapRecipient balance.
    function performZapsWithBalanceChecks(
        address zapRecipient,
        uint256 amountIn,
        uint256 minLastZapAmountIn,
        uint256 deadline,
        ZapParams[] memory zapParams
    )
        external
        payable;

    /// @notice Perform a series of Zap actions in a single transaction.
    /// @dev This function is identical to `performZapsWithBalanceChecks` except that it does not verify that
    /// the ZapRecipient balance for every token in `zapParams` has not increased after the last Zap.
    /// Anyone using this function must validate that the funds are fully spent by ZapRecipient
    /// using other means like separate on-chain checks or off-chain simulation.
    function performZaps(
        address zapRecipient,
        uint256 amountIn,
        uint256 minLastZapAmountIn,
        uint256 deadline,
        ZapParams[] memory zapParams
    )
        external
        payable;
}
