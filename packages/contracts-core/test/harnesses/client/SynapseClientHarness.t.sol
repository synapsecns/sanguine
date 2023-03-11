// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SynapseClient } from "../../../contracts/client/SynapseClient.sol";
import { ClientHarnessEvents } from "../events/ClientHarnessEvents.sol";

contract SynapseClientHarness is ClientHarnessEvents, SynapseClient {
    uint32 internal optimisticPeriod;

    constructor(
        address _origin,
        address _destination,
        uint32 _optimisticPeriod
    ) SynapseClient(_origin, _destination) {
        optimisticPeriod = _optimisticPeriod;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseClientHarness() external {}

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
