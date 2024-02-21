// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainApp} from "../../contracts/interfaces/IInterchainApp.sol";

contract InterchainAppMock is IInterchainApp {
    address[] public receivingModules;

    function setReceivingModule(address _receivingModule) public {
        receivingModules.push(_receivingModule);
    }

    function setAppConfig(
        uint64[] memory chainIDs,
        address[] memory linkedIApps,
        address[] memory sendingModules,
        address[] memory _receivingModules,
        uint256 requiredResponses,
        uint64 optimisticTimePeriod
    )
        external
        virtual
        override
    {}

    function getLinkedIApp(uint64 chainID) external view virtual override returns (address) {}

    function getSendingModules() external view virtual override returns (address[] memory) {}

    function getReceivingModules() public view override returns (address[] memory) {
        return receivingModules;
    }

    function getRequiredResponses() public pure override returns (uint256) {
        return 1;
    }

    function getOptimisticTimePeriod() public pure override returns (uint64) {
        return 0;
    }

    function send(bytes32 receiver, uint256 dstChainId, bytes calldata message) external payable virtual override {}

    function appReceive() external virtual override {}
}
