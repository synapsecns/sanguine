// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ISnapshotHub} from "../../../contracts/interfaces/ISnapshotHub.sol";

// solhint-disable no-empty-blocks
contract SnapshotHubMock is ISnapshotHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testSnapshotHubMock() external {}

    function isValidAttestation(bytes memory attPayload) external view returns (bool isValid) {}

    function getAttestation(uint32 nonce) external view returns (bytes memory attPayload) {}

    function getLatestAgentState(uint32 origin, address agent) external view returns (bytes memory statePayload) {}

    function getLatestNotaryAttestation(address notary) external view returns (bytes memory attPayload) {}

    function getGuardSnapshot(uint256 index) external view returns (bytes memory snapshotPayload) {}

    function getNotarySnapshot(uint256 nonce) external view returns (bytes memory snapshotPayload) {}

    function getNotarySnapshot(bytes memory attPayload) external view returns (bytes memory snapshotPayload) {}

    function getSnapshotProof(uint256 nonce, uint256 stateIndex) external view returns (bytes32[] memory snapProof) {}
}
