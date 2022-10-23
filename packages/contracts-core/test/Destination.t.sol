// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "forge-std/Test.sol";

import { TypedMemView } from "../contracts/libs/TypedMemView.sol";
import { TypeCasts } from "../contracts/libs/TypeCasts.sol";

import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";

import { ISystemRouter } from "../contracts/interfaces/ISystemRouter.sol";
import { DestinationHarness } from "./harnesses/DestinationHarness.sol";

import { AppHarness } from "./harnesses/AppHarness.sol";

import { SynapseTest } from "./utils/SynapseTest.sol";
import { ProofGenerator } from "./utils/ProofGenerator.sol";

// solhint-disable func-name-mixedcase
contract DestinationTest is SynapseTest {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

    DestinationHarness internal destination;
    AppHarness internal dApp;

    uint32 internal constant OPTIMISTIC_PERIOD = 10;

    uint32 internal constant NONCE = 14;
    bytes32 internal constant ROOT = keccak256("test root");

    ProofGenerator internal proofGen;
    bytes internal testMessage;
    bytes32 internal leaf;
    uint32 internal messageIndex;
    bytes32[32] internal proof;
    bytes32 internal merkleRoot;

    ISystemRouter internal systemRouter;

    event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt);
    event SetConfirmation(
        uint32 indexed remoteDomain,
        bytes32 indexed root,
        uint256 previousConfirmAt,
        uint256 newConfirmAt
    );
    event NotaryBlacklisted(
        address indexed notary,
        address indexed guard,
        address indexed reporter,
        bytes report
    );
    event AttestationAccepted(
        uint32 indexed origin,
        uint32 indexed nonce,
        bytes32 indexed root,
        bytes signature
    );
    event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip);

    function setUp() public override {
        super.setUp();
        destination = new DestinationHarness(localDomain);
        destination.initialize();
        dApp = new AppHarness(OPTIMISTIC_PERIOD);
        systemRouter = ISystemRouter(address(1234567890));
        destination.setSystemRouter(systemRouter);
        destination.addNotary(remoteDomain, notary);
        destination.addGuard(guard);

        proofGen = new ProofGenerator();
        merkleRoot = ROOT;
    }

    // ============ INITIAL STATE ============
    function test_correctlyInitialized() public {
        assertEq(uint256(destination.localDomain()), uint256(localDomain));
        assertTrue(destination.isNotary(remoteDomain, notary));
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        destination.initialize();
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

    function test_setConfirmation(uint96 _confirmAt) public {
        vm.startPrank(destination.owner());
        assertEq(destination.submittedAt(remoteDomain, ROOT), 0);
        vm.expectEmit(true, true, true, true);
        emit SetConfirmation(remoteDomain, ROOT, 0, _confirmAt);
        destination.setConfirmation(remoteDomain, ROOT, _confirmAt);
        assertEq(destination.submittedAt(remoteDomain, ROOT), _confirmAt);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SUBMIT REPORT                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitReport() public {
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, ROOT);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        vm.expectEmit(true, true, true, true);
        emit NotaryBlacklisted(notary, guard, address(this), report);
        assertTrue(destination.submitReport(report));
    }

    function test_submitReport_valid() public {
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, ROOT);
        (bytes memory report, ) = signValidReport(guardPK, attestation);
        vm.expectRevert("Not a fraud report");
        destination.submitReport(report);
    }

    function test_submitReport_notGuard() public {
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, ROOT);
        (bytes memory report, ) = signFraudReport(fakeGuardPK, attestation);
        vm.expectRevert("Signer is not a guard");
        destination.submitReport(report);
    }

    function test_submitReport_notNotary() public {
        (bytes memory attestation, ) = signRemoteAttestation(fakeNotaryPK, NONCE, ROOT);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        vm.expectRevert("Signer is not a notary");
        destination.submitReport(report);
    }

    function test_submitReport_twice() public {
        test_submitReport();
        uint32 nonce = NONCE + 1;
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

    // Broadcaster relays a new root signed by notary on Origin chain
    function test_submitAttestation() public {
        assertTrue(destination.isNotary(remoteDomain, vm.addr(notaryPK)));
        (bytes memory attestation, bytes memory sig) = signRemoteAttestation(
            notaryPK,
            NONCE,
            merkleRoot
        );
        // Root doesn't exist yet
        assertEq(destination.submittedAt(remoteDomain, merkleRoot), 0);
        // Broadcaster sends over a root signed by the notary on the Origin chain
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(remoteDomain, NONCE, merkleRoot, sig);
        destination.submitAttestation(attestation);
        // Time at which root was confirmed is set, optimistic timeout starts now
        assertEq(destination.submittedAt(remoteDomain, merkleRoot), block.timestamp);
    }

    function test_submitAttestation_fakeNotary() public {
        (bytes memory attestation, ) = signRemoteAttestation(fakeNotaryPK, NONCE, ROOT);
        vm.expectRevert("Signer is not a notary");
        // Attestation signed by fakeNotary should be rejected
        destination.submitAttestation(attestation);
    }

    function test_submitAttestation_emptyRoot() public {
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, bytes32(0));
        vm.expectRevert("Empty root");
        // Attestations with empty root should be rejected
        destination.submitAttestation(attestation);
    }

    function test_submitAttestation_localDomain() public {
        // Make Notary active on localDomain
        destination.removeNotary(remoteDomain, notary);
        destination.addNotary(localDomain, notary);
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, NONCE, ROOT);
        vm.expectRevert("Attestation is from local chain");
        // Mirror should reject attestations from the chain it's deployed on
        destination.submitAttestation(attestation);
    }

    function test_acceptableRoot() public {
        test_submitAttestation();
        skip(OPTIMISTIC_PERIOD);
        assertTrue(destination.acceptableRoot(remoteDomain, OPTIMISTIC_PERIOD, merkleRoot));
    }

    function test_acceptableRoot_invalidRoot() public {
        vm.expectRevert("Invalid root");
        skip(OPTIMISTIC_PERIOD);
        destination.acceptableRoot(remoteDomain, OPTIMISTIC_PERIOD, merkleRoot);
    }

    function test_acceptableRoot_inactiveNotary() public {
        test_submitAttestation();
        destination.removeNotary(remoteDomain, notary);
        skip(OPTIMISTIC_PERIOD);
        vm.expectRevert("Inactive notary");
        destination.acceptableRoot(remoteDomain, OPTIMISTIC_PERIOD, merkleRoot);
    }

    function test_acceptableRoot_periodNoPassed() public {
        test_submitAttestation();
        uint32 optimisticSeconds = 69;
        vm.warp(block.timestamp + optimisticSeconds - 1);
        vm.expectRevert("!optimisticSeconds");
        assertFalse(destination.acceptableRoot(remoteDomain, optimisticSeconds, merkleRoot));
    }

    function test_execute_firstIndex() public {
        _checkExecute(0);
    }

    function test_execute_midIndex() public {
        _checkExecute(NONCE / 2);
    }

    function test_execute_lastIndex() public {
        _checkExecute(NONCE - 1);
    }

    function test_execute_periodNotPassed() public {
        _prepareTest(0, OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("!optimisticSeconds");
        destination.execute(testMessage, proof, 0);
    }

    function test_execute_forgedPeriod_reduced() public {
        _prepareTest(0, OPTIMISTIC_PERIOD - 1);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD - 1);
        vm.expectRevert("app: !optimisticSeconds");
        destination.execute(testMessage, proof, 0);
    }

    function test_executeForgePeriodZero() public {
        _prepareTest(0, 0);
        vm.expectRevert("app: !optimisticSeconds");
        destination.execute(testMessage, proof, 0);
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

    function _checkExecute(uint32 _index) internal {
        _prepareTest(_index, OPTIMISTIC_PERIOD);
        vm.warp(block.timestamp + OPTIMISTIC_PERIOD);
        vm.expectEmit(true, true, true, true);
        emit LogTips(NOTARY_TIP, BROADCASTER_TIP, PROVER_TIP, EXECUTOR_TIP);
        destination.execute(testMessage, proof, _index);
        // Check that merkle root used for proving was recorded
        assertEq(destination.messageStatus(remoteDomain, leaf), merkleRoot, "!messageStatus");
    }

    function _prepareTest(uint32 _messageIndex, uint32 _optimisticPeriod) internal {
        bytes32 sender = "sender";
        bytes memory messageBody = "message body";
        dApp.prepare(remoteDomain, _messageIndex, sender, messageBody);
        bytes32 recipient = TypeCasts.addressToBytes32(address(dApp));
        bytes memory _header = Header.formatHeader(
            remoteDomain,
            sender,
            _messageIndex,
            localDomain,
            recipient,
            _optimisticPeriod
        );
        testMessage = Message.formatMessage(_header, getDefaultTips(), messageBody);
        leaf = keccak256(testMessage);

        bytes32[] memory leafs = new bytes32[](NONCE);
        for (uint256 i = 0; i < NONCE; ++i) {
            if (i == _messageIndex) {
                leafs[i] = leaf;
            } else {
                leafs[i] = keccak256(abi.encode(NONCE, i));
            }
        }
        proofGen.createTree(leafs);
        merkleRoot = proofGen.getRoot();
        proof = proofGen.getProof(_messageIndex);
        test_submitAttestation();
    }
}
