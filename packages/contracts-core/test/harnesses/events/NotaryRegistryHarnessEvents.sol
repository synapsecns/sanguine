// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

abstract contract NotaryRegistryHarnessEvents {
    event BeforeNotaryAdded(uint32 domain, address notary);
    event AfterNotaryAdded(uint32 domain, address notary);

    event BeforeNotaryRemoved(uint32 domain, address notary);
    event AfterNotaryRemoved(uint32 domain, address notary);
}
