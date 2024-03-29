// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainTransaction, InterchainTransactionLib} from "../../contracts/libs/InterchainTransaction.sol";

contract InterchainTransactionLibHarness {
    function constructLocalTransaction(
        address srcSender,
        uint64 dstChainId,
        bytes32 dstReceiver,
        uint256 dbNonce,
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
}
