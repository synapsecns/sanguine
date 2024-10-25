// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IFastBridge} from "../../contracts/interfaces/IFastBridge.sol";
import {IFastBridgeV2} from "../../contracts/interfaces/IFastBridgeV2.sol";
import {IMulticallTarget} from "../../contracts/interfaces/IMulticallTarget.sol";
import {DisputePeriodNotPassed} from "../../contracts/libs/Errors.sol";

import {MockERC20} from "../MockERC20.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, max-states-count, ordering
abstract contract MulticallTargetIntegrationTest is Test {
    address internal constant ETH_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    uint256 internal constant DEADLINE_PERIOD = 1 days;
    uint256 internal constant SKIP_PERIOD = 1 hours;

    uint32 internal constant LOCAL_CHAIN_ID = 1337;
    uint32 internal constant REMOTE_CHAIN_ID = 7331;

    address internal fastBridge;

    address internal user = makeAddr("User");
    address internal userRemote = makeAddr("User Remote");
    address internal relayer = makeAddr("Relayer");
    address internal claimTo = makeAddr("Claim To");

    MockERC20 internal token;
    address internal remoteToken = makeAddr("Remote Token");

    IFastBridge.BridgeTransaction internal bridgedTokenTx;
    IFastBridge.BridgeTransaction internal provenEthTx;
    IFastBridge.BridgeTransaction internal remoteTokenTx;

    bytes internal encodedBridgedTokenTx;
    bytes internal encodedProvenEthTx;
    bytes internal encodedRemoteTokenTx;

    bytes32 internal bridgedTokenTxId;
    bytes32 internal provenEthTxId;
    bytes32 internal remoteTokenTxId;

    function setUp() public {
        vm.chainId(LOCAL_CHAIN_ID);
        fastBridge = deployAndConfigureFastBridge();
        token = new MockERC20("Token", 18);
        dealTokens(user, 1);
        dealTokens(relayer, 1);
        createFixtures();
        bridge(bridgedTokenTx, user);
        bridge(provenEthTx, user);
        prove(encodedProvenEthTx);
        skip(SKIP_PERIOD);
        // Sanity checks
        checkStatus(bridgedTokenTxId, IFastBridgeV2.BridgeStatus.REQUESTED);
        checkStatus(provenEthTxId, IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        checkStatus(remoteTokenTxId, IFastBridgeV2.BridgeStatus.NULL);
        assertEq(token.balanceOf(user), 0);
        assertEq(user.balance, 0 ether);
        assertEq(token.balanceOf(relayer), 1 ether);
        assertEq(relayer.balance, 1 ether);
    }

    function deployAndConfigureFastBridge() public virtual returns (address);

    function getEncodedBridgeTx(IFastBridge.BridgeTransaction memory bridgeTx)
        public
        view
        virtual
        returns (bytes memory);

    function dealTokens(address to, uint8 amountUnits) public {
        uint256 amountWei = amountUnits * 1 ether;
        token.mint(to, amountWei);
        deal(to, amountWei);
        vm.prank(to);
        token.approve(address(fastBridge), type(uint256).max);
    }

    function createFixtures() public {
        bridgedTokenTx = IFastBridge.BridgeTransaction({
            originChainId: LOCAL_CHAIN_ID,
            destChainId: REMOTE_CHAIN_ID,
            originSender: user,
            destRecipient: userRemote,
            originToken: address(token),
            destToken: remoteToken,
            originAmount: 1 ether,
            destAmount: 0.98 ether,
            originFeeAmount: 0,
            sendChainGas: false,
            deadline: block.timestamp + DEADLINE_PERIOD,
            nonce: 0
        });
        provenEthTx = IFastBridge.BridgeTransaction({
            originChainId: LOCAL_CHAIN_ID,
            destChainId: REMOTE_CHAIN_ID,
            originSender: user,
            destRecipient: userRemote,
            originToken: ETH_ADDRESS,
            destToken: ETH_ADDRESS,
            originAmount: 1 ether,
            destAmount: 0.99 ether,
            originFeeAmount: 0,
            sendChainGas: false,
            deadline: block.timestamp + DEADLINE_PERIOD,
            nonce: 1
        });
        remoteTokenTx = IFastBridge.BridgeTransaction({
            originChainId: REMOTE_CHAIN_ID,
            destChainId: LOCAL_CHAIN_ID,
            originSender: userRemote,
            destRecipient: user,
            originToken: remoteToken,
            destToken: address(token),
            originAmount: 1.01 ether,
            destAmount: 1 ether,
            originFeeAmount: 0,
            sendChainGas: false,
            deadline: block.timestamp + SKIP_PERIOD,
            nonce: 420
        });
        encodedBridgedTokenTx = getEncodedBridgeTx(bridgedTokenTx);
        encodedProvenEthTx = getEncodedBridgeTx(provenEthTx);
        encodedRemoteTokenTx = getEncodedBridgeTx(remoteTokenTx);
        bridgedTokenTxId = keccak256(encodedBridgedTokenTx);
        provenEthTxId = keccak256(encodedProvenEthTx);
        remoteTokenTxId = keccak256(encodedRemoteTokenTx);
    }

    struct TestBridgeTransactionWithMetadata {
        IFastBridge.BridgeTransaction transaction;
        bytes encodedData;
        bytes32 txId;
    }

    function createMany(uint8 countOfTxnsToCreate)
        public
        returns (
            TestBridgeTransactionWithMetadata[] memory toRelay,
            TestBridgeTransactionWithMetadata[] memory toBridgeProveClaimData
        )
    {
        toRelay = new TestBridgeTransactionWithMetadata[](countOfTxnsToCreate);
        toBridgeProveClaimData = new TestBridgeTransactionWithMetadata[](countOfTxnsToCreate);

        // fund user & relayer
        dealTokens(user, countOfTxnsToCreate);
        dealTokens(relayer, countOfTxnsToCreate);

        for (uint8 i = 0; i < countOfTxnsToCreate; i++) {
            // prepare relays for execution later
            IFastBridge.BridgeTransaction memory toRelayTx = IFastBridge.BridgeTransaction({
                originChainId: REMOTE_CHAIN_ID,
                destChainId: LOCAL_CHAIN_ID,
                originSender: userRemote,
                destRecipient: user,
                originToken: remoteToken,
                destToken: address(token),
                originAmount: 1.01 ether,
                destAmount: 1 ether,
                originFeeAmount: 0,
                sendChainGas: false,
                deadline: block.timestamp + SKIP_PERIOD,
                nonce: remoteTokenTx.nonce + i
            });

            bytes memory encoded = getEncodedBridgeTx(toRelayTx);
            bytes32 txId = keccak256(encoded);

            toRelay[i] = TestBridgeTransactionWithMetadata({transaction: toRelayTx, encodedData: encoded, txId: txId});

            // prepare *and* execute bridges, to be subsequently proven & claimed
            IFastBridge.BridgeTransaction memory toBridgeProveClaimTx = IFastBridge.BridgeTransaction({
                originChainId: LOCAL_CHAIN_ID,
                destChainId: REMOTE_CHAIN_ID,
                originSender: user,
                destRecipient: userRemote,
                originToken: address(token),
                destToken: remoteToken,
                originAmount: 1 ether,
                destAmount: 0.98 ether,
                originFeeAmount: 0,
                sendChainGas: false,
                deadline: block.timestamp + DEADLINE_PERIOD,
                nonce: 2 + i
            });

            encoded = getEncodedBridgeTx(toBridgeProveClaimTx);
            txId = keccak256(encoded);

            toBridgeProveClaimData[i] =
            // solhint-disable-next-line max-line-length
             TestBridgeTransactionWithMetadata({transaction: toBridgeProveClaimTx, encodedData: encoded, txId: txId});
        }
    }

    function bridge(IFastBridge.BridgeTransaction memory bridgeTx, address bridgeUser) public {
        uint256 msgValue = bridgeTx.originToken == ETH_ADDRESS ? bridgeTx.originAmount : 0;
        vm.prank(bridgeUser);
        IFastBridge(fastBridge).bridge{value: msgValue}(
            IFastBridge.BridgeParams({
                dstChainId: bridgeTx.destChainId,
                sender: bridgeTx.originSender,
                to: bridgeTx.destRecipient,
                originToken: bridgeTx.originToken,
                destToken: bridgeTx.destToken,
                originAmount: bridgeTx.originAmount,
                destAmount: bridgeTx.destAmount,
                sendChainGas: bridgeTx.sendChainGas,
                deadline: bridgeTx.deadline
            })
        );
    }

    function prove(bytes memory encodedBridgeTx) public {
        vm.prank(relayer);
        IFastBridge(fastBridge).prove(encodedBridgeTx, hex"01");
    }

    function getData() public view returns (bytes[] memory data) {
        data = new bytes[](3);
        data[0] = abi.encodeCall(IFastBridge.prove, (encodedBridgedTokenTx, hex"02"));
        data[1] = abi.encodeCall(IFastBridge.claim, (encodedProvenEthTx, claimTo));
        data[2] = abi.encodeCall(IFastBridge.relay, (encodedRemoteTokenTx));
    }

    enum TestAction {
        bridge,
        prove,
        claim,
        relay
    }

    function getDataMany(
        TestBridgeTransactionWithMetadata[] memory testBridgeTxns,
        TestAction action
    )
        public
        view
        returns (bytes[] memory data)
    {
        data = new bytes[](testBridgeTxns.length);

        for (uint8 i = 0; i < testBridgeTxns.length; i++) {
            if (action == TestAction.bridge) {
                data[i] = abi.encodeCall(
                    IFastBridge.bridge,
                    (
                        IFastBridge.BridgeParams({
                            dstChainId: testBridgeTxns[i].transaction.destChainId,
                            sender: testBridgeTxns[i].transaction.originSender,
                            to: testBridgeTxns[i].transaction.destRecipient,
                            originToken: testBridgeTxns[i].transaction.originToken,
                            destToken: testBridgeTxns[i].transaction.destToken,
                            originAmount: testBridgeTxns[i].transaction.originAmount,
                            destAmount: testBridgeTxns[i].transaction.destAmount,
                            sendChainGas: testBridgeTxns[i].transaction.sendChainGas,
                            deadline: testBridgeTxns[i].transaction.deadline
                        })
                    )
                );
            }
            if (action == TestAction.prove) {
                data[i] = abi.encodeCall(IFastBridge.prove, (testBridgeTxns[i].encodedData, hex"02"));
            }
            if (action == TestAction.claim) {
                data[i] = abi.encodeCall(IFastBridge.claim, (testBridgeTxns[i].encodedData, claimTo));
            }
            if (action == TestAction.relay) {
                data[i] = abi.encodeCall(IFastBridge.relay, (testBridgeTxns[i].encodedData));
            }
        }
    }

    function checkStatus(bytes32 txId, IFastBridgeV2.BridgeStatus expected) public view {
        IFastBridgeV2.BridgeStatus status = IFastBridgeV2(fastBridge).bridgeStatuses(txId);
        assertEq(uint8(status), uint8(expected));
    }

    function test_sequentialExecution() public {
        vm.startPrank(relayer);
        IFastBridge(fastBridge).prove(encodedBridgedTokenTx, hex"02");
        IFastBridge(fastBridge).claim(encodedProvenEthTx, claimTo);
        IFastBridge(fastBridge).relay(encodedRemoteTokenTx);
        vm.stopPrank();
        checkHappyPath();
    }

    enum TestExecutionMode {
        Sequential_NonMulticall,
        Multicall
    }

    function test_manyActions_sequentialNonMulticall() public {
        // send X contiguous bridges, proofs, and claims -- and X non-contiguous relays to the same bridger
        manyActions_flow(5, TestExecutionMode.Sequential_NonMulticall);
    }

    function test_manyActions_multicall() public {
        // send X contiguous bridges, proofs, and claims -- and X non-contiguous relays to the same bridger
        manyActions_flow(5, TestExecutionMode.Multicall);
    }

    // will either execute a single batched multicall, or many sequential txns
    function manyActions_execute(
        TestBridgeTransactionWithMetadata[] memory transactions,
        TestAction action,
        TestExecutionMode mode,
        address pranker
    )
        internal
    {
        bytes[] memory data = getDataMany(transactions, action);
        if (mode == TestExecutionMode.Multicall) {
            vm.prank(pranker);
            IMulticallTarget(fastBridge).multicallNoResults({data: data, ignoreReverts: false});
        } else {
            executeSequentialTransactions(transactions, action, pranker);
        }
    }

    function executeSequentialTransactions(
        TestBridgeTransactionWithMetadata[] memory transactions,
        TestAction action,
        address pranker
    )
        internal
    {
        for (uint8 i = 0; i < transactions.length; i++) {
            if (action == TestAction.bridge) {
                // bridge already pranks as user, no need to prank
                bridge(transactions[i].transaction, user);
            } else if (action == TestAction.prove) {
                vm.prank(pranker);
                IFastBridge(fastBridge).prove(transactions[i].encodedData, hex"02");
            } else if (action == TestAction.claim) {
                vm.prank(pranker);
                IFastBridge(fastBridge).claim(transactions[i].encodedData, claimTo);
            } else if (action == TestAction.relay) {
                vm.prank(pranker);
                IFastBridge(fastBridge).relay(transactions[i].encodedData);
            }
            // Check status after each action
            manyActions_checkStatusAfter(transactions[i].txId, action);
        }
    }

    function manyActions_checkStatusAfter(bytes32 txId, TestAction action) internal view {
        if (action == TestAction.bridge) {
            checkStatus(txId, IFastBridgeV2.BridgeStatus.REQUESTED);
        } else if (action == TestAction.prove) {
            checkStatus(txId, IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        } else if (action == TestAction.claim) {
            checkStatus(txId, IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        }
        // No status check needed for relay action
    }

    // Shared flow & asserts regardless of whether executing manyAction test via MC or a sequence of non-MC txns
    function manyActions_flow(uint8 countOfTxnsToCreate, TestExecutionMode testExecutionMode) internal {
        uint256 relayerOriginalBalWei = token.balanceOf(relayer);
        uint256 userOriginalBalWei = token.balanceOf(user);

        (
            TestBridgeTransactionWithMetadata[] memory toRelay,
            TestBridgeTransactionWithMetadata[] memory toBridgeProveClaim
        ) = createMany(countOfTxnsToCreate);

        manyActions_execute(toBridgeProveClaim, TestAction.bridge, testExecutionMode, user);
        assertEq(token.balanceOf(user), 0 ether);

        manyActions_execute(toBridgeProveClaim, TestAction.prove, testExecutionMode, relayer);

        skip(SKIP_PERIOD);

        manyActions_execute(toBridgeProveClaim, TestAction.claim, testExecutionMode, relayer);

        assertEq(token.balanceOf(relayer), relayerOriginalBalWei + (countOfTxnsToCreate * 1 ether));

        manyActions_execute(toRelay, TestAction.relay, testExecutionMode, relayer);

        assertEq(token.balanceOf(relayer), relayerOriginalBalWei);
        assertEq(token.balanceOf(user), userOriginalBalWei + (countOfTxnsToCreate * 1 ether));
    }

    // ════════════════════════════════════════════════ NO RESULTS ═════════════════════════════════════════════════════

    function checkHappyPath() public view {
        // Check statuses
        checkStatus(bridgedTokenTxId, IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        checkStatus(provenEthTxId, IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        assertTrue(IFastBridgeV2(fastBridge).bridgeRelays(remoteTokenTxId));
        // Check balances
        assertEq(token.balanceOf(user), 1 ether);
        assertEq(user.balance, 0 ether);
        assertEq(token.balanceOf(relayer), 0);
        assertEq(relayer.balance, 1 ether);
        assertEq(claimTo.balance, 1 ether);
    }

    function checkHappyPathNoClaim() public view {
        // Check statuses
        checkStatus(bridgedTokenTxId, IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        checkStatus(provenEthTxId, IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        assertTrue(IFastBridgeV2(fastBridge).bridgeRelays(remoteTokenTxId));
        // Check balances
        assertEq(token.balanceOf(user), 1 ether);
        assertEq(user.balance, 0 ether);
        assertEq(token.balanceOf(relayer), 0);
        assertEq(relayer.balance, 1 ether);
        assertEq(claimTo.balance, 0);
    }

    function test_multicallNoResults_ignoreReverts_success() public {
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget(fastBridge).multicallNoResults({data: data, ignoreReverts: true});
        checkHappyPath();
    }

    function test_multicallNoResults_ignoreReverts_withFailedClaim() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget(fastBridge).multicallNoResults({data: data, ignoreReverts: true});
        checkHappyPathNoClaim();
    }

    function test_multicallNoResults_dontIgnoreReverts_success() public {
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget(fastBridge).multicallNoResults({data: data, ignoreReverts: false});
        checkHappyPath();
    }

    function test_multicallNoResults_dontIgnoreReverts_revert() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.expectRevert(DisputePeriodNotPassed.selector);
        vm.prank(relayer);
        IMulticallTarget(fastBridge).multicallNoResults({data: data, ignoreReverts: false});
    }

    // ═══════════════════════════════════════════════ WITH RESULTS ════════════════════════════════════════════════════

    function assertEq(IMulticallTarget.Result memory result, IMulticallTarget.Result memory expected) public pure {
        assertEq(result.success, expected.success);
        assertEq(result.returnData, expected.returnData);
    }

    function checkHappyPathResults(IMulticallTarget.Result[] memory results) public pure {
        assertEq(results.length, 3);
        assertEq(results[0], IMulticallTarget.Result({success: true, returnData: ""}));
        assertEq(results[1], IMulticallTarget.Result({success: true, returnData: ""}));
        assertEq(results[2], IMulticallTarget.Result({success: true, returnData: ""}));
    }

    function checkHappyPathNoClaimResults(IMulticallTarget.Result[] memory results) public pure {
        assertEq(results.length, 3);
        assertEq(results[0], IMulticallTarget.Result({success: true, returnData: ""}));
        assertEq(
            results[1],
            IMulticallTarget.Result({
                success: false,
                returnData: abi.encodeWithSelector(DisputePeriodNotPassed.selector)
            })
        );
        assertEq(results[2], IMulticallTarget.Result({success: true, returnData: ""}));
    }

    function test_multicallWithResults_ignoreReverts_success() public {
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget.Result[] memory results =
            IMulticallTarget(fastBridge).multicallWithResults({data: data, ignoreReverts: true});
        checkHappyPath();
        checkHappyPathResults(results);
    }

    function test_multicallWithResults_ignoreReverts_withFailedClaim() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget.Result[] memory results =
            IMulticallTarget(fastBridge).multicallWithResults({data: data, ignoreReverts: true});
        checkHappyPathNoClaim();
        checkHappyPathNoClaimResults(results);
    }

    function test_multicallWithResults_dontIgnoreReverts_success() public {
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget.Result[] memory results =
            IMulticallTarget(fastBridge).multicallWithResults({data: data, ignoreReverts: false});
        checkHappyPath();
        checkHappyPathResults(results);
    }

    function test_multicallWithResults_dontIgnoreReverts_revert() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.expectRevert(DisputePeriodNotPassed.selector);
        vm.prank(relayer);
        IMulticallTarget(fastBridge).multicallWithResults({data: data, ignoreReverts: false});
    }
}
