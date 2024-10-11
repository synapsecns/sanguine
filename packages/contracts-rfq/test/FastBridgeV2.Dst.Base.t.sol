// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BridgeTransactionV2Lib} from "../contracts/libs/BridgeTransactionV2.sol";

import {FastBridgeV2, FastBridgeV2Test, IFastBridgeV2} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2DstBaseTest is FastBridgeV2Test {
    uint256 public constant LEFTOVER_BALANCE = 1 ether;

    function setUp() public virtual override {
        vm.chainId(DST_CHAIN_ID);
        super.setUp();
    }

    function deployFastBridge() public override returns (FastBridgeV2) {
        return new FastBridgeV2(address(this));
    }

    function mintTokens() public virtual override {
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

    function relay(address caller, uint256 msgValue, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public {
        bytes memory request = BridgeTransactionV2Lib.encodeV2(bridgeTx);
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.relay{value: msgValue}(request);
    }

    function relayWithAddress(
        address caller,
        address relayer,
        uint256 msgValue,
        IFastBridgeV2.BridgeTransactionV2 memory bridgeTx
    )
        public
    {
        bytes memory request = BridgeTransactionV2Lib.encodeV2(bridgeTx);
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.relay{value: msgValue}(request, relayer);
    }
}
