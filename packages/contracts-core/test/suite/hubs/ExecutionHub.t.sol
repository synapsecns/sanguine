// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IExecutionHub} from "../../../contracts/interfaces/IExecutionHub.sol";
import {SNAPSHOT_MAX_STATES} from "../../../contracts/libs/Snapshot.sol";
import {MessageStatus} from "../../../contracts/libs/Structures.sol";

import {RevertingApp} from "../../harnesses/client/RevertingApp.t.sol";
import {MessageRecipientMock} from "../../mocks/client/MessageRecipientMock.t.sol";
import {SystemContractMock} from "../../mocks/system/SystemContractMock.t.sol";

import {Random} from "../../utils/libs/Random.t.sol";
import {
    ExecutionLib,
    MessageFlag,
    RawAttestation,
    RawBaseMessage,
    RawHeader,
    RawMessage,
    RawState,
    RawSystemMessage
} from "../../utils/libs/SynapseStructs.t.sol";
import {DisputeHubTest, IDisputeHub} from "./DisputeHub.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract ExecutionHubTest is DisputeHubTest {
    struct SnapshotMock {
        RawState rs;
        uint256 stateIndex;
        uint256 statesAmount;
    }

    address internal recipient;

    function setUp() public virtual override {
        super.setUp();
        recipient = address(new MessageRecipientMock());
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testExecutionHub() external {}

    // ═══════════════════════════════════════ TESTS: EXECUTE BASE MESSAGES ════════════════════════════════════════════

    function check_execute_base(
        address hub,
        RawBaseMessage memory rbm,
        RawHeader memory rh,
        SnapshotMock memory sm,
        uint32 timePassed,
        uint64 gasLimit
    ) public {
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, gasLimit);
        verify_executionStatus(hub, keccak256(msgPayload), snapRoot, MessageStatus.Success, executor, executor);
    }

    function check_execute_base_recipientReverted(address hub, Random memory random) public {
        recipient = address(new RevertingApp());
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
        verify_executionStatus(hub, keccak256(msgPayload), snapRoot, MessageStatus.Failed, executor, address(0));
        // Retry the same failed message
        RevertingApp(recipient).toggleRevert(false);
        address executorNew = makeAddr("Executor New");
        vm.prank(executorNew);
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
        verify_executionStatus(hub, keccak256(msgPayload), snapRoot, MessageStatus.Success, executor, executorNew);
    }

    function check_execute_base_revert_alreadyExecuted(address hub, Random memory random) public {
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
        vm.expectRevert("Already executed");
        vm.prank(executor);
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
    }

    function check_execute_base_revert_notaryInDispute(address hub, Random memory random) public {
        address executor = makeAddr("Executor");
        // Create some simple data
        (RawBaseMessage memory rbm, RawHeader memory rh, SnapshotMock memory sm) = createDataRevertTest(random);
        // Create messages and get origin proof
        bytes memory msgPayload = createBaseMessages(rbm, rh, localDomain());
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        (, bytes32[] memory snapProof) = prepareExecution(sm);
        // initiate dispute
        check_submitStateReport(hub, localDomain(), sm.rs, sm.statesAmount, sm.stateIndex);
        // Make sure that optimistic period is over
        uint32 timePassed = random.nextUint32();
        timePassed = uint32(bound(timePassed, rh.optimisticPeriod, rh.optimisticPeriod + 1 days));
        skip(timePassed);
        vm.expectRevert("Notary is in dispute");
        vm.prank(executor);
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
        verify_executionStatusNone(hub, keccak256(msgPayload));
    }

    function check_execute_base_revert_snapRootUnknown(address hub, Random memory random) public {
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
        verify_executionStatusNone(hub, keccak256(msgPayload));
    }

    function check_execute_base_revert_optimisticPeriodNotOver(address hub, Random memory random) public {
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
        verify_executionStatusNone(hub, keccak256(msgPayload));
    }

    function check_execute_base_revert_gasLimitTooLow(address hub, Random memory random) public {
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, gasLimit);
        verify_executionStatusNone(hub, keccak256(msgPayload));
    }

    function check_execute_base_revert_gasSuppliedTooLow(address hub, Random memory random) public {
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute{gas: rbm.request.gasLimit + 20_000}(
            msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit
        );
        verify_executionStatusNone(hub, keccak256(msgPayload));
    }

    function check_execute_base_revert_wrongDestination(address hub, Random memory random, uint32 destination_)
        public
    {
        vm.assume(destination_ != localDomain());
        address executor = makeAddr("Executor");
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
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, rbm.request.gasLimit);
        verify_executionStatusNone(hub, keccak256(msgPayload));
    }

    // ══════════════════════════════════════ TESTS: EXECUTE SYSTEM MESSAGES ═══════════════════════════════════════════

    function check_execute_system(
        address hub,
        address router,
        RawSystemMessage memory rsm,
        RawHeader memory rh,
        SnapshotMock memory sm,
        uint32 timePassed,
        uint64 gasLimit
    ) public {
        address executor = makeAddr("Executor");
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
            router,
            abi.encodeWithSelector(systemRouter.receiveSystemMessage.selector, rh.origin, rh.nonce, timePassed, body)
        );
        vm.expectEmit();
        emit Executed(rh.origin, keccak256(msgPayload));
        vm.prank(executor);
        IExecutionHub(hub).execute(msgPayload, originProof, snapProof, sm.stateIndex, gasLimit);
        verify_executionStatus(hub, keccak256(msgPayload), snapRoot, MessageStatus.Success, executor, executor);
    }

    // ═════════════════════════════════════════════════ VERIFIERS ═════════════════════════════════════════════════════

    function verify_executionStatusNone(address hub, bytes32 messageHash) public {
        verify_executionStatus(hub, messageHash, bytes32(0), MessageStatus.None, address(0), address(0));
    }

    function verify_executionStatus(
        address hub,
        bytes32 messageHash,
        bytes32 snapRoot,
        MessageStatus flag,
        address firstExecutor,
        address finalExecutor
    ) public {
        MessageStatus flag_ = IExecutionHub(hub).messageStatus(messageHash);
        assertEq(uint8(flag_), uint8(flag), "!flag");
        bytes memory data = IExecutionHub(hub).executionData(messageHash);
        if (flag == MessageStatus.None) {
            assertEq(data.length, 0, "!executionData: empty");
        } else {
            assertEq(
                data,
                ExecutionLib.formatExecution(
                    flag, DOMAIN_REMOTE, localDomain(), messageHash, snapRoot, firstExecutor, finalExecutor, ""
                )
            );
        }
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Prepares execution of the created messages
    function prepareExecution(SnapshotMock memory sm)
        public
        virtual
        returns (bytes32 snapRoot, bytes32[] memory snapProof);

    /// @notice Local domain for ExecutionHub tests
    function localDomain() public view virtual returns (uint32);

    function createBaseMessages(RawBaseMessage memory rbm, RawHeader memory rh, uint32 destination_)
        public
        returns (bytes memory msgPayload)
    {
        adjustHeader(rh, destination_);
        rbm.recipient = addressToBytes32(recipient);
        // Set sensible limitations for tips/request
        rbm.tips.boundTips(1e20);
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
        rh.nonce = random.nextUint32();
        rh.optimisticPeriod = random.nextUint32();
        sm = SnapshotMock(random.nextState(), random.nextUint256(), random.nextUint256());
    }

    function createSnapshotProof(SnapshotMock memory sm)
        public
        returns (RawAttestation memory ra, bytes32[] memory snapProof)
    {
        ra = Random(sm.rs.root).nextAttestation(1);
        ra = createAttestation(sm.rs, ra, sm.statesAmount, sm.stateIndex);
        snapProof = genSnapshotProof(sm.stateIndex);
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
        sm.statesAmount = bound(sm.statesAmount, 1, SNAPSHOT_MAX_STATES);
        sm.stateIndex = bound(sm.stateIndex, 0, sm.statesAmount - 1);
        adjustState(sm.rs);
    }
}
