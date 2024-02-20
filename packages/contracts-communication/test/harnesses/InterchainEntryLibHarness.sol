// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, InterchainEntryLib} from "../../contracts/libs/InterchainEntry.sol";

contract InterchainEntryLibHarness {
    function constructLocalEntry(
        uint256 dbNonce,
        address srcWriter,
        bytes32 dataHash
    )
        external
        view
        returns (InterchainEntry memory)
    {
        return InterchainEntryLib.constructLocalEntry(dbNonce, srcWriter, dataHash);
    }

    function entryKey(InterchainEntry memory entry) external pure returns (bytes32) {
        return InterchainEntryLib.entryKey(entry);
    }

    function entryValue(InterchainEntry memory entry) external pure returns (bytes32) {
        return InterchainEntryLib.entryValue(entry);
    }
}
