// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter, SynapseBridgeAdapter, SynapseBridgeAdapterTest} from "./SBA.t.sol";

import {SynapseBridgeMock} from "./mocks/SynapseBridgeMock.sol";
import {TestToken} from "./mocks/TestToken.sol";

import {Origin} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseBridgeAdapterDstTest is SynapseBridgeAdapterTest {
    address internal bridge;
    TestToken internal token;

    address internal executor = makeAddr("Executor");
    address internal recipient = makeAddr("Recipient");
    uint256 internal amount = 0.123456 ether;

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

    function afterAdapterDeployed() internal virtual override {
        adapter.setPeer(SRC_EID, REMOTE_ADAPTER);

        bridge = address(new SynapseBridgeMock());
        token = new TestToken();

        bridgeMessage = bridgeMessageLib.encodeBridgeMessage(recipient, remoteToken, amount);
    }

    function deployAdapter() internal virtual override returns (SynapseBridgeAdapter) {
        return new SynapseBridgeAdapter(endpoint, address(this));
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
        expectRevertRemoteTokenUnknown(SRC_EID, remoteToken);
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
        expectRevertRemoteTokenUnknown(UNKNOWN_EID, remoteToken);
        endpointCallsLzReceive(UNKNOWN_EID, REMOTE_ADAPTER);
    }

    function test_receive_mintBurn_revert_senderUnknown() public withBridgeSet withMintTokenAdded {
        vm.expectRevert();
        endpointCallsLzReceive(SRC_EID, keccak256("Unknown"));
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
        expectRevertRemoteTokenUnknown(SRC_EID, remoteToken);
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
        expectRevertRemoteTokenUnknown(UNKNOWN_EID, remoteToken);
        endpointCallsLzReceive(UNKNOWN_EID, REMOTE_ADAPTER);
    }

    function test_receive_withdrawDeposit_revert_senderUnknown() public withBridgeSet withWithdrawTokenAdded {
        vm.expectRevert();
        endpointCallsLzReceive(SRC_EID, keccak256("Unknown"));
    }
}
