// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { TypedMemView } from "../contracts/libs/TypedMemView.sol";
import { TypeCasts } from "../contracts/libs/TypeCasts.sol";

import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";

import { MirrorLib } from "../contracts/libs/Mirror.sol";
import { ISystemRouter } from "../contracts/interfaces/ISystemRouter.sol";
import { DestinationHarness } from "./harnesses/DestinationHarness.sol";

import { AppHarness } from "./harnesses/AppHarness.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";

// solhint-disable func-name-mixedcase
contract DestinationTest is SynapseTest {
    DestinationHarness destination;
    AppHarness dApp;

    uint32 internal constant OPTIMISTIC_PERIOD = 10;

    bytes32 internal constant ROOT = keccak256("test root");

    uint256 processGas;
    uint256 reserveGas;

    ISystemRouter internal systemRouter;

    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

    event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt);

    function setUp() public override {
        super.setUp();
        destination = new DestinationHarness(localDomain);
        destination.initialize(remoteDomain, notary);
        dApp = new AppHarness(OPTIMISTIC_PERIOD);
        systemRouter = ISystemRouter(address(1234567890));
        destination.setSystemRouter(systemRouter);
        destination.addGuard(guard);
    }

    // ============ INITIAL STATE ============
    function test_correctlyInitialized() public {
        assertEq(uint256(destination.localDomain()), uint256(localDomain));
        assertTrue(destination.isNotary(remoteDomain, notary));
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        destination.initialize(remoteDomain, notary);
    }

    // ============ STATE & PERMISSIONING ============

    function test_cannotSetNotaryAsNotOwner(address _notOwner, address _notary) public {
        vm.assume(_notOwner != destination.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        destination.setNotary(remoteDomain, _notary);
    }

    function test_setNotary(address _notary) public {
        vm.assume(_notary != notary);
        vm.prank(destination.owner());
        destination.setNotary(remoteDomain, _notary);
        assertTrue(destination.isNotary(remoteDomain, _notary));
    }

    function test_cannotSetConfirmationAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != destination.owner());
        vm.prank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        destination.setConfirmation(remoteDomain, ROOT, 0);
    }

    event SetConfirmation(
        uint32 indexed remoteDomain,
        bytes32 indexed root,
        uint256 previousConfirmAt,
        uint256 newConfirmAt
    );

    function test_setConfirmation(uint256 _confirmAt) public {
        vm.startPrank(destination.owner());
        assertEq(destination.activeMirrorConfirmedAt(remoteDomain, ROOT), 0);
        vm.expectEmit(true, true, true, true);
        emit SetConfirmation(remoteDomain, ROOT, 0, _confirmAt);
        destination.setConfirmation(remoteDomain, ROOT, _confirmAt);
        assertEq(destination.activeMirrorConfirmedAt(remoteDomain, ROOT), _confirmAt);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SUBMIT REPORT                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event NotaryBlacklisted(
        address indexed notary,
        address indexed guard,
        address indexed reporter,
        bytes report
    );

    function test_submitReport() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, ROOT);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        vm.expectEmit(true, true, true, true);
        emit NotaryBlacklisted(notary, guard, address(this), report);
        assertTrue(destination.submitReport(report));
    }

    function test_submitReport_valid() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, ROOT);
        (bytes memory report, ) = signValidReport(guardPK, attestation);
        vm.expectRevert("Not a fraud report");
        destination.submitReport(report);
    }

    function test_submitReport_notGuard() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, ROOT);
        (bytes memory report, ) = signFraudReport(fakeGuardPK, attestation);
        vm.expectRevert("Signer is not a guard");
        destination.submitReport(report);
    }

    function test_submitReport_notNotary() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(fakeNotaryPK, nonce, ROOT);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        vm.expectRevert("Signer is not a notary");
        destination.submitReport(report);
    }

    function test_submitReport_twice() public {
        test_submitReport();
        uint32 nonce = 69;
        bytes32 root = "another fraud attestation";
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, root);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        // Reporting already blacklisted Notary will lead to reverting,
        // as Notary is blacklisted
        vm.expectRevert("Signer is not a notary");
        destination.submitReport(report);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SUBMIT ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event AttestationAccepted(
        uint32 indexed origin,
        uint32 indexed nonce,
        bytes32 indexed root,
        bytes signature
    );

    // Broadcaster relays a new root signed by notary on Origin chain
    function test_submitAttestation() public {
        uint32 nonce = 42;
        assertTrue(destination.isNotary(remoteDomain, vm.addr(notaryPK)));
        (bytes memory attestation, bytes memory sig) = signRemoteAttestation(notaryPK, nonce, ROOT);
        // Root doesn't exist yet
        assertEq(destination.activeMirrorConfirmedAt(remoteDomain, ROOT), 0);
        // Broadcaster sends over a root signed by the notary on the Origin chain
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(remoteDomain, nonce, ROOT, sig);
        destination.submitAttestation(attestation);
        // Time at which root was confirmed is set, optimistic timeout starts now
        assertEq(destination.activeMirrorConfirmedAt(remoteDomain, ROOT), block.timestamp);
    }

    function test_submitAttestation_fakeNotary() public {
        uint32 nonce = 42;
        (bytes memory attestation, ) = signRemoteAttestation(fakeNotaryPK, nonce, ROOT);
        vm.expectRevert("Signer is not a notary");
        // Attestation signed by fakeNotary should be rejected
        destination.submitAttestation(attestation);
    }

    function test_submitAttestation_localDomain() public {
        // Make Notary active on localDomain
        destination.removeNotary(remoteDomain, notary);
        destination.addNotary(localDomain, notary);
        uint32 nonce = 42;
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, ROOT);
        vm.expectRevert("Attestation refers to local chain");
        // Mirror should reject attestations from the chain it's deployed on
        destination.submitAttestation(attestation);
    }

    function test_acceptableRoot() public {
        uint32 optimisticSeconds = 69;
        test_submitAttestation();
        vm.warp(block.timestamp + optimisticSeconds);
        assertTrue(destination.acceptableRoot(remoteDomain, optimisticSeconds, ROOT));
    }

    function test_cannotAcceptableRoot() public {
        test_submitAttestation();
        uint32 optimisticSeconds = 69;
        vm.warp(block.timestamp + optimisticSeconds - 1);
        assertFalse(destination.acceptableRoot(remoteDomain, optimisticSeconds, ROOT));
    }

    event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip);

    function test_execute() public {
        bytes memory message = _prepareExecuteTest(OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD);
        vm.expectEmit(true, true, true, true);
        emit LogTips(NOTARY_TIP, BROADCASTER_TIP, PROVER_TIP, EXECUTOR_TIP);
        destination.execute(message);
    }

    function test_executePeriodNotPassed() public {
        bytes memory message = _prepareExecuteTest(OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("!optimisticSeconds");
        destination.execute(message);
    }

    function test_executeForgedPeriodReduced() public {
        bytes memory message = _prepareExecuteTest(OPTIMISTIC_PERIOD - 1);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("app: !optimisticSeconds");
        destination.execute(message);
    }

    function test_executeForgePeriodZero() public {
        bytes memory message = _prepareExecuteTest(0);
        vm.expectRevert("app: !optimisticSeconds");
        destination.execute(message);
    }

    function test_onlySystemRouter() public {
        vm.expectEmit(true, true, true, true);
        emit LogSystemCall(1, 2, 3);
        vm.prank(address(systemRouter));
        destination.setSensitiveValue(1337, 1, 2, 3);
        assertEq(destination.sensitiveValue(), 1337);
    }

    function test_onlySystemRouter_rejectOthers() public {
        vm.expectRevert("!systemRouter");
        destination.setSensitiveValue(1337, 0, 0, 0);
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
        destination.setMessageStatus(remoteDomain, messageHash, ROOT);
    }
}
