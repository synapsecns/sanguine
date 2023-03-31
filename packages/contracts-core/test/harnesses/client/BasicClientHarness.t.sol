// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { BasicClient } from "../../../contracts/client/BasicClient.sol";
import { BasicClientHarnessEvents } from "../events/BasicClientHarnessEvents.sol";

// solhint-disable no-empty-blocks
contract BasicClientHarness is BasicClientHarnessEvents, BasicClient {
    uint32 internal optimisticPeriod;

    constructor(
        address origin_,
        address destination_,
        uint32 optimisticPeriod_
    ) BasicClient(origin_, destination_) {
        optimisticPeriod = optimisticPeriod_;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testBasicClientHarness() external {}

    function sendMessage(
        uint32 destination_,
        bytes memory tipsPayload,
        bytes memory content
    ) public payable {
        _send(destination_, optimisticSeconds(), tipsPayload, content);
    }

    function optimisticSeconds() public view returns (uint32) {
        return optimisticPeriod;
    }

    function trustedSender(uint32 destination_) public pure override returns (bytes32 sender) {
        sender = bytes32(uint256(destination_));
        // bytes32(0) for destination == 0
    }

    function _handleUnsafe(
        uint32 origin_,
        uint32 nonce,
        uint256 rootSubmittedAt,
        bytes memory content
    ) internal override {
        emit LogBasicClientMessage(origin_, nonce, rootSubmittedAt, content);
    }
}
