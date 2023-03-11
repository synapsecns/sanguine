// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    InterfaceDestination,
    ORIGIN_TREE_DEPTH
} from "../../contracts/interfaces/InterfaceDestination.sol";
import { AttestationHubMock } from "./hubs/AttestationHubMock.t.sol";
import { AgentRegistryMock } from "./system/AgentRegistryMock.t.sol";
import { SystemContractMock } from "./system/SystemContractMock.t.sol";

// solhint-disable no-empty-blocks
contract DestinationMock is
    AttestationHubMock,
    AgentRegistryMock,
    SystemContractMock,
    InterfaceDestination
{
    /// @notice Prevents this contract from being included in the coverage report
    function testDestinationMock() external {}

    function execute(
        bytes memory _message,
        bytes32[ORIGIN_TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) external {}

    function submitAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool wasAccepted)
    {}
}
