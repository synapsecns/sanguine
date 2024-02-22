// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { IExecutionFees } from "./interfaces/IExecutionFees.sol";

contract ExecutionFees is IExecutionFees {
    // Interchain Transaction IDs => Total Execution Fees
    mapping(bytes32 => uint256) private _executionFees;
    // Executor Addresses => Total Accumulated Fees
    mapping(address => uint256) private _accumulatedRewards;
    // Executor Addresses => Currently unclaimed fees
    mapping(address => uint256) private _unclaimedRewards;
    // Interchain Transaction IDs => Executor Addresses
    mapping(bytes32 => address) private _transactionsByExecutor;

    address public icClient;

    constructor (address _icClient) {
        icClient = _icClient;
    }

    modifier onlyRecorder() {
        // This is currently set to InterchainClientV1, but will be moved to batched recording later on
        require(msg.sender == icClient, "ExecutionFees: Caller is not the recorder");
        _;
    }

    function addExecutionFee(uint256 dstChainId, bytes32 transactionId) external payable override {
        require(msg.value > 0, "ExecutionFees: Fee must be greater than 0");
        require(_transactionsByExecutor[transactionId] == address(0), "ExecutionFees: Executor already recorded");
        _executionFees[transactionId] += msg.value;
    }

    function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) external override onlyRecorder {
        require(_transactionsByExecutor[transactionId] == address(0), "ExecutionFees: Executor already recorded");
        require(_executionFees[transactionId] > 0, "ExecutionFees: No execution fee found");
        _transactionsByExecutor[transactionId] = executor;
        _accumulatedRewards[executor] += _executionFees[transactionId];
        _unclaimedRewards[executor] += _executionFees[transactionId];
    }

    function claimExecutionFees() external override {
        uint256 amount = _unclaimedRewards[msg.sender];
        require(amount > 0, "ExecutionFees: No unclaimed rewards");
        _unclaimedRewards[msg.sender] = 0;
        payable(msg.sender).transfer(amount);
    }

    function getAccumulatedRewards(address executor) external view override returns (uint256 accumulated) {
        return _accumulatedRewards[executor];
    }

    function getUnclaimedRewards(address executor) external view override returns (uint256 unclaimed) {
        return _unclaimedRewards[executor];
    }
}
