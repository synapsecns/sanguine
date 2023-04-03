// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {LightManager} from "../../../contracts/manager/LightManager.sol";
import {SystemContractHarness} from "../system/SystemContractHarness.t.sol";

// solhint-disable no-empty-blocks
contract LightManagerHarness is LightManager, SystemContractHarness {
    constructor(uint32 domain) LightManager(domain) {}
}
