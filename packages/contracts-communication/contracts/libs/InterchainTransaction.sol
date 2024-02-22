// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {TypeCasts} from "./TypeCasts.sol";

struct InterchainTransaction {
    uint256 srcChainId;
    bytes32 srcSender;
    uint256 dstChainId;
    bytes32 dstReceiver;
    uint64 nonce;
    uint256 dbNonce;
    bytes options;
    bytes message;
}

using InterchainTransactionLib for InterchainTransaction global;

library InterchainTransactionLib {
    function constructLocalTransaction(
        address srcSender,
        uint256 dstChainId,
        bytes32 dstReceiver,
        uint64 nonce,
        uint256 dbNonce,
        bytes memory options,
        bytes memory message
    )
        internal
        view
        returns (InterchainTransaction memory transaction)
    {
        return InterchainTransaction({
            srcChainId: block.chainid,
            srcSender: TypeCasts.addressToBytes32(srcSender),
            dstChainId: dstChainId,
            dstReceiver: dstReceiver,
            nonce: nonce,
            dbNonce: dbNonce,
            options: options,
            message: message
        });
    }

    function encodeTransaction(InterchainTransaction memory transaction) internal pure returns (bytes memory) {
        return abi.encode(transaction);
    }

    function decodeTransaction(bytes memory encodedTx) internal pure returns (InterchainTransaction memory) {
        return abi.decode(encodedTx, (InterchainTransaction));
    }

    function transactionId(InterchainTransaction memory transaction) internal pure returns (bytes32) {
        return keccak256(abi.encode(transaction));
    }
}
