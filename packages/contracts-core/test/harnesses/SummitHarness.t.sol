// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Summit} from "../../contracts/Summit.sol";
import "../utils/SynapseTestConstants.t.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract SummitHarness is Summit, SynapseTestConstants {
    /// @dev Summit could only be deployed on Synapse Domain
    // solhint-disable-next-line no-empty-blocks
    constructor(address agentManager_, address inbox_) Summit(DOMAIN_SYNAPSE, agentManager_, inbox_) {}

    // TODO: add / remove Agents in standalone Go tests
}
