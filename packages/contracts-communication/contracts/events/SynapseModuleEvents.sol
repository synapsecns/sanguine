// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntry} from "../libs/InterchainEntry.sol";

abstract contract SynapseModuleEvents {
    event VerificationRequested(uint256 indexed destChainId, InterchainEntry entry, bytes32 signableEntryHash);
    event EntryVerified(InterchainEntry entry, bytes32 signableEntryHash);
}
