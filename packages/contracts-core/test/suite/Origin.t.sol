// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRegistry } from "../../contracts/interfaces/ISystemRegistry.sol";
import { IStateHub } from "../../contracts/interfaces/IStateHub.sol";
import { SNAPSHOT_MAX_STATES } from "../../contracts/libs/Constants.sol";
import { SystemEntity } from "../../contracts/libs/Structures.sol";
import { TipsLib } from "../../contracts/libs/Tips.sol";

import { InterfaceOrigin } from "../../contracts/Origin.sol";
import { Versioned } from "../../contracts/Version.sol";

import { fakeState, fakeSnapshot } from "../utils/libs/FakeIt.t.sol";
import { Random } from "../utils/libs/Random.t.sol";
import {
    StateFlag,
    RawAttestation,
    RawHeader,
    RawMessage,
    RawSnapshot,
    RawState,
    RawStateReport,
    RawTips
} from "../utils/libs/SynapseStructs.t.sol";
import { addressToBytes32 } from "../utils/libs/SynapseUtilities.t.sol";
import { AgentFlag, ISystemContract, SynapseTest } from "../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract OriginTest is SynapseTest {
    // Deploy Production version of Origin and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_ORIGIN) {}

    function test_setupCorrectly() public {
        // Check Messaging addresses
        assertEq(
            address(ISystemContract(origin).systemRouter()),
            address(systemRouter),
            "!systemRouter"
        );
        // TODO: adjust when Agent Merkle Tree is implemented
        // Check Agents: currently all Agents are known in LightManager
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                checkAgentStatus(
                    agent,
                    ISystemRegistry(origin).agentStatus(agent),
                    AgentFlag.Active
                );
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
            messages[i] = rawMessages[i].formatMessage();
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
            bytes memory state = rs.formatState();
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
            root: bytes32(0),
            origin: DOMAIN_LOCAL,
            nonce: 0,
            blockNumber: uint40(block.number - 1),
            timestamp: uint40(block.timestamp - BLOCK_TIME)
        });
        bytes memory state = rs.formatState();
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
            state = rs.formatState();
            assertEq(hub.suggestState(i + 1), state, "!suggestState");
        }
        assertEq(hub.suggestLatestState(), state, "!suggestLatestState");
    }

    function test_verifySnapshot_valid(
        uint32 nonce,
        uint256 statesAmount,
        uint256 stateIndex
    ) public {
        // Use empty mutation mask
        test_verifySnapshot_existingNonce(nonce, 0, statesAmount, stateIndex);
    }

    function test_verifySnapshot_existingNonce(
        uint32 nonce,
        uint256 mask,
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

    function test_verifyAttestation_valid(Random memory random, uint32 nonce) public {
        test_verifyAttestation_existingNonce(random, nonce, 0);
    }

    function test_verifyAttestation_existingNonce(
        Random memory random,
        uint32 nonce,
        uint256 mask
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

    function test_verifyAttestationWithProof_valid(Random memory random, uint32 nonce) public {
        // Use empty mutation mask
        test_verifyAttestationWithProof_existingNonce(random, nonce, 0);
    }

    function test_verifyAttestationWithProof_existingNonce(
        Random memory random,
        uint32 nonce,
        uint256 mask
    ) public {
        (bool isValid, RawState memory rs) = _prepareExistingState(nonce, mask);
        _verifyAttestationWithProof(random, rs, isValid);
    }

    function test_verifyAttestationWithProof_unknownNonce(Random memory random, RawState memory rs)
        public
    {
        // Restrict nonce to non-existing ones
        rs.nonce = uint32(bound(rs.nonce, MESSAGES + 1, type(uint32).max));
        rs.origin = DOMAIN_LOCAL;
        // Remaining fields are fuzzed
        _verifyAttestationWithProof(random, rs, false);
    }

    function _prepareExistingState(uint32 nonce, uint256 mask)
        internal
        returns (bool isValid, RawState memory rs)
    {
        uint40 initialBN = uint40(block.number - 1);
        uint40 initialTS = uint40(block.timestamp - BLOCK_TIME);
        test_dispatch();
        // State is valid if and only if all three fields match
        isValid = mask & 7 == 0;
        // Restrict nonce to existing ones
        nonce = uint32(bound(nonce, 0, MESSAGES));
        rs = RawState({
            root: getRoot(nonce),
            origin: DOMAIN_LOCAL,
            nonce: nonce,
            blockNumber: initialBN + nonce,
            timestamp: uint40(initialTS + nonce * BLOCK_TIME)
        });
        rs.root = rs.root ^ bytes32(mask & 1);
        rs.blockNumber = rs.blockNumber ^ uint40(mask & 2);
        rs.timestamp = rs.timestamp ^ uint40(mask & 4);
    }

    function _prepareAttestation(Random memory random, RawState memory rawState)
        internal
        returns (
            uint32 domain,
            address notary,
            uint256 stateIndex,
            bytes memory snapshot,
            RawAttestation memory ra
        )
    {
        // Pick random domain expect for 0
        uint256 domainIndex = bound(random.nextUint256(), 1, allDomains.length - 1);
        domain = allDomains[domainIndex];
        // Pick random Notary
        uint256 notaryIndex = bound(random.nextUint256(), 0, DOMAIN_AGENTS - 1);
        notary = domains[domain].agents[notaryIndex];
        // Fuzz the position of invalid state in the snapshot
        uint256 statesAmount = bound(random.nextUint256(), 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(random.nextUint256(), 0, statesAmount - 1);
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, statesAmount, stateIndex);
        snapshot = rawSnap.formatSnapshot();
        // Use random metadata
        ra = random.nextAttestation(rawSnap, random.nextUint32());
        // Save snapshot for Snapshot Proof generation
        acceptSnapshot(rawSnap.formatStates());
    }

    function _verifyAttestation(
        Random memory random,
        RawState memory rawState,
        bool isValid
    ) internal {
        (
            uint32 domain,
            address notary,
            uint256 stateIndex,
            bytes memory snapshot,
            RawAttestation memory ra
        ) = _prepareAttestation(random, rawState);
        bytes memory state = rawState.formatState();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidAttestationState(stateIndex, state, attPayload, attSig);
            // TODO: check that anyone could make the call
            _expectAgentSlashed(domain, notary, address(this));
        }
        vm.recordLogs();
        assertEq(
            InterfaceOrigin(origin).verifyAttestation(stateIndex, snapshot, attPayload, attSig),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    function _verifyAttestationWithProof(
        Random memory random,
        RawState memory rawState,
        bool isValid
    ) internal {
        (
            uint32 domain,
            address notary,
            uint256 stateIndex,
            ,
            RawAttestation memory ra
        ) = _prepareAttestation(random, rawState);
        bytes32[] memory snapProof = genSnapshotProof(stateIndex);
        bytes memory state = rawState.formatState();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidAttestationState(stateIndex, state, attPayload, attSig);
            // TODO: check that anyone could make the call
            _expectAgentSlashed(domain, notary, address(this));
        }
        vm.recordLogs();
        assertEq(
            InterfaceOrigin(origin).verifyAttestationWithProof(
                stateIndex,
                state,
                snapProof,
                attPayload,
                attSig
            ),
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
        (bytes memory snapPayload, bytes memory snapSig) = signSnapshot(notary, rawSnap);
        vm.recordLogs();
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidSnapshotState(stateIndex, snapPayload, snapSig);
            // TODO: check that anyone could make the call
            _expectAgentSlashed(DOMAIN_REMOTE, notary, address(this));
        }
        assertEq(
            InterfaceOrigin(origin).verifySnapshot(stateIndex, snapPayload, snapSig),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
        _verifyStateReport(rawState, isValid);
    }

    function _verifyStateReport(RawState memory rawState, bool isStateValid) internal {
        // Report is valid only if reported state is invalid
        bool isValid = !isStateValid;
        RawStateReport memory rawSR = RawStateReport(uint8(StateFlag.Invalid), rawState);
        address guard = domains[0].agent;
        (bytes memory srPayload, bytes memory srSig) = signStateReport(guard, rawSR);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidStateReport(srPayload, srSig);
            // TODO: check that anyone could make the call
            _expectAgentSlashed(0, guard, address(this));
        }
        vm.recordLogs();
        assertEq(
            InterfaceOrigin(origin).verifyStateReport(srPayload, srSig),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    function _expectAgentSlashed(
        uint32 domain,
        address agent,
        address prover
    ) internal {
        vm.expectEmit(true, true, true, true);
        emit AgentSlashed(domain, agent, prover);
        vm.expectCall(
            address(lightManager),
            abi.encodeWithSelector(lightManager.registrySlash.selector, domain, agent)
        );
    }
}
