// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IExecutionHub} from "../../../contracts/interfaces/IExecutionHub.sol";
import {SNAPSHOT_MAX_STATES} from "../../../contracts/libs/Snapshot.sol";

import {MessageRecipientMock} from "../../mocks/client/MessageRecipientMock.t.sol";
import {SystemContractMock} from "../../mocks/system/SystemContractMock.t.sol";

import {
    MessageFlag,
    RawBaseMessage,
    RawHeader,
    RawMessage,
    RawState,
    RawSystemMessage
} from "../../utils/libs/SynapseStructs.t.sol";
import {DisputeHubTest} from "./DisputeHub.t.sol";

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
        bytes memory msgPayload = createBaseMessages(rbm, rh);
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        bytes32[] memory snapProof = prepareExecution(sm);
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
        bytes memory msgPayload = createSystemMessages(rsm, rh);
        bytes32[] memory originProof = getLatestProof(rh.nonce - 1);
        // Create snapshot proof
        adjustSnapshot(sm);
        bytes32[] memory snapProof = prepareExecution(sm);
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
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Prepares execution of the created messages
    function prepareExecution(SnapshotMock memory sm) public virtual returns (bytes32[] memory snapProof);

    /// @notice Local domain for ExecutionHub tests
    function localDomain() public view virtual returns (uint32);

    function createBaseMessages(RawBaseMessage memory rbm, RawHeader memory rh)
        public
        returns (bytes memory msgPayload)
    {
        adjustHeader(rh);
        rbm.recipient = addressToBytes32(recipient);
        // Set sensible limitations for tips/request
        rbm.tips.boundTips(1e20);
        rbm.request.gasLimit = uint64(bound(rbm.request.gasLimit, 10_000, 1_000_000));
        msgPayload = RawMessage(uint8(MessageFlag.Base), rh, rbm.formatBaseMessage()).formatMessage();
        createMessages(rh.nonce, msgPayload);
    }

    function createSystemMessages(RawSystemMessage memory rsm, RawHeader memory rh)
        public
        returns (bytes memory msgPayload)
    {
        adjustHeader(rh);
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

    /// @notice Sets realistic values for the message header
    function adjustHeader(RawHeader memory rh) public view {
        rh.origin = DOMAIN_REMOTE;
        rh.nonce = uint32(bound(rh.nonce, 1, MESSAGES));
        rh.destination = localDomain();
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
