// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainAppBaseEvents {
    event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod);
    event InterchainClientSet(address interchainClient);
}
