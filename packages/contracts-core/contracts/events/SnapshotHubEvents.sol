// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the SnapshotHub contract
abstract contract SnapshotHubEvents {
    /**
     * @notice Emitted when a new Attestation is saved (derived from a Notary snapshot).
     * @param attestation   Raw payload with attestation data
     */
    event AttestationSaved(bytes attestation);

    /**
     * @notice Emitted when a new Origin State is saved from a Guard snapshot.
     * @param state     Raw payload with state data
     */
    event StateSaved(bytes state);
}
