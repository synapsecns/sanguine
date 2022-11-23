// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../events/NotaryRegistryHarnessEvents.sol";
import { LocalDomainContext } from "../../../contracts/context/LocalDomainContext.sol";
import { DomainNotaryRegistry } from "../../../contracts/registry/DomainNotaryRegistry.sol";

contract DomainNotaryRegistryHarness is
    NotaryRegistryHarnessEvents,
    LocalDomainContext,
    DomainNotaryRegistry
{
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

    /**
     * @notice Hook that is called just before a Notary is added for specified domain.
     */
    function _beforeNotaryAdded(uint32 _domain, address _notary) internal virtual override {
        require(!isNotary(_domain, _notary), "!beforeNotaryAdded");
        emit BeforeNotaryAdded(_domain, _notary);
    }

    /**
     * @notice Hook that is called right after a Notary is added for specified domain.
     */
    function _afterNotaryAdded(uint32 _domain, address _notary) internal virtual override {
        require(isNotary(_domain, _notary), "!afterNotaryAdded");
        emit AfterNotaryAdded(_domain, _notary);
    }

    /**
     * @notice Hook that is called just before a Notary is removed from specified domain.
     */
    function _beforeNotaryRemoved(uint32 _domain, address _notary) internal virtual override {
        require(isNotary(_domain, _notary), "!beforeNotaryRemoved");
        emit BeforeNotaryRemoved(_domain, _notary);
    }

    /**
     * @notice Hook that is called right after a Notary is removed from specified domain.
     */
    function _afterNotaryRemoved(uint32 _domain, address _notary) internal virtual override {
        require(!isNotary(_domain, _notary), "!afterNotaryRemoved");
        emit AfterNotaryRemoved(_domain, _notary);
    }
}
