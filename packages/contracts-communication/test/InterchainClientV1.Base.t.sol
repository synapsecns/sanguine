// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1, InterchainClientV1Events, IInterchainClientV1} from "../contracts/InterchainClientV1.sol";
import {IInterchainApp} from "../contracts/interfaces/IInterchainApp.sol";
import {AppConfigV1} from "../contracts/libs/AppConfig.sol";
import {InterchainTxDescriptor, InterchainTransaction} from "../contracts/libs/InterchainTransaction.sol";
import {OptionsLib} from "../contracts/libs/Options.sol";

import {InterchainTransactionLibHarness} from "./harnesses/InterchainTransactionLibHarness.sol";
import {VersionedPayloadLibHarness} from "./harnesses/VersionedPayloadLibHarness.sol";
import {ExecutionServiceMock} from "./mocks/ExecutionServiceMock.sol";
import {InterchainDBMock} from "./mocks/InterchainDBMock.sol";
import {InterchainModuleMock} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract InterchainClientV1BaseTest is Test, InterchainClientV1Events {
    uint64 public constant LOCAL_CHAIN_ID = 1337;
    uint64 public constant REMOTE_CHAIN_ID = 7331;
    uint64 public constant UNKNOWN_CHAIN_ID = 42;
    bytes32 public constant MOCK_REMOTE_CLIENT = keccak256("RemoteClient");
    uint16 public constant CLIENT_VERSION = 1;

    InterchainTransactionLibHarness public txLibHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    address public mockRemoteClientEVM = makeAddr("RemoteClientEVM");
    bytes32 public mockRemoteClientEVMBytes32 = bytes32(uint256(uint160(mockRemoteClientEVM)));

    InterchainClientV1 public icClient;
    address public icDB;
    address public icModuleA;
    address public icModuleB;

    address public execService;

    address public owner = makeAddr("Owner");
    address public defaultGuard = makeAddr("Default Guard");

    function setUp() public virtual {
        vm.chainId(LOCAL_CHAIN_ID);
        icDB = address(new InterchainDBMock());
        icClient = new InterchainClientV1(icDB, owner);
        execService = address(new ExecutionServiceMock());
        icModuleA = address(new InterchainModuleMock());
        icModuleB = address(new InterchainModuleMock());
        txLibHarness = new InterchainTransactionLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
    }

    function setDefaultGuard(address guard) public {
        vm.prank(owner);
        icClient.setDefaultGuard(guard);
    }

    function setLinkedClient(uint64 chainId, bytes32 client) public {
        vm.prank(owner);
        icClient.setLinkedClient(chainId, client);
    }

    // ══════════════════════════════════════════════════ MOCKING ══════════════════════════════════════════════════════

    function mockReceivingConfig(address receiver, AppConfigV1 memory appConfig, address[] memory modules) internal {
        bytes memory encodedConfig = appConfig.encodeAppConfigV1();
        vm.mockCall({
            callee: receiver,
            data: abi.encodeCall(IInterchainApp.getReceivingConfig, ()),
            returnData: abi.encode(encodedConfig, modules)
        });
    }

    // ═════════════════════════════════════════════ EXPECT (REVERTS) ══════════════════════════════════════════════════

    function expectRevertEntryConflict(address module) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__EntryConflict.selector, module));
    }

    function expectRevertFeeAmountBelowMin(uint256 feeAmount, uint256 minRequired) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__FeeAmountBelowMin.selector, feeAmount, minRequired
            )
        );
    }

    function expectRevertDstChainIdNotLocal(uint64 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__DstChainIdNotLocal.selector, chainId)
        );
    }

    function expectRevertMsgValueMismatch(uint256 actual, uint256 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__MsgValueMismatch.selector, actual, required)
        );
    }

    function expectRevertTxVersionMismatch(uint16 version, uint16 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__TxVersionMismatch.selector, version, required
            )
        );
    }

    function expectRevertChainIdNotLinked(uint64 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__ChainIdNotLinked.selector, chainId)
        );
    }

    function expectRevertGasLeftBelowMin() internal {
        vm.expectRevert(IInterchainClientV1.InterchainClientV1__GasLeftBelowMin.selector);
    }

    function expectRevertResponsesAmountBelowMin(uint256 actual, uint256 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__ResponsesAmountBelowMin.selector, actual, required
            )
        );
    }

    function expectRevertLinkedClientNotEVM(bytes32 client) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__LinkedClientNotEVM.selector, client)
        );
    }

    function expectRevertChainIdNotRemote(uint64 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__ChainIdNotRemote.selector, chainId)
        );
    }

    function expectRevertReceiverNotICApp(address receiver) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__ReceiverNotICApp.selector, receiver)
        );
    }

    function expectRevertReceiverZeroRequiredResponses(address receiver) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__ReceiverZeroRequiredResponses.selector, receiver
            )
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

    function expectRevertGuardZeroAddress() internal {
        vm.expectRevert(IInterchainClientV1.InterchainClientV1__GuardZeroAddress.selector);
    }

    function expectRevertExecutionServiceZeroAddress() internal {
        vm.expectRevert(IInterchainClientV1.InterchainClientV1__ExecutionServiceZeroAddress.selector);
    }

    function expectRevertReceiverZeroAddress() internal {
        vm.expectRevert(IInterchainClientV1.InterchainClientV1__ReceiverZeroAddress.selector);
    }

    function expectRevertVersionInvalid(uint8 version) internal {
        vm.expectRevert(abi.encodeWithSelector(OptionsLib.OptionsLib__VersionInvalid.selector, version));
    }

    function expectRevertOwnableUnauthorizedAccount(address account) internal {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, account));
    }

    // ══════════════════════════════════════════════ EXPECT (EVENTS) ══════════════════════════════════════════════════

    function expectEventGuardSet(address guard) internal {
        vm.expectEmit(address(icClient));
        emit DefaultGuardSet(guard);
    }

    function expectEventLinkedClientSet(uint64 chainId, bytes32 client) internal {
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

    function expectEventExecutionProofWritten(
        bytes32 transactionId,
        uint64 localDbNonce,
        uint64 localEntryIndex,
        address executor
    )
        internal
    {
        vm.expectEmit(address(icClient));
        emit ExecutionProofWritten({
            transactionId: transactionId,
            dbNonce: localDbNonce,
            entryIndex: localEntryIndex,
            executor: executor
        });
    }

    // ════════════════════════════════════════════════ ASSERTIONS ═════════════════════════════════════════════════════

    function assertCorrectDescriptor(
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc
    )
        internal
        view
    {
        assertEq(desc.dbNonce, icTx.dbNonce, "!desc.dbNonce");
        assertEq(desc.entryIndex, icTx.entryIndex, "!desc.entryIndex");
        assertEq(desc.transactionId, keccak256(getEncodedTx(icTx)), "!desc.transactionId");
    }

    function assertEq(InterchainTransaction memory icTx, InterchainTransaction memory expected) internal pure {
        assertEq(icTx.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(icTx.srcSender, expected.srcSender, "!srcSender");
        assertEq(icTx.dstChainId, expected.dstChainId, "!dstChainId");
        assertEq(icTx.dstReceiver, expected.dstReceiver, "!dstReceiver");
        assertEq(icTx.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(icTx.entryIndex, expected.entryIndex, "!entryIndex");
        assertEq(icTx.options, expected.options, "!options");
        assertEq(icTx.message, expected.message, "!message");
    }

    function getEncodedTx(InterchainTransaction memory icTx) internal view returns (bytes memory) {
        return payloadLibHarness.encodeVersionedPayload(CLIENT_VERSION, txLibHarness.encodeTransaction(icTx));
    }

    // ═══════════════════════════════════════════════════ UTILS ═══════════════════════════════════════════════════════

    function toArray(uint256 a) internal pure returns (uint256[] memory arr) {
        arr = new uint256[](1);
        arr[0] = a;
    }

    function toArr(uint256 a, uint256 b) internal pure returns (uint256[] memory arr) {
        arr = new uint256[](2);
        arr[0] = a;
        arr[1] = b;
    }
}
