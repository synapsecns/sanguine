// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { DomainNotaryRegistry } from "../../contracts/registry/DomainNotaryRegistry.sol";

contract DomainNotaryRegistryHarness is DomainNotaryRegistry {
    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 _trackedDomain) DomainNotaryRegistry(_trackedDomain) {}

    function addNotary(address _notary) public returns (bool) {
        return _addNotary(_notary);
    }

    function removeNotary(address _notary) public returns (bool) {
        return _removeNotary(_notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }
}
