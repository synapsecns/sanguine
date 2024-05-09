// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library ModuleEntryLib {
    /// @notice Encodes the versioned entry and the auxiliary module data into a single bytes array
    /// @param versionedEntry       The versioned entry to encode
    /// @param moduleData           The auxiliary module data to encode
    /// @return encodedModuleEntry  The encoded versioned module entry
    function encodeVersionedModuleEntry(
        bytes memory versionedEntry,
        bytes memory moduleData
    )
        internal
        pure
        returns (bytes memory encodedModuleEntry)
    {
        return abi.encode(versionedEntry, moduleData);
    }

    /// @notice Decodes the bytes array into the versioned entry and the auxiliary module data
    /// @param encodedModuleEntry   The bytes array to decode
    /// @return versionedEntry      The decoded versioned entry
    /// @return moduleData          The decoded auxiliary module data
    function decodeVersionedModuleEntry(bytes memory encodedModuleEntry)
        internal
        pure
        returns (bytes memory versionedEntry, bytes memory moduleData)
    {
        return abi.decode(encodedModuleEntry, (bytes, bytes));
    }
}
