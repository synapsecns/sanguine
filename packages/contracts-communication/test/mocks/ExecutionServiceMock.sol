// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionService} from "../../contracts/interfaces/IExecutionService.sol";

// solhint-disable no-empty-blocks
contract ExecutionServiceMock is IExecutionService {
    address public executorEOA;

    function requestExecution(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        uint256 executionFee,
        bytes memory options
    )
        external
    {}

    function getExecutionFee(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        returns (uint256)
    {}
}
