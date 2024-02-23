// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainApp} from "../../contracts/interfaces/IInterchainApp.sol";
import {AppConfigV1} from "../../contracts/libs/AppConfig.sol";

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

    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules) {}

    function send(bytes32 receiver, uint256 dstChainId, bytes calldata message) external payable virtual override {}

    function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes calldata message) external {}
}
