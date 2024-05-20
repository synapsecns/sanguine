// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    IInterchainClientV1,
    InterchainTransaction,
    InterchainTxDescriptor
} from "../../contracts/interfaces/IInterchainClientV1.sol";

// solhint-disable no-empty-blocks
contract InterchainClientV1Mock is IInterchainClientV1 {
    function setDefaultGuard(address guard) external {}

    function setDefaultModule(address module) external {}

    function setLinkedClient(uint64 chainId, bytes32 client) external {}

    function interchainSend(
        uint64 dstChainId,
        bytes32 receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        payable
        returns (InterchainTxDescriptor memory desc)
    {}

    function interchainSendEVM(
        uint64 dstChainId,
        address receiver,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        bytes calldata message
    )
        external
        payable
        returns (InterchainTxDescriptor memory desc)
    {}

    function interchainExecute(uint256 gasLimit, bytes calldata transaction) external payable {}

    function writeExecutionProof(bytes32 transactionId) external returns (uint64 dbNonce) {}

    function isExecutable(bytes calldata transaction) external view returns (bool) {}

    function getTxReadinessV1(InterchainTransaction memory icTx)
        external
        view
        returns (TxReadiness status, bytes32 firstArg, bytes32 secondArg)
    {}

    function getInterchainFee(
        uint64 dstChainId,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        uint256 messageLen
    )
        external
        view
        returns (uint256)
    {}

    function getExecutor(bytes calldata transaction) external view returns (address) {}

    function getExecutorById(bytes32 transactionId) external view returns (address) {}

    function getLinkedClient(uint64 chainId) external view returns (bytes32) {}

    function getLinkedClientEVM(uint64 chainId) external view returns (address) {}
}
