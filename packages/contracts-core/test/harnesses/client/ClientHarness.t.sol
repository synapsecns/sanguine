// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Client } from "../../../contracts/client/Client.sol";
import { ClientHarnessEvents } from "../events/ClientHarnessEvents.sol";

// solhint-disable no-empty-blocks
contract ClientHarness is ClientHarnessEvents, Client {
    uint32 internal optimisticPeriod;

    // solhint-disable-next-line no-empty-blocks
    constructor(
        address origin_,
        address destination_,
        uint32 optimisticPeriod_
    ) Client(origin_, destination_) {
        optimisticPeriod = optimisticPeriod_;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testClientHarness() external {}

    function sendMessage(
        uint32 destination_,
        bytes memory tipsPayload,
        bytes memory content
    ) public payable {
        _send(destination_, tipsPayload, content);
    }

    function optimisticSeconds() public view override returns (uint32) {
        return optimisticPeriod;
    }

    function trustedSender(uint32 destination_) public pure override returns (bytes32 sender) {
        sender = bytes32(uint256(destination_));
        // bytes32(0) for destination == 0
    }

    function _handle(
        uint32 origin_,
        uint32 nonce,
        bytes memory content
    ) internal override {
        emit LogClientMessage(origin_, nonce, content);
    }
}
