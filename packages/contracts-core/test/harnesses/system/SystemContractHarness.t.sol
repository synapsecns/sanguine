// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SystemContract, SystemEntity} from "../../../contracts/system/SystemContract.sol";

// solhint-disable no-empty-blocks
abstract contract SystemContractHarness is SystemContract {
    function remoteMockFunc(uint256 rootSubmittedAt, uint32 origin, SystemEntity sender, bytes32 data)
        external
        onlySystemRouter
    {}
}
