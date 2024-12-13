// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

// solhint-disable no-empty-blocks
/// @notice Pool mock for testing purposes. DO NOT USE IN PRODUCTION.
contract PoolMock {
    using SafeERC20 for IERC20;

    address public immutable token0;
    address public immutable token1;

    uint256 public ratioWei = 1e18;

    error PoolMock__TokenNotSupported();

    constructor(address token0_, address token1_) {
        token0 = token0_;
        token1 = token1_;
    }

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testPoolMock() external {}

    function setRatioWei(uint256 ratioWei_) external {
        ratioWei = ratioWei_;
    }

    function swap(uint256 amountIn, address tokenIn) external returns (uint256 amountOut) {
        address tokenOut;
        if (tokenIn == token0) {
            tokenOut = token1;
            amountOut = amountIn * ratioWei / 1e18;
        } else if (tokenIn == token1) {
            tokenOut = token0;
            amountOut = amountIn * 1e18 / ratioWei;
        } else {
            revert PoolMock__TokenNotSupported();
        }
        IERC20(tokenIn).safeTransferFrom(msg.sender, address(this), amountIn);
        IERC20(tokenOut).safeTransfer(msg.sender, amountOut);
    }
}
