// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { InterfaceOrigin } from "../../contracts/interfaces/InterfaceOrigin.sol";
import { StateHubMock } from "./hubs/StateHubMock.t.sol";
import { SystemContractMock } from "./system/SystemContractMock.t.sol";
import { SystemRegistryMock } from "./system/SystemRegistryMock.t.sol";

// solhint-disable no-empty-blocks
contract OriginMock is StateHubMock, SystemRegistryMock, SystemContractMock, InterfaceOrigin {
    /// @notice Prevents this contract from being included in the coverage report
    function testOriginMock() external {}

    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {}

    function verifyAttestation(
        uint256 _stateIndex,
        bytes memory _snapPayload,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool isValid) {}

    function verifyAttestationWithProof(
        uint256 _stateIndex,
        bytes memory _statePayload,
        bytes32[] memory _snapProof,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool isValid) {}

    function verifySnapshot(
        uint256 _stateIndex,
        bytes memory _snapPayload,
        bytes memory _snapSignature
    ) external returns (bool isValid) {}

    function verifyStateReport(bytes memory _srPayload, bytes memory _srSignature)
        external
        returns (bool isValid)
    {}
}
