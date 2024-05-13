// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IExecutionService {
    function requestTxExecution(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        bytes memory options
    )
        external
        payable;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function executorEOA() external view returns (address);

    function getExecutionFee(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        returns (uint256);
}
