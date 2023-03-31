// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Summit } from "../../contracts/Summit.sol";
import { IAgentManager } from "../../contracts/interfaces/IAgentManager.sol";

import { SystemRouterMock } from "../mocks/system/SystemRouterMock.t.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract SummitHarness is Summit {
    /// @dev Summit could only be deployed on Synapse Domain
    constructor(address agentManager_) Summit(SYNAPSE_DOMAIN, IAgentManager(agentManager_)) {
        // Add Mock for SystemRouter for standalone tests
        systemRouter = new SystemRouterMock();
    }

    // TODO: add / remove Agents in standalone Go tests
}
