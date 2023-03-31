// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SynapseClient } from "../../../contracts/client/SynapseClient.sol";
import { ClientHarnessEvents } from "../events/ClientHarnessEvents.sol";

contract SynapseClientHarness is ClientHarnessEvents, SynapseClient {
    uint32 internal optimisticPeriod;

    constructor(
        address origin,
        address destination,
        uint32 optimisticPeriod
    ) SynapseClient(origin, destination) {
        optimisticPeriod = optimisticPeriod;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseClientHarness() external {}

    function sendMessage(
        uint32 destination,
        bytes memory tips,
        bytes memory message
    ) external payable {
        _send(destination, tips, message);
    }

    function optimisticSeconds() public view override returns (uint32) {
        return optimisticPeriod;
    }

    function _handle(
        uint32 origin,
        uint32 nonce,
        bytes memory message
    ) internal override {
        emit LogClientMessage(origin, nonce, message);
    }
}
