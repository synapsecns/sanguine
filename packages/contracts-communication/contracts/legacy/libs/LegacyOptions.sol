// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library LegacyOptionsLib {
    /// @notice Supported version of the legacy options format.
    uint16 internal constant LEGACY_OPTIONS_VERSION = 1;
    /// @dev The length of the legacy options format: 2 bytes for the version and 32 bytes for the gas limit.
    uint256 private constant LEGACY_OPTIONS_LENGTH = 34;
    /// @dev The offset of the gas limit in the legacy options format.
    uint256 private constant GAS_LIMIT_OFFSET = 2;

    error LegacyOptionsLib__PayloadInvalid(bytes legacyOpts);

    /// @notice Encodes the gas limit into a legacy options format.
    /// @param gasLimit     The gas limit to encode.
    /// @return legacyOpts  The encoded legacy options
    function encodeLegacyOptions(uint256 gasLimit) internal pure returns (bytes memory legacyOpts) {
        return abi.encodePacked(LEGACY_OPTIONS_VERSION, gasLimit);
    }

    /// @notice Decodes the gas limit from a legacy options format.
    /// @param legacyOpts   The encoded legacy options
    /// @return gasLimit    The gas limit
    function decodeLegacyOptions(bytes calldata legacyOpts) internal pure returns (uint256 gasLimit) {
        if (legacyOpts.length != LEGACY_OPTIONS_LENGTH) {
            revert LegacyOptionsLib__PayloadInvalid(legacyOpts);
        }
        uint16 version = uint16(bytes2(legacyOpts[:GAS_LIMIT_OFFSET]));
        if (version != LEGACY_OPTIONS_VERSION) {
            revert LegacyOptionsLib__PayloadInvalid(legacyOpts);
        }
        gasLimit = uint256(bytes32(legacyOpts[GAS_LIMIT_OFFSET:]));
    }
}
