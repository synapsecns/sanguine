// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, Dispute, DisputeFlag} from "../../../contracts/libs/Structures.sol";
import {IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";
import {AgentManagerHarness} from "../../harnesses/manager/AgentManagerHarness.t.sol";
import {SynapseTest} from "../../utils/SynapseTest.t.sol";

import {fakeSnapshot} from "../../utils/libs/FakeIt.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {
    AttestationFlag,
    StateFlag,
    RawAttestation,
    RawAttestationReport,
    RawSnapshot,
    RawState,
    RawStateIndex,
    RawStateReport
} from "../../utils/libs/SynapseStructs.t.sol";

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

    function expectDisputeOpened(address guard, address notary) public {
        vm.expectEmit();
        emit DisputeUpdated(guard, Dispute(DisputeFlag.Pending, agentIndex[notary], address(0)));
        vm.expectEmit();
        emit DisputeUpdated(notary, Dispute(DisputeFlag.Pending, agentIndex[guard], address(0)));
        // TODO: check if Summit is called, when separated
        bytes memory expectedCall =
            abi.encodeWithSelector(IAgentSecured.openDispute.selector, agentIndex[guard], agentIndex[notary]);
        vm.expectCall(localDestination(), expectedCall);
    }

    function expectDisputeResolved(address slashed, address honest, address prover) public {
        vm.expectEmit();
        emit DisputeUpdated(slashed, Dispute(DisputeFlag.Slashed, agentIndex[honest], prover));
        if (honest != address(0)) {
            vm.expectEmit();
            emit DisputeUpdated(honest, Dispute(DisputeFlag.None, 0, address(0)));
        }
        // TODO: check if Summit is called, when separated
        bytes memory expectedCall =
            abi.encodeWithSelector(IAgentSecured.resolveDispute.selector, agentIndex[slashed], agentIndex[honest]);
        vm.expectCall(localDestination(), expectedCall);
    }

    // ══════════════════════════════════════════════ DISPUTE CHEATS ═══════════════════════════════════════════════════

    function openDispute(address guard, address notary) public {
        AgentManagerHarness(localAgentManager()).openDisputeExposed(guard, notary);
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

    function createSignedAttestationReport(address guard, RawAttestation memory ra)
        public
        view
        returns (bytes memory arPayload, bytes memory arSig)
    {
        RawAttestationReport memory rawAR = RawAttestationReport(uint8(AttestationFlag.Invalid), ra);
        return signAttestationReport(guard, rawAR);
    }

    function createSignedStateReport(address guard, RawState memory rs)
        public
        view
        returns (bytes memory srPayload, bytes memory srSig)
    {
        RawStateReport memory rawSR = RawStateReport(uint8(StateFlag.Invalid), rs);
        return signStateReport(guard, rawSR);
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
