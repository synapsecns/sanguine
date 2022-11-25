// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Destination } from "../../contracts/Destination.sol";
import { AbstractNotaryRegistry } from "../../contracts/registry/AbstractNotaryRegistry.sol";

import { Tips } from "../../contracts/libs/Tips.sol";
import { ISystemRouter } from "../../contracts/interfaces/ISystemRouter.sol";

import { GuardRegistryHarness } from "./registry/GuardRegistryHarness.t.sol";
import { SystemContractHarness } from "./system/SystemContractHarness.t.sol";
import { DestinationHarnessEvents } from "./events/DestinationHarnessEvents.sol";

contract DestinationHarness is
    DestinationHarnessEvents,
    Destination,
    SystemContractHarness,
    GuardRegistryHarness
{
    using Tips for bytes29;

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) Destination(_domain) {}

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

    function addNotary(uint32 _domain, address _notary) public returns (bool) {
        return _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) public returns (bool) {
        return _removeNotary(_domain, _notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }

    function _storeTips(bytes29 _tips) internal override {
        emit LogTips(
            _tips.notaryTip(),
            _tips.broadcasterTip(),
            _tips.proverTip(),
            _tips.executorTip()
        );
    }
}
