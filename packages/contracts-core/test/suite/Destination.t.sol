// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AttestationLib, SummitAttestation } from "../../contracts/libs/Attestation.sol";
import { SNAPSHOT_MAX_STATES } from "../../contracts/libs/Constants.sol";
import { Snapshot, SnapshotLib } from "../../contracts/libs/Snapshot.sol";
import { State, SummitState } from "../../contracts/libs/State.sol";
import { AgentInfo, SystemEntity } from "../../contracts/libs/Structures.sol";
import { IAgentRegistry } from "../../contracts/interfaces/IAgentRegistry.sol";

import { InterfaceDestination, ORIGIN_TREE_DEPTH } from "../../contracts/Destination.sol";
import { Versioned } from "../../contracts/Version.sol";

import { MessageRecipientMock } from "../mocks/client/MessageRecipientMock.t.sol";

import { fakeStates } from "../utils/libs/FakeIt.t.sol";
import { RawHeader, RawMessage, RawTips } from "../utils/libs/SynapseStructs.t.sol";
import { addressToBytes32 } from "../utils/libs/SynapseUtilities.t.sol";
import { SynapseProofs } from "../utils/SynapseProofs.t.sol";
import { ISystemContract, SynapseTest } from "../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract DestinationTest is SynapseTest, SynapseProofs {
    using SnapshotLib for bytes;

    uint32 internal constant PERIOD = 1 minutes;
    bytes internal constant BODY = "Test Body";

    RawMessage[] internal rawMessages;
    bytes[] internal messages;
    SummitAttestation internal summitAtt;

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
        // Check Agents
        // Destination should know about local Notaries and Guards
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                if (domain == 0) {
                    assertTrue(IAgentRegistry(destination).isActiveAgent(domain, agent), "!guard");
                } else if (domain == DOMAIN_LOCAL) {
                    assertTrue(
                        IAgentRegistry(destination).isActiveAgent(domain, agent),
                        "!local notary"
                    );
                } else {
                    // Remote Notaries are unknown to Destination
                    assertFalse(
                        IAgentRegistry(destination).isActiveAgent(domain, agent),
                        "!remote notary"
                    );
                }
            }
        }
        // Check version
        assertEq(Versioned(destination).version(), LATEST_VERSION, "!version");
    }

    function test_slashAgent() public {
        address notary = domains[DOMAIN_LOCAL].agent;
        vm.expectEmit(true, true, true, true);
        emit AgentRemoved(DOMAIN_LOCAL, notary);
        vm.expectEmit(true, true, true, true);
        emit AgentSlashed(DOMAIN_LOCAL, notary);
        vm.recordLogs();
        vm.prank(address(systemRouter));
        ISystemContract(destination).slashAgent({
            _rootSubmittedAt: block.timestamp,
            _callOrigin: DOMAIN_LOCAL,
            _caller: SystemEntity.BondingManager,
            _info: AgentInfo(DOMAIN_LOCAL, notary, false)
        });
        assertEq(vm.getRecordedLogs().length, 2, "Emitted extra logs");
    }

    function test_execute(
        SummitState memory state,
        uint256 statesAmount,
        uint256 stateIndex,
        uint32 attNonce,
        uint16 skipTime
    ) public {
        address executor = makeAddr("Executor");

        statesAmount = bound(statesAmount, 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(stateIndex, 0, statesAmount - 1);

        createMessages();
        state.root = getRoot(MESSAGES);
        state.origin = DOMAIN_REMOTE;
        // Remainder of State struct is fuzzed
        createAttestation(state, statesAmount, stateIndex);
        bytes32[] memory snapProof = genSnapshotProof(stateIndex);

        // Attestation Nonce is fuzzed as well
        bytes memory attPayload = AttestationLib.formatSummitAttestation(summitAtt, attNonce);
        bytes memory attSignature = signMessage(domains[DOMAIN_LOCAL].agent, keccak256(attPayload));

        skip(skipTime);
        uint256 rootTimestamp = block.timestamp;
        // Should emit event when attestation is accepted
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(
            DOMAIN_LOCAL,
            domains[DOMAIN_LOCAL].agent,
            attPayload,
            attSignature
        );
        InterfaceDestination(destination).submitAttestation(attPayload, attSignature);
        skip(PERIOD);
        for (uint256 i = 0; i < MESSAGES; ++i) {
            bytes32[ORIGIN_TREE_DEPTH] memory originProof = getLatestProof(i);
            // (_origin, _nonce, _sender, _rootTimestamp, _message)
            vm.expectCall(
                recipient,
                abi.encodeWithSelector(
                    MessageRecipientMock.handle.selector,
                    DOMAIN_REMOTE,
                    i + 1,
                    sender,
                    rootTimestamp,
                    BODY
                )
            );
            // Should emit event when message is executed
            vm.expectEmit(true, true, true, true);
            emit Executed(DOMAIN_REMOTE, keccak256(messages[i]));
            vm.prank(executor);
            InterfaceDestination(destination).execute(
                messages[i],
                originProof,
                snapProof,
                stateIndex
            );
        }
    }

    function createAttestation(
        SummitState memory state,
        uint256 statesAmount,
        uint256 stateIndex
    ) public {
        (bytes[] memory states, State[] memory ptrs) = fakeStates(state, statesAmount, stateIndex);
        Snapshot snapshot = SnapshotLib.formatSnapshot(ptrs).castToSnapshot();
        acceptSnapshot(states);
        summitAtt = snapshot.toSummitAttestation();
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
            (bytes memory message, ) = rm.castToMessage();
            rawMessages.push(rm);
            messages.push(message);
            insertMessage(message);
        }
    }
}
