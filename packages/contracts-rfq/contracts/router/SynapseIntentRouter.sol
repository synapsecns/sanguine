// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {ISynapseIntentRouter} from "../interfaces/ISynapseIntentRouter.sol";
import {ISynapseIntentRouterErrors} from "../interfaces/ISynapseIntentRouterErrors.sol";
import {IZapRecipient} from "../interfaces/IZapRecipient.sol";

// ═════════════════════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════════════════════════

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

contract SynapseIntentRouter is ISynapseIntentRouter, ISynapseIntentRouterErrors {
    using SafeERC20 for IERC20;

    /// @notice The address reserved for the native gas token (ETH on Ethereum and most L2s, AVAX on Avalanche, etc.).
    address public constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /// @dev Amount value that signals that the Zap step should be performed using the full ZapRecipient balance.
    uint256 internal constant FULL_BALANCE = type(uint256).max;

    /// @inheritdoc ISynapseIntentRouter
    function completeIntentWithBalanceChecks(
        address zapRecipient,
        uint256 amountIn,
        uint256 minLastStepAmountIn,
        uint256 deadline,
        StepParams[] calldata steps
    )
        external
        payable
    {
        // TODO: record and check balances
        completeIntent(zapRecipient, amountIn, minLastStepAmountIn, deadline, steps);
    }

    /// @inheritdoc ISynapseIntentRouter
    function completeIntent(
        address zapRecipient,
        uint256 amountIn,
        uint256 minLastStepAmountIn,
        uint256 deadline,
        StepParams[] calldata steps
    )
        public
        payable
    {
        // Validate the input parameters before proceeding.
        uint256 length = steps.length;
        if (block.timestamp > deadline) revert SIR__DeadlineExceeded();
        if (length == 0) revert SIR__StepsNotProvided();

        // Transfer the input asset from the user to ZapRecipient. `steps[0]` exists as per check above.
        _transferInputAsset(zapRecipient, steps[0].token, amountIn);

        // Perform the Zap steps, using predetermined amounts or the full balance of ZapRecipient, if instructed.
        uint256 totalUsedMsgValue = 0;
        for (uint256 i = 0; i < length; i++) {
            address token = steps[i].token;
            uint256 msgValue = steps[i].msgValue;

            // Adjust amount to be the full balance, if needed.
            amountIn = steps[i].amount;
            if (amountIn == FULL_BALANCE) {
                amountIn = token == NATIVE_GAS_TOKEN
                    // Existing native balance + msg.value that will be forwarded
                    ? zapRecipient.balance + msgValue
                    : IERC20(token).balanceOf(zapRecipient);
            }

            _performZap({
                zapRecipient: zapRecipient,
                msgValue: msgValue,
                zapRecipientCallData: abi.encodeCall(IZapRecipient.zap, (token, amountIn, steps[i].zapData))
            });
            unchecked {
                // Can do unchecked addition here since we're guaranteed that the sum of all msg.value
                // used for the Zaps won't overflow.
                totalUsedMsgValue += msgValue;
            }
        }

        // Verify amountIn used for the last step, and that we fully spent `msg.value`.
        if (amountIn < minLastStepAmountIn) revert SIR__AmountInsufficient();
        if (totalUsedMsgValue < msg.value) revert SIR__MsgValueIncorrect();
    }

    // ═════════════════════════════════════════════ INTERNAL METHODS ══════════════════════════════════════════════════

    /// @notice Transfers the input asset from the user into ZapRecipient custody. This asset will later be
    /// used to perform the zap steps.
    function _transferInputAsset(address zapRecipient, address token, uint256 amount) internal {
        if (token == NATIVE_GAS_TOKEN) {
            // For the native gas token, we just need to check that the supplied `msg.value` is correct.
            // We will later forward `msg.value` in the series of the steps using `StepParams.msgValue`.
            if (amount != msg.value) revert SIR__MsgValueIncorrect();
        } else {
            // For ERC20s, token is transferred from the user to ZapRecipient before performing the zap steps.
            // Throw an explicit error if the provided token address is not a contract.
            if (token.code.length == 0) revert SIR__TokenNotContract();
            IERC20(token).safeTransferFrom(msg.sender, zapRecipient, amount);
        }
    }

    /// @notice Performs a Zap step, using the provided msg.value and calldata.
    /// Validates the return data from ZapRecipient as per `IZapRecipient` specification.
    function _performZap(address zapRecipient, uint256 msgValue, bytes memory zapRecipientCallData) internal {
        // Perform the low-level call to ZapRecipient, bubbling up any revert reason.
        bytes memory returnData =
            Address.functionCallWithValue({target: zapRecipient, data: zapRecipientCallData, value: msgValue});

        // Explicit revert if no return data at all.
        if (returnData.length == 0) revert SIR__ZapNoReturnValue();
        // Check that exactly a single return value was returned.
        if (returnData.length != 32) revert SIR__ZapIncorrectReturnValue();
        // Return value should be abi-encoded hook function selector.
        if (bytes32(returnData) != bytes32(IZapRecipient.zap.selector)) {
            revert SIR__ZapIncorrectReturnValue();
        }
    }
}
