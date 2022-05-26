// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { TypedMemView } from "../contracts/libs/TypedMemView.sol";
import { Message } from "../contracts/libs/Message.sol";

import { ReplicaHarness } from "./harnesses/ReplicaHarness.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";

contract ReplicaTest is SynapseTest {
    ReplicaHarness replica;

    uint256 optimisticSeconds;
    bytes32 committedRoot;
    uint256 processGas;
    uint256 reserveGas;

    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

    function setUp() public override {
        super.setUp();
        optimisticSeconds = 10;
        committedRoot = "commited root";
        processGas = 850_000;
        reserveGas = 15_000;
        replica = new ReplicaHarness(localDomain, processGas, reserveGas);
        replica.initialize(remoteDomain, updater, committedRoot, optimisticSeconds);
    }

    // ============ INITIAL STATE ============
    function test_correctlyInitialized() public {
        assertEq(uint256(replica.PROCESS_GAS()), processGas);
        assertEq(uint256(replica.RESERVE_GAS()), reserveGas);
        assertEq(uint256(replica.localDomain()), uint256(localDomain));
        assertEq(uint256(replica.remoteDomain()), uint256(remoteDomain));
        assertEq(replica.committedRoot(), committedRoot);
        assertEq(replica.optimisticSeconds(), optimisticSeconds);
        assertEq(replica.confirmAt(committedRoot), 1);
        assertEq(replica.updater(), updater);
        // replica set to active
        assertEq(uint256(replica.state()), 1);
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        replica.initialize(remoteDomain, updater, committedRoot, optimisticSeconds);
    }

    // ============ STATE & PERMISSIONING ============
    function test_setOptimisticNotOwner(address _notOwner, uint256 _optimisticSeconds) public {
        vm.assume(_notOwner != replica.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replica.setOptimisticTimeout(_optimisticSeconds);
    }

    function test_setOptimistic(uint256 _optimisticSeconds) public {
        vm.assume(_optimisticSeconds != replica.optimisticSeconds());
        assertFalse(replica.optimisticSeconds() == _optimisticSeconds);
        vm.prank(replica.owner());
        replica.setOptimisticTimeout(_optimisticSeconds);
        assertEq(replica.optimisticSeconds(), _optimisticSeconds);
    }

    function test_cannotSetUpdaterAsNotOwner(address _notOwner, address _updater) public {
        vm.assume(_notOwner != replica.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replica.setUpdater(_updater);
    }

    function test_setUpdater(address _updater) public {
        vm.assume(_updater != replica.updater());
        vm.prank(replica.owner());
        replica.setUpdater(_updater);
        assertEq(replica.updater(), _updater);
    }

    function test_cannotSetConfirmationAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != replica.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replica.setConfirmation(committedRoot, 0);
    }

    function test_setConfirmation(uint256 _confirmAt) public {
        assertEq(replica.confirmAt(committedRoot), 1);
        replica.setConfirmation(committedRoot, _confirmAt);
        assertEq(replica.confirmAt(committedRoot), _confirmAt);
    }
}
