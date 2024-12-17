// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {ISwapQuoter, LimitedToken, SwapQuery} from "../../contracts/legacy/rfq/interfaces/ISwapQuoter.sol";

// solhint-disable no-empty-blocks
contract SwapQuoterMock is ISwapQuoter {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testSwapQuoterMock() external {}

    function getAmountOut(
        LimitedToken memory tokenIn,
        address tokenOut,
        uint256 amountIn
    )
        external
        view
        returns (SwapQuery memory query)
    {}
}
