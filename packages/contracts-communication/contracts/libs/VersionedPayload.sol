// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// solhint-disable no-inline-assembly
// solhint-disable ordering
library VersionedPayloadLib {
    /// @notice Amount of bytes reserved for the version (uint16) in the versioned payload
    uint256 internal constant VERSION_LENGTH = 2;

    error VersionedPayload__PayloadTooShort(bytes versionedPayload);
    error VersionedPayload__PrecompileFailed();

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
            revert VersionedPayload__PayloadTooShort(versionedPayload);
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
            revert VersionedPayload__PayloadTooShort(versionedPayload);
        }
        return versionedPayload[VERSION_LENGTH:];
    }

    /// @notice Extracts the version from the versioned payload (memory reference).
    /// @param versionedPayload     The versioned payload (memory reference).
    function getVersionFromMemory(bytes memory versionedPayload) internal pure returns (uint16 version) {
        if (versionedPayload.length < VERSION_LENGTH) {
            revert VersionedPayload__PayloadTooShort(versionedPayload);
        }
        assembly {
            // We are only interested in the highest 16 bits of the loaded full 32 bytes word.
            // We add 0x20 to skip the length of the bytes array.
            version := shr(240, mload(add(versionedPayload, 0x20)))
        }
    }

    /// @notice Extracts the payload from the versioned payload (memory reference).
    /// @dev The extracted payload is copied into a new memory location. Use `getPayload` when possible
    /// to avoid extra memory allocation.
    /// @param versionedPayload     The versioned payload (memory reference).
    function getPayloadFromMemory(bytes memory versionedPayload) internal view returns (bytes memory payload) {
        if (versionedPayload.length < VERSION_LENGTH) {
            revert VersionedPayload__PayloadTooShort(versionedPayload);
        }
        // Figure how many bytes to copy and allocate the memory for the extracted payload.
        uint256 toCopy;
        unchecked {
            toCopy = versionedPayload.length - VERSION_LENGTH;
        }
        payload = new bytes(toCopy);
        // Use identity precompile (0x04) to copy the payload. Unlike MCOPY, this is available on all EVM chains.
        bool res;
        assembly {
            // We add 0x20 to skip the length of the bytes array.
            // We add 0x02 to skip the 2 bytes reserved for the version.
            // Copy the payload to the previously allocated memory.
            res := staticcall(gas(), 0x04, add(versionedPayload, 0x22), toCopy, add(payload, 0x20), toCopy)
        }
        if (!res) {
            revert VersionedPayload__PrecompileFailed();
        }
    }
}
