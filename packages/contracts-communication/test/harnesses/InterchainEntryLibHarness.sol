// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, InterchainEntryLib} from "../../contracts/libs/InterchainEntry.sol";

contract InterchainEntryLibHarness {
    function constructLocalEntry(
        address srcWriter,
        uint256 writerNonce,
        bytes32 dataHash
    )
        external
        view
        returns (InterchainEntry memory)
    {
        return InterchainEntryLib.constructLocalEntry(srcWriter, writerNonce, dataHash);
    }

    function entryId(InterchainEntry memory entry) external pure returns (bytes32) {
        return InterchainEntryLib.entryId(entry);
    }
}
