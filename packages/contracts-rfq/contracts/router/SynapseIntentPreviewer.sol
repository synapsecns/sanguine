// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {ISynapseIntentRouter} from "../interfaces/ISynapseIntentRouter.sol";

contract SynapseIntentPreviewer {
    function previewIntent(
        address swapQuoter,
        address tokenIn,
        address tokenOut,
        uint256 amountIn
    )
        external
        view
        returns (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps)
    {
        // TODO: implement
    }
}
