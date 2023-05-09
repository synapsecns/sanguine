// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Destination} from "../../contracts/Destination.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract DestinationHarness is Destination {
    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 domain, address agentManager_, address inbox_) Destination(domain, agentManager_, inbox_) {}

    // TODO: add / remove Agents in standalone Go tests
}
