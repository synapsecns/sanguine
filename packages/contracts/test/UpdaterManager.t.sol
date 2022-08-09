// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";
import { SynapseTestWithUpdaterManager } from "./utils/SynapseTest.sol";

import { HomeHarness } from "./harnesses/HomeHarness.sol";

import { UpdaterManager } from "../contracts/UpdaterManager.sol";
import { IUpdaterManager } from "../contracts/interfaces/IUpdaterManager.sol";

contract UpdaterManagerTest is SynapseTestWithUpdaterManager {
    HomeHarness home;
    HomeHarness fakeHome;

    function setUp() public override {
        super.setUp();
        home = new HomeHarness(localDomain);
        fakeHome = new HomeHarness(localDomain);
        home.initialize(IUpdaterManager(updaterManager));
        updaterManager.setHome(address(home));
    }

    function test_cannotSetHomeAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != updaterManager.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        updaterManager.setHome(address(fakeHome));
    }

    event NewHome(address home);

    function test_setHomeAsOwner() public {
        vm.startPrank(updaterManager.owner());
        assertEq(updaterManager.home(), address(home));
        vm.expectEmit(false, false, false, true);
        emit NewHome(address(fakeHome));
        updaterManager.setHome(address(fakeHome));
        assertEq(updaterManager.home(), address(fakeHome));
    }

    function test_cannotSetUpdaterAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != updaterManager.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        updaterManager.setUpdater(fakeUpdater);
    }

    event NewUpdater(address updater);

    function test_setUpdaterAsOwner() public {
        vm.startPrank(updaterManager.owner());
        assertEq(updaterManager.updater(), address(updater));
        assertTrue(home.isNotary(updater));
        vm.expectEmit(false, false, false, true);
        emit NewUpdater(fakeUpdater);
        updaterManager.setUpdater(address(fakeUpdater));
        assertEq(updaterManager.updater(), address(fakeUpdater));
        assertTrue(home.isNotary(fakeUpdater));
    }

    function test_cannotSlashUpdaterAsNotHome(address _notHome) public {
        vm.assume(_notHome != updaterManager.home());
        vm.startPrank(_notHome);
        vm.expectRevert("!home");
        updaterManager.slashUpdater(payable(address(this)));
    }

    event FakeSlashed(address reporter);

    function test_slashUpdaterAsHome() public {
        vm.startPrank(updaterManager.home());
        vm.expectEmit(false, false, false, true);
        emit FakeSlashed(address(this));
        updaterManager.slashUpdater(payable(address(this)));
    }
}
