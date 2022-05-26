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

    event SetOptimisticTimeout(uint256 timeout);

    function test_setOptimistic(uint256 _optimisticSeconds) public {
        vm.assume(_optimisticSeconds != replica.optimisticSeconds());
        assertFalse(replica.optimisticSeconds() == _optimisticSeconds);
        vm.startPrank(replica.owner());
        vm.expectEmit(false, false, false, true);
        emit SetOptimisticTimeout(_optimisticSeconds);
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

    event SetConfirmation(bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt);

    function test_setConfirmation(uint256 _confirmAt) public {
        assertEq(replica.confirmAt(committedRoot), 1);
        vm.expectEmit(true, false, false, true);
        emit SetConfirmation(committedRoot, 1, _confirmAt);
        replica.setConfirmation(committedRoot, _confirmAt);
        assertEq(replica.confirmAt(committedRoot), _confirmAt);
    }

    event Update(
        uint32 indexed homeDomain,
        bytes32 indexed oldRoot,
        bytes32 indexed newRoot,
        bytes signature
    );

    // Relayer relays a new root signed by updater on Home chain
    function test_update() public {
        bytes32 newRoot = "new root";
        assertEq(replica.updater(), vm.addr(updaterPK));
        bytes memory sig = signRemoteUpdate(updaterPK, committedRoot, newRoot);
        // Root doesn't exist yet
        assertEq(replica.confirmAt(newRoot), 0);
        // Relayer sends over a root signed by the updater on the Home chain
        vm.expectEmit(true, true, true, true);
        emit Update(remoteDomain, committedRoot, newRoot, sig);
        replica.update(committedRoot, newRoot, sig);
        // Root set with optimistic latency allowing it to be processed at T+10
        assertEq(replica.confirmAt(newRoot), block.timestamp + 10);
        assertEq(replica.committedRoot(), newRoot);
    }

    function test_updateWithIncorrectRoot() public {
        bytes32 newRoot = "new root";
        vm.expectRevert("not current update");
        replica.update(newRoot, newRoot, bytes(""));
    }

    function test_updateWithIncorrectSig() public {
        bytes32 newRoot = "new root";
        bytes memory sig = signRemoteUpdate(fakeUpdaterPK, committedRoot, newRoot);
        vm.expectRevert("!updater sig");
        replica.update(committedRoot, newRoot, sig);
    }

    function test_acceptableRoot() public {
        bytes32 newRoot = "new root";
        test_update();
        vm.warp(block.timestamp + optimisticSeconds + 1);
        assertTrue(replica.acceptableRoot(newRoot));
    }

    function test_cannotAcceptableRoot() public {
        bytes32 newRoot = "new root";
        test_update();
        vm.warp(block.timestamp + optimisticSeconds - 1);
        assertFalse(replica.acceptableRoot(newRoot));
    }
}
