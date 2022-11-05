// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract NotaryManagerEvents {
    /**
     * @notice Emitted when a new origin is set
     * @param origin The address of the new origin contract
     */
    event NewOrigin(address origin);

    /**
     * @notice Emitted when a new notary is set
     * @param notary The address of the new notary
     */
    event NewNotary(address notary);

    /**
     * @notice Emitted when slashNotary is called
     */
    event FakeSlashed(address reporter);
}
