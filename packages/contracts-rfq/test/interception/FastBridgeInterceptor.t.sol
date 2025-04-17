// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {FastBridgeInterceptor, IFastBridge, IFastBridgeInterceptor} from "../../contracts/FastBridgeInterceptor.sol";
import {MockERC20} from "../mocks/MockERC20.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeInterceptorTest is Test {
    uint32 internal constant SRC_CHAIN_ID = 1337;
    uint32 internal constant DST_CHAIN_ID = 7331;
    uint256 internal constant DEADLINE = 1 days;
    address internal constant ETH = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    FastBridgeInterceptor internal fbi;
    address internal fastBridge;

    MockERC20 internal srcToken;
    address internal dstToken = makeAddr("DstToken");

    address internal user = makeAddr("User");
    address internal recipient = makeAddr("Recipient");

    event BridgeRequested(
        bytes32 indexed transactionId,
        address indexed sender,
        bytes request,
        uint32 destChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        bool sendChainGas
    );

    function setUp() public {
        vm.chainId(SRC_CHAIN_ID);
        fbi = new FastBridgeInterceptor();
        // Use the cheatcode to deploy 0.8.20 contract within a 0.8.24 test
        fastBridge = deployCode({what: "FastBridge", args: abi.encode(address(this))});
        srcToken = new MockERC20("SrcToken", 18);
        srcToken.mint(user, 10 ether);
        vm.prank(user);
        srcToken.approve(address(fbi), type(uint256).max);
        deal(user, 10 ether);
    }

    function getOriginValue(uint256 valueWei) internal pure virtual returns (uint256) {
        return valueWei;
    }

    function getDestValue(uint256 valueWei) internal pure virtual returns (uint256) {
        return valueWei;
    }

    function expectEventBridgeRequested(IFastBridge.BridgeParams memory params, uint256 expectedDestAmount) internal {
        IFastBridge.BridgeTransaction memory expectedTx = IFastBridge.BridgeTransaction({
            originChainId: SRC_CHAIN_ID,
            destChainId: params.dstChainId,
            originSender: params.sender,
            destRecipient: params.to,
            originToken: params.originToken,
            destToken: params.destToken,
            originAmount: params.originAmount,
            destAmount: expectedDestAmount,
            originFeeAmount: 0,
            sendChainGas: params.sendChainGas,
            deadline: params.deadline,
            nonce: 0
        });
        bytes memory request = abi.encode(expectedTx);
        bytes32 transactionId = keccak256(request);
        vm.expectEmit(fastBridge);
        emit BridgeRequested({
            transactionId: transactionId,
            sender: params.sender,
            request: request,
            destChainId: params.dstChainId,
            originToken: params.originToken,
            destToken: params.destToken,
            originAmount: params.originAmount,
            destAmount: expectedDestAmount,
            sendChainGas: params.sendChainGas
        });
    }

    function expectRevertOriginAmountOutOfRange(uint256 originAmount, uint256 quoteOriginAmount) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IFastBridgeInterceptor.FBI__OriginAmountOutOfRange.selector, originAmount, quoteOriginAmount
            )
        );
    }

    function expectRevertTokenNotContract(address token) internal {
        vm.expectRevert(abi.encodeWithSelector(IFastBridgeInterceptor.FBI__TokenNotContract.selector, token));
    }

    function generateBridgeParams(
        address token,
        uint256 originAmount,
        uint256 destAmount
    )
        internal
        view
        returns (IFastBridge.BridgeParams memory)
    {
        return IFastBridge.BridgeParams({
            dstChainId: DST_CHAIN_ID,
            sender: user,
            to: recipient,
            originToken: token,
            destToken: dstToken,
            originAmount: originAmount,
            destAmount: destAmount,
            sendChainGas: false,
            deadline: block.timestamp + DEADLINE
        });
    }

    function generateInterceptorParams(uint256 quoteOriginAmount)
        internal
        view
        returns (IFastBridgeInterceptor.InterceptorParams memory)
    {
        return IFastBridgeInterceptor.InterceptorParams({fastBridge: fastBridge, quoteOriginAmount: quoteOriginAmount});
    }

    function checkHappyPathToken(
        uint256 quoteOriginAmount,
        uint256 quoteDestAmount,
        uint256 originAmount,
        uint256 expectedDestAmount
    )
        internal
    {
        IFastBridge.BridgeParams memory params = generateBridgeParams(address(srcToken), originAmount, quoteDestAmount);
        IFastBridgeInterceptor.InterceptorParams memory icParams = generateInterceptorParams(quoteOriginAmount);
        expectEventBridgeRequested(params, expectedDestAmount);
        vm.prank(user);
        fbi.bridgeWithInterception(params, icParams);
    }

    function checkRevertTokenOriginAmountOutOfRange(
        uint256 quoteOriginAmount,
        uint256 quoteDestAmount,
        uint256 originAmount
    )
        internal
    {
        IFastBridge.BridgeParams memory params = generateBridgeParams(address(srcToken), originAmount, quoteDestAmount);
        IFastBridgeInterceptor.InterceptorParams memory icParams = generateInterceptorParams(quoteOriginAmount);
        expectRevertOriginAmountOutOfRange(originAmount, quoteOriginAmount);
        vm.prank(user);
        fbi.bridgeWithInterception(params, icParams);
    }

    function checkRevertTokenNotContract(
        uint256 quoteOriginAmount,
        uint256 quoteDestAmount,
        uint256 originAmount
    )
        internal
    {
        IFastBridge.BridgeParams memory params = generateBridgeParams(address(srcToken), originAmount, quoteDestAmount);
        IFastBridgeInterceptor.InterceptorParams memory icParams = generateInterceptorParams(quoteOriginAmount);
        params.originToken = makeAddr("Not Contract");
        expectRevertTokenNotContract(params.originToken);
        vm.prank(user);
        fbi.bridgeWithInterception(params, icParams);
    }

    function test_bridge_token_1_1_originAmountSame() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1 ether),
            expectedDestAmount: getDestValue(1 ether)
        });
    }

    function test_bridge_token_1_1_originAmountLower() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(0.995 ether),
            expectedDestAmount: getDestValue(0.995 ether)
        });
    }

    function test_bridge_token_1_1_originAmountMinAllowed() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(0.99 ether),
            expectedDestAmount: getDestValue(0.99 ether)
        });
    }

    function test_bridge_token_1_1_originAmountHigher() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1.005 ether),
            expectedDestAmount: getDestValue(1.005 ether)
        });
    }

    function test_bridge_token_1_1_originAmountMaxAllowed() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1.01 ether),
            expectedDestAmount: getDestValue(1.01 ether)
        });
    }

    function test_bridge_token_1_1_revert_originAmountTooLow() public {
        checkRevertTokenOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(0.99 ether) - 1 wei
        });
    }

    function test_bridge_token_1_1_revert_originAmountTooHigh() public {
        checkRevertTokenOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1.01 ether) + 1 wei
        });
    }

    function test_bridge_token_1_1_revert_tokenNotContract() public {
        checkRevertTokenNotContract({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1 ether)
        });
    }

    function test_bridge_token_1_2_originAmountSame() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1 ether),
            expectedDestAmount: getDestValue(2 ether)
        });
    }

    function test_bridge_token_1_2_originAmountLower() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(0.995 ether),
            expectedDestAmount: getDestValue(1.99 ether)
        });
    }

    function test_bridge_token_1_2_originAmountMinAllowed() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(0.99 ether),
            expectedDestAmount: getDestValue(1.98 ether)
        });
    }

    function test_bridge_token_1_2_originAmountHigher() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1.005 ether),
            expectedDestAmount: getDestValue(2.01 ether)
        });
    }

    function test_bridge_token_1_2_originAmountMaxAllowed() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1.01 ether),
            expectedDestAmount: getDestValue(2.02 ether)
        });
    }

    function test_bridge_token_1_2_revert_originAmountTooLow() public {
        checkRevertTokenOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(0.99 ether) - 1 wei
        });
    }

    function test_bridge_token_1_2_revert_originAmountTooHigh() public {
        checkRevertTokenOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1.01 ether) + 1 wei
        });
    }

    function test_bridge_token_1_2_revert_tokenNotContract() public {
        checkRevertTokenNotContract({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1 ether)
        });
    }

    function test_bridge_token_2_1_originAmountSame() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1 ether),
            expectedDestAmount: getDestValue(0.5 ether)
        });
    }

    function test_bridge_token_2_1_originAmountLower() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(0.995 ether),
            expectedDestAmount: getDestValue(0.4975 ether)
        });
    }

    function test_bridge_token_2_1_originAmountMinAllowed() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(0.99 ether),
            expectedDestAmount: getDestValue(0.495 ether)
        });
    }

    function test_bridge_token_2_1_originAmountHigher() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1.005 ether),
            expectedDestAmount: getDestValue(0.5025 ether)
        });
    }

    function test_bridge_token_2_1_originAmountMaxAllowed() public {
        checkHappyPathToken({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1.01 ether),
            expectedDestAmount: getDestValue(0.505 ether)
        });
    }

    function test_bridge_token_2_1_revert_originAmountTooLow() public {
        checkRevertTokenOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(0.99 ether) - 1 wei
        });
    }

    function test_bridge_token_2_1_revert_originAmountTooHigh() public {
        checkRevertTokenOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1.01 ether) + 1 wei
        });
    }

    function test_bridge_token_2_1_revert_tokenNotContract() public {
        checkRevertTokenNotContract({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1 ether)
        });
    }

    function checkHappyPathEth(
        uint256 quoteOriginAmount,
        uint256 quoteDestAmount,
        uint256 originAmount,
        uint256 expectedDestAmount
    )
        public
    {
        IFastBridge.BridgeParams memory params = generateBridgeParams(ETH, originAmount, quoteDestAmount);
        IFastBridgeInterceptor.InterceptorParams memory icParams = generateInterceptorParams(quoteOriginAmount);
        expectEventBridgeRequested(params, expectedDestAmount);
        vm.prank(user);
        fbi.bridgeWithInterception{value: originAmount}(params, icParams);
    }

    function checkRevertEthOriginAmountOutOfRange(
        uint256 quoteOriginAmount,
        uint256 quoteDestAmount,
        uint256 originAmount
    )
        public
    {
        IFastBridge.BridgeParams memory params = generateBridgeParams(ETH, originAmount, quoteDestAmount);
        IFastBridgeInterceptor.InterceptorParams memory icParams = generateInterceptorParams(quoteOriginAmount);
        expectRevertOriginAmountOutOfRange(originAmount, quoteOriginAmount);
        vm.prank(user);
        fbi.bridgeWithInterception{value: originAmount}(params, icParams);
    }

    function test_bridge_eth_1_1_originAmountSame() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1 ether),
            expectedDestAmount: getDestValue(1 ether)
        });
    }

    function test_bridge_eth_1_1_originAmountLower() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(0.995 ether),
            expectedDestAmount: getDestValue(0.995 ether)
        });
    }

    function test_bridge_eth_1_1_originAmountMinAllowed() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(0.99 ether),
            expectedDestAmount: getDestValue(0.99 ether)
        });
    }

    function test_bridge_eth_1_1_originAmountHigher() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1.005 ether),
            expectedDestAmount: getDestValue(1.005 ether)
        });
    }

    function test_bridge_eth_1_1_originAmountMaxAllowed() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1.01 ether),
            expectedDestAmount: getDestValue(1.01 ether)
        });
    }

    function test_bridge_eth_1_1_revert_originAmountTooLow() public {
        checkRevertEthOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(0.99 ether) - 1 wei
        });
    }

    function test_bridge_eth_1_1_revert_originAmountTooHigh() public {
        checkRevertEthOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(1 ether),
            originAmount: getOriginValue(1.01 ether) + 1 wei
        });
    }

    function test_bridge_eth_1_2_originAmountSame() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1 ether),
            expectedDestAmount: getDestValue(2 ether)
        });
    }

    function test_bridge_eth_1_2_originAmountLower() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(0.995 ether),
            expectedDestAmount: getDestValue(1.99 ether)
        });
    }

    function test_bridge_eth_1_2_originAmountMinAllowed() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(0.99 ether),
            expectedDestAmount: getDestValue(1.98 ether)
        });
    }

    function test_bridge_eth_1_2_originAmountHigher() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1.005 ether),
            expectedDestAmount: getDestValue(2.01 ether)
        });
    }

    function test_bridge_eth_1_2_originAmountMaxAllowed() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1.01 ether),
            expectedDestAmount: getDestValue(2.02 ether)
        });
    }

    function test_bridge_eth_1_2_revert_originAmountTooLow() public {
        checkRevertEthOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(0.99 ether) - 1 wei
        });
    }

    function test_bridge_eth_1_2_revert_originAmountTooHigh() public {
        checkRevertEthOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(2 ether),
            originAmount: getOriginValue(1.01 ether) + 1 wei
        });
    }

    function test_bridge_eth_2_1_originAmountSame() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1 ether),
            expectedDestAmount: getDestValue(0.5 ether)
        });
    }

    function test_bridge_eth_2_1_originAmountLower() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(0.995 ether),
            expectedDestAmount: getDestValue(0.4975 ether)
        });
    }

    function test_bridge_eth_2_1_originAmountMinAllowed() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(0.99 ether),
            expectedDestAmount: getDestValue(0.495 ether)
        });
    }

    function test_bridge_eth_2_1_originAmountHigher() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1.005 ether),
            expectedDestAmount: getDestValue(0.5025 ether)
        });
    }

    function test_bridge_eth_2_1_originAmountMaxAllowed() public {
        checkHappyPathEth({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1.01 ether),
            expectedDestAmount: getDestValue(0.505 ether)
        });
    }

    function test_bridge_eth_2_1_revert_originAmountTooLow() public {
        checkRevertEthOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(0.99 ether) - 1 wei
        });
    }

    function test_bridge_eth_2_1_revert_originAmountTooHigh() public {
        checkRevertEthOriginAmountOutOfRange({
            quoteOriginAmount: getOriginValue(1 ether),
            quoteDestAmount: getDestValue(0.5 ether),
            originAmount: getOriginValue(1.01 ether) + 1 wei
        });
    }
}
