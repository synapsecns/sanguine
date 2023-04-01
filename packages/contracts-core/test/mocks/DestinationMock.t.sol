// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ExecutionAttestation, InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";
import {ExecutionHubMock} from "./hubs/ExecutionHubMock.t.sol";
import {DisputeHubMock} from "./hubs/DisputeHubMock.t.sol";
import {SystemContractMock} from "./system/SystemContractMock.t.sol";
import {SystemRegistryMock} from "./system/SystemRegistryMock.t.sol";

// solhint-disable no-empty-blocks
contract DestinationMock is
    ExecutionHubMock,
    DisputeHubMock,
    SystemRegistryMock,
    SystemContractMock,
    InterfaceDestination
{
    /// @notice Prevents this contract from being included in the coverage report
    function testDestinationMock() external {}

    function passAgentRoot() external returns (bool rootPassed, bool rootPending) {}

    function submitAttestation(bytes memory attPayload, bytes memory attSignature)
        external
        returns (bool wasAccepted)
    {}

    function submitAttestationReport(bytes memory arPayload, bytes memory arSignature, bytes memory attSignature)
        external
        returns (bool wasAccepted)
    {}

    /**
     * @notice Returns the total amount of Notaries attestations that have been accepted.
     */
    function attestationsAmount() external view returns (uint256) {}

    /**
     * @notice Returns an attestation from the list of all accepted Notary attestations.
     * @dev Index refers to attestation's snapshot root position in `roots` array.
     * @param index    Attestation index
     * @return root    Snapshot root for the attestation
     * @return execAtt Rest of attestation data that Destination keeps track of
     */
    function getAttestation(uint256 index) external view returns (bytes32 root, ExecutionAttestation memory execAtt) {}

    function destStatus() external view returns (uint48 snapRootTime, uint48 agentRootTime, address notary) {}

    function nextAgentRoot() external view returns (bytes32) {}
}
