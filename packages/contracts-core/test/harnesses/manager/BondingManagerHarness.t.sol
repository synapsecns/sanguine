// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BondingManager} from "../../../contracts/manager/BondingManager.sol";

import {AgentManagerHarness} from "./AgentManagerHarness.t.sol";

// solhint-disable no-empty-blocks
contract BondingManagerHarness is BondingManager, AgentManagerHarness {
    constructor(uint32 synapseDomain) BondingManager(synapseDomain) {}
}
