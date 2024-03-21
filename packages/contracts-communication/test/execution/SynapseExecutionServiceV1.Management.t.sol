// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceV1Test} from "./SynapseExecutionServiceV1.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseExecutionServiceV1ManagementTest is SynapseExecutionServiceV1Test {
    address public executorEOA = makeAddr("ExecutorEOA");
    address public gasOracle = makeAddr("GasOracle");

    function setUp() public override {
        super.setUp();
        service.initialize(admin);
        vm.prank(admin);
        service.grantRole(GOVERNOR_ROLE, governor);
    }

    function test_setExecutorEOA() public {
        expectEventExecutorEOASet(executorEOA);
        vm.prank(governor);
        service.setExecutorEOA(executorEOA);
        assertEq(service.executorEOA(), executorEOA);
    }

    function test_setExecutorEOA_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertNotGovernor(caller);
        vm.prank(caller);
        service.setExecutorEOA(executorEOA);
    }

    function test_setExecutorEOA_revert_zeroAddress() public {
        expectRevertZeroAddress();
        vm.prank(governor);
        service.setExecutorEOA(address(0));
    }

    function test_setGasOracle() public {
        expectEventGasOracleSet(gasOracle);
        vm.prank(governor);
        service.setGasOracle(gasOracle);
        assertEq(service.gasOracle(), gasOracle);
    }

    function test_setGasOracle_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertNotGovernor(caller);
        vm.prank(caller);
        service.setGasOracle(gasOracle);
    }

    function test_setGasOracle_revert_zeroAddress() public {
        expectRevertZeroAddress();
        vm.prank(governor);
        service.setGasOracle(address(0));
    }

    function test_getExecutionFee_revert_gasOracleNotSet() public {
        expectRevertGasOracleNotSet();
        service.getExecutionFee(1, 2, "");
    }
}
