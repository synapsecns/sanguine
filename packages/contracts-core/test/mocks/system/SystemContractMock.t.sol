// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../../contracts/interfaces/ISystemContract.sol";
import "../ExcludeCoverage.sol";

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

// solhint-disable no-empty-blocks
contract SystemContractMock is ExcludeCoverage, Ownable, ISystemContract {
    InterfaceSystemRouter public systemRouter;

    modifier onlySystemRouter() {
        require(msg.sender == address(systemRouter), "!systemRouter");
        _;
    }

    function setSystemRouter(InterfaceSystemRouter _systemRouter) external {
        systemRouter = _systemRouter;
    }

    function slashAgent(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external onlySystemRouter {}

    function syncAgents(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        uint256 _requestID,
        bool _removeExisting,
        AgentInfo[] memory _infos
    ) external onlySystemRouter {}
}
