// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { AbstractNotaryRegistry } from "../../../contracts/registry/AbstractNotaryRegistry.sol";

abstract contract NotaryRegistryHarness is AbstractNotaryRegistry {
    function addNotary(uint32 _domain, address _notary) public returns (bool) {
        return _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) public returns (bool) {
        return _removeNotary(_domain, _notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }
}
