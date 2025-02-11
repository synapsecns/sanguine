// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {UniversalTokenLib} from "../../contracts/libs/UniversalToken.sol";

// solhint-disable no-empty-blocks, ordering
contract UniversalTokenLibHarness {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testUniversalTokenLibHarness() external {}

    function universalTransfer(address token, address to, uint256 value) public {
        UniversalTokenLib.universalTransfer(token, to, value);
    }

    function universalApproveInfinity(address token, address spender, uint256 amountToSpend) public {
        UniversalTokenLib.universalApproveInfinity(token, spender, amountToSpend);
    }

    function ethAddress() public pure returns (address) {
        return UniversalTokenLib.ETH_ADDRESS;
    }

    function universalBalanceOf(address token, address account) public view returns (uint256) {
        return UniversalTokenLib.universalBalanceOf(token, account);
    }

    function assertIsContract(address token) public view {
        UniversalTokenLib.assertIsContract(token);
    }
}
