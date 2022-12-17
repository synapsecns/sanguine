// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Destination } from "../../contracts/Destination.sol";
import { DestinationHub } from "../../contracts/hubs/DestinationHub.sol";

import { AgentSet } from "../../contracts/libs/AgentSet.sol";
import { Tips } from "../../contracts/libs/Tips.sol";
import { ISystemRouter } from "../../contracts/interfaces/ISystemRouter.sol";

import { AgentRegistryExtended } from "./system/AgentRegistryExtended.t.sol";
import { SystemContractHarness } from "./system/SystemContractHarness.t.sol";
import { DestinationHarnessEvents } from "./events/DestinationHarnessEvents.sol";

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract DestinationHarness is
    DestinationHarnessEvents,
    Destination,
    AgentRegistryExtended,
    SystemContractHarness
{
    using AgentSet for AgentSet.DomainAddressSet;
    using Tips for bytes29;

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) Destination(_domain) {}

    function addRemoteNotary(uint32 _domain, address _notary) external {
        agents[_currentEpoch()].add(_domain, _notary);
    }

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

    function _isIgnoredAgent(uint32 _domain, address _account)
        internal
        view
        override(AgentRegistryExtended, DestinationHub)
        returns (bool)
    {
        return DestinationHub._isIgnoredAgent(_domain, _account);
    }
}
