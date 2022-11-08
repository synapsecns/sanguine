// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../events/GlobalNotaryRegistryHarnessEvents.sol";
import { GlobalNotaryRegistry } from "../../../contracts/registry/GlobalNotaryRegistry.sol";

contract GlobalNotaryRegistryHarness is GlobalNotaryRegistryHarnessEvents, GlobalNotaryRegistry {
    function addNotary(uint32 _domain, address _notary) public returns (bool) {
        return _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) public returns (bool) {
        return _removeNotary(_domain, _notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }

    /**
     * @notice Hook that is called after the specified domain becomes active,
     * i.e. when a Notary is added to the domain, which previously had no active Notaries.
     */
    function _afterDomainBecomesActive(uint32 _domain, address _notary) internal virtual override {
        emit HookDomainActive(_domain, _notary);
    }

    /**
     * @notice Hook that is called after the specified domain becomes inactive,
     * i.e. when the last Notary is removed from the domain.
     */
    function _afterDomainBecomesInactive(uint32 _domain, address _notary)
        internal
        virtual
        override
    {
        emit HookDomainInactive(_domain, _notary);
    }
}
