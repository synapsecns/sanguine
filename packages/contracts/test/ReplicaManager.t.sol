// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { TypedMemView } from "../contracts/libs/TypedMemView.sol";
import { TypeCasts } from "../contracts/libs/TypeCasts.sol";

import { Header } from "../contracts/libs/Header.sol";
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

    bytes32 internal constant ROOT = keccak256("test root");

    uint256 processGas;
    uint256 reserveGas;

    ISystemMessenger internal systemMessenger;

    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

    function setUp() public override {
        super.setUp();
        replicaManager = new ReplicaManagerHarness(localDomain);
        replicaManager.initialize(remoteDomain, notary);
        dApp = new AppHarness(OPTIMISTIC_PERIOD);
        systemMessenger = ISystemMessenger(address(1234567890));
        replicaManager.setSystemMessenger(systemMessenger);
    }

    // ============ INITIAL STATE ============
    function test_correctlyInitialized() public {
        assertEq(uint256(replicaManager.localDomain()), uint256(localDomain));
        assertTrue(replicaManager.isNotary(remoteDomain, notary));
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        replicaManager.initialize(remoteDomain, notary);
    }

    // ============ STATE & PERMISSIONING ============

    function test_cannotSetNotaryAsNotOwner(address _notOwner, address _notary) public {
        vm.assume(_notOwner != replicaManager.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replicaManager.setNotary(remoteDomain, _notary);
    }

    function test_setNotary(address _notary) public {
        vm.assume(_notary != notary);
        vm.prank(replicaManager.owner());
        replicaManager.setNotary(remoteDomain, _notary);
        assertTrue(replicaManager.isNotary(remoteDomain, _notary));
    }

    function test_cannotSetConfirmationAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != replicaManager.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        replicaManager.setConfirmation(remoteDomain, ROOT, 0);
    }

    event SetConfirmation(
        uint32 indexed remoteDomain,
        bytes32 indexed root,
        uint256 previousConfirmAt,
        uint256 newConfirmAt
    );

    function test_setConfirmation(uint256 _confirmAt) public {
        vm.startPrank(replicaManager.owner());
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, ROOT), 0);
        vm.expectEmit(true, true, true, true);
        emit SetConfirmation(remoteDomain, ROOT, 0, _confirmAt);
        replicaManager.setConfirmation(remoteDomain, ROOT, _confirmAt);
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, ROOT), _confirmAt);
    }

    event AttestationAccepted(
        uint32 indexed originDomain,
        uint32 indexed nonce,
        bytes32 indexed root,
        bytes signature
    );

    // Broadcaster relays a new root signed by notary on Origin chain
    function test_submitAttestation() public {
        uint32 nonce = 42;
        assertTrue(replicaManager.isNotary(remoteDomain, vm.addr(notaryPK)));
        (bytes memory attestation, bytes memory sig) = signRemoteAttestation(notaryPK, nonce, ROOT);
        // Root doesn't exist yet
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, ROOT), 0);
        // Broadcaster sends over a root signed by the notary on the Origin chain
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(remoteDomain, nonce, ROOT, sig);
        replicaManager.submitAttestation(attestation);
        // Time at which root was confirmed is set, optimistic timeout starts now
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, ROOT), block.timestamp);
    }

    function test_submitAttestation_fakeNotary() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(fakeNotaryPK, nonce, ROOT);
        vm.expectRevert("Signer is not a notary");
        // Attestation signed by fakeNotary should be rejected
        replicaManager.submitAttestation(attestation);
    }

    function test_submitAttestation_localDomain() public {
        replicaManager.addNotary(localDomain, notary);
        uint32 nonce = 42;
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, ROOT);
        vm.expectRevert("Attestation refers to local chain");
        // Replica should reject attestations from the chain it's deployed on
        replicaManager.submitAttestation(attestation);
    }

    function test_acceptableRoot() public {
        uint32 optimisticSeconds = 69;
        test_submitAttestation();
        vm.warp(block.timestamp + optimisticSeconds);
        assertTrue(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, ROOT));
    }

    function test_cannotAcceptableRoot() public {
        test_submitAttestation();
        uint32 optimisticSeconds = 69;
        vm.warp(block.timestamp + optimisticSeconds - 1);
        assertFalse(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, ROOT));
    }

    event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip);

    function test_execute() public {
        bytes memory message = _prepareExecuteTest(OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD);
        vm.expectEmit(true, true, true, true);
        emit LogTips(NOTARY_TIP, BROADCASTER_TIP, PROVER_TIP, EXECUTOR_TIP);
        replicaManager.execute(message);
    }

    function test_executePeriodNotPassed() public {
        bytes memory message = _prepareExecuteTest(OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("!optimisticSeconds");
        replicaManager.execute(message);
    }

    function test_executeForgedPeriodReduced() public {
        bytes memory message = _prepareExecuteTest(OPTIMISTIC_PERIOD - 1);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("app: !optimisticSeconds");
        replicaManager.execute(message);
    }

    function test_executeForgePeriodZero() public {
        bytes memory message = _prepareExecuteTest(0);
        vm.expectRevert("app: !optimisticSeconds");
        replicaManager.execute(message);
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

    function _prepareExecuteTest(uint32 optimisticPeriod) internal returns (bytes memory message) {
        test_submitAttestation();

        uint32 nonce = 1234;
        bytes32 sender = "sender";
        bytes memory messageBody = "message body";
        dApp.prepare(remoteDomain, nonce, sender, messageBody);
        bytes32 recipient = TypeCasts.addressToBytes32(address(dApp));

        bytes memory _header = Header.formatHeader(
            remoteDomain,
            sender,
            nonce,
            localDomain,
            recipient,
            optimisticPeriod
        );

        message = Message.formatMessage(_header, getDefaultTips(), messageBody);
        bytes32 messageHash = keccak256(message);
        // Let's imagine message was proved against current root
        replicaManager.setMessageStatus(remoteDomain, messageHash, ROOT);
    }
}
