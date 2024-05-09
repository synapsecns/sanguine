// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {PingPongApp} from "../../contracts/apps/examples/PingPongApp.sol";

import {ICIntegrationTest, InterchainEntry, InterchainTransaction} from "./ICIntegration.t.sol";

import {OptionsV1} from "../../contracts/libs/Options.sol";

// solhint-disable custom-errors
// solhint-disable ordering
abstract contract PingPongIntegrationTest is ICIntegrationTest {
    uint256 public constant PING_PONG_BALANCE = 1000 ether;
    uint256 public constant COUNTER = 42;

    OptionsV1 public ppOptions = OptionsV1({gasLimit: 500_000, gasAirdrop: 0});

    event PingReceived(uint256 counter, uint64 dbNonce, uint64 entryIndex);
    event PingSent(uint256 counter, uint64 dbNonce, uint64 entryIndex);

    /// @dev Should deploy the tested app and return its address.
    function deployApp() internal override returns (address app) {
        app = address(new PingPongApp(address(this)));
        deal(app, PING_PONG_BALANCE);
    }

    function expectPingPongEventPingReceived(uint256 counter, FullEntry memory fullEntry) internal {
        vm.expectEmit(localApp());
        emit PingReceived(counter, fullEntry.dbNonce, fullEntry.entryIndex);
    }

    function expectPingPongEventPingSent(uint256 counter, FullEntry memory fullEntry) internal {
        vm.expectEmit(localApp());
        emit PingSent(counter, fullEntry.dbNonce, fullEntry.entryIndex);
    }

    // ═══════════════════════════════════════════ COMPLEX SERIES CHECKS ═══════════════════════════════════════════════

    function expectEventsPingSent(
        uint256 counter,
        InterchainTransaction memory icTx,
        FullEntry memory fullEntry,
        uint256 verificationFee,
        uint256 executionFee
    )
        internal
    {
        expectEventsMessageSent(icTx, fullEntry, verificationFee, executionFee);
        expectPingPongEventPingSent(counter, fullEntry);
    }

    // ═══════════════════════════════════════════════ DATA HELPERS ════════════════════════════════════════════════════

    function srcPingPongApp() internal view returns (PingPongApp) {
        return PingPongApp(payable(srcApp));
    }

    function dstPingPongApp() internal view returns (PingPongApp) {
        return PingPongApp(payable(dstApp));
    }

    function getSrcOptions() internal view override returns (OptionsV1 memory) {
        return ppOptions;
    }

    /// @notice Message that source chain PingPongApp sends to destination chain.
    function getSrcMessage() internal pure override returns (bytes memory) {
        return abi.encode(COUNTER);
    }

    function getDstOptions() internal view override returns (OptionsV1 memory) {
        return ppOptions;
    }

    /// @notice Message that destination chain PingPongApp sends back to source chain.
    function getDstMessage() internal pure override returns (bytes memory) {
        return abi.encode(COUNTER - 1);
    }
}
