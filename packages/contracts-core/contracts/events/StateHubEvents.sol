// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the StateHub contract
abstract contract StateHubEvents {
    /**
     * @notice Emitted when a new Origin State is saved after a message was sent.
     * @param state     Raw payload with state data
     */
    event StateSaved(bytes state);
}
