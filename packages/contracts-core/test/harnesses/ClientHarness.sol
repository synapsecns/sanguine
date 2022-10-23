// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Client } from "../../contracts/client/Client.sol";

contract ClientHarness is Client {
    uint32 internal optimisticPeriod;

    event LogMessage(uint32, uint32, bytes);

    // solhint-disable-next-line no-empty-blocks
    constructor(
        address _origin,
        address _destination,
        uint32 _optimisticPeriod
    ) Client(_origin, _destination) {
        optimisticPeriod = _optimisticPeriod;
    }

    function sendMessage(
        uint32 _destination,
        bytes memory _tips,
        bytes memory _message
    ) public payable {
        _send(_destination, _tips, _message);
    }

    function optimisticSeconds() public view override returns (uint32) {
        return optimisticPeriod;
    }

    function trustedSender(uint32 _destination) public pure override returns (bytes32 sender) {
        if (_destination != 0) {
            sender = keccak256(abi.encode(_destination));
        }
        // bytes32(0) for _destination == 0
    }

    function _handle(
        uint32 _origin,
        uint32 _nonce,
        bytes memory _message
    ) internal override {
        emit LogMessage(_origin, _nonce, _message);
    }
}
