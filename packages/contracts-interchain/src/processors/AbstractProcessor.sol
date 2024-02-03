// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IDefaultPool} from "../interfaces/IDefaultPool.sol";
import {InterchainERC20} from "../interfaces/InterchainERC20.sol";

/// @notice AbstractProcessor is an abstraction for a contract that enables the conversion between
/// the ERC20 token (underlying) and its InterchainERC20 counterpart. The exact implementation
/// of the conversion mechanism is defined in the derived contracts.
/// NOTE: InterchainERC20 token issuance by the Bridge is rate limited by burn and mint limits,
/// while the Processor is simply a tool for Token<>InterchainToken conversion.
/// Therefore, the minting/burning of both tokens does not require a separate
/// minting/burning limit for the Processor.
/// The AbstractProcessor implements the IDefaultPool interface, allowing a seamless integration
/// with the SynapseBridge contract.
abstract contract AbstractProcessor is IDefaultPool {
    address public immutable interchainToken;
    address public immutable underlyingToken;

    constructor(address interchainToken_, address underlyingToken_) {
        interchainToken = interchainToken_;
        underlyingToken = underlyingToken_;
    }

    /// @inheritdoc IDefaultPool
    function swap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx,
        uint256 minDy,
        uint256 deadline
    )
        external
        returns (uint256 amountOut)
    {
        // TODO: implement
    }

    /// @inheritdoc IDefaultPool
    function calculateSwap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx
    )
        external
        view
        returns (uint256 amountOut)
    {
        // TODO: implement
    }

    /// @inheritdoc IDefaultPool
    function getToken(uint8 index) external view returns (address token) {
        // TODO: implement
    }
}
