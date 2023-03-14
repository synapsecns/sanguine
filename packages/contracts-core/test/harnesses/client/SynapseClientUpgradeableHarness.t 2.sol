// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SynapseClientUpgradeable } from "../../../contracts/client/SynapseClientUpgradeable.sol";
import { ClientHarnessEvents } from "../events/ClientHarnessEvents.sol";

contract SynapseClientUpgradeableHarness is ClientHarnessEvents, SynapseClientUpgradeable {
    uint32 internal immutable optimisticPeriod;

    constructor(
        address _origin,
        address _destination,
        uint32 _optimisticPeriod
    ) SynapseClientUpgradeable(_origin, _destination) {
        optimisticPeriod = _optimisticPeriod;
    }

    function initialize() external initializer {
        __SynapseClient_init();
    }

    function sendMessage(
        uint32 _destination,
        bytes memory _tips,
        bytes memory _message
    ) external payable {
        _send(_destination, _tips, _message);
    }

    function optimisticSeconds() public view override returns (uint32) {
        return optimisticPeriod;
    }

    function _handle(
        uint32 _origin,
        uint32 _nonce,
        bytes memory _message
    ) internal override {
        emit LogClientMessage(_origin, _nonce, _message);
    }
}
