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

    function test_setExecutorEOA_correctSlotERC7201() public {
        bytes32 slot = getExpectedLocationERC7201({namespaceId: "Synapse.ExecutionService.V1", stolOffset: 0});
        vm.prank(governor);
        service.setExecutorEOA(executorEOA);
        assertStorageAddress(address(service), slot, executorEOA);
    }

    function test_setExecutorEOA_revert_notGovernor(address caller) public {
        assumeNotProxyAdmin({target: address(service), caller: caller});
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

    function test_setGasOracle_correctSlotERC7201() public {
        bytes32 slot = getExpectedLocationERC7201({namespaceId: "Synapse.ExecutionService.V1", stolOffset: 1});
        vm.prank(governor);
        service.setGasOracle(gasOracle);
        assertStorageAddress(address(service), slot, gasOracle);
    }

    function test_setGasOracle_revert_notGovernor(address caller) public {
        assumeNotProxyAdmin({target: address(service), caller: caller});
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

    function test_setGlobalMarkup() public {
        uint256 globalMarkup = 100;
        expectEventGlobalMarkupSet(globalMarkup);
        vm.prank(governor);
        service.setGlobalMarkup(globalMarkup);
        assertEq(service.globalMarkup(), globalMarkup);
    }

    function test_setGlobalMarkup_correctSlotERC7201() public {
        uint256 globalMarkup = 100;
        bytes32 slot = getExpectedLocationERC7201({namespaceId: "Synapse.ExecutionService.V1", stolOffset: 2});
        vm.prank(governor);
        service.setGlobalMarkup(globalMarkup);
        assertStorageUint(address(service), slot, globalMarkup);
    }

    function test_setGlobalMarkup_toZero() public {
        test_setGlobalMarkup();
        expectEventGlobalMarkupSet(0);
        vm.prank(governor);
        service.setGlobalMarkup(0);
        assertEq(service.globalMarkup(), 0);
    }

    function test_setGlobalMarkup_revert_notGovernor(address caller) public {
        assumeNotProxyAdmin({target: address(service), caller: caller});
        vm.assume(caller != governor);
        expectRevertNotGovernor(caller);
        vm.prank(caller);
        service.setGlobalMarkup(100);
    }

    function test_setClaimerFraction() public {
        uint256 claimerFraction = 1e16;
        expectEventClaimerFractionSet(claimerFraction);
        vm.prank(governor);
        service.setClaimerFraction(claimerFraction);
        assertEq(service.getClaimerFraction(), claimerFraction);
    }

    function test_setClaimerFraction_correctSlotERC7201() public {
        uint256 claimerFraction = 1e16;
        bytes32 slot = getExpectedLocationERC7201({namespaceId: "Synapse.ExecutionService.V1", stolOffset: 3});
        vm.prank(governor);
        service.setClaimerFraction(claimerFraction);
        assertStorageUint(address(service), slot, claimerFraction);
    }

    function test_setClaimerFraction_revert_overOnePercent() public {
        expectRevertClaimerFractionExceedsMax(1e16 + 1);
        vm.prank(governor);
        service.setClaimerFraction(1e16 + 1);
    }

    function test_setClaimerFraction_revert_notGovernor(address caller) public {
        assumeNotProxyAdmin({target: address(service), caller: caller});
        vm.assume(caller != governor);
        expectRevertNotGovernor(caller);
        vm.prank(caller);
        service.setClaimerFraction(1e16);
    }

    function test_getExecutionFee_revert_gasOracleNotSet() public {
        expectRevertGasOracleNotSet();
        service.getExecutionFee(1, 2, "");
    }
}
