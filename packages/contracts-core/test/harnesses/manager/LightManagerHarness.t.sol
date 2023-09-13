// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {LightManager} from "../../../contracts/manager/LightManager.sol";

import {AgentManager, AgentManagerHarness} from "./AgentManagerHarness.t.sol";

// solhint-disable no-empty-blocks
contract LightManagerHarness is LightManager, AgentManagerHarness {
    constructor(uint32 synapseDomain) LightManager(synapseDomain) {}

    function _afterAgentSlashed(uint32 domain, address agent, address prover)
        internal
        override(AgentManager, LightManager)
    {
        LightManager._afterAgentSlashed(domain, agent, prover);
    }
}
