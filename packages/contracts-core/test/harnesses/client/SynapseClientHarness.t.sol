// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SynapseClient } from "../../../contracts/client/SynapseClient.sol";
import { ClientHarnessEvents } from "../events/ClientHarnessEvents.sol";

// solhint-disable no-empty-blocks
contract SynapseClientHarness is ClientHarnessEvents, SynapseClient {
    uint32 internal optimisticPeriod;

    constructor(
        address origin_,
        address destination_,
        uint32 optimisticPeriod_
    ) SynapseClient(origin_, destination_) {
        optimisticPeriod = optimisticPeriod_;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseClientHarness() external {}

    function sendMessage(
        uint32 destination_,
        bytes memory tipsPayload,
        bytes memory content
    ) external payable {
        _send(destination_, tipsPayload, content);
    }

    function optimisticSeconds() public view override returns (uint32) {
        return optimisticPeriod;
    }

    function _handle(
        uint32 origin_,
        uint32 nonce,
        bytes memory content
    ) internal override {
        emit LogClientMessage(origin_, nonce, content);
    }
}
