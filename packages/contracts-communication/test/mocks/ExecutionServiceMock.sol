// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionService} from "../../contracts/interfaces/IExecutionService.sol";

// solhint-disable no-empty-blocks
contract ExecutionServiceMock is IExecutionService {
    address public executorEOA;

    function requestTxExecution(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        bytes memory options
    )
        external
        payable
    {}

    function requestExecution(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        uint256 executionFee,
        bytes memory options
    )
        external
    {}

    function getExecutionFee(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        returns (uint256)
    {}
}
