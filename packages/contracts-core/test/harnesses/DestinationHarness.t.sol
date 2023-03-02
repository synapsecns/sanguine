// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Destination } from "../../contracts/Destination.sol";

// TODO: remove/adapt when "go generate" is updated
contract DestinationHarness is Destination {
    constructor(uint32 _domain) Destination(_domain) {}
}
