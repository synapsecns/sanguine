// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1} from "../contracts/libs/Options.sol";

import {
    InterchainClientV1,
    InterchainClientV1BaseTest,
    InterchainTransaction,
    InterchainTxDescriptor
} from "./InterchainClientV1.Base.t.sol";

import {ExecutionFeesMock} from "./mocks/ExecutionFeesMock.sol";
import {ExecutionServiceMock} from "./mocks/ExecutionServiceMock.sol";
import {InterchainDBMock} from "./mocks/InterchainDBMock.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering

/// @notice InterchainClientV1 source chain tests.
/// Happy path for interchainSend and interchainSendEVM:
/// 1. Check that provided msg.value covers the interchain fees for the given modules.
/// 2. Construct the interchain transaction struct, and write its UID to the InterchainDB,
/// requesting verification from the source modules.
/// 3. Allocate the rest of the msg.value to the execution fees,
/// and move it to the ExecutionFees contract (if non-zero).
/// 4. If execution service is provided, request execution from it, passing correct execution fee value.
/// NOTE: it is the execution service's responsibility to check that the provided execution fee is enough.
/// We are using the mocks in this test to verify that the correct values are passed to the contracts,
/// the actual revert "not enough execution fee" test is in the ExecutionService tests, and integration tests.
/// 5. Finally, the event should be emitted.
contract InterchainClientV1SourceTest is InterchainClientV1BaseTest {
    uint256 public constant MOCK_EXECUTION_FEE = 1 ether;
    uint256 public constant MOCK_INTERCHAIN_FEE = 0.5 ether;

    uint256 public constant MOCK_DB_NONCE = 444;
    uint64 public constant MOCK_ENTRY_INDEX = 4;

    OptionsV1 public options = OptionsV1({gasLimit: 100_000, gasAirdrop: 1 ether});
    bytes public encodedOptions = options.encodeOptionsV1();
    bytes public message = "Hello, World!";

    address public srcSender = makeAddr("SrcSender");
    bytes32 public srcSenderBytes32 = bytes32(uint256(uint160(srcSender)));

    bytes32 public dstReceiver = keccak256("DstReceiver");
    address public dstReceiverEVM = makeAddr("DstReceiverEVM");
    bytes32 public dstReceiverEVMBytes32 = bytes32(uint256(uint160(dstReceiverEVM)));

    address[] public twoModules;

    function setUp() public override {
        super.setUp();
        setExecutionFees(execFees);
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        twoModules.push(icModuleA);
        twoModules.push(icModuleB);
        deal(srcSender, MOCK_EXECUTION_FEE + MOCK_INTERCHAIN_FEE);
    }

    /// @dev Override the DB's returned interchain fee for the given destination chain and modules.
    function mockInterchainFee(uint256 dstChainId, address[] memory modules, uint256 interchainFee) internal {
        vm.mockCall(
            icDB, abi.encodeCall(InterchainDBMock.getInterchainFee, (dstChainId, modules)), abi.encode(interchainFee)
        );
    }

    /// @dev Override the DB's returned next entry index (both for reads and writes)
    function mockNextEntryIndex(uint256 dbNonce, uint64 entryIndex) internal {
        bytes memory returnData = abi.encode(dbNonce, entryIndex);
        // Use partial calldata to override return values for calls to these functions with any arguments.
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.getNextEntryIndex.selector), returnData);
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.writeEntry.selector), returnData);
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.writeEntryWithVerification.selector), returnData);
    }

    /// @dev Constructs an interchain transaction and its descriptor for testing.
    function constructInterchainTx(bytes32 receiver)
        internal
        view
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        icTx = InterchainTransaction({
            srcChainId: LOCAL_CHAIN_ID,
            srcSender: srcSenderBytes32,
            dstChainId: REMOTE_CHAIN_ID,
            dstReceiver: receiver,
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            options: encodedOptions,
            message: message
        });
        desc = InterchainTxDescriptor({
            transactionId: icTx.transactionId(),
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX
        });
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function expectWriteEntryWithVerificationCall(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc,
        uint256 interchainFee,
        address[] memory srcModules
    )
        internal
    {
        bytes memory expectedCalldata = abi.encodeCall(
            InterchainDBMock.writeEntryWithVerification, (icTx.dstChainId, desc.transactionId, srcModules)
        );
        vm.expectCall({callee: icDB, msgValue: interchainFee, data: expectedCalldata, count: 1});
    }

    function expectAddExecutionFeeCall(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc,
        uint256 executionFee
    )
        internal
    {
        bytes memory expectedCalldata =
            abi.encodeCall(ExecutionFeesMock.addExecutionFee, (icTx.dstChainId, desc.transactionId));
        vm.expectCall({callee: execFees, msgValue: executionFee, data: expectedCalldata, count: 1});
    }

    function expectNoAddExecutionFeeCall() internal {
        vm.expectCall({
            callee: execFees,
            data: abi.encodeWithSelector(ExecutionFeesMock.addExecutionFee.selector),
            count: 0
        });
    }

    function expectRequestExecutionCall(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc,
        uint256 executionFee
    )
        internal
    {
        uint256 txPayloadSize = abi.encode(icTx).length;
        bytes memory expectedCalldata = abi.encodeCall(
            ExecutionServiceMock.requestExecution,
            (icTx.dstChainId, txPayloadSize, desc.transactionId, executionFee, icTx.options)
        );
        vm.expectCall({callee: execService, data: expectedCalldata, count: 1});
    }

    // ══════════════════════════════════════════ TESTS: INTERCHAIN SEND ═══════════════════════════════════════════════

    function prepareSendTest(
        bytes32 receiver,
        uint256 interchainFee,
        address[] memory srcModules
    )
        internal
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        (icTx, desc) = constructInterchainTx(receiver);
        mockInterchainFee(icTx.dstChainId, srcModules, interchainFee);
        mockNextEntryIndex(MOCK_DB_NONCE, MOCK_ENTRY_INDEX);
    }

    function test_interchainSend_withExecService_icFeeNonZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectRequestExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_EXECUTION_FEE + MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_withExecService_icFeeNonZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectNoAddExecutionFeeCall();
        expectRequestExecutionCall(icTx, desc, 0);
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, 0);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_withExecService_icFeeZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: 0, srcModules: twoModules});
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectRequestExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, 0, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_withExecService_icFeeZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: 0, srcModules: twoModules});
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectNoAddExecutionFeeCall();
        expectRequestExecutionCall(icTx, desc, 0);
        expectEventInterchainTransactionSent(icTx, desc, 0, 0);
        vm.prank(srcSender);
        icClient.interchainSend({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_noExecService_icFeeNonZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_EXECUTION_FEE + MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_noExecService_icFeeNonZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectNoAddExecutionFeeCall();
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, 0);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_noExecService_icFeeZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: 0, srcModules: twoModules});
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, 0, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_noExecService_icFeeZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiver, interchainFee: 0, srcModules: twoModules});
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectNoAddExecutionFeeCall();
        expectEventInterchainTransactionSent(icTx, desc, 0, 0);
        vm.prank(srcSender);
        icClient.interchainSend({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    // ═══════════════════════════════════════ TESTS: INTERCHAIN SEND TO EVM ═══════════════════════════════════════════

    function test_interchainSendEVM_withExecService_icFeeNonZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareSendTest({
            receiver: dstReceiverEVMBytes32,
            interchainFee: MOCK_INTERCHAIN_FEE,
            srcModules: twoModules
        });
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectRequestExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_EXECUTION_FEE + MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_withExecService_icFeeNonZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareSendTest({
            receiver: dstReceiverEVMBytes32,
            interchainFee: MOCK_INTERCHAIN_FEE,
            srcModules: twoModules
        });
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectNoAddExecutionFeeCall();
        expectRequestExecutionCall(icTx, desc, 0);
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, 0);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_withExecService_icFeeZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: 0, srcModules: twoModules});
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectRequestExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, 0, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_withExecService_icFeeZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: 0, srcModules: twoModules});
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectNoAddExecutionFeeCall();
        expectRequestExecutionCall(icTx, desc, 0);
        expectEventInterchainTransactionSent(icTx, desc, 0, 0);
        vm.prank(srcSender);
        icClient.interchainSendEVM({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_noExecService_icFeeNonZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareSendTest({
            receiver: dstReceiverEVMBytes32,
            interchainFee: MOCK_INTERCHAIN_FEE,
            srcModules: twoModules
        });
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_EXECUTION_FEE + MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_noExecService_icFeeNonZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareSendTest({
            receiver: dstReceiverEVMBytes32,
            interchainFee: MOCK_INTERCHAIN_FEE,
            srcModules: twoModules
        });
        expectWriteEntryWithVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectNoAddExecutionFeeCall();
        expectEventInterchainTransactionSent(icTx, desc, MOCK_INTERCHAIN_FEE, 0);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_noExecService_icFeeZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: 0, srcModules: twoModules});
        // Anything paid on top of the interchain fee is considered an execution fee.
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectAddExecutionFeeCall(icTx, desc, MOCK_EXECUTION_FEE);
        expectEventInterchainTransactionSent(icTx, desc, 0, MOCK_EXECUTION_FEE);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_noExecService_icFeeZero_execFeeZero() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: 0, srcModules: twoModules});
        expectWriteEntryWithVerificationCall(icTx, desc, 0, twoModules);
        expectNoAddExecutionFeeCall();
        expectEventInterchainTransactionSent(icTx, desc, 0, 0);
        vm.prank(srcSender);
        icClient.interchainSendEVM({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }
}
