// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IExecutionHub} from "../../../contracts/interfaces/IExecutionHub.sol";
import {SNAPSHOT_MAX_STATES} from "../../../contracts/libs/Snapshot.sol";
import {MessageStatus} from "../../../contracts/libs/Structures.sol";

import {RevertingApp} from "../../harnesses/client/RevertingApp.t.sol";
import {MessageRecipientMock} from "../../mocks/client/MessageRecipientMock.t.sol";
import {ISystemContract, SystemContractMock} from "../../mocks/system/SystemContractMock.t.sol";
import {SystemRouterMock} from "../../mocks/system/SystemRouterMock.t.sol";

import {Random} from "../../utils/libs/Random.t.sol";
import {
    ReceiptLib,
    MessageFlag,
    RawAttestation,
    RawBaseMessage,
    RawReceiptBody,
    RawExecReceipt,
    RawHeader,
    RawMessage,
    RawState,
    RawStateIndex,
    RawSystemMessage,
    RawTips
} from "../../utils/libs/SynapseStructs.t.sol";
import {DisputeHubTest, IDisputeHub} from "./DisputeHub.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract ExecutionHubTest is DisputeHubTest {
    struct SnapshotMock {
        RawState rs;
        RawStateIndex rsi;
    }

    address internal recipient;
    address internal executor;
    address internal executorNew;

    uint8 internal cachedStateIndex;

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
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        gasLimit = uint64(bound(gasLimit, rbm.request.gasLimit, 2_000_000));
        // receiveBaseMessage(origin, nonce, sender, proofMaturity, message)
        bytes memory expectedCall = abi.encodeWithSelector(
            MessageRecipientMock.receiveBaseMessage.selector, rh.origin, rh.nonce, rbm.sender, timePassed, rbm.content
        );
        // expectCall(address callee, uint256 msgValue, uint64 gas, bytes calldata data)
        vm.expectCall(recipient, 0, gasLimit, expectedCall);
        vm.expectEmit();
        emit Executed(rh.origin, keccak256(msgPayload));
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, gasLimit);
        bytes memory receiptBody = verify_messageStatus(
            keccak256(msgPayload), snapRoot, sm.rsi.stateIndex, MessageStatus.Success, executor, executor
        );
        verify_receipt_valid(receiptBody, rbm.tips);
    }

    function test_execute_base_recipientReverted(Random memory random) public {
        recipient = address(new RevertingApp());
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectEmit();
        emit Executed(rh.origin, keccak256(msgPayload));
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        bytes memory receiptBodyFirst = verify_messageStatus(
            keccak256(msgPayload), snapRoot, sm.rsi.stateIndex, MessageStatus.Failed, executor, address(0)
        );
        verify_receipt_valid(receiptBodyFirst, rbm.tips);
        // Retry the same failed message
        RevertingApp(payable(recipient)).toggleRevert(false);
        vm.prank(executorNew);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        bytes memory receiptBodySecond = verify_messageStatus(
            keccak256(msgPayload), snapRoot, sm.rsi.stateIndex, MessageStatus.Success, executor, executorNew
        );
        // Both receipts (historical and current) should be valid
        verify_receipt_valid(receiptBodyFirst, rbm.tips);
        verify_receipt_valid(receiptBodySecond, rbm.tips);
        cachedStateIndex = uint8(sm.rsi.stateIndex);
    }

    function test_execute_base_revert_alreadyExecuted(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        vm.expectRevert("Already executed");
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
    }

    function test_execute_base_revert_notaryInDispute(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // initiate dispute
        check_submitStateReport(systemContract(), localDomain(), sm.rs, sm.rsi);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert("Notary is in dispute");
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(keccak256(msgPayload));
    }

    function test_execute_base_revert_snapRootUnknown(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = createSnapshotProof(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert("Invalid snapshot root");
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(keccak256(msgPayload));
    }

    function test_execute_base_revert_optimisticPeriodNotOver(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        vm.assume(rh.optimisticPeriod != 0);
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is NOT over
        uint32 timePassed = random.nextUint32() % rh.optimisticPeriod;
        skip(timePassed);
        vm.expectRevert("!optimisticPeriod");
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(keccak256(msgPayload));
    }

    function test_execute_base_revert_gasLimitTooLow(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
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
        vm.expectRevert("Gas limit too low");
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, gasLimit);
        verify_messageStatusNone(keccak256(msgPayload));
    }

    function test_execute_base_revert_gasSuppliedTooLow(Random memory random) public {
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert("Not enough gas supplied");
        vm.prank(executor);
        // Limit amount of gas for the whole call
        testedEH().execute{gas: rbm.request.gasLimit + 20_000}(
            msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit
        );
        verify_messageStatusNone(keccak256(msgPayload));
    }

    function test_execute_base_revert_wrongDestination(Random memory random, uint32 destination_) public {
        vm.assume(destination_ != localDomain());
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, destination_);
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert("!destination");
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        verify_messageStatusNone(keccak256(msgPayload));
    }

    // ══════════════════════════════════════ TESTS: EXECUTE SYSTEM MESSAGES ═══════════════════════════════════════════

    function test_execute_system(
        RawSystemMessage memory rsm,
        RawHeader memory rh,
        SnapshotMock memory sm,
        uint32 timePassed,
        uint64 gasLimit
    ) public {
        // Use System Router Mock for this test
        SystemRouterMock router = (new SystemRouterMock());
        ISystemContract(systemContract()).setSystemRouter(router);
        // Create messages and get origin proof
        bytes memory msgPayload = createSystemMessages(rsm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (bytes32 snapRoot, bytes32[] memory snapProof) = prepareExecution(sm);
        // Make sure that optimistic period is over
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        bytes memory body = rsm.formatSystemMessage();
        // expectCall(address callee, bytes calldata data)
        // receiveSystemMessage(origin, nonce, proofMaturity, body)
        vm.expectCall(
            address(router),
            abi.encodeWithSelector(systemRouter.receiveSystemMessage.selector, rh.origin, rh.nonce, timePassed, body)
        );
        vm.expectEmit();
        emit Executed(rh.origin, keccak256(msgPayload));
        vm.prank(executor);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, gasLimit);
        verify_messageStatus(
            keccak256(msgPayload), snapRoot, sm.rsi.stateIndex, MessageStatus.Success, executor, executor
        );
    }

    // ══════════════════════════════════════════ TESTS: INVALID RECEIPTS ══════════════════════════════════════════════

    function test_verifyReceipt_invalid_msgStatusNone(RawExecReceipt memory re) public {
        vm.assume(testedEH().messageStatus(re.body.messageHash) == MessageStatus.None);
        vm.assume(re.body.origin != localDomain());
        re.body.destination = localDomain();
        verify_receipt_invalid(re);
    }

    function test_verifyReceipt_invalid_msgStatusSuccess(uint256 mask) public {
        test_execute_base_recipientReverted(Random(bytes32(mask)));
        RawReceiptBody memory rrb = RawReceiptBody({
            origin: DOMAIN_REMOTE,
            destination: localDomain(),
            messageHash: getLeaf(0),
            snapshotRoot: getSnapshotRoot(),
            stateIndex: cachedStateIndex,
            attNotary: domains[localDomain()].agent,
            firstExecutor: executor,
            finalExecutor: executorNew
        });
        RawExecReceipt memory re = RawExecReceipt({body: rrb, tips: RawTips(0, 0, 0, 0)});
        // Check that data we start with is valid. Use require() to break the test execution early.
        require(testedEH().isValidReceipt(re.formatReceipt()), "Incorrect initial receipt data");
        RawReceiptBody memory mrb = rrb.modifyReceiptBody(mask);
        RawExecReceipt memory mre = RawExecReceipt({body: mrb, tips: RawTips(0, 0, 0, 0)});
        verify_receipt_invalid(mre);
    }

    // ═════════════════════════════════════════════════ VERIFIERS ═════════════════════════════════════════════════════

    function verify_messageStatusNone(bytes32 messageHash) public {
        verify_messageStatus(messageHash, bytes32(0), 0, MessageStatus.None, address(0), address(0));
    }

    function verify_messageStatus(
        bytes32 messageHash,
        bytes32 snapRoot,
        uint256 stateIndex,
        MessageStatus flag,
        address firstExecutor,
        address finalExecutor
    ) public returns (bytes memory receiptBody) {
        MessageStatus flag_ = testedEH().messageStatus(messageHash);
        assertEq(uint8(flag_), uint8(flag), "!flag");
        receiptBody = testedEH().receiptBody(messageHash);
        if (flag == MessageStatus.None) {
            assertEq(receiptBody.length, 0, "!receiptBody: empty");
        } else {
            address notary = domains[localDomain()].agent;
            assertEq(
                receiptBody,
                ReceiptLib.formatReceiptBody(
                    DOMAIN_REMOTE,
                    localDomain(),
                    messageHash,
                    snapRoot,
                    uint8(stateIndex),
                    notary,
                    firstExecutor,
                    finalExecutor
                )
            );
        }
    }

    function verify_receipt_valid(bytes memory receiptBody, RawTips memory rt) public {
        bytes memory rcptPayload = abi.encodePacked(receiptBody, rt.encodeTips());
        assertTrue(testedEH().isValidReceipt(rcptPayload));
    }

    function verify_receipt_invalid(RawExecReceipt memory re) public {
        bytes memory rcptPayload = re.formatReceipt();
        assertFalse(testedEH().isValidReceipt(rcptPayload));
        address notary = domains[localDomain()].agent;
        bytes memory rcptSignature = signReceipt(notary, rcptPayload);
        // TODO: check that anyone could make the call
        expectAgentSlashed(localDomain(), notary, address(this));
        testedEH().verifyReceipt(rcptPayload, rcptSignature);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Prepares execution of the created messages
    function prepareExecution(SnapshotMock memory sm)
        public
        virtual
        returns (bytes32 snapRoot, bytes32[] memory snapProof);

    function createBaseMessages(RawBaseMessage memory rbm, RawHeader memory rh, uint32 destination_)
        public
        returns (bytes memory msgPayload)
    {
        adjustHeader(rh, destination_);
        rbm.recipient = addressToBytes32(recipient);
        // Set sensible limitations for tips/request
        rbm.tips.boundTips(2 ** 32);
        rbm.request.gasLimit = uint64(bound(rbm.request.gasLimit, 50_000, 200_000));
        msgPayload = RawMessage(uint8(MessageFlag.Base), rh, rbm.formatBaseMessage()).formatMessage();
        createMessages(rh.nonce, msgPayload);
    }

    function createSystemMessages(RawSystemMessage memory rsm, RawHeader memory rh, uint32 destination_)
        public
        returns (bytes memory msgPayload)
    {
        adjustHeader(rh, destination_);
        rsm.boundEntities();
        rsm.callData.selector = SystemContractMock.remoteMockFunc.selector;
        rsm.callData.args = abi.encode(rh.nonce);
        msgPayload = RawMessage(uint8(MessageFlag.System), rh, rsm.formatSystemMessage()).formatMessage();
        createMessages(rh.nonce, msgPayload);
    }

    function createMessages(uint32 msgNonce, bytes memory msgPayload) public {
        for (uint32 nonce = 1; nonce <= MESSAGES; ++nonce) {
            if (nonce == msgNonce) {
                insertMessage(msgPayload);
            } else {
                insertMessage(abi.encode("Mocked payload", nonce));
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
        rh.nonce = 1;
        rh.optimisticPeriod = random.nextUint32();
        sm = SnapshotMock(random.nextState(), RawStateIndex(random.nextUint256(), random.nextUint256()));
        sm.rsi.boundStateIndex();
    }

    function createSnapshotProof(SnapshotMock memory sm)
        public
        returns (RawAttestation memory ra, bytes32[] memory snapProof)
    {
        ra = Random(sm.rs.root).nextAttestation(1);
        ra = createAttestation(sm.rs, ra, sm.rsi);
        snapProof = genSnapshotProof(sm.rsi.stateIndex);
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

    /// @notice Returns tested system contract as IExecutionHub
    function testedEH() public view returns (IExecutionHub) {
        return IExecutionHub(systemContract());
    }
}
