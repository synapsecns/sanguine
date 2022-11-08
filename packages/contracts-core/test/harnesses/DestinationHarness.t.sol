// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Destination } from "../../contracts/Destination.sol";
import { GlobalNotaryRegistry } from "../../contracts/registry/GlobalNotaryRegistry.sol";

import { Tips } from "../../contracts/libs/Tips.sol";
import { ISystemRouter } from "../../contracts/interfaces/ISystemRouter.sol";

import { GuardRegistryHarness } from "./registry/GuardRegistryHarness.t.sol";
import { GlobalNotaryRegistryHarness } from "./registry/GlobalNotaryRegistryHarness.t.sol";
import { SystemContractHarness } from "./system/SystemContractHarness.t.sol";
import { DestinationHarnessEvents } from "./events/DestinationHarnessEvents.sol";

contract DestinationHarness is
    DestinationHarnessEvents,
    Destination,
    SystemContractHarness,
    GlobalNotaryRegistryHarness,
    GuardRegistryHarness
{
    using Tips for bytes29;

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _localDomain) Destination(_localDomain) {}

    function setSensitiveValue(uint256 _newValue) external onlySystemRouter {
        sensitiveValue = _newValue;
    }

    function setMessageStatus(
        uint32 _originDomain,
        bytes32 _messageHash,
        bytes32 _status
    ) external {
        messageStatus[_originDomain][_messageHash] = _status;
    }

    function _storeTips(bytes29 _tips) internal override {
        emit LogTips(
            _tips.notaryTip(),
            _tips.broadcasterTip(),
            _tips.proverTip(),
            _tips.executorTip()
        );
    }

    /**
     * @notice Hook that is called after the specified domain becomes active,
     * i.e. when a Notary is added to the domain, which previously had no active Notaries.
     */
    function _afterDomainBecomesActive(uint32 _domain, address _notary)
        internal
        override(GlobalNotaryRegistry, GlobalNotaryRegistryHarness)
    {
        GlobalNotaryRegistry._afterDomainBecomesActive(_domain, _notary);
    }

    /**
     * @notice Hook that is called after the specified domain becomes inactive,
     * i.e. when the last Notary is removed from the domain.
     */
    function _afterDomainBecomesInactive(uint32 _domain, address _notary)
        internal
        override(GlobalNotaryRegistry, GlobalNotaryRegistryHarness)
    {
        GlobalNotaryRegistry._afterDomainBecomesInactive(_domain, _notary);
    }
}
