// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseClientUpgradeable } from "../../contracts/client/SynapseClientUpgradeable.sol";

contract SynapseClientUpgradeableHarness is SynapseClientUpgradeable {
    constructor(address _home, address _replicaManager)
        SynapseClientUpgradeable(_home, _replicaManager)
    {}

    function initialize() external initializer {
        __SynapseClient_init();
    }

    function _handle(
        uint32,
        uint32,
        bytes32,
        bytes memory
    ) internal override {}

    function optimisticSeconds() public pure override returns (uint32) {
        return 0;
    }

    function send(uint32 _destination, bytes memory _message) external {
        _send(_destination, _message);
    }
}
