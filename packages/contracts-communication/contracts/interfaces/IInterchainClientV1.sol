// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainTransaction, InterchainTxDescriptor} from "../libs/InterchainTransaction.sol";

interface IInterchainClientV1 {
    enum TxReadiness {
        Ready,
        AlreadyExecuted,
        EntryAwaitingResponses,
        EntryConflict,
        ReceiverNotICApp,
        ReceiverZeroRequiredResponses,
        TxWrongDstChainId,
        UndeterminedRevert
    }

    error InterchainClientV1__ChainIdNotLinked(uint64 chainId);
    error InterchainClientV1__ChainIdNotRemote(uint64 chainId);
    error InterchainClientV1__DstChainIdNotLocal(uint64 chainId);
    error InterchainClientV1__EntryConflict(address module);
    error InterchainClientV1__ExecutionServiceZeroAddress();
    error InterchainClientV1__FeeAmountBelowMin(uint256 feeAmount, uint256 minRequired);
    error InterchainClientV1__GasLeftBelowMin(uint256 gasLeft, uint256 minRequired);
    error InterchainClientV1__GuardZeroAddress();
    error InterchainClientV1__LinkedClientNotEVM(bytes32 client);
    error InterchainClientV1__MsgValueMismatch(uint256 msgValue, uint256 required);
    error InterchainClientV1__ReceiverNotICApp(address receiver);
    error InterchainClientV1__ReceiverZeroAddress();
    error InterchainClientV1__ReceiverZeroRequiredResponses(address receiver);
    error InterchainClientV1__ResponsesAmountBelowMin(uint256 responsesAmount, uint256 minRequired);
    error InterchainClientV1__TxAlreadyExecuted(bytes32 transactionId);
    error InterchainClientV1__TxNotExecuted(bytes32 transactionId);
    error InterchainClientV1__TxVersionMismatch(uint16 txVersion, uint16 required);

    function setDefaultGuard(address guard_) external;
    function setLinkedClient(uint64 chainId, bytes32 client) external;

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
        returns (InterchainTxDescriptor memory desc);

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
        returns (InterchainTxDescriptor memory desc);

    function interchainExecute(uint256 gasLimit, bytes calldata transaction) external payable;

    function writeExecutionProof(bytes32 transactionId) external returns (uint64 dbNonce);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function isExecutable(bytes calldata transaction) external view returns (bool);
    function getTxReadinessV1(InterchainTransaction memory icTx)
        external
        view
        returns (TxReadiness status, bytes32 firstArg, bytes32 secondArg);

    function getInterchainFee(
        uint64 dstChainId,
        address srcExecutionService,
        address[] calldata srcModules,
        bytes calldata options,
        uint256 messageLen
    )
        external
        view
        returns (uint256);

    function getExecutor(bytes calldata transaction) external view returns (address);
    function getExecutorById(bytes32 transactionId) external view returns (address);
    function getLinkedClient(uint64 chainId) external view returns (bytes32);
    function getLinkedClientEVM(uint64 chainId) external view returns (address);
}
