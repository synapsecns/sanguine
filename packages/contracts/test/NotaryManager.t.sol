// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { OriginHarness } from "./harnesses/OriginHarness.sol";

import { NotaryManager } from "../contracts/NotaryManager.sol";
import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";

contract NotaryManagerTest is SynapseTestWithNotaryManager {
    OriginHarness origin;
    OriginHarness fakeOrigin;

    function setUp() public override {
        super.setUp();
        origin = new OriginHarness(localDomain);
        fakeOrigin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));
    }

    function test_cannotSetOriginAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != notaryManager.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        notaryManager.setOrigin(address(fakeOrigin));
    }

    event NewOrigin(address origin);

    function test_setOriginAsOwner() public {
        vm.startPrank(notaryManager.owner());
        assertEq(notaryManager.origin(), address(origin));
        vm.expectEmit(false, false, false, true);
        emit NewOrigin(address(fakeOrigin));
        notaryManager.setOrigin(address(fakeOrigin));
        assertEq(notaryManager.origin(), address(fakeOrigin));
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
        assertTrue(origin.isNotary(notary));
        vm.expectEmit(false, false, false, true);
        emit NewNotary(fakeNotary);
        notaryManager.setNotary(address(fakeNotary));
        assertEq(notaryManager.notary(), address(fakeNotary));
        assertTrue(origin.isNotary(fakeNotary));
    }

    function test_cannotSlashNotaryAsNotOrigin(address _notOrigin) public {
        vm.assume(_notOrigin != notaryManager.origin());
        vm.startPrank(_notOrigin);
        vm.expectRevert("!origin");
        notaryManager.slashNotary(payable(address(this)));
    }

    event FakeSlashed(address reporter);

    function test_slashNotaryAsOrigin() public {
        vm.startPrank(notaryManager.origin());
        vm.expectEmit(false, false, false, true);
        emit FakeSlashed(address(this));
        notaryManager.slashNotary(payable(address(this)));
    }
}
