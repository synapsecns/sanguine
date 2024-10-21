// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {FastBridge, IFastBridge} from "../contracts/FastBridge.sol";
import {IMulticallTarget} from "../contracts/interfaces/IMulticallTarget.sol";
import {DisputePeriodNotPassed} from "../contracts/libs/Errors.sol";

import {MockERC20} from "./MockERC20.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeMulticallTargetTest is Test {
    address internal constant ETH_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    uint256 internal constant DEADLINE_PERIOD = 1 days;
    uint256 internal constant SKIP_PERIOD = 1 hours;

    uint32 internal constant LOCAL_CHAIN_ID = 1337;
    uint32 internal constant REMOTE_CHAIN_ID = 7331;

    FastBridge internal fastBridge;

    address internal user = makeAddr("User");
    address internal userRemote = makeAddr("User Remote");
    address internal relayer = makeAddr("Relayer");
    address internal claimTo = makeAddr("Claim To");

    MockERC20 internal token;
    address internal remoteToken = makeAddr("Remote Token");

    IFastBridge.BridgeTransaction internal bridgedTokenTx;
    IFastBridge.BridgeTransaction internal provenEthTx;
    IFastBridge.BridgeTransaction internal remoteTokenTx;

    bytes32 internal bridgedTokenTxId;
    bytes32 internal provenEthTxId;
    bytes32 internal remoteTokenTxId;

    function setUp() public {
        vm.chainId(LOCAL_CHAIN_ID);
        fastBridge = new FastBridge(address(this));
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayer);
        token = new MockERC20("Token", 18);
        dealTokens(user);
        dealTokens(relayer);
        createFixtures();
        bridge(bridgedTokenTx);
        bridge(provenEthTx);
        prove(provenEthTx);
        skip(SKIP_PERIOD);
        // Sanity checks
        checkStatus(bridgedTokenTxId, FastBridge.BridgeStatus.REQUESTED);
        checkStatus(provenEthTxId, FastBridge.BridgeStatus.RELAYER_PROVED);
        checkStatus(remoteTokenTxId, FastBridge.BridgeStatus.NULL);
        assertEq(token.balanceOf(user), 0);
        assertEq(user.balance, 0 ether);
        assertEq(token.balanceOf(relayer), 1 ether);
        assertEq(relayer.balance, 1 ether);
    }

    function dealTokens(address to) public {
        token.mint(to, 1 ether);
        deal(to, 1 ether);
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
            originSender: userRemote,
            destRecipient: user,
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

        bridgedTokenTxId = keccak256(abi.encode(bridgedTokenTx));
        provenEthTxId = keccak256(abi.encode(provenEthTx));
        remoteTokenTxId = keccak256(abi.encode(remoteTokenTx));
    }

    function bridge(IFastBridge.BridgeTransaction memory bridgeTx) public {
        uint256 msgValue = bridgeTx.originToken == ETH_ADDRESS ? bridgeTx.originAmount : 0;
        vm.prank(user);
        fastBridge.bridge{value: msgValue}(
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

    function prove(IFastBridge.BridgeTransaction memory bridgeTx) public {
        bytes memory request = abi.encode(bridgeTx);
        vm.prank(relayer);
        fastBridge.prove(request, hex"01");
    }

    function getData() public view returns (bytes[] memory data) {
        data = new bytes[](3);
        data[0] = abi.encodeCall(fastBridge.prove, (abi.encode(bridgedTokenTx), hex"02"));
        data[1] = abi.encodeCall(fastBridge.claim, (abi.encode(provenEthTx), claimTo));
        data[2] = abi.encodeCall(fastBridge.relay, (abi.encode(remoteTokenTx)));
    }

    function checkStatus(bytes32 txId, FastBridge.BridgeStatus expected) public view {
        FastBridge.BridgeStatus status = fastBridge.bridgeStatuses(txId);
        assertEq(uint8(status), uint8(expected));
    }

    function test_sequentialExecution() public {
        vm.startPrank(relayer);
        fastBridge.prove(abi.encode(bridgedTokenTx), hex"02");
        fastBridge.claim(abi.encode(provenEthTx), claimTo);
        fastBridge.relay(abi.encode(remoteTokenTx));
        vm.stopPrank();
        checkHappyPath();
    }

    // ════════════════════════════════════════════════ NO RESULTS ═════════════════════════════════════════════════════

    function checkHappyPath() public view {
        // Check statuses
        checkStatus(bridgedTokenTxId, FastBridge.BridgeStatus.RELAYER_PROVED);
        checkStatus(provenEthTxId, FastBridge.BridgeStatus.RELAYER_CLAIMED);
        assertTrue(fastBridge.bridgeRelays(remoteTokenTxId));
        // Check balances
        assertEq(token.balanceOf(user), 1 ether);
        assertEq(user.balance, 0 ether);
        assertEq(token.balanceOf(relayer), 0);
        assertEq(relayer.balance, 1 ether);
        assertEq(claimTo.balance, 1 ether);
    }

    function checkHappyPathNoClaim() public view {
        // Check statuses
        checkStatus(bridgedTokenTxId, FastBridge.BridgeStatus.RELAYER_PROVED);
        checkStatus(provenEthTxId, FastBridge.BridgeStatus.RELAYER_PROVED);
        assertTrue(fastBridge.bridgeRelays(remoteTokenTxId));
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
        fastBridge.multicallNoResults({data: data, ignoreReverts: true});
        checkHappyPath();
    }

    function test_multicallNoResults_ignoreReverts_withFailedClaim() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.prank(relayer);
        fastBridge.multicallNoResults({data: data, ignoreReverts: true});
        checkHappyPathNoClaim();
    }

    function test_multicallNoResults_dontIgnoreReverts_success() public {
        bytes[] memory data = getData();
        vm.prank(relayer);
        fastBridge.multicallNoResults({data: data, ignoreReverts: false});
        checkHappyPath();
    }

    function test_multicallNoResults_dontIgnoreReverts_revert() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.expectRevert(DisputePeriodNotPassed.selector);
        vm.prank(relayer);
        fastBridge.multicallNoResults({data: data, ignoreReverts: false});
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
        IMulticallTarget.Result[] memory results = fastBridge.multicallWithResults({data: data, ignoreReverts: true});
        checkHappyPath();
        checkHappyPathResults(results);
    }

    function test_multicallWithResults_ignoreReverts_withFailedClaim() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget.Result[] memory results = fastBridge.multicallWithResults({data: data, ignoreReverts: true});
        checkHappyPathNoClaim();
        checkHappyPathNoClaimResults(results);
    }

    function test_multicallWithResults_dontIgnoreReverts_success() public {
        bytes[] memory data = getData();
        vm.prank(relayer);
        IMulticallTarget.Result[] memory results = fastBridge.multicallWithResults({data: data, ignoreReverts: false});
        checkHappyPath();
        checkHappyPathResults(results);
    }

    function test_multicallWithResults_dontIgnoreReverts_revert() public {
        // Rewind time to make claim fail
        rewind(SKIP_PERIOD - 15 minutes);
        bytes[] memory data = getData();
        vm.expectRevert(DisputePeriodNotPassed.selector);
        vm.prank(relayer);
        fastBridge.multicallWithResults({data: data, ignoreReverts: false});
    }
}
