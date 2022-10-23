// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SynapseClientUpgradeable } from "../../contracts/client/SynapseClientUpgradeable.sol";

// solhint-disable no-empty-blocks
contract SynapseClientUpgradeableHarness is SynapseClientUpgradeable {
    constructor(address _origin, address _destination)
        SynapseClientUpgradeable(_origin, _destination)
    {}

    function initialize() external initializer {
        __SynapseClient_init();
    }

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
