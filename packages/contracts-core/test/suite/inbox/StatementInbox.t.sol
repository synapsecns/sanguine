// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {StatementInbox} from "../../../contracts/inbox/StatementInbox.sol";
import {AgentNotGuard, GuardInDispute, NotaryInDispute} from "../../../contracts/libs/Errors.sol";
import {AgentFlag} from "../../../contracts/libs/Structures.sol";
import {IExecutionHub} from "../../../contracts/interfaces/IExecutionHub.sol";

import {MessagingBaseTest} from "../base/MessagingBase.t.sol";

import {fakeSnapshot} from "../../utils/libs/FakeIt.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {
    RawAttestation, RawExecReceipt, RawSnapshot, RawState, RawStateIndex
} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract StatementInboxTest is MessagingBaseTest {
    uint256 private _signatureIndex;

    function test_setup() public virtual {
        assertEq(testedInbox().owner(), address(this));
        assertEq(testedInbox().localDomain(), localDomain());
        assertEq(testedInbox().origin(), localOrigin());
        assertEq(testedInbox().destination(), localDestination());
        assertEq(testedInbox().agentManager(), localAgentManager());
    }

    // ════════════════════════════════════════════ TESTS: OPEN DISPUTE ════════════════════════════════════════════════

    function test_submitStateReportWithSnapshot(RawState memory rs, RawStateIndex memory rsi) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        // Create Notary signature for the snapshot
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory statePayload, bytes memory srSig) = signStateReport(guard, rs);
        expectDisputeOpened(0, guard, notary);
        vm.prank(prover);
        testedInbox().submitStateReportWithSnapshot(rsi.stateIndex, srSig, snapPayload, snapSig);
        assertEq(testedInbox().getReportsAmount(), 1, "!reportsAmount");
        (bytes memory reportPayload, bytes memory reportSignature) = testedInbox().getGuardReport(0);
        assertEq(reportPayload, statePayload, "!reportPayload");
        assertEq(reportSignature, srSig, "!reportSignature");
    }

    function test_submitStateReportWithSnapshot_revert_signedByNotary(Random memory random) public {
        RawState memory rs = random.nextState();
        RawStateIndex memory rsi = random.nextStateIndex();
        // Create Notary signature for the snapshot
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        // Force a random Notary to sign the report
        address reportSigner = getNotary(random.nextUint256(), random.nextUint256());
        (, bytes memory srSig) = signStateReport(reportSigner, rs);
        vm.expectRevert(AgentNotGuard.selector);
        testedInbox().submitStateReportWithSnapshot(rsi.stateIndex, srSig, snapPayload, snapSig);
    }

    function test_submitStateReportWithAttestation(
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        ra = createAttestation(rs, ra, rsi);
        bytes memory snapPayload = fakeSnapshot(rs, rsi).formatSnapshot();
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory statePayload, bytes memory srSig) = signStateReport(guard, rs);
        expectDisputeOpened(0, guard, notary);
        vm.prank(prover);
        testedInbox().submitStateReportWithAttestation(rsi.stateIndex, srSig, snapPayload, attPayload, attSig);
        assertEq(testedInbox().getReportsAmount(), 1, "!reportsAmount");
        (bytes memory reportPayload, bytes memory reportSignature) = testedInbox().getGuardReport(0);
        assertEq(reportPayload, statePayload, "!reportPayload");
        assertEq(reportSignature, srSig, "!reportSignature");
    }

    function test_submitStateReportWithAttestation_revert_signedByNotary(Random memory random) public {
        RawState memory rs = random.nextState();
        RawStateIndex memory rsi = random.nextStateIndex();
        RawSnapshot memory rawSnap = fakeSnapshot(rs, rsi);
        bytes memory snapPayload = rawSnap.formatSnapshot();
        RawAttestation memory ra = random.nextAttestation(rawSnap, random.nextUint32());
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Force a random Notary to sign the report
        address reportSigner = getNotary(random.nextUint256(), random.nextUint256());
        (, bytes memory srSig) = signStateReport(reportSigner, rs);
        vm.expectRevert(AgentNotGuard.selector);
        testedInbox().submitStateReportWithAttestation(rsi.stateIndex, srSig, snapPayload, attPayload, attSig);
    }

    function test_submitStateReportWithSnapshotProof(
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        ra = createAttestation(rs, ra, rsi);
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory statePayload, bytes memory srSig) = signStateReport(guard, rs);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        expectDisputeOpened(0, guard, notary);
        vm.prank(prover);
        testedInbox().submitStateReportWithSnapshotProof(
            rsi.stateIndex, statePayload, srSig, snapProof, attPayload, attSig
        );
        assertEq(testedInbox().getReportsAmount(), 1, "!reportsAmount");
        (bytes memory reportPayload, bytes memory reportSignature) = testedInbox().getGuardReport(0);
        assertEq(reportPayload, statePayload, "!reportPayload");
        assertEq(reportSignature, srSig, "!reportSignature");
    }

    function test_submitStateReportWithSnapshotProof_revert_signedByNotary(Random memory random) public {
        RawState memory rs = random.nextState();
        RawStateIndex memory rsi = random.nextStateIndex();
        RawSnapshot memory rawSnap = fakeSnapshot(rs, rsi);
        RawAttestation memory ra = random.nextAttestation(rawSnap, random.nextUint32());
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Force a random Notary to sign the report
        address reportSigner = getNotary(random.nextUint256(), random.nextUint256());
        (bytes memory statePayload, bytes memory srSig) = signStateReport(reportSigner, rs);
        // Generate Snapshot Proof
        acceptSnapshot(rawSnap);
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        vm.expectRevert(AgentNotGuard.selector);
        testedInbox().submitStateReportWithSnapshotProof(
            rsi.stateIndex, statePayload, srSig, snapProof, attPayload, attSig
        );
    }

    // ═════════════════════════════════════════ TESTS: ALREADY IN DISPUTE ═════════════════════════════════════════════

    function test_submitStateReportWithSnapshot_revert_guardInDispute(RawState memory rs, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        address prover = makeAddr("Prover");
        // Create Notary signature for the snapshot
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (, bytes memory srSig) = signStateReport(guard, rs);
        openDispute(guard, domains[DOMAIN_LOCAL].agents[1]);
        vm.expectRevert(GuardInDispute.selector);
        vm.prank(prover);
        testedInbox().submitStateReportWithSnapshot(rsi.stateIndex, srSig, snapPayload, snapSig);
    }

    function test_submitStateReportWithSnapshot_revert_notaryInDispute(RawState memory rs, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        address prover = makeAddr("Prover");
        // Create Notary signature for the snapshot
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (, bytes memory srSig) = signStateReport(guard, rs);
        openDispute(domains[0].agents[1], notary);
        vm.expectRevert(NotaryInDispute.selector);
        vm.prank(prover);
        testedInbox().submitStateReportWithSnapshot(rsi.stateIndex, srSig, snapPayload, snapSig);
    }

    function test_submitStateReportWithAttestation_revert_guardInDispute(
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        ra = createAttestation(rs, ra, rsi);
        bytes memory snapPayload = fakeSnapshot(rs, rsi).formatSnapshot();
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (, bytes memory srSig) = signStateReport(guard, rs);
        openDispute(guard, domains[DOMAIN_LOCAL].agents[1]);
        vm.expectRevert(GuardInDispute.selector);
        vm.prank(prover);
        testedInbox().submitStateReportWithAttestation(rsi.stateIndex, srSig, snapPayload, attPayload, attSig);
    }

    function test_submitStateReportWithAttestation_revert_notaryInDispute(
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        ra = createAttestation(rs, ra, rsi);
        bytes memory snapPayload = fakeSnapshot(rs, rsi).formatSnapshot();
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (, bytes memory srSig) = signStateReport(guard, rs);
        openDispute(domains[0].agents[1], notary);
        vm.expectRevert(NotaryInDispute.selector);
        vm.prank(prover);
        testedInbox().submitStateReportWithAttestation(rsi.stateIndex, srSig, snapPayload, attPayload, attSig);
    }

    function test_submitStateReportWithSnapshotProof_revert_guardInDispute(
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        ra = createAttestation(rs, ra, rsi);
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory statePayload, bytes memory srSig) = signStateReport(guard, rs);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        openDispute(guard, domains[DOMAIN_LOCAL].agents[1]);
        vm.expectRevert(GuardInDispute.selector);
        vm.prank(prover);
        testedInbox().submitStateReportWithSnapshotProof(
            rsi.stateIndex, statePayload, srSig, snapProof, attPayload, attSig
        );
    }

    function test_submitStateReportWithSnapshotProof_revert_notaryInDispute(
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        ra = createAttestation(rs, ra, rsi);
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory statePayload, bytes memory srSig) = signStateReport(guard, rs);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        openDispute(domains[0].agents[1], notary);
        vm.expectRevert(NotaryInDispute.selector);
        vm.prank(prover);
        testedInbox().submitStateReportWithSnapshotProof(
            rsi.stateIndex, statePayload, srSig, snapProof, attPayload, attSig
        );
    }

    // ═════════════════════════════════════════ TESTS: VERIFY STATEMENTS ══════════════════════════════════════════════

    function test_verifyReceiptReport_validReceipt(Random memory random) public {
        address guard = randomGuard(random);
        RawExecReceipt memory re = random.nextReceipt(localDomain());
        bytes memory rcptPayload = re.formatReceipt();
        bytes memory rrSignature = signReceiptReport(guard, rcptPayload);
        // Force isValidReceipt(rcptPayload) to return true
        vm.mockCall(
            localDestination(),
            abi.encodeWithSelector(IExecutionHub.isValidReceipt.selector, rcptPayload),
            abi.encode(true)
        );
        vm.expectEmit();
        emit InvalidReceiptReport(rcptPayload, rrSignature);
        expectStatusUpdated(AgentFlag.Fraudulent, 0, guard);
        expectDisputeResolved(0, guard, address(0), address(this));
        // Should return false, as the report is invalid
        assertFalse(testedInbox().verifyReceiptReport(rcptPayload, rrSignature));
    }

    function test_verifyReceiptReport_invalidReceipt(Random memory random) public {
        address guard = randomGuard(random);
        RawExecReceipt memory re = random.nextReceipt(localDomain());
        bytes memory rcptPayload = re.formatReceipt();
        bytes memory rrSignature = signReceiptReport(guard, rcptPayload);
        // Force isValidReceipt(rcptPayload) to return false
        vm.mockCall(
            localDestination(),
            abi.encodeWithSelector(IExecutionHub.isValidReceipt.selector, rcptPayload),
            abi.encode(false)
        );
        vm.recordLogs();
        // Should return true, as the report is valid
        assertTrue(testedInbox().verifyReceiptReport(rcptPayload, rrSignature));
        assertEq(vm.getRecordedLogs().length, 0, "Logs should be empty");
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function nextSignatureIndex() public returns (uint256 sigIndex) {
        sigIndex = _signatureIndex++;
    }

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return localInbox();
    }

    /// @notice Returns tested contract as StatementInbox
    function testedInbox() public view returns (StatementInbox) {
        return StatementInbox(localInbox());
    }
}
