// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SystemRouter } from "../../contracts/system/SystemRouter.sol";

contract SystemRouterHarness is SystemRouter {
    constructor(
        address _origin,
        address _destination,
        uint32 _optimisticSeconds
    ) SystemRouter(_origin, _destination, _optimisticSeconds) {}
}
