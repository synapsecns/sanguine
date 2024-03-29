// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Minimal interface for the Interchain App to work with the Interchain Client.
interface IInterchainApp {
    /// @notice Allows the Interchain Client to pass the message to the Interchain App.
    /// @dev App is responsible for keeping track of interchain clients, and must verify the message sender.
    /// @param srcChainId   Chain ID of the source chain, where the message was sent from.
    /// @param sender       Sender address on the source chain, as a bytes32 value.
    /// @param dbNonce      The Interchain DB nonce of the batch containing the message entry.
    /// @param entryIndex   The index of the message entry within the batch.
    /// @param message      The message being sent.
    function appReceive(
        uint64 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        external
        payable;

    /// @notice Returns the verification configuration of the Interchain App.
    /// @dev This configuration is used by the Interchain Client to verify that message has been confirmed
    /// by the Interchain Modules on the destination chain.
    /// Note: V1 version of AppConfig includes the required responses count, and optimistic period after which
    /// the message is considered confirmed by the module. Following versions may include additional fields.
    /// @return appConfig    The versioned configuration of the Interchain App, encoded as bytes.
    /// @return modules      The list of Interchain Modules that app is trusting to confirm the messages.
    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules);
}
