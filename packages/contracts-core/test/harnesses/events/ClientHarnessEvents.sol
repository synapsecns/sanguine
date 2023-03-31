// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract ClientHarnessEvents {
    event LogClientMessage(uint32 origin, uint32 nonce, bytes content);
}
