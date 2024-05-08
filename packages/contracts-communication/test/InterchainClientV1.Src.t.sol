// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1} from "../contracts/libs/Options.sol";
import {VersionedPayloadLib} from "../contracts/libs/VersionedPayload.sol";

import {
    InterchainClientV1BaseTest, InterchainTransaction, InterchainTxDescriptor
} from "./InterchainClientV1.Base.t.sol";

import {ExecutionServiceMock} from "./mocks/ExecutionServiceMock.sol";
import {InterchainDBMock} from "./mocks/InterchainDBMock.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering

/// @notice InterchainClientV1 source chain tests.
/// # Happy path for interchainSend and interchainSendEVM.
/// 1. Check that provided msg.value covers the interchain fees for the given modules.
/// 2. Construct the interchain transaction struct, and write its UID to the InterchainDB,
/// requesting verification from the source modules.
/// 3. Allocate the rest of the msg.value to the execution fees,
/// and move it to the ExecutionService contract.
/// 4. If execution service is provided, request execution from it, passing correct execution fee value.
/// NOTE: it is the execution service's responsibility to check that the provided execution fee is enough.
/// We are using the mocks in this test to verify that the correct values are passed to the contracts,
/// the actual revert "not enough execution fee" test is in the ExecutionService tests, and integration tests.
/// 5. Finally, the event should be emitted.
/// # Revert cases for interchainSend and interchainSendEVM.
/// - Provided msg.value is less than the required interchain fee.
/// - Destination chain ID is not a remote chain ID.
/// - Destination chain ID does not have a linked client.
/// - Receiver is a zero address.
/// - Options are not encoded correctly.
contract InterchainClientV1SourceTest is InterchainClientV1BaseTest {
    uint256 public constant MOCK_EXECUTION_FEE = 1 ether;
    uint256 public constant MOCK_INTERCHAIN_FEE = 0.5 ether;

    uint64 public constant MOCK_DB_NONCE = 444;
    uint64 public constant MOCK_ENTRY_INDEX = 4;

    OptionsV1 public options = OptionsV1({gasLimit: 100_000, gasAirdrop: 1 ether});
    bytes public encodedOptions = options.encodeOptionsV1();

    bytes public invalidOptionsV0 = VersionedPayloadLib.encodeVersionedPayload(0, abi.encode(options));
    bytes public invalidOptionsV1 = VersionedPayloadLib.encodeVersionedPayload(1, abi.encode(options.gasLimit));

    bytes public message = "Hello, World!";

    address public srcSender = makeAddr("SrcSender");
    bytes32 public srcSenderBytes32 = bytes32(uint256(uint160(srcSender)));

    bytes32 public dstReceiver = keccak256("DstReceiver");
    address public dstReceiverEVM = makeAddr("DstReceiverEVM");
    bytes32 public dstReceiverEVMBytes32 = bytes32(uint256(uint160(dstReceiverEVM)));

    address[] public twoModules;

    function setUp() public override {
        super.setUp();
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        twoModules.push(icModuleA);
        twoModules.push(icModuleB);
        deal(srcSender, MOCK_EXECUTION_FEE + MOCK_INTERCHAIN_FEE);
    }

    /// @dev Override the DB's returned interchain fee for the given destination chain and modules.
    function mockInterchainFee(uint64 dstChainId, address[] memory modules, uint256 interchainFee) internal {
        vm.mockCall(
            icDB, abi.encodeCall(InterchainDBMock.getInterchainFee, (dstChainId, modules)), abi.encode(interchainFee)
        );
    }

    /// @dev Override the ExecutionService's returned execution fee for the given destination chain and transaction.
    function mockExecutionFee(uint64 dstChainId, InterchainTransaction memory icTx, uint256 executionFee) internal {
        uint256 txPayloadSize = getEncodedTx(icTx).length;
        vm.mockCall(
            execService,
            abi.encodeCall(ExecutionServiceMock.getExecutionFee, (dstChainId, txPayloadSize, icTx.options)),
            abi.encode(executionFee)
        );
    }

    /// @dev Override the DB's returned next entry index (both for reads and writes)
    function mockNextEntryIndex(uint64 dbNonce, uint64 entryIndex) internal {
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
            transactionId: keccak256(getEncodedTx(icTx)),
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX
        });
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function expectWriteEntryRequestVerificationCall(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc,
        uint256 interchainFee,
        address[] memory srcModules
    )
        internal
    {
        bytes memory expectedCalldata = abi.encodeCall(
            InterchainDBMock.writeEntryRequestVerification, (icTx.dstChainId, desc.transactionId, srcModules)
        );
        vm.expectCall({callee: icDB, msgValue: interchainFee, data: expectedCalldata, count: 1});
    }

    function expectRequestTxExecutionCall(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc,
        uint256 executionFee
    )
        internal
    {
        uint256 txPayloadSize = getEncodedTx(icTx).length;
        bytes memory expectedCalldata = abi.encodeCall(
            ExecutionServiceMock.requestTxExecution, (icTx.dstChainId, txPayloadSize, desc.transactionId, icTx.options)
        );
        vm.expectCall({callee: execService, msgValue: executionFee, data: expectedCalldata, count: 1});
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
        expectWriteEntryRequestVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectRequestTxExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
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
        expectWriteEntryRequestVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectRequestTxExecutionCall(icTx, desc, 0);
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
        expectWriteEntryRequestVerificationCall(icTx, desc, 0, twoModules);
        expectRequestTxExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
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
        expectWriteEntryRequestVerificationCall(icTx, desc, 0, twoModules);
        expectRequestTxExecutionCall(icTx, desc, 0);
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

    // ═════════════════════════════════════════ INTERCHAIN SEND: REVERTS ══════════════════════════════════════════════

    // Pay 1 wei less than required and expect a revert with the correct fee amount.
    function test_interchainSend_revert_feeTooLow() public {
        prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        deal(srcSender, MOCK_INTERCHAIN_FEE - 1);
        expectRevertFeeAmountBelowMin(MOCK_INTERCHAIN_FEE - 1, MOCK_INTERCHAIN_FEE);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE - 1}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_revert_ChainIdNotRemote() public {
        prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: LOCAL_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_revert_ChainIdNotLinked() public {
        prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertChainIdNotLinked(UNKNOWN_CHAIN_ID);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: UNKNOWN_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_revert_ReceiverZeroAddress() public {
        prepareSendTest({receiver: 0, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertReceiverZeroAddress();
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: 0,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_revert_ExecutionServiceZeroAddress() public {
        prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertExecutionServiceZeroAddress();
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSend_revert_emptyOptions() public {
        prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: "",
            message: message
        });
    }

    function test_interchainSend_revert_invalidOptionsV0() public {
        prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertVersionInvalid(0);
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: invalidOptionsV0,
            message: message
        });
    }

    function test_interchainSend_revert_invalidOptionsV1() public {
        prepareSendTest({receiver: dstReceiver, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        vm.prank(srcSender);
        icClient.interchainSend{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiver,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: invalidOptionsV1,
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
        expectWriteEntryRequestVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectRequestTxExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
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
        expectWriteEntryRequestVerificationCall(icTx, desc, MOCK_INTERCHAIN_FEE, twoModules);
        expectRequestTxExecutionCall(icTx, desc, 0);
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
        expectWriteEntryRequestVerificationCall(icTx, desc, 0, twoModules);
        expectRequestTxExecutionCall(icTx, desc, MOCK_EXECUTION_FEE);
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
        expectWriteEntryRequestVerificationCall(icTx, desc, 0, twoModules);
        expectRequestTxExecutionCall(icTx, desc, 0);
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

    // ═══════════════════════════════════════ INTERCHAIN SEND EVM: REVERTS ════════════════════════════════════════════

    // Pay 1 wei less than required and expect a revert with the correct fee amount.
    function test_interchainSendEVM_revert_feeTooLow() public {
        prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        deal(srcSender, MOCK_INTERCHAIN_FEE - 1);
        expectRevertFeeAmountBelowMin(MOCK_INTERCHAIN_FEE - 1, MOCK_INTERCHAIN_FEE);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE - 1}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_revert_ChainIdNotRemote() public {
        prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: LOCAL_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_revert_ChainIdNotLinked() public {
        prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertChainIdNotLinked(UNKNOWN_CHAIN_ID);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: UNKNOWN_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_revert_ReceiverZeroAddress() public {
        prepareSendTest({receiver: 0, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertReceiverZeroAddress();
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: address(0),
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_revert_ExecutionServiceZeroAddress() public {
        prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertExecutionServiceZeroAddress();
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            message: message
        });
    }

    function test_interchainSendEVM_revert_emptyOptions() public {
        prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: "",
            message: message
        });
    }

    function test_interchainSendEVM_revert_invalidOptionsV0() public {
        prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        expectRevertVersionInvalid(0);
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: invalidOptionsV0,
            message: message
        });
    }

    function test_interchainSendEVM_revert_invalidOptionsV1() public {
        prepareSendTest({receiver: dstReceiverEVMBytes32, interchainFee: MOCK_INTERCHAIN_FEE, srcModules: twoModules});
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        vm.prank(srcSender);
        icClient.interchainSendEVM{value: MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE}({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: dstReceiverEVM,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: invalidOptionsV1,
            message: message
        });
    }

    // ═════════════════════════════════════════ TESTS: GET INTERCHAIN FEE ═════════════════════════════════════════════

    function test_getInterchainFee_withExecService_icFeeNonZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(dstReceiver);
        mockInterchainFee(REMOTE_CHAIN_ID, twoModules, MOCK_INTERCHAIN_FEE);
        mockExecutionFee(REMOTE_CHAIN_ID, icTx, MOCK_EXECUTION_FEE);
        assertEq(
            icClient.getInterchainFee({
                dstChainId: REMOTE_CHAIN_ID,
                srcExecutionService: execService,
                srcModules: twoModules,
                options: encodedOptions,
                messageLen: message.length
            }),
            MOCK_INTERCHAIN_FEE + MOCK_EXECUTION_FEE
        );
    }

    function test_getInterchainFee_withExecService_icFeeNonZero_execFeeZero() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(dstReceiver);
        mockInterchainFee(REMOTE_CHAIN_ID, twoModules, MOCK_INTERCHAIN_FEE);
        mockExecutionFee(REMOTE_CHAIN_ID, icTx, 0);
        assertEq(
            icClient.getInterchainFee({
                dstChainId: REMOTE_CHAIN_ID,
                srcExecutionService: execService,
                srcModules: twoModules,
                options: encodedOptions,
                messageLen: message.length
            }),
            MOCK_INTERCHAIN_FEE
        );
    }

    function test_getInterchainFee_withExecService_icFeeZero_execFeeNonZero() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(dstReceiver);
        mockInterchainFee(REMOTE_CHAIN_ID, twoModules, 0);
        mockExecutionFee(REMOTE_CHAIN_ID, icTx, MOCK_EXECUTION_FEE);
        assertEq(
            icClient.getInterchainFee({
                dstChainId: REMOTE_CHAIN_ID,
                srcExecutionService: execService,
                srcModules: twoModules,
                options: encodedOptions,
                messageLen: message.length
            }),
            MOCK_EXECUTION_FEE
        );
    }

    function test_getInterchainFee_withExecService_icFeeZero_execFeeZero() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(dstReceiver);
        mockInterchainFee(REMOTE_CHAIN_ID, twoModules, 0);
        mockExecutionFee(REMOTE_CHAIN_ID, icTx, 0);
        assertEq(
            icClient.getInterchainFee({
                dstChainId: REMOTE_CHAIN_ID,
                srcExecutionService: execService,
                srcModules: twoModules,
                options: encodedOptions,
                messageLen: message.length
            }),
            0
        );
    }

    function test_getInterchainFee_revert_ChainIdNotRemote() public {
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        icClient.getInterchainFee({
            dstChainId: LOCAL_CHAIN_ID,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            messageLen: message.length
        });
    }

    function test_getInterchainFee_revert_ChainIdNotLinked() public {
        expectRevertChainIdNotLinked(UNKNOWN_CHAIN_ID);
        icClient.getInterchainFee({
            dstChainId: UNKNOWN_CHAIN_ID,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: encodedOptions,
            messageLen: message.length
        });
    }

    function test_getInterchainFee_revert_ExecutionServiceZeroAddress() public {
        expectRevertExecutionServiceZeroAddress();
        icClient.getInterchainFee({
            dstChainId: REMOTE_CHAIN_ID,
            srcExecutionService: address(0),
            srcModules: twoModules,
            options: encodedOptions,
            messageLen: message.length
        });
    }

    function test_getInterchainFee_revert_emptyOptions() public {
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        icClient.getInterchainFee({
            dstChainId: REMOTE_CHAIN_ID,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: "",
            messageLen: message.length
        });
    }

    function test_getInterchainFee_revert_invalidOptionsV0() public {
        expectRevertVersionInvalid(0);
        icClient.getInterchainFee({
            dstChainId: REMOTE_CHAIN_ID,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: invalidOptionsV0,
            messageLen: message.length
        });
    }

    function test_getInterchainFee_revert_invalidOptionsV1() public {
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        icClient.getInterchainFee({
            dstChainId: REMOTE_CHAIN_ID,
            srcExecutionService: execService,
            srcModules: twoModules,
            options: invalidOptionsV1,
            messageLen: message.length
        });
    }
}
