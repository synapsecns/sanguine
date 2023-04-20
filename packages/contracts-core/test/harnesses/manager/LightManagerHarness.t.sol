// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {LightManager} from "../../../contracts/manager/LightManager.sol";

// solhint-disable no-empty-blocks
contract LightManagerHarness is LightManager {
    constructor(uint32 domain) LightManager(domain) {}

    function remoteMockFunc(uint32, uint256, bytes32) external view returns (bytes4) {
        require(msg.sender == destination, "!destination");
        return this.remoteMockFunc.selector;
    }
}
