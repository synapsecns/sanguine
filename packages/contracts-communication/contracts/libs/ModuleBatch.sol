// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainBatch} from "./InterchainBatch.sol";

library ModuleBatchLib {
    /// @notice Encodes the InterchainBatch and the auxiliary module data into a single bytes array
    /// @param batch       The InterchainBatch to encode
    /// @param moduleData  The auxiliary module data to encode
    function encodeModuleBatch(
        InterchainBatch memory batch,
        bytes memory moduleData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encode(batch, moduleData);
    }

    /// @notice Decodes the bytes array into the InterchainBatch and the auxiliary module data
    /// @param encodedModuleBatch  The bytes array to decode
    /// @return batch              The decoded InterchainBatch
    /// @return moduleData         The decoded auxiliary module data
    function decodeModuleBatch(bytes memory encodedModuleBatch)
        internal
        pure
        returns (InterchainBatch memory batch, bytes memory moduleData)
    {
        return abi.decode(encodedModuleBatch, (InterchainBatch, bytes));
    }

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
