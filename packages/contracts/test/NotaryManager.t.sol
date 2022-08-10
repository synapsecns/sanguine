// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { HomeHarness } from "./harnesses/HomeHarness.sol";

import { NotaryManager } from "../contracts/NotaryManager.sol";
import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";

contract NotaryManagerTest is SynapseTestWithNotaryManager {
    HomeHarness home;
    HomeHarness fakeHome;

    function setUp() public override {
        super.setUp();
        home = new HomeHarness(localDomain);
        fakeHome = new HomeHarness(localDomain);
        home.initialize(INotaryManager(notaryManager));
        notaryManager.setHome(address(home));
    }

    function test_cannotSetHomeAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != notaryManager.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        notaryManager.setHome(address(fakeHome));
    }

    event NewHome(address home);

    function test_setHomeAsOwner() public {
        vm.startPrank(notaryManager.owner());
        assertEq(notaryManager.home(), address(home));
        vm.expectEmit(false, false, false, true);
        emit NewHome(address(fakeHome));
        notaryManager.setHome(address(fakeHome));
        assertEq(notaryManager.home(), address(fakeHome));
    }

    function test_cannotSetNotaryAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != notaryManager.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        notaryManager.setNotary(fakeNotary);
    }

    event NewNotary(address notary);

    function test_setNotaryAsOwner() public {
        vm.startPrank(notaryManager.owner());
        assertEq(notaryManager.notary(), address(notary));
        assertTrue(home.isNotary(notary));
        vm.expectEmit(false, false, false, true);
        emit NewNotary(fakeNotary);
        notaryManager.setNotary(address(fakeNotary));
        assertEq(notaryManager.notary(), address(fakeNotary));
        assertTrue(home.isNotary(fakeNotary));
    }

    function test_cannotSlashNotaryAsNotHome(address _notHome) public {
        vm.assume(_notHome != notaryManager.home());
        vm.startPrank(_notHome);
        vm.expectRevert("!home");
        notaryManager.slashNotary(payable(address(this)));
    }

    event FakeSlashed(address reporter);

    function test_slashNotaryAsHome() public {
        vm.startPrank(notaryManager.home());
        vm.expectEmit(false, false, false, true);
        emit FakeSlashed(address(this));
        notaryManager.slashNotary(payable(address(this)));
    }
}
