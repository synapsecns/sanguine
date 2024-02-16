// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainModuleEvents {
    event VerificationRequested(uint256 indexed destChainId, bytes entry);
}
