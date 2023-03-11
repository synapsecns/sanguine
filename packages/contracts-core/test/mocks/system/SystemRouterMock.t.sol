// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    InterfaceSystemRouter,
    SystemEntity
} from "../../../contracts/interfaces/InterfaceSystemRouter.sol";

contract SystemRouterMock is InterfaceSystemRouter {
    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRouterMock() external {}

    function systemCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity _recipient,
        bytes memory _data
    ) external {}

    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes[] memory _dataArray
    ) external {}

    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes memory _data
    ) external {}

    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity _recipient,
        bytes[] memory _dataArray
    ) external {}
}
