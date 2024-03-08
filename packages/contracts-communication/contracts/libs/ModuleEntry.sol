// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntry} from "./InterchainEntry.sol";

library ModuleEntryLib {
    /// @notice Encodes the InterchainEntry and the auxiliary module data into a single bytes array
    /// @param entry       The InterchainEntry to encode
    /// @param moduleData  The auxiliary module data to encode
    function encodeModuleEntry(
        InterchainEntry memory entry,
        bytes memory moduleData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encode(entry, moduleData);
    }

    /// @notice Decodes the bytes array into the InterchainEntry and the auxiliary module data
    /// @param encodedModuleEntry  The bytes array to decode
    /// @return entry              The decoded InterchainEntry
    /// @return moduleData         The decoded auxiliary module data
    function decodeModuleEntry(bytes memory encodedModuleEntry)
        internal
        pure
        returns (InterchainEntry memory entry, bytes memory moduleData)
    {
        return abi.decode(encodedModuleEntry, (InterchainEntry, bytes));
    }
}
