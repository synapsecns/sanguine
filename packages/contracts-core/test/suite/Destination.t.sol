// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SNAPSHOT_MAX_STATES } from "../../contracts/libs/Snapshot.sol";
import { SystemEntity } from "../../contracts/libs/Structures.sol";
import { ISystemRegistry } from "../../contracts/interfaces/ISystemRegistry.sol";
import { IDisputeHub } from "../../contracts/interfaces/IDisputeHub.sol";
import { IExecutionHub } from "../../contracts/interfaces/IExecutionHub.sol";

import { InterfaceDestination } from "../../contracts/Destination.sol";
import { Versioned } from "../../contracts/Version.sol";

import { MessageRecipientMock } from "../mocks/client/MessageRecipientMock.t.sol";

import { fakeSnapshot } from "../utils/libs/FakeIt.t.sol";
import {
    AttestationFlag,
    StateFlag,
    RawAttestation,
    RawAttestationReport,
    RawHeader,
    RawMessage,
    RawSnapshot,
    RawState,
    RawStateReport,
    RawTips
} from "../utils/libs/SynapseStructs.t.sol";
import { addressToBytes32 } from "../utils/libs/SynapseUtilities.t.sol";
import { ISystemContract, SynapseTest } from "../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract DestinationTest is SynapseTest {
    uint32 internal constant PERIOD = 1 minutes;
    bytes internal constant BODY = "Test Body";

    RawMessage[] internal rawMessages;
    bytes[] internal messages;

    address internal sender;
    address internal recipient;

    // Deploy Production version of Destination and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_DESTINATION) {}

    function setUp() public override {
        super.setUp();
        sender = makeAddr("Sender");
        recipient = address(new MessageRecipientMock());
    }

    function test_setupCorrectly() public {
        // Check Messaging addresses
        assertEq(
            address(ISystemContract(destination).systemRouter()),
            address(systemRouter),
            "!systemRouter"
        );
        // TODO: adjust when Agent Merkle Tree is implemented
        // Check Agents: currently all Agents are known in LightManager
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                assertTrue(ISystemRegistry(destination).isActiveAgent(domain, agent));
            }
        }
        // Check version
        assertEq(Versioned(destination).version(), LATEST_VERSION, "!version");
    }

    function test_submitAttestationReport(RawAttestationReport memory rawAR) public {
        address reporter = makeAddr("Reporter");
        // Make sure Flag fits in AttestationFlag enum
        rawAR.flag = uint8(bound(rawAR.flag, 0, uint8(type(AttestationFlag).max)));
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (, bytes memory attSig) = signAttestation(notary, rawAR.attestation);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory arPayload, bytes memory arSig) = signAttestationReport(guard, rawAR);
        // TODO: complete the test when Dispute is implemented
        vm.expectEmit(true, true, true, true);
        emit Dispute(guard, DOMAIN_LOCAL, notary);
        vm.prank(reporter);
        InterfaceDestination(destination).submitAttestationReport(arPayload, arSig, attSig);
    }

    function test_submitStateReport(
        RawState memory rs,
        RawAttestation memory ra,
        uint256 statesAmount,
        uint256 stateIndex
    ) public {
        address reporter = makeAddr("Reporter");
        // Make sure statesAmount, stateIndex are valid entires
        statesAmount = bound(statesAmount, 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(stateIndex, 0, statesAmount - 1);
        ra = createAttestation(rs, ra, statesAmount, stateIndex);
        // Create Notary signature for the snapshot
        address notary = domains[DOMAIN_LOCAL].agent;
        // fakeSnapshot is deterministic, so this will be the same snapshot
        RawSnapshot memory rawSnap = fakeSnapshot(rs, statesAmount, stateIndex);
        (bytes memory snapPayload, bytes memory snapSig) = signSnapshot(notary, rawSnap);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        RawStateReport memory rawSR = RawStateReport(uint8(StateFlag.Invalid), rs);
        (bytes memory srPayload, bytes memory srSig) = signStateReport(guard, rawSR);
        // TODO: complete the test when Dispute is implemented
        vm.expectEmit(true, true, true, true);
        emit Dispute(guard, DOMAIN_LOCAL, notary);
        vm.prank(reporter);
        IDisputeHub(destination).submitStateReport(
            stateIndex,
            srPayload,
            srSig,
            snapPayload,
            snapSig
        );
    }

    function test_submitStateReportWithProof(
        RawState memory rs,
        RawAttestation memory ra,
        uint256 statesAmount,
        uint256 stateIndex
    ) public {
        address reporter = makeAddr("Reporter");
        // Make sure statesAmount, stateIndex are valid entires
        statesAmount = bound(statesAmount, 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(stateIndex, 0, statesAmount - 1);
        ra = createAttestation(rs, ra, statesAmount, stateIndex);
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        RawStateReport memory rawSR = RawStateReport(uint8(StateFlag.Invalid), rs);
        (bytes memory srPayload, bytes memory srSig) = signStateReport(guard, rawSR);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(stateIndex);
        // TODO: complete the test when Dispute is implemented
        vm.expectEmit(true, true, true, true);
        emit Dispute(guard, DOMAIN_LOCAL, notary);
        vm.prank(reporter);
        IDisputeHub(destination).submitStateReportWithProof(
            stateIndex,
            srPayload,
            srSig,
            snapProof,
            attPayload,
            attSig
        );
    }

    function test_execute(
        RawState memory rs,
        RawAttestation memory ra,
        uint256 statesAmount,
        uint256 stateIndex,
        uint32 rootSubmittedAt
    ) public {
        address executor = makeAddr("Executor");

        statesAmount = bound(statesAmount, 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(stateIndex, 0, statesAmount - 1);

        createMessages();
        rs.root = getRoot(MESSAGES);
        rs.origin = DOMAIN_REMOTE;
        // Remainder of State struct is fuzzed
        ra = createAttestation(rs, ra, statesAmount, stateIndex);
        bytes32[] memory snapProof = genSnapshotProof(stateIndex);

        // Attestation Nonce is fuzzed as well
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);

        vm.warp(rootSubmittedAt);
        // Should emit event when attestation is accepted
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(DOMAIN_LOCAL, notary, attPayload, attSig);
        InterfaceDestination(destination).submitAttestation(attPayload, attSig);
        skip(PERIOD);
        for (uint256 i = 0; i < MESSAGES; ++i) {
            bytes32[] memory originProof = getLatestProof(i);
            // (_origin, _nonce, _sender, _rootSubmittedAt, _message)
            vm.expectCall(
                recipient,
                abi.encodeWithSelector(
                    MessageRecipientMock.handle.selector,
                    DOMAIN_REMOTE,
                    i + 1,
                    sender,
                    rootSubmittedAt,
                    BODY
                )
            );
            // Should emit event when message is executed
            vm.expectEmit(true, true, true, true);

            emit Executed(DOMAIN_REMOTE, keccak256(messages[i]));
            vm.prank(executor);
            IExecutionHub(destination).execute(messages[i], originProof, snapProof, stateIndex);
        }
    }

    function createAttestation(
        RawState memory rawState,
        RawAttestation memory ra,
        uint256 statesAmount,
        uint256 stateIndex
    ) public returns (RawAttestation memory) {
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, statesAmount, stateIndex);
        bytes[] memory states = rawSnap.formatStates();
        acceptSnapshot(states);
        // Reuse existing metadata in RawAttestation
        return rawSnap.castToRawAttestation(ra.nonce, ra.blockNumber, ra.timestamp);
    }

    function createMessages() public {
        for (uint32 i = 0; i < MESSAGES; ++i) {
            RawMessage memory rm = RawMessage(
                RawHeader({
                    origin: DOMAIN_REMOTE,
                    sender: addressToBytes32(sender),
                    nonce: i + 1,
                    destination: DOMAIN_LOCAL,
                    recipient: addressToBytes32(recipient),
                    optimisticSeconds: PERIOD
                }),
                RawTips(0, 0, 0, 0),
                BODY
            );
            bytes memory message = rm.formatMessage();
            rawMessages.push(rm);
            messages.push(message);
            insertMessage(message);
        }
    }
}
