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
        bytes memory receiptData =
            verify_messageStatus(keccak256(msgPayload), snapRoot, MessageStatus.Success, executor, executor);
        verify_receipt_valid(receiptData, rbm.tips);
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
        bytes memory receiptDataFirst =
            verify_messageStatus(keccak256(msgPayload), snapRoot, MessageStatus.Failed, executor, address(0));
        verify_receipt_valid(receiptDataFirst, rbm.tips);
        // Retry the same failed message
        RevertingApp(recipient).toggleRevert(false);
        vm.prank(executorNew);
        testedEH().execute(msgPayload, originProof, snapProof, sm.rsi.stateIndex, rbm.request.gasLimit);
        bytes memory receiptDataSecond =
            verify_messageStatus(keccak256(msgPayload), snapRoot, MessageStatus.Success, executor, executorNew);
        // Both receipts (historical and current) should be valid
        verify_receipt_valid(receiptDataFirst, rbm.tips);
        verify_receipt_valid(receiptDataSecond, rbm.tips);
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
        verify_messageStatus(keccak256(msgPayload), snapRoot, MessageStatus.Success, executor, executor);
    }

    // ══════════════════════════════════════════ TESTS: INVALID RECEIPTS ══════════════════════════════════════════════

    function test_verifyReceipt_invalid_msgStatusNone(RawExecReceipt memory re) public {
        vm.assume(testedEH().messageStatus(re.messageHash) == MessageStatus.None);
        vm.assume(re.origin != localDomain());
        re.destination = localDomain();
        verify_receipt_invalid(re);
    }

    function test_verifyReceipt_invalid_msgStatusSuccess(uint256 mask) public {
        test_execute_base_recipientReverted(Random(bytes32(mask)));
        RawExecReceipt memory re = RawExecReceipt({
            origin: DOMAIN_REMOTE,
            destination: DOMAIN_LOCAL,
            messageHash: getLeaf(0),
            snapshotRoot: getSnapshotRoot(),
            firstExecutor: executor,
            finalExecutor: executorNew,
            tips: RawTips(0, 0, 0, 0)
        });
        // Check that data we start with is valid. Use require() to break the test execution early.
        require(testedEH().isValidReceipt(re.formatReceipt()), "Incorrect initial receipt data");
        RawExecReceipt memory mre = re.modifyReceipt(mask);
        verify_receipt_invalid(mre);
    }

    // ═════════════════════════════════════════════════ VERIFIERS ═════════════════════════════════════════════════════

    function verify_messageStatusNone(bytes32 messageHash) public {
        verify_messageStatus(messageHash, bytes32(0), MessageStatus.None, address(0), address(0));
    }

    function verify_messageStatus(
        bytes32 messageHash,
        bytes32 snapRoot,
        MessageStatus flag,
        address firstExecutor,
        address finalExecutor
    ) public returns (bytes memory receiptData) {
        MessageStatus flag_ = testedEH().messageStatus(messageHash);
        assertEq(uint8(flag_), uint8(flag), "!flag");
        receiptData = testedEH().receiptData(messageHash);
        if (flag == MessageStatus.None) {
            assertEq(receiptData.length, 0, "!receiptData: empty");
        } else {
            assertEq(
                receiptData,
                ReceiptLib.formatReceipt(
                    DOMAIN_REMOTE, localDomain(), messageHash, snapRoot, firstExecutor, finalExecutor, ""
                )
            );
        }
    }

    function verify_receipt_valid(bytes memory receiptData, RawTips memory rt) public {
        bytes memory rcptPayload = abi.encodePacked(receiptData, rt.formatTips());
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
