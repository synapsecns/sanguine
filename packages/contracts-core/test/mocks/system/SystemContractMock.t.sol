// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceSystemRouter, ISystemContract} from "../../../contracts/interfaces/ISystemContract.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

// solhint-disable no-empty-blocks
contract SystemContractMock is Ownable, ISystemContract {
    InterfaceSystemRouter public systemRouter;

    modifier onlySystemRouter() {
        require(msg.sender == address(systemRouter), "!systemRouter");
        _;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSystemContractMock() external {}

    function setSystemRouter(InterfaceSystemRouter systemRouter_) external {
        systemRouter = systemRouter_;
    }
}
