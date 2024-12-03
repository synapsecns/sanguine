// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {ISynapseIntentRouter} from "../interfaces/ISynapseIntentRouter.sol";
import {ISynapseIntentRouterErrors} from "../interfaces/ISynapseIntentRouterErrors.sol";

contract SynapseIntentRouter is ISynapseIntentRouter, ISynapseIntentRouterErrors {
    /// @inheritdoc ISynapseIntentRouter
    function completeIntent(
        address zapRecipient,
        uint256 amountIn,
        uint256 minLastStepAmountIn,
        uint256 deadline,
        StepParams[] calldata steps
    )
        external
        payable
    {
        // TODO: implement
    }
}
