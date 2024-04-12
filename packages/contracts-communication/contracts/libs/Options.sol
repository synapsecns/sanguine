// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {VersionedPayloadLib} from "./VersionedPayload.sol";

/// @notice Struct to hold V1 of options data.
/// @dev Next versions have to use the fields from the previous version and add new fields at the end.
/// @param gasLimit The gas limit for the transaction.
/// @param gasAirdrop The amount of gas to airdrop.
struct OptionsV1 {
    uint256 gasLimit;
    uint256 gasAirdrop;
}

using OptionsLib for OptionsV1 global;

/// @title OptionsLib
/// @notice A library for encoding and decoding Interchain options related to interchain messages.
library OptionsLib {
    using VersionedPayloadLib for bytes;

    uint16 internal constant OPTIONS_V1 = 1;

    error OptionsLib__IncorrectVersion(uint16 version);

    /// @notice Decodes options (V1 or higher) from a bytes format back into an OptionsV1 struct.
    /// @param data         The options data in bytes format.
    function decodeOptionsV1(bytes memory data) internal view returns (OptionsV1 memory) {
        uint16 version = data.getVersionFromMemory();
        if (version < OPTIONS_V1) {
            revert OptionsLib__IncorrectVersion(version);
        }
        // Structs of the same version will always be decoded correctly.
        // Following versions will be decoded correctly if they have the same fields as the previous version,
        // and new fields at the end: abi.decode ignores the extra bytes in the decoded payload.
        return abi.decode(data.getPayloadFromMemory(), (OptionsV1));
    }

    /// @notice Encodes V1 options into a bytes format.
    /// @param options      The OptionsV1 to encode.
    function encodeOptionsV1(OptionsV1 memory options) internal pure returns (bytes memory) {
        return VersionedPayloadLib.encodeVersionedPayload(OPTIONS_V1, abi.encode(options));
    }
}
