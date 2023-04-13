// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentManager, InterfaceSummit} from "../../contracts/Summit.sol";

import {AgentFlag, ISystemContract, Summit, SynapseTest} from "../utils/SynapseTest.t.sol";
import {IDisputeHub, DisputeHubTest} from "./hubs/DisputeHub.t.sol";

import {fakeState} from "../utils/libs/FakeIt.t.sol";
import {RawExecReceipt, RawState, RawStateIndex, RawSnapshot, RawTips} from "../utils/libs/SynapseStructs.t.sol";

import {stdStorage, StdStorage} from "forge-std/Test.sol";

contract SummitCheats is Summit {
    constructor(uint32 domain, IAgentManager agentManager_) Summit(domain, agentManager_) {}

    function setActorTips(address actor, uint32 origin, uint128 earned, uint128 claimed) external {
        actorTips[actor][origin].earned = earned;
        actorTips[actor][origin].claimed = claimed;
    }
}

// solhint-disable code-complexity
// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract SummitTipsTest is DisputeHubTest {
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

    // Deploy Production version of Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_SUMMIT) {}

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
        acceptSnapshot(snapshot0.formatStates());
        snapRoot0 = getSnapshotRoot();
        acceptSnapshot(snapshot1.formatStates());
        snapRoot1 = getSnapshotRoot();
        // Deploy Summit implementation with Cheats
        summitCheats = address(new SummitCheats(DOMAIN_SYNAPSE, bondingManager));
    }

    // ══════════════════════════════════════════ TESTS: SUBMIT RECEIPTS ═══════════════════════════════════════════════

    function test_submitReceipt(
        RawExecReceipt memory re,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex,
        bool isSuccess
    ) public checkQueueLength(1) {
        prepareReceipt(re, originZero, attNotaryIndex, isSuccess);
        rcptNotary = domains[DOMAIN_REMOTE].agents[rcptNotaryIndex % DOMAIN_AGENTS];
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotary, re);
        vm.expectEmit();
        emit ReceiptAccepted(DOMAIN_REMOTE, rcptNotary, rcptPayload, rcptSignature);
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_notAccepted_pending() public checkQueueLength(1) {
        RawExecReceipt memory re = mockReceipt("First");
        test_submitReceipt(re, false, 0, 0, false);
        re.finalExecutor = createExecutorEOA(re.finalExecutor, "Final Executor");
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotary, re);
        vm.recordLogs();
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function test_submitReceipt_notAccepted_outdatedStatus() public checkQueueLength(0) {
        RawExecReceipt memory re = mockReceipt("First");
        test_distributeTips_success(re, false, 0, 0);
        re.finalExecutor = address(0);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotary, re);
        vm.recordLogs();
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function test_submitReceipt_revert_signedByGuard() public {
        RawExecReceipt memory re = mockReceipt("First");
        prepareReceipt(re, false, 0, false);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(guard0, re);
        vm.expectRevert("Signer is not a Notary");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_revert_wrongNotaryDomain() public {
        // TODO: remove when Notary restrictions are revisited
        RawExecReceipt memory re = mockReceipt("First");
        prepareReceipt(re, false, 0, false);
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert("Wrong Notary domain");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_revert_notaryInDispute() public {
        RawExecReceipt memory re = mockReceipt("First");
        prepareReceipt(re, false, 0, false);
        // Put DOMAIN_REMOTE notary in Dispute
        check_submitStateReport(summit, DOMAIN_REMOTE, state0, RawStateIndex(0, 1));
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert("Notary is in dispute");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_revert_unknownSnapRoot() public {
        RawExecReceipt memory re = mockReceipt("First");
        prepareReceipt(re, false, 0, false);
        re.snapshotRoot = "Some fake snapshot root";
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert("Unknown snapshot root");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    // ═══════════════════════════════════════════ TESTS: TIPS AWARDING ════════════════════════════════════════════════

    function test_distributeTips_success(
        RawExecReceipt memory re,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex
    ) public checkQueueLength(0) {
        test_submitReceipt(re, originZero, rcptNotaryIndex, attNotaryIndex, true);
        skip(BONDING_OPTIMISTIC_PERIOD);
        assertTrue(InterfaceSummit(summit).distributeTips());
        rcptNotaryFinal = rcptNotary;
        checkAwardedTips(re, true);
    }

    function test_distributeTips_failed(
        RawExecReceipt memory re,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex
    ) public checkQueueLength(0) {
        test_submitReceipt(re, originZero, rcptNotaryIndex, attNotaryIndex, false);
        skip(BONDING_OPTIMISTIC_PERIOD);
        assertTrue(InterfaceSummit(summit).distributeTips());
        rcptNotaryFinal = address(0);
        checkAwardedTips(re, false);
    }

    function test_distributeTips_failedThenSuccess(
        RawExecReceipt memory re,
        bool originZero,
        uint256 rcptNotaryIndex,
        uint256 attNotaryIndex,
        uint256 rcptNotaryIndexFinal,
        address finalExecutor
    ) public checkQueueLength(0) {
        test_distributeTips_failed(re, originZero, rcptNotaryIndex, attNotaryIndex);
        re.finalExecutor = createExecutorEOA(finalExecutor, "Final Executor");
        rcptNotaryFinal = domains[DOMAIN_REMOTE].agents[rcptNotaryIndexFinal % DOMAIN_AGENTS];
        emit log_named_address("Receipt Notary", rcptNotaryFinal);
        emit log_named_address("Attestation Notary", re.attNotary);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(rcptNotaryFinal, re);
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
        skip(BONDING_OPTIMISTIC_PERIOD);
        assertTrue(InterfaceSummit(summit).distributeTips());
        checkAwardedTips(re, true);
    }

    function test_distributeTips_emptyQueue() public checkQueueLength(0) {
        RawExecReceipt memory re = mockReceipt("First");
        test_distributeTips_success(re, true, 0, 0);
        assertFalse(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_optimisticPeriodNotOver(uint256 timePassed) public checkQueueLength(1) {
        RawExecReceipt memory re = mockReceipt("First");
        test_submitReceipt(re, false, 0, 0, false);
        timePassed = timePassed % BONDING_OPTIMISTIC_PERIOD;
        skip(timePassed);
        assertFalse(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_attestationNotaryDispute() public checkQueueLength(2) {
        // rcptNotary: agents[1], attNotary: agents[0]
        prepareTwoReceiptTest(1, 0);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Put DOMAIN_REMOTE agents[0] in Dispute
        check_submitStateReport(summit, DOMAIN_REMOTE, state0, RawStateIndex(0, 1));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_attestationNotaryFraudulent() public checkQueueLength(1) {
        // rcptNotary: agents[1], attNotary: agents[0]
        address attNotary = prepareTwoReceiptTest(1, 0);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set attNotary status to Fraudulent
        vm.prank(originSynapse);
        bondingManager.registrySlash(DOMAIN_REMOTE, attNotary, address(0));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_attestationNotarySlashed() public checkQueueLength(1) {
        // rcptNotary: agents[1], attNotary: agents[0]
        address attNotary = prepareTwoReceiptTest(1, 0);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set attNotary status to Slashed
        vm.prank(originSynapse);
        bondingManager.registrySlash(DOMAIN_REMOTE, attNotary, address(0));
        bondingManager.completeSlashing(DOMAIN_REMOTE, attNotary, bondingManager.getProof(attNotary));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_receiptNotaryDispute() public checkQueueLength(2) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Put DOMAIN_REMOTE agents[0] in Dispute
        check_submitStateReport(summit, DOMAIN_REMOTE, state0, RawStateIndex(0, 1));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_receiptNotaryFraudulent() public checkQueueLength(1) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set rcptNotary status to Fraudulent
        vm.prank(originSynapse);
        bondingManager.registrySlash(DOMAIN_REMOTE, rcptNotary, address(0));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function test_distributeTips_receiptNotarySlashed() public checkQueueLength(1) {
        // rcptNotary: agents[0], attNotary: agents[1]
        prepareTwoReceiptTest(0, 1);
        skip(BONDING_OPTIMISTIC_PERIOD);
        // Set rcptNotary status to Slashed
        vm.prank(originSynapse);
        bondingManager.registrySlash(DOMAIN_REMOTE, rcptNotary, address(0));
        bondingManager.completeSlashing(DOMAIN_REMOTE, rcptNotary, bondingManager.getProof(rcptNotary));
        assertTrue(InterfaceSummit(summit).distributeTips());
    }

    function prepareTwoReceiptTest(uint256 rcptNotaryIndex, uint256 attNotaryIndex)
        public
        returns (address attNotary)
    {
        RawExecReceipt memory re = mockReceipt("First");
        test_submitReceipt(re, false, rcptNotaryIndex, attNotaryIndex, false);
        re.messageHash = keccak256("Second");
        test_submitReceipt(re, false, rcptNotaryIndex, attNotaryIndex, false);
        attNotary = re.attNotary;
    }

    function checkAwardedTips(RawExecReceipt memory re, bool isFinal) public {
        logTips(re.tips);
        checkSnapshotTips(re);
        uint64 receiptTipFull = splitTip({tip: re.tips.summitTip, parts: 3, roundUp: true});
        uint64 receiptTipFirst = splitTip({tip: receiptTipFull, parts: 2, roundUp: false});
        uint64 receiptTipFinal = splitTip({tip: receiptTipFull, parts: 2, roundUp: true});
        if (rcptNotary == rcptNotaryFinal) {
            if (rcptNotary == re.attNotary) {
                // rcptNotary == rcptNotaryFinal == attNotary
                checkActorTips(rcptNotary, re.origin, receiptTipFirst + receiptTipFinal + re.tips.attestationTip, 0);
            } else {
                // rcptNotary == rcptNotaryFinal != attNotary
                checkActorTips(rcptNotary, re.origin, receiptTipFirst + receiptTipFinal, 0);
                checkActorTips(re.attNotary, re.origin, re.tips.attestationTip, 0);
            }
        } else if (re.attNotary == rcptNotaryFinal) {
            // rcptNotaryFinal == attNotary != rcptNotary
            checkActorTips(rcptNotary, re.origin, receiptTipFirst, 0);
            checkActorTips(re.attNotary, re.origin, receiptTipFinal + re.tips.attestationTip, 0);
        } else {
            if (rcptNotary == re.attNotary) {
                // rcptNotary == attNotary != rcptNotaryFinal
                checkActorTips(rcptNotary, re.origin, receiptTipFirst + re.tips.attestationTip, 0);
            } else {
                // rcptNotary != attNotary != rcptNotaryFinal
                checkActorTips(rcptNotary, re.origin, receiptTipFirst, 0);
                checkActorTips(re.attNotary, re.origin, re.tips.attestationTip, 0);
            }
            if (isFinal) checkActorTips(rcptNotaryFinal, re.origin, receiptTipFinal, 0);
        }
        // Check non-bonded actors
        if (re.firstExecutor == re.finalExecutor) {
            checkActorTips(re.firstExecutor, re.origin, re.tips.executionTip + (isFinal ? re.tips.deliveryTip : 0), 0);
        } else {
            checkActorTips(re.firstExecutor, re.origin, re.tips.executionTip, 0);
            if (isFinal) checkActorTips(re.finalExecutor, re.origin, re.tips.deliveryTip, 0);
        }
    }

    function checkSnapshotTips(RawExecReceipt memory re) public {
        uint64 snapshotTip = splitTip({tip: re.tips.summitTip, parts: 3, roundUp: false});
        if (re.origin == origin0) {
            // Tips for origin0 go to guard0 and notary0 (they were first to use it),
            // regardless of what attestation was used
            checkActorTips(guard0, re.origin, snapshotTip, 0);
            checkActorTips(snapNotary0, re.origin, snapshotTip, 0);
        } else if (re.origin == origin1) {
            // Tips for origin1 go to guard1 and notary1 (they were first to use it)
            checkActorTips(guard1, re.origin, snapshotTip, 0);
            checkActorTips(snapNotary1, re.origin, snapshotTip, 0);
        } else {
            revert("Incorrect origin value");
        }
    }

    function checkActorTips(address actor, uint32 origin_, uint128 earned, uint128 claimed) public {
        (uint128 earned_, uint128 claimed_) = InterfaceSummit(summit).actorTips(actor, origin_);
        assertEq(earned_, earned, "!earned");
        assertEq(claimed_, claimed, "!claimed");
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
        vm.prank(actor);
        InterfaceSummit(summit).withdrawTips(domain, amount);
        (uint128 earned_, uint128 claimed_) = InterfaceSummit(summit).actorTips(actor, domain);
        assertEq(earned_, earned, "!earned");
        assertEq(claimed_, claimed + amount, "!claimed");
    }

    function test_withdrawTips_revert_zeroAmount(address actor, uint32 domain) public {
        vm.expectRevert("Amount is zero");
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
        vm.expectRevert("Tips balance too low");
        vm.prank(actor);
        InterfaceSummit(summit).withdrawTips(domain, amount);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function prepareReceipt(RawExecReceipt memory re, bool originZero, uint256 attNotaryIndex, bool isSuccess) public {
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
        re.tips.boundTips(type(uint32).max);
        re.tips.floorTips(1);
    }

    function mockReceipt(bytes memory salt) public pure returns (RawExecReceipt memory re) {
        re.messageHash = keccak256(salt);
        // Leave everything else as zero, prepareReceipt() takes care of that
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
        InterfaceSummit(summit).submitSnapshot(snapPayload, snapSignature);
    }

    // ═════════════════════════════════════════════════ OVERRIDES ═════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested system contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
