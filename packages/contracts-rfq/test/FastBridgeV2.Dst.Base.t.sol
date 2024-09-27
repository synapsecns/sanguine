// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2, FastBridgeV2Test, IFastBridge} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2DstBaseTest is FastBridgeV2Test {
    uint256 public constant LEFTOVER_BALANCE = 1 ether;

    function setUp() public override {
        vm.chainId(DST_CHAIN_ID);
        super.setUp();
    }

    function deployFastBridge() public override returns (FastBridgeV2) {
        return new FastBridgeV2(address(this));
    }

    function mintTokens() public override {
        dstToken.mint(address(relayerA), LEFTOVER_BALANCE + tokenParams.destAmount);
        dstToken.mint(address(relayerB), LEFTOVER_BALANCE + tokenParams.destAmount);
        deal(relayerA, LEFTOVER_BALANCE + ethParams.destAmount);
        deal(relayerB, LEFTOVER_BALANCE + ethParams.destAmount);
        vm.prank(relayerA);
        dstToken.approve(address(fastBridge), type(uint256).max);
        vm.prank(relayerB);
        dstToken.approve(address(fastBridge), type(uint256).max);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function relay(address caller, uint256 msgValue, IFastBridge.BridgeTransaction memory bridgeTx) public {
        bytes memory request = abi.encode(bridgeTx);
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.relay{value: msgValue}(request);
    }

    function relayWithAddress(
        address caller,
        address relayer,
        uint256 msgValue,
        IFastBridge.BridgeTransaction memory bridgeTx
    )
        public
    {
        bytes memory request = abi.encode(bridgeTx);
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.relay{value: msgValue}(request, relayer);
    }
}
