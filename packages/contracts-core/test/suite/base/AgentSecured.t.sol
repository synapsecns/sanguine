// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MessagingBaseTest} from "./MessagingBase.t.sol";
import {AgentFlag} from "../../utils/SynapseTest.t.sol";

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

abstract contract AgentSecuredTest is MessagingBaseTest {
    struct SnapshotMock {
        RawState rs;
        RawStateIndex rsi;
    }

    function expectAgentSlashed(uint32 domain, address agent, address prover) public {
        vm.expectEmit(localAgentManager());
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
        // vm.expectEmit(systemContract());
        // emit AgentSlashed(domain, agent, prover);
    }

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
}
