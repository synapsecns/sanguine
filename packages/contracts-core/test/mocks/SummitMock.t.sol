// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../contracts/interfaces/InterfaceSummit.sol";
import "./hubs/SnapshotHubMock.t.sol";
import "./system/SystemContractMock.t.sol";

// solhint-disable no-empty-blocks
contract SummitMock is SnapshotHubMock, SystemContractMock, InterfaceSummit {
    function submitSnapshot(bytes memory _snapPayload, bytes memory _snapSignature)
        external
        returns (bool wasAccepted)
    {}

    function verifyAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool isValid)
    {}

    function addAgent(uint32 _domain, address _account) external returns (bool) {}

    function removeAgent(uint32 _domain, address _account) external returns (bool) {}
}
