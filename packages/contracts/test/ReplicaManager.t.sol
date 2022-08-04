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
        replicaManager.initialize(remoteDomain, updater, watchtower);
        dApp = new AppHarness(OPTIMISTIC_PERIOD);
        systemMessenger = ISystemMessenger(address(1234567890));
        replicaManager.setSystemMessenger(systemMessenger);
    }

    // ============ INITIAL STATE ============
    function test_correctlyInitialized() public {
        assertEq(uint256(replicaManager.localDomain()), uint256(localDomain));
        assertEq(replicaManager.updater(), updater);
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        replicaManager.initialize(remoteDomain, updater, watchtower);
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

    event Update(
        uint32 indexed homeDomain,
        uint32 indexed nonce,
        bytes32 indexed root,
        bytes signature
    );

    // Relayer relays a new root signed by updater on Home chain
    function test_successfulUpdate() public {
        uint32 nonce = 42;
        assertEq(replicaManager.updater(), vm.addr(updaterPK));
        (bytes memory attestation, bytes memory sig) = signRemoteAttestation(
            updaterPK,
            nonce,
            ROOT
        );
        // Root doesn't exist yet
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, ROOT), 0);
        // Relayer sends over a root signed by the updater on the Home chain
        vm.expectEmit(true, true, true, true);
        emit Update(remoteDomain, nonce, ROOT, sig);
        replicaManager.submitAttestation(attestation);
        // Time at which root was confirmed is set, optimistic timeout starts now
        assertEq(replicaManager.activeReplicaConfirmedAt(remoteDomain, ROOT), block.timestamp);
    }

    function test_updateWithFakeSigner() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(fakeUpdaterPK, nonce, ROOT);
        vm.expectRevert("Signer is not an updater");
        // Update signed by fakeUpdater should be rejected
        replicaManager.submitAttestation(attestation);
    }

    function test_updateWithLocalDomain() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signHomeAttestation(updaterPK, nonce, ROOT);
        vm.expectRevert("Update refers to local chain");
        // Replica should reject updates from the chain it's deployed on
        replicaManager.submitAttestation(attestation);
    }

    event UpdaterBlacklisted(
        address indexed updater,
        address indexed reporter,
        address indexed watchtower,
        bytes report
    );

    function test_submitReport() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(updaterPK, nonce, ROOT);
        bytes memory report = signReport(updaterPK, watchtowerPK, attestation);
        vm.expectEmit(true, true, true, true);
        emit UpdaterBlacklisted(updater, address(this), watchtower, report);
        replicaManager.submitReport(report);
    }

    function test_acceptableRoot() public {
        uint32 optimisticSeconds = 69;
        test_successfulUpdate();
        vm.warp(block.timestamp + optimisticSeconds);
        assertTrue(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, ROOT));
    }

    function test_cannotAcceptableRoot() public {
        test_successfulUpdate();
        uint32 optimisticSeconds = 69;
        vm.warp(block.timestamp + optimisticSeconds - 1);
        assertFalse(replicaManager.acceptableRoot(remoteDomain, optimisticSeconds, ROOT));
    }

    event LogTips(uint96 updaterTip, uint96 relayerTip, uint96 proverTip, uint96 processorTip);

    function test_process() public {
        bytes memory message = _prepareProcessTest(OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD);
        vm.expectEmit(true, true, true, true);
        emit LogTips(UPDATER_TIP, RELAYER_TIP, PROVER_TIP, PROCESSOR_TIP);
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
