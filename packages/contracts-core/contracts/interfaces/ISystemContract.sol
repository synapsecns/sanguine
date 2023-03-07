// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../libs/Structures.sol";
import { InterfaceSystemRouter } from "./InterfaceSystemRouter.sol";

interface ISystemContract {
    /**
     * @notice Receive a system call indicating the off-chain agent needs to be slashed.
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _info             Information about agent to slash
     */
    function slashAgent(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external;

    /**
     * @notice Receive a system call indicating the list of off-chain agents needs to be synced.
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _requestID        Unique ID of the sync request
     * @param _removeExisting   Whether the existing agents need to be removed first
     * @param _infos            Information about a list of agents to sync
     */
    function syncAgents(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        uint256 _requestID,
        bool _removeExisting,
        AgentInfo[] memory _infos
    ) external;

    function setSystemRouter(InterfaceSystemRouter _systemRouter) external;

    function systemRouter() external view returns (InterfaceSystemRouter);
}
