// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IDisputeHub} from "../../../contracts/interfaces/IDisputeHub.sol";
import {DisputeFlag, DisputeStatus, SystemEntity} from "../../../contracts/libs/Structures.sol";

import {fakeSnapshot} from "../../utils/libs/FakeIt.t.sol";
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
import {SystemRegistryTest} from "../system/SystemRegistry.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract DisputeHubTest is SystemRegistryTest {
    /// @notice Prevents this contract from being included in the coverage report
    function testDisputeHub() external {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         SUBMIT DATA HELPERS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function check_submitStateReport(address hub, uint32 notaryDomain, RawState memory rs, RawStateIndex memory rsi)
        public
    {
        address prover = makeAddr("Prover");
        // Create Notary signature for the snapshot
        address notary = domains[notaryDomain].agent;
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        vm.expectEmit();
        emit Dispute(guard, notaryDomain, notary);
        vm.prank(prover);
        IDisputeHub(hub).submitStateReport(rsi.stateIndex, srPayload, srSig, snapPayload, snapSig);
        checkDisputeOpened(hub, guard, notary);
    }

    function check_submitStateReportWithProof(
        address hub,
        uint32 notaryDomain,
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public {
        address prover = makeAddr("Prover");
        ra = createAttestation(rs, ra, rsi);
        // Create Notary signature for the attestation
        address notary = domains[notaryDomain].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, rs);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        vm.expectEmit();
        emit Dispute(guard, notaryDomain, notary);
        vm.prank(prover);
        IDisputeHub(hub).submitStateReportWithProof(rsi.stateIndex, srPayload, srSig, snapProof, attPayload, attSig);
        checkDisputeOpened(hub, guard, notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CREATE DATA HELPERS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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

    /// @notice Checks that the Dispute was opened between a Guard and a Notary.
    function checkDisputeOpened(address hub, address guard, address notary) public {
        DisputeStatus memory guardStatus = IDisputeHub(hub).disputeStatus(guard);
        assertEq(uint8(guardStatus.flag), uint8(DisputeFlag.Pending), "!guard flag");
        assertEq(guardStatus.counterpart, notary, "!guard counterpart");
        DisputeStatus memory notaryStatus = IDisputeHub(hub).disputeStatus(notary);
        assertEq(uint8(notaryStatus.flag), uint8(DisputeFlag.Pending), "!notary flag");
        assertEq(notaryStatus.counterpart, guard, "!notary counterpart");
    }

    /// @notice Checks that the Dispute between a Guard and a Notary was resolved.
    function checkDisputeResolved(address hub, address honest, address slashed) public {
        DisputeStatus memory honestStatus = IDisputeHub(hub).disputeStatus(honest);
        assertEq(uint8(honestStatus.flag), uint8(DisputeFlag.None), "!honest flag");
        assertEq(honestStatus.counterpart, address(0), "!honest counterpart");
        DisputeStatus memory slashedStatus = IDisputeHub(hub).disputeStatus(slashed);
        assertEq(uint8(slashedStatus.flag), uint8(DisputeFlag.Slashed), "!honest flag");
        assertEq(slashedStatus.counterpart, honest, "!honest counterpart");
    }

    // ═════════════════════════════════════════════════ OVERRIDES ═════════════════════════════════════════════════════

    /// @notice Returns address of the tested system contract
    function systemContract() public view override returns (address) {
        return localDestination();
    }
}
