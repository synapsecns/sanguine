// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceOrigin} from "../../contracts/interfaces/InterfaceOrigin.sol";
import {StateHubMock} from "./hubs/StateHubMock.t.sol";
import {SystemContractMock} from "./system/SystemContractMock.t.sol";
import {SystemRegistryMock} from "./system/SystemRegistryMock.t.sol";

// solhint-disable no-empty-blocks
contract OriginMock is StateHubMock, SystemRegistryMock, SystemContractMock, InterfaceOrigin {
    /// @notice Prevents this contract from being included in the coverage report
    function testOriginMock() external {}

    function dispatch(
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds,
        bytes memory tipsPayload,
        bytes memory content
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {}

    function verifyAttestation(
        uint256 stateIndex,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValid) {}

    function verifyAttestationWithProof(
        uint256 stateIndex,
        bytes memory statePayload,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValid) {}

    function verifySnapshot(uint256 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValid)
    {}

    function verifyStateReport(bytes memory srPayload, bytes memory srSignature) external returns (bool isValid) {}
}
