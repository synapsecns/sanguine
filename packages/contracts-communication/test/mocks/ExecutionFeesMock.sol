// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionFees} from "../../contracts/interfaces/IExecutionFees.sol";

contract ExecutionFeesMock is IExecutionFees {
    function addExecutionFee(uint256 dstChainId, bytes32 transactionId) external payable {}

    function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) external {}

    function claimExecutionFees(address executor) external {}

    function accumulatedRewards(address executor) external view returns (uint256 accumulated) {}

    function unclaimedRewards(address executor) external view returns (uint256 unclaimed) {}

    function executionFee(uint256 dstChainId, bytes32 transactionId) external view returns (uint256 fee) {}

    function recordedExecutor(uint256 dstChainId, bytes32 transactionId) external view returns (address executor) {}
}
