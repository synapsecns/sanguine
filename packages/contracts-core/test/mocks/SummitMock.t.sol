// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { InterfaceSummit } from "../../contracts/interfaces/InterfaceSummit.sol";
import { SnapshotHubMock } from "./hubs/SnapshotHubMock.t.sol";
import { SystemContractMock } from "./system/SystemContractMock.t.sol";

// solhint-disable no-empty-blocks
contract SummitMock is SnapshotHubMock, SystemContractMock, InterfaceSummit {
    /// @notice Prevents this contract from being included in the coverage report
    function testSummitMock() external {}

    function submitSnapshot(bytes memory _snapPayload, bytes memory _snapSignature)
        external
        returns (bytes memory attPayload)
    {}

    function verifyAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool isValid)
    {}

    function verifyAttestationReport(bytes memory _arPayload, bytes memory _arSignature)
        external
        returns (bool isValid)
    {}

    function addAgent(uint32 _domain, address _account) external returns (bool) {}

    function removeAgent(uint32 _domain, address _account) external returns (bool) {}

    function getLatestState(uint32 _origin) external view returns (bytes memory statePayload) {}
}
