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
     * @notice Receive a system call indicating the off-chain agent status needs to be updated.
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _info             Information about agent to sync
     */
    function syncAgent(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external;

    function setSystemRouter(InterfaceSystemRouter _systemRouter) external;

    function systemRouter() external view returns (InterfaceSystemRouter);
}
