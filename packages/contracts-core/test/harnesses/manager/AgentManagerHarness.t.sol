// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManager} from "../../../contracts/manager/AgentManager.sol";

abstract contract AgentManagerHarness is AgentManager {
    /// @notice Exposes _openDispute for testing.
    function openDisputeExposed(address guard, address notary) external {
        _openDispute(guard, _storedAgentStatus(guard).index, notary, _storedAgentStatus(notary).index);
    }

    /// @notice Exposes _slashAgent for testing.
    function slashAgentExposed(uint32 domain, address agent, address prover) external {
        _slashAgent(domain, agent, prover);
    }
}
