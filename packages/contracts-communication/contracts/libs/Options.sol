// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title OptionsLib
/// @notice A library for encoding and decoding Interchain options related to interchain messages.
library OptionsLib {
    struct NativeDrop {
        address recipient;
        uint256 amount;
    }

    struct OptionsV1 {
        uint8 version;
        uint256 gasLimit;
        uint256 msgValue;
        NativeDrop[] nativeDrops;
    }

    /// @dev Struct to hold V1 of options data.
    /// @param version The version of the options.
    /// @param gasLimit The gas limit for the transaction.
    /// @param gasAirdrop The amount of gas to airdrop.
    struct Options {
        uint8 version;
        uint256 gasLimit;
        // uint256 msgValue;
        NativeDrop[] nativeDrops;

    }

    /// @notice Encodes options into a bytes format.
    /// @param options The Options to encode.
    /// @return The encoded options as bytes.
    function encodeOptions(Options memory options) internal pure returns (bytes memory) {
        return abi.encode(options.version, options.gasLimit, options.nativeDrops);
    }

    /// @notice Decodes options from a bytes format back into an Options struct.
    /// @param data The options data in bytes format.
    /// @return The decoded options as an Options struct.
    function decodeOptions(bytes memory data) internal pure returns (Options memory) {
        (uint8 version, uint256 gasLimit, NativeDrop[] memory nativeDrops) = abi.decode(data, (uint8, uint256, NativeDrop[]));
        return Options(version, gasLimit, nativeDrops);
    }
}

