// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface ISnapshotHub {
    /**
     * @notice Returns the state with the highest known nonce submitted by a given Guard.
     * @param _origin       Domain of origin chain
     * @param _guard        Guard address
     * @return statePayload Raw payload with guard latest state for origin
     */
    function getLatestState(uint32 _origin, address _guard)
        external
        view
        returns (bytes memory statePayload);

    /**
     * @notice Returns Notary snapshot that was used for creating an attestation with a given nonce.
     * @dev Reverts if attestation with given nonce hasn't been created yet.
     * @param _nonce            Nonce for the attestation
     * @return snapshotPayload  Raw payload with Notary snapshot used for creating the attestation
     */
    function getSnapshot(uint256 _nonce) external view returns (bytes memory snapshotPayload);

    /**
     * @notice Returns Notary snapshot that was used for creating a given attestation.
     * @dev Reverts if either of this is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation is invalid (doesn't have a matching Notary snapshot).
     * @param _snapAttPayload   Raw payload with attestation data
     * @return snapshotPayload  Raw payload with Notary snapshot used for creating the attestation
     */
    function getSnapshot(bytes memory _snapAttPayload)
        external
        view
        returns (bytes memory snapshotPayload);
}
