// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { BasicClient } from "../../../contracts/client/BasicClient.sol";
import { BasicClientHarnessEvents } from "../events/BasicClientHarnessEvents.sol";

contract BasicClientHarness is BasicClientHarnessEvents, BasicClient {
    uint32 internal optimisticPeriod;

    // solhint-disable-next-line no-empty-blocks
    constructor(
        address origin,
        address destination,
        uint32 _optimisticPeriod
    ) BasicClient(origin, destination) {
        optimisticPeriod = _optimisticPeriod;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testBasicClientHarness() external {}

    function sendMessage(
        uint32 destination,
        bytes memory tips,
        bytes memory message
    ) public payable {
        _send(destination, optimisticSeconds(), tips, message);
    }

    function optimisticSeconds() public view returns (uint32) {
        return optimisticPeriod;
    }

    function trustedSender(uint32 destination) public pure override returns (bytes32 sender) {
        sender = bytes32(uint256(destination));
        // bytes32(0) for destination == 0
    }

    function _handleUnsafe(
        uint32 origin,
        uint32 nonce,
        uint256 rootSubmittedAt,
        bytes memory message
    ) internal override {
        emit LogBasicClientMessage(origin, nonce, rootSubmittedAt, message);
    }
}
