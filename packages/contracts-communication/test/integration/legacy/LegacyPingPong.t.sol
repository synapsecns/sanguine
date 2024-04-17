// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ICAppV1} from "../../../contracts/apps/ICAppV1.sol";
import {LegacyPingPong} from "../../../contracts/legacy/LegacyPingPong.sol";
import {LegacyReceiver} from "../../../contracts/legacy/LegacyReceiver.sol";
import {LegacyMessageLib} from "../../../contracts/legacy/libs/LegacyMessage.sol";
import {LegacyOptionsLib} from "../../../contracts/legacy/libs/LegacyOptions.sol";
import {TypeCasts} from "../../../contracts/libs/TypeCasts.sol";

import {
    ICIntegrationTest,
    InterchainBatch,
    InterchainEntry,
    InterchainTransaction,
    InterchainTxDescriptor,
    OptionsV1
} from "../ICIntegration.t.sol";
import {MessageBusHarness} from "../../harnesses/MessageBusHarness.sol";

// solhint-disable custom-errors
// solhint-disable ordering
abstract contract LegacyPingPongIntegrationTest is ICIntegrationTest {
    uint256 public constant PING_PONG_BALANCE = 1000 ether;
    uint256 public constant COUNTER = 42;
    uint256 public constant GAS_LIMIT = 500_000;

    uint64 public constant SRC_MSG_BUS_NONCE = 5;
    uint64 public constant DST_MSG_BUS_NONCE = 15;

    OptionsV1 public icOptions = OptionsV1({gasLimit: GAS_LIMIT, gasAirdrop: 0});
    bytes public legacyOptions = LegacyOptionsLib.encodeLegacyOptions(GAS_LIMIT);

    address public srcPingPong;
    address public dstPingPong;

    event PingReceived(uint256 counter);
    event PingSent(uint256 counter);

    function setUp() public virtual override {
        super.setUp();
        srcPingPong = deployLegacyPingPong();
        vm.label(srcPingPong, "Src PingPong");
        dstPingPong = deployLegacyPingPong();
        vm.label(dstPingPong, "Dst PingPong");
        configureLocalPingPong();
        setMsgBusNonce();
    }

    /// @dev Should deploy the tested app and return its address.
    function deployApp() internal override returns (address app) {
        return address(new MessageBusHarness(address(this)));
    }

    function deployLegacyPingPong() internal returns (address pingPong) {
        pingPong = address(new LegacyPingPong(address(this)));
        deal(pingPong, PING_PONG_BALANCE);
    }

    /// @dev Local app is MessageBus, need to grant the IC_GOVERNOR_ROLE, as it's not done in the constructor.
    function configureLocalApp() internal virtual override {
        address app = localApp();
        ICAppV1(app).grantRole(ICAppV1(app).IC_GOVERNOR_ROLE(), address(this));
        super.configureLocalApp();
    }

    function configureLocalPingPong() internal {
        address pingPong = localLegacyPingPong();
        bytes32 remotePingPongBytes32 = TypeCasts.addressToBytes32(isSourceChainTest() ? dstPingPong : srcPingPong);
        LegacyReceiver(pingPong).setMessageBus(localApp());
        LegacyReceiver(pingPong).setTrustedRemote(remoteChainId(), remotePingPongBytes32);
    }

    function setMsgBusNonce() internal {
        uint64 nonce = isSourceChainTest() ? SRC_MSG_BUS_NONCE : DST_MSG_BUS_NONCE;
        MessageBusHarness(localApp()).setNonce(nonce);
    }

    function expectPingPongEventPingReceived(uint256 counter) internal {
        vm.expectEmit(localLegacyPingPong());
        emit PingReceived(counter);
    }

    function expectPingPongEventPingSent(uint256 counter) internal {
        vm.expectEmit(localLegacyPingPong());
        emit PingSent(counter);
    }

    function expectPingPongCall() internal {
        bytes memory expectedCalldata = abi.encodeCall(
            LegacyReceiver.executeMessage,
            (TypeCasts.addressToBytes32(srcPingPong), SRC_CHAIN_ID, getSrcLegacyMessage(), address(icClient))
        );
        vm.expectCall({callee: dstPingPong, msgValue: 0, data: expectedCalldata, count: 1});
    }

    // ═══════════════════════════════════════════ COMPLEX SERIES CHECKS ═══════════════════════════════════════════════

    function expectEventsPingSent(
        uint256 counter,
        InterchainTransaction memory icTx,
        InterchainEntry memory entry,
        uint256 verificationFee,
        uint256 executionFee
    )
        internal
    {
        expectEventsMessageSent(icTx, entry, verificationFee, executionFee);
        expectPingPongEventPingSent(counter);
    }

    // ═══════════════════════════════════════════════ DATA HELPERS ════════════════════════════════════════════════════

    function localLegacyPingPong() internal view returns (address) {
        return isSourceChainTest() ? srcPingPong : dstPingPong;
    }

    function srcLegacyPingPong() internal view returns (LegacyPingPong) {
        return LegacyPingPong(payable(srcPingPong));
    }

    function dstLegacyPingPong() internal view returns (LegacyPingPong) {
        return LegacyPingPong(payable(dstPingPong));
    }

    function getSrcOptions() internal pure override returns (OptionsV1 memory) {
        return OptionsV1({gasLimit: GAS_LIMIT, gasAirdrop: 0});
    }

    function getSrcMessage() internal view override returns (bytes memory) {
        return LegacyMessageLib.encodeLegacyMessage({
            srcSender: srcPingPong,
            dstReceiver: dstPingPong,
            srcNonce: SRC_MSG_BUS_NONCE,
            message: getSrcLegacyMessage()
        });
    }

    function getSrcLegacyMessage() internal pure returns (bytes memory) {
        return abi.encode(COUNTER);
    }

    function getDstOptions() internal pure override returns (OptionsV1 memory) {
        return OptionsV1({gasLimit: GAS_LIMIT, gasAirdrop: 0});
    }

    function getDstMessage() internal view override returns (bytes memory) {
        return LegacyMessageLib.encodeLegacyMessage({
            srcSender: dstPingPong,
            dstReceiver: srcPingPong,
            srcNonce: DST_MSG_BUS_NONCE,
            message: getDstLegacyMessage()
        });
    }

    function getDstLegacyMessage() internal pure returns (bytes memory) {
        return abi.encode(COUNTER - 1);
    }
}
