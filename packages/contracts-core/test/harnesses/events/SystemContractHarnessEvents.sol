// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract SystemContractHarnessEvents {
    event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt);

    event UsualCall(address recipient, uint256 newValue);
    event OnlyLocalCall(address recipient, uint256 newValue);
    event OnlyOriginCall(address recipient, uint256 newValue);
    event OnlyDestinationCall(address recipient, uint256 newValue);
    event OnlyOriginDestinationCall(address recipient, uint256 newValue);
    event OnlyTwoHoursCall(address recipient, uint256 newValue);
    event OnlySynapseChainCall(address recipient, uint256 newValue);
}
