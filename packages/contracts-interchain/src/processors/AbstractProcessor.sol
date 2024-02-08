// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IDefaultPool} from "../interfaces/IDefaultPool.sol";
import {IInterchainFactory} from "../interfaces/IInterchainFactory.sol";

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice AbstractProcessor is an abstraction for a contract that enables the conversion between
/// the ERC20 token (underlying) and its ICERC20 counterpart (interchain). The exact implementation
/// of the conversion mechanism is defined in the derived contracts.
/// NOTE: ICERC20 token issuance by the Bridge is rate limited by burn and mint limits,
/// while the Processor is simply a tool for Token<>InterchainToken conversion.
/// Therefore, the minting/burning of both tokens does not require a separate
/// minting/burning limit for the Processor.
/// The AbstractProcessor implements the IDefaultPool interface, allowing a seamless integration
/// with the SynapseBridge contract.
abstract contract AbstractProcessor is IDefaultPool {
    using SafeERC20 for IERC20;

    address public immutable INTERCHAIN_TOKEN;
    address public immutable UNDERLYING_TOKEN;

    error AbstractProcessor__EqualIndices(uint8 index);
    error AbstractProcessor__IndexOutOfBounds(uint8 index);

    constructor() {
        (INTERCHAIN_TOKEN, UNDERLYING_TOKEN) = IInterchainFactory(msg.sender).getProcessorDeployParameters();
    }

    /// @inheritdoc IDefaultPool
    function swap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx,
        uint256, // minDy
        uint256 // deadline
    )
        external
        returns (uint256 amountOut)
    {
        if (tokenIndexFrom == tokenIndexTo) {
            revert AbstractProcessor__EqualIndices(tokenIndexFrom);
        }
        address tokenFrom = getToken(tokenIndexFrom);
        address tokenTo = getToken(tokenIndexTo);
        // After checks above: indices are [0, 1] or [1, 0]
        // Transfer tokenFrom to this contract, use the balance difference as amountOut
        uint256 balanceBefore = IERC20(tokenFrom).balanceOf(address(this));
        IERC20(tokenFrom).safeTransferFrom(msg.sender, address(this), dx);
        amountOut = IERC20(tokenFrom).balanceOf(address(this)) - balanceBefore;
        if (tokenTo == INTERCHAIN_TOKEN) {
            _mintInterchainToken(amountOut);
        } else {
            _burnInterchainToken(amountOut);
        }
    }

    /// @inheritdoc IDefaultPool
    function calculateSwap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx
    )
        external
        pure
        returns (uint256 amountOut)
    {
        // InterchainToken (0) -> UnderlyingToken (1)
        if (tokenIndexFrom == 0 && tokenIndexTo == 1) {
            return dx;
        }
        // UnderlyingToken (1) -> InterchainToken (0)
        if (tokenIndexFrom == 1 && tokenIndexTo == 0) {
            return dx;
        }
        // Return 0 for unsupported operations
        return 0;
    }

    /// @inheritdoc IDefaultPool
    function getToken(uint8 index) public view returns (address token) {
        if (index == 0) {
            return INTERCHAIN_TOKEN;
        } else if (index == 1) {
            return UNDERLYING_TOKEN;
        }
        revert AbstractProcessor__IndexOutOfBounds(index);
    }

    /// @dev Burns the ICERC20 token taken from `msg.sender`, then
    /// transfers the same amount of the underlying token to `msg.sender`.
    function _burnInterchainToken(uint256 amount) internal virtual;

    /// @dev Handles the underlying token taken from `msg.sender`, then
    /// mints the same amount of the ICERC20 token to `msg.sender`.
    function _mintInterchainToken(uint256 amount) internal virtual;
}
