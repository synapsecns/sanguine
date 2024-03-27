// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1} from "../../contracts/libs/Options.sol";
import {InterchainTxDescriptor} from "../../contracts/libs/InterchainTransaction.sol";

import {InterchainAppV1Test} from "./InterchainAppV1.t.sol";

import {InterchainClientV1Mock} from "../mocks/InterchainClientV1Mock.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract InterchainAppV1MessagingTest is InterchainAppV1Test {
    uint256 public constant MOCK_DB_NONCE = 42;
    uint64 public constant MOCK_ENTRY_INDEX = 4;

    uint256 public constant MOCK_IC_FEE = 4844;

    OptionsV1 public options = OptionsV1({gasLimit: 100_000, gasAirdrop: 0.2 ether});
    bytes public encodedOptions = options.encodeOptionsV1();
    bytes public message = "Hello, Interchain!";

    address public extraClient = makeAddr("Extra Client");

    InterchainTxDescriptor public mockDesc = InterchainTxDescriptor({
        transactionId: keccak256("Mock Transaction ID"),
        dbNonce: MOCK_DB_NONCE,
        entryIndex: MOCK_ENTRY_INDEX
    });

    event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message);

    function configureICAppV1() internal override {
        appHarness.addInterchainClient({client: icClient, updateLatest: true});
        appHarness.addInterchainClient({client: extraClient, updateLatest: false});
        appHarness.linkRemoteApp({chainId: REMOTE_CHAIN_ID, remoteApp: linkedAppMockBytes32});
        appHarness.addTrustedModule(moduleMock);
        appHarness.setAppConfigV1(appConfig);
        appHarness.setExecutionService(execServiceMock);
    }

    function mockInterchainSendResult() internal {
        vm.mockCall(icClient, getExpectedCalldata(), abi.encode(mockDesc));
    }

    function mockInterchainFeeResult() internal {
        bytes memory expectedCalldata = abi.encodeCall(
            InterchainClientV1Mock.getInterchainFee,
            (REMOTE_CHAIN_ID, execServiceMock, toArray(moduleMock), encodedOptions, message)
        );
        vm.mockCall(icClient, expectedCalldata, abi.encode(MOCK_IC_FEE));
    }

    function expectClientCallInterchainSend() internal {
        bytes memory expectedCalldata = getExpectedCalldata();
        vm.expectCall({callee: icClient, msgValue: MOCK_IC_FEE, data: expectedCalldata, count: 1});
    }

    function getExpectedCalldata() internal view returns (bytes memory) {
        return abi.encodeCall(
            InterchainClientV1Mock.interchainSend,
            (REMOTE_CHAIN_ID, linkedAppMockBytes32, execServiceMock, toArray(moduleMock), encodedOptions, message)
        );
    }

    function assertEq(InterchainTxDescriptor memory desc, InterchainTxDescriptor memory expected) internal {
        assertEq(desc.transactionId, expected.transactionId);
        assertEq(desc.dbNonce, expected.dbNonce);
        assertEq(desc.entryIndex, expected.entryIndex);
    }

    // ══════════════════════════════════════════ TESTS: RECEIVE MESSAGES ══════════════════════════════════════════════

    function test_appReceive_noAirdrop() public {
        vm.expectEmit(address(appHarness));
        emit MessageReceived(REMOTE_CHAIN_ID, linkedAppMockBytes32, MOCK_DB_NONCE, MOCK_ENTRY_INDEX, message);
        vm.prank(icClient);
        appHarness.appReceive({
            srcChainId: REMOTE_CHAIN_ID,
            sender: linkedAppMockBytes32,
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            message: message
        });
    }

    function test_appReceive_withAirdrop() public {
        deal(icClient, 1 ether);
        vm.expectEmit(address(appHarness));
        emit MessageReceived(REMOTE_CHAIN_ID, linkedAppMockBytes32, MOCK_DB_NONCE, MOCK_ENTRY_INDEX, message);
        vm.prank(icClient);
        appHarness.appReceive{value: 1 ether}({
            srcChainId: REMOTE_CHAIN_ID,
            sender: linkedAppMockBytes32,
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            message: message
        });
    }

    function test_appReceive_revert_notInterchainClients(address caller) public {
        vm.assume(caller != icClient && caller != extraClient);
        expectRevertNotInterchainClient(caller);
        vm.prank(caller);
        appHarness.appReceive({
            srcChainId: REMOTE_CHAIN_ID,
            sender: linkedAppMockBytes32,
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            message: message
        });
    }

    function test_appReceive_revert_sameChainId() public {
        expectRevertSameChainId(LOCAL_CHAIN_ID);
        vm.prank(icClient);
        appHarness.appReceive({
            srcChainId: LOCAL_CHAIN_ID,
            sender: linkedAppMockBytes32,
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            message: message
        });
    }

    function test_appReceive_revert_senderNotAllowed(bytes32 sender) public {
        vm.assume(sender != linkedAppMockBytes32);
        expectRevertSenderNotAllowed(REMOTE_CHAIN_ID, sender);
        vm.prank(icClient);
        appHarness.appReceive({
            srcChainId: REMOTE_CHAIN_ID,
            sender: sender,
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            message: message
        });
    }

    function test_getReceivingConfig() public {
        (bytes memory encodedConfig, address[] memory modules) = appHarness.getReceivingConfig();
        assertEq(encodedConfig, appConfig.encodeAppConfigV1());
        assertEq(modules, toArray(moduleMock));
    }

    // ══════════════════════════════════════════ TESTS: SENDING MESSAGES ══════════════════════════════════════════════

    function test_sendInterchainMessage() public {
        deal(address(appHarness), MOCK_IC_FEE);
        mockInterchainSendResult();
        expectClientCallInterchainSend();
        InterchainTxDescriptor memory desc = appHarness.exposed__sendInterchainMessage({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: linkedAppMockBytes32,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
        assertEq(desc, mockDesc);
    }

    function test_sendInterchainMessage_revert_latestClientNotSet() public {
        vm.prank(governor);
        appHarness.setLatestInterchainClient(address(0));
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertInterchainClientZeroAddress();
        appHarness.exposed__sendInterchainMessage({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: linkedAppMockBytes32,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendInterchainMessage_revert_balanceTooLow() public {
        deal(address(appHarness), MOCK_IC_FEE - 1);
        expectRevertBalanceTooLow({actual: MOCK_IC_FEE - 1, required: MOCK_IC_FEE});
        appHarness.exposed__sendInterchainMessage({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: linkedAppMockBytes32,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendInterchainMessage_revert_sameChainId() public {
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertSameChainId(LOCAL_CHAIN_ID);
        appHarness.exposed__sendInterchainMessage({
            dstChainId: LOCAL_CHAIN_ID,
            receiver: linkedAppMockBytes32,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendInterchainMessage_revert_receiverNotSet() public {
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertReceiverNotSet(REMOTE_CHAIN_ID);
        appHarness.exposed__sendInterchainMessage({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: 0,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendInterchainMessageEVM() public {
        deal(address(appHarness), MOCK_IC_FEE);
        mockInterchainSendResult();
        expectClientCallInterchainSend();
        InterchainTxDescriptor memory desc = appHarness.exposed__sendInterchainMessageEVM({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: linkedAppMock,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
        assertEq(desc, mockDesc);
    }

    function test_sendInterchainMessageEVM_revert_latestClientNotSet() public {
        vm.prank(governor);
        appHarness.setLatestInterchainClient(address(0));
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertInterchainClientZeroAddress();
        appHarness.exposed__sendInterchainMessageEVM({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: linkedAppMock,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendInterchainMessageEVM_revert_balanceTooLow() public {
        deal(address(appHarness), MOCK_IC_FEE - 1);
        expectRevertBalanceTooLow({actual: MOCK_IC_FEE - 1, required: MOCK_IC_FEE});
        appHarness.exposed__sendInterchainMessageEVM({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: linkedAppMock,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendInterchainMessageEVM_revert_sameChainId() public {
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertSameChainId(LOCAL_CHAIN_ID);
        appHarness.exposed__sendInterchainMessageEVM({
            dstChainId: LOCAL_CHAIN_ID,
            receiver: linkedAppMock,
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendInterchainMessageEVM_revert_receiverNotSet() public {
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertReceiverNotSet(REMOTE_CHAIN_ID);
        appHarness.exposed__sendInterchainMessageEVM({
            dstChainId: REMOTE_CHAIN_ID,
            receiver: address(0),
            messageFee: MOCK_IC_FEE,
            options: encodedOptions,
            message: message
        });
    }

    function test_sendToLinkedApp() public {
        deal(address(appHarness), MOCK_IC_FEE);
        mockInterchainSendResult();
        expectClientCallInterchainSend();
        InterchainTxDescriptor memory desc = appHarness.exposed__sendToLinkedApp({
            dstChainId: REMOTE_CHAIN_ID,
            messageFee: MOCK_IC_FEE,
            options: options,
            message: message
        });
        assertEq(desc, mockDesc);
    }

    function test_sendToLinkedApp_revert_latestClientNotSet() public {
        vm.prank(governor);
        appHarness.setLatestInterchainClient(address(0));
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertInterchainClientZeroAddress();
        appHarness.exposed__sendToLinkedApp({
            dstChainId: REMOTE_CHAIN_ID,
            messageFee: MOCK_IC_FEE,
            options: options,
            message: message
        });
    }

    function test_sendToLinkedApp_revert_balanceTooLow() public {
        deal(address(appHarness), MOCK_IC_FEE - 1);
        expectRevertBalanceTooLow({actual: MOCK_IC_FEE - 1, required: MOCK_IC_FEE});
        appHarness.exposed__sendToLinkedApp({
            dstChainId: REMOTE_CHAIN_ID,
            messageFee: MOCK_IC_FEE,
            options: options,
            message: message
        });
    }

    function test_sendToLinkedApp_revert_sameChainId() public {
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertSameChainId(LOCAL_CHAIN_ID);
        appHarness.exposed__sendToLinkedApp({
            dstChainId: LOCAL_CHAIN_ID,
            messageFee: MOCK_IC_FEE,
            options: options,
            message: message
        });
    }

    function test_sendToLinkedApp_revert_receiverNotSet() public {
        deal(address(appHarness), MOCK_IC_FEE);
        expectRevertReceiverNotSet(UNKNOWN_CHAIN_ID);
        appHarness.exposed__sendToLinkedApp({
            dstChainId: UNKNOWN_CHAIN_ID,
            messageFee: MOCK_IC_FEE,
            options: options,
            message: message
        });
    }

    function test_getInterchainFee() public {
        mockInterchainFeeResult();
        uint256 fee = appHarness.exposed__getInterchainFee({
            dstChainId: REMOTE_CHAIN_ID,
            options: encodedOptions,
            message: message
        });
        assertEq(fee, MOCK_IC_FEE);
    }

    function test_getInterchainFee_revert_latestClientNotSet() public {
        vm.prank(governor);
        appHarness.setLatestInterchainClient(address(0));
        expectRevertInterchainClientZeroAddress();
        appHarness.exposed__getInterchainFee({dstChainId: REMOTE_CHAIN_ID, options: encodedOptions, message: message});
    }

    function test_getMessageFee() public {
        mockInterchainFeeResult();
        uint256 fee =
            appHarness.exposed__getMessageFee({dstChainId: REMOTE_CHAIN_ID, options: options, message: message});
        assertEq(fee, MOCK_IC_FEE);
    }

    function test_getMessageFee_revert_latestClientNotSet() public {
        vm.prank(governor);
        appHarness.setLatestInterchainClient(address(0));
        expectRevertInterchainClientZeroAddress();
        appHarness.exposed__getMessageFee({dstChainId: REMOTE_CHAIN_ID, options: options, message: message});
    }
}
