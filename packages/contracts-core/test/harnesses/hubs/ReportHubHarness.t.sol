// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";
import { ReportHub } from "../../../contracts/hubs/ReportHub.sol";
import { AbstractNotaryRegistry } from "../../../contracts/registry/AbstractNotaryRegistry.sol";

import { GuardRegistryHarness } from "../registry/GuardRegistryHarness.t.sol";
import { GlobalNotaryRegistryHarness } from "../registry/GlobalNotaryRegistryHarness.t.sol";
import { ReportHubHarnessEvents } from "../events/ReportHubHarnessEvents.sol";

contract ReportHubHarness is
    ReportHubHarnessEvents,
    ReportHub,
    GuardRegistryHarness,
    GlobalNotaryRegistryHarness
{
    using TypedMemView for bytes29;

    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal override returns (bool) {
        emit LogReport(_guard, _notary, _attestationView.clone(), _reportView.clone(), _report);
        return true;
    }

    /**
     * @notice Hook that is called just before a Notary is added for specified domain.
     */
    function _beforeNotaryAdded(uint32 _domain, address _notary)
        internal
        virtual
        override(AbstractNotaryRegistry, GlobalNotaryRegistryHarness)
    {
        AbstractNotaryRegistry._beforeNotaryAdded(_domain, _notary);
    }

    /**
     * @notice Hook that is called right after a Notary is added for specified domain.
     */
    function _afterNotaryAdded(uint32 _domain, address _notary)
        internal
        virtual
        override(AbstractNotaryRegistry, GlobalNotaryRegistryHarness)
    {
        AbstractNotaryRegistry._afterNotaryAdded(_domain, _notary);
    }

    /**
     * @notice Hook that is called just before a Notary is removed from specified domain.
     */
    function _beforeNotaryRemoved(uint32 _domain, address _notary)
        internal
        virtual
        override(AbstractNotaryRegistry, GlobalNotaryRegistryHarness)
    {
        AbstractNotaryRegistry._beforeNotaryRemoved(_domain, _notary);
    }

    /**
     * @notice Hook that is called right after a Notary is removed from specified domain.
     */
    function _afterNotaryRemoved(uint32 _domain, address _notary)
        internal
        virtual
        override(AbstractNotaryRegistry, GlobalNotaryRegistryHarness)
    {
        AbstractNotaryRegistry._afterNotaryRemoved(_domain, _notary);
    }
}
