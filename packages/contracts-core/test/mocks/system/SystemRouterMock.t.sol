// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceSystemRouter, SystemEntity} from "../../../contracts/interfaces/InterfaceSystemRouter.sol";

// solhint-disable no-empty-blocks
contract SystemRouterMock is InterfaceSystemRouter {
    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRouterMock() external {}

    function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 rootSubmittedAt, bytes memory body) external {}

    function systemCall(uint32 destination, uint32 optimisticSeconds, SystemEntity recipient, bytes memory data)
        external
    {}
}
