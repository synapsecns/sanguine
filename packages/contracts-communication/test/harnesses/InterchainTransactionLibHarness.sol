// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    InterchainTransaction, InterchainTransactionLib, ICTxHeader
} from "../../contracts/libs/InterchainTransaction.sol";

contract InterchainTransactionLibHarness {
    function constructLocalTransaction(
        address srcSender,
        uint64 dstChainId,
        bytes32 dstReceiver,
        uint64 dbNonce,
        uint64 entryIndex,
        bytes memory options,
        bytes memory message
    )
        external
        view
        returns (InterchainTransaction memory transaction)
    {
        return InterchainTransactionLib.constructLocalTransaction(
            srcSender, dstChainId, dstReceiver, dbNonce, entryIndex, options, message
        );
    }

    function encodeTransaction(InterchainTransaction memory transaction) external pure returns (bytes memory) {
        return InterchainTransactionLib.encodeTransaction(transaction);
    }

    function decodeTransaction(bytes calldata encodedTx) external pure returns (InterchainTransaction memory) {
        return InterchainTransactionLib.decodeTransaction(encodedTx);
    }

    function payloadSize(uint256 optionsLen, uint256 messageLen) external pure returns (uint256) {
        return InterchainTransactionLib.payloadSize(optionsLen, messageLen);
    }

    function encodeTxHeader(
        uint64 srcChainId,
        uint64 dstChainId,
        uint64 dbNonce,
        uint64 entryIndex
    )
        external
        pure
        returns (ICTxHeader)
    {
        return InterchainTransactionLib.encodeTxHeader(srcChainId, dstChainId, dbNonce, entryIndex);
    }

    function decodeTxHeader(ICTxHeader header)
        external
        pure
        returns (uint64 srcChainId, uint64 dstChainId, uint64 dbNonce, uint64 entryIndex)
    {
        return InterchainTransactionLib.decodeTxHeader(header);
    }
}
