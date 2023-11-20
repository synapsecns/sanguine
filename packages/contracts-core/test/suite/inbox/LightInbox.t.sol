// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    AgentNotGuard,
    AgentNotNotary,
    IncorrectAgentDomain,
    IncorrectDataHash,
    GuardInDispute,
    NotaryInDispute,
    SynapseDomainForbidden
} from "../../../contracts/libs/Errors.sol";
import {InterfaceDestination} from "../../../contracts/interfaces/InterfaceDestination.sol";

import {StatementInboxTest} from "./StatementInbox.t.sol";

import {BaseMock} from "../../mocks/base/BaseMock.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {RawAttestation, RawSnapshot} from "../../utils/libs/SynapseStructs.t.sol";

import {LightInbox, SynapseTest} from "../../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract LightInboxTest is StatementInboxTest {
    // Deploy mocks for everything except LightManager and LightInbox
    constructor() SynapseTest(0) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        vm.assume(domain != DOMAIN_SYNAPSE);
        vm.chainId(domain);

        address caller = random.nextAddress();
        address agentManager = random.nextAddress();
        address origin_ = random.nextAddress();
        address destination_ = random.nextAddress();
        address owner_ = random.nextAddress();
        LightInbox lightInbox_ = new LightInbox(DOMAIN_SYNAPSE);
        vm.prank(caller);
        lightInbox_.initialize(agentManager, origin_, destination_, owner_);
        assertEq(lightInbox_.owner(), owner_);
        assertEq(lightInbox_.localDomain(), domain);
        assertEq(lightInbox_.origin(), origin_);
        assertEq(lightInbox_.destination(), destination_);
        assertEq(lightInbox_.agentManager(), agentManager);
    }

    function test_setup() public override {
        super.test_setup();
        assertEq(lightInbox.version(), LATEST_VERSION);
    }

    function test_constructor_revert_onSynapseChain() public {
        vm.chainId(DOMAIN_SYNAPSE);
        vm.expectRevert(SynapseDomainForbidden.selector);
        new LightInbox({synapseDomain_: DOMAIN_SYNAPSE});
    }

    function test_constructor_revert_chainIdOverflow() public {
        vm.chainId(2 ** 32);
        vm.expectRevert("SafeCast: value doesn't fit in 32 bits");
        new LightInbox({synapseDomain_: 1});
    }

    function initializeLocalContract() public override {
        LightInbox(lightInbox).initialize(address(0), address(0), address(0), address(0));
    }

    // ══════════════════════════════════════════ TEST: SUBMIT STATEMENTS ══════════════════════════════════════════════

    function test_submitAttestation(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        // Should pass to Destination: acceptAttestation(status, sigIndex, attestation, agentRoot, snapGas)
        vm.expectCall(
            destination,
            abi.encodeWithSelector(
                InterfaceDestination.acceptAttestation.selector,
                agentIndex[notary],
                nextSignatureIndex(),
                attPayload,
                ra._agentRoot,
                snapGas
            )
        );
        lightInbox.submitAttestation(attPayload, attSignature, ra._agentRoot, snapGas);
    }

    function test_submitAttestation_revert_agentRootMismatch(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        uint256 malformedBit = random.nextUint8();
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        vm.expectRevert(IncorrectDataHash.selector);
        // Try to feed the agent root with a single malformed bit
        lightInbox.submitAttestation(attPayload, attSignature, ra._agentRoot ^ bytes32(1 << malformedBit), snapGas);
    }

    function test_submitAttestation_revert_snapGasMismatch(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        uint256 malformedBit = random.nextUint8();
        uint256 malformedIndex = random.nextUint256() % snapGas.length;
        snapGas[malformedIndex] ^= 1 << malformedBit;
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        vm.expectRevert(IncorrectDataHash.selector);
        // Try to feed the gas data with a single malformed bit
        lightInbox.submitAttestation(attPayload, attSignature, ra._agentRoot, snapGas);
    }

    function test_submitAttestation_revert_signedByGuard(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        address guard = domains[0].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(guard, ra);
        vm.expectRevert(AgentNotNotary.selector);
        lightInbox.submitAttestation(attPayload, attSignature, ra._agentRoot, snapGas);
    }

    function test_submitAttestation_revert_signedByRemoteNotary(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        vm.expectRevert(IncorrectAgentDomain.selector);
        lightInbox.submitAttestation(attPayload, attSignature, ra._agentRoot, snapGas);
    }

    // ════════════════════════════════════════════ TEST: OPEN DISPUTES ════════════════════════════════════════════════

    function test_submitAttestationReport(Random memory random) public {
        address prover = makeAddr("Prover");
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (, bytes memory attSignature) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory attPayload, bytes memory arSignature) = signAttestationReport(guard, ra);
        expectDisputeOpened(0, guard, notary);
        vm.prank(prover);
        lightInbox.submitAttestationReport(attPayload, arSignature, attSignature);
        assertEq(lightInbox.getReportsAmount(), 1, "!reportsAmount");
        (bytes memory reportPayload, bytes memory reportSignature) = lightInbox.getGuardReport(0);
        assertEq(reportPayload, attPayload, "!reportPayload");
        assertEq(reportSignature, arSignature, "!reportSig");
    }

    function test_submitAttestationReport_revert_signedByNotary(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (, bytes memory attSignature) = signAttestation(notary, ra);
        // Force a random Notary to sign the report
        address reportSigner = getNotary(random.nextUint256(), random.nextUint256());
        (bytes memory attPayload, bytes memory arSignature) = signAttestationReport(reportSigner, ra);
        vm.expectRevert(AgentNotGuard.selector);
        lightInbox.submitAttestationReport(attPayload, arSignature, attSignature);
    }

    function test_submitAttestationReport_revert_guardInDispute(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agents[0];
        (, bytes memory attSignature) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory attPayload, bytes memory arSignature) = signAttestationReport(guard, ra);
        // Put the Guard in Dispute with another Notary
        openDispute({guard: guard, notary: domains[DOMAIN_LOCAL].agents[1]});
        vm.expectRevert(GuardInDispute.selector);
        lightInbox.submitAttestationReport(attPayload, arSignature, attSignature);
    }

    function test_submitAttestationReport_revert_notaryInDispute(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agents[0];
        (, bytes memory attSignature) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agents[0];
        (bytes memory attPayload, bytes memory arSignature) = signAttestationReport(guard, ra);
        // Put the Notary in Dispute with another Guard
        openDispute({guard: domains[0].agents[1], notary: notary});
        vm.expectRevert(NotaryInDispute.selector);
        lightInbox.submitAttestationReport(attPayload, arSignature, attSignature);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
