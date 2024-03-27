// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// solhint-disable no-inline-assembly
library VersionedPayloadLib {
    /// @notice Amount of bytes reserved for the version (uint16) in the versioned payload
    uint256 internal constant VERSION_LENGTH = 2;

    error VersionedPayload__TooShort(bytes versionedPayload);

    /// @notice Encodes the versioned payload into a single bytes array.
    /// @param version  The payload's version.
    /// @param payload  The payload to encode.
    function encodeVersionedPayload(uint16 version, bytes memory payload) internal pure returns (bytes memory) {
        return abi.encodePacked(version, payload);
    }

    /// @notice Extracts the version from the versioned payload (calldata reference).
    /// @param versionedPayload     The versioned payload (calldata reference).
    function getVersion(bytes calldata versionedPayload) internal pure returns (uint16 version) {
        if (versionedPayload.length < VERSION_LENGTH) {
            revert VersionedPayload__TooShort(versionedPayload);
        }
        assembly {
            // We are only interested in the highest 16 bits of the loaded full 32 bytes word.
            version := shr(240, calldataload(versionedPayload.offset))
        }
    }

    /// @notice Extracts the payload from the versioned payload (calldata reference).
    /// @dev The extracted payload is also returned as a calldata reference.
    /// @param versionedPayload     The versioned payload.
    function getPayload(bytes calldata versionedPayload) internal pure returns (bytes calldata) {
        if (versionedPayload.length < VERSION_LENGTH) {
            revert VersionedPayload__TooShort(versionedPayload);
        }
        return versionedPayload[VERSION_LENGTH:];
    }

    /// @notice Extracts the version from the versioned payload (memory reference).
    /// @param versionedPayload     The versioned payload (memory reference).
    function getVersionFromMemory(bytes memory versionedPayload) internal pure returns (uint16) {
        // TODO: implement
    }

    /// @notice Extracts the payload from the versioned payload (memory reference).
    /// @dev The extracted payload is copied into a new memory location. Use `getPayload` when possible
    /// to avoid extra memory allocation.
    /// @param versionedPayload     The versioned payload (memory reference).
    function getPayloadFromMemory(bytes memory versionedPayload) internal pure returns (bytes memory) {
        // TODO: implement
    }
}
