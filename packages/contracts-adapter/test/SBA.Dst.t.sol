// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseHyperCoreComposer} from "../src/SynapseHyperCoreComposer.sol";
import {ICoreWriter} from "../src/interfaces/ICoreWriter.sol";
import {ISynapseBridgeAdapter, SynapseBridgeAdapterTest, SynapseBridgeAdapterV2} from "./SBA.t.sol";

import {SynapseBridgeMock} from "./mocks/SynapseBridgeMock.sol";
import {TestToken} from "./mocks/TestToken.sol";

import {Origin} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";

// solhint-disable func-name-mixedcase, ordering, max-line-length
contract SynapseBridgeAdapterDstTest is SynapseBridgeAdapterTest {
    uint64 internal constant CORE_INDEX = 1337;
    uint64 internal constant CORE_AMOUNT = 12_345_600;

    address internal bridge;
    TestToken internal token;
    SynapseHyperCoreComposer internal composer;

    address internal executor = makeAddr("Executor");
    address internal recipient = makeAddr("Recipient");
    uint256 internal amount = 0.123_456 ether;

    bytes internal bridgeMessage;

    modifier withBridgeSet() {
        adapter.setBridge(bridge);
        _;
    }

    modifier withMintTokenAdded() {
        adapter.addToken(
            address(token),
            ISynapseBridgeAdapter.TokenType.MintBurn,
            toArray(ISynapseBridgeAdapter.RemoteToken(SRC_EID, remoteToken))
        );
        _;
    }

    modifier withWithdrawTokenAdded() {
        adapter.addToken(
            address(token),
            ISynapseBridgeAdapter.TokenType.WithdrawDeposit,
            toArray(ISynapseBridgeAdapter.RemoteToken(SRC_EID, remoteToken))
        );
        _;
    }

    modifier withHyperCoreComposerSet() {
        composer = new SynapseHyperCoreComposer(address(adapter), address(token), CORE_INDEX);
        adapter.setHyperCoreComposer(address(token), address(composer));
        vm.etch(composer.CORE_WRITER(), hex"00");
        _;
    }

    function afterAdapterDeployed() internal virtual override {
        adapter.setPeer(SRC_EID, REMOTE_ADAPTER);

        bridge = address(new SynapseBridgeMock());
        token = new TestToken();

        bridgeMessage = bridgeMessageLib.encodeBridgeMessage(recipient, remoteToken, amount);
    }

    function deployAdapter() internal virtual override returns (SynapseBridgeAdapterV2) {
        return new SynapseBridgeAdapterV2(endpoint, address(this));
    }

    function endpointCallsLzReceive() internal {
        endpointCallsLzReceive(SRC_EID, REMOTE_ADAPTER);
    }

    function endpointCallsLzReceive(uint32 srcEid, bytes32 sender) internal {
        vm.prank({msgSender: endpoint, txOrigin: executor});
        adapter.lzReceive({
            _origin: Origin({srcEid: srcEid, sender: sender, nonce: 1}),
            _guid: MOCK_GUID,
            _message: bridgeMessage,
            _executor: executor,
            _extraData: ""
        });
    }

    function useHyperCoreMessage() internal {
        bridgeMessage = hyperCoreMessageLib.encodeHyperCoreMessage(recipient, remoteToken, amount);
    }

    function mockCoreUser(address user, bool exists) internal {
        vm.mockCall(composer.CORE_USER_EXISTS_PRECOMPILE(), abi.encode(user), abi.encode(exists));
    }

    function mockSpotBalance(uint64 total) internal {
        vm.mockCall(
            composer.SPOT_BALANCE_PRECOMPILE(),
            abi.encode(composer.assetBridge(), CORE_INDEX),
            abi.encode(total, uint64(0), uint64(0))
        );
    }

    // ═══════════════════════════════════════════ TEST: MINT-BURN TOKEN ═══════════════════════════════════════════════

    function test_receive_mintBurn() public withBridgeSet withMintTokenAdded {
        // Expected action: bridge.mint
        vm.expectCall({
            callee: bridge,
            data: abi.encodeCall(SynapseBridgeMock.mint, (recipient, address(token), amount, 0, MOCK_GUID))
        });
        // Expected event
        expectEventTokenReceived(SRC_EID, recipient, address(token), amount, MOCK_GUID);
        endpointCallsLzReceive();
    }

    function test_receive_mintBurn_revert_bridgeNotSet() public withMintTokenAdded {
        expectRevertBridgeNotSet();
        endpointCallsLzReceive();
    }

    function test_receive_mintBurn_revert_tokenNotAdded() public withBridgeSet {
        expectRevertLocalPairNotFound(SRC_EID, remoteToken);
        endpointCallsLzReceive();
    }

    function test_receive_mintBurn_revert_bridgeNotSetTokenNotAdded() public {
        expectRevertBridgeNotSet();
        endpointCallsLzReceive();
    }

    function test_receive_mintBurn_revert_eidUnknown() public withBridgeSet withMintTokenAdded {
        vm.expectRevert();
        endpointCallsLzReceive(UNKNOWN_EID, REMOTE_ADAPTER);
    }

    function test_receive_mintBurn_revert_eidUnknown_withPeerAdded() public withBridgeSet withMintTokenAdded {
        adapter.setPeer(UNKNOWN_EID, REMOTE_ADAPTER);
        expectRevertLocalPairNotFound(UNKNOWN_EID, remoteToken);
        endpointCallsLzReceive(UNKNOWN_EID, REMOTE_ADAPTER);
    }

    function test_receive_mintBurn_revert_senderUnknown() public withBridgeSet withMintTokenAdded {
        vm.expectRevert();
        endpointCallsLzReceive(SRC_EID, keccak256("Unknown"));
    }

    // ═════════════════════════════════════════ TEST: HYPERCORE DELIVERY ═════════════════════════════════════════════

    function test_receiveHyperCore_mintBurn() public withBridgeSet withMintTokenAdded withHyperCoreComposerSet {
        useHyperCoreMessage();
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        mockSpotBalance(type(uint64).max);

        bytes memory payload =
            abi.encodePacked(composer.SPOT_SEND_HEADER(), abi.encode(recipient, CORE_INDEX, CORE_AMOUNT));
        vm.expectCall({
            callee: bridge,
            data: abi.encodeCall(SynapseBridgeMock.mint, (address(composer), address(token), amount, 0, MOCK_GUID))
        });
        vm.expectCall(composer.CORE_WRITER(), abi.encodeCall(ICoreWriter.sendRawAction, (payload)));

        endpointCallsLzReceive();

        assertEq(token.totalSupply(), amount);
        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.balanceOf(composer.assetBridge()), amount);
        assertEq(token.balanceOf(recipient), 0);
    }

    function test_receiveHyperCore_revert_composerNotSet() public withBridgeSet withMintTokenAdded {
        useHyperCoreMessage();
        expectRevertHyperCoreComposerNotSet(address(token));
        endpointCallsLzReceive();
        assertEq(token.totalSupply(), 0);
    }

    function test_receiveHyperCore_revert_capacity_rollsBackMint()
        public
        withBridgeSet
        withMintTokenAdded
        withHyperCoreComposerSet
    {
        useHyperCoreMessage();
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        mockSpotBalance(CORE_AMOUNT - 1);
        vm.expectRevert(
            abi.encodeWithSelector(
                SynapseHyperCoreComposer.SHCC__AssetBridgeCapacityInsufficient.selector, CORE_AMOUNT - 1, CORE_AMOUNT
            )
        );

        endpointCallsLzReceive();

        assertEq(token.totalSupply(), 0);
        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.balanceOf(composer.assetBridge()), 0);
    }

    function test_receiveHyperCore_revert_recipientNotActivated_rollsBackMint()
        public
        withBridgeSet
        withMintTokenAdded
        withHyperCoreComposerSet
    {
        useHyperCoreMessage();
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, false);
        vm.expectRevert(abi.encodeWithSelector(SynapseHyperCoreComposer.SHCC__CoreUserNotActivated.selector, recipient));

        endpointCallsLzReceive();

        assertEq(token.totalSupply(), 0);
        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.balanceOf(composer.assetBridge()), 0);
    }

    function test_receiveHyperCore_revert_dust_rollsBackMint()
        public
        withBridgeSet
        withMintTokenAdded
        withHyperCoreComposerSet
    {
        amount += 1;
        useHyperCoreMessage();
        vm.expectRevert(
            abi.encodeWithSelector(
                SynapseHyperCoreComposer.SHCC__AmountNotRepresentable.selector, amount, composer.AMOUNT_SCALE()
            )
        );

        endpointCallsLzReceive();

        assertEq(token.totalSupply(), 0);
        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.balanceOf(composer.assetBridge()), 0);
    }

    // ═══════════════════════════════════════ TEST: WITHDRAW-DEPOSIT TOKEN ════════════════════════════════════════════

    function test_receive_withdrawDeposit() public withBridgeSet withWithdrawTokenAdded {
        // Expected action: bridge.withdraw
        vm.expectCall({
            callee: bridge,
            data: abi.encodeCall(SynapseBridgeMock.withdraw, (recipient, address(token), amount, 0, MOCK_GUID))
        });
        // Expected event
        expectEventTokenReceived(SRC_EID, recipient, address(token), amount, MOCK_GUID);
        endpointCallsLzReceive();
    }

    function test_receive_withdrawDeposit_revert_bridgeNotSet() public withWithdrawTokenAdded {
        expectRevertBridgeNotSet();
        endpointCallsLzReceive();
    }

    function test_receive_withdrawDeposit_revert_tokenNotAdded() public withBridgeSet {
        expectRevertLocalPairNotFound(SRC_EID, remoteToken);
        endpointCallsLzReceive();
    }

    function test_receive_withdrawDeposit_revert_bridgeNotSetTokenNotAdded() public {
        expectRevertBridgeNotSet();
        endpointCallsLzReceive();
    }

    function test_receive_withdrawDeposit_revert_eidUnknown() public withBridgeSet withWithdrawTokenAdded {
        vm.expectRevert();
        endpointCallsLzReceive(UNKNOWN_EID, REMOTE_ADAPTER);
    }

    function test_receive_withdrawDeposit_revert_eidUnknown_withPeerAdded()
        public
        withBridgeSet
        withWithdrawTokenAdded
    {
        adapter.setPeer(UNKNOWN_EID, REMOTE_ADAPTER);
        expectRevertLocalPairNotFound(UNKNOWN_EID, remoteToken);
        endpointCallsLzReceive(UNKNOWN_EID, REMOTE_ADAPTER);
    }

    function test_receive_withdrawDeposit_revert_senderUnknown() public withBridgeSet withWithdrawTokenAdded {
        vm.expectRevert();
        endpointCallsLzReceive(SRC_EID, keccak256("Unknown"));
    }
}
