// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Origin } from "../../contracts/Origin.sol";

// TODO: remove/adapt when "go generate" is updated
contract OriginHarness is Origin {
    constructor(uint32 _domain) Origin(_domain) {}
}
