// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

contract GlobalNotaryRegistryHarnessEvents {
    event HookDomainActive(uint32 domain, address notary);

    event HookDomainInactive(uint32 domain, address notary);
}
