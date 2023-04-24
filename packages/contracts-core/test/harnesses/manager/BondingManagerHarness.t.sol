// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BondingManager} from "../../../contracts/manager/BondingManager.sol";

import {AgentManagerHarness} from "./AgentManagerHarness.t.sol";

// solhint-disable no-empty-blocks
contract BondingManagerHarness is BondingManager, AgentManagerHarness {
    constructor(uint32 domain) BondingManager(domain) {}

    function remoteMockFunc(uint32, uint256, bytes32) external view returns (bytes4) {
        require(msg.sender == destination, "!destination");
        return this.remoteMockFunc.selector;
    }
}
