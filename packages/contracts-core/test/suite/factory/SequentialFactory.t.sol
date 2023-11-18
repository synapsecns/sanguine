// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SequentialFactory} from "../../../contracts/factory/SequentialFactory.sol";

import {A, B, UnDeployable} from "./Contracts.sol";

import {Test} from "forge-std/Test.sol";

contract SequentialFactoryTest is Test {
    SequentialFactory public factory;
    address public owner;
    address public attacker;

    function setUp() public {
        owner = makeAddr("Owner");
        attacker = makeAddr("Attacker");
        factory = new SequentialFactory(owner);
    }

    function testSetup() public {
        assertEq(address(factory.owner()), owner);
        // Nonce starts at 1
        assertEq(factory.getNonce(), 1);
    }

    // ═════════════════════════════════════════════ DEFINE TEST CASES ═════════════════════════════════════════════════

    function getOneArgCode(uint256 value) internal pure returns (bytes memory code) {
        return abi.encodePacked(type(B).creationCode, abi.encode(value));
    }

    function getCodeTenDeployments(uint256 index) internal pure returns (bytes memory code) {
        // First 10 deployments are A, B, A, B, A, B, A, B, A, B
        if (index % 2 == 0) {
            return type(A).creationCode;
        } else {
            return getOneArgCode(index);
        }
    }

    function deploy(bytes memory code) internal returns (address deployedAt) {
        uint256 nonce = factory.getNonce();
        vm.prank(owner);
        return factory.deploy(nonce, code);
    }

    function deployTenContracts() internal {
        for (uint256 i = 0; i < 10; i++) {
            deploy(getCodeTenDeployments(i));
        }
    }

    // ═══════════════════════════════════════ HAPPY PATH: FIRST DEPLOYMENT ════════════════════════════════════════════

    function test_deploy_nonce01_contractA_correctContract() public {
        address deployedAt = deploy(type(A).creationCode);
        A a = A(deployedAt);
        assertEq(a.value(), 1337);
    }

    function test_deploy_nonce01_contractA_incrementsNonce() public {
        deploy(type(A).creationCode);
        assertEq(factory.getNonce(), 2);
    }

    function test_deploy_nonce01_contractA_predictDeployment() public {
        address predicted = factory.predictDeployment(1);
        address deployedAt = deploy(type(A).creationCode);
        assertEq(predicted, deployedAt);
    }

    function test_deploy_nonce01_contractA_predictNextDeployment() public {
        address predicted = factory.predictNextDeployment();
        address deployedAt = deploy(type(A).creationCode);
        assertEq(predicted, deployedAt);
    }

    function test_deploy_nonce01_contractB_correctContract() public {
        address deployedAt = deploy(getOneArgCode(42));
        B b = B(deployedAt);
        assertEq(b.value(), 42);
    }

    function test_deploy_nonce01_contractB_incrementsNonce() public {
        deploy(getOneArgCode(42));
        assertEq(factory.getNonce(), 2);
    }

    function test_deploy_nonce01_contractB_predictDeployment() public {
        address predicted = factory.predictDeployment(1);
        address deployedAt = deploy(getOneArgCode(42));
        assertEq(predicted, deployedAt);
    }

    function test_deploy_nonce01_contractB_predictNextDeployment() public {
        address predicted = factory.predictNextDeployment();
        address deployedAt = deploy(getOneArgCode(42));
        assertEq(predicted, deployedAt);
    }

    // ═════════════════════════════════════ HAPPY PATH: NON-FIRST DEPLOYMENT ══════════════════════════════════════════

    function test_deploy_nonce11_contractA_correctContract() public {
        deployTenContracts();
        address deployedAt = deploy(type(A).creationCode);
        A a = A(deployedAt);
        assertEq(a.value(), 1337);
    }

    function test_deploy_nonce11_contractA_incrementsNonce() public {
        deployTenContracts();
        deploy(type(A).creationCode);
        assertEq(factory.getNonce(), 12);
    }

    function test_deploy_nonce11_contractA_predictDeployment() public {
        deployTenContracts();
        address predicted = factory.predictDeployment(11);
        address deployedAt = deploy(type(A).creationCode);
        assertEq(predicted, deployedAt);
    }

    function test_deploy_nonce11_contractA_predictNextDeployment() public {
        deployTenContracts();
        address predicted = factory.predictNextDeployment();
        address deployedAt = deploy(type(A).creationCode);
        assertEq(predicted, deployedAt);
    }

    function test_deploy_nonce11_contractB_correctContract() public {
        deployTenContracts();
        address deployedAt = deploy(getOneArgCode(42));
        B b = B(deployedAt);
        assertEq(b.value(), 42);
    }

    function test_deploy_nonce11_contractB_incrementsNonce() public {
        deployTenContracts();
        deploy(getOneArgCode(42));
        assertEq(factory.getNonce(), 12);
    }

    function test_deploy_nonce11_contractB_predictDeployment() public {
        deployTenContracts();
        address predicted = factory.predictDeployment(11);
        address deployedAt = deploy(getOneArgCode(42));
        assertEq(predicted, deployedAt);
    }

    function test_deploy_nonce11_contractB_predictNextDeployment() public {
        deployTenContracts();
        address predicted = factory.predictNextDeployment();
        address deployedAt = deploy(getOneArgCode(42));
        assertEq(predicted, deployedAt);
    }

    // ════════════════════════════════════════ PREDICT DEPLOYMENT ADDRESS ═════════════════════════════════════════════

    function test_predictDeployment_tenContracts() public {
        for (uint256 i = 0; i < 10; i++) {
            address predicted = factory.predictDeployment(i + 1);
            address deployedAt = deploy(getCodeTenDeployments(i));
            assertEq(predicted, deployedAt);
        }
    }

    function test_predictNextDeployment_tenContracts() public {
        for (uint256 i = 0; i < 10; i++) {
            address predicted = factory.predictNextDeployment();
            address deployedAt = deploy(getCodeTenDeployments(i));
            assertEq(predicted, deployedAt);
        }
    }

    // ══════════════════════════════════════════════════ REVERTS ══════════════════════════════════════════════════════

    function test_deploy_nonce01_reverts_whenCallerNotOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(attacker);
        factory.deploy(1, type(A).creationCode);
    }

    function test_deploy_nonce11_reverts_whenCallerNotOwner() public {
        deployTenContracts();
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(attacker);
        factory.deploy(11, type(A).creationCode);
    }

    function test_deploy_nonce01_reverts_whenNonceIsLower() public {
        vm.expectRevert(abi.encodeWithSelector(SequentialFactory.SequentialFactory__InvalidNonce.selector, 0, 1));
        vm.prank(owner);
        factory.deploy(0, type(A).creationCode);
    }

    function test_deploy_nonce01_reverts_whenNonceIsHigher() public {
        vm.expectRevert(abi.encodeWithSelector(SequentialFactory.SequentialFactory__InvalidNonce.selector, 2, 1));
        vm.prank(owner);
        factory.deploy(2, type(A).creationCode);
    }

    function test_deploy_nonce11_reverts_whenNonceIsLower() public {
        deployTenContracts();
        vm.expectRevert(abi.encodeWithSelector(SequentialFactory.SequentialFactory__InvalidNonce.selector, 10, 11));
        vm.prank(owner);
        factory.deploy(10, type(A).creationCode);
    }

    function test_deploy_nonce11_reverts_whenNonceIsHigher() public {
        deployTenContracts();
        vm.expectRevert(abi.encodeWithSelector(SequentialFactory.SequentialFactory__InvalidNonce.selector, 12, 11));
        vm.prank(owner);
        factory.deploy(12, type(A).creationCode);
    }

    function test_deploy_nonce01_reverts_whenDeploymentFails() public {
        vm.expectRevert(abi.encodeWithSelector(SequentialFactory.SequentialFactory__DeploymentFailed.selector, 1));
        vm.prank(owner);
        factory.deploy(1, type(UnDeployable).creationCode);
    }

    function test_deploy_nonce11_reverts_whenDeploymentFails() public {
        deployTenContracts();
        vm.expectRevert(abi.encodeWithSelector(SequentialFactory.SequentialFactory__DeploymentFailed.selector, 11));
        vm.prank(owner);
        factory.deploy(11, type(UnDeployable).creationCode);
    }
}
