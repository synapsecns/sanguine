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

        // See initExistingTxs for why these start from 1, not 0
        bridgedTokenTx.nonce = 1;
        bridgedEthTx.nonce = 2;
        provenTokenTx.nonce = 3;
        provenEthTx.nonce = 4;
        // Next nonce for userA tx would be 5 (either token or eth)
        tokenTx.nonce = 5;
        ethTx.nonce = 5;
    }

    function mintTokens() public virtual override {
        super.mintTokens();
        srcToken.mint(relayerA, INITIAL_RELAYER_BALANCE);
        srcToken.mint(relayerB, INITIAL_RELAYER_BALANCE);
        deal(relayerA, INITIAL_RELAYER_BALANCE);
        deal(relayerB, INITIAL_RELAYER_BALANCE);
    }

    function initExistingTxs() public {
        // Set userA nonce to 1 so that the first bridge tx doesn't have inflated gas costs due to
        // the storage write from the zero initial value
        cheatSenderNonce(userA, 1);
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
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

    function checkAfterBridgeToken() public view {
        assertEq(fastBridge.bridgeStatuses(getTxId(tokenTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken - tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken + tokenParams.originAmount);
    }

    function test_bridge_token() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
        checkAfterBridgeToken();
    }

    function test_bridge_token_withExclusivity() public {
        setTokenTestExclusivityParams(relayerA, EXCLUSIVITY_PERIOD);
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
        checkAfterBridgeToken();
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

    function test_cancelPermissioned_token() public {
        bytes32 txId = getTxId(bridgedTokenTx);
        skipTimeAtLeast({time: DEADLINE});
        cancel({caller: canceler, bridgeTx: bridgedTokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken - tokenParams.originAmount);
    }

    function test_cancelPermissionless_token() public {
        bytes32 txId = getTxId(bridgedTokenTx);
        skipTimeAtLeast({time: DEADLINE + PERMISSIONLESS_CANCEL_DELAY});
        cancel({caller: userB, bridgeTx: bridgedTokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(srcToken.balanceOf(userA), initialUserBalanceToken + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), initialFastBridgeBalanceToken - tokenParams.originAmount);
    }

    // ════════════════════════════════════════════════════ ETH ════════════════════════════════════════════════════════

    function checkAfterBridgeEth() public view {
        assertEq(fastBridge.bridgeStatuses(getTxId(ethTx)), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(userA.balance, initialUserBalanceEth - ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth + ethParams.originAmount);
    }

    function test_bridge_eth() public {
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
        checkAfterBridgeEth();
    }

    function test_bridge_eth_withExclusivity() public {
        setEthTestExclusivityParams(relayerA, EXCLUSIVITY_PERIOD);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
        checkAfterBridgeEth();
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

    function test_cancelPermissioned_eth() public {
        bytes32 txId = getTxId(bridgedEthTx);
        skipTimeAtLeast({time: DEADLINE});
        cancel({caller: canceler, bridgeTx: bridgedEthTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(userA.balance, initialUserBalanceEth + ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth - ethParams.originAmount);
    }

    function test_cancelPermissionless_eth() public {
        bytes32 txId = getTxId(bridgedEthTx);
        skipTimeAtLeast({time: DEADLINE + PERMISSIONLESS_CANCEL_DELAY});
        cancel({caller: userB, bridgeTx: bridgedEthTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(userA.balance, initialUserBalanceEth + ethParams.originAmount);
        assertEq(address(fastBridge).balance, initialFastBridgeBalanceEth - ethParams.originAmount);
    }
}
