// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {HyperLiquidComposer} from "@layerzerolabs/hyperliquid-composer/contracts/HyperLiquidComposer.sol";
import {RecoverableComposer} from "@layerzerolabs/hyperliquid-composer/contracts/extensions/RecoverableComposer.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Transfers SynapseERC20 tokens received through an OFT adapter to HyperCore, with authorized recovery.
/// @dev Activate the composer on HyperCore before use. Recipients must also be activated.
contract SynapseComposer is RecoverableComposer {
    using SafeERC20 for IERC20;

    /// @param oft The HyperEVM OFT adapter for the linked SynapseERC20 token.
    /// @param coreIndexId The linked HyperCore token index.
    /// @param assetDecimalDiff ERC20 decimals minus HyperCore weiDecimals.
    /// @param recoveryAddress The account authorized to recover composer-held assets.
    constructor(
        address oft,
        uint64 coreIndexId,
        int8 assetDecimalDiff,
        address recoveryAddress
    )
        HyperLiquidComposer(oft, coreIndexId, assetDecimalDiff)
        RecoverableComposer(recoveryAddress)
    {
        // Source refunds call the adapter, which needs allowance to burnFrom the composer.
        IERC20(ERC20).forceApprove(oft, type(uint256).max);
    }
}
