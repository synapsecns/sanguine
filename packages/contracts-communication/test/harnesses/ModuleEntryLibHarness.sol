// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ModuleEntryLib} from "../../contracts/libs/ModuleEntry.sol";

contract ModuleEntryLibHarness {
    function encodeVersionedModuleEntry(
        bytes memory versionedEntry,
        bytes memory moduleData
    )
        external
        pure
        returns (bytes memory)
    {
        return ModuleEntryLib.encodeVersionedModuleEntry(versionedEntry, moduleData);
    }

    function decodeVersionedModuleEntry(bytes memory encodedModuleEntry)
        external
        pure
        returns (bytes memory, bytes memory)
    {
        return ModuleEntryLib.decodeVersionedModuleEntry(encodedModuleEntry);
    }
}
