// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BridgeTransactionV2Lib} from "../contracts/libs/BridgeTransactionV2.sol";

import {FastBridgeV2, FastBridgeV2Test, IFastBridge, IFastBridgeV2} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
abstract contract FastBridgeV2SrcBaseTest is FastBridgeV2Test {
    uint256 public constant MIN_DEADLINE = 30 minutes;
    uint256 public constant CLAIM_DELAY = 30 minutes;
    // Use a value different from the default to ensure it's being set correctly.
    uint256 public constant PERMISSIONLESS_CANCEL_DELAY = 13.37 hours;

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
        fastBridge.grantRole(fastBridge.PROVER_ROLE(), relayerA);
        fastBridge.grantRole(fastBridge.PROVER_ROLE(), relayerB);
        fastBridge.grantRole(fastBridge.GUARD_ROLE(), guard);
        fastBridge.grantRole(fastBridge.CANCELER_ROLE(), canceler);

        fastBridge.grantRole(fastBridge.GOVERNOR_ROLE(), address(this));
        fastBridge.setCancelDelay(PERMISSIONLESS_CANCEL_DELAY);
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
        fastBridge.bridgeV2{value: msgValue}(params, paramsV2);
    }

    function prove(address caller, bytes32 transactionId, bytes32 destTxHash, address relayer) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.proveV2(transactionId, destTxHash, relayer);
    }

    function prove(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx, bytes32 destTxHash) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.prove(BridgeTransactionV2Lib.encodeV2(bridgeTx), destTxHash);
    }

    function claim(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.claimV2(BridgeTransactionV2Lib.encodeV2(bridgeTx));
    }

    function claim(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx, address to) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.claim(BridgeTransactionV2Lib.encodeV2(bridgeTx), to);
    }

    function dispute(address caller, bytes32 txId) public {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.dispute(txId);
    }

    function cancel(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public virtual {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.cancelV2(BridgeTransactionV2Lib.encodeV2(bridgeTx));
    }

    function test_nonce() public view {
        uint256 result = fastBridge.nonce();
        // deprecated. should always return zero in FbV2.
        assertEq(result, 0);
    }

    function assertEq(FastBridgeV2.BridgeStatus a, FastBridgeV2.BridgeStatus b) public pure {
        assertEq(uint8(a), uint8(b));
    }
}
