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
}
