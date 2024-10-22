// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IFastBridgeRouter, SwapQuery} from "../contracts/legacy/rfq/interfaces/IFastBridgeRouter.sol";

import {FastBridgeV2SrcTest, IFastBridge} from "./FastBridgeV2.Src.t.sol";

// solhint-disable func-name-mixedcase
contract FastBridgeV2SrcFBRouterV2Test is FastBridgeV2SrcTest {
    address public router;

    error FastBridgeRouterV2__OriginSenderNotSpecified();

    function setUp() public virtual override {
        super.setUp();
        router = deployCode("FastBridgeRouterV2", abi.encode(address(this)));
        IFastBridgeRouter(router).setFastBridge(address(fastBridge));
        vm.prank(userA);
        srcToken.approve(router, type(uint256).max);
        vm.prank(userB);
        srcToken.approve(router, type(uint256).max);
    }

    /// @notice Override bridge function to leverage the old FastBridgeRouterV2 contract
    function bridge(address caller, uint256 msgValue, IFastBridge.BridgeParams memory params) public virtual override {
        bytes memory destQueryParams = abi.encodePacked(bytes1(0x00), params.sender);
        vm.prank({msgSender: caller, txOrigin: caller});
        IFastBridgeRouter(router).bridge{value: msgValue}({
            recipient: params.to,
            chainId: params.dstChainId,
            token: params.originToken,
            amount: params.originAmount,
            originQuery: SwapQuery({
                routerAdapter: address(0),
                tokenOut: params.originToken,
                minAmountOut: params.originAmount,
                deadline: block.timestamp,
                rawParams: ""
            }),
            destQuery: SwapQuery({
                routerAdapter: address(0),
                tokenOut: params.destToken,
                minAmountOut: params.destAmount,
                deadline: params.deadline,
                rawParams: destQueryParams
            })
        });
    }

    // ══════════════════════════════════════════════ TEST OVERRIDES ═══════════════════════════════════════════════════

    function test_bridge_token_revert_approvedZero() public virtual override {
        vm.prank(userA);
        srcToken.approve(router, 0);
        vm.expectRevert();
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_token_revert_approvedNotEnough() public virtual override {
        vm.prank(userA);
        srcToken.approve(router, tokenParams.originAmount - 1);
        vm.expectRevert();
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_eth_revert_zeroMsgValue() public virtual override {
        vm.expectRevert(TokenNotContract.selector);
        bridge({caller: userA, msgValue: 0, params: ethParams});
    }

    function test_bridge_revert_zeroOriginToken() public virtual override {
        tokenParams.originToken = address(0);
        vm.expectRevert(TokenNotContract.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroSender() public virtual override {
        vm.etch(userA, "Mock contract code");
        tokenParams.sender = address(0);
        vm.expectRevert(FastBridgeRouterV2__OriginSenderNotSpecified.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }
}
