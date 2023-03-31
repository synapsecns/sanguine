// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { InterfaceSystemRouter } from "./InterfaceSystemRouter.sol";

interface ISystemContract {
    function setSystemRouter(InterfaceSystemRouter systemRouter_) external;

    function systemRouter() external view returns (InterfaceSystemRouter);
}
