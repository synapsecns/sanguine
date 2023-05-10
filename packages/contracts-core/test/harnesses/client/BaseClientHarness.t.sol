// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BaseClient} from "../../../contracts/client/BaseClient.sol";
import {RequestLib} from "../../../contracts/libs/Request.sol";
import {TipsLib} from "../../../contracts/libs/Tips.sol";
import {BaseClientHarnessEvents} from "../events/BaseClientHarnessEvents.t.sol";

// solhint-disable no-empty-blocks
contract BaseClientHarness is BaseClient, BaseClientHarnessEvents {
    uint32 private constant OPTIMISTIC_PERIOD = 1 hours;

    constructor(address origin_, address destination_) BaseClient(origin_, destination_) {}

    /// @notice Exposes _sendBaseMessage for testing
    function sendBaseMessage(uint32 destination_, uint256 paddedRequest, bytes memory content) external payable {
        _sendBaseMessage(destination_, RequestLib.wrapPadded(paddedRequest), content);
    }

    /// @inheritdoc BaseClient
    function optimisticPeriod() public pure override returns (uint32) {
        return OPTIMISTIC_PERIOD;
    }

    /// @inheritdoc BaseClient
    function trustedSender(uint32 destination_) public pure override returns (bytes32) {
        // Return different address for different destination
        return bytes32(uint256(destination_));
    }

    /// @inheritdoc BaseClient
    function _receiveBaseMessage(uint32 origin_, uint32 nonce, uint32 version, bytes memory content)
        internal
        override
    {
        emit BaseMessageReceived(msg.value, origin_, nonce, version, content);
    }
}
