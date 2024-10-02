// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2, FastBridgeV2Test, IFastBridge, IFastBridgeV2} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
abstract contract FastBridgeV2SrcBaseTest is FastBridgeV2Test {
    uint256 public constant MIN_DEADLINE = 30 minutes;
    uint256 public constant CLAIM_DELAY = 30 minutes;
    uint256 public constant PERMISSIONLESS_REFUND_DELAY = 7 days;

    uint256 public constant LEFTOVER_BALANCE = 10 ether;
    uint256 public constant INITIAL_PROTOCOL_FEES_TOKEN = 456_789;
    uint256 public constant INITIAL_PROTOCOL_FEES_ETH = 0.123 ether;

    function setUp() public virtual override {
        vm.chainId(SRC_CHAIN_ID);
        super.setUp();
    }

    function deployFastBridge() public virtual override returns (FastBridgeV2) {
        return new FastBridgeV2(address(this));
    }

    function configureFastBridge() public virtual override {
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayerA);
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayerB);
        fastBridge.grantRole(fastBridge.GUARD_ROLE(), guard);
        fastBridge.grantRole(fastBridge.REFUNDER_ROLE(), refunder);
    }

    function mintTokens() public virtual override {
        // Prior Protocol fees
        srcToken.mint(address(fastBridge), INITIAL_PROTOCOL_FEES_TOKEN);
        deal(address(fastBridge), INITIAL_PROTOCOL_FEES_ETH);
        cheatCollectedProtocolFees(address(srcToken), INITIAL_PROTOCOL_FEES_TOKEN);
        cheatCollectedProtocolFees(ETH_ADDRESS, INITIAL_PROTOCOL_FEES_ETH);
        // Users
        srcToken.mint(userA, LEFTOVER_BALANCE + tokenParams.originAmount);
        srcToken.mint(userB, LEFTOVER_BALANCE + tokenParams.originAmount);
        deal(userA, LEFTOVER_BALANCE + ethParams.originAmount);
        deal(userB, LEFTOVER_BALANCE + ethParams.originAmount);
        vm.prank(userA);
        srcToken.approve(address(fastBridge), type(uint256).max);
        vm.prank(userB);
        srcToken.approve(address(fastBridge), type(uint256).max);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function bridge(address caller, uint256 msgValue, IFastBridge.BridgeParams memory params) public virtual {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.bridge{value: msgValue}(params);
    }

    function bridge(
        address caller,
        uint256 msgValue,
        IFastBridge.BridgeParams memory params,
        IFastBridgeV2.BridgeParamsV2 memory paramsV2
    )
        public
    {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.bridge{value: msgValue}(params, paramsV2);
    }

    function prove(address caller, bytes32 transactionId, bytes32 destTxHash, address relayer) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.prove(transactionId, destTxHash, relayer);
    }

    function prove(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx, bytes32 destTxHash) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.prove(abi.encode(bridgeTx), destTxHash);
    }

    function claim(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.claim(abi.encode(bridgeTx));
    }

    function claim(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx, address to) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.claim(abi.encode(bridgeTx), to);
    }

    function dispute(address caller, bytes32 txId) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.dispute(txId);
    }

    function refund(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.refund(abi.encode(bridgeTx));
    }

    function assertEq(FastBridgeV2.BridgeStatus a, FastBridgeV2.BridgeStatus b) public pure {
        assertEq(uint8(a), uint8(b));
    }
}
