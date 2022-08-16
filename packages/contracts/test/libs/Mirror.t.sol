// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { MirrorLib } from "../../contracts/libs/Mirror.sol";

import { SynapseTest } from "../utils/SynapseTest.sol";

contract MirrorTest is SynapseTest {
    using MirrorLib for MirrorLib.Mirror;

    MirrorLib.Mirror mirror;

    function setUp() public override {
        super.setUp();
        mirror.setupMirror(remoteDomain);
    }

    function test_setup() public {
        assertEq(mirror.nonce, 0);
        assertEq(mirror.remoteDomain, remoteDomain);
        assertEq(uint256(mirror.status), 1);
    }

    function test_setNonce(uint32 _nonce) public {
        mirror.setNonce(_nonce);
        assertEq(mirror.nonce, _nonce);
    }

    function test_setConfirmAt(bytes32 _committedRoot, uint256 _confirmAt) public {
        mirror.setConfirmAt(_committedRoot, _confirmAt);
        assertEq(mirror.confirmAt[_committedRoot], _confirmAt);
    }

    function test_setMessageStatus(bytes32 _messageHash, bytes32 _status) public {
        mirror.setMessageStatus(_messageHash, _status);
        assertEq(mirror.messageStatus[_messageHash], _status);
    }

    function test_setStatus() public {
        mirror.setStatus(MirrorLib.MirrorStatus.Failed);
        assertEq(uint256(mirror.status), 2);
    }
}
