// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Minimal interface for the Interchain App to work with the Interchain Client.
interface IInterchainApp {
    function appReceive(
        uint64 srcChainId,
        bytes32 sender,
        uint64 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        external
        payable;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules);
}
