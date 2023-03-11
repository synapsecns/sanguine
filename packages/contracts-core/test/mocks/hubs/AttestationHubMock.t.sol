// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    DestinationAttestation,
    IAttestationHub
} from "../../../contracts/interfaces/IAttestationHub.sol";

// solhint-disable no-empty-blocks
contract AttestationHubMock is IAttestationHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testAttestationHubMock() external {}

    function attestationsAmount() external view returns (uint256) {}

    function getAttestation(uint256 _index)
        external
        view
        returns (bytes32 root, DestinationAttestation memory destAtt)
    {}
}
