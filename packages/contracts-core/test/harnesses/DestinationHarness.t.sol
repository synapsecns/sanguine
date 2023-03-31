// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Destination} from "../../contracts/Destination.sol";
import {IAgentManager} from "../../contracts/interfaces/IAgentManager.sol";

import {SystemRouterMock} from "../mocks/system/SystemRouterMock.t.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract DestinationHarness is Destination {
    constructor(uint32 domain, address agentManager_) Destination(domain, IAgentManager(agentManager_)) {
        // Add Mock for SystemRouter for standalone tests
        systemRouter = new SystemRouterMock();
    }

    // TODO: add / remove Agents in standalone Go tests
}
