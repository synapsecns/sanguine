// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1Events} from "../../contracts/events/InterchainClientV1Events.sol";

contract InterchainClientEventMock is InterchainClientV1Events {
    function emitInterchainTransactionSent(
        bytes32 transactionId,
        uint256 dbNonce,
        uint64 entryIndex,
        uint256 dstChainId,
        bytes32 srcSender,
        bytes32 dstReceiver,
        uint256 verificationFee,
        uint256 executionFee,
        bytes memory options,
        bytes memory message
    ) public {
        emit InterchainTransactionSent(transactionId, dbNonce, entryIndex, dstChainId, srcSender, dstReceiver, verificationFee, executionFee, options, message);
    }

    function emitInterchainTransactionReceived(
        bytes32 transactionId,
        uint256 dbNonce,
        uint64 entryIndex,
        uint256 srcChainId,
        bytes32 srcSender,
        bytes32 dstReceiver
    ) public {
        emit InterchainTransactionReceived(transactionId, dbNonce, entryIndex, srcChainId, srcSender, dstReceiver);
    }
}
