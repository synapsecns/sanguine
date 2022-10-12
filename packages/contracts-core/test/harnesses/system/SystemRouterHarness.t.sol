// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SystemRouter } from "../../../contracts/system/SystemRouter.sol";

// solhint-disable no-empty-blocks
contract SystemRouterHarness is SystemRouter {
    constructor(
        uint32 _localDomain,
        address _origin,
        address _destination
    ) SystemRouter(_localDomain, _origin, _destination) {}
}
