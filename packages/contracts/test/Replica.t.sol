// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { Replica } from "../contracts/Replica.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";

contract ReplicaTest is SynapseTest {
    Replica replica;
    uint256 optimisticSeconds;

    function setUp() public override {
        super.setUp();
        optimisticSeconds = 10;
        replica = new Replica(address(1337), remoteDomain, optimisticSeconds);
    }

    function test_constructor() public {
        assertEq(replica.committedRoot(), bytes32(""));
        assertEq(replica.remoteDomain(), remoteDomain);
        assertEq(replica.optimisticSeconds(), optimisticSeconds);
        assertEq(uint256(replica.status()), 1);
    }

    function test_onlyReplicaManager(address notReplicaManager) public {
        vm.assume(notReplicaManager != replica.replicaManager());
        vm.expectRevert("!replica");
        replica.setCommittedRoot(bytes32(""));
        vm.expectRevert("!replica");
        replica.setConfirmAt(bytes32(""), 0);
        vm.expectRevert("!replica");
        replica.setConfirmAt(bytes32(""), 0);
        vm.expectRevert("!replica");
        replica.setMessageStatus(bytes32(""), Replica.MessageStatus.Processed);
        vm.expectRevert("!replica");
        replica.setOptimisticTimeout(10);
        vm.expectRevert("!replica");
        replica.setStatus(Replica.ReplicaStatus.Failed);
    }

    function test_setCommittedRoot(bytes32 _committedRoot) public {
        vm.prank(address(1337));
        replica.setCommittedRoot(_committedRoot);
        assertEq(replica.committedRoot(), _committedRoot);
    }

    function test_setConfirmAt(bytes32 _committedRoot, uint256 _confirmAt) public {
        vm.prank(address(1337));
        replica.setConfirmAt(_committedRoot, _confirmAt);
        assertEq(replica.confirmAt(_committedRoot), _confirmAt);
    }

    function test_setMessageStatus(bytes32 _messageHash) public {
        vm.prank(address(1337));
        replica.setMessageStatus(_messageHash, Replica.MessageStatus.Processed);
        assertEq(uint256(replica.messages(_messageHash)), 2);
    }

    function test_setOptimisticTimeout(uint256 _optimisticSeconds) public {
        vm.prank(address(1337));
        replica.setOptimisticTimeout(_optimisticSeconds);
        assertEq(replica.optimisticSeconds(), _optimisticSeconds);
    }

    function test_setStatus() public {
        vm.prank(address(1337));
        replica.setStatus(Replica.ReplicaStatus.Failed);
        assertEq(uint256(replica.status()), 2);
    }

    function test_setReplicaManagerOnlyOwner() public {
        replica.setReplicaManager(address(9999));
        assertEq(replica.replicaManager(), address(9999));
    }

    function test_cannotSetReplicaManager(address _notOwner) public {
        vm.assume(_notOwner != replica.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replica.setReplicaManager(address(9999));
    }
}
