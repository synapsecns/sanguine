// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntry} from "../libs/InterchainEntry.sol";

interface ISynapseModuleEvents {
    event VerificationRequested(uint256 indexed destChainId, InterchainEntry entry, bytes32 signedEntryHash);
    event EntryVerified(InterchainEntry entry);
}
