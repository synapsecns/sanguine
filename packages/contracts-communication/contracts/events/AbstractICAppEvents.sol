// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract AbstractICAppEvents {
    event InterchainClientAdded(address client);
    event InterchainClientRemoved(address client);
    event LatestClientSet(address client);
}
