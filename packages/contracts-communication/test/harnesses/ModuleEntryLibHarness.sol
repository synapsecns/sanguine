// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, ModuleEntryLib} from "../../contracts/libs/ModuleEntry.sol";

contract ModuleEntryLibHarness {
    function encodeModuleEntry(
        InterchainEntry memory entry,
        bytes memory moduleData
    )
        external
        pure
        returns (bytes memory)
    {
        return ModuleEntryLib.encodeModuleEntry(entry, moduleData);
    }

    function decodeModuleEntry(bytes memory encodedModuleEntry)
        external
        pure
        returns (InterchainEntry memory, bytes memory)
    {
        return ModuleEntryLib.decodeModuleEntry(encodedModuleEntry);
    }
}
