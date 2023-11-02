// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DISPUTE_TIMEOUT_NOTARY} from "../../../contracts/libs/Constants.sol";
import {
    AlreadyExecuted,
    AlreadyFailed,
    DisputeTimeoutNotOver,
    GasLimitTooLow,
    GasSuppliedTooLow,
    IncorrectDestinationDomain,
    IncorrectOriginDomain,
    IncorrectMagicValue,
    IncorrectSnapshotRoot,
    MessageOptimisticPeriod,
    NotaryInDispute
} from "../../../contracts/libs/Errors.sol";
import {AgentFlag} from "../../../contracts/libs/Structures.sol";
import {IExecutionHub} from "../../../contracts/interfaces/IExecutionHub.sol";
import {IStatementInbox} from "../../../contracts/interfaces/IStatementInbox.sol";
import {SNAPSHOT_MAX_STATES} from "../../../contracts/libs/memory/Snapshot.sol";
import {MessageStatus} from "../../../contracts/libs/Structures.sol";

import {ReentrantApp} from "../../harnesses/client/ReentrantApp.t.sol";
import {RevertingApp} from "../../harnesses/client/RevertingApp.t.sol";
import {MessageRecipientMock} from "../../mocks/client/MessageRecipientMock.t.sol";

import {fakeSnapshot} from "../../utils/libs/FakeIt.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {
    ReceiptLib,
    MessageFlag,
    RawAttestation,
    RawBaseMessage,
    RawCallData,
    RawExecReceipt,
    RawHeader,
    RawMessage,
    RawSnapshot,
    RawState,
    RawStateIndex,
    RawTips
} from "../../utils/libs/SynapseStructs.t.sol";
import {AgentSecuredTest} from "../base/AgentSecured.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract ExecutionHubTest is AgentSecuredTest {
    address internal recipient;
    address internal executor;
    address internal executorNew;

    uint8 internal cachedStateIndex;

    bytes internal msgPayload;
    bytes32 internal msgLeaf;

    function setUp() public virtual override {
        super.setUp();
        recipient = address(new MessageRecipientMock());
        executor = makeAddr("Executor");
        executorNew = makeAddr("Executor New");
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testExecutionHub() external {}

    // ═══════════════════════════════════════ TESTS: EXECUTE BASE MESSAGES ════════════════════════════════════════════

    function test_execute_base(
        RawBaseMessage memory rbm,
        RawHeader memory rh,
        SnapshotMock memory sm,
        uint32 timePassed,
        uint64 gasLimit
    ) public {
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        gasLimit = uint64(bound(gasLimit, rbm.request.gasLimit, 2_000_000));
        // receiveBaseMessage(origin, nonce, sender, proofMaturity, version, message)
        bytes memory expectedCall = abi.encodeWithSelector(
            MessageRecipientMock.receiveBaseMessage.selector,
            rh.origin,
            rh.nonce,
            rbm.sender,
            timePassed,
            rbm.request.version,
            rbm.content
        );
        // expectCall(address callee, uint256 msgValue, uint64 gas, bytes calldata data)
        vm.expectCall(recipient, 0, gasLimit, expectedCall);
        vm.expectEmit();
        emit Executed(rh.origin, msgLeaf, true);
        vm.prank(executor);
        msgLeaf = rm.castToMessage().leaf();
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, gasLimit);
        bytes memory rcptPayload =
            verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Success, executor, executor);
        verify_receipt_valid(rcptPayload);
    }

    function test_execute_base_recipientEOA(Random memory random) public {
        recipient = random.nextAddress();
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectEmit();
        emit Executed(rh.origin, msgLeaf, false);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        bytes memory rcptPayloadFirst =
            verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Failed, executor, address(0));
        verify_receipt_valid(rcptPayloadFirst);
    }

    function test_execute_base_recipientReverted_thenSuccess(Random memory random) public {
        recipient = address(new RevertingApp());
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectEmit();
        emit Executed(rh.origin, msgLeaf, false);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        bytes memory rcptPayloadFirst =
            verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Failed, executor, address(0));
        verify_receipt_valid(rcptPayloadFirst);
        // Retry the same failed message
        RevertingApp(payable(recipient)).toggleRevert(false);
        vm.expectEmit();
        emit Executed(rh.origin, msgLeaf, true);
        vm.prank(executorNew);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        bytes memory rcptPayloadSecond =
            verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Success, executor, executorNew);
        // Both receipts (historical and current) should be valid
        verify_receipt_valid(rcptPayloadFirst);
        verify_receipt_valid(rcptPayloadSecond);
        cachedStateIndex = uint8(sm.rsi.stateIndex);
    }

    function test_execute_base_recipientReverted_twice(Random memory random) public {
        recipient = address(new RevertingApp());
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectEmit();
        emit Executed(rh.origin, msgLeaf, false);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        bytes memory rcptPayload =
            verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Failed, executor, address(0));
        verify_receipt_valid(rcptPayload);
        // Retry the same failed message
        vm.expectRevert(AlreadyFailed.selector);
        vm.prank(executorNew);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
    }

    function test_execute_base_revert_alreadyExecuted(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        vm.expectRevert(AlreadyExecuted.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
    }

    function test_execute_base_revert_notaryInDispute(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // initiate dispute
        openDispute({guard: domains[0].agent, notary: domains[DOMAIN_LOCAL].agent});
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert(NotaryInDispute.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(msgLeaf);
    }

    function test_execute_base_revert_notaryWonDisputeTimeout() public {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;
        Random memory random = Random("Random Salt 1");
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // initiate dispute
        openTestDispute({guardIndex: agentIndex[guard], notaryIndex: agentIndex[notary]});
        // Make sure that optimistic period is over
        skip(rh.optimisticPeriod);
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        vm.expectRevert(DisputeTimeoutNotOver.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
    }

    function test_execute_base_afterNotaryDisputeTimeout() public {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;
        Random memory random = Random("Random Salt 2");
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // initiate dispute
        openTestDispute({guardIndex: agentIndex[guard], notaryIndex: agentIndex[notary]});
        // Make sure that optimistic period is over
        skip(rh.optimisticPeriod);
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
        skip(DISPUTE_TIMEOUT_NOTARY);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Success, executor, executor);
    }

    function test_execute_base_revert_snapRootUnknown(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = createSnapshotProof(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert(IncorrectSnapshotRoot.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(msgLeaf);
    }

    function test_execute_base_revert_optimisticPeriodNotOver(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        vm.assume(rh.optimisticPeriod != 0);
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is NOT over
        uint32 timePassed = random.nextUint32() % rh.optimisticPeriod;
        skip(timePassed);
        vm.expectRevert(MessageOptimisticPeriod.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(msgLeaf);
    }

    function test_execute_base_revert_optimisticPeriodMinus1Second() public {
        Random memory random = Random("random salt");
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        vm.assume(rh.optimisticPeriod != 0);
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is NOT over
        uint32 timePassed = rh.optimisticPeriod - 1;
        skip(timePassed);
        vm.expectRevert(MessageOptimisticPeriod.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
    }

    function test_execute_base_revert_gasLimitTooLow(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        // Make sure gas limit is lower than requested
        uint64 gasLimit = random.nextUint64() % rbm.request.gasLimit;
        vm.expectRevert(GasLimitTooLow.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, gasLimit);
        verify_messageStatusNone(msgLeaf);
    }

    function test_execute_base_revert_gasSuppliedTooLow(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert(GasSuppliedTooLow.selector);
        vm.prank(executor);
        // Limit amount of gas for the whole call
        testedEH().execute{gas: rbm.request.gasLimit + 20_000}(
            msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit
        );
        verify_messageStatusNone(msgLeaf);
    }

    function test_execute_base_revert_wrongDestination(Random memory random, uint32 destination_) public {
        vm.assume(destination_ != localDomain());
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, destination_);
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert(IncorrectDestinationDomain.selector);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(msgLeaf);
    }

    function test_execute_base_revert_originSameDomain() public {
        RawBaseMessage memory rbm;
        rbm.sender = rbm.recipient = addressToBytes32(recipient);
        RawMessage memory rm;
        rm.header.flag = uint8(MessageFlag.Base);
        rm.header.origin = localDomain();
        rm.header.destination = localDomain();
        rm.header.nonce = 1;
        rm.header.optimisticPeriod = 1 seconds;
        rm.body = rbm.formatBaseMessage();
        msgPayload = rm.formatMessage();
        vm.expectRevert(IncorrectOriginDomain.selector);
        testedEH().execute(msgPayload, new bytes32[](0), new bytes32[](0), 0, 0);
    }

    function test_execute_base_revert_reentrancy(Random memory random) public {
        recipient = address(new ReentrantApp());
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        RawMessage memory rm = createBaseMessages(rbm, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        ReentrantApp(recipient).prepare(msgPayload, originProof, snapProof, sm.rsi.stateIndex);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        // We allow message recipient to revert, which is what supposed to happen if they try to reenter
        verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Failed, executor, address(0));
    }

    // ══════════════════════════════════════ TESTS: EXECUTE MANAGER MESSAGES ══════════════════════════════════════════

    function test_execute_manager(RawHeader memory rh, SnapshotMock memory sm, uint32 timePassed, uint64 gasLimit)
        public
    {
        // Create messages and get origin proof
        RawMessage memory rm = createManagerMessages(lightManager.remoteMockFunc.selector, rh, localDomain());
        msgPayload = rm.formatMessage();
        msgLeaf = rm.castToMessage().leaf();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        // expectCall(address callee, bytes calldata data)
        // remoteMockFunc(msgOrigin, proofMaturity, data); data = rh.nonce
        vm.expectCall(
            localAgentManager(),
            abi.encodeWithSelector(bondingManager.remoteMockFunc.selector, rh.origin, timePassed, rh.nonce)
        );
        vm.expectEmit();
        emit Executed(rh.origin, msgLeaf, true);
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, gasLimit);
        verify_messageStatus(msgLeaf, snapRoot, sm.rsi.stateIndex, MessageStatus.Success, executor, executor);
    }

    function test_execute_manager_revert_incorrectMagicValue() public {
        // Use empty values - these would be filled later in test preparation
        RawHeader memory rh;
        SnapshotMock memory sm;
        // Create messages and get origin proof
        RawMessage memory rm = createManagerMessages(lightManager.sensitiveMockFunc.selector, rh, localDomain());
        msgPayload = rm.formatMessage();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        vm.expectRevert(IncorrectMagicValue.selector);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, 0);
    }

    function test_execute_manager_revert_noMagicValue() public {
        // Use empty values - these would be filled later in test preparation
        RawHeader memory rh;
        SnapshotMock memory sm;
        // Create messages and get origin proof
        RawMessage memory rm = createManagerMessages(lightManager.sensitiveMockFuncVoid.selector, rh, localDomain());
        msgPayload = rm.formatMessage();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        vm.expectRevert(IncorrectMagicValue.selector);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, 0);
    }

    function test_execute_manager_revert_magicMoreThan32Bytes() public {
        // Use empty values - these would be filled later in test preparation
        RawHeader memory rh;
        SnapshotMock memory sm;
        // Create messages and get origin proof
        RawMessage memory rm =
            createManagerMessages(lightManager.sensitiveMockFuncOver32Bytes.selector, rh, localDomain());
        msgPayload = rm.formatMessage();
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        vm.expectRevert(IncorrectMagicValue.selector);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, 0);
    }

    // ══════════════════════════════════════════ TESTS: INVALID RECEIPTS ══════════════════════════════════════════════

    function test_verifyReceipt_invalid_msgStatusNone(RawExecReceipt memory re) public {
        vm.assume(testedEH().messageStatus(re.messageHash) == MessageStatus.None);
        vm.assume(re.origin != localDomain());
        re.destination = localDomain();
        verify_receipt_invalid(re);
    }

    function test_verifyReceipt_invalid_msgStatusSuccess(uint256 mask) public {
        test_execute_base_recipientReverted_thenSuccess(Random(bytes32(mask)));
        RawExecReceipt memory re = RawExecReceipt({
            origin: DOMAIN_REMOTE,
            destination: localDomain(),
            messageHash: getLeaf(0),
            snapshotRoot: getSnapshotRoot(),
            stateIndex: cachedStateIndex,
            attNotary: domains[DOMAIN_LOCAL].agent,
            firstExecutor: executor,
            finalExecutor: executorNew
        });
        // Check that data we start with is valid. Use require() to break the test execution early.
        require(testedEH().isValidReceipt(re.formatReceipt()), "Incorrect initial receipt data");
        RawExecReceipt memory mre = re.modifyReceipt(mask);
        verify_receipt_invalid(mre);
    }

    // ═════════════════════════════════════════════════ VERIFIERS ═════════════════════════════════════════════════════

    function verify_messageStatusNone(bytes32 messageHash) public {
        verify_messageStatus(messageHash, bytes32(0), 0, MessageStatus.None, address(0), address(0));
    }

    function verify_messageStatus(
        bytes32 messageHash,
        bytes32 snapRoot,
        uint8 stateIndex,
        MessageStatus flag,
        address firstExecutor,
        address finalExecutor
    ) public returns (bytes memory rcptPayload) {
        MessageStatus flag_ = testedEH().messageStatus(messageHash);
        assertEq(uint8(flag_), uint8(flag), "!flag");
        rcptPayload = testedEH().messageReceipt(messageHash);
        if (flag == MessageStatus.None) {
            assertEq(rcptPayload.length, 0, "!receipt: empty");
        } else {
            address attNotary = domains[DOMAIN_LOCAL].agent;
            assertEq(
                rcptPayload,
                ReceiptLib.formatReceipt(
                    DOMAIN_REMOTE,
                    localDomain(),
                    messageHash,
                    snapRoot,
                    uint8(stateIndex),
                    attNotary,
                    firstExecutor,
                    finalExecutor
                )
            );
        }
    }

    function verify_receipt_valid(bytes memory rcptPayload) public {
        assertTrue(testedEH().isValidReceipt(rcptPayload));
        address notary = domains[DOMAIN_LOCAL].agent;
        bytes memory rcptSignature = signReceipt(notary, rcptPayload);
        vm.recordLogs();
        assertTrue(IStatementInbox(localInbox()).verifyReceipt(rcptPayload, rcptSignature));
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function verify_receipt_invalid(RawExecReceipt memory re) public {
        bytes memory rcptPayload = re.formatReceipt();
        assertFalse(testedEH().isValidReceipt(rcptPayload));
        address notary = domains[DOMAIN_LOCAL].agent;
        bytes memory rcptSignature = signReceipt(notary, rcptPayload);
        // TODO: check that anyone could make the call
        expectStatusUpdated(AgentFlag.Fraudulent, DOMAIN_LOCAL, notary);
        expectDisputeResolved(0, notary, address(0), address(this));
        // expectAgentSlashed(localDomain(), notary, address(this));
        assertFalse(IStatementInbox(localInbox()).verifyReceipt(rcptPayload, rcptSignature));
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Prepares execution of the created messages
    function prepareExecution(SnapshotMock memory sm)
        public
        virtual
        returns (bytes32 snapRoot, bytes32[] memory snapProof);

    function createBaseMessages(RawBaseMessage memory rbm, RawHeader memory rh, uint32 destination_)
        public
        returns (RawMessage memory rm)
    {
        adjustHeader(rh, destination_);
        rbm.recipient = addressToBytes32(recipient);
        // Set sensible limitations for tips/request
        rbm.tips.boundTips(2 ** 32);
        rbm.request.gasLimit = uint64(bound(rbm.request.gasLimit, 50_000, 200_000));
        rh.flag = uint8(MessageFlag.Base);
        rm = RawMessage(rh, rbm.formatBaseMessage());
        createMessages(rh.nonce, rm);
    }

    function createManagerMessages(bytes4 selector, RawHeader memory rh, uint32 destination_)
        public
        returns (RawMessage memory rm)
    {
        adjustHeader(rh, destination_);
        RawCallData memory rcd = RawCallData({selector: selector, args: abi.encode(rh.nonce)});
        rh.flag = uint8(MessageFlag.Manager);
        rm = RawMessage(rh, rcd.formatCallData());
        createMessages(rh.nonce, rm);
    }

    function createMessages(uint32 msgNonce, RawMessage memory rm) public {
        for (uint32 nonce = 1; nonce <= MESSAGES; ++nonce) {
            if (nonce == msgNonce) {
                insertMessage(rm.castToMessage().leaf());
            } else {
                insertMessage(keccak256(abi.encode("Mocked payload", nonce)));
            }
        }
    }

    /// @notice Creates some simple data for the revert test
    function createDataRevertTest(Random memory random)
        public
        pure
        returns (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm)
    {
        rbm.sender = random.next();
        rbm.content = "Test content";
        rbm.tips = RawTips(1, 1, 1, 1);
        rh.nonce = 1;
        rh.optimisticPeriod = random.nextUint32();
        sm = SnapshotMock(random.nextState(), RawStateIndex(random.nextUint8(), random.nextUint256()));
        sm.rsi.boundStateIndex();
    }

    /// @notice Sets realistic values for the message header
    function adjustHeader(RawHeader memory rh, uint32 destination_) public view {
        rh.origin = DOMAIN_REMOTE;
        rh.nonce = uint32(bound(rh.nonce, 1, MESSAGES));
        rh.destination = destination_;
        rh.optimisticPeriod = rh.optimisticPeriod % 1 days;
    }

    function adjustState(RawState memory rs) public view {
        rs.origin = DOMAIN_REMOTE;
        rs.nonce = MESSAGES;
        rs.root = getRoot(rs.nonce);
    }

    function adjustSnapshot(SnapshotMock memory sm) public view {
        adjustState(sm.rs);
        sm.rsi.boundStateIndex();
    }

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return localDestination();
    }

    /// @notice Returns tested contract as IExecutionHub
    function testedEH() public view returns (IExecutionHub) {
        return IExecutionHub(localContract());
    }
}
