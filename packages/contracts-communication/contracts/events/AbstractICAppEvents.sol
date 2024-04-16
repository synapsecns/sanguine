// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract AbstractICAppEvents {
    /// @notice Emitted when a new interchain client is added.
    /// This client will be able to send messages to the app.
    /// @param client   The address of the client.
    event InterchainClientAdded(address client);

    /// @notice Emitted when an interchain client is removed.
    /// @param client   The address of the client.
    event InterchainClientRemoved(address client);

    /// @notice Emitted when the latest interchain client is set.
    /// This client will be used by the app to send messages.
    /// @param client   The address of the client.
    event LatestClientSet(address client);
}
