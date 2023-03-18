// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISnapshotHub } from "../../../contracts/interfaces/ISnapshotHub.sol";

// solhint-disable no-empty-blocks
contract SnapshotHubMock is ISnapshotHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testSnapshotHubMock() external {}

    function isValidAttestation(bytes memory _attPayload) external view returns (bool isValid) {}

    function getAttestation(uint32 _nonce) external view returns (bytes memory attPayload) {}

    function getLatestAgentState(uint32 _origin, address _agent)
        external
        view
        returns (bytes memory statePayload)
    {}

    function getGuardSnapshot(uint256 _index)
        external
        view
        returns (bytes memory snapshotPayload)
    {}

    function getNotarySnapshot(uint256 _nonce)
        external
        view
        returns (bytes memory snapshotPayload)
    {}

    function getNotarySnapshot(bytes memory _attPayload)
        external
        view
        returns (bytes memory snapshotPayload)
    {}

    function getSnapshotProof(uint256 _nonce, uint256 _stateIndex)
        external
        view
        returns (bytes32[] memory snapProof)
    {}
}
