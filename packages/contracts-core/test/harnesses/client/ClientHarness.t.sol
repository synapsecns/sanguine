// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Client } from "../../../contracts/client/Client.sol";
import { ClientHarnessEvents } from "../events/ClientHarnessEvents.sol";

contract ClientHarness is ClientHarnessEvents, Client {
    uint32 internal optimisticPeriod;

    // solhint-disable-next-line no-empty-blocks
    constructor(
        address origin,
        address destination,
        uint32 optimisticPeriod
    ) Client(origin, destination) {
        optimisticPeriod = optimisticPeriod;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testClientHarness() external {}

    function sendMessage(
        uint32 destination,
        bytes memory tips,
        bytes memory message
    ) public payable {
        _send(destination, tips, message);
    }

    function optimisticSeconds() public view override returns (uint32) {
        return optimisticPeriod;
    }

    function trustedSender(uint32 destination) public pure override returns (bytes32 sender) {
        sender = bytes32(uint256(destination));
        // bytes32(0) for destination == 0
    }

    function _handle(
        uint32 origin,
        uint32 nonce,
        bytes memory message
    ) internal override {
        emit LogClientMessage(origin, nonce, message);
    }
}
