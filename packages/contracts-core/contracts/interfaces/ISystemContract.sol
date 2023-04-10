// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceSystemRouter} from "./InterfaceSystemRouter.sol";

interface ISystemContract {
    /**
     * @notice Sets System Router address in for a contract.
     * @dev This function should be protected. System Router is granted the ability
     * to pass the cross-chain system messages to the contract.
     */
    function setSystemRouter(InterfaceSystemRouter systemRouter_) external;

    /**
     * @notice Returns current System Router.
     */
    function systemRouter() external view returns (InterfaceSystemRouter);
}
