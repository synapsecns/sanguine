// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {IDefaultPool} from "../../contracts/legacy/router/interfaces/IDefaultPool.sol";

// solhint-disable no-empty-blocks
contract DefaultPoolMock is IDefaultPool {
    uint8 private constant TOKENS = 3;

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testDefaultPoolMock() external {}

    function swap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx,
        uint256 minDy,
        uint256 deadline
    )
        external
        returns (uint256 amountOut)
    {}

    function calculateSwap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx
    )
        external
        view
        returns (uint256 amountOut)
    {}

    function swapStorage()
        external
        view
        returns (
            uint256 initialA,
            uint256 futureA,
            uint256 initialATime,
            uint256 futureATime,
            uint256 swapFee,
            uint256 adminFee,
            address lpToken
        )
    {}

    function getToken(uint8 index) external pure returns (address token) {
        if (index < TOKENS) {
            // Will be overridden by vm.mockCall
            return address(uint160(1 + index));
        }
        revert("Token does not exist");
    }
}
