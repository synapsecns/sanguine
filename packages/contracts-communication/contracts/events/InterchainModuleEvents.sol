// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntry} from "../libs/InterchainEntry.sol";

abstract contract InterchainModuleEvents {
    event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash);

    event EntryVerified(InterchainEntry entry);
}
