// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Origin } from "../../contracts/Origin.sol";
import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";
import { SystemContractHarness } from "./SystemContractHarness.sol";

contract OriginHarness is Origin, SystemContractHarness, GuardRegistryHarness {
    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) Origin(_domain) {}
}
