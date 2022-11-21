// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract GuardRegistryEvents {
    /**
     * @notice Emitted when a new Guard is added.
     * @param guard    Address of the added guard
     */
    event GuardAdded(address guard);

    /**
     * @notice Emitted when a Guard is removed.
     * @param guard    Address of the removed guard
     */
    event GuardRemoved(address guard);
}
