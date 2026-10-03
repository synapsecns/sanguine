// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridge} from "../../src/interfaces/ISynapseBridge.sol";
import {TestToken} from "./TestToken.sol";

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

// solhint-disable no-empty-blocks
contract SynapseBridgeMock is ISynapseBridge {
    using SafeERC20 for IERC20;

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testSynapseBridgeMock() external {}

    function mint(address to, address token, uint256 amount, uint256, bytes32) external override {
        TestToken(token).mintTestTokens(to, amount);
    }

    function withdraw(address to, address token, uint256 amount, uint256, bytes32) external override {
        IERC20(token).safeTransfer(to, amount);
    }
}
