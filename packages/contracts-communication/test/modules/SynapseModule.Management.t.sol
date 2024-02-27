// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseModuleEvents} from "../../contracts/events/SynapseModuleEvents.sol";
import {SynapseModule, ISynapseModule} from "../../contracts/modules/SynapseModule.sol";
import {ThresholdECDSALib} from "../../contracts/libs/ThresholdECDSA.sol";

import {GasOracleMock} from "../mocks/GasOracleMock.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseModuleManagementTest is Test, SynapseModuleEvents {
    SynapseModule public module;
    GasOracleMock public gasOracle;

    address public interchainDB = makeAddr("InterchainDB");
    address public owner = makeAddr("Owner");
    address public feeCollector = makeAddr("FeeCollector");

    address public constant VERIFIER_1 = address(1);
    address public constant VERIFIER_2 = address(2);
    address public constant VERIFIER_3 = address(3);

    address[] public allVerifiers;

    function setUp() public {
        module = new SynapseModule(interchainDB, owner);
        gasOracle = new GasOracleMock();
        allVerifiers.push(VERIFIER_1);
        allVerifiers.push(VERIFIER_2);
        allVerifiers.push(VERIFIER_3);
    }

    function basicSetup() internal {
        vm.startPrank(owner);
        module.addVerifier(VERIFIER_1);
        module.addVerifier(VERIFIER_2);
        module.addVerifier(VERIFIER_3);
        vm.stopPrank();
    }

    function test_setup() public {
        assertEq(module.owner(), owner);
        assertEq(module.INTERCHAIN_DB(), interchainDB);
        assertEq(module.getThreshold(), 0);
        assertEq(module.gasOracle(), address(0));
    }

    function test_basicSetup() public {
        basicSetup();
        assertTrue(module.isVerifier(VERIFIER_1));
        assertTrue(module.isVerifier(VERIFIER_2));
        assertTrue(module.isVerifier(VERIFIER_3));
        address[] memory verifiers = module.getVerifiers();
        assertEq(verifiers.length, 3);
        assertEq(verifiers[0], VERIFIER_1);
        assertEq(verifiers[1], VERIFIER_2);
        assertEq(verifiers[2], VERIFIER_3);
    }

    function test_addVerifier_addsVerifier() public {
        vm.prank(owner);
        module.addVerifier(VERIFIER_1);
        assertTrue(module.isVerifier(VERIFIER_1));
    }

    function test_addVerifier_emitsEvent() public {
        vm.expectEmit(address(module));
        emit VerifierAdded(VERIFIER_1);
        vm.prank(owner);
        module.addVerifier(VERIFIER_1);
    }

    function test_addVerifier_revert_alreadyAdded() public {
        basicSetup();
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__AlreadySigner.selector, VERIFIER_1));
        vm.prank(owner);
        module.addVerifier(VERIFIER_1);
    }

    function test_addVerifier_revert_notOwner(address notOwner) public {
        vm.assume(notOwner != owner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        module.addVerifier(VERIFIER_1);
    }

    function test_addVerifiers_addsVerifiers() public {
        vm.prank(owner);
        module.addVerifiers(allVerifiers);
        assertTrue(module.isVerifier(VERIFIER_1));
        assertTrue(module.isVerifier(VERIFIER_2));
        assertTrue(module.isVerifier(VERIFIER_3));
    }

    function test_addVerifiers_emitsEvents() public {
        vm.expectEmit(address(module));
        emit VerifierAdded(VERIFIER_1);
        vm.expectEmit(address(module));
        emit VerifierAdded(VERIFIER_2);
        vm.expectEmit(address(module));
        emit VerifierAdded(VERIFIER_3);
        vm.prank(owner);
        module.addVerifiers(allVerifiers);
    }

    function test_addVerifiers_revert_alreadyAdded() public {
        vm.prank(owner);
        module.addVerifier(VERIFIER_2);
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__AlreadySigner.selector, VERIFIER_2));
        vm.prank(owner);
        module.addVerifiers(allVerifiers);
    }

    function test_addVerifiers_revert_notOwner(address notOwner) public {
        vm.assume(notOwner != owner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        module.addVerifiers(allVerifiers);
    }

    function test_removeVerifier_removesVerifier() public {
        basicSetup();
        vm.prank(owner);
        module.removeVerifier(VERIFIER_1);
        assertFalse(module.isVerifier(VERIFIER_1));
    }

    function test_removeVerifier_emitsEvent() public {
        basicSetup();
        vm.expectEmit(address(module));
        emit VerifierRemoved(VERIFIER_1);
        vm.prank(owner);
        module.removeVerifier(VERIFIER_1);
    }

    function test_removeVerifier_revert_notAdded() public {
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__NotSigner.selector, VERIFIER_1));
        vm.prank(owner);
        module.removeVerifier(VERIFIER_1);
    }

    function test_removeVerifier_revert_notOwner(address notOwner) public {
        vm.assume(notOwner != owner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        module.removeVerifier(VERIFIER_1);
    }

    function test_removeVerifiers_removesVerifiers() public {
        basicSetup();
        vm.prank(owner);
        module.removeVerifiers(allVerifiers);
        assertFalse(module.isVerifier(VERIFIER_1));
        assertFalse(module.isVerifier(VERIFIER_2));
        assertFalse(module.isVerifier(VERIFIER_3));
    }

    function test_removeVerifiers_emitsEvents() public {
        basicSetup();
        vm.expectEmit(address(module));
        emit VerifierRemoved(VERIFIER_1);
        vm.expectEmit(address(module));
        emit VerifierRemoved(VERIFIER_2);
        vm.expectEmit(address(module));
        emit VerifierRemoved(VERIFIER_3);
        vm.prank(owner);
        module.removeVerifiers(allVerifiers);
    }

    function test_removeVerifiers_revert_notAdded() public {
        vm.prank(owner);
        module.addVerifier(VERIFIER_2);
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__NotSigner.selector, VERIFIER_1));
        vm.prank(owner);
        module.removeVerifiers(allVerifiers);
    }

    function test_removeVerifiers_revert_notOwner(address notOwner) public {
        vm.assume(notOwner != owner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        module.removeVerifiers(allVerifiers);
    }

    function test_setThreshold_setsThreshold() public {
        vm.prank(owner);
        module.setThreshold(3);
        assertEq(module.getThreshold(), 3);
    }

    function test_setThreshold_emitsEvent() public {
        vm.expectEmit(address(module));
        emit ThresholdChanged(3);
        vm.prank(owner);
        module.setThreshold(3);
    }

    function test_setThreshold_revert_zeroThreshold() public {
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__ZeroThreshold.selector));
        vm.prank(owner);
        module.setThreshold(0);
    }

    function test_setThreshold_revert_notOwner(address notOwner) public {
        vm.assume(notOwner != owner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        module.setThreshold(3);
    }

    function test_setFeeCollector_setsFeeCollector() public {
        vm.prank(owner);
        module.setFeeCollector(feeCollector);
        assertEq(module.feeCollector(), feeCollector);
    }

    function test_setFeeCollector_emitsEvent() public {
        vm.expectEmit(address(module));
        emit FeeCollectorChanged(feeCollector);
        vm.prank(owner);
        module.setFeeCollector(feeCollector);
    }

    function test_setFeeCollector_revert_notOwner(address notOwner) public {
        vm.assume(notOwner != owner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        module.setFeeCollector(feeCollector);
    }

    function test_setGasOracle_setsGasOracle() public {
        vm.prank(owner);
        module.setGasOracle(address(gasOracle));
        assertEq(module.gasOracle(), address(gasOracle));
    }

    function test_setGasOracle_emitsEvent() public {
        vm.expectEmit(address(module));
        emit GasOracleChanged(address(gasOracle));
        vm.prank(owner);
        module.setGasOracle(address(gasOracle));
    }

    function test_setGasOracle_revert_notOwner(address notOwner) public {
        vm.assume(notOwner != owner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        module.setGasOracle(address(gasOracle));
    }

    function test_setGasOracle_revert_notContract() public {
        address notContract = makeAddr("NotContract");
        // Sanity check
        assert(notContract.code.length == 0);
        vm.expectRevert(
            abi.encodeWithSelector(ISynapseModule.SynapseModule__GasOracleNotContract.selector, notContract)
        );
        vm.prank(owner);
        module.setGasOracle(notContract);
    }
}
