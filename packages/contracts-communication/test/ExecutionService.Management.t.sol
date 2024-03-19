// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionService, ExecutionServiceEvents} from "../contracts/ExecutionService.sol";

import {SynapseGasOracleMock} from "./mocks/SynapseGasOracleMock.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract ExecutionServiceManagementTest is ExecutionServiceEvents, Test {
    ExecutionService public service;
    SynapseGasOracleMock public gasOracle;

    address public icClient = makeAddr("InterchainClient");
    address public executorEOA = makeAddr("ExecutorEOA");
    address public owner = makeAddr("Owner");

    function setUp() public {
        gasOracle = new SynapseGasOracleMock();
        service = new ExecutionService(address(this));
        service.setInterchainClient(icClient);
        service.setExecutorEOA(executorEOA);
        service.setGasOracle(address(gasOracle));
        service.transferOwnership(owner);
    }

    function expectUnauthorizedAccountRevert(address caller) internal {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, caller));
    }

    function test_setInterchainClient_emitsEvent() public {
        vm.expectEmit(address(service));
        emit InterchainClientUpdated(address(1));
        vm.prank(owner);
        service.setInterchainClient(address(1));
    }

    function test_setInterchainClient_setsInterchainClient() public {
        vm.prank(owner);
        service.setInterchainClient(address(1));
        assertEq(service.interchainClient(), address(1));
    }

    function test_setInterchainClient_revert_callerNotOwner(address caller) public {
        vm.assume(caller != owner);
        expectUnauthorizedAccountRevert(caller);
        vm.prank(caller);
        service.setInterchainClient(address(1));
    }

    function test_setExecutorEOA_emitsEvent() public {
        vm.expectEmit(address(service));
        emit ExecutorEOAUpdated(address(1));
        vm.prank(owner);
        service.setExecutorEOA(address(1));
    }

    function test_setExecutorEOA_setsExecutorEOA() public {
        vm.prank(owner);
        service.setExecutorEOA(address(1));
        assertEq(service.executorEOA(), address(1));
    }

    function test_setExecutorEOA_revert_callerNotOwner(address caller) public {
        vm.assume(caller != owner);
        expectUnauthorizedAccountRevert(caller);
        vm.prank(caller);
        service.setExecutorEOA(address(1));
    }

    function test_setGasOracle_emitsEvent() public {
        vm.expectEmit(address(service));
        emit GasOracleUpdated(address(1));
        vm.prank(owner);
        service.setGasOracle(address(1));
    }

    function test_setGasOracle_setsGasOracle() public {
        vm.prank(owner);
        service.setGasOracle(address(1));
        assertEq(address(service.gasOracle()), address(1));
    }

    function test_setGasOracle_revert_callerNotOwner(address caller) public {
        vm.assume(caller != owner);
        expectUnauthorizedAccountRevert(caller);
        vm.prank(caller);
        service.setGasOracle(address(1));
    }
}
