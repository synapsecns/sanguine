// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {Create2Harness, Create2} from "../harnesses/Create2Harness.sol";

import {OneArgContract, RevertingContract} from "../mocks/SimpleContracts.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract Create2LibraryTest is Test {
    Create2Harness public libHarness;
    address public oneArgContract;

    function setUp() public {
        libHarness = new Create2Harness();
        oneArgContract = address(new OneArgContract(1337));
    }

    function test_create2(uint256 ethValue, bytes32 salt, uint256 argValue) public {
        deal(address(this), ethValue);
        bytes memory creationCode = abi.encodePacked(type(OneArgContract).creationCode, abi.encode(argValue));
        address predicted = libHarness.predictDeployment(address(libHarness), salt, creationCode);
        address deployed = libHarness.deploy{value: ethValue}(ethValue, salt, creationCode);
        assertEq(deployed, predicted);
        assertEq(OneArgContract(deployed).value(), argValue);
        assertEq(deployed.code, oneArgContract.code);
        assertEq(deployed.balance, ethValue);
    }

    function test_create2_revert_alreadyDeployed(uint256 ethValue, bytes32 salt, uint256 argValue) public {
        deal(address(this), ethValue);
        bytes memory creationCode = abi.encodePacked(type(OneArgContract).creationCode, abi.encode(argValue));
        address deployed = libHarness.deploy{value: ethValue}(ethValue, salt, creationCode);
        deal(address(this), ethValue);
        vm.expectRevert(abi.encodeWithSelector(Create2.Create2__AlreadyDeployed.selector, deployed));
        libHarness.deploy{value: ethValue}(ethValue, salt, creationCode);
    }

    function test_create2_revert_deploymentFailed(uint256 ethValue, bytes32 salt) public {
        deal(address(this), ethValue);
        bytes memory creationCode = type(RevertingContract).creationCode;
        vm.expectRevert(abi.encodeWithSelector(Create2.Create2__DeploymentFailed.selector));
        libHarness.deploy{value: ethValue}(ethValue, salt, creationCode);
    }
}
