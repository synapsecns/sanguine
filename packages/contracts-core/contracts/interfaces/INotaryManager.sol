// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface INotaryManager {
    function slashNotary(address payable _reporter) external;

    function notary() external view returns (address);
}
