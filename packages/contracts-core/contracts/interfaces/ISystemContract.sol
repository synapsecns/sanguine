// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { InterfaceSystemRouter } from "./InterfaceSystemRouter.sol";

interface ISystemContract {
    function setSystemRouter(InterfaceSystemRouter _systemRouter) external;

    function systemRouter() external view returns (InterfaceSystemRouter);
}
