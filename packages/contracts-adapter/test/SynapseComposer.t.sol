// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseComposer} from "../src/SynapseComposer.sol";
import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";
import {SynapseOFTAdapterFactory} from "../src/SynapseOFTAdapterFactory.sol";

import {EndpointMock} from "./mocks/EndpointMock.sol";
import {TestToken} from "./mocks/TestToken.sol";

import {ICoreWriter} from "@layerzerolabs/hyperliquid-composer/contracts/interfaces/ICoreWriter.sol";
import {IHyperLiquidComposer} from "@layerzerolabs/hyperliquid-composer/contracts/interfaces/IHyperLiquidComposer.sol";
import {IRecoverableComposer} from "@layerzerolabs/hyperliquid-composer/contracts/interfaces/IRecoverableComposer.sol";
import {SendParam} from "@layerzerolabs/oft-evm/contracts/interfaces/IOFT.sol";
import {OFTComposeMsgCodec} from "@layerzerolabs/oft-evm/contracts/libs/OFTComposeMsgCodec.sol";
import {Test} from "forge-std/Test.sol";

contract SynapseComposerTest is Test {
    uint32 internal constant SRC_EID = 30_101;
    uint64 internal constant CORE_INDEX_ID = 235;
    int8 internal constant ASSET_DECIMAL_DIFF = 10;
    uint64 internal constant HYPE_CORE_INDEX_TESTNET = 1105;
    address internal constant CORE_WRITER = 0x3333333333333333333333333333333333333333;
    address internal constant SPOT_BALANCE_PRECOMPILE = address(0x0801);
    bytes4 internal constant SPOT_SEND_HEADER = 0x01000006;
    bytes4 internal constant ENDPOINT_SEND_SELECTOR =
        bytes4(keccak256("send((uint32,bytes32,bytes,bytes,bool),address)"));
    bytes32 internal constant GUID = keccak256("Malformed compose");

    EndpointMock internal endpoint;
    TestToken internal token;
    SynapseOFTAdapter internal adapter;
    SynapseComposer internal composer;

    address internal recovery = makeAddr("Recovery");
    address internal user = makeAddr("User");

    function setUp() public {
        endpoint = new EndpointMock();
        token = new TestToken();

        SynapseOFTAdapterFactory factory = new SynapseOFTAdapterFactory(address(this));
        factory.initialize(address(endpoint));
        adapter = SynapseOFTAdapter(factory.deploy(address(token), bytes32(0)));
        adapter.setPeer(SRC_EID, bytes32(uint256(uint160(address(adapter)))));

        composer = new SynapseComposer(address(adapter), CORE_INDEX_ID, ASSET_DECIMAL_DIFF, recovery);
        vm.etch(CORE_WRITER, hex"00");
    }

    function testConstructorRevertsForZeroRecoveryAddress() public {
        vm.expectRevert(IRecoverableComposer.InvalidRecoveryAddress.selector);
        new SynapseComposer(address(adapter), CORE_INDEX_ID, ASSET_DECIMAL_DIFF, address(0));
    }

    function testLzComposeRevertsForNonEndpoint() public {
        vm.expectRevert(IHyperLiquidComposer.OnlyEndpoint.selector);
        vm.prank(user);
        composer.lzCompose(address(adapter), GUID, _malformedMessage(1 ether), user, "");
    }

    function testLzComposeRevertsForUnauthorizedOFT() public {
        address unauthorizedOFT = makeAddr("Unauthorized OFT");

        vm.expectRevert(
            abi.encodeWithSelector(
                IHyperLiquidComposer.InvalidComposeCaller.selector, address(adapter), unauthorizedOFT
            )
        );
        vm.prank(address(endpoint));
        composer.lzCompose(unauthorizedOFT, GUID, _malformedMessage(1 ether), user, "");
    }

    function testMalformedComposeRefundToSrcBurnsComposerTokensThroughAdapter() public {
        uint256 amount = 1 ether;
        token.mintTestTokens(address(composer), amount);

        vm.prank(address(endpoint));
        composer.lzCompose(address(adapter), GUID, _malformedMessage(amount), user, "");

        (SendParam memory refundSendParam, uint256 msgValue) = composer.failedMessages(GUID);
        assertEq(refundSendParam.dstEid, SRC_EID);
        assertEq(refundSendParam.to, bytes32(uint256(uint160(user))));
        assertEq(refundSendParam.amountLD, amount);
        assertEq(msgValue, 0);
        assertEq(token.balanceOf(address(composer)), amount);
        assertEq(token.totalSupply(), amount);

        composer.refundToSrc(GUID);

        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.totalSupply(), 0);
        assertEq(token.allowance(address(composer), address(adapter)), type(uint256).max);
        (SendParam memory clearedSendParam,) = composer.failedMessages(GUID);
        assertEq(clearedSendParam.dstEid, 0);
    }

    function testRefundToSrcEndpointFailureRollsBackAndAllowsRetry() public {
        uint256 amount = 1 ether;
        bytes memory endpointError = abi.encodeWithSignature("Error(string)", "Endpoint send failed");
        token.mintTestTokens(address(composer), amount);

        vm.prank(address(endpoint));
        composer.lzCompose(address(adapter), GUID, _malformedMessage(amount), user, "");

        vm.mockCallRevert(address(endpoint), abi.encodeWithSelector(ENDPOINT_SEND_SELECTOR), endpointError);
        vm.expectRevert(endpointError);
        composer.refundToSrc(GUID);

        (SendParam memory retainedSendParam,) = composer.failedMessages(GUID);
        assertEq(retainedSendParam.dstEid, SRC_EID);
        assertEq(token.balanceOf(address(composer)), amount);
        assertEq(token.totalSupply(), amount);

        vm.clearMockedCalls();
        composer.refundToSrc(GUID);

        (SendParam memory clearedSendParam,) = composer.failedMessages(GUID);
        assertEq(clearedSendParam.dstEid, 0);
        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.totalSupply(), 0);
    }

    function testRecoveryFunctionsRevertForNonRecoveryAddress() public {
        vm.startPrank(user);

        vm.expectRevert(IRecoverableComposer.NotRecoveryAddress.selector);
        composer.recoverEvmERC20(0);

        vm.expectRevert(IRecoverableComposer.NotRecoveryAddress.selector);
        composer.recoverEvmNative(0);

        vm.expectRevert(IRecoverableComposer.NotRecoveryAddress.selector);
        composer.retrieveCoreERC20(0);

        vm.expectRevert(IRecoverableComposer.NotRecoveryAddress.selector);
        composer.retrieveCoreHYPE(0);

        vm.stopPrank();
    }

    function testRecoverEvmERC20SupportsPartialAndFullRecovery() public {
        uint256 amount = 5 ether;
        uint256 fullTransfer = composer.FULL_TRANSFER();
        token.mintTestTokens(address(composer), amount);

        vm.prank(recovery);
        composer.recoverEvmERC20(2 ether);

        assertEq(token.balanceOf(recovery), 2 ether);
        assertEq(token.balanceOf(address(composer)), 3 ether);

        vm.prank(recovery);
        composer.recoverEvmERC20(fullTransfer);

        assertEq(token.balanceOf(recovery), amount);
        assertEq(token.balanceOf(address(composer)), 0);
    }

    function testRecoverEvmNativeSupportsPartialAndFullRecovery() public {
        uint256 fullTransfer = composer.FULL_TRANSFER();
        vm.deal(address(composer), 5 ether);

        vm.prank(recovery);
        composer.recoverEvmNative(2 ether);

        assertEq(recovery.balance, 2 ether);
        assertEq(address(composer).balance, 3 ether);

        vm.prank(recovery);
        composer.recoverEvmNative(fullTransfer);

        assertEq(recovery.balance, 5 ether);
        assertEq(address(composer).balance, 0);
    }

    function testRetrieveCoreERC20FullTransferSubmitsCoreWriterAction() public {
        uint64 coreBalance = 123_456;
        uint64 fullTransfer = uint64(composer.FULL_TRANSFER());
        _mockSpotBalance(CORE_INDEX_ID, coreBalance);
        _expectCoreWriterTransfer(composer.ERC20_ASSET_BRIDGE(), CORE_INDEX_ID, coreBalance);

        vm.prank(recovery);
        composer.retrieveCoreERC20(fullTransfer);
    }

    function testRetrieveCoreHYPESubmitsRequestedCoreWriterAction() public {
        uint64 coreBalance = 500_000;
        uint64 requestedAmount = 400_000;
        _mockSpotBalance(HYPE_CORE_INDEX_TESTNET, coreBalance);
        _expectCoreWriterTransfer(composer.NATIVE_ASSET_BRIDGE(), HYPE_CORE_INDEX_TESTNET, requestedAmount);

        vm.prank(recovery);
        composer.retrieveCoreHYPE(requestedAmount);
    }

    function testRetrieveCoreERC20RevertsWhenRequestedAmountExceedsBalance() public {
        uint64 coreBalance = 100;
        uint64 requestedAmount = 101;
        _mockSpotBalance(CORE_INDEX_ID, coreBalance);

        vm.expectRevert(
            abi.encodeWithSelector(
                IRecoverableComposer.MaxRetrieveAmountExceeded.selector, coreBalance, requestedAmount
            )
        );
        vm.prank(recovery);
        composer.retrieveCoreERC20(requestedAmount);
    }

    function testConstructorSetsComposerConfigurationAndAdapterAllowance() public view {
        assertEq(composer.ENDPOINT(), address(endpoint));
        assertEq(composer.OFT(), address(adapter));
        assertEq(composer.ERC20(), address(token));
        assertEq(composer.ERC20_CORE_INDEX_ID(), CORE_INDEX_ID);
        assertEq(composer.ERC20_DECIMAL_DIFF(), ASSET_DECIMAL_DIFF);
        assertEq(composer.RECOVERY_ADDRESS(), recovery);
        assertEq(token.allowance(address(composer), address(adapter)), type(uint256).max);
    }

    function _mockSpotBalance(uint64 coreIndexId, uint64 balance) internal {
        vm.mockCall(
            SPOT_BALANCE_PRECOMPILE,
            abi.encode(address(composer), coreIndexId),
            abi.encode(balance, uint64(0), uint64(0))
        );
    }

    function _expectCoreWriterTransfer(address to, uint64 coreIndexId, uint64 amount) internal {
        bytes memory payload = abi.encodePacked(SPOT_SEND_HEADER, abi.encode(to, coreIndexId, amount));
        vm.expectCall(CORE_WRITER, abi.encodeCall(ICoreWriter.sendRawAction, (payload)));
    }

    function _malformedMessage(uint256 amount) internal view returns (bytes memory) {
        return OFTComposeMsgCodec.encode(
            1, SRC_EID, amount, abi.encodePacked(bytes32(uint256(uint160(user))), hex"deadbeef")
        );
    }
}
