// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IFastBridgeV2} from "../interfaces/IFastBridgeV2.sol";

// solhint-disable no-inline-assembly
library BridgeTransactionV2Lib {
    uint16 internal constant VERSION = 2;

    // Offsets of the fields in the packed BridgeTransactionV2 struct
    // uint16   version                 [000 .. 002)
    // uint32   originChainId           [002 .. 006)
    // uint32   destChainId             [006 .. 010)
    // address  originSender            [010 .. 030)
    // address  destRecipient           [030 .. 050)
    // address  originToken             [050 .. 070)
    // address  destToken               [070 .. 090)
    // uint256  originAmount            [090 .. 122)
    // uint256  destAmount              [122 .. 154)
    // uint256  originFeeAmount         [154 .. 186)
    // uint256  deadline                [186 .. 218)
    // uint256  nonce                   [218 .. 250)
    // address  exclusivityRelayer      [250 .. 270)
    // uint256  exclusivityEndTime      [270 .. 302)
    // uint256  zapNative               [302 .. 334)
    // bytes    zapData                 [334 .. ***)

    // forgefmt: disable-start
    uint256 private constant OFFSET_ORIGIN_CHAIN_ID      = 2;
    uint256 private constant OFFSET_DEST_CHAIN_ID        = 6;
    uint256 private constant OFFSET_ORIGIN_SENDER        = 10;
    uint256 private constant OFFSET_DEST_RECIPIENT       = 30;
    uint256 private constant OFFSET_ORIGIN_TOKEN         = 50;
    uint256 private constant OFFSET_DEST_TOKEN           = 70;
    uint256 private constant OFFSET_ORIGIN_AMOUNT        = 90;
    uint256 private constant OFFSET_DEST_AMOUNT          = 122;
    uint256 private constant OFFSET_ORIGIN_FEE_AMOUNT    = 154;
    uint256 private constant OFFSET_DEADLINE             = 186;
    uint256 private constant OFFSET_NONCE                = 218;
    uint256 private constant OFFSET_EXCLUSIVITY_RELAYER  = 250;
    uint256 private constant OFFSET_EXCLUSIVITY_END_TIME = 270;
    uint256 private constant OFFSET_ZAP_NATIVE           = 302;
    uint256 private constant OFFSET_ZAP_DATA             = 334;
    // forgefmt: disable-end

    error BridgeTransactionV2__InvalidEncodedTx();
    error BridgeTransactionV2__UnsupportedVersion(uint16 version);

    /// @notice Validates that the encoded transaction is a tightly packed encoded payload for BridgeTransactionV2.
    /// @dev Checks the minimum length and the version, use this function before decoding any of the fields.
    function validateV2(bytes calldata encodedTx) internal pure {
        // Check the minimum length: must at least include all static fields.
        if (encodedTx.length < OFFSET_ZAP_DATA) revert BridgeTransactionV2__InvalidEncodedTx();
        // Once we validated the length, we can be sure that the version field is present.
        uint16 version_ = version(encodedTx);
        if (version_ != VERSION) revert BridgeTransactionV2__UnsupportedVersion(version_);
    }

    /// @notice Encodes the BridgeTransactionV2 struct by tightly packing the fields.
    /// @dev `abi.decode` will not work as a result of the tightly packed fields. Use `decodeV2` to decode instead.
    function encodeV2(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) internal pure returns (bytes memory) {
        // We split the encoding into two parts to avoid stack-too-deep error
        bytes memory firstPart = abi.encodePacked(
            VERSION,
            bridgeTx.originChainId,
            bridgeTx.destChainId,
            bridgeTx.originSender,
            bridgeTx.destRecipient,
            bridgeTx.originToken,
            bridgeTx.destToken,
            bridgeTx.originAmount
        );
        return abi.encodePacked(
            firstPart,
            bridgeTx.destAmount,
            bridgeTx.originFeeAmount,
            // Note: we skip the deprecated `sendChainGas` flag, which was present in BridgeTransaction V1
            bridgeTx.deadline,
            bridgeTx.nonce,
            // New V2 fields: exclusivity
            bridgeTx.exclusivityRelayer,
            bridgeTx.exclusivityEndTime,
            // New V2 fields: Zap
            bridgeTx.zapNative,
            bridgeTx.zapData
        );
    }

    /// @notice Decodes the BridgeTransactionV2 struct from the encoded transaction.
    /// @dev Encoded BridgeTransactionV2 struct must be tightly packed.
    /// Use `validateV2` before decoding to ensure the encoded transaction is valid.
    function decodeV2(bytes calldata encodedTx)
        internal
        pure
        returns (IFastBridgeV2.BridgeTransactionV2 memory bridgeTx)
    {
        bridgeTx.originChainId = originChainId(encodedTx);
        bridgeTx.destChainId = destChainId(encodedTx);
        bridgeTx.originSender = originSender(encodedTx);
        bridgeTx.destRecipient = destRecipient(encodedTx);
        bridgeTx.originToken = originToken(encodedTx);
        bridgeTx.destToken = destToken(encodedTx);
        bridgeTx.originAmount = originAmount(encodedTx);
        bridgeTx.destAmount = destAmount(encodedTx);
        bridgeTx.originFeeAmount = originFeeAmount(encodedTx);
        bridgeTx.deadline = deadline(encodedTx);
        bridgeTx.nonce = nonce(encodedTx);
        bridgeTx.exclusivityRelayer = exclusivityRelayer(encodedTx);
        bridgeTx.exclusivityEndTime = exclusivityEndTime(encodedTx);
        bridgeTx.zapNative = zapNative(encodedTx);
        bridgeTx.zapData = zapData(encodedTx);
    }

    /// @notice Extracts the version from the encoded transaction.
    function version(bytes calldata encodedTx) internal pure returns (uint16 version_) {
        // Load 32 bytes from the start and shift it 240 bits to the right to get the highest 16 bits.
        assembly {
            version_ := shr(240, calldataload(encodedTx.offset))
        }
    }

    /// @notice Extracts the origin chain ID from the encoded transaction.
    function originChainId(bytes calldata encodedTx) internal pure returns (uint32 originChainId_) {
        // Load 32 bytes from the offset and shift it 224 bits to the right to get the highest 32 bits.
        assembly {
            originChainId_ := shr(224, calldataload(add(encodedTx.offset, OFFSET_ORIGIN_CHAIN_ID)))
        }
    }

    /// @notice Extracts the destination chain ID from the encoded transaction.
    function destChainId(bytes calldata encodedTx) internal pure returns (uint32 destChainId_) {
        // Load 32 bytes from the offset and shift it 224 bits to the right to get the highest 32 bits.
        assembly {
            destChainId_ := shr(224, calldataload(add(encodedTx.offset, OFFSET_DEST_CHAIN_ID)))
        }
    }

    /// @notice Extracts the origin sender from the encoded transaction.
    function originSender(bytes calldata encodedTx) internal pure returns (address originSender_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            originSender_ := shr(96, calldataload(add(encodedTx.offset, OFFSET_ORIGIN_SENDER)))
        }
    }

    /// @notice Extracts the destination recipient from the encoded transaction.
    function destRecipient(bytes calldata encodedTx) internal pure returns (address destRecipient_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            destRecipient_ := shr(96, calldataload(add(encodedTx.offset, OFFSET_DEST_RECIPIENT)))
        }
    }

    /// @notice Extracts the origin token from the encoded transaction.
    function originToken(bytes calldata encodedTx) internal pure returns (address originToken_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            originToken_ := shr(96, calldataload(add(encodedTx.offset, OFFSET_ORIGIN_TOKEN)))
        }
    }

    /// @notice Extracts the destination token from the encoded transaction.
    function destToken(bytes calldata encodedTx) internal pure returns (address destToken_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            destToken_ := shr(96, calldataload(add(encodedTx.offset, OFFSET_DEST_TOKEN)))
        }
    }

    /// @notice Extracts the origin amount from the encoded transaction.
    function originAmount(bytes calldata encodedTx) internal pure returns (uint256 originAmount_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            originAmount_ := calldataload(add(encodedTx.offset, OFFSET_ORIGIN_AMOUNT))
        }
    }

    /// @notice Extracts the destination amount from the encoded transaction.
    function destAmount(bytes calldata encodedTx) internal pure returns (uint256 destAmount_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            destAmount_ := calldataload(add(encodedTx.offset, OFFSET_DEST_AMOUNT))
        }
    }

    /// @notice Extracts the origin fee amount from the encoded transaction.
    function originFeeAmount(bytes calldata encodedTx) internal pure returns (uint256 originFeeAmount_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            originFeeAmount_ := calldataload(add(encodedTx.offset, OFFSET_ORIGIN_FEE_AMOUNT))
        }
    }

    /// @notice Extracts the deadline from the encoded transaction.
    function deadline(bytes calldata encodedTx) internal pure returns (uint256 deadline_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            deadline_ := calldataload(add(encodedTx.offset, OFFSET_DEADLINE))
        }
    }

    /// @notice Extracts the nonce from the encoded transaction.
    function nonce(bytes calldata encodedTx) internal pure returns (uint256 nonce_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            nonce_ := calldataload(add(encodedTx.offset, OFFSET_NONCE))
        }
    }

    /// @notice Extracts the exclusivity relayer from the encoded transaction.
    function exclusivityRelayer(bytes calldata encodedTx) internal pure returns (address exclusivityRelayer_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            exclusivityRelayer_ := shr(96, calldataload(add(encodedTx.offset, OFFSET_EXCLUSIVITY_RELAYER)))
        }
    }

    /// @notice Extracts the exclusivity end time from the encoded transaction.
    function exclusivityEndTime(bytes calldata encodedTx) internal pure returns (uint256 exclusivityEndTime_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            exclusivityEndTime_ := calldataload(add(encodedTx.offset, OFFSET_EXCLUSIVITY_END_TIME))
        }
    }

    /// @notice Extracts the Zap's native value from the encoded transaction.
    function zapNative(bytes calldata encodedTx) internal pure returns (uint256 zapNative_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            zapNative_ := calldataload(add(encodedTx.offset, OFFSET_ZAP_NATIVE))
        }
    }

    /// @notice Extracts the Zap's data from the encoded transaction.
    function zapData(bytes calldata encodedTx) internal pure returns (bytes calldata zapData_) {
        zapData_ = encodedTx[OFFSET_ZAP_DATA:];
    }
}
