// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {MathLib} from "./Math.sol";
import {TypeCasts} from "./TypeCasts.sol";
import {VersionedPayloadLib} from "./VersionedPayload.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

type ICTxHeader is uint256;

struct InterchainTransaction {
    uint64 srcChainId;
    uint64 dstChainId;
    uint64 dbNonce;
    uint64 entryIndex;
    bytes32 srcSender;
    bytes32 dstReceiver;
    bytes options;
    bytes message;
}

struct InterchainTxDescriptor {
    bytes32 transactionId;
    uint64 dbNonce;
    uint64 entryIndex;
}

using InterchainTransactionLib for InterchainTransaction global;

library InterchainTransactionLib {
    using MathLib for uint256;
    using VersionedPayloadLib for bytes;

    function constructLocalTransaction(
        address srcSender,
        uint64 dstChainId,
        bytes32 dstReceiver,
        uint64 dbNonce,
        uint64 entryIndex,
        bytes memory options,
        bytes memory message
    )
        internal
        view
        returns (InterchainTransaction memory transaction)
    {
        return InterchainTransaction({
            srcChainId: SafeCast.toUint64(block.chainid),
            srcSender: TypeCasts.addressToBytes32(srcSender),
            dstChainId: dstChainId,
            dstReceiver: dstReceiver,
            dbNonce: dbNonce,
            entryIndex: entryIndex,
            options: options,
            message: message
        });
    }

    function encodeTransaction(InterchainTransaction memory transaction) internal pure returns (bytes memory) {
        return abi.encode(
            encodeTxHeader(transaction.srcChainId, transaction.dstChainId, transaction.dbNonce, transaction.entryIndex),
            transaction.srcSender,
            transaction.dstReceiver,
            transaction.options,
            transaction.message
        );
    }

    function decodeTransaction(bytes calldata transaction) internal pure returns (InterchainTransaction memory icTx) {
        ICTxHeader header;
        (header, icTx.srcSender, icTx.dstReceiver, icTx.options, icTx.message) =
            abi.decode(transaction, (ICTxHeader, bytes32, bytes32, bytes, bytes));
        (icTx.srcChainId, icTx.dstChainId, icTx.dbNonce, icTx.entryIndex) = decodeTxHeader(header);
    }

    function payloadSize(uint256 optionsLen, uint256 messageLen) internal pure returns (uint256) {
        // 2 bytes are reserved for the transaction version
        // + 5 fields * 32 bytes (3 values for static, 2 offsets for dynamic) + 2 * 32 bytes (lengths for dynamic) = 226
        // (srcChainId, dstChainId, dbNonce, entryIndex) are merged into a single 32 bytes field
        // Both options and message are dynamic fields, which are padded up to 32 bytes
        return 226 + optionsLen.roundUpToWord() + messageLen.roundUpToWord();
    }

    function encodeTxHeader(
        uint64 srcChainId,
        uint64 dstChainId,
        uint64 dbNonce,
        uint64 entryIndex
    )
        internal
        pure
        returns (ICTxHeader)
    {
        return ICTxHeader.wrap(
            (uint256(srcChainId) << 192) | (uint256(dstChainId) << 128) | (uint256(dbNonce) << 64) | uint256(entryIndex)
        );
    }

    function decodeTxHeader(ICTxHeader header)
        internal
        pure
        returns (uint64 srcChainId, uint64 dstChainId, uint64 dbNonce, uint64 entryIndex)
    {
        srcChainId = uint64(ICTxHeader.unwrap(header) >> 192);
        dstChainId = uint64(ICTxHeader.unwrap(header) >> 128);
        dbNonce = uint64(ICTxHeader.unwrap(header) >> 64);
        entryIndex = uint64(ICTxHeader.unwrap(header));
    }
}
