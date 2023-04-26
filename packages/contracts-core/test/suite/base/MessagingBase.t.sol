// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

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
        bytes[] memory states = rawSnap.formatStates();
        acceptSnapshot(states);
        // Reuse existing metadata in RawAttestation
        return rawSnap.castToRawAttestation(ra.agentRoot, ra.nonce, ra.blockNumber, ra.timestamp);
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

    /// @notice Returns local domain for the tested system contract
    function localDomain() public view virtual returns (uint32);

    /// @notice Returns address of the tested system contract
    function systemContract() public view virtual returns (address);

    /// @notice Returns address of Agent Manager on the tested domain
    function localAgentManager() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(lightManager) : address(bondingManager);
    }

    /// @notice Returns address of Destination on the tested domain
    function localDestination() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(destination) : address(summit);
    }

    /// @notice Returns address of Origin on the tested domain
    function localOrigin() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(origin) : address(originSynapse);
    }

    /// @notice Checks if contract is a local SystemContract
    function isLocalSystemContract(address addr) public view returns (bool) {
        return addr == localAgentManager() || addr == localDestination() || addr == localOrigin();
    }

    /// @notice Checks if contract is a local SystemRegistry
    function isLocalSystemRegistry(address addr) public view returns (bool) {
        return addr == localDestination() || addr == localOrigin();
    }
}
