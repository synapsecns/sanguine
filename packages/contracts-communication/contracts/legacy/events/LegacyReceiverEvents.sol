// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract LegacyReceiverEvents {
    event MessageBusSet(address messageBus);
    event TrustedRemoteSet(uint256 chainId, bytes32 trustedRemote);
}
