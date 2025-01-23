// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseIntentRouter {
    /// @notice Parameters for a single Zap step.
    /// @param token    Address of the token to use for the step
    /// @param amount   Amount of tokens to use for the step (type(uint256).max to use the full ZapRecipient balance)
    /// @param msgValue Amount of native token to supply for the step, out of the total `msg.value` used for the
    ///                 `fulfillIntent` call (could differ from `amount` regardless of the token type)
    /// @param zapData  Instructions for the ZapRecipient contract on how to execute the Zap
    struct StepParams {
        address token;
        uint256 amount;
        uint256 msgValue;
        bytes zapData;
    }

    /// @notice Kindly ask SIR to complete the provided intent by completing a series of Zap steps using the
    /// provided ZapRecipient contract.
    /// - Each step is verified to be a correct Zap as per `IZapRecipient` specification.
    /// - The amounts used for each step can be predetermined or based on the proceeds from the previous steps.
    /// - SIR does not perform any checks on the Zap Data; the user is responsible for ensuring correct encoding.
    /// - The user is responsible for selecting the correct ZapRecipient for their intent: ZapRecipient must be
    ///   able to modify the Zap Data to adjust to possible changes in the passed amount value.
    /// - SIR checks that the ZapRecipient balance for every token in `steps` has not increased after the last step.
    /// - SIR does not perform any slippage checks. If required, the slippage settings must be embedded in any of
    ///   the Zap steps to be used by ZapRecipient.
    /// @dev Typical workflow involves a series of preparation steps followed by the last step representing the user
    /// intent such as bridging, depositing, or a simple transfer to the final recipient. The ZapRecipient must be
    /// the funds recipient for the preparation steps, while the final recipient must be used for the last step.
    /// @dev This function will revert in any of the following cases:
    /// - The deadline has passed.
    /// - The array of StepParams is empty.
    /// - Any step fails.
    /// - `msg.value` does not match `sum(steps[i].msgValue)`.
    /// @param zapRecipient         Address of the IZapRecipient contract to use for the Zap steps
    /// @param amountIn             Initial amount of tokens (steps[0].token) to transfer into ZapRecipient
    /// @param deadline             Deadline for the intent to be completed
    /// @param steps                Parameters for each step. Use amount = type(uint256).max for steps that
    ///                             should use the full ZapRecipient balance.
    function completeIntentWithBalanceChecks(
        address zapRecipient,
        uint256 amountIn,
        uint256 deadline,
        StepParams[] memory steps
    )
        external
        payable;

    /// @notice Kindly ask SIR to complete the provided intent by completing a series of Zap steps using the
    /// provided ZapRecipient contract.
    /// @dev This function is identical to `completeIntentWithBalanceChecks` except that it does not verify that
    /// the ZapRecipient balance for every token in `steps` has not increased after the last Zap.
    /// Anyone using this function must validate that the funds are fully spent by ZapRecipient
    /// using other means like separate on-chain checks or off-chain simulation.
    function completeIntent(
        address zapRecipient,
        uint256 amountIn,
        uint256 deadline,
        StepParams[] memory steps
    )
        external
        payable;
}
