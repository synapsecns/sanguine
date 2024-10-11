// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2SrcBaseTest, IFastBridgeV2} from "./FastBridgeV2.Src.Base.t.sol";

// solhint-disable func-name-mixedcase, ordering
/// @notice This test is used to estimate the gas cost of FastBridgeV2 source chain functions.
/// Very little state checks are performed, make sure to do full coverage in different tests.
contract FastBridgeV2GasBenchmarkSrcTest is FastBridgeV2SrcBaseTest {
    uint256 public constant BLOCK_TIME = 12 seconds;
    uint256 public constant INITIAL_RELAYER_BALANCE = 100 ether;
    uint256 public constant EXCLUSIVITY_PERIOD = 60 seconds;

    IFastBridgeV2.BridgeTransactionV2 internal bridgedTokenTx;
    IFastBridgeV2.BridgeTransactionV2 internal bridgedEthTx;

    IFastBridgeV2.BridgeTransactionV2 internal provenTokenTx;
    IFastBridgeV2.BridgeTransactionV2 internal provenEthTx;

    uint256 public initialUserBalanceToken;
    uint256 public initialUserBalanceEth;
    uint256 public initialFastBridgeBalanceToken;
    uint256 public initialFastBridgeBalanceEth;

    function setUp() public virtual override {
        super.setUp();
        initExistingTxs();
        initialUserBalanceToken = srcToken.balanceOf(userA);
        initialUserBalanceEth = userA.balance;
        initialFastBridgeBalanceToken = srcToken.balanceOf(address(fastBridge));
        initialFastBridgeBalanceEth = address(fastBridge).balance;
    }

    function createFixtures() public virtual override {
        super.createFixtures();
        bridgedTokenTx = tokenTx;
        provenTokenTx = tokenTx;
        bridgedEthTx = ethTx;
        provenEthTx = ethTx;

        bridgedTokenTx.nonce = 0;
        bridgedEthTx.nonce = 1;
        provenTokenTx.nonce = 2;
        provenEthTx.nonce = 3;
        // Next nonce for userA tx would be 4 (either token or eth)
        tokenTx.nonce = 4;
        ethTx.nonce = 4;
    }

    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        bridgedTokenTx.exclusivityEndTime = block.timestamp;
        provenTokenTx.exclusivityEndTime = block.timestamp;
        bridgedEthTx.exclusivityEndTime = block.timestamp;
        provenEthTx.exclusivityEndTime = block.timestamp;
        // Actual tx will be submitted one block later
        tokenTx.exclusivityEndTime = block.timestamp + BLOCK_TIME;
        ethTx.exclusivityEndTime = block.timestamp + BLOCK_TIME;
    }

    function mintTokens() public virtual override {
        super.mintTokens();
        srcToken.mint(relayerA, INITIAL_RELAYER_BALANCE);
        srcToken.mint(relayerB, INITIAL_RELAYER_BALANCE);
        deal(relayerA, INITIAL_RELAYER_BALANCE);
        deal(relayerB, INITIAL_RELAYER_BALANCE);
    }

    function initExistingTxs() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        skipBlocksExactly(1);
        prove({caller: relayerA, bridgeTx: provenTokenTx, destTxHash: hex"01"});
        prove({caller: relayerB, transactionId: getTxId(provenEthTx), destTxHash: hex"02", relayer: relayerA});
        // Status checks
        assertEq(fastBridge.bridgeStatuses(getTxId(bridgedTokenTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(fastBridge.bridgeStatuses(getTxId(bridgedEthTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(fastBridge.bridgeStatuses(getTxId(provenTokenTx)), IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        assertEq(fastBridge.bridgeStatuses(getTxId(provenEthTx)), IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        assertEq(fastBridge.bridgeStatuses(getTxId(tokenTx)), IFastBridgeV2.BridgeStatus.NULL);
        assertEq(fastBridge.bridgeStatuses(getTxId(ethTx)), IFastBridgeV2.BridgeStatus.NULL);
    }

    function skipBlocksExactly(uint256 blocks) public {
        vm.roll(block.number + blocks);
        vm.warp(block.timestamp + blocks * BLOCK_TIME);
    }

    function skipTimeAtLeast(uint256 time) public {
        uint256 blocksToSkip = time / BLOCK_TIME;
        if (blocksToSkip * BLOCK_TIME < time) blocksToSkip++;
        skipBlocksExactly(blocksToSkip);
    }

    // ═══════════════════════════════════════════════════ TOKEN ═══════════════════════════════════════════════════════

    function test_bridge_token() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.bridgeStatuses(getTxId(tokenTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken - tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken + tokenParams.originAmount);
    }

    function test_bridge_token_withArbitraryCall() public {
        setTokenTestCallParams(abi.encode(userB));
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
        assertEq(fastBridge.bridgeStatuses(getTxId(tokenTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken - tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken + tokenParams.originAmount);
    }

    function test_bridge_token_withExclusivity() public {
        setTokenTestExclusivityParams(relayerA, EXCLUSIVITY_PERIOD);
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
        assertEq(fastBridge.bridgeStatuses(getTxId(tokenTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken - tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken + tokenParams.originAmount);
    }

    function test_prove_token() public {
        bytes32 txId = getTxId(bridgedTokenTx);
        prove({caller: relayerA, bridgeTx: bridgedTokenTx, destTxHash: hex"03"});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken);
    }

    function test_proveWithAddress_token() public {
        bytes32 txId = getTxId(bridgedTokenTx);
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"03", relayer: relayerA});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
    }

    function test_claim_token() public {
        bytes32 txId = getTxId(provenTokenTx);
        skipTimeAtLeast({time: CLAIM_DELAY + 1});
        assertTrue(fastBridge.canClaim(txId, relayerA));
        claim({caller: relayerA, bridgeTx: provenTokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        assertEq(srcToken.balanceOf(relayerA), INITIAL_RELAYER_BALANCE + tokenTx.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken - tokenTx.originAmount);
    }

    function test_claimWithAddress_token() public {
        bytes32 txId = getTxId(provenTokenTx);
        skipTimeAtLeast({time: CLAIM_DELAY + 1});
        assertTrue(fastBridge.canClaim(txId, relayerA));
        claim({caller: relayerA, bridgeTx: provenTokenTx, to: relayerB});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        assertEq(srcToken.balanceOf(relayerB), INITIAL_RELAYER_BALANCE + tokenTx.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken - tokenTx.originAmount);
    }

    function test_dispute_token() public {
        bytes32 txId = getTxId(provenTokenTx);
        dispute({caller: guard, txId: txId});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken);
    }

    function test_refundPermissioned_token() public {
        bytes32 txId = getTxId(bridgedTokenTx);
        skipTimeAtLeast({time: DEADLINE});
        refund({caller: refunder, bridgeTx: bridgedTokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken - tokenParams.originAmount);
    }

    function test_refundPermissionless_token() public {
        bytes32 txId = getTxId(bridgedTokenTx);
        skipTimeAtLeast({time: DEADLINE + PERMISSIONLESS_REFUND_DELAY});
        refund({caller: userB, bridgeTx: bridgedTokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken - tokenParams.originAmount);
    }

    // ════════════════════════════════════════════════════ ETH ════════════════════════════════════════════════════════

    function test_bridge_eth() public {
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.bridgeStatuses(getTxId(ethTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(userA.balance, initialUserBalanceEth - ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth + ethParams.originAmount);
    }

    function test_bridge_eth_withArbitraryCall() public {
        setEthTestCallParams(abi.encode(userB));
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
        assertEq(fastBridge.bridgeStatuses(getTxId(ethTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(userA.balance, initialUserBalanceEth - ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth + ethParams.originAmount);
    }

    function test_bridge_eth_withExclusivity() public {
        setEthTestExclusivityParams(relayerA, EXCLUSIVITY_PERIOD);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
        assertEq(fastBridge.bridgeStatuses(getTxId(ethTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(userA.balance, initialUserBalanceEth - ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth + ethParams.originAmount);
    }

    function test_prove_eth() public {
        bytes32 txId = getTxId(bridgedEthTx);
        prove({caller: relayerA, bridgeTx: bridgedEthTx, destTxHash: hex"03"});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth);
    }

    function test_proveWithAddress_eth() public {
        bytes32 txId = getTxId(bridgedEthTx);
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"03", relayer: relayerA});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth);
    }

    function test_claim_eth() public {
        bytes32 txId = getTxId(provenEthTx);
        skipTimeAtLeast({time: CLAIM_DELAY + 1});
        assertTrue(fastBridge.canClaim(txId, relayerA));
        claim({caller: relayerA, bridgeTx: provenEthTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        assertEq(relayerA.balance, INITIAL_RELAYER_BALANCE + ethTx.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth - ethTx.originAmount);
    }

    function test_claimWithAddress_eth() public {
        bytes32 txId = getTxId(provenEthTx);
        skipTimeAtLeast({time: CLAIM_DELAY + 1});
        assertTrue(fastBridge.canClaim(txId, relayerA));
        claim({caller: relayerA, bridgeTx: provenEthTx, to: relayerB});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        assertEq(relayerB.balance, INITIAL_RELAYER_BALANCE + ethTx.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth - ethTx.originAmount);
    }

    function test_dispute_eth() public {
        bytes32 txId = getTxId(provenEthTx);
        dispute({caller: guard, txId: txId});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth);
    }

    function test_refundPermissioned_eth() public {
        bytes32 txId = getTxId(bridgedEthTx);
        skipTimeAtLeast({time: DEADLINE});
        refund({caller: refunder, bridgeTx: bridgedEthTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(userA.balance, initialUserBalanceEth + ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth - ethParams.originAmount);
    }

    function test_refundPermissionless_eth() public {
        bytes32 txId = getTxId(bridgedEthTx);
        skipTimeAtLeast({time: DEADLINE + PERMISSIONLESS_REFUND_DELAY});
        refund({caller: userB, bridgeTx: bridgedEthTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(userA.balance, initialUserBalanceEth + ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth - ethParams.originAmount);
    }
}
