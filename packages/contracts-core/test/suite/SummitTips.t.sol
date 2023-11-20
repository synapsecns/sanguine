// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DISPUTE_TIMEOUT_NOTARY} from "../../contracts/libs/Constants.sol";
import {
    AgentNotNotary,
    CallerNotInbox,
    DisputeTimeoutNotOver,
    IncorrectSnapshotRoot,
    IncorrectTipsProof,
    NotaryInDispute,
    TipsClaimMoreThanEarned,
    TipsClaimZero
} from "../../contracts/libs/Errors.sol";
import {IAgentManager, InterfaceSummit} from "../../contracts/Summit.sol";

import {AgentFlag, AgentStatus, Summit, SynapseTest} from "../utils/SynapseTest.t.sol";
import {AgentSecuredTest} from "./hubs/ExecutionHub.t.sol";

import {fakeState} from "../utils/libs/FakeIt.t.sol";
import {Random} from "../utils/libs/Random.t.sol";
import {
    RawExecReceipt,
    RawState,
    RawStateIndex,
    RawSnapshot,
    RawTips,
    RawTipsProof
} from "../utils/libs/SynapseStructs.t.sol";

import {stdStorage, StdStorage} from "forge-std/Test.sol";

contract SummitCheats is Summit {
    constructor(uint32 domain, address agentManager_, address inbox_) Summit(domain, agentManager_, inbox_) {}

    function setActorTips(address actor, uint32 origin, uint128 earned, uint128 claimed) external {
        actorTips[actor][origin].earned = earned;
        actorTips[actor][origin].claimed = claimed;
    }
}

// solhint-disable code-complexity
// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract SummitTipsTest is AgentSecuredTest {
    using stdStorage for StdStorage;

    RawState internal state0;
    address internal guard0;
    uint32 internal origin0;

    RawState internal state1;
    address internal guard1;
    uint32 internal origin1;

    // Notary[0] snapshot: (state0)
    RawSnapshot internal snapshot0;
    address internal snapNotary0;
    bytes32 internal snapRoot0;

    // Notary[1] snapshot: (state0, state1)
    RawSnapshot internal snapshot1;
    address internal snapNotary1;
    bytes32 internal snapRoot1;

    // Notary who posted Receipt to Summit
    address internal rcptNotary;
    address internal rcptNotaryFinal;

    address internal summitCheats;

    // Deploy Production version of Destination and Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_DESTINATION_SYNAPSE | DEPLOY_PROD_SUMMIT) {}

    modifier checkQueueLength(int256 diff) {
        uint256 len = InterfaceSummit(summit).receiptQueueLength();
        _;
        assertEq(InterfaceSummit(summit).receiptQueueLength(), uint256(int256(len) + diff), "!queueLength");
    }

    function setUp() public override {
        super.setUp();
        guard0 = domains[0].agents[0];
        guard1 = domains[0].agents[1];
        snapNotary0 = domains[DOMAIN_LOCAL].agents[0];
        snapNotary1 = domains[DOMAIN_LOCAL].agents[1];
        // Prepare test snapshot data
        origin0 = 1;
        state0 = fakeState(origin0);
        origin1 = 2;
        state1 = fakeState(origin1);
        snapshot0.states.push(state0);
        snapshot1.states.push(state0);
        snapshot1.states.push(state1);
        // Submit snapshots to Summit
        submitGuardSnapshot(guard0, state0);
        submitGuardSnapshot(guard1, state1);
        submitSnapshot(snapNotary0, snapshot0);
        submitSnapshot(snapNotary1, snapshot1);
        // Extract snapshot roots
        acceptSnapshot(snapshot0);
        snapRoot0 = getSnapshotRoot();
        acceptSnapshot(snapshot1);
        snapRoot1 = getSnapshotRoot();
        // Deploy Summit implementation with Cheats
        vm.chainId(DOMAIN_SYNAPSE);
        summitCheats = address(new SummitCheats(DOMAIN_SYNAPSE, address(bondingManager), address(inbox)));
    }

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = DOMAIN_SYNAPSE;
        vm.chainId(domain);
        address agentManager = random.nextAddress();
        address inbox_ = random.nextAddress();
        address caller = random.nextAddress();
        address owner_ = random.nextAddress();
        Summit cleanContract = new Summit(domain, agentManager, inbox_);
        vm.prank(caller);
        cleanContract.initialize(owner_);
        assertEq(cleanContract.owner(), owner_, "!owner");
        assertEq(cleanContract.agentManager(), agentManager, "!agentManager");
        assertEq(cleanContract.inbox(), inbox_, "!inbox");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
    }

    function initializeLocalContract() public override {
        Summit(localContract()).initialize(address(0));
    }

    function prepareNotaryInDisputeTest() internal {
        address notary = domains[DOMAIN_REMOTE].agent;
        address guard = domains[0].agent;
        openTestDispute({guardIndex: agentIndex[guard], notaryIndex: agentIndex[notary]});
    }

    /// @dev Resolves test dispute above in favor of the Notary.
    function prepareNotaryWonDisputeTest() internal {
        address notary = domains[DOMAIN_REMOTE].agent;
        address guard = domains[0].agent;
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
    }

    // ══════════════════════════════════════════ TESTS: SUBMIT RECEIPTS ═══════════════════════════════════════════════

    function test_submitReceipt(
        RawExecReceipt memory re,
        RawTips memory tips,
        RawTipsProof memory rtp,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex,
        bool isSuccess
    ) public checkQueueLength(1) {
        prepareReceipt(re, tips, rtp, originZero, attNotaryIndex, isSuccess);
        rcptNotary = domains[DOMAIN_REMOTE].agents[rcptNotaryIndex % DOMAIN_AGENTS];
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotary, re);
        vm.expectEmit();
        emit ReceiptAccepted(DOMAIN_REMOTE, rcptNotary, rcptPayload, rcptSignature);
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
    }

    function test_submitReceipt_notAccepted_pending() public checkQueueLength(1) {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        test_submitReceipt(re, tips, rtp, false, 0, 0, false);
        re.finalExecutor = createExecutorEOA(re.finalExecutor, "Final Executor");
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotary, re);
        vm.recordLogs();
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function test_submitReceipt_notAccepted_outdatedStatus() public checkQueueLength(0) {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        test_distributeTips_success(re, tips, rtp, false, 0, 0);
        re.finalExecutor = address(0);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotary, re);
        vm.recordLogs();
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function test_submitReceipt_revert_signedByGuard() public {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        prepareReceipt(re, tips, rtp, false, 0, false);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(guard0, re);
        vm.expectRevert(AgentNotNotary.selector);
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
    }

    function test_submitReceipt_revert_notaryInDispute() public {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        prepareReceipt(re, tips, rtp, false, 0, false);
        // Put DOMAIN_REMOTE notary in Dispute
        address notary = domains[DOMAIN_REMOTE].agent;
        openDispute({guard: domains[0].agent, notary: notary});
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert(NotaryInDispute.selector);
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
    }

    function test_submitReceipt_revert_notaryWonDisputeTimeout() public {
        address notary = domains[DOMAIN_REMOTE].agent;
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        prepareReceipt(re, tips, rtp, false, 0, false);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        // Put DOMAIN_REMOTE notary in Dispute
        prepareNotaryInDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        vm.expectRevert(DisputeTimeoutNotOver.selector);
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
    }

    function test_submitReceipt_afterNotaryDisputeTimeout() public checkQueueLength(1) {
        address notary = domains[DOMAIN_REMOTE].agent;
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        prepareReceipt(re, tips, rtp, false, 0, false);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        // Put DOMAIN_REMOTE notary in Dispute
        prepareNotaryInDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        bool wasAccepted =
            inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
        assertTrue(wasAccepted);
    }

    function test_submitReceipt_revert_unknownSnapRoot() public {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        prepareReceipt(re, tips, rtp, false, 0, false);
        re.snapshotRoot = "Some fake snapshot root";
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert(IncorrectSnapshotRoot.selector);
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
    }

    function test_submitReceipt_revert_incorrectTipsProof(uint256 corruptedId, uint256 corruptedBit) public {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        prepareReceipt(re, tips, rtp, false, 0, false);
        uint256 encodedTips = tips.encodeTips();
        corruptedBit = corruptedBit % 256;
        corruptedId = corruptedId % 4;
        // Corrupt a single bit in either of those
        if (corruptedId == 0) {
            // Corrupt message hash
            re.messageHash ^= bytes32(1 << corruptedBit);
        } else if (corruptedId == 1) {
            // Corrupt tips
            encodedTips ^= 1 << corruptedBit;
        } else if (corruptedId == 2) {
            // Corrupt header hash
            rtp.headerHash ^= bytes32(1 << corruptedBit);
        } else {
            // Corrupt body hash
            rtp.bodyHash ^= bytes32(1 << corruptedBit);
        }
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert(IncorrectTipsProof.selector);
        inbox.submitReceipt(rcptPayload, rcptSignature, encodedTips, rtp.headerHash, rtp.bodyHash);
    }

    function test_acceptReceipt_revert_notInbox(address caller) public {
        vm.assume(caller != localInbox());
        vm.expectRevert(CallerNotInbox.selector);
        vm.prank(caller);
        InterfaceSummit(summit).acceptReceipt(0, 0, 0, 0, 0, "");
    }

    // ═══════════════════════════════════════════ TESTS: TIPS AWARDING ════════════════════════════════════════════════

    function test_distributeTips_success(
        RawExecReceipt memory re,
        RawTips memory tips,
        RawTipsProof memory rtp,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex
    ) public checkQueueLength(0) {
        test_submitReceipt(re, tips, rtp, originZero, rcptNotaryIndex, attNotaryIndex, true);
        skip(BONDING_OPTIMISTIC_PERIOD);
        rcptNotaryFinal = rcptNotary;
        expectAwardedTipsEvents({re: re, tips: tips, isFirst: true, isFinal: true});
        assertTrue(InterfaceSummit(summit).distributeTips());
        checkAwardedTips(re, tips, true);
    }

    function test_distributeTips_failed(
        RawExecReceipt memory re,
        RawTips memory tips,
        RawTipsProof memory rtp,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex
    ) public checkQueueLength(0) {
        test_submitReceipt(re, tips, rtp, originZero, rcptNotaryIndex, attNotaryIndex, false);
        skip(BONDING_OPTIMISTIC_PERIOD);
        rcptNotaryFinal = address(0);
        expectAwardedTipsEvents({re: re, tips: tips, isFirst: true, isFinal: false});
        assertTrue(InterfaceSummit(summit).distributeTips());
        checkAwardedTips(re, tips, false);
    }

    function test_distributeTips_failedThenSuccess(
        RawExecReceipt memory re,
        RawTips memory tips,
        RawTipsProof memory rtp,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex,
        uint256 rcptNotaryIndexFinal,
        address finalExecutor
    ) public checkQueueLength(0) {
        test_distributeTips_failed(re, tips, rtp, originZero, rcptNotaryIndex, attNotaryIndex);
        re.finalExecutor = createExecutorEOA(finalExecutor, "Final Executor");
        rcptNotaryFinal = domains[DOMAIN_REMOTE].agents[rcptNotaryIndexFinal % DOMAIN_AGENTS];
        emit log_named_address("Receipt Notary", rcptNotaryFinal);
        emit log_named_address("Attestation Notary", re.attNotary);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotaryFinal, re);
        inbox.submitReceipt(rcptPayload, rcptSignature, tips.encodeTips(), rtp.headerHash, rtp.bodyHash);
        skip(BONDING_OPTIMISTIC_PERIOD);
        expectAwardedTipsEvents({re: re, tips: tips, isFirst: false, isFinal: true});
        assertTrue(InterfaceSummit(summit).distributeTips());
        checkAwardedTips(re, tips, true);
    }

    function test_distributeTips_emptyQueue() public checkQueueLength(0) {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        test_distributeTips_success(re, tips, rtp, true, 0, 0);
        assertFalse(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_optimisticPeriodNotOver(uint256 timePassed) public checkQueueLength(1) {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        test_submitReceipt(re, tips, rtp, false, 0, 0, false);
        timePassed = timePassed % BONDING_OPTIMISTIC_PERIOD;
        skip(timePassed);
        assertFalse(InterfaceSummit(summit).distributeTips());
    }

    /// @dev When Notary in "Dispute Won Timeout" mode, the receipt is moved to the end of the queue,
    /// therefore two receipts in the queue are expected in the modifier.
    function test_distributeTips_attestationNotary_wonDisputeTimeout() public checkQueueLength(2) {
        // rcptNotary: agents[1], attNotary: agents[0]
        prepareTwoReceiptTest(1, 0);
        prepareNotaryInDisputeTest();
        skip(BONDING_OPTIMISTIC_PERIOD);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    /// @dev Should distribute tips if the "dispute won" timeout is over.
    function test_distributeTips_attestationNotary_afterNotaryDisputeTimeout() public checkQueueLength(1) {
        // rcptNotary: agents[1], attNotary: agents[0]
        prepareTwoReceiptTest(1, 0);
        prepareNotaryInDisputeTest();
        skip(BONDING_OPTIMISTIC_PERIOD);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        vm.recordLogs();
        assertTrue(InterfaceSummit(summit).distributeTips());
        // Should end up emitting TipsAwarded logs (TODO: better way to check this)
        assertGt(vm.getRecordedLogs().length, 0);
    }

    function test_distributeTips_attestationNotaryDispute() public checkQueueLength(2) {
        // rcptNotary: agents[1], attNotary: agents[0]
        prepareTwoReceiptTest(1, 0);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Put DOMAIN_REMOTE agents[0] in Dispute
        openDispute({guard: domains[0].agent, notary: domains[DOMAIN_REMOTE].agents[0]});
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_attestationNotaryFraudulent() public checkQueueLength(1) {
        // rcptNotary: agents[1], attNotary: agents[0]
        address attNotary = prepareTwoReceiptTest(1, 0);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set attNotary status to Fraudulent
        vm.prank(originSynapse);
        bondingManager.slashAgentExposed(DOMAIN_REMOTE, attNotary, address(0));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_attestationNotarySlashed() public checkQueueLength(1) {
        // rcptNotary: agents[1], attNotary: agents[0]
        address attNotary = prepareTwoReceiptTest(1, 0);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set attNotary status to Slashed
        vm.prank(originSynapse);
        bondingManager.slashAgentExposed(DOMAIN_REMOTE, attNotary, address(0));
        bondingManager.completeSlashing(DOMAIN_REMOTE, attNotary, bondingManager.getProof(attNotary));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_receiptNotaryDispute() public checkQueueLength(2) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Put DOMAIN_REMOTE agents[0] in Dispute
        openDispute({guard: domains[0].agent, notary: domains[DOMAIN_REMOTE].agents[0]});
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    /// @dev When Notary in "Dispute Won Timeout" mode, the receipt is moved to the end of the queue,
    /// therefore two receipts in the queue are expected in the modifier.
    function test_distributeTips_receiptNotary_wonDisputeTimeout() public checkQueueLength(2) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        prepareNotaryInDisputeTest();
        skip(BONDING_OPTIMISTIC_PERIOD);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    /// @dev Should distribute tips if the "dispute won" timeout is over.
    function test_distributeTips_receiptNotary_afterNotaryDisputeTimeout() public checkQueueLength(1) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        prepareNotaryInDisputeTest();
        skip(BONDING_OPTIMISTIC_PERIOD);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        vm.recordLogs();
        assertTrue(InterfaceSummit(summit).distributeTips());
        // Should end up emitting TipsAwarded logs (TODO: better way to check this)
        assertGt(vm.getRecordedLogs().length, 0);
    }

    function test_distributeTips_receiptNotaryFraudulent() public checkQueueLength(1) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set rcptNotary status to Fraudulent
        vm.prank(originSynapse);
        bondingManager.slashAgentExposed(DOMAIN_REMOTE, rcptNotary, address(0));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_receiptNotarySlashed() public checkQueueLength(1) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set rcptNotary status to Slashed
        vm.prank(originSynapse);
        bondingManager.slashAgentExposed(DOMAIN_REMOTE, rcptNotary, address(0));
        bondingManager.completeSlashing(DOMAIN_REMOTE, rcptNotary, bondingManager.getProof(rcptNotary));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function prepareTwoReceiptTest(uint256 rcptNotaryIndex, uint256 attNotaryIndex)
        public
        returns (address attNotary)
    {
        (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp) = mockReceipt("First");
        test_submitReceipt(re, tips, rtp, false, rcptNotaryIndex, attNotaryIndex, false);
        re.messageHash = keccak256("Second");
        test_submitReceipt(re, tips, rtp, false, rcptNotaryIndex, attNotaryIndex, false);
        attNotary = re.attNotary;
    }

    function checkAwardedTips(RawExecReceipt memory re, RawTips memory tips, bool isFinal) public {
        logTips(tips);
        checkSnapshotTips(re, tips);
        uint64 receiptTipFull = splitTip({tip: tips.summitTip, parts: 3, roundUp: true});
        uint64 receiptTipFirst = splitTip({tip: receiptTipFull, parts: 2, roundUp: false});
        uint64 receiptTipFinal = splitTip({tip: receiptTipFull, parts: 2, roundUp: true});
        if (rcptNotary == rcptNotaryFinal) {
            if (rcptNotary == re.attNotary) {
                // rcptNotary == rcptNotaryFinal == attNotary
                checkEarnedActorTips(rcptNotary, re.origin, receiptTipFirst + receiptTipFinal + tips.attestationTip);
            } else {
                // rcptNotary == rcptNotaryFinal != attNotary
                checkEarnedActorTips(rcptNotary, re.origin, receiptTipFirst + receiptTipFinal);
                checkEarnedActorTips(re.attNotary, re.origin, tips.attestationTip);
            }
        } else if (re.attNotary == rcptNotaryFinal) {
            // rcptNotaryFinal == attNotary != rcptNotary
            checkEarnedActorTips(rcptNotary, re.origin, receiptTipFirst);
            checkEarnedActorTips(re.attNotary, re.origin, receiptTipFinal + tips.attestationTip);
        } else {
            if (rcptNotary == re.attNotary) {
                // rcptNotary == attNotary != rcptNotaryFinal
                checkEarnedActorTips(rcptNotary, re.origin, receiptTipFirst + tips.attestationTip);
            } else {
                // rcptNotary != attNotary != rcptNotaryFinal
                checkEarnedActorTips(rcptNotary, re.origin, receiptTipFirst);
                checkEarnedActorTips(re.attNotary, re.origin, tips.attestationTip);
            }
            if (isFinal) checkEarnedActorTips(rcptNotaryFinal, re.origin, receiptTipFinal);
        }
        // Check non-bonded actors
        if (re.firstExecutor == re.finalExecutor) {
            checkEarnedActorTips(re.firstExecutor, re.origin, tips.executionTip + (isFinal ? tips.deliveryTip : 0));
        } else {
            checkEarnedActorTips(re.firstExecutor, re.origin, tips.executionTip);
            if (isFinal) checkEarnedActorTips(re.finalExecutor, re.origin, tips.deliveryTip);
        }
    }

    function checkSnapshotTips(RawExecReceipt memory re, RawTips memory tips) public {
        uint64 snapshotTip = splitTip({tip: tips.summitTip, parts: 3, roundUp: false});
        if (re.origin == origin0) {
            // Tips for origin0 go to guard0 and notary0 (they were first to use it),
            // regardless of what attestation was used
            checkEarnedActorTips(guard0, re.origin, snapshotTip);
            checkEarnedActorTips(snapNotary0, re.origin, snapshotTip);
        } else if (re.origin == origin1) {
            // Tips for origin1 go to guard1 and notary1 (they were first to use it)
            checkEarnedActorTips(guard1, re.origin, snapshotTip);
            checkEarnedActorTips(snapNotary1, re.origin, snapshotTip);
        } else {
            revert("Incorrect origin value");
        }
    }

    function expectAwardedTipsEvents(RawExecReceipt memory re, RawTips memory tips, bool isFirst, bool isFinal)
        public
    {
        if (isFirst) {
            expectAwardedTipsEventsFirstSubmit(re, tips);
        }
        expectAwardedTipsEventsReceiptTips(re, tips, isFirst, isFinal);
        if (isFinal) {
            expectAwardedTipsEventsFinalSubmit(re, tips);
        }
    }

    function expectAwardedTipsEventsFirstSubmit(RawExecReceipt memory re, RawTips memory tips) public {
        // In the first submit, tips are awarded to
        // 1. The Guard and Notary who submitted snapshot (summit tips)
        // 2. Notary who submitted attestation (attestation tips)
        // 3. First Executor (execution tips)
        uint64 snapshotTip = splitTip({tip: tips.summitTip, parts: 3, roundUp: false});
        if (re.origin == origin0) {
            // Tips for origin0 go to guard0 and notary0 (they were first to use it),
            // regardless of what attestation was used
            expectAwardedTipsEvent(guard0, re.origin, snapshotTip);
            expectAwardedTipsEvent(snapNotary0, re.origin, snapshotTip);
        } else if (re.origin == origin1) {
            // Tips for origin1 go to guard1 and notary1 (they were first to use it)
            expectAwardedTipsEvent(guard1, re.origin, snapshotTip);
            expectAwardedTipsEvent(snapNotary1, re.origin, snapshotTip);
        } else {
            revert("Incorrect origin value");
        }
        expectAwardedTipsEvent(re.attNotary, re.origin, tips.attestationTip);
        expectAwardedTipsEvent(re.firstExecutor, re.origin, tips.executionTip);
    }

    function expectAwardedTipsEventsReceiptTips(
        RawExecReceipt memory re,
        RawTips memory tips,
        bool isFirst,
        bool isFinal
    ) public {
        // The receipt tips are awarded to Notary who submitted receipt
        // The receipt tips are 1/3 of the snapshot tips (rounded up). These are then split:
        // - First half (rounded down) of receipt  tips are awarded to Receipt Notary who submitted first receipt
        // - Second half (rounded up) of receipt tips are awarded to Receipt Notary who submitted final receipt
        uint64 receiptTipFull = splitTip({tip: tips.summitTip, parts: 3, roundUp: true});
        uint64 receiptTip;
        address notary = rcptNotary;
        if (isFirst && isFinal) {
            receiptTip = receiptTipFull;
        } else if (isFirst) {
            receiptTip = splitTip({tip: receiptTipFull, parts: 2, roundUp: false});
        } else if (isFinal) {
            receiptTip = splitTip({tip: receiptTipFull, parts: 2, roundUp: true});
            notary = rcptNotaryFinal;
        } else {
            revert("Incorrect isFirst and isFinal values");
        }
        expectAwardedTipsEvent(notary, re.origin, receiptTip);
    }

    function expectAwardedTipsEventsFinalSubmit(RawExecReceipt memory re, RawTips memory tips) public {
        // In the final submit (successful execution) tips are awarded to
        // 1. Final Executor (delivery tips)
        expectAwardedTipsEvent(re.finalExecutor, re.origin, tips.deliveryTip);
    }

    function expectAwardedTipsEvent(address actor, uint32 origin_, uint128 earnedScaledDown) public {
        uint256 earnedTips = 2 ** 32 * earnedScaledDown;
        vm.expectEmit(summit);
        emit TipAwarded(actor, origin_, earnedTips);
    }

    /// @dev We calculate the "scaled down" version of earned tips, i.e. divided by 2^32
    /// `Summit` is supposed to store the full value, so we scale calculated value up by 2^32.
    function checkEarnedActorTips(address actor, uint32 origin_, uint128 earnedScaledDown) public {
        (uint256 earnedTips,) = InterfaceSummit(summit).actorTips(actor, origin_);
        assertEq(earnedTips, 2 ** 32 * earnedScaledDown);
    }

    function logTips(RawTips memory tips) public {
        emit log_named_uint("Summit tip", tips.summitTip);
        emit log_named_uint("Attestation tip", tips.attestationTip);
        emit log_named_uint("Execution tip", tips.executionTip);
        emit log_named_uint("Delivery tip", tips.deliveryTip);
    }

    function splitTip(uint64 tip, uint64 parts, bool roundUp) public pure returns (uint64) {
        return tip / parts + (roundUp ? tip % parts : 0);
    }

    // ═══════════════════════════════════════════ TESTS: WITHDRAW TIPS ════════════════════════════════════════════════

    function test_withdrawTips(address actor, uint32 domain, uint128 earned, uint128 claimed, uint128 amount) public {
        // Etch the contract with cheat codes to set actor tips
        vm.etch(summit, summitCheats.code);
        earned = uint128(bound(earned, 1, type(uint128).max));
        claimed = claimed % earned;
        amount = uint128(bound(amount, 1, earned - claimed));
        SummitCheats(summit).setActorTips(actor, domain, earned, claimed);
        bytes memory expectedCall = abi.encodeWithSelector(bondingManager.withdrawTips.selector, actor, domain, amount);
        vm.expectCall(address(bondingManager), expectedCall);
        vm.expectEmit(summit);
        emit TipWithdrawalInitiated(actor, domain, amount);
        vm.prank(actor);
        InterfaceSummit(summit).withdrawTips(domain, amount);
        (uint128 earned_, uint128 claimed_) = InterfaceSummit(summit).actorTips(actor, domain);
        assertEq(earned_, earned, "!earned");
        assertEq(claimed_, claimed + amount, "!claimed");
    }

    function test_withdrawTips_revert_zeroAmount(address actor, uint32 domain) public {
        vm.expectRevert(TipsClaimZero.selector);
        vm.prank(actor);
        InterfaceSummit(summit).withdrawTips(domain, 0);
    }

    function test_withdrawTips_revert_tipsBalanceTooLow(
        address actor,
        uint32 domain,
        uint128 earned,
        uint128 claimed,
        uint128 amount
    ) public {
        earned = uint128(bound(earned, 1, type(uint64).max));
        claimed = claimed % earned;
        amount = uint128(bound(amount, 1, earned - claimed));
        amount = uint128(bound(amount, earned - claimed + 1, type(uint128).max));
        vm.expectRevert(TipsClaimMoreThanEarned.selector);
        vm.prank(actor);
        InterfaceSummit(summit).withdrawTips(domain, amount);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function prepareReceipt(
        RawExecReceipt memory re,
        RawTips memory tips,
        RawTipsProof memory rtp,
        bool originZero,
        uint256 attNotaryIndex,
        bool isSuccess
    ) public {
        if (originZero) {
            // For Origin0's state0 we could use (state0) or (state0, state1) attestations
            re.origin = origin0;
            re.snapshotRoot = attNotaryIndex % 2 == 0 ? snapRoot0 : snapRoot1;
            re.stateIndex = 0;
        } else {
            // For Origin1's state1 we could only use (state0, state1) attestation
            re.origin = origin1;
            re.snapshotRoot = snapRoot1;
            re.stateIndex = 1;
        }
        re.destination = DOMAIN_REMOTE;
        re.attNotary = domains[DOMAIN_REMOTE].agents[attNotaryIndex % DOMAIN_AGENTS];
        re.firstExecutor = createExecutorEOA(re.firstExecutor, "First Executor");
        if (isSuccess) {
            re.finalExecutor = createExecutorEOA(re.finalExecutor, "Final Executor");
        } else {
            re.finalExecutor = address(0);
        }
        // Make every tip component non-zero and not too big
        tips.boundTips(type(uint32).max);
        tips.floorTips(1);
        // Make sure tips could be proven against mocked message hash
        re.messageHash = tips.getMessageHash(rtp);
    }

    function mockReceipt(bytes memory salt)
        public
        pure
        returns (RawExecReceipt memory re, RawTips memory tips, RawTipsProof memory rtp)
    {
        rtp.headerHash = keccak256(abi.encodePacked(salt, "header"));
        rtp.bodyHash = keccak256(abi.encodePacked(salt, "body"));
        // Leave everything else as zero, prepareReceipt() takes care of that
        re;
        tips;
    }

    /// @notice Creates an EOA address that should not collide with existing agents
    function createExecutorEOA(address addr, string memory label) public returns (address eoa) {
        bytes32 addrHash = keccak256(abi.encode(addr));
        eoa = address(uint160(uint256(addrHash)));
        vm.label(eoa, label);
    }

    function submitGuardSnapshot(address guard, RawState memory rs) public {
        RawSnapshot memory rawSnap;
        rawSnap.states = new RawState[](1);
        rawSnap.states[0] = rs;
        submitSnapshot(guard, rawSnap);
    }

    function submitSnapshot(address agent, RawSnapshot memory rawSnap) public {
        (bytes memory snapPayload, bytes memory snapSignature) = signSnapshot(agent, rawSnap);
        inbox.submitSnapshot(snapPayload, snapSignature);
    }

    // ═════════════════════════════════════════════════ OVERRIDES ═════════════════════════════════════════════════════

    function localContract() public view override returns (address) {
        return summit;
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
