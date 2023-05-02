// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, AgentStatus, SystemEntity} from "../../../contracts/libs/Structures.sol";
import {InterfaceDestination} from "../../../contracts/interfaces/InterfaceDestination.sol";
import {InterfaceOrigin} from "../../../contracts/interfaces/InterfaceOrigin.sol";
import {GAS_DATA_LENGTH} from "../../../contracts/libs/Constants.sol";
import {ChainGas, GasDataLib} from "../../../contracts/libs/GasData.sol";

import {AgentManagerTest} from "./AgentManager.t.sol";

import {
    AgentFlag,
    AgentStatus,
    LightManager,
    LightManagerHarness,
    IAgentSecured,
    SynapseTest
} from "../../utils/SynapseTest.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {RawAttestation, RawSnapshot} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract LightManagerTest is AgentManagerTest {
    // Deploy mocks for every messaging contract
    constructor() SynapseTest(0) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        vm.assume(domain != DOMAIN_SYNAPSE);
        address caller = random.nextAddress();
        address origin_ = random.nextAddress();
        address destination_ = random.nextAddress();
        LightManager cleanContract = new LightManager(domain);
        vm.prank(caller);
        cleanContract.initialize(origin_, destination_);
        assertEq(cleanContract.localDomain(), domain);
        assertEq(cleanContract.owner(), caller);
        assertEq(cleanContract.origin(), origin_);
        assertEq(cleanContract.destination(), destination_);
    }

    function initializeLocalContract() public override {
        LightManager(localContract()).initialize(address(0), address(0));
    }

    // ═══════════════════════════════════════════════ TESTS: SETUP ════════════════════════════════════════════════════

    function test_constructor_revert_onSynapseChain() public {
        // Should not be able to deploy on Synapse Chain
        vm.expectRevert("Can't be deployed on SynChain");
        new LightManagerHarness(DOMAIN_SYNAPSE);
    }

    function test_setup() public override {
        super.test_setup();
        assertEq(lightManager.version(), LATEST_VERSION, "!version");
    }

    // ══════════════════════════════════ TESTS: UNAUTHORIZED ACCESS (NOT OWNER) ═══════════════════════════════════════

    function test_setAgentRoot_revert_notDestination(address caller) public {
        vm.assume(caller != destination);
        vm.expectRevert("Only Destination sets agent root");
        vm.prank(caller);
        lightManager.setAgentRoot(bytes32(uint256(1)));
    }

    // ═════════════════════════════════════════ TESTS: ADD/REMOVE AGENTS ══════════════════════════════════════════════

    function test_addAgent_new(address caller, uint32 domain, address agent) public {
        // Should not be an already added agent
        vm.assume(agent != address(0));
        vm.assume(lightManager.agentStatus(agent).flag == AgentFlag.Unknown);
        bytes32 root = addNewAgent(domain, agent);
        test_setAgentRoot(root);
        bytes32[] memory proof = getAgentProof(agent);
        expectStatusUpdated(AgentFlag.Active, domain, agent);
        // Anyone could add agents in Light Manager
        vm.prank(caller);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
        checkAgentStatus(agent, lightManager.agentStatus(agent), AgentFlag.Active);
    }

    function test_updateAgentStatus_slashed(address caller, uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        // Set flag to Slashed in the Merkle Tree
        bytes32 root = updateAgent(AgentFlag.Slashed, agent);
        test_setAgentRoot(root);
        bytes32[] memory proof = getAgentProof(agent);
        expectStatusUpdated(AgentFlag.Slashed, domain, agent);
        expectDisputeResolved(agent, address(0), caller);
        vm.prank(caller);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
        checkAgentStatus(agent, lightManager.agentStatus(agent), AgentFlag.Slashed);
    }

    function test_setAgentRoot(bytes32 root) public {
        bool isDifferent = root != lightManager.agentRoot();
        if (isDifferent) {
            vm.expectEmit();
            emit RootUpdated(root);
        }
        vm.recordLogs();
        vm.prank(destination);
        lightManager.setAgentRoot(root);
        if (!isDifferent) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't have");
        }
        assertEq(lightManager.agentRoot(), root, "!agentRoot");
    }

    function test_setAgentRoot_equal() public {
        test_setAgentRoot(lightManager.agentRoot());
    }

    // ═══════════════════════════════════════ TEST: UPDATE AGENTS (REVERTS) ═══════════════════════════════════════════

    function test_addAgent_revert_invalidProof(uint256 domainId, uint256 agentId) public {
        (, address agent) = getAgent(domainId, agentId);
        bytes32[] memory proof = getAgentProof(agent);
        AgentStatus memory status = getAgentStatus(agent);
        // This succeeds, but doesn't do anything, as agent was already added
        lightManager.updateAgentStatus(agent, status, proof);
        // Change agent root, so old proofs are no longer valid
        test_setAgentRoot(bytes32(0));
        assertEq(uint8(lightManager.agentStatus(agent).flag), uint8(AgentFlag.Unknown));
        vm.expectRevert("Invalid proof");
        lightManager.updateAgentStatus(agent, status, proof);
    }

    // ══════════════════════════════════════════ TEST: SUBMIT STATEMENTS ══════════════════════════════════════════════

    function test_submitAttestation(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        // Should pass the attestation to Destination: acceptAttestation(status, sigIndex, attestation, snapGas)
        vm.expectCall(
            destination,
            abi.encodeWithSelector(
                InterfaceDestination.acceptAttestation.selector,
                agentIndex[notary],
                nextSignatureIndex(),
                attPayload,
                snapGas
            )
        );
        lightManager.submitAttestation(attPayload, attSignature, snapGas);
    }

    function test_submitAttestation_success_snapGasHighBitsMalformed(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        uint256[] memory snapGasMalformed = new uint256[](snapGas.length);
        uint256 chainGasBits = 8 * (4 + GAS_DATA_LENGTH);
        for (uint256 i = 0; i < snapGas.length; i++) {
            // This will not revert as the malformed bit is outside of ChainGas struct: (domain, gasData)
            uint256 malformedBit = chainGasBits + random.nextUint8() % (256 - chainGasBits);
            ChainGas cg0 = GasDataLib.wrapChainGas(snapGas[i]);
            snapGasMalformed[i] = snapGas[i] ^ (1 << malformedBit);
            ChainGas cg1 = GasDataLib.wrapChainGas(snapGas[i]);
            // Malformed bit should not affect ChainGas
            assert(ChainGas.unwrap(cg0) == ChainGas.unwrap(cg1));
        }
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        // Should pass the attestation to Destination: acceptAttestation(status, sigIndex, attestation, snapGas)
        // Note: the malformed highest bits are ignored, so will be passing `snapGas` instead of `snapGasMalformed`
        vm.expectCall(
            destination,
            abi.encodeWithSelector(
                InterfaceDestination.acceptAttestation.selector,
                agentIndex[notary],
                nextSignatureIndex(),
                attPayload,
                snapGas
            )
        );
        // Try to feed the gas data with malformed highest bits
        lightManager.submitAttestation(attPayload, attSignature, snapGasMalformed);
    }

    function test_submitAttestation_revert_snapGasMismatch(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        // This should revert only if the malformed bit is within ChainGas struct: (domain, gasData)
        uint256 chainGasBits = 8 * (4 + GAS_DATA_LENGTH);
        uint256 malformedBit = random.nextUint8() % chainGasBits;
        uint256 malformedIndex = random.nextUint256() % snapGas.length;
        snapGas[malformedIndex] ^= 1 << malformedBit;
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        vm.expectRevert("Invalid snapGas");
        // Try to feed the gas data with malformed lowest bits
        lightManager.submitAttestation(attPayload, attSignature, snapGas);
    }

    function test_submitAttestation_revert_signedByGuard(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        address guard = domains[0].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(guard, ra);
        vm.expectRevert("Signer is not a Notary");
        lightManager.submitAttestation(attPayload, attSignature, snapGas);
    }

    function test_submitAttestation_revert_signedByRemoteNotary(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        vm.expectRevert("Wrong Notary domain");
        lightManager.submitAttestation(attPayload, attSignature, snapGas);
    }

    function test_submitAttestation_revert_notaryInDispute(Random memory random) public {
        RawSnapshot memory rs = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation(rs, random.nextUint32());
        uint256[] memory snapGas = rs.snapGas();
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        openDispute({guard: domains[0].agent, notary: notary});
        vm.expectRevert("Notary is in dispute");
        lightManager.submitAttestation(attPayload, attSignature, snapGas);
    }

    // ════════════════════════════════════════════ TEST: WITHDRAW TIPS ════════════════════════════════════════════════

    function test_remoteWithdrawTips(address actor, uint256 amount, uint32 proofMaturity) public {
        proofMaturity = uint32(bound(proofMaturity, BONDING_OPTIMISTIC_PERIOD, type(uint32).max));
        skip(proofMaturity);
        bytes memory msgPayload = managerMsgPayload(DOMAIN_SYNAPSE, remoteWithdrawTipsCalldata(actor, amount));
        bytes memory expectedCall = abi.encodeWithSelector(InterfaceOrigin.withdrawTips.selector, actor, amount);
        vm.expectCall(origin, expectedCall);
        managerMsgPrank(msgPayload);
    }

    function test_remoteWithdrawTips_revert_notDestination(address caller) public {
        vm.assume(caller != destination);
        skip(BONDING_OPTIMISTIC_PERIOD);
        vm.expectRevert("!destination");
        vm.prank(caller);
        lightManager.remoteWithdrawTips(DOMAIN_SYNAPSE, BONDING_OPTIMISTIC_PERIOD, address(0), 0);
    }

    function test_remoteWithdrawTips_revert_notSynapseChain(uint32 msgOrigin) public {
        vm.assume(msgOrigin != DOMAIN_SYNAPSE);
        skip(BONDING_OPTIMISTIC_PERIOD);
        bytes memory msgPayload = managerMsgPayload(msgOrigin, remoteWithdrawTipsCalldata(address(0), 0));
        vm.expectRevert("!synapseDomain");
        managerMsgPrank(msgPayload);
    }

    function test_remoteWithdrawTips_revert_optimisticPeriodNotOver(uint32 proofMaturity) public {
        proofMaturity = proofMaturity % BONDING_OPTIMISTIC_PERIOD;
        skip(proofMaturity);
        bytes memory msgPayload = managerMsgPayload(DOMAIN_SYNAPSE, remoteWithdrawTipsCalldata(address(0), 0));
        vm.expectRevert("!optimisticPeriod");
        managerMsgPrank(msgPayload);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
