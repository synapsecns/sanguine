// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentSecured} from "../../contracts/interfaces/IAgentSecured.sol";
import {InterfaceGasOracle} from "../../contracts/interfaces/InterfaceGasOracle.sol";
import {IStateHub} from "../../contracts/interfaces/IStateHub.sol";
import {SNAPSHOT_MAX_STATES} from "../../contracts/libs/Constants.sol";
import {SystemEntity} from "../../contracts/libs/Structures.sol";
import {TipsLib} from "../../contracts/libs/Tips.sol";

import {InterfaceOrigin} from "../../contracts/Origin.sol";
import {Versioned} from "../../contracts/base/Version.sol";

import {RevertingApp} from "../harnesses/client/RevertingApp.t.sol";
import {BaseMock, GasOracleMock} from "../mocks/GasOracleMock.t.sol";

import {fakeState, fakeSnapshot} from "../utils/libs/FakeIt.t.sol";
import {Random} from "../utils/libs/Random.t.sol";
import {
    MessageFlag,
    StateFlag,
    RawAttestation,
    RawBaseMessage,
    RawHeader,
    RawMessage,
    RawRequest,
    RawSnapshot,
    RawState,
    RawStateIndex,
    RawStateReport,
    RawTips
} from "../utils/libs/SynapseStructs.t.sol";
import {AgentFlag, Origin, SynapseTest} from "../utils/SynapseTest.t.sol";
import {AgentSecuredTest} from "./base/AgentSecured.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract OriginTest is AgentSecuredTest {
    address public sender = makeAddr("Sender");
    address public recipient = makeAddr("Recipient");
    uint32 public period = 1 minutes;
    RawTips public tips = RawTips(0, 0, 0, 0);
    RawRequest public request = RawRequest({gasLimit: 100_000, gasDrop: 0});

    // Deploy Production version of Origin and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_ORIGIN) {}

    function test_setupCorrectly() public {
        // Check Agents: currently all Agents are known in LightManager
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                checkAgentStatus(agent, IAgentSecured(origin).agentStatus(agent), AgentFlag.Active);
            }
        }
        // Check version
        assertEq(Versioned(origin).version(), LATEST_VERSION, "!version");
    }

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        address caller = random.nextAddress();
        address agentManager = random.nextAddress();
        address gasOracle_ = random.nextAddress();
        Origin cleanContract = new Origin(domain, agentManager, gasOracle_);
        vm.prank(caller);
        cleanContract.initialize();
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.agentManager(), agentManager, "!agentManager");
        assertEq(cleanContract.gasOracle(), gasOracle_, "!gasOracle");
        assertEq(cleanContract.statesAmount(), 1, "!statesAmount");
    }

    function initializeLocalContract() public override {
        Origin(localContract()).initialize();
    }

    function test_sendBaseMessage_revert_tipsTooLow(RawTips memory minTips, uint256 msgValue) public {
        minTips.boundTips(1 ** 32);
        minTips.floorTips(1);
        msgValue = msgValue % minTips.castToTips().value();
        GasOracleMock(gasOracle).setMockReturnValue(minTips.encodeTips());
        deal(sender, msgValue);
        vm.expectRevert("Tips value too low");
        vm.prank(sender);
        InterfaceOrigin(origin).sendBaseMessage{value: msgValue}(
            DOMAIN_REMOTE, addressToBytes32(recipient), period, request.encodeRequest(), "test content"
        );
    }

    function test_getMinimumTipsValue(
        uint32 destination_,
        uint256 paddedRequest,
        uint256 contentLength,
        RawTips memory minTips
    ) public {
        minTips.boundTips(1 ** 32);
        GasOracleMock(gasOracle).setMockReturnValue(minTips.encodeTips());
        vm.expectCall(
            address(gasOracle),
            abi.encodeWithSelector(
                InterfaceGasOracle.getMinimumTips.selector, destination_, paddedRequest, contentLength
            )
        );
        assertEq(
            InterfaceOrigin(origin).getMinimumTipsValue(destination_, paddedRequest, contentLength),
            minTips.castToTips().value(),
            "!getMinimumTipsValue"
        );
    }

    function test_sendMessages() public {
        uint160 encodedRequest = request.encodeRequest();
        bytes memory content = "test content";
        bytes memory body = RawBaseMessage({
            sender: addressToBytes32(sender),
            recipient: addressToBytes32(recipient),
            tips: tips,
            request: request,
            content: content
        }).formatBaseMessage();
        bytes[] memory messages = new bytes[](MESSAGES);
        bytes32[] memory roots = new bytes32[](MESSAGES);
        for (uint32 i = 0; i < MESSAGES; ++i) {
            messages[i] = RawMessage(
                uint8(MessageFlag.Base),
                RawHeader({origin: DOMAIN_LOCAL, nonce: i + 1, destination: DOMAIN_REMOTE, optimisticPeriod: period}),
                body
            ).formatMessage();
            insertMessage(messages[i]);
            roots[i] = getRoot(i + 1);
        }

        // Expect Origin Events
        for (uint32 i = 0; i < MESSAGES; ++i) {
            // 1 block is skipped after each sent message
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
            emit Sent(keccak256(messages[i]), i + 1, DOMAIN_REMOTE, messages[i]);
        }

        for (uint32 i = 0; i < MESSAGES; ++i) {
            vm.prank(sender);
            (uint32 messageNonce, bytes32 messageHash) = InterfaceOrigin(origin).sendBaseMessage(
                DOMAIN_REMOTE, addressToBytes32(recipient), period, encodedRequest, content
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
        // Send some messages
        test_sendMessages();
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

    function test_verifySnapshot_valid(uint32 nonce, RawStateIndex memory rsi) public {
        // Use empty mutation mask
        test_verifySnapshot_existingNonce(nonce, 0, rsi);
    }

    function test_verifySnapshot_existingNonce(uint32 nonce, uint256 mask, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        (bool isValid, RawState memory rs) = _prepareExistingState(nonce, mask);
        _verifySnapshot(rs, isValid, rsi);
    }

    function test_verifySnapshot_unknownNonce(RawState memory rs, RawStateIndex memory rsi) public boundIndex(rsi) {
        // Restrict nonce to non-existing ones
        rs.nonce = uint32(bound(rs.nonce, MESSAGES + 1, type(uint32).max));
        rs.origin = DOMAIN_LOCAL;
        // Remaining fields are fuzzed
        _verifySnapshot(rs, false, rsi);
    }

    function test_verifyAttestation_valid(Random memory random, uint32 nonce) public {
        test_verifyAttestation_existingNonce(random, nonce, 0);
    }

    function test_verifyAttestation_existingNonce(Random memory random, uint32 nonce, uint256 mask) public {
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

    function test_verifyAttestationWithProof_existingNonce(Random memory random, uint32 nonce, uint256 mask) public {
        (bool isValid, RawState memory rs) = _prepareExistingState(nonce, mask);
        _verifyAttestationWithProof(random, rs, isValid);
    }

    function test_verifyAttestationWithProof_unknownNonce(Random memory random, RawState memory rs) public {
        // Restrict nonce to non-existing ones
        rs.nonce = uint32(bound(rs.nonce, MESSAGES + 1, type(uint32).max));
        rs.origin = DOMAIN_LOCAL;
        // Remaining fields are fuzzed
        _verifyAttestationWithProof(random, rs, false);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function _prepareExistingState(uint32 nonce, uint256 mask) internal returns (bool isValid, RawState memory rs) {
        uint40 initialBN = uint40(block.number - 1);
        uint40 initialTS = uint40(block.timestamp - BLOCK_TIME);
        test_sendMessages();
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
            RawStateIndex memory rsi,
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
        rsi = random.nextStateIndex();
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, rsi);
        snapshot = rawSnap.formatSnapshot();
        // Use random metadata
        ra = random.nextAttestation(rawSnap, random.nextUint32());
        // Save snapshot for Snapshot Proof generation
        acceptSnapshot(rawSnap.formatStates());
    }

    function _verifyAttestation(Random memory random, RawState memory rawState, bool isValid) internal {
        (uint32 domain, address notary, RawStateIndex memory rsi, bytes memory snapshot, RawAttestation memory ra) =
            _prepareAttestation(random, rawState);
        bytes memory state = rawState.formatState();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidStateWithAttestation(rsi.stateIndex, state, attPayload, attSig);
            // TODO: check that anyone could make the call
            expectStatusUpdated(AgentFlag.Fraudulent, domain, notary);
            expectDisputeResolved(notary, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(
            lightManager.verifyStateWithAttestation(rsi.stateIndex, snapshot, attPayload, attSig),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    function _verifyAttestationWithProof(Random memory random, RawState memory rawState, bool isValid) internal {
        (uint32 domain, address notary, RawStateIndex memory rsi,, RawAttestation memory ra) =
            _prepareAttestation(random, rawState);
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        bytes memory state = rawState.formatState();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidStateWithAttestation(rsi.stateIndex, state, attPayload, attSig);
            // TODO: check that anyone could make the call
            expectStatusUpdated(AgentFlag.Fraudulent, domain, notary);
            expectDisputeResolved(notary, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(
            lightManager.verifyStateWithSnapshotProof(rsi.stateIndex, state, snapProof, attPayload, attSig),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    function _verifySnapshot(RawState memory rawState, bool isValid, RawStateIndex memory rsi) internal {
        address notary = domains[DOMAIN_REMOTE].agent;
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, rsi);
        (bytes memory snapPayload, bytes memory snapSig) = signSnapshot(notary, rawSnap);
        vm.recordLogs();
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidStateWithSnapshot(rsi.stateIndex, snapPayload, snapSig);
            // TODO: check that anyone could make the call
            expectStatusUpdated(AgentFlag.Fraudulent, DOMAIN_REMOTE, notary);
            expectDisputeResolved(notary, address(0), address(this));
        }
        assertEq(lightManager.verifyStateWithSnapshot(rsi.stateIndex, snapPayload, snapSig), isValid, "!returnValue");
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
            expectStatusUpdated(AgentFlag.Fraudulent, 0, guard);
            expectDisputeResolved(guard, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(lightManager.verifyStateReport(srPayload, srSig), isValid, "!returnValue");
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    // ════════════════════════════════════════════ TEST: WITHDRAW TIPS ════════════════════════════════════════════════

    function test_withdrawTips(uint256 amount) public {
        vm.deal(origin, amount);
        vm.prank(address(lightManager));
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
        assertEq(recipient.balance, amount);
    }

    function test_remoteWithdrawTips_revert_insufficientBalance(uint256 balance, uint256 amount) public {
        amount = bound(amount, 1, type(uint256).max);
        balance = balance % amount;
        vm.deal(origin, balance);
        vm.expectRevert("Insufficient balance");
        vm.prank(address(lightManager));
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
    }

    function test_withdrawTips_revert_recipientReverted(uint256 amount) public {
        address revertingRecipient = address(new RevertingApp());
        vm.deal(origin, amount);
        vm.expectRevert("Recipient reverted");
        vm.prank(address(lightManager));
        InterfaceOrigin(origin).withdrawTips(revertingRecipient, amount);
    }

    // ═════════════════════════════════════════════════ OVERRIDES ═════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return localOrigin();
    }
}
