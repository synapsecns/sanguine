// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentNotGuard, GuardInDispute, NotaryInDispute} from "../../../contracts/libs/Errors.sol";
import {IAgentManager} from "../../../contracts/interfaces/IAgentManager.sol";
import {IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";
import {AgentFlag, AgentStatus, SystemEntity} from "../../../contracts/libs/Structures.sol";

import {MessagingBaseTest} from "../base/MessagingBase.t.sol";
import {AgentManagerHarness} from "../../harnesses/manager/AgentManagerHarness.t.sol";

import {fakeSnapshot} from "../../utils/libs/FakeIt.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {
    RawAttestation,
    RawCallData,
    RawManagerCall,
    RawSnapshot,
    RawState,
    RawStateIndex
} from "../../utils/libs/SynapseStructs.t.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract AgentManagerTest is MessagingBaseTest {
    using Address for address;

    uint256 internal rootSubmittedAt;
    uint256 private _signatureIndex;

    function test_setup() public virtual {
        assertEq(address(testedAM().destination()), localDestination());
        assertEq(address(testedAM().origin()), localOrigin());
        assertEq(testedAM().agentRoot(), getAgentRoot());
    }

    // ═══════════════════════════════════════════════ TESTS: VIEWS ════════════════════════════════════════════════════

    function test_getAgent_notExistingIndex() public {
        (address agent, AgentStatus memory status) = testedAM().getAgent(0);
        assertEq(agent, address(0));
        assertEq(uint8(status.flag), 0);
        assertEq(status.domain, 0);
        assertEq(status.index, 0);
        // Last agent has index DOMAIN_AGENTS * allDomains.length
        (agent, status) = testedAM().getAgent(DOMAIN_AGENTS * allDomains.length + 1);
        assertEq(agent, address(0));
        assertEq(uint8(status.flag), 0);
        assertEq(status.domain, 0);
        assertEq(status.index, 0);
    }

    // ════════════════════════════════════════════ TESTS: OPEN DISPUTE ════════════════════════════════════════════════

    function test_submitStateReportWithSnapshot(RawState memory rs, RawStateIndex memory rsi) public boundIndex(rsi) {
        address prover = makeAddr("Prover");
        // Create Notary signature for the snapshot
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        expectDisputeOpened(guard, notary);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshot(
            rsi.stateIndex, srPayload, srSig, snapPayload, snapSig
        );
    }

    function test_submitStateReportWithSnapshot_revert_signedByNotary(Random memory random) public {
        RawState memory rs = random.nextState();
        RawStateIndex memory rsi = random.nextStateIndex();
        // Create Notary signature for the snapshot
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        // Force a random Notary to sign the report
        address reportSigner = getNotary(random.nextUint256(), random.nextUint256());
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(reportSigner, rs);
        vm.expectRevert(AgentNotGuard.selector);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshot(
            rsi.stateIndex, srPayload, srSig, snapPayload, snapSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        expectDisputeOpened(guard, notary);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithAttestation(
            rsi.stateIndex, srPayload, srSig, snapPayload, attPayload, attSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(reportSigner, rs);
        vm.expectRevert(AgentNotGuard.selector);
        IAgentManager(localAgentManager()).submitStateReportWithAttestation(
            rsi.stateIndex, srPayload, srSig, snapPayload, attPayload, attSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        expectDisputeOpened(guard, notary);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshotProof(
            rsi.stateIndex, srPayload, srSig, snapProof, attPayload, attSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(reportSigner, rs);
        // Generate Snapshot Proof
        acceptSnapshot(rawSnap);
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        vm.expectRevert(AgentNotGuard.selector);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshotProof(
            rsi.stateIndex, srPayload, srSig, snapProof, attPayload, attSig
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        openDispute(guard, domains[DOMAIN_LOCAL].agents[1]);
        vm.expectRevert(GuardInDispute.selector);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshot(
            rsi.stateIndex, srPayload, srSig, snapPayload, snapSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        openDispute(domains[0].agents[1], notary);
        vm.expectRevert(NotaryInDispute.selector);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshot(
            rsi.stateIndex, srPayload, srSig, snapPayload, snapSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        openDispute(guard, domains[DOMAIN_LOCAL].agents[1]);
        vm.expectRevert(GuardInDispute.selector);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithAttestation(
            rsi.stateIndex, srPayload, srSig, snapPayload, attPayload, attSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        openDispute(domains[0].agents[1], notary);
        vm.expectRevert(NotaryInDispute.selector);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithAttestation(
            rsi.stateIndex, srPayload, srSig, snapPayload, attPayload, attSig
        );
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        openDispute(guard, domains[DOMAIN_LOCAL].agents[1]);
        vm.expectRevert(GuardInDispute.selector);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshotProof(
            rsi.stateIndex, srPayload, srSig, snapProof, attPayload, attSig
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
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        openDispute(domains[0].agents[1], notary);
        vm.expectRevert(NotaryInDispute.selector);
        vm.prank(prover);
        IAgentManager(localAgentManager()).submitStateReportWithSnapshotProof(
            rsi.stateIndex, srPayload, srSig, snapProof, attPayload, attSig
        );
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function nextSignatureIndex() public returns (uint256 sigIndex) {
        sigIndex = _signatureIndex++;
    }

    function checkAgentStatus(address agent, AgentStatus memory status, AgentFlag flag) public virtual override {
        super.checkAgentStatus(agent, status, flag);
        (address agent_, AgentStatus memory status_) = testedAM().getAgent(status.index);
        assertEq(agent_, agent, "!agent");
        super.checkAgentStatus(agent, status_, flag);
    }

    function skipBondingOptimisticPeriod() public {
        skipPeriod(BONDING_OPTIMISTIC_PERIOD);
    }

    function skipPeriod(uint256 period) public {
        rootSubmittedAt = block.timestamp;
        skip(period);
    }

    function managerMsgPrank(bytes memory payload) public {
        vm.prank(localDestination());
        localContract().functionCall(payload);
    }

    function managerMsgPayload(uint32 msgOrigin, RawCallData memory rcd) public view returns (bytes memory) {
        RawManagerCall memory rmc =
            RawManagerCall({origin: msgOrigin, proofMaturity: block.timestamp - rootSubmittedAt, callData: rcd});
        return rmc.callPayload();
    }

    function remoteSlashAgentCalldata(uint32 domain, address agent, address prover)
        public
        view
        returns (RawCallData memory)
    {
        // (msgOrigin, proofMaturity) are omitted => (domain, agent, prover)
        return
            RawCallData({selector: bondingManager.remoteSlashAgent.selector, args: abi.encode(domain, agent, prover)});
    }

    function remoteWithdrawTipsCalldata(address actor, uint256 amount) public view returns (RawCallData memory) {
        // (msgOrigin, proofMaturity) are omitted => (actor, amount)
        return RawCallData({selector: lightManager.remoteWithdrawTips.selector, args: abi.encode(actor, amount)});
    }

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return localAgentManager();
    }

    /// @notice Returns tested contract as AgentManager
    function testedAM() public view returns (AgentManagerHarness) {
        return AgentManagerHarness(localAgentManager());
    }
}
