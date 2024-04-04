// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {MathLib} from "./Math.sol";
import {TypeCasts} from "./TypeCasts.sol";

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

    function encodeTransaction(InterchainTransaction memory transaction) internal pure returns (bytes memory) {
        return abi.encode(transaction);
    }

    function decodeTransaction(bytes memory encodedTx) internal pure returns (InterchainTransaction memory) {
        return abi.decode(encodedTx, (InterchainTransaction));
    }

    function payloadSize(uint256 optionsLen, uint256 messageLen) internal pure returns (uint256) {
        // 8 fields * 32 bytes (6 values for static, 2 offsets for dynamic) + 2 * 32 bytes (lengths for dynamic) = 320
        // abi.encode() also prepends the global offset (which is always 0x20) if there's a dynamic field, making it 352
        // Both options and message are dynamic fields, which are padded up to 32 bytes
        return 352 + optionsLen.roundUpToWord() + messageLen.roundUpToWord();
    }

    function transactionId(InterchainTransaction memory transaction) internal pure returns (bytes32) {
        return keccak256(abi.encode(transaction));
    }
}
