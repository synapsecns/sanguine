// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {MathLib} from "./Math.sol";
import {TypeCasts} from "./TypeCasts.sol";
import {VersionedPayloadLib} from "./VersionedPayload.sol";

struct InterchainTransaction {
    uint256 srcChainId;
    bytes32 srcSender;
    uint256 dstChainId;
    bytes32 dstReceiver;
    uint256 dbNonce;
    uint64 entryIndex;
    bytes options;
    bytes message;
}

struct InterchainTxDescriptor {
    bytes32 transactionId;
    uint256 dbNonce;
    uint64 entryIndex;
}

using InterchainTransactionLib for InterchainTransaction global;

library InterchainTransactionLib {
    using MathLib for uint256;
    using VersionedPayloadLib for bytes;

    function constructLocalTransaction(
        address srcSender,
        uint256 dstChainId,
        bytes32 dstReceiver,
        uint256 dbNonce,
        uint64 entryIndex,
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
            dbNonce: dbNonce,
            entryIndex: entryIndex,
            options: options,
            message: message
        });
    }

    function encodeVersionedTransaction(
        uint16 clientVersion,
        InterchainTransaction memory transaction
    )
        internal
        pure
        returns (bytes memory)
    {
        return VersionedPayloadLib.encodeVersionedPayload(clientVersion, abi.encode(transaction));
    }

    function decodeVersionedTransaction(bytes calldata versionedTx)
        internal
        pure
        returns (uint16 clientVersion, InterchainTransaction memory transaction)
    {
        clientVersion = versionedTx.getVersion();
        transaction = abi.decode(versionedTx.getPayload(), (InterchainTransaction));
    }

    function payloadSize(uint256 optionsLen, uint256 messageLen) internal pure returns (uint256) {
        // 2 bytes are reserved for the transaction version
        // + 8 fields * 32 bytes (6 values for static, 2 offsets for dynamic) + 2 * 32 bytes (lengths for dynamic) = 322
        // abi.encode() also prepends the global offset (which is always 0x20) if there's a dynamic field, making it 354
        // Both options and message are dynamic fields, which are padded up to 32 bytes
        return 354 + optionsLen.roundUpToWord() + messageLen.roundUpToWord();
    }
}
