// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

interface IUpdaterManager {
    function slashUpdater(address payable _reporter) external;

    function updater() external view returns (address);
}
