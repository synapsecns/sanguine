// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainClientV1, InterchainTxDescriptor} from "../../contracts/interfaces/IInterchainClientV1.sol";

// solhint-disable no-empty-blocks
contract InterchainClientV1Mock is IInterchainClientV1 {
    function setExecutionFees(address executionFees_) external {}

    function setLinkedClient(uint256 chainId, bytes32 client) external {}

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

    function interchainExecute(
        uint256 gasLimit,
        bytes calldata transaction,
        bytes32[] calldata proof
    )
        external
        payable
    {}

    function writeExecutionProof(bytes32 transactionId) external returns (uint256 dbNonce, uint64 entryIndex) {}

    function isExecutable(bytes calldata transaction, bytes32[] calldata proof) external view returns (bool) {}

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

    function getLinkedClient(uint256 chainId) external view returns (bytes32) {}

    function getLinkedClientEVM(uint256 chainId) external view returns (address) {}
}
