// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainTransaction, InterchainTransactionLib} from "../../contracts/libs/InterchainTransaction.sol";

contract InterchainTransactionLibHarness {
    function constructLocalTransaction(
        address srcSender,
        uint256 dstChainId,
        bytes32 dstReceiver,
        uint64 nonce,
        uint256 dbNonce,
        bytes memory options,
        bytes memory message
    )
        external
        view
        returns (InterchainTransaction memory transaction)
    {
        return InterchainTransactionLib.constructLocalTransaction(
            srcSender, dstChainId, dstReceiver, nonce, dbNonce, options, message
        );
    }

    function encodeTransaction(InterchainTransaction memory transaction) external pure returns (bytes memory) {
        return InterchainTransactionLib.encodeTransaction(transaction);
    }

    function decodeTransaction(bytes memory encodedTx) external pure returns (InterchainTransaction memory) {
        return InterchainTransactionLib.decodeTransaction(encodedTx);
    }

    function transactionId(InterchainTransaction memory transaction) external pure returns (bytes32) {
        return InterchainTransactionLib.transactionId(transaction);
    }
}
