// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Summit, SYNAPSE_DOMAIN} from "../../contracts/Summit.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract SummitHarness is Summit {
    /// @dev Summit could only be deployed on Synapse Domain
    // solhint-disable-next-line no-empty-blocks
    constructor(address agentManager_, address inbox_) Summit(SYNAPSE_DOMAIN, agentManager_, inbox_) {}

    // TODO: add / remove Agents in standalone Go tests
}
