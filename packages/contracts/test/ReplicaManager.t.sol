// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { TypedMemView } from "../contracts/libs/TypedMemView.sol";
import { TypeCasts } from "../contracts/libs/TypeCasts.sol";

import { Message } from "../contracts/libs/Message.sol";

import { ReplicaLib } from "../contracts/libs/Replica.sol";
import { ISystemMessenger } from "../contracts/interfaces/ISystemMessenger.sol";
import { ReplicaManagerHarness } from "./harnesses/ReplicaManagerHarness.sol";

import { AppHarness } from "./harnesses/AppHarness.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";

contract ReplicaManagerTest is SynapseTest {
    ReplicaManagerHarness replicaManager;
    AppHarness dApp;

    uint32 internal constant OPTIMISTIC_PERIOD = 10;

    bytes32 committedRoot;
    uint256 processGas;
    uint256 reserveGas;

    ISystemMessenger internal systemMessenger;

    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

    function setUp() public override {
        super.setUp();
        committedRoot = "";
        replicaManager = new ReplicaManagerHarness(localDomain);
        replicaManager.initialize(remoteDomain, updater);
        dApp = new AppHarness(OPTIMISTIC_PERIOD);
        systemMessenger = ISystemMessenger(address(1234567890));
        replicaManager.setSystemMessenger(systemMessenger);
    }

    // ============ INITIAL STATE ============
    function test_correctlyInitialized() public {
        assertEq(uint256(replicaManager.localDomain()), uint256(localDomain));
        assertEq(replicaManager.updater(), updater);
        // replicaManager set to active
        // assertEq(uint256(replicaManager.state()), 1);
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        replicaManager.initialize(remoteDomain, updater);
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
        bytes32 newRoot = keccak256(newMessage);
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
        bytes32 newRoot = "new root";
        vm.expectRevert("not current update");
        replicaManager.update(remoteDomain, newRoot, newRoot, bytes(""));
    }

    function test_updateWithIncorrectSig() public {
        bytes memory newMessage = "new root";
        bytes32 newRoot = keccak256(newMessage);
        bytes memory sig = signRemoteUpdate(fakeUpdaterPK, committedRoot, newRoot);
        vm.expectRevert("!updater sig");
        replicaManager.update(remoteDomain, committedRoot, newRoot, sig);
    }

    function test_acceptableRoot() public {
        bytes memory newMessage = "new root";
        bytes32 newRoot = keccak256(newMessage);
        uint32 optimisticSeconds = 69;
        test_successfulUpdate();
        vm.warp(block.timestamp + optimisticSeconds);
        assertTrue(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, newRoot));
    }

    function test_cannotAcceptableRoot() public {
        bytes32 newRoot = "new root";
        test_successfulUpdate();
        uint32 optimisticSeconds = 69;
        vm.warp(block.timestamp + optimisticSeconds - 1);
        assertFalse(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, newRoot));
    }

    function test_process() public {
        bytes memory message = _prepareProcessTest(OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD);
        replicaManager.process(message);
    }

    function test_processPeriodNotPassed() public {
        bytes memory message = _prepareProcessTest(OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("!optimisticSeconds");
        replicaManager.process(message);
    }

    function test_processForgedPeriodReduced() public {
        bytes memory message = _prepareProcessTest(OPTIMISTIC_PERIOD - 1);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("app: !optimisticSeconds");
        replicaManager.process(message);
    }

    function test_processForgePeriodZero() public {
        bytes memory message = _prepareProcessTest(0);
        vm.expectRevert("app: !optimisticSeconds");
        replicaManager.process(message);
    }

    function test_onlySystemMessenger() public {
        vm.prank(address(systemMessenger));
        replicaManager.setSensitiveValue(1337);
        assertEq(replicaManager.sensitiveValue(), 1337);
    }

    function test_onlySystemMessenger_rejectOthers() public {
        vm.expectRevert("!systemMessenger");
        replicaManager.setSensitiveValue(1337);
    }

    function _prepareProcessTest(uint32 optimisticPeriod) internal returns (bytes memory message) {
        test_successfulUpdate();

        bytes32 root = replicaManager.activeReplicaCommittedRoot(remoteDomain);
        assert(root != bytes32(0));

        uint32 nonce = 1234;
        bytes32 sender = "sender";
        bytes memory messageBody = "message body";
        dApp.prepare(remoteDomain, nonce, sender, messageBody);
        bytes32 recipient = TypeCasts.addressToBytes32(address(dApp));

        message = Message.formatMessage(
            remoteDomain,
            sender,
            nonce,
            localDomain,
            recipient,
            optimisticPeriod,
            messageBody
        );
        bytes32 messageHash = keccak256(message);
        // Let's imagine message was proved against current root
        replicaManager.setMessageStatus(remoteDomain, messageHash, root);
    }
}
