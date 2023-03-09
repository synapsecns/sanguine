// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { InterfaceOrigin } from "../../contracts/interfaces/InterfaceOrigin.sol";
import { StateHubMock } from "./hubs/StateHubMock.t.sol";
import { AgentRegistryMock } from "./system/AgentRegistryMock.t.sol";
import { SystemContractMock } from "./system/SystemContractMock.t.sol";

// solhint-disable no-empty-blocks
contract OriginMock is StateHubMock, AgentRegistryMock, SystemContractMock, InterfaceOrigin {
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {}

    function verifyAttestation(
        bytes memory _snapPayload,
        uint256 _stateIndex,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool isValid) {}

    function verifySnapshot(
        bytes memory _snapPayload,
        uint256 _stateIndex,
        bytes memory _snapSignature
    ) external returns (bool isValid) {}
}
