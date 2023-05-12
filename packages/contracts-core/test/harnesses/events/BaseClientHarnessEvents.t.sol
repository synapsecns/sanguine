// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract BaseClientHarnessEvents {
    event BaseMessageReceived(uint256 msgValue, uint32 origin, uint32 nonce, uint32 version, bytes content);
}
