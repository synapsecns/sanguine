// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract ClientHarnessEvents {
    event LogMessage(uint32 origin, uint32 nonce, bytes message);
}
