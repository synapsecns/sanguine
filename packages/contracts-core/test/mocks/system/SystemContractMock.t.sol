// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceSystemRouter, ISystemContract} from "../../../contracts/interfaces/ISystemContract.sol";
import {SystemEntity} from "./SystemRouterMock.t.sol";

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

    function remoteMockFunc(uint256 proofMaturity, uint32 origin, SystemEntity sender, bytes32 data)
        external
        onlySystemRouter
    {}
}
