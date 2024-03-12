// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1, InterchainClientV1Events, IInterchainClientV1} from "../contracts/InterchainClientV1.sol";
import {InterchainTxDescriptor, InterchainTransaction} from "../contracts/libs/InterchainTransaction.sol";

import {ExecutionFeesMock} from "./mocks/ExecutionFeesMock.sol";
import {InterchainDBMock} from "./mocks/InterchainDBMock.sol";

import {Test} from "forge-std/Test.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract InterchainClientV1BaseTest is Test, InterchainClientV1Events {
    uint256 public constant LOCAL_CHAIN_ID = 1337;
    uint256 public constant REMOTE_CHAIN_ID = 7331;
    uint256 public constant UNKNOWN_CHAIN_ID = 42;
    bytes32 public constant MOCK_REMOTE_CLIENT = keccak256("RemoteClient");

    address public mockRemoteClientEVM = makeAddr("RemoteClientEVM");
    bytes32 public mockRemoteClientEVMBytes32 = bytes32(uint256(uint160(mockRemoteClientEVM)));

    InterchainClientV1 public icClient;
    address public icDB;
    address public execFees;

    address public owner = makeAddr("Owner");

    function setUp() public virtual {
        vm.chainId(LOCAL_CHAIN_ID);
        icDB = address(new InterchainDBMock());
        icClient = new InterchainClientV1(icDB, owner);
        execFees = address(new ExecutionFeesMock());
    }

    function setExecutionFees(address executionFees) public {
        vm.prank(owner);
        icClient.setExecutionFees(executionFees);
    }

    function setLinkedClient(uint256 chainId, bytes32 client) public {
        vm.prank(owner);
        icClient.setLinkedClient(chainId, client);
    }

    // ═════════════════════════════════════════════ EXPECT (REVERTS) ══════════════════════════════════════════════════

    function expectRevertFeeAmountTooLow(uint256 actual, uint256 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__FeeAmountTooLow.selector, actual, required)
        );
    }

    function expectRevertIncorrectDstChainId(uint256 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__IncorrectDstChainId.selector, chainId)
        );
    }

    function expectRevertIncorrectMsgValue(uint256 actual, uint256 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__IncorrectMsgValue.selector, actual, required)
        );
    }

    function expectRevertNoLinkedClient(uint256 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__NoLinkedClient.selector, chainId)
        );
    }

    function expectRevertNotEnoughResponses(uint256 actual, uint256 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__NotEnoughResponses.selector, actual, required
            )
        );
    }

    function expectRevertNotEVMClient(bytes32 client) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__NotEVMClient.selector, client));
    }

    function expectRevertNotRemoteChainId(uint256 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__NotRemoteChainId.selector, chainId)
        );
    }

    function expectRevertTxAlreadyExecuted(bytes32 transactionId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__TxAlreadyExecuted.selector, transactionId)
        );
    }

    function expectRevertTxNotExecuted(bytes32 transactionId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__TxNotExecuted.selector, transactionId)
        );
    }

    function expectRevertOwnableUnauthorizedAccount(address account) internal {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, account));
    }

    // ══════════════════════════════════════════════ EXPECT (EVENTS) ══════════════════════════════════════════════════

    function expectEventExecutionFeesSet(address executionFees) internal {
        vm.expectEmit(address(icClient));
        emit ExecutionFeesSet(executionFees);
    }

    function expectEventLinkedClientSet(uint256 chainId, bytes32 client) internal {
        vm.expectEmit(address(icClient));
        emit LinkedClientSet(chainId, client);
    }

    function expectEventInterchainTransactionSent(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc,
        uint256 verificationFee,
        uint256 executionFee
    )
        internal
    {
        // Sanity check
        assertCorrectDescriptor(icTx, desc);
        vm.expectEmit(address(icClient));
        emit InterchainTransactionSent({
            transactionId: desc.transactionId,
            dbNonce: desc.dbNonce,
            entryIndex: desc.entryIndex,
            dstChainId: icTx.dstChainId,
            srcSender: icTx.srcSender,
            dstReceiver: icTx.dstReceiver,
            verificationFee: verificationFee,
            executionFee: executionFee,
            options: icTx.options,
            message: icTx.message
        });
    }

    function expectEventInterchainTransactionReceived(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc
    )
        internal
    {
        // Sanity check
        assertCorrectDescriptor(icTx, desc);
        vm.expectEmit(address(icClient));
        emit InterchainTransactionReceived({
            transactionId: desc.transactionId,
            dbNonce: desc.dbNonce,
            entryIndex: desc.entryIndex,
            srcChainId: icTx.srcChainId,
            srcSender: icTx.srcSender,
            dstReceiver: icTx.dstReceiver
        });
    }

    function expectEventExecutionProofWritten(InterchainTxDescriptor memory desc, address executor) internal {
        vm.expectEmit(address(icClient));
        emit ExecutionProofWritten({
            transactionId: desc.transactionId,
            dbNonce: desc.dbNonce,
            entryIndex: desc.entryIndex,
            executor: executor
        });
    }

    // ════════════════════════════════════════════════ ASSERTIONS ═════════════════════════════════════════════════════

    function assertCorrectDescriptor(InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) internal {
        assertEq(desc.dbNonce, icTx.dbNonce, "!desc.dbNonce");
        assertEq(desc.entryIndex, icTx.entryIndex, "!desc.entryIndex");
        assertEq(desc.transactionId, icTx.transactionId(), "!desc.transactionId");
    }

    function assertEq(InterchainTransaction memory icTx, InterchainTransaction memory expected) internal {
        assertEq(icTx.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(icTx.srcSender, expected.srcSender, "!srcSender");
        assertEq(icTx.dstChainId, expected.dstChainId, "!dstChainId");
        assertEq(icTx.dstReceiver, expected.dstReceiver, "!dstReceiver");
        assertEq(icTx.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(icTx.entryIndex, expected.entryIndex, "!entryIndex");
        assertEq(icTx.options, expected.options, "!options");
        assertEq(icTx.message, expected.message, "!message");
    }
}
