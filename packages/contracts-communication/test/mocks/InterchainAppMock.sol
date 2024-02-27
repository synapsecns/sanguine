// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainApp} from "../../contracts/interfaces/IInterchainApp.sol";

// solhint-disable no-empty-blocks
contract InterchainAppMock is IInterchainApp {
    address[] public receivingModules;

    function setReceivingModule(address _receivingModule) external {
        receivingModules.push(_receivingModule);
    }

    function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, bytes calldata message) external payable {}

    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules) {}
}
