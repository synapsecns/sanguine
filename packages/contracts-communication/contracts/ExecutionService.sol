// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;


import { IExecutionService } from "./interfaces/IExecutionService.sol";
import { ExecutionServiceEvents } from "./events/ExecutionServiceEvents.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";


contract ExecutionService is ExecutionServiceEvents, Ownable, IExecutionService {
    address public executorEOA;

    constructor() Ownable(msg.sender) {}

    function setExecutorEOA(address _executorEOA) external onlyOwner {
        executorEOA = _executorEOA;
        emit ExecutorEOAUpdated(executorEOA);
    }

    function requestExecution(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        uint256 executionFee,
        bytes memory options
    )
        external
        override
    {
        emit ExecutionRequested(dstChainId, txPayloadSize, transactionId, executionFee, options);
    }

    function getExecutionFee(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        override
        returns (uint256)
    {
        // ...
    }
}
