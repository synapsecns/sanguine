// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Origin } from "../../contracts/Origin.sol";
import { IAgentManager } from "../../contracts/interfaces/IAgentManager.sol";

import { SystemRouterMock } from "../mocks/system/SystemRouterMock.t.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract OriginHarness is Origin {
    constructor(uint32 _domain, address _agentManager)
        Origin(_domain, IAgentManager(_agentManager))
    {
        // Add Mock for SystemRouter for standalone tests
        systemRouter = new SystemRouterMock();
    }

    // TODO: add / remove Agents in standalone Go tests
}
