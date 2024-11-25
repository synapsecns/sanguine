// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {IZapRouterV1} from "../interfaces/IZapRouterV1.sol";
import {IZapRouterV1Errors} from "../interfaces/IZapRouterV1Errors.sol";

contract ZapRouterV1 is IZapRouterV1, IZapRouterV1Errors {
    /// @inheritdoc IZapRouterV1
    function performZaps(
        address zapRecipient,
        uint256 amountIn,
        uint256 minLastZapAmountIn,
        uint256 deadline,
        ZapParams[] calldata zapParams
    )
        external
        payable
    {
        // TODO: implement
    }
}
