// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {AppConfigV1} from "../libs/AppConfig.sol";

interface IInterchainAppV1 {
    error InterchainApp__ModuleAlreadyAdded(address module);
    error InterchainApp__ModuleNotAdded(address module);

    /// @notice Allows the owner to add the interchain client to the allowed clients set,
    /// and optionally set the latest client to this one.
    /// Note: only the allowed clients can send messages to this app.
    /// Note: the latest client is used for sending messages from this app.
    /// @param client       The address of the interchain client to add.
    /// @param updateLatest Whether to set the latest client to this one.
    function addInterchainClient(address client, bool updateLatest) external;

    /// @notice Allows the owner to remove the interchain client from the allowed clients set.
    /// If the client is the latest client, the latest client is set to the zero address.
    /// @param client       The address of the interchain client to remove.
    function removeInterchainClient(address client) external;

    /// @notice Allows the owner to set the address of the latest interchain client.
    /// @dev The new latest client must be an allowed client or zero address.
    /// Setting the client to zero address effectively pauses the app ability to send messages,
    /// while allowing to receive them.
    /// @param client       The address of the latest interchain client.
    function setLatestInterchainClient(address client) external;

    /// @notice Allows the owner to link the remote app for the given chain ID.
    /// - This address will be used as the receiver for the messages sent from this chain.
    /// - This address will be the only trusted sender for the messages sent to this chain.
    /// @param chainId      The remote chain ID.
    /// @param remoteApp    The address of the remote app to link.
    function linkRemoteApp(uint256 chainId, bytes32 remoteApp) external;

    /// @notice Thin wrapper for `linkRemoteApp` to accept EVM address as a parameter.
    function linkRemoteAppEVM(uint256 chainId, address remoteApp) external;

    /// @notice Allows the owner to add the module to the trusted modules set.
    /// - This set of modules will be used to verify both sent and received messages.
    function addTrustedModule(address module) external;

    /// @notice Allows the owner to remove the module from the trusted modules set.
    function removeTrustedModule(address module) external;

    /// @notice Allows the owner to set the app config for the current app. App config includes:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    function setAppConfigV1(AppConfigV1 memory appConfig) external;

    /// @notice Allows the owner to set the address of the Execution Service.
    /// This address will be used to request execution of the messages sent from this chain,
    /// by supplying the Service's execution fee.
    function setExecutionService(address executionService) external;
}
