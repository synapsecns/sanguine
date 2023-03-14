// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { BasicClient } from "../../../contracts/client/BasicClient.sol";
import { BasicClientHarnessEvents } from "../events/BasicClientHarnessEvents.sol";

contract BasicClientHarness is BasicClientHarnessEvents, BasicClient {
    uint32 internal optimisticPeriod;

    // solhint-disable-next-line no-empty-blocks
    constructor(
        address _origin,
        address _destination,
        uint32 _optimisticPeriod
    ) BasicClient(_origin, _destination) {
        optimisticPeriod = _optimisticPeriod;
    }

    function sendMessage(
        uint32 _destination,
        bytes memory _tips,
        bytes memory _message
    ) public payable {
        _send(_destination, optimisticSeconds(), _tips, _message);
    }

    function optimisticSeconds() public view returns (uint32) {
        return optimisticPeriod;
    }

    function trustedSender(uint32 _destination) public pure override returns (bytes32 sender) {
        sender = bytes32(uint256(_destination));
        // bytes32(0) for _destination == 0
    }

    function _handleUnsafe(
        uint32 _origin,
        uint32 _nonce,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) internal override {
        emit LogBasicClientMessage(_origin, _nonce, _rootSubmittedAt, _message);
    }
}
