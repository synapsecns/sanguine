// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library LegacyOptionsLib {
    uint16 internal constant LEGACY_OPTIONS_VERSION = 1;

    error LegacyOptionsLib__InvalidOptions(bytes legacyOpts);

    /// @notice Encodes the gas limit into a legacy options format.
    /// @param gasLimit     The gas limit to encode.
    /// @return legacyOpts  The encoded legacy options
    function encodeLegacyOptions(uint256 gasLimit) internal pure returns (bytes memory legacyOpts) {
        // TODO: Implement
    }

    /// @notice Decodes the gas limit from a legacy options format.
    /// @param legacyOpts   The encoded legacy options
    /// @return gasLimit    The gas limit
    function decodeLegacyOptions(bytes calldata legacyOpts) internal pure returns (uint256 gasLimit) {
        // TODO: Implement
    }
}
