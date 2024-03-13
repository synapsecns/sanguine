pragma solidity =0.8.20 ^0.8.0 ^0.8.13;

// contracts/libs/Options.sol

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
    uint8 constant OPTIONS_V1 = 1;

    error OptionsLib__IncorrectVersion(uint8 version);

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
}

// contracts/libs/TypeCasts.sol

library TypeCasts {
    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    function bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }
}

// contracts/mocks/OptionsLibExport.sol

contract OptionsLibMocks {
    function encodeOptions(OptionsV1 memory options) public view returns (bytes memory) {
        return OptionsLib.encodeOptionsV1(options);
    }

    function decodeOptions(bytes memory data) public view returns (OptionsV1 memory) {
        return OptionsLib.decodeOptionsV1(data);
    }

    function addressToBytes32(address convertable) public view returns (bytes32) {
        return TypeCasts.addressToBytes32(convertable);
    }
}
