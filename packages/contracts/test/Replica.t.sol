// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { ReplicaLib } from "../contracts/libs/Replica.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";

contract ReplicaTest is SynapseTest {
    using ReplicaLib for ReplicaLib.Replica;

    ReplicaLib.Replica replica;

    function setUp() public override {
        super.setUp();
        replica.setupReplica(remoteDomain);
    }

    function test_setup() public {
        assertEq(replica.nonce, 0);
        assertEq(replica.remoteDomain, remoteDomain);
        assertEq(uint256(replica.status), 1);
    }

    function test_setNonce(uint32 _nonce) public {
        replica.setNonce(_nonce);
        assertEq(replica.nonce, _nonce);
    }

    function test_setConfirmAt(bytes32 _committedRoot, uint256 _confirmAt) public {
        replica.setConfirmAt(_committedRoot, _confirmAt);
        assertEq(replica.confirmAt[_committedRoot], _confirmAt);
    }

    function test_setMessageStatus(bytes32 _messageHash, bytes32 _status) public {
        replica.setMessageStatus(_messageHash, _status);
        assertEq(replica.messageStatus[_messageHash], _status);
    }

    function test_setStatus() public {
        replica.setStatus(ReplicaLib.ReplicaStatus.Failed);
        assertEq(uint256(replica.status), 2);
    }
}
