// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../../contracts/interfaces/ISnapshotHub.sol";
import "../ExcludeCoverage.sol";

// solhint-disable no-empty-blocks
contract SnapshotHubMock is ExcludeCoverage, ISnapshotHub {
    function isValidAttestation(bytes memory _attPayload) external view returns (bool isValid) {}

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
