// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

library ZapDataV1 {
    /// @notice Version of the Zap Data struct.
    uint16 internal constant VERSION = 1;

    /// @notice Value that indicates the amount is not present in the target function's payload.
    uint16 internal constant AMOUNT_NOT_PRESENT = 0xFFFF;

    // Offsets of the fields in the packed ZapData struct
    // uint16   version                 [000 .. 002)
    // uint16   amountPosition          [002 .. 004)
    // address  target                  [004 .. 024)
    // bytes    payload                 [024 .. ***)

    // forgefmt: disable-start
    uint256 private constant OFFSET_AMOUNT_POSITION = 2;
    uint256 private constant OFFSET_TARGET          = 4;
    uint256 private constant OFFSET_PAYLOAD         = 24;
    // forgefmt: disable-end

    error ZapDataV1__InvalidEncoding();
    error ZapDataV1__UnsupportedVersion(uint16 version);

    /// @notice Validates the encodedZapData to be a tightly packed encoded payload for ZapData struct.
    /// @dev Checks that all the required fields are present, version is correct and amount position is valid.
    function validateV1(bytes calldata encodedZapData) internal pure {
        // TODO: implement
    }

    /// @notice Encodes the ZapData struct by tightly packing the fields.
    /// Note: we don't know the exact amount that will be used for the Zap at the time of encoding,
    /// so we provide the reference index where the amount is encoded within `payload_`. This allows up to
    /// hot-swap the amount in the payload, when the Zap is performed.
    /// @dev `abi.decode` will not work as a result of the tightly packed fields. Use `decodeZapData` instead.
    /// @param amountPosition_  Position (start index) where the amount is encoded within `payload_`.
    ///                         This will usually be `4 + 32 * n`, where `n` is the position of the amount in
    ///                         the list of parameters of the target function (starting from 0).
    ///                         Or `AMOUNT_NOT_PRESENT` if the amount is not encoded within `payload_`.
    /// @param target_          Address of the target contract.
    /// @param payload_         Payload to be used as a calldata for the `target_` contract call.
    function encodeV1(
        uint16 amountPosition_,
        address target_,
        bytes memory payload_
    )
        internal
        pure
        returns (bytes memory encodedZapData)
    {
        // TODO: implement
    }

    /// @notice Extracts the version from the encoded Zap Data.
    function version(bytes calldata encodedZapData) internal pure returns (uint16) {
        // TODO: implement
    }

    /// @notice Extracts the target address from the encoded Zap Data.
    function target(bytes calldata encodedZapData) internal pure returns (address) {
        // TODO: implement
    }

    /// @notice Extracts the payload from the encoded Zap Data. Replaces the amount with the provided value,
    /// if it was present in the original data (amountPosition is not AMOUNT_NOT_PRESENT).
    /// @dev This payload will be used as a calldata for the target contract.
    function payload(bytes calldata encodedZapData, uint256 amount) internal pure returns (bytes memory) {
        // TODO: implement
    }
}
