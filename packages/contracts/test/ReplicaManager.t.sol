// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { TypedMemView } from "../contracts/libs/TypedMemView.sol";
import { Message } from "../contracts/libs/Message.sol";

import { ReplicaLib } from "../contracts/libs/Replica.sol";

import { ReplicaManagerHarness } from "./harnesses/ReplicaManagerHarness.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";

contract ReplicaManagerTest is SynapseTest {
    ReplicaManagerHarness replicaManager;

    uint32 optimisticSeconds;
    bytes32 committedRoot;
    uint256 processGas;
    uint256 reserveGas;

    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

    function setUp() public override {
        super.setUp();
        optimisticSeconds = 10;
        committedRoot = "";
        processGas = 850_000;
        reserveGas = 15_000;
        replicaManager = new ReplicaManagerHarness(localDomain, processGas, reserveGas);
        replicaManager.initialize(remoteDomain, updater, optimisticSeconds);
    }

    // ============ INITIAL STATE ============
    function test_correctlyInitialized() public {
        assertEq(uint256(replicaManager.PROCESS_GAS()), processGas);
        assertEq(uint256(replicaManager.RESERVE_GAS()), reserveGas);
        assertEq(uint256(replicaManager.localDomain()), uint256(localDomain));
        assertEq(replicaManager.updater(), updater);
        // replicaManager set to active
        // assertEq(uint256(replicaManager.state()), 1);
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        replicaManager.initialize(remoteDomain, updater, optimisticSeconds);
    }

    // ============ STATE & PERMISSIONING ============

    function test_cannotSetUpdaterAsNotOwner(address _notOwner, address _updater) public {
        vm.assume(_notOwner != replicaManager.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replicaManager.setUpdater(_updater);
    }

    function test_setUpdater(address _updater) public {
        vm.assume(_updater != replicaManager.updater());
        vm.prank(replicaManager.owner());
        replicaManager.setUpdater(_updater);
        assertEq(replicaManager.updater(), _updater);
    }

    function test_cannotSetOptimisticTimeoutAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != replicaManager.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replicaManager.setOptimisticTimeout(remoteDomain, 10);
    }

    event SetOptimisticTimeout(uint32 indexed remoteDomain, uint32 timeout);

    function test_setOptimisticTimeout(uint32 _optimisticSeconds) public {
        vm.startPrank(replicaManager.owner());
        assertEq(replicaManager.activeReplicaOptimisticSeconds(remoteDomain), 10);
        vm.expectEmit(true, false, false, true);
        emit SetOptimisticTimeout(remoteDomain, _optimisticSeconds);
        replicaManager.setOptimisticTimeout(remoteDomain, _optimisticSeconds);
        assertEq(replicaManager.activeReplicaOptimisticSeconds(remoteDomain), _optimisticSeconds);
    }

    function test_cannotSetConfirmationAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != replicaManager.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replicaManager.setConfirmation(remoteDomain, committedRoot, 0);
    }

    event SetConfirmation(
        uint32 indexed remoteDomain,
        bytes32 indexed root,
        uint256 previousConfirmAt,
        uint256 newConfirmAt
    );

    function test_setConfirmation(uint256 _confirmAt) public {
        vm.startPrank(replicaManager.owner());
        bytes32 activeCommittedRoot = replicaManager.activeReplicaCommittedRoot(remoteDomain);
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, activeCommittedRoot), 0);
        vm.expectEmit(true, true, false, true);
        emit SetConfirmation(remoteDomain, committedRoot, 0, _confirmAt);
        replicaManager.setConfirmation(remoteDomain, activeCommittedRoot, _confirmAt);
        assertEq(
            replicaManager.activeReplicaConfirmedAt(remoteDomain, activeCommittedRoot),
            _confirmAt
        );
    }

    event Update(
        uint32 indexed homeDomain,
        bytes32 indexed oldRoot,
        bytes32 indexed newRoot,
        bytes signature
    );

    // Relayer relays a new root signed by updater on Home chain
    function test_successfulUpdate() public {
        bytes memory newMessage = "new root";
        bytes32 newRoot = keccak256(abi.encodePacked(newMessage, optimisticSeconds));
        assertEq(replicaManager.updater(), vm.addr(updaterPK));
        bytes memory sig = signRemoteUpdate(updaterPK, committedRoot, newRoot);
        // Root doesn't exist yet
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, newRoot), 0);
        // Relayer sends over a root signed by the updater on the Home chain
        vm.expectEmit(true, true, true, true);
        emit Update(remoteDomain, committedRoot, newRoot, sig);
        replicaManager.update(remoteDomain, committedRoot, newRoot, sig);
        // Time at which root was confirmed is set, optimistic timeout starts now
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, newRoot), block.timestamp);
        assertEq(replicaManager.activeReplicaCommittedRoot(remoteDomain), newRoot);
    }

    function test_updateWithIncorrectRoot() public {
        bytes memory newMessage = "new root";
        bytes32 newRoot = "new root";
        vm.expectRevert("not current update");
        replicaManager.update(remoteDomain, newRoot, newRoot, bytes(""));
    }

    function test_updateWithIncorrectSig() public {
        bytes memory newMessage = "new root";
        bytes32 newRoot = keccak256(abi.encodePacked(newMessage, optimisticSeconds));
        bytes memory sig = signRemoteUpdate(fakeUpdaterPK, committedRoot, newRoot);
        vm.expectRevert("!updater sig");
        replicaManager.update(remoteDomain, committedRoot, newRoot, sig);
    }

    function test_acceptableRoot() public {
        bytes memory newMessage = "new root";
        bytes32 newRoot = keccak256(abi.encodePacked(newMessage, optimisticSeconds));
        test_successfulUpdate();
        vm.warp(block.timestamp + optimisticSeconds + 1);
        assertTrue(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, newRoot));
    }

    function test_cannotAcceptableRoot() public {
        bytes32 newRoot = "new root";
        test_successfulUpdate();
        vm.warp(block.timestamp + optimisticSeconds - 1);
        assertFalse(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, newRoot));
    }
}
