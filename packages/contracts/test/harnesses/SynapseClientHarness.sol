// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseClient } from "../../contracts/client/SynapseClient.sol";

contract SynapseClientHarness is SynapseClient {
    constructor(address _home, address _replicaManager) SynapseClient(_home, _replicaManager) {}

    function _handle(
        uint32,
        uint32,
        bytes32,
        bytes memory
    ) internal override {}

    function optimisticSeconds() public pure override returns (uint32) {
        return 0;
    }

    function send(
        uint32 _destination,
        bytes memory _tips,
        bytes memory _message
    ) external payable {
        _send(_destination, _tips, _message);
    }
}
