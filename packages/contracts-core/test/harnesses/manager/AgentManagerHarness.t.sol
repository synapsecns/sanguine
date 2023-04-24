// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManager} from "../../../contracts/manager/AgentManager.sol";

abstract contract AgentManagerHarness is AgentManager {
    /// @notice Exposes _slashAgent for testing.
    function slashAgentExposed(uint32 domain, address agent, address prover) external {
        _slashAgent(domain, agent, prover);
    }
}
