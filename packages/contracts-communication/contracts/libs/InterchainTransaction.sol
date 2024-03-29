// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {MathLib} from "./Math.sol";
import {TypeCasts} from "./TypeCasts.sol";
import {VersionedPayloadLib} from "./VersionedPayload.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

struct InterchainTransaction {
    uint64 srcChainId;
    bytes32 srcSender;
    uint64 dstChainId;
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
        uint64 dstChainId,
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
        return abi.encode(transaction);
    }

    function decodeTransaction(bytes calldata transaction) internal pure returns (InterchainTransaction memory) {
        return abi.decode(transaction, (InterchainTransaction));
    }

    function payloadSize(uint256 optionsLen, uint256 messageLen) internal pure returns (uint256) {
        // 2 bytes are reserved for the transaction version
        // + 8 fields * 32 bytes (6 values for static, 2 offsets for dynamic) + 2 * 32 bytes (lengths for dynamic) = 322
        // abi.encode() also prepends the global offset (which is always 0x20) if there's a dynamic field, making it 354
        // Both options and message are dynamic fields, which are padded up to 32 bytes
        return 354 + optionsLen.roundUpToWord() + messageLen.roundUpToWord();
    }
}
