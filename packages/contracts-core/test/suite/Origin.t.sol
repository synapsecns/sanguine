// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentSecured} from "../../contracts/interfaces/IAgentSecured.sol";
import {InterfaceGasOracle} from "../../contracts/interfaces/InterfaceGasOracle.sol";
import {IStateHub} from "../../contracts/interfaces/IStateHub.sol";
import {
    EthTransferFailed,
    IncorrectDestinationDomain,
    InsufficientEthBalance,
    TipsValueTooLow
} from "../../contracts/libs/Errors.sol";
import {SNAPSHOT_MAX_STATES} from "../../contracts/libs/Constants.sol";
import {TipsLib} from "../../contracts/libs/stack/Tips.sol";

import {InterfaceOrigin} from "../../contracts/Origin.sol";
import {Versioned} from "../../contracts/base/Version.sol";

import {RevertingApp} from "../harnesses/client/RevertingApp.t.sol";
import {console, stdJson} from "forge-std/Script.sol";

import {fakeState, fakeSnapshot} from "../utils/libs/FakeIt.t.sol";
import {Random} from "../utils/libs/Random.t.sol";
import {
    MessageFlag,
    RawAttestation,
    RawBaseMessage,
    RawGasData,
    RawHeader,
    RawMessage,
    RawRequest,
    RawSnapshot,
    RawState,
    RawStateIndex,
    RawTips
} from "../utils/libs/SynapseStructs.t.sol";
import {AgentFlag, GasOracle, Origin, SynapseTest} from "../utils/SynapseTest.t.sol";
import {AgentSecuredTest} from "./base/AgentSecured.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract OriginTest is AgentSecuredTest {
    address public sender = makeAddr("Sender");
    address public recipient = makeAddr("Recipient");
    uint32 public period = 1 minutes;
    RawTips public tips = RawTips(0, 0, 0, 0);
    RawRequest public request = RawRequest({gasLimit: 100_000, gasDrop: 0, version: 0});

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

    function test_constructor_revert_chainIdOverflow() public {
        vm.chainId(2 ** 32);
        vm.expectRevert("SafeCast: value doesn't fit in 32 bits");
        new Origin({synapseDomain_: 1, agentManager_: address(2), inbox_: address(3), gasOracle_: address(4)});
    }

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = uint32(block.chainid);
        address caller = random.nextAddress();
        address agentManager = random.nextAddress();
        address inbox_ = random.nextAddress();
        address gasOracle_ = address(new GasOracle(DOMAIN_SYNAPSE, random.nextAddress()));
        address owner_ = random.nextAddress();

        Origin cleanContract = new Origin(DOMAIN_SYNAPSE, agentManager, inbox_, gasOracle_);
        vm.prank(caller);
        cleanContract.initialize(owner_);
        assertEq(cleanContract.owner(), owner_, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.agentManager(), agentManager, "!agentManager");
        assertEq(cleanContract.inbox(), inbox_, "!inbox");
        assertEq(cleanContract.gasOracle(), gasOracle_, "!gasOracle");
        assertEq(cleanContract.statesAmount(), 1, "!statesAmount");
    }

    function initializeLocalContract() public override {
        Origin(localContract()).initialize(address(0));
    }

    function test_sendBaseMessage_revert_blockTimestampOverflow() public {
        vm.warp(2 ** 40);
        vm.prank(sender);
        vm.expectRevert("SafeCast: value doesn't fit in 40 bits");
        InterfaceOrigin(origin).sendBaseMessage(
            DOMAIN_REMOTE, addressToBytes32(recipient), period, request.encodeRequest(), "test content"
        );
    }

    function test_sendBaseMessage_revert_blockNumberOverflow() public {
        vm.roll(2 ** 40);
        vm.prank(sender);
        vm.expectRevert("SafeCast: value doesn't fit in 40 bits");
        InterfaceOrigin(origin).sendBaseMessage(
            DOMAIN_REMOTE, addressToBytes32(recipient), period, request.encodeRequest(), "test content"
        );
    }

    function test_sendBaseMessage_revert_tipsTooLow(RawTips memory minTips, uint256 msgValue) public {
        minTips.boundTips(1 ** 32);
        minTips.floorTips(1);
        msgValue = msgValue % minTips.castToTips().value();
        // Force gasOracle.getMinimumTips(DOMAIN_REMOTE, *, *) to return minTips
        vm.mockCall(
            gasOracle,
            abi.encodeWithSelector(InterfaceGasOracle.getMinimumTips.selector, DOMAIN_REMOTE),
            abi.encode(minTips.encodeTips())
        );
        deal(sender, msgValue);
        vm.expectRevert(TipsValueTooLow.selector);
        vm.prank(sender);
        InterfaceOrigin(origin).sendBaseMessage{value: msgValue}(
            DOMAIN_REMOTE, addressToBytes32(recipient), period, request.encodeRequest(), "test content"
        );
    }

    function test_sendBaseMessage_revert_sameDestination() public {
        vm.expectRevert(IncorrectDestinationDomain.selector);
        vm.prank(sender);
        InterfaceOrigin(origin).sendBaseMessage(
            localDomain(), addressToBytes32(recipient), period, request.encodeRequest(), "test content"
        );
    }

    function test_sendManagementMessage_revert_sameDestination() public {
        vm.expectRevert(IncorrectDestinationDomain.selector);
        vm.prank(localAgentManager());
        InterfaceOrigin(origin).sendManagerMessage(localDomain(), period, "test payload");
    }

    function test_getMinimumTipsValue(
        uint32 destination_,
        uint256 paddedRequest,
        uint256 contentLength,
        RawTips memory minTips
    ) public {
        minTips.boundTips(1 ** 32);
        // Force gasOracle.getMinimumTips(destination_, *, *) to return minTips
        vm.mockCall(
            gasOracle,
            abi.encodeWithSelector(InterfaceGasOracle.getMinimumTips.selector, destination_),
            abi.encode(minTips.encodeTips())
        );
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

    function test_sendMessages(RawGasData memory rgd) public {
        // Force gasOracle.getGasData() to return rgd
        vm.mockCall(
            gasOracle, abi.encodeWithSelector(InterfaceGasOracle.getGasData.selector), abi.encode(rgd.encodeGasData())
        );
        uint192 encodedRequest = request.encodeRequest();
        bytes memory content = "test content";
        bytes memory body = RawBaseMessage({
            tips: tips,
            sender: addressToBytes32(sender),
            recipient: addressToBytes32(recipient),
            request: request,
            content: content
        }).formatBaseMessage();
        bytes[] memory messages = new bytes[](MESSAGES);
        bytes32[] memory leafs = new bytes32[](MESSAGES);
        bytes32[] memory roots = new bytes32[](MESSAGES);
        for (uint32 i = 0; i < MESSAGES; ++i) {
            RawMessage memory rm = RawMessage(
                RawHeader({
                    flag: uint8(MessageFlag.Base),
                    origin: DOMAIN_LOCAL,
                    nonce: i + 1,
                    destination: DOMAIN_REMOTE,
                    optimisticPeriod: period
                }),
                body
            );
            messages[i] = rm.formatMessage();
            leafs[i] = rm.castToMessage().leaf();
            insertMessage(leafs[i]);
            roots[i] = getRoot(i + 1);
        }

        for (uint32 i = 0; i < MESSAGES; ++i) {
            // Expect Origin Events
            RawState memory rs = RawState({
                root: roots[i],
                origin: DOMAIN_LOCAL,
                nonce: i + 1,
                blockNumber: uint40(block.number),
                timestamp: uint40(block.timestamp),
                gasData: rgd
            });
            bytes memory state = rs.formatState();
            vm.expectEmit();
            emit StateSaved(state);
            vm.expectEmit();
            emit Sent(leafs[i], i + 1, DOMAIN_REMOTE, messages[i]);
            vm.expectCall(gasOracle, abi.encodeCall(InterfaceGasOracle.updateGasData, (DOMAIN_REMOTE)));
            vm.prank(sender);
            (uint32 messageNonce, bytes32 messageHash) = InterfaceOrigin(origin).sendBaseMessage(
                DOMAIN_REMOTE, addressToBytes32(recipient), period, encodedRequest, content
            );
            // Check return values
            assertEq(messageNonce, i + 1, "!messageNonce");
            assertEq(messageHash, leafs[i], "!messageHash");
            skipBlock();
        }
    }

    function test_states(RawGasData memory rgd) public {
        IStateHub hub = IStateHub(origin);
        // Check initial States
        assertEq(hub.statesAmount(), 1, "!initial statesAmount");
        // Initial state was saved "1 block ago"
        RawState memory rs;
        rs.origin = DOMAIN_LOCAL;
        rs.blockNumber = uint40(block.number - 1);
        rs.timestamp = uint40(block.timestamp - BLOCK_TIME);
        bytes memory state = rs.formatState();
        assertEq(hub.suggestState(0), state, "!state: 0");
        assertEq(hub.suggestState(0), hub.suggestLatestState(), "!latest state: 0");
        // Send some messages
        test_sendMessages(rgd);
        // Check saved States
        assertEq(hub.statesAmount(), MESSAGES + 1, "!statesAmount");
        assertEq(hub.suggestState(0), state, "!suggestState: 0");
        for (uint32 i = 0; i < MESSAGES; ++i) {
            rs.nonce += 1;
            rs.root = getRoot(rs.nonce);
            rs.blockNumber += 1;
            rs.timestamp += uint40(BLOCK_TIME);
            rs.gasData = rgd;
            state = rs.formatState();
            assertEq(hub.suggestState(i + 1), state, "!suggestState");
        }
        assertEq(hub.suggestLatestState(), state, "!suggestLatestState");
    }

    function test_verifySnapshot_valid(uint32 nonce, RawGasData memory rgd, RawStateIndex memory rsi) public {
        // Use empty mutation mask
        test_verifySnapshot_existingNonce(nonce, 0, rgd, rsi);
    }

    function test_verifySnapshot_existingNonce(
        uint32 nonce,
        uint256 mask,
        RawGasData memory rgd,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        (bool isValid, RawState memory rs) = _prepareExistingState(rgd, nonce, mask);
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
        (bool isValid, RawState memory rs) = _prepareExistingState(random.nextGasData(), nonce, mask);
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
        (bool isValid, RawState memory rs) = _prepareExistingState(random.nextGasData(), nonce, mask);
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

    function _prepareExistingState(RawGasData memory rgd, uint32 nonce, uint256 mask)
        internal
        returns (bool isValid, RawState memory rs)
    {
        uint40 initialBN = uint40(block.number - 1);
        uint40 initialTS = uint40(block.timestamp - BLOCK_TIME);
        test_sendMessages(rgd);
        // State is valid if and only if all three fields match
        isValid = mask & 7 == 0;
        // Restrict nonce to existing ones
        nonce = uint32(bound(nonce, 0, MESSAGES));
        rs.origin = DOMAIN_LOCAL;
        rs.nonce = nonce;
        rs.root = getRoot(nonce) ^ bytes32(mask & 1);
        rs.blockNumber = (initialBN + nonce) ^ uint40(mask & 2);
        rs.timestamp = uint40(initialTS + nonce * BLOCK_TIME) ^ uint40(mask & 4);
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
        acceptSnapshot(rawSnap);
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
            expectDisputeResolved(0, notary, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(
            lightInbox.verifyStateWithAttestation(rsi.stateIndex, snapshot, attPayload, attSig), isValid, "!returnValue"
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
            expectDisputeResolved(0, notary, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(
            lightInbox.verifyStateWithSnapshotProof(rsi.stateIndex, state, snapProof, attPayload, attSig),
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
            expectDisputeResolved(0, notary, address(0), address(this));
        }
        assertEq(lightInbox.verifyStateWithSnapshot(rsi.stateIndex, snapPayload, snapSig), isValid, "!returnValue");
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
        _verifyStateReport(rawState, isValid);
    }

    function _verifyStateReport(RawState memory rawState, bool isStateValid) internal {
        // Report is valid only if reported state is invalid
        bool isValid = !isStateValid;
        address guard = domains[0].agent;
        (bytes memory statePayload, bytes memory srSig) = signStateReport(guard, rawState);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidStateReport(statePayload, srSig);
            // TODO: check that anyone could make the call
            expectStatusUpdated(AgentFlag.Fraudulent, 0, guard);
            expectDisputeResolved(0, guard, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(lightInbox.verifyStateReport(statePayload, srSig), isValid, "!returnValue");
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    // ════════════════════════════════════════════ TEST: WITHDRAW TIPS ════════════════════════════════════════════════

    function test_withdrawTips(uint256 amount) public {
        vm.deal(origin, amount);
        vm.expectEmit(origin);
        emit TipWithdrawalCompleted(recipient, amount);
        vm.prank(address(lightManager));
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
        assertEq(recipient.balance, amount);
    }

    function test_remoteWithdrawTips_revert_insufficientBalance(uint256 balance, uint256 amount) public {
        amount = bound(amount, 1, type(uint256).max);
        balance = balance % amount;
        vm.deal(origin, balance);
        vm.expectRevert(InsufficientEthBalance.selector);
        vm.prank(address(lightManager));
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
    }

    function test_withdrawTips_revert_recipientReverted(uint256 amount) public {
        address revertingRecipient = address(new RevertingApp());
        vm.deal(origin, amount);
        vm.expectRevert(EthTransferFailed.selector);
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
