// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    InterfaceSystemRouter,
    SystemEntity
} from "../../../contracts/interfaces/InterfaceSystemRouter.sol";

// solhint-disable no-empty-blocks
contract SystemRouterMock is InterfaceSystemRouter {
    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRouterMock() external {}

    function systemCall(
        uint32 destination,
        uint32 optimisticSeconds,
        SystemEntity recipient,
        bytes memory data
    ) external {}

    function systemMultiCall(
        uint32 destination,
        uint32 optimisticSeconds,
        SystemEntity[] memory recipients,
        bytes[] memory dataArray
    ) external {}

    function systemMultiCall(
        uint32 destination,
        uint32 optimisticSeconds,
        SystemEntity[] memory recipients,
        bytes memory data
    ) external {}

    function systemMultiCall(
        uint32 destination,
        uint32 optimisticSeconds,
        SystemEntity recipient,
        bytes[] memory dataArray
    ) external {}
}
