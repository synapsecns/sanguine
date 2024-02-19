// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Struct to hold V1 of options data.
/// @dev Next versions have to use the fields from the previous version and add new fields at the end.
/// @param gasLimit The gas limit for the transaction.
/// @param gasAirdrop The amount of gas to airdrop.
struct OptionsV1 {
    uint256 gasLimit;
    uint256 gasAirdrop;
}

/// @title OptionsLib
/// @notice A library for encoding and decoding Interchain options related to interchain messages.
library OptionsLib {
    error OptionsLib__IncorrectVersion(uint8 version);

    uint8 constant OPTIONS_V1 = 1;

    /// @dev Struct to hold V1 of options data.
    /// @param version The version of the options.
    /// @param gasLimit The gas limit for the transaction.
    /// @param gasAirdrop The amount of gas to airdrop.
    struct Options {
        uint8 version;
        uint256 gasLimit;
        // uint256 msgValue;
        uint256 gasAirdrop;
    }

    /// @notice Encodes versioned options into a bytes format.
    /// @param version      The version of the options.
    /// @param options      The options to encode.
    function encodeVersionedOptions(uint8 version, bytes memory options) internal pure returns (bytes memory) {
        return abi.encode(version, options);
    }

    /// @notice Decodes versioned options from a bytes format back into a version and options.
    /// @param data         The versioned options data in bytes format.
    /// @return version     The version of the options.
    /// @return options     The options as bytes.
    function decodeVersionedOptions(bytes memory data) internal pure returns (uint8 version, bytes memory options) {
        (version, options) = abi.decode(data, (uint8, bytes));
    }

    /// @notice Encodes V1 options into a bytes format.
    /// @param options      The OptionsV1 to encode.
    function encodeOptionsV1(OptionsV1 memory options) internal pure returns (bytes memory) {
        return encodeVersionedOptions(OPTIONS_V1, abi.encode(options));
    }

    /// @notice Decodes options (V1 or higher) from a bytes format back into an OptionsV1 struct.
    /// @param data         The options data in bytes format.
    function decodeOptionsV1(bytes memory data) internal pure returns (OptionsV1 memory) {
        (uint8 version, bytes memory options) = decodeVersionedOptions(data);
        if (version < OPTIONS_V1) {
            revert OptionsLib__IncorrectVersion(version);
        }
        // Structs of the same version will always be decoded correctly.
        // Following versions will be decoded correctly if they have the same fields as the previous version,
        // and new fields at the end: abi.decode ignores the extra bytes in the decoded payload.
        return abi.decode(options, (OptionsV1));
    }

    /// @notice Encodes options into a bytes format.
    /// @param options The Options to encode.
    /// @return The encoded options as bytes.
    function encodeOptions(Options memory options) internal pure returns (bytes memory) {
        return abi.encode(options.version, options.gasLimit, options.gasAirdrop);
    }

    /// @notice Decodes options from a bytes format back into an Options struct.
    /// @param data The options data in bytes format.
    /// @return The decoded options as an Options struct.
    function decodeOptions(bytes memory data) internal pure returns (Options memory) {
        (uint8 version, uint256 gasLimit, uint256 gasAirdrop) = abi.decode(data, (uint8, uint256, uint256));
        return Options(version, gasLimit, gasAirdrop);
    }
}
