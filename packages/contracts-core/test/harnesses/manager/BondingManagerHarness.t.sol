// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BondingManager} from "../../../contracts/manager/BondingManager.sol";
import {SystemContractHarness} from "../system/SystemContractHarness.t.sol";

// solhint-disable no-empty-blocks
contract BondingManagerHarness is BondingManager, SystemContractHarness {
    constructor(uint32 domain) BondingManager(domain) {}
}
