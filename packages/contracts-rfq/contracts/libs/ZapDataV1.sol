// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// solhint-disable no-inline-assembly
library ZapDataV1 {
    /// @notice Version of the Zap Data struct.
    uint16 internal constant VERSION = 1;

    /// @notice Value that indicates the amount is not present in the target function's payload.
    uint16 internal constant AMOUNT_NOT_PRESENT = 0xFFFF;

    // Offsets of the fields in the packed ZapData struct
    // uint16   version                 [000 .. 002)
    // uint16   amountPosition          [002 .. 004)
    // address  finalToken              [004 .. 024)
    // address  forwardTo               [024 .. 044)
    // uint256  minFinalBalance         [044 .. 076)
    // address  target                  [076 .. 096)
    // bytes    payload                 [096 .. ***)

    // forgefmt: disable-start
    uint256 private constant OFFSET_AMOUNT_POSITION   = 2;
    uint256 private constant OFFSET_FINAL_TOKEN       = 4;
    uint256 private constant OFFSET_FORWARD_TO        = 24;
    uint256 private constant OFFSET_MIN_FINAL_BALANCE = 44;
    uint256 private constant OFFSET_TARGET            = 76;
    uint256 private constant OFFSET_PAYLOAD           = 96;
    // forgefmt: disable-end

    error ZapDataV1__InvalidEncoding();
    error ZapDataV1__FinalTokenNotSpecified();
    error ZapDataV1__PayloadLengthAboveMax();
    error ZapDataV1__TargetZeroAddress();
    error ZapDataV1__UnsupportedVersion(uint16 version);

    /// @notice Validates that encodedZapData is a tightly packed encoded payload for ZapData struct.
    /// @dev Checks that all the required fields are present and the version is correct.
    function validateV1(bytes calldata encodedZapData) internal pure {
        // Check the minimum length: must at least include all static fields.
        if (encodedZapData.length < OFFSET_PAYLOAD) revert ZapDataV1__InvalidEncoding();
        // Once we validated the length, we can be sure that the version field is present.
        uint16 version_ = version(encodedZapData);
        if (version_ != VERSION) revert ZapDataV1__UnsupportedVersion(version_);
    }

    /// @notice Encodes the ZapData struct by tightly packing the fields.
    /// Note: we don't know the exact amount of tokens that will be used for the Zap at the time of encoding,
    /// so we provide the reference index where the token amount is encoded within `payload_`. This allows us to
    /// hot-swap the token amount in the payload, when the Zap is performed.
    /// @dev `abi.decode` will not work as a result of the tightly packed fields. Use `decodeZapData` instead.
    /// @param amountPosition_  Position (start index) where the token amount is encoded within `payload_`.
    ///                         This will usually be `4 + 32 * n`, where `n` is the position of the token amount in
    ///                         the list of parameters of the target function (starting from 0).
    ///                         Or `AMOUNT_NOT_PRESENT` if the token amount is not encoded within `payload_`.
    /// @param finalToken_      The token produced as a result of the Zap action (ERC20 or native gas token).
    ///                         A zero address value signals that the Zap action doesn't result in any asset per se,
    ///                         like bridging or depositing into a vault without an LP token.
    ///                         Note: this parameter must be set to a non-zero value if either `forwardTo_` or
    ///                         `minFinalBalance_` is set to a non-zero value.
    /// @param forwardTo_       The address to which `finalToken` should be forwarded. This parameter is required only
    ///                         if the Zap action does not automatically transfer the token to the intended recipient.
    ///                         Otherwise, it must be set to address(0).
    /// @param minFinalBalance_ The minimum balance of `finalToken` to have after performing the Zap action.
    ///                         This parameter could be used to implement slippage protection for a Zap action that does
    ///                         not automatically transfer the token to the intended recipient. Otherwise, it must be
    ///                         set to 0.
    /// @param target_          Address of the target contract.
    /// @param payload_         ABI-encoded calldata to be used for the `target_` contract call.
    ///                         If the target function has the token amount as an argument, any placeholder amount value
    ///                         can be used for the original ABI encoding of `payload_`. The placeholder amount will
    ///                         be replaced with the actual amount, when the Zap Data is decoded.
    function encodeV1(
        uint16 amountPosition_,
        address finalToken_,
        address forwardTo_,
        uint256 minFinalBalance_,
        address target_,
        bytes memory payload_
    )
        internal
        pure
        returns (bytes memory encodedZapData)
    {
        if (target_ == address(0)) revert ZapDataV1__TargetZeroAddress();
        // We use uint16 to represent the amount position within the payload, so its length must fit in uint16.
        if (payload_.length > AMOUNT_NOT_PRESENT) revert ZapDataV1__PayloadLengthAboveMax();
        // Amount is encoded in [amountPosition_ .. amountPosition_ + 32), which should be within the payload.
        if (amountPosition_ != AMOUNT_NOT_PRESENT && (uint256(amountPosition_) + 32 > payload_.length)) {
            revert ZapDataV1__InvalidEncoding();
        }
        // Final token needs to be specified if forwarding is required.
        if (forwardTo_ != address(0) && finalToken_ == address(0)) {
            revert ZapDataV1__FinalTokenNotSpecified();
        }
        // Final token needs to be specified if minimum balance check is required.
        if (minFinalBalance_ != 0 && finalToken_ == address(0)) {
            revert ZapDataV1__FinalTokenNotSpecified();
        }
        return abi.encodePacked(VERSION, amountPosition_, finalToken_, forwardTo_, minFinalBalance_, target_, payload_);
    }

    /// @notice Extracts the version from the encoded Zap Data.
    function version(bytes calldata encodedZapData) internal pure returns (uint16 version_) {
        // Load 32 bytes from the start and shift it 240 bits to the right to get the highest 16 bits.
        assembly {
            version_ := shr(240, calldataload(encodedZapData.offset))
        }
    }

    /// @notice Extracts the finalToken address from the encoded Zap Data.
    function finalToken(bytes calldata encodedZapData) internal pure returns (address finalToken_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            finalToken_ := shr(96, calldataload(add(encodedZapData.offset, OFFSET_FINAL_TOKEN)))
        }
    }

    /// @notice Extracts the forwardTo address from the encoded Zap Data.
    function forwardTo(bytes calldata encodedZapData) internal pure returns (address forwardTo_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            forwardTo_ := shr(96, calldataload(add(encodedZapData.offset, OFFSET_FORWARD_TO)))
        }
    }

    /// @notice Extracts the minimum final balance from the encoded Zap Data.
    function minFinalBalance(bytes calldata encodedZapData) internal pure returns (uint256 minFinalBalance_) {
        // Load 32 bytes from the offset. No shift is applied, as we need the full 256 bits.
        assembly {
            minFinalBalance_ := calldataload(add(encodedZapData.offset, OFFSET_MIN_FINAL_BALANCE))
        }
    }

    /// @notice Extracts the target address from the encoded Zap Data.
    function target(bytes calldata encodedZapData) internal pure returns (address target_) {
        // Load 32 bytes from the offset and shift it 96 bits to the right to get the highest 160 bits.
        assembly {
            target_ := shr(96, calldataload(add(encodedZapData.offset, OFFSET_TARGET)))
        }
    }

    /// @notice Extracts the payload from the encoded Zap Data. Replaces the token amount with the provided value,
    /// if it was present in the original data (if amountPosition is not AMOUNT_NOT_PRESENT).
    /// @dev This payload will be used as a calldata for the target contract.
    function payload(bytes calldata encodedZapData, uint256 amount) internal pure returns (bytes memory) {
        // The original payload is located at encodedZapData[OFFSET_PAYLOAD:].
        uint16 amountPosition = _amountPosition(encodedZapData);
        // If the amount was not present in the original payload, return the payload as is.
        if (amountPosition == AMOUNT_NOT_PRESENT) {
            return encodedZapData[OFFSET_PAYLOAD:];
        }
        // Calculate the start and end indexes of the amount in ZapData from its position within the payload.
        // Note: we use inclusive start and exclusive end indexes for easier slicing of the ZapData.
        uint256 amountStartIndexIncl = OFFSET_PAYLOAD + amountPosition;
        uint256 amountEndIndexExcl = amountStartIndexIncl + 32;
        // Check that the amount is within the ZapData.
        if (amountEndIndexExcl > encodedZapData.length) revert ZapDataV1__InvalidEncoding();
        // Otherwise we need to replace the amount in the payload with the provided value.
        return abi.encodePacked(
            // Copy the original payload up to the amount
            encodedZapData[OFFSET_PAYLOAD:amountStartIndexIncl],
            // Replace the originally encoded amount with the provided value
            amount,
            // Copy the rest of the payload after the amount
            encodedZapData[amountEndIndexExcl:]
        );
    }

    /// @notice Extracts the amount position from the encoded Zap Data.
    function _amountPosition(bytes calldata encodedZapData) private pure returns (uint16 amountPosition) {
        // Load 32 bytes from the offset and shift it 240 bits to the right to get the highest 16 bits.
        assembly {
            amountPosition := shr(240, calldataload(add(encodedZapData.offset, OFFSET_AMOUNT_POSITION)))
        }
    }
}
