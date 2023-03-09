// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    DestinationAttestation,
    IAttestationHub
} from "../../../contracts/interfaces/IAttestationHub.sol";
import { ExcludeCoverage } from "../ExcludeCoverage.sol";

// solhint-disable no-empty-blocks
contract AttestationHubMock is ExcludeCoverage, IAttestationHub {
    function attestationsAmount() external view returns (uint256) {}

    function getAttestation(uint256 _index)
        external
        view
        returns (bytes32 root, DestinationAttestation memory destAtt)
    {}
}
