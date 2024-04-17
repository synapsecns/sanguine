// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainBatch} from "./InterchainBatch.sol";

library ModuleBatchLib {
    /// @notice Encodes the versioned batch and the auxiliary module data into a single bytes array
    /// @param versionedBatch       The versioned batch to encode
    /// @param moduleData           The auxiliary module data to encode
    /// @return encodedModuleBatch  The encoded versioned module batch
    function encodeVersionedModuleBatch(
        bytes memory versionedBatch,
        bytes memory moduleData
    )
        internal
        pure
        returns (bytes memory encodedModuleBatch)
    {
        return abi.encode(versionedBatch, moduleData);
    }

    /// @notice Decodes the bytes array into the versioned batch and the auxiliary module data
    /// @param encodedModuleBatch   The bytes array to decode
    /// @return versionedBatch      The decoded versioned batch
    /// @return moduleData          The decoded auxiliary module data
    function decodeVersionedModuleBatch(bytes memory encodedModuleBatch)
        internal
        pure
        returns (bytes memory versionedBatch, bytes memory moduleData)
    {
        return abi.decode(encodedModuleBatch, (bytes, bytes));
    }
}
