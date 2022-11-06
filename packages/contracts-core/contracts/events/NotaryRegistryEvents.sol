// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract NotaryRegistryEvents {
    /*
     * @notice Emitted when a new Notary is added.
     * @param domain    Domain where a Notary was added
     * @param notary    Address of the added notary
     */
    event NotaryAdded(uint32 indexed domain, address notary);

    /**
     * @notice Emitted when a new Notary is removed.
     * @param domain    Domain where a Notary was removed
     * @param notary    Address of the removed notary
     */
    event NotaryRemoved(uint32 indexed domain, address notary);
}
