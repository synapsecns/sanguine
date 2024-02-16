// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainModuleEvents {
    event VerificationRequested(uint256 indexed destChainId, bytes entry);
}
