// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { LocalDomainContext } from "../../../contracts/context/LocalDomainContext.sol";
import { DomainNotaryRegistry } from "../../../contracts/registry/DomainNotaryRegistry.sol";

contract DomainNotaryRegistryHarness is LocalDomainContext, DomainNotaryRegistry {
    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 _trackedDomain) LocalDomainContext(_trackedDomain) {}

    function addNotary(address _notary) public returns (bool) {
        return _addNotary(_notary);
    }

    function addNotary(uint32 _domain, address _notary) public returns (bool) {
        return _addNotary(_domain, _notary);
    }

    function removeNotary(address _notary) public returns (bool) {
        return _removeNotary(_notary);
    }

    function removeNotary(uint32 _domain, address _notary) public returns (bool) {
        return _removeNotary(_domain, _notary);
    }

    function isNotary(address _notary) public view returns (bool) {
        return _isNotary(_notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }
}
