// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAgentRegistry } from "../../contracts/interfaces/IAgentRegistry.sol";
import { IStateHub } from "../../contracts/interfaces/IStateHub.sol";
import { EMPTY_ROOT, SNAPSHOT_MAX_STATES } from "../../contracts/libs/Constants.sol";
import { AgentInfo, SystemEntity } from "../../contracts/libs/Structures.sol";
import { TipsLib } from "../../contracts/libs/Tips.sol";

import { InterfaceOrigin } from "../../contracts/Origin.sol";
import { Versioned } from "../../contracts/Version.sol";

import { OriginStateMask } from "./libs/State.t.sol";
import { fakeState, fakeSnapshot } from "../utils/libs/FakeIt.t.sol";
import { Random } from "../utils/libs/Random.t.sol";
import {
    RawAttestation,
    RawHeader,
    RawMessage,
    RawSnapshot,
    RawState,
    RawTips
} from "../utils/libs/SynapseStructs.t.sol";
import { addressToBytes32 } from "../utils/libs/SynapseUtilities.t.sol";
import { SynapseProofs } from "../utils/SynapseProofs.t.sol";
import { ISystemContract, SynapseTest } from "../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract OriginTest is SynapseTest, SynapseProofs {
    // Deploy Production version of Origin and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_ORIGIN) {}

    function test_setupCorrectly() public {
        // Check Messaging addresses
        assertEq(
            address(ISystemContract(origin).systemRouter()),
            address(systemRouter),
            "!systemRouter"
        );
        // Check Agents
        // Origin should know about agents from all domains, including Guards
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                assertTrue(IAgentRegistry(origin).isActiveAgent(domain, agent), "!agent");
            }
        }
        // Check version
        assertEq(Versioned(origin).version(), LATEST_VERSION, "!version");
    }

    function test_dispatch() public {
        address sender = makeAddr("Sender");
        address recipient = makeAddr("Recipient");
        uint32 period = 1 minutes;
        bytes memory tips = TipsLib.emptyTips();
        bytes memory body = "test body";

        RawMessage[] memory rawMessages = new RawMessage[](MESSAGES);
        bytes[] memory messages = new bytes[](MESSAGES);
        bytes32[] memory roots = new bytes32[](MESSAGES);
        for (uint32 i = 0; i < MESSAGES; ++i) {
            rawMessages[i] = RawMessage(
                RawHeader({
                    origin: DOMAIN_LOCAL,
                    sender: addressToBytes32(sender),
                    nonce: i + 1,
                    destination: DOMAIN_REMOTE,
                    recipient: addressToBytes32(recipient),
                    optimisticSeconds: period
                }),
                RawTips(0, 0, 0, 0),
                body
            );
            (messages[i], ) = rawMessages[i].castToMessage();
            insertMessage(messages[i]);
            roots[i] = getRoot(i + 1);
        }

        // Expect Origin Events
        for (uint32 i = 0; i < MESSAGES; ++i) {
            // 1 block is skipped after each dispatched message
            RawState memory rs = RawState({
                root: roots[i],
                origin: DOMAIN_LOCAL,
                nonce: i + 1,
                blockNumber: uint40(block.number + i),
                timestamp: uint40(block.timestamp + i * BLOCK_TIME)
            });
            (bytes memory state, ) = rs.castToState();
            vm.expectEmit(true, true, true, true);
            emit StateSaved(state);
            vm.expectEmit(true, true, true, true);
            emit Dispatched(keccak256(messages[i]), i + 1, DOMAIN_REMOTE, messages[i]);
        }

        for (uint32 i = 0; i < MESSAGES; ++i) {
            vm.prank(sender);
            (uint32 messageNonce, bytes32 messageHash) = InterfaceOrigin(origin).dispatch(
                DOMAIN_REMOTE,
                addressToBytes32(recipient),
                period,
                tips,
                body
            );
            // Check return values
            assertEq(messageNonce, i + 1, "!messageNonce");
            assertEq(messageHash, keccak256(messages[i]), "!messageHash");
            skipBlock();
        }
    }

    function test_states() public {
        IStateHub hub = IStateHub(origin);
        // Check initial States
        assertEq(hub.statesAmount(), 1, "!initial statesAmount");
        // Initial state was saved "1 block ago"
        RawState memory rs = RawState({
            root: EMPTY_ROOT,
            origin: DOMAIN_LOCAL,
            nonce: 0,
            blockNumber: uint40(block.number - 1),
            timestamp: uint40(block.timestamp - BLOCK_TIME)
        });
        (bytes memory state, ) = rs.castToState();
        assertEq(hub.suggestState(0), state, "!state: 0");
        assertEq(hub.suggestState(0), hub.suggestLatestState(), "!latest state: 0");
        // Dispatch some messages
        test_dispatch();
        // Check saved States
        assertEq(hub.statesAmount(), MESSAGES + 1, "!statesAmount");
        assertEq(hub.suggestState(0), state, "!suggestState: 0");
        for (uint32 i = 0; i < MESSAGES; ++i) {
            rs.nonce += 1;
            rs.root = getRoot(rs.nonce);
            rs.blockNumber += 1;
            rs.timestamp += uint40(BLOCK_TIME);
            (state, ) = rs.castToState();
            assertEq(hub.suggestState(i + 1), state, "!suggestState");
        }
        assertEq(hub.suggestLatestState(), state, "!suggestLatestState");
    }

    function test_slashAgent() public {
        address notary = domains[DOMAIN_REMOTE].agent;
        vm.expectEmit(true, true, true, true);
        emit AgentRemoved(DOMAIN_REMOTE, notary);
        vm.expectEmit(true, true, true, true);
        emit AgentSlashed(DOMAIN_REMOTE, notary);
        vm.recordLogs();
        vm.prank(address(systemRouter));
        ISystemContract(origin).slashAgent({
            _rootSubmittedAt: block.timestamp,
            _callOrigin: DOMAIN_LOCAL,
            _caller: SystemEntity.BondingManager,
            _info: AgentInfo(DOMAIN_REMOTE, notary, false)
        });
        assertEq(vm.getRecordedLogs().length, 2, "Emitted extra logs");
    }

    function test_verifySnapshot_existingNonce(
        uint32 nonce,
        OriginStateMask memory mask,
        uint256 statesAmount,
        uint256 stateIndex
    ) public {
        (bool isValid, RawState memory rs) = _prepareExistingState(nonce, mask);
        _verifySnapshot(rs, isValid, statesAmount, stateIndex);
    }

    function test_verifySnapshot_unknownNonce(
        RawState memory rs,
        uint256 statesAmount,
        uint256 stateIndex
    ) public {
        // Restrict nonce to non-existing ones
        rs.nonce = uint32(bound(rs.nonce, MESSAGES + 1, type(uint32).max));
        rs.origin = DOMAIN_LOCAL;
        // Remaining fields are fuzzed
        _verifySnapshot(rs, false, statesAmount, stateIndex);
    }

    function test_verifyAttestation_existingNonce(
        Random memory random,
        uint32 nonce,
        OriginStateMask memory mask
    ) public {
        (bool isValid, RawState memory rs) = _prepareExistingState(nonce, mask);
        _verifyAttestation(random, rs, isValid);
    }

    function test_verifyAttestation_unknownNonce(Random memory random, RawState memory rs) public {
        // Restrict nonce to non-existing ones
        rs.nonce = uint32(bound(rs.nonce, MESSAGES + 1, type(uint32).max));
        rs.origin = DOMAIN_LOCAL;
        // Remaining fields are fuzzed
        _verifyAttestation(random, rs, false);
    }

    function _prepareExistingState(uint32 nonce, OriginStateMask memory mask)
        internal
        returns (bool isValid, RawState memory rs)
    {
        uint40 initialBN = uint40(block.number - 1);
        uint40 initialTS = uint40(block.timestamp - BLOCK_TIME);
        test_dispatch();
        // State is valid if and only if all three fields match
        isValid = !(mask.diffRoot || mask.diffBlockNumber || mask.diffTimestamp);
        // Restrict nonce to existing ones
        nonce = uint32(bound(nonce, 0, MESSAGES));
        rs = RawState({
            root: getRoot(nonce),
            origin: DOMAIN_LOCAL,
            nonce: nonce,
            blockNumber: initialBN + nonce,
            timestamp: uint40(initialTS + nonce * BLOCK_TIME)
        });
        if (mask.diffRoot) rs.root = rs.root ^ bytes32(uint256(1));
        if (mask.diffBlockNumber) rs.blockNumber = rs.blockNumber ^ 1;
        if (mask.diffTimestamp) rs.timestamp = rs.timestamp ^ 1;
    }

    function _verifyAttestation(
        Random memory random,
        RawState memory rawState,
        bool isValid
    ) internal {
        // Pick random domain expect for 0
        uint256 domainIndex = bound(random.nextUint256(), 1, allDomains.length - 1);
        uint32 domain = allDomains[domainIndex];
        // Pick random Notary
        uint256 notaryIndex = bound(random.nextUint256(), 0, DOMAIN_AGENTS - 1);
        address notary = domains[domain].agents[notaryIndex];
        // Fuzz the position of invalid state in the snapshot
        uint256 statesAmount = bound(random.nextUint256(), 1, SNAPSHOT_MAX_STATES);
        uint256 stateIndex = bound(random.nextUint256(), 0, statesAmount - 1);
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, statesAmount, stateIndex);
        (bytes memory snapshot, ) = rawSnap.castToSnapshot();
        // Use random metadata
        RawAttestation memory ra = random.nextAttestation(rawSnap, random.nextUint32());
        (bytes memory attestation, ) = ra.castToAttestation();
        bytes memory signature = signAttestation(notary, attestation);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidAttestationState(stateIndex, snapshot, attestation, signature);
            vm.expectEmit(true, true, true, true);
            emit AgentRemoved(domain, notary);
            vm.expectEmit(true, true, true, true);
            emit AgentSlashed(domain, notary);
        }
        vm.recordLogs();
        assertEq(
            InterfaceOrigin(origin).verifyAttestation(snapshot, stateIndex, attestation, signature),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    function _verifySnapshot(
        RawState memory rawState,
        bool isValid,
        uint256 statesAmount,
        uint256 stateIndex
    ) internal {
        // Fuzz the position of invalid state in the snapshot
        statesAmount = bound(statesAmount, 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(stateIndex, 0, statesAmount - 1);
        address notary = domains[DOMAIN_REMOTE].agent;
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, statesAmount, stateIndex);
        (bytes memory snapshot, ) = rawSnap.castToSnapshot();
        bytes memory signature = signSnapshot(notary, snapshot);
        vm.recordLogs();
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidSnapshotState(stateIndex, snapshot, signature);
            vm.expectEmit(true, true, true, true);
            emit AgentRemoved(DOMAIN_REMOTE, notary);
            vm.expectEmit(true, true, true, true);
            emit AgentSlashed(DOMAIN_REMOTE, notary);
        }
        assertEq(
            InterfaceOrigin(origin).verifySnapshot(snapshot, stateIndex, signature),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }
}
