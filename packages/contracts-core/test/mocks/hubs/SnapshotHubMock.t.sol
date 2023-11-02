// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ISnapshotHub} from "../../../contracts/interfaces/ISnapshotHub.sol";
import {BaseMock} from "../base/BaseMock.t.sol";

// solhint-disable no-empty-blocks
contract SnapshotHubMock is BaseMock, ISnapshotHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testSnapshotHubMock() external {}

    function isValidAttestation(bytes memory) external view returns (bool isValid) {
        return getReturnValueBool();
    }

    function getAttestation(uint32 attNonce)
        external
        view
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas)
    {}

    function getLatestAgentState(uint32 origin, address agent) external view returns (bytes memory statePayload) {}

    function getLatestNotaryAttestation(address notary)
        external
        view
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas)
    {}

    function getGuardSnapshot(uint256 index)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {}

    function getNotarySnapshot(uint256 index)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {}

    function getNotarySnapshot(bytes memory attPayload)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {}

    function getSnapshotProof(uint32 attNonce, uint8 stateIndex) external view returns (bytes32[] memory snapProof) {}
}
