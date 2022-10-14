// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseClient } from "../../contracts/client/SynapseClient.sol";

// solhint-disable no-empty-blocks
contract SynapseClientHarness is SynapseClient {
    constructor(address _origin, address _destination) SynapseClient(_origin, _destination) {}

    function sendMessage(
        uint32 _destination,
        bytes memory _tips,
        bytes memory _message
    ) external payable {
        _send(_destination, _tips, _message);
    }

    function optimisticSeconds() public pure override returns (uint32) {
        return 0;
    }

    function _handle(
        uint32,
        uint32,
        bytes memory
    ) internal override {}
}
