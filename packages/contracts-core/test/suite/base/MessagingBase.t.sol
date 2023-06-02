// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MessagingBase} from "../../../contracts/base/MessagingBase.sol";
import {AgentFlag, DisputeFlag} from "../../../contracts/libs/Structures.sol";
import {IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";
import {IStatementInbox} from "../../../contracts/interfaces/IStatementInbox.sol";

import {AgentManagerHarness} from "../../harnesses/manager/AgentManagerHarness.t.sol";
import {SynapseTest} from "../../utils/SynapseTest.t.sol";

import {fakeSnapshot} from "../../utils/libs/FakeIt.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {RawAttestation, RawSnapshot, RawState, RawStateIndex} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
abstract contract MessagingBaseTest is SynapseTest {
    struct SnapshotMock {
        RawState rs;
        RawStateIndex rsi;
    }

    modifier boundIndex(RawStateIndex memory rsi) {
        rsi.boundStateIndex();
        _;
    }

    modifier onlySupportedDomain() virtual {
        require(localDomain() == DOMAIN_LOCAL || localDomain() == DOMAIN_SYNAPSE, "Unsupported local domain");
        _;
    }

    // To ensure that every MessagingBase contract has this test
    function test_cleanSetup(Random memory random) public virtual;

    // To ensure that every MessagingBase contract has this test
    function test_initializer_revert_alreadyInitialized() public {
        expectRevertAlreadyInitialized();
        initializeLocalContract();
    }

    function initializeLocalContract() public virtual;

    // ═══════════════════════════════════════════════ EXPECTATIONS ════════════════════════════════════════════════════

    function expectStatusUpdated(AgentFlag flag, uint32 domain, address agent) public {
        vm.expectEmit();
        emit StatusUpdated(flag, domain, agent);
    }

    function expectDisputeOpened(uint256 disputeIndex, address guard, address notary) public {
        vm.expectEmit();
        emit DisputeOpened(disputeIndex, agentIndex[guard], agentIndex[notary]);
        bytes memory expectedCall =
            abi.encodeWithSelector(IAgentSecured.openDispute.selector, agentIndex[guard], agentIndex[notary]);
        vm.expectCall(localDestination(), expectedCall);
        if (localSummit() != address(0)) {
            vm.expectCall(localSummit(), expectedCall);
        }
    }

    function expectDisputeResolved(uint256 disputePtr, address slashed, address honest, address prover) public {
        if (disputePtr != 0) {
            vm.expectEmit();
            emit DisputeResolved(disputePtr - 1, agentIndex[slashed], agentIndex[honest], prover);
        }
        bytes memory expectedCall =
            abi.encodeWithSelector(IAgentSecured.resolveDispute.selector, agentIndex[slashed], agentIndex[honest]);
        vm.expectCall(localDestination(), expectedCall);
        if (localSummit() != address(0)) {
            vm.expectCall(localSummit(), expectedCall);
        }
    }

    // ══════════════════════════════════════════════ DISPUTE CHEATS ═══════════════════════════════════════════════════

    function openDispute(address guard, address notary) public {
        require(agentIndex[guard] != 0 && agentDomain[guard] == 0, "Invalid Guard");
        require(agentIndex[notary] != 0 && agentDomain[notary] != 0, "Invalid Notary");
        RawSnapshot memory rs = fakeSnapshot({statesAmount: 1});
        (bytes memory snapPayload, bytes memory snapSignature) = signSnapshot(notary, rs);
        (, bytes memory srSignature) = signStateReport(guard, rs.states[0]);
        IStatementInbox(localInbox()).submitStateReportWithSnapshot(0, srSignature, snapPayload, snapSignature);
    }

    // ═══════════════════════════════════════════════ AGENT GETTERS ═══════════════════════════════════════════════════

    function randomAgent(Random memory random) public view returns (address agent) {
        // Pick a random Agent
        (, agent) = getAgent(random.nextUint256(), random.nextUint256());
    }

    function randomGuard(Random memory random) public view returns (address guard) {
        // Pick a random Guard
        guard = getGuard(random.nextUint256());
    }

    function randomNotary(Random memory random) public view returns (address notary) {
        if (localDomain() != DOMAIN_SYNAPSE) {
            // Pick a random Notary from local domain
            notary = getDomainAgent(localDomain(), random.nextUint256());
        } else {
            // Pick a random Notary from any domain
            notary = getNotary(random.nextUint256(), random.nextUint256());
        }
    }

    // ═══════════════════════════════════════════════ DATA CREATION ═══════════════════════════════════════════════════

    /// @notice Creates attestation for snapshot having given rawState at given index,
    /// with some fake data for other states in the snapshots.
    function createAttestation(RawState memory rawState, RawAttestation memory ra, RawStateIndex memory rsi)
        public
        returns (RawAttestation memory)
    {
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, rsi);
        acceptSnapshot(rawSnap);
        // Reuse existing metadata in RawAttestation
        return rawSnap.castToRawAttestation(ra._agentRoot, ra.nonce, ra.blockNumber, ra.timestamp);
    }

    function createSnapshotProof(SnapshotMock memory sm)
        public
        returns (RawAttestation memory ra, bytes32[] memory snapProof)
    {
        ra = Random(sm.rs.root).nextAttestation(1);
        ra = createAttestation(sm.rs, ra, sm.rsi);
        snapProof = genSnapshotProof(sm.rsi.stateIndex);
    }

    function createSignedSnapshot(address notary, RawState memory rs, RawStateIndex memory rsi)
        public
        view
        returns (bytes memory snapPayload, bytes memory snapSig)
    {
        RawSnapshot memory rawSnap = fakeSnapshot(rs, rsi);
        return signSnapshot(notary, rawSnap);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public view virtual returns (uint32);

    /// @notice Returns address of the tested contract
    function localContract() public view virtual returns (address);

    /// @notice Returns address of Agent Manager on the tested domain
    function localAgentManager() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(lightManager) : address(bondingManager);
    }

    /// @notice Returns address of Inbox on the tested domain
    function localInbox() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(lightInbox) : address(inbox);
    }

    /// @notice Returns address of Destination on the tested domain
    function localDestination() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(destination) : address(destinationSynapse);
    }

    /// @notice Returns address of Origin on the tested domain
    function localOrigin() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(origin) : address(originSynapse);
    }

    /// @notice Returns address of Summit on the tested domain
    function localSummit() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_SYNAPSE ? address(summit) : address(0);
    }

    /// @notice Checks if contract is a local SystemContract
    function isLocalSystemContract(address addr) public view returns (bool) {
        return addr == localAgentManager() || addr == localDestination() || addr == localOrigin();
    }

    /// @notice Checks if contract is a local AgentSecured contract
    function isLocalAgentSecured(address addr) public view returns (bool) {
        return addr == localDestination() || addr == localOrigin();
    }
}
