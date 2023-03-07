// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../contracts/interfaces/InterfaceDestination.sol";
import "./hubs/AttestationHubMock.t.sol";
import "./system/AgentRegistryMock.t.sol";
import "./system/SystemContractMock.t.sol";

// solhint-disable no-empty-blocks
contract DestinationMock is
    AttestationHubMock,
    AgentRegistryMock,
    SystemContractMock,
    InterfaceDestination
{
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
